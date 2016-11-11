<?php

$hardLevelTownLinks = array();
$hardLevelTownLinks[0]['text'] = "未使用";
$hardLevelTownLinks[0]['id'] = "hard_level&m=0";
$sql = "select `id`,`name` from `town`";
$allTowns = db_query($db, $sql);
$townMap = array();
$i = 1;
foreach ($allTowns as $value) {
	$townMap[$value['id']] = $value['name'];
	$hardLevelTownLinks[$i]['text'] = $value['name'];
	$hardLevelTownLinks[$i]['id'] = "hard_level&m=".$value['id'];
	$i++;
}

$pconfig = array(
	'title'   => '难度关卡',
	'table'   => 'hard_level',
	'links'   => $hardLevelTownLinks,

	'columns' => array(
		array('name' => 'mission_level_lock', 'text' => '区域关卡权值', 'width' => '80px'),
		array('name' => 'hard_level_lock', 'text' => '难度关卡权值', 'width' => '80px'),
		array('name' => 'award_hard_level_lock', 'text' => '奖励难度关卡权值', 'width' => '80px'),
		array('name' => 'desc', 'text' => '关卡描述', 'width' => '80px'),
	),


	'opera' => array(
		array('type' => 'link', 'text' => '关卡', 'render' => 'leveltable'),
	),

	'not_delete' 	=> true,
	'insert' => 'sql_insert',
	'where' => 'sql_where',
	'location' => 'location',
);
?>
