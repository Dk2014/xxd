<?php
class GameDBUpgrader {
	private $_dbcfg;
	private $_db;
	private $_dir;
	private $_needBackup;
	private $_singleFileMode;

	private $_gsid;

	private $_IDTables = array();
	private $_selectdbs = array(/*array(id, free, db)*/);

	const _MAX_SELECTDB_NUM = 5;
	const _php_dbdump_header = "<?php /* dump file */\n";

	public function __construct($dir, $backup, $host, $user, $pass, $name, $port, $vendor='default') {
		$this->_dir   = $dir;
		$this->_dbcfg = array('host'=>$host, 'user'=>$user, 'pass'=>$pass, 'name'=>$name, 'port'=>$port);
		$this->_db    = $this->newDB();
		$this->_needBackup = $backup;
		$this->_singleFileMode = isset($GLOBALS['singleFileLogArray']);
		$this->_vendor = $vendor;
	}

	public function __destruct() {
		$this->_db->close();
		foreach ($this->_selectdbs as $dbinfo) {
			$dbinfo['db']->close();
		}
	}
	public function AddSQL($sql, $needReviewInLog = false, $exclude_vendors=array()) {
		if ($needReviewInLog){

			echo "[DEBUG]" . trim($sql, "\r\n") . "\n";

		}
		if(in_array($this->_vendor, $exclude_vendors)) {
			return;
		}

		if ($this->_db->multi_query($sql)) {
			do {
				if (!$this->_db->more_results()) {
					break;
				}
			} while ($this->_db->next_result());
		}

		if ($this->_db->errno) {
			echo("\nupgrade database failed ({$this->_db->errno}) {$this->_db->error} gsid: {$this->_gsid} db: {$this->_dbcfg['name']}\n\n");
			$this->_db->rollback();
			die();
		}		
	}

	public function getServerInfo($db) {
		$result = $db->query("select round((id>>32)/10) as gsid from player limit 1;");
		if ($result) {
			$row    = $result->fetch_array(MYSQLI_ASSOC);
			$result->close();
			if ($row == false) {
				return 0;
			} else {
				return $row['gsid'];
			}
		}
		return 0;
	}

	public function getOnlineVersion() {
		$result = $this->_db->query("SELECT `version` FROM `db_version`;");
		$row    = $result->fetch_array(MYSQLI_ASSOC);
		$result->close();

		if ($row == false) {
			$this->_db->query("INSERT INTO `db_version` VALUES (0);");
			return 0;
		} else {
			return $row['version'];
		}
	}

	public function NewQuery($sql) {
		if (count($this->_selectdbs) > self::_MAX_SELECTDB_NUM) {
			die("超过db连接限制");
		}

		$db_idx = null;
		foreach ($this->_selectdbs as $key=>&$dbinfo) {
			if ($dbinfo['free'] == true) {
				$dbinfo['free'] = false;
				$db_idx = $key;
				break;
			}
		}

		if ($db_idx === null) {
			$db_id = count($this->_selectdbs)+1;
			$this->_selectdbs[] = array('id'=>$db_id, 'free'=>false, 'db'=>$this->newDB());
			$db_idx = count($this->_selectdbs) - 1;
		}

		return new DBQuery($this->_selectdbs[$db_idx], $sql);
	}

	public function DropQuery(&$dbQuery) {
		$dbQuery->Drop();
	}

	public function GetAutoID($pid, $tableName ) {
		$fixid = ($pid>>32)<<32;
		$idkey = "{$fixid}";

		if (!isset($this->_IDTables[$tableName]) || !isset($this->_IDTables[$tableName][$idkey])){
			if(!isset($this->_IDTables[$tableName])) {
				$this->_IDTables[$tableName] = array();
			}
			$maxQuery = $this->NewQuery("SELECT MAX(`id`) as id FROM `{$tableName}` WHERE (pid >> 32) = ".($pid>>32));
			$max = $maxQuery->GoNext();
			$this->DropQuery($maxQuery);
			$this->_IDTables[$tableName][$idkey] = ($max['id']==NULL) ? $fixid : $max['id'];
		}

		$this->_IDTables[$tableName][$idkey] += 1;
		return $this->_IDTables[$tableName][$idkey];
	}

