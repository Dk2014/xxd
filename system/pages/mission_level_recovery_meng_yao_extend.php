<?php
$extend_columns = array(
/*   '字段' => 配置 */
'my_effect' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(1=> '恢复灵宠使用次数', 2=> '恢复全体生命值', 3=>'恢复伙伴绝招次数')
	),
	'my_dir' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('r'=>'右', 'rb'=>'右下方', 'b'=>'下', 'lb'=>'左下方', 'l'=>'左', 'lt'=>'左上方', 't'=>'上', 'rt'=>'右上方'),
	),
	'talk' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function editor_my_effect($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_my_effect($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_my_dir($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_my_dir($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function render_talk($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_talk($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}
?>