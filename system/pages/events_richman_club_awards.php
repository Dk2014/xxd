<?php

$pconfig = array(
	'title'   => '土豪俱乐部',
	'table'   => 'events_richman_club_awards',
	'links'		=> array(),
	
	'columns' => array(
		array('name' => 'require_vip_level', 'text' => '目标vip', 'width' => '80px'),
		array('name' => 'require_vip_count', 'text' => '目标vip人数', 'width' => '80px'),
		array('name' => 'award_vip_level1', 'text' => '1阶vip等级', 'width' => '80px'),
		array('name' => 'award_vip_item1_id', 'text' => '1阶奖励物品', 'width' => '80px'),
		array('name' => 'award_vip_item1_num', 'text' => '1阶奖励数量', 'width' => '80px'),
		array('name' => 'award_vip_level2', 'text' => '2阶vip等级', 'width' => '80px'),
		array('name' => 'award_vip_item2_id', 'text' => '2阶奖励物品', 'width' => '80px'),
		array('name' => 'award_vip_item2_num', 'text' => '2阶奖励数量', 'width' => '80px'),
		array('name' => 'award_vip_level3', 'text' => '3阶vip等级', 'width' => '80px'),
		array('name' => 'award_vip_item3_id', 'text' => '3阶奖励物品', 'width' => '80px'),
		array('name' => 'award_vip_item3_num', 'text' => '3阶奖励数量', 'width' => '80px'),
		array('name' => 'award_vip_level4', 'text' => '4阶vip等级', 'width' => '80px'),
		array('name' => 'award_vip_item4_id', 'text' => '4阶奖励物品', 'width' => '80px'),
		array('name' => 'award_vip_item4_num', 'text' => '4阶奖励数量', 'width' => '80px'),
		array('name' => 'award_vip_level5', 'text' => '5阶vip等级', 'width' => '80px'),
		array('name' => 'award_vip_item5_id', 'text' => '5阶奖励物品', 'width' => '80px'),
		array('name' => 'award_vip_item5_num', 'text' => '5阶奖励数量', 'width' => '80px'),

	),
	'where' 		=> 'sql_where',
	'js' 			=> 'js_function',
);

?>