<?php
$db = db_connect();

$sql = "select * from town_npc_item where item_id not in (select item_id from purchase_limit);";
$rows = db_query_array($db, $sql);
$if(count($rows)>0) {
	die("城镇商人：有物品未配置次数");
}




?>
