<?php

$pconfig = array(
	'title'   => '命锁',
	'table'   => 'fate_box',
	'links'   => array(),

	'columns' => array(
		array('name' => 'name', 'text' => '名称', 'width' => '200px'),
		array('name' => 'sign', 'text' => '唯一表识', 'width' => '200px'),
		array('name' => 'level', 'text' => '要求等级', 'width' => '100px'),
		array('name' => 'require_lock', 'text' => '要求命锁权值', 'width' => '100px'),
		array('name' => 'award_lock', 'text' => '奖励命锁权值', 'width' => '100px'),
		array('name' => 'fixed_prop', 'text' => '固定道具', 'width' => '100px'),
	),


	'opera' => array(
		array('type' => 'link', 'text' => '关卡', 'render' => 'operation'),
	),

  'not_delete' 	=> true,

  'where' => 'sql_where',
  'js' => 'jsFunction',
);
?>
