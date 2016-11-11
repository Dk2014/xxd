<?php

//require_once "lang_rawstr.php";

//$lang_rawstr_php_dq = new lang_rawstr("php_dq");

require_once "nation_tool.php";

// 导出文件的头部
$dump_file_head = "<?php /* dump file */\n";

if (!isset($target)) {
	//
	// 命令行参数处理
	//
	$command = $argv[1];  // 命令
	$target = $argv[2];   // 目标数据库

	if ($command != 'install') {
		require_once dirname(__FILE__)."/config.php";
	} else {
		$command = 'upgrade';
		$target  = $argv[2];

		$db_config = array(
			$target => array(
				'host' => $argv[3],
				'user' => $argv[4],
				'pass' => $argv[5],
				'name' => $argv[6],
				'port' => '3306',
			)
		);
	}

	$db_config['changes'] = dirname(__FILE__).'/servers/'.$db_config['local_game']['server'];

	if ($command == "invoke") {
		$plugin_name = "./plugins/".$argv[3]."/main.php";
	
		if (!is_file($plugin_name)) {
			echo "调用不存在的插件\n";
			exit;
		}
	
		require_once $plugin_name;
	
		plugin_main();
	}
	else if ($command == "upgrade") {
		db_upgrade();
	}
	else if ($command == "export") {
		exec('hg pull -u -f');
		exec('sh zupgrade');
		db_export();
	}
	else if ($command == "clean") {
		db_clean();
	}
	else if ($command == "check") {
		db_check($argv[3]);
	}
	else {
		echo "不支持的操作\n";
	}
}

// 保存文件，并记录文件名，用于自动删除旧文件和避免重复生成
function save_code($file, $code) {
	global $saved_files;
	if (!is_file($file) || file_get_contents($file) != $code) {
		file_put_contents($file, $code);
	}
	$saved_files[] = realpath($file);
}

// 删除旧文件
function remove_old_files($dirPath){
	global $saved_files;

	if(file_exists($dirPath)){
		$dirFiles=scandir($dirPath);
		foreach ($dirFiles as $dirFile) {
			$dirFilePath = $dirPath."/".$dirFile;
			if (filetype($dirFilePath) == "file" ) {
				if (!in_array(realpath($dirFilePath), $saved_files)) {
					unlink($dirFilePath);
				}
			}
		}
	}
}

//
// 连接当前目标数据库
//
function db_connect()
{
	global $target, $db_config;
	
	$host = $db_config[$target]['host'];
	$user = $db_config[$target]['user'];
	$pass = $db_config[$target]['pass'];
	$name = $db_config[$target]['name'];
	$port = $db_config[$target]['port'];
	
	$db = new mysqli($host, $user, $pass, $name, $port);

	if ($db->connect_error) {
		die('连接数据库失败：Error '.$db->connect_errno.', '.$db->connect_error);
	}
	
	$db->query("set names utf8mb4");
	
	return $db;
}

//
// 预备好数据库
//
function db_prepare() {
	global $target, $db_config;
	
	$db_host = $db_config[$target]['host'];
	$db_user = $db_config[$target]['user'];
	$db_pass = $db_config[$target]['pass'];
	$db_name = $db_config[$target]['name'];
	$db_port = $db_config[$target]['port'];
	
	$db = new mysqli($db_host, $db_user, $db_pass, 'information_schema', $db_port);
	if ($db->connect_error) {
		die('open information_schema failed (' . $db->connect_errno . ') ' . $db->connect_error);
	}
	
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
		$db = new mysqli($db_host, $db_user, $db_pass, $db_name, $db_port);
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

	$db->close();
	
	return $version;
}

function get_changes($db_version = 0) {
	global $dump_file_head, $target, $db_config;

	$dir = './servers/'.$db_config[$target]['server'];

	// 按版本号差异获取变更文件列表
	$changes = array();
	$last_version = 0;
	$last_version_name = '';
	$last_export_version = 0;

	if ($handle = opendir($dir)) {
		while (FALSE !== ($file = readdir($handle))) {
			if ($file == "." || $file == ".." || is_dir($file))
				continue;
			
			if (strrpos($file, ".php") == 13) {
				if (strlen($file) != 17)
					continue;
				
				$name = substr($file, 0, -4);
				$version = (int)str_replace('-', '', $name);
				$is_dump = false;
			}
			else {
				continue;
			}

			$file_handler = fopen($dir.'/'.$file, 'r');
			$is_dump = (fgets($file_handler) == $dump_file_head);
			fclose($file_handler);

			if ($version > $last_version) {
				$last_version = $version;
				$last_version_name = $name;
			}
			
			if ($version > $db_version) {
				if ($is_dump && $last_export_version <= $version)
					$last_export_version = $version;
				$changes[$version] = array('name' => $name, 'dir' => $dir, 'is_dump' => $is_dump);
			}
		}
	}
	
    asort($changes);

    return array(
    	'changes' => $changes,
    	'last_version' => $last_version,
		'last_version_name' => $last_version_name,
    	'last_export_version' => $last_export_version,
    );
}

