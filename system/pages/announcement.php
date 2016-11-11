<?php

include "mail_common.php";
$announcementLinks = array(
	array('text' => '跑马灯公告', 'id' => 'announcement'),
	array('text' => '活动公告', 'id' => 'opera_announcement'),
);




function sql_where($params){
	return "where `type`=0 ";
}

function sql_insert($params) {
	return " `type='0'";
}


$pconfig = array(
	'title'   => '公告模板',
	'table'   => 'announcement',
	'links'   => $announcementLinks,
	'where'	  => 'sql_where',
	'columns' => array(
		array('name' => 'name', 'text' => '公告名称', 'width' => '150px'),
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
		array('name' => 'duration', 'text' => '公告有效时间（秒）', 'width' => '90px'),
		array('name' => 'show_cyle', 'text' => '重复展示间隔（秒）', 'width' => '90px'),
	),
);
/*

// 公告类型
$types[0] = '模块公告';
$types[1] = '后台公告';

function announcement_type_render($row){
	global $types;
	return $types[$row['type']];
}

function announcement_type_editor($row){
	global $types;

	$code = '';
	if ($row != null) {
		$code .= '<select name="type[]" >';
	foreach ($types as $key => $value){

		if ($key == $row['type']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="type[]" >';
		$code .= '<option value="0" selected="selected">模块公告</option>';
		$code .= '<option value="1" >后台公告</option>';
		$code .= '</select>';
	}
	
	return $code;
}

?>

*/
