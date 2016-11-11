<?php
/*
$host = $argv[1];
$user = $argv[2];
$pass = $argv[3];
$name = $argv[4];
$port = $argv[5];

$gsid = $argv[6];
$ip = $argv[7];
$game_port = $argv[8];
$sid = $gsid/10;
 */

// 内网测试组专用
//$host = "172.26.38.101";
//$user = "root";
//$pass = "mysql";

//厦门数值
$host = "172.26.166.143";
$user = "root";
$pass = "";

$name = "v3";
$port = 3306;

$server_info = array(
	100021 => array(
		'gsid' => 100021,
		'ip' => '172.26.160.11',
		'port' => '7080',
	),
	100022 => array(
		'gsid' => 100022,
		'ip' => '172.26.160.11',
		'port' => '8080',
	),

	100011 => array(
		'gsid' => 110011,
		'ip' => '172.26.166.143',
		'port' => '7080',
	),

	100012 => array(
		'gsid' => 110012,
		'ip' => '172.26.166.143',
		'port' => '8080',
	),

);


$db = new mysqli($host, $user, $pass, $name, $port);
if ($db->connect_error) {
	die('连接数据库失败：Error '.$db->connect_errno.', '.$db->connect_error);
}

$db->query("set names utf8mb4");

$sql = "select p.id as pid, p.user as user, p.nick as nick, pr.role_id as role_id, pr.level as level from player p left join player_role pr on p.`main_role_id`=pr.id;";


$result = $db->query($sql, MYSQLI_USE_RESULT);

if (!$result) {
	die('Query Error (' . $db->errno . ') ' . $db->error);
}


$json_dat = array();

while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
	$gsid = intval($row['pid']) >> 32;
	$json_dat[$row['user']] = array(
		'Nick'=> $row['nick'],
		'RoleId' => intval($row['role_id']),
		'Gsid' => $gsid,
		'Ip' => $server_info[$gsid]['ip'],
		'Port' => $server_info[$gsid]['port'],
		'RoleLevel' => intval($row['level']),
		'Sid' => intval($gsid/10),
	);
}

$result->close();


echo json_encode($json_dat, JSON_UNESCAPED_UNICODE);

echo "\n";


?>