//
// 升级数据库
//
function db_upgrade() {
	global $target, $db_config;
	
	$db_version = db_prepare();
	$renew = $db_version == 0;
	
	$db = db_connect();

	$changesInfo = get_changes($db_version);
	
	$changes = $changesInfo['changes'];
	$last_export_version = $changesInfo['last_export_version'];
    
	$main_stime = microtime(true);

	// 效验数据库脚本是否符合规范 不符合就拒绝执行
	if ($db_version != 0 ){
		db_upgrade_validate($changes);  // 新库不做效验
	}
	

	$db->autocommit(0); // 关闭mysql的自动提交
	$latestVersion=0;   // 最后版本号 注意 不是mysql的版本 是脚本的版本
	foreach (array_keys($changes) as $version) {
		if ($db_version == 0 || $renew) {
			if ($version < $last_export_version)
				continue;
		} else if ($changes[$version]['is_dump'] && $last_export_version > $version) {
			continue;
		}

		echo "\napply change: ".$changes[$version]['name'].".php";
		if ($changes[$version]['is_dump']){ // 数据导出脚本自动提交
			$db->autocommit(1);
			require_once("./".$changes[$version]['dir']."/".$changes[$version]['name'].".php");
			$sql = "UPDATE `db_version` SET `version` = ".$version . ";";
			db_execute($db, $sql);
			echo " ........................ [done]";
			$db->autocommit(0);
			continue;
		}

		$success=TRUE; // 变量:本脚本的语句是否全部执行成功
		require_once("./".$changes[$version]['dir']."/".$changes[$version]['name'].".php");

		if ($db->errno) {
			$success=FALSE;
			echo("\n 报错原因: ".$db->error."\n");
		}

		if ($success){ // 如果脚本内语句全部执行成功，才把脚本版本号写入数据库
			$sql = "UPDATE `db_version` SET `version` = ".$version . ";" ;
			db_execute($db, $sql);
			if ($db->errno){
				$success=FALSE;
				echo("\n报错原因: ".$db->error."\n");
			}
		}

		if($success){
			$db->commit(); // 提交
			echo " ........................ [done]";
    	}else{
    		$db->rollback(); // 回滚
    		echo "=== 数据已回滚 ===";
    		die("\nupgrade database fialed (" . $db->errno . ') ' . $db->error);
    	}
	}
    $db->autocommit(1);  // 开启mysql自动提交
	
	$main_etime = microtime(true);
	echo "\n\napply changes complete in ".round($main_etime - $main_stime, 2)."s\n";
}

function db_execute($db, $sql) {
	if ($db->multi_query($sql)) {
	    do { 
			if (!$db->more_results()) {
				break;
			}
		} while ($db->next_result());	
	}
}

function db_query($db, $sql) {
	$result = $db->query($sql, MYSQLI_USE_RESULT);

	if (!$result)
		die("\nQuery ($sql) Error (" . $db->errno . "):\n" . $db->error . "\n");

	$rows = array();

	while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
		$rows[] = $row;
	}

	$result->close();

	return $rows;
}

function db_query_array($db, $sql) {
	$result = $db->query($sql, MYSQLI_USE_RESULT);

	if (!$result)
		die("\nQuery ($sql) Error (" . $db->errno . "):\n" . $db->error . "\n");

	$rows = array();

	while ($row = $result->fetch_array(MYSQLI_NUM)) {
		$rows[] = $row;
	}

	$result->close();

	return $rows;
}

