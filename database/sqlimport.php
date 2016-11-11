<?php

$db_name		  = $argv[1];
$db_sql_file      = $argv[2];

require_once dirname(__FILE__)."/config.php";

$cfg = $db_config[$db_name];
echo "\nexcute upgrade game database to mysql...\n";
echo shell_exec("mysql -h{$cfg['host']} -u{$cfg['user']} -p{$cfg['pass']} -P{$cfg['port']} --default-character-set=utf8mb4 {$cfg['name']} < {$db_sql_file}");


$db = new mysqli($cfg['host'], $cfg['user'], $cfg['pass'], $cfg['name'], $cfg['port']);
if ($db->connect_error) {
	die('game database upgrade error. mysql connect error. '.$db->connect_error);
}

$result = $db->query("select * from db_version");
$row = $result->fetch_array(MYSQLI_ASSOC);
$result->close();

echo "\n\n[info] database ** {$cfg['name']} ** upgrade to ** {$row['version']} **\n\n\n";

?>