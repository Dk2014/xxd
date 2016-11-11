<?php

$pconfig = array(
	'title'   => '推送通知',
	'table'   => 'push_notify',
	'links'   => array(),
	'columns' => array(
		array('name' => 'sign', 'text' => '唯一标示', 'width' => '150px'),
		array('name' => 'name', 'text' => '推送通知名称', 'width' => '150px'),
		array('name' => 'content', 'text' => '内容', 'width' => '200px'),
		array('name' => 'trigger_time', 'text' => '一天中的第几秒[0,86400) -1非定时推送', 'width' => '90px'),
	),
);
