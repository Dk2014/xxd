<?php


$extend_columns = array(
/*   '字段' => 配置 */

	'draw_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('1'=>'大转盘','2'=>'刮刮卡'),
	),
);


function render_draw_type($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_draw_type($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

function op_render($row){
	return html_get_link("奖品配置", "?p=heart_draw_award&m=".$row['id'], true);
}


?>