<?php
	require_once dirname(__FILE__).'/../lib/global.php';

	// 提交
 	if (isset($_POST['item_config']) && 
 		isset($_POST['item_id']) && 
 		isset($_POST['target_type']) && 
 		isset($_POST['effect_type']) &&
 		isset($_POST['can_use_level'])) {

 		$sql = "select * from battle_item_config where id='{$_POST['item_id']}'";
 		$rows = db_query($db, $sql);

 		$keep = isset($_POST['keep'])? $_POST['keep'] : 0;
 		$cost_power = isset($_POST['cost_power'])? $_POST['cost_power'] : 0;
 		$max_override = isset($_POST['max_override'])? $_POST['max_override'] : 0;
 		$can_use_level = isset($_POST['can_use_level'])?$_POST['can_use_level']:0;

 		$set_sql = "target_type='{$_POST['target_type']}', effect_type='{$_POST['effect_type']}', 
 		config='{$_POST['item_config']}',cost_power='{$cost_power}',keep='{$keep}',max_override='{$max_override}',can_use_level='{$can_use_level}'";
 		if (count($rows) > 0) {
 			$sql = "update battle_item_config set {$set_sql} where id='{$_POST['item_id']}'";
 		} else {
 			$sql = "insert battle_item_config set {$set_sql},id='{$_POST['item_id']}'";
 		}
 		db_query($db, $sql);
 	} else if (isset($_GET['item_id'])) {
		$sql = "select * from battle_item_config where id='{$_GET['item_id']}'";
 		$rows = db_query($db, $sql);
 		echo json_encode($rows[0]);
 	}
?>