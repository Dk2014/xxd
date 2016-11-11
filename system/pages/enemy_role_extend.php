<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'jump_attack' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("否","是"),
	),
	'show_shader' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("否","是"),
	),
	'body_size' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("1"=>"小","2"=>"中","3"=>"大"),
	),
	'is_boss' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("否","是"),
	),
	'prop' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array("仙","灵","人","兽","妖","魔"),
	),
);

function editor_prop($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_prop($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function render_is_boss($column_name, $column_val, $row, $data){
	return $data[$column_val];
}
function render_show_shader($column_name, $column_val, $row, $data){
	return $data[$column_val];
}
function editor_show_shader($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_is_boss($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_jump_attack($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_jump_attack($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}


function render_body_size($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_body_size($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function op_render($row){
	$html = '<a href="javascript:;" onclick="open_skillEditor('.$row['id'].')">技能</a>';
	$html = '<a href="javascript:;" onclick="open_passive_skill_editor('.$row['id'].')">被动技能</a>';

	if ($row['is_boss']) {
		$html .= ' | <a href="javascript:;" onclick="open_editor('.$row['id'].')">脚本</a>';	
	}
	return $html;
}

function foot() {
	require dirname(__FILE__).'/enemy_boss_script_editor.php';
	require dirname(__FILE__).'/enemy_skill_editor.php';
	require dirname(__FILE__).'/enemy_passive_skill_editor.php';
}

function monster_property_add($row){
  $rtn='';
  $rtn .= ' | <a href="?p=monster_property_addition&m='.$row['id'].'">属性加成</a>';
  return $rtn;
} 
?>
