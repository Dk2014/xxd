<?php
include "battle_pet_links.php";


$pconfig = array(
	'title'   => '灵宠',
	'table'   => 'battle_pet',
	'links'   => $battle_pet_links,
	
	'columns' => array(
		array('name' => 'pet_id', 'text' => '灵宠', 'width' => '150px'),
		//array('name' => 'parent_pet_id', 'text' => '父灵宠', 'width' => '150px'),
		//array('name' => 'star', 'text' => '星', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'skill', 'text' => '技能', 'width' => '40px'),
		//array('name' => 'force', 'text' => '威力', 'width' => '40px'),

		//array('name' => 'health', 'text' => '生命', 'width' => '40px'),
		//array('name' => 'attack', 'text' => '攻击', 'width' => '40px'),
		//array('name' => 'defence', 'text' => '防御', 'width' => '40px'),
		//array('name' => 'speed', 'text' => '速度', 'width' => '40px'),

		array('name' => 'round_attack', 'text' => '单回合行动次数', 'width' => '80px'),
		array('name' => 'cost_power', 'text' => '召唤时消耗精气', 'width' => '80px'),
		array('name' => 'live_round', 'text' => '召唤后存活回合数', 'width' => '80px'),
		array('name' => 'live_pos', 'text' => '召唤出现位置', 'width' => '80px'),
		array('name' => 'desc', 'text' => '灵宠描述', 'width' => '80px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '捕捉概率', 'render' => 'battle_pet_opera'),
	    array('type' => 'link', 'text' => '等级数据', 'render' => 'battle_pet_level_info'),
	),
	
	'not_delete' 	=> true,
	'js'			=> 'jsFunction',
	'where' => 'sql_where',
);

function battle_pet_opera($row) {
	$html = html_get_link("捕捉概率", "?p=battle_pet_catch_rate&m={$row['pet_id']}", false);
	$html.=' | ';
	return $html;
}

function battle_pet_level_info($row) {
    $html = html_get_link("等级数据", "?p=battle_pet_level_info&m={$row['pet_id']}", false);
    return $html;
}

?>
