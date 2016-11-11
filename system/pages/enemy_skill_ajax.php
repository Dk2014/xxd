<?php
	require_once dirname(__FILE__).'/../lib/global.php';

	// 提交
 	if (isset($_POST['config'])
 		&& isset($_POST['id']) && $_POST['id'] > 0) {

	 	$sql = "update enemy_role set skill_config='{$_POST['config']}' where id='{$_POST['id']}'";
 		db_query($db, $sql);

 	} else if (isset($_GET['id']) && $_GET['id']>0) {
		$sql = "select skill_config from enemy_role where id='{$_GET['id']}'";
 		$rows = db_query($db, $sql);
 		if (!empty($rows[0]['skill_config']))
	 		echo $rows[0]['skill_config'];
	 	else
	 		echo "[]";
 	}
?>