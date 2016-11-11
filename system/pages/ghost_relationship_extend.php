<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'ghost_a' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'tip_render'=> array('params' => array()),
		'range' => array('params' => array()),
	),

	'ghost_b' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'ghost_a_star' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(
				1=>'一星',
				2=>'二星',
				3=>'三星',
				4=>'四星',
				5=>'五星',
			),
	),

	'ghost_b_star' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(
				1=>'一星',
				2=>'二星',
				3=>'三星',
				4=>'四星',
				5=>'五星',
			),
	),
	);

$allGhosts = get_all_ghost();
function editor_ghost_a($column_name, $column_val, $row, $data){
	global $allGhosts;
	return html_get_select($column_name, $allGhosts, $column_val);
}

function render_ghost_a($column_name, $column_val, $row, $data){
	global $allGhosts;
	return $allGhosts[$column_val];
}

function range_ghost_a(){
	global $allGhosts;
	$allGhostsTemp = $allGhosts;
	return $allGhostsTemp;
}

function editor_ghost_b($column_name, $column_val, $row, $data){
	global $allGhosts;
	return html_get_select($column_name, $allGhosts, $column_val);
}

function render_ghost_b($column_name, $column_val, $row, $data){
	global $allGhosts;
	return $allGhosts[$column_val];
}

function range_ghost_b(){
	global $allGhosts;
	$allGhostsTemp = $allGhosts;
	return $allGhostsTemp;
}

function tip_render_ghost_a(){

}

function tip_render_ghost_b(){
	
}

function editor_ghost_a_star($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_ghost_a_star($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_ghost_b_star($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_ghost_b_star($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

?>