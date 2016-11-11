<?php
include "common_links.php";

function sql_where($params) {
	return " order by `realm_level`";
}

$pconfig = array(
	'title'   => '角色境界等级',
	'table'   => 'role_realm_level',
	'links'   => $role_realm_links,
	'columns' => array(
		array('name' => 'realm_level',      'text' => '境界等级',    'width' => '60px'),
		array('name' => 'exp',              'text' => '升级所需经验', 'width' => '100px'),
		array('name' => 'need_realm_class', 'text' => '升级所需阶级', 'width' => '100px'),
		array('name' => 'item_num',         'text' => '道具数量',    'width' => '100px'),

		array('name' => 'add_health', 'text' => '增加生命', 'width' => '100px'),
		array('name' => 'add_attack', 'text' => '增加攻击', 'width' => '100px'),
		array('name' => 'add_defence', 'text' => '增加防御', 'width' => '100px'),
		array('name' => 'add_cultivation', 'text' => '增加内力', 'width' => '100px'),
		array('name' => 'add_speed', 'text' => '增加速度'),
	),
	'where' => 'sql_where',
);
?>