	// 使用此接口限制php >= 5.4.0
	public function ExportSingleUpgradeFile($file) {
		$target_vendor = exec("hg branch");
		$StripPHPTag = function($code) {
			$code = str_replace("<?php", "", $code);
			$code = str_replace("<?PHP", "", $code);
			$code = str_replace("<?", "", $code);
			$code = str_replace("?>", "", $code);
			return $code;
		};

		$fp = fopen($file, 'w');
		fwrite($fp, "<?php\n\n");

		$code = <<<FNCODE
error_reporting(E_ALL);
set_time_limit(0);
ini_set('memory_limit', '256M');

\$host = \$argv[1];
\$user = \$argv[2];
\$pass = \$argv[3];
\$name = \$argv[4];
\$port = \$argv[5];

\$target_vendor = '{$target_vendor}';

\$singleFileLogArray = array();
\$upgrader = new GameDBUpgrader(null, false, \$host, \$user, \$pass, \$name, \$port, \$target_vendor);

FNCODE;
		fwrite($fp, $code);

		$changesets  = $this->getChanges();
		foreach ($changesets as $version => $file) {
			if (!is_file($file)) {
				die("$file is not file");
			}

			$code = file_get_contents($file);
			$this->checkUpgradeFile($file, $code);

			$code = trim($StripPHPTag($code));
			$head = ($this->checkFileIsDBDump($file)) ? "\$this->_db->autocommit(1);\n" : "";
			$tail = ($this->checkFileIsDBDump($file)) ? "return false;\n" : "return true;\n";
			$fnCode =<<<FNCODE

\$singleFileLogArray[] = $version;
\$updatelog_$version = function(\$isNewDatabase) {
	$head
	$code
	$tail
};
\$do_updatelog_$version = Closure::bind(\$updatelog_$version, \$upgrader, 'GameDBUpgrader');

FNCODE;
			fwrite($fp, $fnCode);
		}

		$code = trim($StripPHPTag(file_get_contents(__FILE__)));
		fwrite($fp, $code);
		fwrite($fp, "\n\$upgrader->UpgraderBySingleFile();\n");
		fwrite($fp, "\n\n?>");
		fclose($fp);
	}

	public function UpgraderBySingleFile() {
		if (!$this->_singleFileMode) {
			die("不能是单文件模式进行升级");
		}

		$lastVersion = $this->gamedb_prepare();
		$isNewDatabase = $lastVersion == 0;

		foreach ($GLOBALS['singleFileLogArray'] as $version) {
			if ($lastVersion >= $version) {
				continue;
			}

			echo "\nupgrade: " . $version . "\n";
			$this->_db->autocommit(0);

			// 重置ID自增长ID
			$this->_IDTables = array();
			$fn = "do_updatelog_$version";
			if (!$GLOBALS[$fn]($isNewDatabase)) {
				$this->AddSQL("UPDATE `db_version` SET `version` = {$version};");
				echo " ........................ [done]\n";
				$this->_db->autocommit(0);
				continue;
			}

			$this->AddSQL("UPDATE `db_version` SET `version` = " . $version . ";");

			if ($this->_db->commit()) {
				echo " ........................ [done]\n";
			} else {
				echo("\nupgrade database failed ({$this->_db->errno}) {$this->_db->error} gsid: {$this->_gsid} db: {$this->_dbcfg['name']}\n\n");
				$this->_db->rollback();
				die();
			}
		}
		$this->_db->autocommit(1);
		echo "\nupgrade complete {$this->_gsid}\n";
	}

	public function UpgradeByLogFile() {
		if ($this->_singleFileMode) {
			die("这是多文件升级模式");
		}
		$lastVersion = $this->gamedb_prepare();
		$isNewDatabase = $lastVersion == 0;

		$changesets  = $this->getChanges();

		if ($this->_needBackup) {
			foreach ($changesets as $version => $file) {
				if ($version > $lastVersion) {
					if (!$this->dumpDB()) {
						die("can't dump db.");
					}
					break;
				}
			}
		}

		$this->_db->autocommit(0);

		foreach ($changesets as $version => $file) {
			if ($lastVersion >= $version) {
				continue;
			}

			if (!is_file($file)) {
				die("$file is not file");
			}

			echo "\nupgrade: " . $file;

			// 执行导出的数据
			if ($this->checkFileIsDBDump($file)) {
				$this->_db->autocommit(1);
				require_once ($file);
				$this->AddSQL("UPDATE `db_version` SET `version` = {$version};");
				echo " ........................ [done]\n";
				$this->_db->autocommit(0);
				continue;				
			}

			// 重置ID自增长ID
			$this->_IDTables = array();
			require_once ($file);

			$this->AddSQL("UPDATE `db_version` SET `version` = " . $version . ";");

			if ($this->_db->commit()) {
				echo " ........................ [done]\n";
			} else {
				echo("\nupgrade database failed ({$this->_db->errno}) {$this->_db->error} gsid: {$this->_gsid} db: {$this->_dbcfg['name']}\n\n");
				$this->_db->rollback();
				die();
			}
		}

		$this->_db->autocommit(1);
		echo "\nupgrade complete {$this->_gsid}\n";
	}

	private function checkUpgradeFile($file, &$code) {
		if ($this->checkFileIsDBDump($file) || strpos($code, "\$this->NewQuery")===false) { // dump file 和没有NewQuery的file不检查
			return;
		}

		$TagDDL = array('ALTER', 'CREATE', 'DECLARE', 'TRUNCATE');
		$TagDML = array('INSERT', 'UPDATE','DELETE');

		$haveDDL = false;
		$haveDML = false;

		foreach ($TagDDL as $tag) {
			if (stripos($code, $tag)!==false) {
				$haveDDL = true;
				break;
			}
		}

		foreach ($TagDML as $tag) {
			if (stripos($code, $tag)!==false) {
				$haveDML = true;
				break;
			}
		}

		if ($haveDDL && $haveDML) {
			die("\n !!!!!!!! ERROR: DDL 不能和 DML 同时并存在同一个更新文件中.".$file."\n\n\n");
		}
	}

