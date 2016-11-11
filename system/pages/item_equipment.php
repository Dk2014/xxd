<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备数据',
	'table'   => 'item',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'name', 'text' => '装备名称', 'width' => '150px'),
		array('name' => 'type_id', 'text' => '类型ID','width' => '70px',),
		array('name' => 'equip_type_id', 'text' => '等级类型','width' => '70px',),
		array('name' => 'quality', 'text' => '品质','width' => '50px',),
		array('name' => 'equip_role_id', 'text' => '限定装备角色','width' => '50px',),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'price', 'text' => '价格', 'width' => '50px'),
		array('name' => 'level', 'text' => '需求等级', 'width' => '60px'),
		array('name' => 'refine_base', 'text' => '精炼值', 'width' => '60px'),
		array('name' => 'health', 'text' => '生命', 'width' => '50px'),
		array('name' => 'speed', 'text' => '速度', 'width' => '50px'),
		array('name' => 'attack', 'text' => '攻击', 'width' => '50px'),
		array('name' => 'defence', 'text' => '防御', 'width' => '50px'),
		array('name' => 'appendix_num', 'text' => '追加属性数', 'width' => '50px'),
		array('name' => 'appendix_level', 'text' => '追加属等级', 'width' => '50px'),
		array('name' => 'desc', 'text' => '装备描述', 'width' => '80px'),
	),

	'where' 		=> 'sql_where',
	'not_delete' => true,
);
?>
