<?php

include "role_common.php";

$pconfig = array(
	'title'   => '伙伴协力',
	'table'   => 'buddy_cooperation',
	'links'   =>$role_links,
        	
	'columns' => array(
		array('name' => 'name', 'text' => '组合名', 'width' => '150px'),
		array('name' => 'desc', 'text' => '描述', 'width' => '150px'),
		array('name' => 'role_id1', 'text' => '角色1', 'width' => '150px'),
		array('name' => 'role_id2', 'text' => '角色2', 'width' => '150px'),
		array('name' => 'require_friendship_level', 'text' => '羁绊等级', 'width' => '150px'),
		
		array('name' => 'skill_level', 'text' => '技能熟练度加成', 'width' => '40px'),
		array('name' => 'ghost_power', 'text' => '初始魂力','width' => '60px'),

		array('name' => 'health', 'text' => '生命加成', 'width' => '40px'),
		array('name' => 'attack', 'text' => '攻击加成', 'width' => '40px'),
		array('name' => 'defence', 'text' => '防御加成', 'width' => '40px'),
    		array('name' => 'speed',   'text' => '速度加成', 'width' => '60px'),
    		array('name' => 'sunder',   'text' => '护甲加成', 'width' => '60px'),
    		array('name' => 'cultivation',   'text' => '内力加成', 'width' => '60px'),

		array('name' => 'tenacity_level', 'text' => '韧性等级','width' => '60px'),
		array('name' => 'destroy_level', 'text' => '破击等级','width' => '60px'),
		array('name' => 'block_level',   'text' => '格挡等级', 'width' => '60px'),
		array('name' => 'hit_level',   'text' => '命中等级', 'width' => '60px'),
		array('name' => 'dodge_level',   'text' => '闪避等级', 'width' => '60px'),
    		array('name' => 'critical_level',   'text' => '暴击等级', 'width' => '60px'),
		array('name' => 'critical_hurt_level',   'text' => '必杀等级', 'width' => '60px'),
	),
	
	'where' => 'sql_where',
	'js'	        => 'js_function',
);
