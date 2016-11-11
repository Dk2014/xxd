<?php
$extend_columns = array(
/*   '字段' => 配置 */
	'condition' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(1 => '翻倍活动关卡铜钱',2 => '翻倍活动关卡经验', 3 => 'QQ会员关卡加成铜钱', 4 => 'QQ会员关卡加成经验', 5 => 'QQ超级会员加成铜钱', 6 => 'QQ超级会员加成经验'),
	),
);

function editor_condition($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_condition($column_name, $column_val, $row, $data){
	return $data[$column_val];
}
?>