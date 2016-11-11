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
	return "where `mail_id`={$params['m']} and `attachment_type` != 0 ";
}

function sql_insert($params) {
	return "`mail_id` = {$params['m']}, `item_id` = 0 ";
}

$special_attachment = array(
	1 => '铜钱',
	2 => '元宝',
	3 => '爱心',
	4 => '经验(所有上阵角色)',
	5 => '经验(单个角色)',
	8 => '好友爱心',
);


function attachment_type_editor($row){
	global $special_attachment;
	$name = "attachment_type";

	$html = '<select name="'.$name.'[]" >';
	foreach ($special_attachment as $id  => $attach) {
		if ( isset($row[$name]) && $row[$name] == $id) {
			$html .= '<option value="'.$id.'" selected="selected" >'.$attach.'</option>';
		} else {
			$html .= '<option value="'.$id.'">'.$attach.'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

function attachment_type_render($row){
	global $special_attachment;
	$name = "attachment_type";

	if (empty($row[$name])) {
		return '';
	}
	return $special_attachment[$row[$name]];
}

$pconfig = array(
	'title'       => '特殊附件',
	'table'       => 'mail_attachments',
	'links'       => array(),
	'columns'     => array(
		array('name' => 'attachment_type', 'text' => '类型', 'width' => '80px',
			'editor' => 'attachment_type_editor',
			'render' => 'attachment_type_render',
		),
		array('name' => 'item_num', 'text' => '数量', 'width' => '40px'),
	),
	'where'   => 'sql_where',
	'insert'  => 'sql_insert',
);
?>
