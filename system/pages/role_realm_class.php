<?php
include "common_links.php";

$pconfig = array(
	'title'   => '角色境界阶级',
	'table'   => 'role_realm_class',
	'links'   => $role_realm_links,
	'columns' => array(
		array('name' => 'name',      'text' => '名称',        'width' => '60px'),
		array('name' => 'realm_class',      'text' => '境界阶级',        'width' => '60px'),
		array('name' => 'need_realm_level', 'text' => '突破至下一阶所需境界等级', 'width' => '120px'),
		array('name' => 'item_id',          'text' => '突破至下一阶所需道具',    'width' => '100px',),
		array('name' => 'item_num',                 'text' => '道具数量',    'width' => '60px'),
		array('name' => 'add_hth',                  'text' => '生命加成',    'width' => '50px'),
		array('name' => 'add_sunder_value',         'text' => '护甲加成',    'width' => '50px'),
		array('name' => 'add_cultivation',          'text' => '内力加成',    'width' => '50px'),
		array('name' => 'add_power',                'text' => '初始精气加成', 'width' => '50px'),
		array('name' => 'add_max_power',            'text' => '精气上限加成', 'width' => '50px'),
		array('name' => 'add_sunder_min_hurt_rate', 'text' => '破甲前免伤',   'width' => '50px'),
		array('name' => 'add_aoe_reduce',           'text' => '范围免伤加成', 'width' => '50px'),
		array('name' => 'add_critial_level',        'text' => '暴击等级加成', 'width' => '50px'),
		array('name' => 'add_dodge_level',          'text' => '闪避等级加成', 'width' => '50px'),
		array('name' => 'add_hit_level',            'text' => '命中等级加成', 'width' => '50px'),
		array('name' => 'add_block_level',          'text' => '格挡等级加成', 'width' => '50px'),
		array('name' => 'add_tenacity_level',       'text' => '韧性等级加成', 'width' => '50px'),
		array('name' => 'add_destroy_level',        'text' => '破击等级加成', 'width' => '50px'),
		array('name' => 'desc',                     'text' => '简介', 'width' => '200px'),

	),
	'where' => 'sql_where',
);
?>