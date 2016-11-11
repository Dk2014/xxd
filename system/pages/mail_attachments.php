 <?php

 include "mail_common.php";

$links = array(
	array('text' => '邮件', 'id' => 'mail'),
	array('text' => '模块公告', 'id' => 'mail_announcement'),
);

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}
	return "where `mail_id`={$params['m']} and `attachment_type` = 0";
}

function sql_insert($params) {
	return "`attachment_type` = 0, `mail_id` = {$params['m']}";
}


$pconfig = array(
	'title'       => '附件',
	'table'       => 'mail_attachments',
	'links'       => array(),
	'columns'     => array(
		array('name' => 'item_id', 'text' => '物品', 'width' => '80px',
			'editor' => 'item_id_editor',
			'render' => 'item_id_render',
		),
		array('name' => 'item_num', 'text' => '数量', 'width' => '40px'),
	),
	'where'   => 'sql_where',
	'insert'  => 'sql_insert',
);
?>