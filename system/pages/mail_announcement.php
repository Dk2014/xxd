<?php

include "mail_common.php";

function sql_where(){
	return " WHERE `type` = 1 ORDER BY `sign`";
}

$pconfig = array(
	'title'   => '模块公告',
	'table'   => 'mail',
	'links'   => $links,
	'where'	  => 'sql_where',
	'columns' => array(
		array('name' => 'title', 'text' => '标题', 'width' => '150px'),
		array('name' => 'sign', 'text' => '唯一标识', 'width' => '90px'),
		array('name' => 'type', 'text' => '类型', 'width' => '60px',
			'editor' => 'mail_type_editor',
			'render' => 'mail_type_render',
		),
		array(
			'name' => 'parameters', 'text' => '参数', 'width' => '200px',
			'editor' => 'parameters_editor',
			'render' => 'parameters_render',
		),
		array('name' => 'content', 'text' => '内容', 'width' => '200px',
			'editor' => 'content_editor',
			'render' => 'content_render',
		),
		array('name' => 'announcement_secs', 'text' => '公告持续秒数', 'width' => '90px'),
	),
);

?>
