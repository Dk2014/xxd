<?php

include "role_common.php";

$pconfig = array(
	'title'   => '主角协力',
	'table'   => 'main_role_cooperation',
	'links'   =>$role_links,
        	
	'columns' => array(
		array('name' => 'buddy_num', 'text' => '伙伴人数', 'width' => '150px'),
		
		array('name' => 'health', 'text' => '生命加成', 'width' => '40px'),
		array('name' => 'attack', 'text' => '攻击加成', 'width' => '40px'),
		array('name' => 'defence', 'text' => '防御加成', 'width' => '40px'),
    		array('name' => 'speed',   'text' => '速度加成', 'width' => '60px'),
    		array('name' => 'cultivation',   'text' => '内力加成', 'width' => '60px'),

	),
);
