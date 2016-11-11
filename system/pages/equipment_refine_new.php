<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备强化',
	'table'   => 'equipment_refine_new',
	'links'   =>$links,
	
	'columns' => array(
		array('name' => 'equ_type', 'text' => '装备类型', 'width' => '180px'),
		array('name' => 'equ_kind', 'text' => '装备种类', 'width' => '180px'),
		array('name' => 'base_price', 'text' => '装备基础价格', 'width' => '100px'),
		array('name' => 'incre_val', 'text' => '强化提升单位属性', 'width' => '180px'),
		array('name' => 'incre_price', 'text' => '强化提升单位价格', 'width' => '180px'),

	),
);

?>
