<?php
require_once dirname(__FILE__).'/../lib/global.php';

switch ($_SERVER['REQUEST_METHOD']) {
case 'POST':
	assert(isset($_POST['id']), "save_enemy_passive_skill need `id`");
	error_log($_POST['config']);
	$sql = "replace into `enemy_passive_skill` set `config`='{$_POST['config']}', `id` = {$_POST['id']}";
	error_log($sql);
	db_execute($db, $sql);
	break;
case 'GET':
	assert(isset($_GET['id']), "load_enemy_passive_skill need `id`");
	$sql = "select config from enemy_passive_skill where id='{$_GET['id']}'";
	$rows = db_query($db, $sql);
	if (!empty($rows[0])) {
 		if (!empty($rows[0]['config']))
	 		echo $rows[0]['config'];
		//echo json_encode($rows[0], true);
		//error_log(json_decode($rows[0]), true);
	}
	else {
		echo "[]";
	}
	break;
default:
}
?>
