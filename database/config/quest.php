<?php

//任务状态统一使用主线任务的常量
// 0 -- 未领取
// 1 -- 已领取
// 2 -- 已完成但未领取奖励
// 3 -- 已领取奖励
//
// TODO 
// 伙伴任务 extend_quest 的返回统一的任务状态



$config = array(
	array( 'const',
		// 任务状态
		'QUEST_STATUS_NO_RECEIVE'      => 0, // 任务状态未接
		'QUEST_STATUS_ALREADY_RECEIVE' => 1, // 任务状态已接
		'QUEST_STATUS_FINISH'          => 2, // 任务状态已完成
		'QUEST_STATUS_AWARD'           => 3, // 任务状态已奖励

		'QUEST_ID_NO_MORE' => 0, //没有任务的任务ID

		'QUEST_TYPE_NONE'		   => -1,
		'QUEST_TYPE_TOWN'          => 0, // 进入城镇任务
		'QUEST_TYPE_TOWN_NPC'      => 1, // 城镇NPC任务 (跟指定NPC对话完成)
		'QUEST_TYPE_MISSION'       => 2, // 进入副本任务 (进入副本完成)
		'QUEST_TYPE_WIN_MISSION'   => 3, // 副本通关任务 (通关副本完成)
		'QUEST_TYPE_MISSION_ENEMY' => 4, // 消灭敌人任务


		'DAILY_QUEST_TYPE_MISSION_LEVEL' 	=> 0, // 区域关卡
		'DAILY_QUEST_TYPE_RESOURCE_LEVEL' 	=> 1, // 资源关卡
		'DAILY_QUEST_TYPE_LIMIT_LEVEL' 		=> 2, // 极限关卡
		'DAILY_QUEST_TYPE_HARD_LEVEL' 		=> 8, // 难度关卡
		'DAILY_QUEST_TYPE_BUDDY_LEVEL'	 	=> 9, // 伙伴关卡
		'DAILY_QUEST_TYPE_PET_LEVEL' 		=> 10, // 灵宠关卡
		'DAILY_QUEST_TYPE_GHOST_LEVEL' 		=> 11, // 魂侍关卡
		'DAILY_QUEST_TYPE_PET_VIRENV' 		=> 12, // 灵宠幻境


		'DAILY_QUEST_STATUS_NONE' => -1,
		'DAILY_QUEST_STATUS_CAN_AWARD' => 0,
		'DAILY_QUEST_STATUS_HAVE_AWARD' => 1,


		'DAILY_QUEST_CLASS_MISSION_LEVEL_BOSS' 	=> 0, // 区域BOSS关卡
		'DAILY_QUEST_CLASS_RESOURCE_COIN_LEVEL' => 1, // 资源铜钱关卡
		'DAILY_QUEST_CLASS_RESOURCE_EXP_LEVEL' 	=> 2, // 资源经验关卡
		'DAILY_QUEST_CLASS_LIMIT_LEVEL' 		=> 3, // 极限关卡
		'DAILY_QUEST_CLASS_HARD_LEVEL' 			=> 4, // 难度关卡
		'DAILY_QUEST_CLASS_BUDDY_LEVEL'	 		=> 5, // 伙伴关卡
		'DAILY_QUEST_CLASS_PET_LEVEL' 			=> 6, // 灵宠关卡
		'DAILY_QUEST_CLASS_GHOST_LEVEL' 		=> 7, // 魂侍关卡
		'DAILY_QUEST_CLASS_AREAN'				=> 8, // 比武场
		'DAILY_QUEST_CLASS_MULTI_LEVEL'			=> 9, // 多人关卡

		'DAILY_QUEST_CLASS_DRAW_INGOT_BOX'		=> 10, // 神龙宝箱
		'DAILY_QUEST_CLASS_DRAW_HEART_BOX'		=> 11, // 爱心宝箱
		'DAILY_QUEST_CLASS_SEND_FRIEND_HEART'	=> 12, // 赠送爱心
		'DAILY_QUEST_CLASS_DRAW_SWORD'			=> 13, // 拔剑
		'DAILY_QUEST_CLASS_TRAIN_GHOST'			=> 14, // 魂侍培养
		'DAILY_QUEST_CLASS_REFINE_EQ'			=> 15, // 精炼装备
		'DAILY_QUEST_CLASS_RECAST_EQ'			=> 16, // 重铸装备
		'DAILY_QUEST_CLASS_PET_VIRENV'			=> 17, // 灵宠幻境
		'DAILY_QUEST_CLASS_GRID_UPGRADE'		=> 18, // 契约之力
		'DAILY_QUEST_CLASS_LEVELUP_PET'         => 19, // 灵宠培养
		'DAILY_QUEST_CLASS_SKILL_TRAINING'         => 20, // 绝招训练
		'DAILY_QUEST_CLASS_TEAMSHIP_TRAINING'         => 21, //团队友情
		'DAILY_QUEST_CLASS_RECHARGE' 			=> 22, //每日充值
		'DAILY_QUEST_CLASS_CALL_TOTEM' 			=> 23, //召唤阵印
		'DAILY_QUEST_CLASS_DRIVING_SWORD' 			=> 24, //云海愈剑

		'EXTEND_QUEST_STATUS_NONE' =>0, //未完成
		'EXTEND_QUEST_STATUS_CAN_AWARD' =>1, //已完成
		'EXTEND_QUEST_STATUS_AWARDED' =>2, //已奖励

		'EXTEND_QUEST_TYPE_MISSION_START' => 1, //区域评星
		'EXTEND_QUEST_TYPE_LOGIN' => 2, //连续登陆
		'EXTEND_QUEST_TYPE_PAY_INGOT' => 3, //元宝购买

		'RECRUIT_BUDDY_REQUIRED_QUEST' => 9999, //招募伙伴图标开启任务

		'ADDITION_QUEST_TYPE_CONVERSION' => 1, //NPC对话
		'ADDITION_QUEST_TYPE_KILL_ENEMY' => 2, //消灭敌人
		'ADDITION_QUEST_TYPE_PASS_LEVEL' => 3, //通过关卡
		'ADDITION_QUEST_TYPE_COLLECT_ITEM' => 4, //收集物品
		'ADDITION_QUEST_TYPE_SHOW_ITEM' => 5, //展示物品
		'ADDITION_QUEST_TYPE_MISSION_STAR' => 6, //区域评星
		'ADDITION_QUEST_TYPE_RECUIT_BUDDY' => 7, //招募伙伴

		'ADDITION_QUEST_STATE_NOT_RECIVE' => 0, //未领取
		'ADDITION_QUEST_STATE_RECIVED' => 1, //已领取
		'ADDITION_QUEST_STATE_NOT_AWARD' => 2, //已完成 未奖励
		'ADDITION_QUEST_STATE_AWARDED' => 3, //已奖励

		//任务大类
		'QUEST_CLASS_MAIN' =>  1, // 主线任务
		'QUEST_CLASS_DAILY' => 2, // 每日任务
		'QUEST_CLASS_EXTEND' => 3, // 伙伴招募任务
		'QUEST_CLASS_ADDITION' => 4, // 支线任务

		'QUEST_REQUIRE_ITEM_TYPE_EQUIPMENT' => 1, //装备
		'QUEST_REQUIRE_ITEM_TYPE_ITEM' => 2, //非装备物品

		'MAX_ADDITION_QUEST_NUM' => 5, //玩家最多同时领取5个任务

		'MAIN_QUEST_AWARDE_STARS' => 10, //主线任务获得星数奖励
		'ADDITION_QUEST_AWARDE_STARS' => 5, //支线任务获得星数奖励


	),
);

$db = db_connect();
$row = db_query($db, "select * from 
	(select n.showup_quest as quest_id , n.id as npc_id from recruit_buddy r 
	left join town_npc n on r.related_npc=n.`id`) qid 
	left join (select id, `order` from quest ) q on qid.quest_id=q.id 
	order by `order` limit 1;");
if(isset($row[0]) && isset($row[0]['quest_id'])) {
	$config[0]['RECRUIT_BUDDY'] = $row[0]['quest_id'];
}

?>