//
// 获取数据库表结构信息
//
// 返回格式：
// array(
//   array(
//	 'name' => '表名',
//	 'desc' => '注释',
//   'index' => '序号',
//	 'cols' => array(
//	   array(
//		 'name' => '字段名',
//		 'type' => '类型',
//		 'desc' => '注释'
//	   ),
//	 ...)
//   )
// ...)
//
function db_tables($only_player_table = false, $only_log_table = false)
{
	global $target, $db_config;
	
	$db_name = $db_config[$target]['name'];
	
	$db = db_connect();
	$db->query("use information_schema");
	
	$tables = array();
	$tables_index = array();
	$index = 0 ;
	
	$sql = "select TABLE_NAME, TABLE_COMMENT from TABLES where TABLE_SCHEMA = '$db_name'";
	
	if ($only_player_table)
		$sql .= " and (TABLE_NAME like 'player%' or TABLE_NAME like 'global_%') and TABLE_NAME not like '%_log'";

	if ($only_log_table)
		$sql .= " and (TABLE_NAME like '%_log')";
	
	if ($result = $db->query($sql)) {
		while ($row = $result->fetch_assoc()) {
			$tables[] = array(
				'name' => $row['TABLE_NAME'],
				'desc' => $row['TABLE_COMMENT'],
				'cols' => array()
			);

			$tables_index[ $row['TABLE_NAME'] ] = $index;
			$index++;
		}
		$result->close();
	}
	else {
		die("获取数据库表信息时失败：Error ".$db->connect_errno.', '.$db->connect_error);
	}
	
	$sql = "select TABLE_NAME, COLUMN_NAME, DATA_TYPE, COLUMN_DEFAULT, COLUMN_COMMENT, COLUMN_KEY from COLUMNS where `TABLE_SCHEMA` = '".$db_name."' order by ORDINAL_POSITION asc";
		
	if ($result = $db->query($sql)) {
		while ($row = $result->fetch_assoc()) {
			$table_name = $row['TABLE_NAME'];
			if (!isset($tables_index[$table_name])){
				continue;
			}

			$i =  $tables_index[$table_name];

			$col = array(
				'name' => $row['COLUMN_NAME'],
				'type' => $row['DATA_TYPE'],
				'desc' => $row['COLUMN_COMMENT']
			);
			$tables[$i]['cols'][] = $col;
			if ($row['COLUMN_KEY'] == 'PRI') {
				$tables[$i]['pri'] = $col;
			}
		}
		$result->close();
	}
	
	$db->close();
	
	return $tables;
}

function db_config_tables()
{
	global $target, $db_config;
	
	$db_name = $db_config[$target]['name'];
	
	$db = db_connect();
	$db->query("use information_schema");
	
	$tables = array();
	
	$sql = "select TABLE_NAME, TABLE_COMMENT from TABLES where TABLE_SCHEMA = '$db_name'";
	
	$sql .= " and TABLE_NAME not like 'player%' and TABLE_NAME not like 'global_%' and TABLE_NAME not like '%_log'";

	if ($result = $db->query($sql)) {
		while ($row = $result->fetch_assoc()) {
			$tables[] = array(
				'name' => $row['TABLE_NAME'],
				'desc' => $row['TABLE_COMMENT'],
				'cols' => array()
			);
		}
		$result->close();
	}
	else {
		die("获取数据库表信息时失败：Error ".$db->connect_errno.', '.$db->connect_error);
	}
	
	$tables_count = count($tables);
	
	for ($i = 0; $i < $tables_count; $i ++) {
		$sql = "select COLUMN_NAME, DATA_TYPE, COLUMN_DEFAULT, COLUMN_COMMENT, COLUMN_KEY from COLUMNS where TABLE_NAME = '{$tables[$i]['name']}' AND `TABLE_SCHEMA` = '".$db_name."' order by ORDINAL_POSITION asc";
		
		if ($result = $db->query($sql)) {
			while ($row = $result->fetch_assoc()) {
				$col = array(
					'name' => $row['COLUMN_NAME'],
					'type' => $row['DATA_TYPE'],
					'desc' => $row['COLUMN_COMMENT']
				);
				$tables[$i]['cols'][] = $col;
				if ($row['COLUMN_KEY'] == 'PRI') {
					$tables[$i]['pri'] = $col;
				}
			}
			$result->close();
		}
	}
	
	$db->close();
	
	return $tables;
}

