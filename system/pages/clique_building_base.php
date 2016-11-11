<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派建筑物 -- 总舵',
	'table'   => 'clique_center_building_level_info',
	'links'   => $clique_buildings_link,
	
	'columns' => array(
		array('name' => 'level', 'text' => '总舵等级', 'width' => '200px'),
		array('name' => 'max_member', 'text' => '总舵人数上限', 'width' => '200px'),
		array('name' => 'daily_max_coins', 'text' => '总舵每日贡献上限', 'width' => '200px'),
		array('name' => 'levelup_coins', 'text' => '总舵升级所需铜钱', 'width' => '200px'),
		array('name' => 'desc', 'text' => '对应等级描述', 'width' => '200px'),
		array('name' => 'cur_level_desc', 'text' => '当前总舵描述', 'width' => '200px'),
		array('name' => 'next_level_desc', 'text' => '下一等级总舵描述', 'width' => '200px'),
	),
	
	'opera' => array(
	),
	

	'where' 		=> 'sql_where',
);

function sql_where(){
	return ' ORDER BY level ASC';
}

?>