	private function checkFileIsDBDump($file) {
		$f = fopen($file, 'r');
		$is_dump = (fgets($f) == self::_php_dbdump_header);
		fclose($f);
		return $is_dump;	
	}

	private function getChanges() {
		if ($this->_singleFileMode) {
			die("getChanges()是多文件升级模式调用接口");
		}
		$changes = array();
		foreach (glob($this->_dir . "/*.php") as $file) {
			$name              = basename($file, '.php');
			$version           = (int) str_replace('-', '', $name);
			$changes[$version] = $file;
		}

		asort($changes);
		return $changes;
	}

	private function dumpDB() {
		if ($this->_singleFileMode) {
			die("dumpDB()是多文件升级模式调用接口");
		}
		$cmd = "mysqldump -h{$this->_dbcfg['host']} -P{$this->_dbcfg['port']} -u{$this->_dbcfg['user']} -p{$this->_dbcfg['pass']} --skip-add-locks {$this->_dbcfg['name']}";
		ob_start();
		system($cmd, $ret);
		$data = ob_get_contents();
		ob_end_clean();
		if ($ret == 0) {
			if (!is_dir($this->_dir . "/db_backup")) {
				mkdir($this->_dir . "/db_backup");
			}

			$version = $this->getOnlineVersion();
			$sqlfile = $version . ".sql";
			$i       = 0;
			while (file_exists($this->_dir . "/db_backup/$sqlfile")) {
				$i++;
				$sqlfile = $version . "_$i.sql";
			}
			file_put_contents($this->_dir . "/db_backup/$sqlfile", $data);
		}
		return $ret == 0;
	}

	private function newDB() {
		$db = new mysqli($this->_dbcfg['host'], $this->_dbcfg['user'], $this->_dbcfg['pass'], $this->_dbcfg['name'], $this->_dbcfg['port']);
		if ($db->connect_error) {
			die('连接数据库失败：Error ' . $db->connect_errno . ', ' . $db->connect_error);
		}
		$db->query("set names utf8mb4");
		return $db;
	}

	private function gamedb_prepare() {
		$db = $this->newDB();
		$db_name = $this->_dbcfg['name'];

		if ($db->connect_error) {
			die('open information_schema failed (' . $db->connect_errno . ') ' . $db->connect_error);
		}

		$db->query("USE `information_schema`");
		// 检查数据库是否存在
		$sql = "SELECT `SCHEMA_NAME` FROM `SCHEMATA` WHERE `SCHEMA_NAME` = '".$db_name."'";
		$result = $db->query($sql);
		$need_init_db = $result->fetch_array(MYSQLI_ASSOC) == false;
		$result->close();

		// 检查版本表是否存在
		$sql = "SELECT `TABLE_NAME` FROM `TABLES` WHERE `TABLE_NAME` = 'db_version' AND `TABLE_SCHEMA` = '".$db_name."'";
		$result = $db->query($sql);
		$need_init_table = $result->fetch_array(MYSQLI_ASSOC) == false;
		$result->close();

		// 初始化数据库
		if ($need_init_db) {
			$sql = "CREATE DATABASE `".$db_name."` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';";
			if ($db->query($sql) === false)
				die("prepare database fialed (" . $db->errno . ') ' . $db->error);
		}

		$db->query("USE `".$db_name."`");

		// 初始化版本表
		if ($need_init_table) {
			$sql = "CREATE TABLE `db_version` ( `version` INT ) ENGINE 'InnoDb' CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci'";
			if ($db->multi_query($sql) === false)
				die("prepare database fialed (" . $db->errno . ') ' . $db->error);
			$db->close();
			$db = $this->newDB();
		}

		// 获取版本号
		$sql = "SELECT `version` FROM `db_version`;";
		$result = $db->query($sql);
		$row = $result->fetch_array(MYSQLI_ASSOC);
		$result->close();

		if ($row == false) {
			$db->query("INSERT INTO `db_version` VALUES (0);");
			$version = 0;
		} else {
			$version = $row['version'];
		}

		$this->_gsid = $this->getServerInfo($db);

		$db->close();

		return $version;
	}
}

class DBQuery {
	public $db_info;
	private $_query_result = NULL;

	public function __construct(&$db_info, $sql) {
		$this->db_info   = &$db_info;
		$this->_query_result = $this->db_info['db']->query($sql, MYSQLI_USE_RESULT);
		if (!$this->_query_result) {
			die("\nQuery ($sql) Error (" . $this->db_info['db']->errno . "):\n" . $this->db_info['db']->error . "\n");
		}
	}

	public function __destruct() {
		if ($this->_query_result) {
			$this->Drop();
		}
	}

	public function GoNext() {
		return $this->_query_result->fetch_array(MYSQLI_ASSOC);
	}

	public function Have() {
		return count($this->_query_result->fetch_row()) > 0;
	}

	public function Drop() {
		$this->db_info['free'] = true;
		$this->_query_result->close();
		$this->_query_result = NULL;
	}
}
?>