function db_export() {
	//global $lang_rawstr_php_dq;

	//from def_config.php or config.php
	global $target, $db_config, $dump_file_head, $vendor_special_table;

    	$target_vendor = exec("hg branch");
	$code = $dump_file_head."\n";
	
	$host = $db_config[$target]['host'];
	$user = $db_config[$target]['user'];
	$pass = $db_config[$target]['pass'];
	$name = $db_config[$target]['name'];
	$port = $db_config[$target]['port'];
	$dir  = "./servers/".$db_config[$target]['server'].'/';

	$db = db_connect();

	// 导出前重置event_version = 0
	$db->query("use ".$name);
	$db->query("UPDATE `server_info` SET `event_version` = 0;" );


	$db->query("use information_schema");

	// 获取玩家表
	$sql = "SELECT `TABLE_NAME` FROM `TABLES` WHERE (`TABLE_NAME` LIKE 'player%' OR `TABLE_NAME` LIKE 'global_%' OR `TABLE_NAME` LIKE 'stat_%_log') AND `TABLE_SCHEMA` = '".$name."'";
	$result = $db->query($sql);
	$rows = array();
	while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
		$rows[] = $row;
	}
	//特殊表
	foreach($vendor_special_table as $table_name => $vendors) {
		$rows[] = array("TABLE_NAME" => $table_name);
	}
	$result->close();
	
	$cmd = "mysqldump -h$host -P$port -u$user -p$pass --skip-add-locks --skip-comments --ignore-table={$name}.db_version ";

	foreach ($rows as $row) {
		$cmd .= " --ignore-table=".$name.".".$row['TABLE_NAME']." ";
	}

	$cmd .= "$name";

	//$code .= "	\$this->AddSQL(\"\n";
	$code .= "	\$this->AddSQL(<<<THESQL1\n";
	ob_start();
	passthru($cmd);
	$dump_sql  = ob_get_contents();
	//$code .= $lang_rawstr_php_dq->replace($dump_sql);
	$code .= str_replace('$', '\$', $dump_sql);
	ob_end_clean();
	//$code .= "\n	\");\n\n";
	$code .= "\nTHESQL1\n\t);\n\n";

	//vendor special table  
	foreach($vendor_special_table as $table_name => $vendors) {
		$exclude_vendors = 'array(';
		foreach($vendors as $vendor) {
			if ($vendor != $target_vendor) {
				$exclude_vendors .= "'$vendor', ";
			}
		}
		$exclude_vendors .= ' )';
		$cmd = "mysqldump -h$host -P$port -u$user -p$pass --skip-add-locks --skip-comments {$name} {$table_name}";
		$code .= "	\$this->AddSQL(<<<THESQLVS{$table_name}\n";
		ob_start();
		passthru($cmd);
		$dump_sql  = ob_get_contents();
		$code .= str_replace('$', '\$', $dump_sql);
		ob_end_clean();
		$code .= "\nTHESQLVS{$table_name}\n\t ,false, {$exclude_vendors});\n\n";
	}

	// 获取非玩家表
	$sql = "SELECT `TABLE_NAME` FROM `TABLES` WHERE `TABLE_NAME` NOT LIKE 'player%' AND `TABLE_NAME` NOT LIKE 'global_%' AND `TABLE_NAME` NOT LIKE 'stat_%_log' AND `TABLE_SCHEMA` = '".$name."'";
	$result = $db->query($sql);
	$rows = array();
	while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
		$rows[] = $row;
	}
	$result->close();
	
	$cmd = "mysqldump -h$host -P$port -u$user -p$pass --skip-add-locks --skip-comments --no-data ";

	foreach ($rows as $row) {
		$cmd .= " --ignore-table=".$name.".".$row['TABLE_NAME']." ";
	}

	$cmd .= "$name";

	$code .= "if (\$isNewDatabase) {\n";
	$code .= "	\$this->AddSQL(<<<THESQL2\n";
	ob_start();
	passthru($cmd);
	$code .= ob_get_contents();
	ob_end_clean();
	$code .= "\nTHESQL2\n\t);\n";
	$code .= "}\n";

	$code .= "?>";

	$changesInfo = get_changes();

	$last_version = $changesInfo['last_version'];
	$last_version_name = $changesInfo['last_version_name'];

	$prefix = substr($last_version_name, 0, 11);
	$version = sprintf("%02d", substr($last_version_name, 11, 2) + 1);

	date_default_timezone_set('Asia/Shanghai');
	if (date('Y-m-d') > $prefix) {
		$prefix = date('Y-m-d').'-';
		$version = '01';
		$last_version = date('Ymd').'00';
	}

	$output_file = $dir.$prefix.$version.".php";

	file_put_contents($output_file, $code);

	echo $output_file."\n";

	$db->query("use ".$name);
	$db->query("UPDATE `db_version` SET `version` = ".($last_version + 1) . ";" );

	//================================
	//if ($name == 'lms_trunk') {
	if (is_exist_reward($db,$name)){
		$code = get_insert_code_for_global_table($db, $db_config[$target]['name']);

		$version = sprintf("%02d", $version+1);

		$output_file2 = $dir.$prefix.($version).".php";

		file_put_contents($output_file2, $code);

		echo $output_file2."\n";
	}
	//================================

	$db->close();
}

