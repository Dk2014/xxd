<?php

$addition_quest_link = array(
	array('text' => '全部支线任务', 'id' => 'addition_quest'),
	array('text' => '孤立支线任务', 'id' => 'addition_quest&m=0')
);

$sql = "select * from ( select serial_number as sn, name, require_lock from addition_quest where award_lock>0 order by require_lock ) q group by q.sn;";
$query = db_query($db, $sql);
foreach($query as $key => $row) {
	$addition_quest_link[] = array(
		'text' => "{$row['name']} 系列",
		'id' => "addition_quest&m={$row['sn']}",
	);
}

$pconfig = array(
	'title'   => '支线任务',
	'table'   => 'addition_quest',
	'links'   => $addition_quest_link,
	
	'columns' => array(
		array('name' => 'serial_number', 'text' => '任务系列ID', 'width' => '80px'),
		array('name' => 'award_lock', 'text' => '奖励任务系列权值', 'width' => '40px'),
		array('name' => 'require_serial_number', 'text' => '前置任务系列', 'width' => '80px'),
		array('name' => 'require_lock', 'text' => '要求任务系列权值', 'width' => '40px'),
		array('name' => 'require_level', 'text' => '要求等级', 'width' => '40px'),

		array('name' => 'publish_npc', 'text' => '发布NPC', 'width' => '180px'),
		array('name' => 'name', 'text' => '名称', 'width' => '150px'),
		array('name' => 'description', 'text' => '简介', 'width' => '180px'),
		array('name' => 'showup_main_quest', 'text' => '出现任务', 'width' => '80px'),
		array('name' => 'disappear_main_quest', 'text' => '消失任务', 'width' => '80px'),


		array('name' => 'type', 'text' => '类型', 'width' => '80px'),
		array('name' => 'npc_id', 'text' => '对话NPC', 'width' => '80px'),
		array('name' => 'mission_id', 'text' => '任务区域', 'width' => '120px'),
		array('name' => 'role_id', 'text' => '招募伙伴', 'width' => '120px'),
		array('name' => 'npc_role', 'text' => '招募伙伴NPC', 'width' => '120px'),
		array('name' => 'mission_level_id', 'text' => '任务关卡', 'width' => '120px'),
		array('name' => 'enemy_id', 'text' => '击杀怪物', 'width' => '120px'),
		array('name' => 'mission_enemy_id', 'text' => '任务关卡怪物组', 'width' => '120px'),
		array('name' => 'quest_item_id', 'text' => '任务关卡掉落物品', 'width' => '80px'),
		array('name' => 'quest_item_num', 'text' => '任务关卡掉落数量', 'width' => '40px'),
		array('name' => 'quest_item_rate', 'text' => '任务关卡掉落概率', 'width' => '40px'),
		array('name' => 'require_item_type', 'text' => '任务需求物品类型', 'width' => '40px'),
		array('name' => 'require_item_id', 'text' => '任务需求物品', 'width' => '200px'),
		array('name' => 'required_progress', 'text' => '任务要求进度', 'width' => '40px'),

		array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
		array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
		array('name' => 'award_item_1', 'text' => '奖励道具1', 'width' => '80px'),
		array('name' => 'award_num_1', 'text' => '奖励道具1数量', 'width' => '80px'),
		array('name' => 'award_item_2', 'text' => '奖励道具2', 'width' => '80px'),
		array('name' => 'award_num_2', 'text' => '奖励道具2数量', 'width' => '80px'),

		array('name' => 'conversion_reciving_quest', 'text' => '领取任务对话', 'width' => '180px'),
		array('name' => 'conversion_recived_quest', 'text' => '任务中对话', 'width' => '180px'),
		array('name' => 'conversion_finish_quest', 'text' => '领奖励对话', 'width' => '180px'),
	),


	'where' 		=> 'sql_where',
	'js'			=> 'jsFunction',
	'batch_update_key' => 'id',
	'batch_update_password' => '789',

);
?>
