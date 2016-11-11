<?php

include "common_links.php";

function totem_opera($row) {
	return html_get_link("技能配置", "?p=totem_skill&m=".$row['id'], false);
}

$pconfig = array(
	'title' => '阵印等级配置',
	'table' => 'totem_level_info',
	'links' => array(),

	'columns' => array(
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'level', 'text' => '等级', 'width' => '80px'),
		array('name' => 'health', 'text' => '生命', 'width' => '60px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '60px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '60px'),
		array('name' => 'cultivation', 'text' => '内力', 'width' => '60px'),
		array('name' => 'rock_rune_rate', 'text' => '分解获得石符文概率', 'width' => '60px'),
		array('name' => 'rock_rune_num', 'text' => '分解获得石符文数量', 'width' => '60px'),
		array('name' => 'jade_rune_rate', 'text' => '分解获得玉符文概率', 'width' => '60px'),
		array('name' => 'jade_rune_num', 'text' => '分解获得玉符文数量', 'width' => '60px'),
		array('name' => 'upgrade_need_rock', 'text' => '铭刻所需石符文数量', 'width' => '60px'),
		array('name' => 'upgrade_need_jade', 'text' => '铭刻所需玉符文数量', 'width' => '60px')
	),
	'js' => 'js_function',
	'links' => $totem_links,

);