/*
  导出同步悬赏表的脚本
*/
function is_exist_reward($db,$db_name){
	$db->query("use `information_schema`");
	$count = db_query($db, "select count(1) as `num` from `TABLES` where `TABLE_SCHEMA`='{$db_name}' and `TABLE_NAME`='reward';");
	return $count[0]['num']==1;
}

function get_insert_code_for_global_table($db,$db_name)
{	
	$code = "<?php \n";
	$db->query("use ".$db_name);

	$rewards = db_query($db, "SELECT * FROM `reward` where `reward_type` = 1;");	
	$code .= "\$rewardIds = array(";
	foreach ($rewards as $reward) {
		$code .= "{$reward['id']},";
	}
	$code .= ");\n";
	
	$code .= <<<EOT
		foreach (\$rewardIds as \$rewardId) {
			\$sql1 = "SELECT * FROM `global_reward` WHERE id = {\$rewardId};";
			\$res = db_query(\$db, \$sql1);		
			if (count(\$res) == 0) {
				\$sql2 = "INSERT INTO `global_reward` (`id`,`hunter_num`,`hunter_id`,`hunter_nick`) VALUES ({\$rewardId},0,0,'');";
				db_execute(\$db, \$sql2);
			}
		}\n
EOT;

	$kendos = db_query($db, "select * from `kendo` order by `order`;");	
	$code .= "\$kendoIds = array(";
	foreach ($kendos as $kendo) {
		$code .= "{$kendo['id']},";
	}
	$code .= ");\n";
	
	$code .= <<<EOT
		foreach (\$kendoIds as \$kendoId) {
			\$sql3 = "SELECT * FROM `global_kendo` WHERE id = {\$kendoId};";
			\$res = db_query(\$db, \$sql3);
			if (count(\$res) == 0) {
				\$sql4 = "INSERT INTO `global_kendo` (`id`,`fail_num`) VALUES ({\$kendoId},0);";
				db_execute(\$db, \$sql4);
			}
		}\n
EOT;
	
	$code .= <<<EOT
		\$sql5 = "SELECT * FROM `global_no_vip_group_num` WHERE id = 1;";
		\$res = db_query(\$db, \$sql5);
		if (count(\$res) == 0) {
			\$sql6 = "INSERT INTO `global_no_vip_group_num` (`id`,`num`) VALUES (1,0);";
			db_execute(\$db, \$sql6);
		}\n
EOT;

	$code .= <<<EOT
		\$sql7 = "SELECT * FROM `global_platform` WHERE id = 1;";
		\$res = db_query(\$db, \$sql7);
		if (count(\$res) == 0) {
			\$sql8 = "INSERT INTO `global_platform` (`id`,`value1`,`value2`) VALUES (1,0,0);";
			db_execute(\$db, \$sql8);
		}\n
EOT;

$code .= <<<EOT
		\$sql9 = "SELECT * FROM `global_service_activity` WHERE id = 1;";
		\$res = db_query(\$db, \$sql9);
		if (count(\$res) == 0) {
			\$sql10 = "INSERT INTO `global_service_activity` (`id`,`vip_num`,`gold_ghost_num`,`purple_ghost_num`,`knight_num`) VALUES (1,0,0,0,0);";
			db_execute(\$db, \$sql10);
		}\n
EOT;

	$code .= "?>";

	return $code;
}

