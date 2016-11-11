<?php
$extend_columns = array(
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => $qualityTypes,
	),
);

function render_quality($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function sql_where($params) {
	return ' order by `quality` asc';
}

?>
