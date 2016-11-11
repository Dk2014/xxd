<?php


//城镇
$sql = "select `id`,`name`, `lock` from `town`";
$rows = db_query($db, $sql);

$i = 0;
$allTown = array('0'=>'无');
$allTownLock = array('0'=>'无');

foreach ($rows as $value) {
	$allTown[$value['id']] = $value['name'];
	$allTownLock[$value['lock']] = $value['name'];
	$linksTown[$i]['text'] = $value['name'];
	$linksTown[$i]['id'] = "quest&m=".$value['id'];
	$i++;
}
$allTownLock['0'] = '无';


$pconfig = array(
	'title'   => '主线任务',
	'table'   => 'quest',
	'links'   => $linksTown,
	
	'columns' => array(
		array('name' => 'related_town', 'text' => '城镇分类', 'width' => '50px'),
		array('name' => 'order', 'text' => '排序', 'width' => '50px'),
		array('name' => 'name', 'text' => '标题', 'width' => '100px'),
		array('name' => 'type', 'text' => '类型', 'width' => '50px'),
		array('name' => 'require_level', 'text' => '需求等级', 'width' => '80px'),
		array('name' => 'town_id', 'text' => '城镇', 'width' => '50px'),
		array('name' => 'town_npc_id', 'text' => '城镇NPC', 'width' => '50px'),
		array('name' => 'mission_level_id', 'text' => '关卡', 'width' => '50px'),
		array('name' => 'enemy_id', 'text' => '敌人', 'width' => '50px'),
		array('name' => 'enemy_num', 'text' => '敌人组数', 'width' => '80px'),
		array('name' => 'drama_mode', 'text' => '剧情模式', 'width' => '50px'),

		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
		array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_physical', 'text' => '奖励体力', 'width' => '80px'),

		array('name' => 'award_item1_id', 'text' => '奖励物品1', 'width' => '50px'),
		array('name' => 'award_item1_num', 'text' => '奖励物品1数量', 'width' => '50px'),		

		array('name' => 'award_item2_id', 'text' => '奖励物品2', 'width' => '50px'),
		array('name' => 'award_item2_num', 'text' => '奖励物品2数量', 'width' => '50px'),

		array('name' => 'award_item3_id', 'text' => '奖励物品3', 'width' => '50px'),
		array('name' => 'award_item3_num', 'text' => '奖励物品3数量', 'width' => '50px'),

		array('name' => 'award_item4_id', 'text' => '奖励物品4', 'width' => '50px'),
		array('name' => 'award_item4_num', 'text' => '奖励物品4数量', 'width' => '50px'),		

		array('name' => 'award_func_key', 'text' => '奖励功能权值', 'width' => '50px'),
		array('name' => 'award_town_key', 'text' => '奖励城镇权值', 'width' => '50px'),
		array('name' => 'award_mission_level_lock', 'text' => '奖励关卡权值', 'width' => '50px'),

		array('name' => 'award_role_id', 'text' => '奖励角色', 'width' => '50px'),
		array('name' => 'award_role_level', 'text' => '奖励角色等级', 'width' => '80px'),

		array('name' => 'auto_fight', 'text' => '自动打怪', 'width' => '50px'),
		array('name' => 'show_black_curtain', 'text' => '黑幕', 'width' => '50px'),
		array('name' => 'pass_mission_share', 'text' => '通关区域分享', 'width' => '50px'),
		array('name' => 'desc', 'text' => '简介', 'width' => '0px'),
		array('name' => 'mission_drama_talk', 'text' => '区域对话', 'width' => '0px'),
	),
	
	'opera' => array(),

	'where' 		=> 'sql_where',
	'js'			=> 'jsFunction',
	'batch_update_key' => 'id',
	'batch_update_password' => '789',
);




?>
