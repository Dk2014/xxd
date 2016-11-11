<?php
$p = dirname(__FILE__);
$config_path = $p."/config.php";
$database_name = $argv[1];

$file = fopen($config_path, 'r');
$data = fread($file, filesize($config_path));
fclose($file);

$pattern = '/(local_game.+name\W+)(\w+)/s';
$replace = '${1}'.$database_name;
$data = preg_replace($pattern, $replace, $data);
$file = fopen($config_path, 'w');
fwrite($file, $data);
fclose($file);
?>
