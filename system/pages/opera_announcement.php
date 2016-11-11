<?php

include "mail_common.php";
$announcementLinks = array(
	array('text' => '跑马灯公告', 'id' => 'announcement'),
	array('text' => '活动公告', 'id' => 'opera_announcement'),
);




$pconfig = array(
	'title'   => '活动公告',
	'table'   => 'opera_announcement',
	'links'   => $announcementLinks,
	'columns' => array(
		array('name' => 'title', 'text' => '公告标题', 'width' => '150px'),
		array('name' => 'content', 'text' => '内容', 'width' => '200px'),
		array('name' => 'effect_time', 'text' => '生效时间戳', 'width' => '90px'),
		array('name' => 'expire_time', 'text' => '结束时间(0永远有效）', 'width' => '90px'),
	),
);
