<?php

set_time_limit(0);

$db_name = $argv[1];
$upgrade_dir = $argv[2];
$needBackup_db = (!isset($argv[3]) || !$argv[3]) ? false : true;
 
require_once dirname(__FILE__) . "/config.php";
require_once dirname(__FILE__) . "/cls_dbupgrader.php";

$upgrader = new GameDBUpgrader($upgrade_dir, $needBackup_db, $db_config[$db_name]['host'], $db_config[$db_name]['user'],
 $db_config[$db_name]['pass'], $db_config[$db_name]['name'], $db_config[$db_name]['port']);

$upgrader->ExportSingleUpgradeFile("./database-game_server-upgrade.php");

?>