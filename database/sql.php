<?php

//
// 命令行参数处理
//
$dbtpl_name      = $argv[1];  // db模版名称, eg: game_server, platform_server
$dbversion_start = $argv[2];  // 起点 db_version
$dbversion_end   = $argv[3];  // 终点 db_version
$db              = $argv[4];  // 目标SQL

if (!$argv[4]) {
  echo "usage: php sql.php game_server 2014050401 2015010101 ~/upgrade_2014050401_to_2015010101.sql";
  exit(-1);
}

function sql_append($filename, $line) {
  file_put_contents($filename, $line, FILE_APPEND);
}

function db_execute($filename, $sql) {
  sql_append($filename, $sql."\n");
}

$dir = './servers/'.$dbtpl_name;

// 按版本号差异获取变更文件列表
$changes = array();

if ($handle = opendir($dir)) {
	while (FALSE !== ($file = readdir($handle))) {
		if ($file == "." || $file == ".." || is_dir($file))
			continue;
		
		if (strrpos($file, ".php") == 13) {
			if (strlen($file) != 17)
				continue;
			
			$name = substr($file, 0, -4);
			$version = (int)str_replace('-', '', $name);
		}
		else {
			continue;
		}
    
    if ($version <= $dbversion_start) 
      continue;

    if ($version > $dbversion_end) 
      continue;
    
		$changes[$version] = array('name' => $name, 'dir' => $dir);
	}
} else {
  echo "error: {$dir} not exist";
  exit(-1);
}

$renew = false;
@unlink($db);

asort($changes);

$last_dbversion = 0;

db_execute($db, "SET NAMES utf8mb4;");
  
foreach (array_keys($changes) as $version) {
  db_execute($db, "/* only exec following sql when db_version.version = {$last_dbversion} */");
  
  sql_append($db, <<<EOT
    DELIMITER //
    DROP PROCEDURE IF EXISTS `UPGRADE_XXDS_PROC`;
    CREATE PROCEDURE `UPGRADE_XXDS_PROC`()
    BEGIN
      DECLARE DBVersion INT;
      DECLARE DBVersionCount INT;
  
      CREATE TABLE IF NOT EXISTS `db_version` ( `version` INT ) ENGINE 'InnoDb' CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

      SELECT COUNT(*) INTO DBVersionCount FROM `db_version`;
      if (DBVersionCount = 0) THEN
        INSERT INTO `db_version` SET `version` = 0; 
      END IF;
      SELECT `version` INTO DBVersion FROM `db_version` WHERE 1 LIMIT 1;
      IF (DBVersion = {$last_dbversion}) THEN 
      
EOT
);
  
  require_once("./".$changes[$version]['dir']."/".$changes[$version]['name'].".php");
  db_execute($db, "UPDATE `db_version` SET `version` = ".$version.";");
  $last_dbversion = $version;
  
  sql_append($db, <<<EOT
    
    END IF;
  END //

  DELIMITER ;
  CALL UPGRADE_XXDS_PROC();
  DROP PROCEDURE IF EXISTS `UPGRADE_XXDS_PROC`;
  
EOT
  );
}

echo "done";
?>