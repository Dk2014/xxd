<?php

include "mail_common.php";

$pconfig = array(
	'title'   => '世界频道消息模板',
	'table'   => 'world_channel_message',
	'links'   => array(),
	'columns' => array(
		array('name' => 'desc', 'text' => '描述', 'width' => '150px'),
		array('name' => 'sign', 'text' => '唯一标识', 'width' => '150px'),
		array(
			'name' => 'parameters', 'text' => '参数', 'width' => '200px',
			'editor' => 'parameters_editor',
			'render' => 'parameters_render',
		),
		array('name' => 'content', 'text' => '内容', 'width' => '200px',
			'editor' => 'content_editor',
			'render' => 'content_render',
		),
	),
);
