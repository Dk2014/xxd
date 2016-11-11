<?php



$pconfig = array(
	'title'   => '区域关卡宝箱',
	'table'   => 'mission_level_box',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'order', 'text' => '品质顺序', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖励类型', 'width' => '80px'),
		array('name' => 'award_chance', 'text' => '奖励概率%', 'width' => '80px'),
		array('name' => 'award_num', 'text' => '奖励数量', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '物品ID(仅物品奖励填写)', 'width' => '150px'),

	),
	
	'opera' => array(
	),
	
	'js'=>'js_function',
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'location' => 'location',
);
global $db;
$sql = "SELECT parent_type FROM `mission_level` WHERE id = " . htmlspecialchars($_GET['m']);
$finds = db_query_array($db, $sql);
if(intval($finds[0][0]) == 8 ||intval($finds[0][0]) == 0 || intval($finds[0][0]) == 1){  //是否为难度关卡或者区域关 或者资源关卡
	array_push($pconfig['columns'], array('name' => 'must_in_first', 'text' => '第一次必然获得？', 'width' => '80px'));
}
?>
