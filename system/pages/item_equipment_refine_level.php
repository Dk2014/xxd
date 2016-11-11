<?php

include "item_equipment_links.php";

$pconfig = array(
	'title'   => '装备精练等级',
	'table'   => 'equipment_refine_level',
	'links'   => $links,
	
	'columns' => array(
		array('name' => 'level', 'text' => '精练级别', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
		array('name' => 'probability', 'text' => '精练成功概率（%）', 'width' => '80px'),
		array('name' => 'gain_pct', 'text' => '精炼系数（%）', 'width' => '80px'),
	),
);