/*
  数据库脚本升级的时候，效验脚本是否符合规范
*/
function db_upgrade_validate($changes){
	foreach (array_keys($changes) as $version) {
		$file = "./".$changes[$version]['dir']."/".$changes[$version]['name'].".php";
		$data = file_get_contents($file);

		if ($changes[$version]['is_dump']){ // 忽略数据导出脚本
			continue;
		}

		$haveDDL = FALSE;
		$haveDML = FALSE;

		$dataList = explode('`', $data);
		foreach ($dataList as $key => $value) {
			if ($key%2 == 0){
				// echo($value."\n");
				$datatmp = strtoupper($value);

				if (strpos($datatmp,"ALTER")!== false || strpos($datatmp,"CREATE")!== false || strpos($datatmp,"DECLARE")!== false){
					$haveDDL = true;
				}

				if (strpos($datatmp,"INSERT")!== false || strpos($datatmp,"UPDATE")!== false || strpos($datatmp,"DELETE")!== false){
					$haveDML = true;
				}
				if ($haveDDL && $haveDML){
					die("\nuvalidate fialed at ".$version );
				}
			}
		}
	}
}

/*
  清除玩家表
*/
function db_clean() {
	global $target, $db_config;

	$code = "<?php\n";
	
	$name = $db_config[$target]['name'];

	$db = db_connect();
	$db->query("use information_schema");

	// 获取玩家表
	$sql = "SELECT `TABLE_NAME` FROM `TABLES` WHERE (`TABLE_NAME` LIKE 'player%' OR `TABLE_NAME` LIKE 'global_%' OR `TABLE_NAME` LIKE 'stat_%_log') AND `TABLE_SCHEMA` = '".$name."'";
	$result = $db->query($sql);
	$rows = array();
	while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
		$rows[] = $row;
	}
	$result->close();
	
	$db->query("use ".$name);
	
	foreach ($rows as $row) {
		$sql = "DELETE FROM ".$row['TABLE_NAME'].";\n";
		echo $sql;
		db_execute($db, $sql);
	}
	
	$db->close();
}

function db_check($check_item) {
	$check_file = "./updates/check_".$check_item.".php";

	if (!is_file($check_file)) {
		echo "无检查文件: $check_file\n";
		exit;
	}

	$db = db_connect();

	require_once $check_file;
}

// 
// 测试文本是否是utf8编码
// 
// 返回值：
//   1 - 有BOM头的内容
//   2 - 纯utf8的内容
//   3 - 较可能是utf8的内容
//   4 - 较不可能是utf8的内容
// 
function utf8_check($text)
{
	$utf8_bom = chr(0xEF).chr(0xBB).chr(0xBF);
	
	// BOM头检查
	if (strstr($text, $utf8_bom) === 0)
		return 1;
	
	$text_len = strlen($text);
	
	// UTF-8是一种变长字节编码方式。对于某一个字符的UTF-8编码，如果只有一个字节则其最高二进制位为0；
	// 如果是多字节，其第一个字节从最高位开始，连续的二进制位值为1的个数决定了其编码的位数，其余各字节均以10开头。
	// UTF-8最多可用到6个字节。
	//
	// 如表：
	// < 0x80 1字节 0xxxxxxx
	// < 0xE0 2字节 110xxxxx 10xxxxxx
	// < 0xF0 3字节 1110xxxx 10xxxxxx 10xxxxxx
	// < 0xF8 4字节 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
	// < 0xFC 5字节 111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
	// < 0xFE 6字节 1111110x 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
	
	$bad   = 0; // 不符合utf8规范的字符数
	$good  = 0; // 符号utf8规范的字符数
	
	$need_check = 0; // 遇到多字节的utf8字符后，需要检查的连续字节数
	$have_check = 0; // 已经检查过的连续字节数
	
	for ($i = 0; $i < $text_len; $i ++) {
		$c = ord($text[$i]);

		if ($need_check > 0) {
			$c = ord($text[$i]);
			$c = ($c >> 6) << 6;
			
			$have_check ++;
			
			// 10xxxxxx ~ 10111111
			if ($c != 0x80) {
				$i -= $have_check;
				$need_check = 0;
				$have_check = 0;
				$bad ++;
			}
			else if ($need_check == $have_check) {
				$need_check = 0;
				$have_check = 0;
				$good ++;
			}
			
			continue;
		}
		
		if ($c < 0x80)	  // 0xxxxxxx
			$good ++;
		else if ($c < 0xE0) // 110xxxxx
			$need_check = 1;
		else if ($c < 0xF0) // 1110xxxx
			$need_check = 2;
		else if ($c < 0xF8) // 11110xxx
			$need_check = 3;
		else if ($c < 0xFC) // 111110xx
			$need_check = 4;
		else if ($c < 0xFE) // 1111110x
			$need_check = 5;
		else
			$bad ++;
	}
	
	if ($bad == 0)
		return 2;
	else if ($good > $bad)
		return 3;
	else
		return 4;
}
?>
