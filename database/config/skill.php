<?php
$config = array(
	array( 'const',
		/*角色身上技能位置*/
		'POS_SKILL_1' => 1,  // 基础技能槽
		'POS_SKILL_2' => 2,  // 奥义技能槽
		'POS_SKILL_3' => 3,  // 体术技能槽
		'POS_SKILL_4' => 4,  // 咒术技能槽

		/*角色身上技能槽可装备技能类型 参见 TYPE_XXX 常量*/
		'SKILL_KIND_FOR_POS_1' => 1,  // 基础技
		'SKILL_KIND_FOR_POS_2' => 3,  // 体术技
		'SKILL_KIND_FOR_POS_3' => 4,  // 咒术技
		'SKILL_KIND_FOR_POS_4' => 2,  // 奥义技

		'MAX_SKILL_LEVEL' => 100, //最大技能训练等级

		'MAX_SKILL_NUM' => 100, //技能总数

		/* 战斗相关 */
		'SKILL_IS_NULL' => 0,

		/* 技能所属类型 对应 `skill`.`type` 字段*/
		'OWNER_SYS'	  => 0, // 技能不在客户端展示（特殊技能）
		'OWNER_ROLE'  => 1,  // 玩家角色技能
		'OWNER_ENEMY' => 5,  // 怪物技能
		'OWNER_GHOST' => 7,  // 魂侍技能

		/*技能子类型*/
		'SKILL_KIND_BASIC'       => 1,  //基础
		'SKILL_KIND_ULTIMATE'    => 2,  //奥义
		'SKILL_KIND_ADVANCE'     => 3,  //体术
		'SKILL_KIND_CURSE'       => 4,  //咒术

		/*伙伴的进阶技能*一半*有次数限制， -1 代表这种进阶技能是可以无限使用的*/
		'UNLIMIT_BUDDY_SKILL' => -1 ,
		

		/*技能作用*/
		'SKILL_CHILD_TYPE_ATTACK'  => 1,  //进攻
		'SKILL_CHILD_TYPE_DENFEND' => 3,  //防御
		'SKILL_CHILD_TYPE_HEAL'    => 4,  //治疗
		'SKILL_CHILD_TYPE_SUPPORT' => 5,  //辅助
		'SKILL_CHILD_TYPE_SUNDER'  => 6,  //破甲

		/*角色类型*/
		'SKILL_MAIN_ROLE_ID' => -2, //绝招表中主角id
		'SKILL_ENEMY_ID'     => -1, //绝招表中敌人id

		/*主角默认技能*/
		'SKILL_SHANJI' => 1,           //闪击
		'SKILL_JINGANG_SHANJI' => 90,  //金刚闪击
		'SKILL_DAOXUAN_SHANJI' => 93,  //道玄闪击
		'SKILL_WANXIANG_SHANJI' => 96, //万象闪击
		'SKILL_TIANREN_SHANJI' => 99,  //天人闪击
		'SKILL_BISHAJI' => 4, //必杀技,特殊绝招，15级获得可升级

		/*绝招攻击方式*/
		'SKILL_TARGET_MODE_SINGLE'                 => 0, //单体
		'SKILL_TARGET_MODE_ALL'                    => 1, //全体
		'SKILL_TARGET_MODE_HORIZONTAL'             => 2, //横向
		'SKILL_TARGET_MODE_VERTICAL'               => 3, //纵向
		'SKILL_TARGET_MODE_STILL'                  => 4, //不攻击
		'SKILL_TARGET_MODE_HORIZONTAL_BACK_ROW'    => 5, //后排横向
		'SKILL_TARGET_MODE_MIN_HP'                 => 6, //血量最少
		'SKILL_TARGET_MODE_TWICE_SINGLE_SAME'      => 7, //单体重复两次
		'SKILL_TARGET_MODE_PASSIVE'                => 8, //被动技
		'SKILL_TARGET_MODE_VERTICAL_EXTRA'         => 9, //纵向攻击，并额外随机攻击另一列
		'SKILL_TARGET_MODE_TWICE_SINGLE_RANDOM'    => 10, //随机攻击两个目标
		'SKILL_TARGET_MODE_SINGLE_EXTRA'           => 11, //单体攻击，并额外随机攻击另一人
		'SKILL_TARGET_MODE_MAX_HP'                 => 12, //血量最多
		'SKILL_TARGET_MODE_V_ATTACK'               => 13, //V字攻击
		'SKILL_TARGET_MODE_SINGLE_BACK_ROW'        => 14, //后排单体攻击
		'SKILL_TARGET_MODE_CROSS'                  => 15, //十字攻击
		'SKILL_TARGET_MODE_HORIZONTAL_PENETRATION' => 16, //横排穿透
		'SKILL_TARGET_MODE_CROSS_2'                => 17, //十字攻击(固定)

		/*绝招品质*/
		'SKILL_QUALITY_BLUE'   => 0,
		'SKILL_QUALITY_PURPLE' => 1,
		'SKILL_QUALITY_GOLDEN' => 2,


		/*攻击目标，客户端表现使用*/
		'SKILL_TARGET_SINGLE'  => 1, //单体
		'SKILL_TARGET_ALL'  => 2, //全体
		'SKILL_TARGET_HORIZON'  => 3, //横向
		'SKILL_TARGET_VERTICAL'  => 4, //纵向
		'SKILL_TARGET_ENEMY_FRONTEND_ROW'  => 5, //敌方前排固定
		'SKILL_TARGET_ENEMY_MIDDLE_ROW'  => 6, //敌方中排固定
		'SKILL_TARGET_ENEMY_BACKEND_ROW'  => 7, //敌方后排固定
		'SKILL_TARGET_BUDDY_FRONTEND_ROW'  => 8, //我方前排固定
		'SKILL_TARGET_BUDDY_MIDDLE_ROW'  => 9, //我中排固定
		'SKILL_TARGET_BUDDY_BACKEND_ROW'  => 10, //我方后排固定

		'SKILL_CAN_AUTO_LEARN' => 1, //技能到等级和自动学习

		"FLUSH_SKILL_CD" => 30*60, // 角色技能洗点冷却时间
		"FLUSH_SKILL_BACK_PERCENT" => _S(1.0, 1.0, 0.8, 0.8) // 角色技能洗点冷却时间
	),
);
?>
