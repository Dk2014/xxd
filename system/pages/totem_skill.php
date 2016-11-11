<?php

include "common_links.php";
include "skill_links.php";

function sql_where($params) {
	return "where `type`=9 and `role_id`={$params['m']} order by `required_level`,type";
}

function sql_insert($params) {
	return "`role_id` = {$params['m']}, `type`=9";
}

function location($params){
	global $db;

	$sql = "select * from `totem` where id = {$params['m']}";
	$totem = db_query($db, $sql);


	$html = html_get_link("阵印", "?p=totem", false);
	$html .= '->';
	@$html .=html_get_link($totem[0]['name'], '#', false);

	return $html;
}

function foot() {
	require dirname(__FILE__).'/skill_editor.php';
	require dirname(__FILE__).'/skill_info_editor.php';
}

function skill_opera($row) {
	return '<a href="javascript:;" onclick="open_editor(\''.$row['name'].'\', \''.$row['info'].'\', \''.$row['id'].'\', false)">配置绝招 </a>';
}

function skill_info($row) {
    return '<a href="javascript:;" onclick="edit_skill_info(\''.$row['id'].'\')">绝招描述</a>';
}



function target_render($row){
	global $skill_target_type;
	if (!isset($row['target'])){
		return '';
	}
	return $skill_target_type[$row['target']];
}


function target_editor($row){
	global $skill_target_type;
	$code = '';
	$code .= '<select name="target[]" >';
	if ($row != null) {
		foreach ($skill_target_type as $key => $value){
			if ($key == $row['target']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($skill_target_type as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function skill_type_render($row){
	global $skillTypes;
	if (!isset($row['child_type'])){
		return '';
	}
	return $skillTypes[$row['child_type']];
}

function skill_type_editor($row){
	global $skillTypes;
	$code = '';
	$code .= '<select name="child_type[]" >';
	if ($row != null) {
		foreach ($skillTypes as $key => $value){
			if ($key == $row['child_type']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($skillTypes as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

$pconfig = array(
	'title'   => '阵印技能',
	'table'   => 'skill',
	'links'   => array(),
	'location'=> 'location',
	'columns' => array(
		array('name' => 'name', 'text' => '绝招名称', 'width' => '120px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '100px'),
		array('name' => 'music_sign', 'text' => '音乐资源标识', 'width' => '100px'),
		array('name' => 'child_type', 'text' => '绝招类型', 'width' => '80px',
		'editor' => "skill_type_editor",
		'render' => "skill_type_render",
		),
		array('name' => 'target', 'text' => '攻击方式', 'width' => '60px',
		'editor' => "target_editor",
		'render' => "target_render",
		),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '绝招配置', 'render' => 'skill_opera'),
	    array('type' => 'link', 'text' => '绝招描述', 'render' => 'skill_info'),
	),
//'not_delete' => true,
//'not_new' => true,
	'foot' => 'foot',
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
);
?>
