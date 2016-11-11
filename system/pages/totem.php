<?php

include "common_links.php";
include "skill_links.php";

function totem_opera($row) {
	return html_get_link("技能配置", "?p=totem_skill&m=".$row['id'], false);
}

$pconfig = array(
	'title' => '阵印',
	'table' => 'totem',

	'columns' => array(
		array('name' => 'name', 'text' => '名称', 'width' => '80px'),
		array('name' => 'sign', 'text' => '资源标示', 'width' => '80px'),
		array('name' => 'music_sign', 'text' => '音乐资源标示', 'width' => '80px'),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px'),
	),
	'js' => 'js_function',
	'links' => $totem_links,
	'opera' => array(
		array('type' => 'link', 'text' => '技能设置', 'render' => 'totem_opera'),
	),

);
