<?php
require_once '../lib/global.php';
if (isset($_GET['id'])) {
	// 查询数据
	$sql = "SELECT `desc` FROM `clique_building` WHERE `id`={$_GET['id']} LIMIT 1";
	$rows = db_query($db, $sql);
	print_r($rows[0]['desc']);
}else if(isset($_POST['id'])){
	// 提交数据
	$descs = $_POST['descs'];
	$sql = "UPDATE `clique_building` SET `desc`='{$descs}' WHERE `id`={$_POST['id']}";
	db_execute($db, $sql);
}

?>