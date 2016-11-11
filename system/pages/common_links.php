<?php

// 品质
$qualityTypes[-1] = '无';
$qualityTypes[0] = '蓝色';
$qualityTypes[1] = '紫色';
$qualityTypes[2] = '金色';
$qualityTypes[3] = '橙色';

// 绝招类型
$skillKind[0] = '无';
$skillKind[1] = '基础';
$skillKind[2] = '奥义';
$skillKind[3] = '体术';
$skillKind[4] = '咒术';


// 绝招子类型 
// 改变请 grep battle/battle.go SKILL_KIND_
$skillTypes[1] = '进攻';
$skillTypes[3] = '防御';
$skillTypes[4] = '治疗';
$skillTypes[5] = '辅助';
$skillTypes[6] = '破甲';

// 二进制下拉框 (ghost 唯一值)
$uniqueKeyConfig = array(
	0        => 0,
	pow(2,0) => 1,
	pow(2,1) => 2,
	pow(2,2) => 3,
	pow(2,3) => 4,
	pow(2,4) => 5,
	pow(2,5) => 6,
);

// 是否可以在影界获得
$canMissionGet['0']  = '不能在影界获得';
$canMissionGet['1']  = '可以在影界获得';

// 魂侍
$ghost_links = array(
	array('text' => '魂侍', 'id' => 'ghost'),
	array('text' => '魂侍等级', 'id' => 'ghost_level'),
	array('text' => '魂侍进阶', 'id' => 'ghost_star'),
	array('text' => '魂侍被动技能', 'id' => 'ghost_passive_skill'),
	array('text' => '魂侍提示信息', 'id' => 'ghost_tip'),
	array('text' => '技能训练价格', 'id' => 'ghost_skill_train_price'),
	array('text' => '魂侍连锁', 'id' => 'ghost_relationship'),
	array('text' => '魂侍洗练', 'id' => 'ghost_baptize'),
);

// vip、体力、招财神符的共用链接
$price_links = array(
	array('text' => 'VIP信息', 'id' => 'vip_info'),
	array('text' => '招财神符获得铜钱', 'id' => 'rune_coins'),
	array('text' => '招财神符价格', 'id' => 'rune_price'),
	array('text' => '体力价格', 'id' => 'physical_price'),
	array('text' => '魂侍培养价格', 'id' => 'ghost_train_price'),
	array('text' => '金蛋价格', 'id' => 'ghost_egg_price'),
	array('text' => '影界价格', 'id' => 'ghost_mission_price'),
	array('text' => '斩妖奖励次数价格', 'id' => 'chop_goblin_price'),
	array('text' => '感悟价格', 'id' => 'gnosis_price'),
	array('text' => '元宝作画价格', 'id' => 'painting_price'),
    array('text' => '元宝寻访价格', 'id' => 'skill_visit_price'),
);

// 剑心
$sword_soul_links = array(
	array('text' => '剑心', 'id' => 'sword_soul'),
	array('text' => '剑心类型', 'id' => 'sword_soul_type'),
	array('text' => '剑心品质', 'id' => 'sword_soul_quality'),
    array('text' => '剑山拔剑概率', 'id' => 'sword_soul_probability')
);

// 角色境界
$role_realm_links = array(
	array('text' => '角色境界等级', 'id' => 'role_realm_level',),
    array('text' => '角色境界阶级', 'id' => 'role_realm_class',),	
);

$display_params[1] = '黑幕';
$display_params[2] = '震动';

//阵印链接
$totem_links = array(
	array('text' => '阵印', 'id' => 'totem'),
	array('text' => '阵印元宝召唤', 'id' => 'totem_call_cost_config'),
	array('text' => '阵印等级数据', 'id' => 'totem_level_info'),
);

$skill_target_type = array(1=>'单体', 2=>'全体', 3=>'横向', 4=>'纵向', 5=>'敌方前排固定', 6=>'敌方中排固定', 7=>'敌方后排固定',
	8=>'我方前排固定', 9=>'我方中排固定', 10=>'我方后排固定');
?>
