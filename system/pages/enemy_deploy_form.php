<?php


$pconfig = array(
	'title'   => '怪物阵法站位表',
	'table'   => 'enemy_deploy_form',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'pos1', 'hide' => true, 'text' => '位置1上的敌人ID', 'width' => '80px'),
		array('name' => 'pos2', 'hide' => true, 'text' => '位置2上的敌人ID', 'width' => '80px'),
		array('name' => 'pos3', 'hide' => true, 'text' => '位置3上的敌人ID', 'width' => '80px'),
		array('name' => 'pos4', 'hide' => true, 'text' => '位置4上的敌人ID', 'width' => '80px'),
		array('name' => 'pos5', 'hide' => true, 'text' => '位置5上的敌人ID', 'width' => '80px'),
		array('name' => 'pos6', 'hide' => true, 'text' => '位置6上的敌人ID', 'width' => '80px'),
		array('name' => 'pos7', 'hide' => true, 'text' => '位置7上的敌人ID', 'width' => '80px'),
		array('name' => 'pos8', 'hide' => true, 'text' => '位置8上的敌人ID', 'width' => '80px'),
		array('name' => 'pos9', 'hide' => true, 'text' => '位置9上的敌人ID', 'width' => '80px'),
		array('name' => 'pos10', 'hide' => true, 'text' => '位置10上的敌人ID', 'width' => '80px'),
		array('name' => 'pos11', 'hide' => true, 'text' => '位置11上的敌人ID', 'width' => '80px'),
		array('name' => 'pos12', 'hide' => true, 'text' => '位置12上的敌人ID', 'width' => '80px'),
		array('name' => 'pos13', 'hide' => true, 'text' => '位置13上的敌人ID', 'width' => '80px'),
		array('name' => 'pos14', 'hide' => true, 'text' => '位置14上的敌人ID', 'width' => '80px'),
		array('name' => 'pos15', 'hide' => true, 'text' => '位置15上的敌人ID', 'width' => '80px'),

	),
	
	'opera' => array(
		array('text' => '站位', 'render' => 'enemy_group_form_opera'),
	),
	
	'js' 			=> 'jsFunction',
	'where'	  		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'location' 		=> 'location',
	'not_template' => true,
);


?>