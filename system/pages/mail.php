<?php

include "mail_common.php";

function sql_where(){
	return " WHERE `type` = 0 ORDER BY `sign`";
}

function sql_insert() {
	return ' `expire_time`=0 ';
}

function add_attachment_reader($row){
	return html_get_link("添加附件 | ", "?p=mail_attachments&m=".$row['id'], false);
}

function add_special_attachment_reader($row){
	return html_get_link("特殊附件", "?p=mail_special_attachments&m=".$row['id'], false);
}


$pconfig = array(
	'title'   => '邮件',
	'table'   => 'mail',
	'links'   => $links,
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'columns' => array(
		array('name' => 'title', 'text' => '标题', 'width' => '150px'),
		array('name' => 'sign', 'text' => '唯一标识', 'width' => '90px'),
		array('name' => 'type', 'text' => '类型', 'width' => '60px',
			'editor' => 'mail_type_editor',
			'render' => 'mail_type_render',
		),
		array('name' => 'priority', 'text' => '优先级', 'width' => '60px',),
		//array('name' => 'title', 'text' => '标题', 'width' => '150px'),
		//array('name' => 'expire_time', 'text' => '过期时机', 'width' => '60px',
		//	'editor' => 'mail_expire_time_editor',
		//	'render' => 'mail_expire_time_render',
		//),
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

	'opera' => array(
		array('type' => 'link', 'text' => '添加附件', 'render' => 'add_attachment_reader'),
		array('type' => 'link', 'text' => '特殊附件', 'render' => 'add_special_attachment_reader'),
	),
	'not_delete' => true,

);
?>
