<?php

$battle_item_target = array(
	'1'=> array('name' => '道具使用者', 'code'=>'[]*Fighter{user}'),
	
	'2'=> array('name' => '我方全体', 'code'=>'findSideLiveFighter(user.battle.Attackers.Fighters)'),
	'3'=> array('name' => '我方血量最少的', 'code'=>'[]*Fighter{findLeastHealthFighter(user.battle.Attackers.Fighters)}'),

	'4'=> array('name' => '敌方全体', 'code'=>'findSideLiveFighter(user.battle.Defenders.Fighters)'),
	'5'=> array('name' => '敌方血量最少的', 'code'=>'[]*Fighter{findLeastHealthFighter(user.battle.Defenders.Fighters)}'),
	
	'6'=> array('name' => '我方阵亡角色', 'code'=>'[]*Fighter{findDeadFighter(user.battle.Attackers.Fighters)}'),
	'7'=> array('name' => '灵宠', 'code'=>'[]*Fighter{findBattlePet(user.battle.Defenders.Fighters)}'),
	'8'=> array('name' => '我方伙伴', 'code'=>'[]*Fighter{findOurBuddies(user)}'),
	'9'=> array('name' => '主角和伙伴', 'code'=>'findOurSideWithoutPet(user)'),
	'10'=> array('name' => '敌方单体', 'code'=>'findOneTarget(user)'),
);

$battle_item_effect_type = array(
	'1'=>array('name' => '基础属性', 'call' => 'configNormalCode'),
	'2'=>array('name' => '捕捉球','call' => 'catchCode'),	
	'3'=>array('name' => 'Buff增益','call' => 'configBuffCode'),
);

$battle_item_effect_normal_mod = array(
	'0'=> array('name' => '精气', 'var' => 'Power', 'max' => 'target.MaxPower'), // var对应battle.Fighter基础属性变量
	'1'=> array('name' => '生命', 'var' => 'Health', 'max' => 'target.MaxHealth', 'min'=>'MIN_HEALTH'),
	'3'=> array('name' => '复活', 'call'=>'reliveFighter'),// 对应的值为恢复生命的%
	'4'=> array('name' => '捕捉概率(%)','var' => '', 'call'=>'catchPet'),
);

// 对应buff类型
$battle_item_effect_buff_mod = array(
	'0'=>array('name'=>'精气', 'const_var'=>'BUFF_POWER'),
	'4'=>array('name'=>'生命','const_var'=>'BUFF_HEALTH'),
	'7'=>array('name'=>'清除负面状态', 'const_var'=>'BUFF_CLEAN_BAD'),
	'13'=>array('name'=>'闪避等级','const_var'=>'BUFF_DODGE_LEVEL'),
	'16'=>array('name'=>'伤害加成(%)','const_var'=>'BUFF_HURT_ADD'),
	'29'=>array('name'=>'伙伴技能','const_var'=>'BUFF_BUDDY_SKILL'),
	'36'=>array('name'=>'回复伙伴技能','const_var'=>'BUFF_RECOVER_BUDDY_SKILL'),
	'37'=>array('name'=>'增加精气到满','const_var'=>'BUFF_MAKE_POWER_FULL'),
);
?>
