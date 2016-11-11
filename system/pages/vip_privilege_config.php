<?php
include "vip_common.php";

$level = $_REQUEST['m'];
if (isset($level) && $level > 1 ) {
	$sql = "select id from vip_privilege_config where level={$level}";
	$configRows = db_query($db, $sql);
	if(count($configRows) == 0) {
		$preLevel = $level - 1;
		$sql = "select * from vip_privilege_config where level={$preLevel}";
		$preLevelConfigRows = db_query($db, $sql);
		foreach($preLevelConfigRows as $row) {
			$sql = "insert into vip_privilege_config(`privilege_id`, `level`, `times`, `unit`) 
				values('{$row['privilege_id']}', '{$level}', '{$row['times']}', '{$row['unit']}') ";
			db_execute($db, $sql);
		}
	}

}

$pconfig = array(
	'title'   => 'VIP特权配置',
	'table'   => 'vip_privilege_config',
	'links'   => array(),

	'columns' => array(
		array('name' => 'privilege_id', 'text' => '特权ID', 'width' => '80px'),
		array('name' => 'times', 'text' => '特权次数', 'width' => '80px'),
		array('name' => 'unit', 'text' => '单位', 'width' => '80px'),
	),
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'location' => 'location',
);

