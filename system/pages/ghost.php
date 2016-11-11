<?php

include "common_links.php";
include "skill_links.php";

function sql_where($params) {
	return " order by `quality`";
}

function ghost_unique_key_render($row){
	global $uniqueKeyConfig;
	if ($row['unique_key'] == 0){
		return "无";
	}
	return $uniqueKeyConfig[$row['unique_key']];
}

function ghost_unique_key_editor($row){
	global $uniqueKeyConfig;

	$code = '';
	$code .= '<select name="unique_key[]" >';
	if ($row != null) {
		foreach ($uniqueKeyConfig as $key => $value){
			if ($key == $row['unique_key']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($uniqueKeyConfig as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function ghost_opera($row) {
	return html_get_link('绝招', "?p=ghost_skill&m={$row['id']}", false) .  '|' .
		//html_get_link('威力', "?p=ghost_skill_force&m={$row['id']}", false) .  '|' .
	 ghost_skill_info("select * from `skill` where `type` = 7 and `role_id`={$row['id']} order by `required_level` limit 1");

}


function _ghost_opera($row) {
	$html = '';
	$html .= '<table border="0"><tr><td width="110">';
	$html .= '<a href="?p=ghost_skill&m='.$row['id'].'">魂侍绝招</a> ';
	$html .= ghost_skill_info("select * from `skill` where `type` = 7 and `role_id`={$row['id']} order by `required_level`");
	$html .= '</td></tr></table>';

	return $html;
}

function ghost_skill_info($sql){
	global $db;
	global $skillTypes;
	$html = '';
	$skillInfos = db_query($db, $sql);

	foreach ($skillInfos as $skillInfo) {
		if (isset($skillInfo['name'])){
			$html .= "</br>名称:".$skillInfo['name'];
		}
	
		if (isset($skillInfo['child_type'])){
			$html .= " 类型:". $skillTypes[$skillInfo['child_type']];
		}	
	}

	return $html;
}

function ghost_desc_editor($row){
	$code  = '<textarea name="desc[]" rows="7" cols="20" >';
	$code .= $row['desc'];
	$code .= '</textarea>';
	return $code;
}

function ghost_desc_render($row){
	$code  = '<textarea disabled="disabled" name="desc[]" rows="7" cols="20">';
	$code .= $row['desc'];
	$code .= '</textarea>';
	return $code;
}

function ghost_quality_render($row){
	global $qualityTypes;
	return $qualityTypes[$row['quality']];
}

function ghost_quality_editor($row){
	global $qualityTypes;

	$code = '';
	$code .= '<select name="quality[]" >';
	if ($row != null) {
		foreach ($qualityTypes as $key => $value){
			if ($key == $row['quality']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($qualityTypes as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

$all_stars[1] = '1';
$all_stars[2] = '2';
$all_stars[3] = '3';
$all_stars[4] = '4';
$all_stars[5] = '5';
$all_stars[6] = '6';


function ghost_init_star_render($row){
	global $all_stars;
	return $all_stars[$row['init_star']];
}

function ghost_init_star_editor($row){
	global $all_stars;

	$code = '';
	$code .= '<select name="init_star[]" >';
	if ($row != null) {
		foreach ($all_stars as $key => $value){
			if ($key == $row['init_star']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($all_stars as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function _ghost_get_all_role() {
	global $db;
	$all_roles = array();

	$tmp = db_query($db, "select `id`, `name` from `role` where `id` >=3 order by `id`");
	foreach ($tmp as $value)
	{
		$all_roles[$value['id']] = $value['name'];
	}

	return $all_roles;
}

$all_roles = _ghost_get_all_role();
$all_roles[-1] = '主角';

function ghost_role_id_render($row){
	global $all_roles;
	return $all_roles[$row['role_id']];
}

function ghost_role_id_editor($row){
	global $all_roles;

	$code = '';
	$code .= '<select name="role_id[]" >';
	if ($row != null) {
		foreach ($all_roles as $key => $value){
			if ($key == $row['role_id']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($all_roles as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function get_all_ghost_fragment() {
	global $db;
	$all_ghost_fragments = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id` = 13");
	foreach ($tmp as $value)
	{
		$all_ghost_fragments[$value['id']] = $value['name'];
	}

	return $all_ghost_fragments;
}

$all_ghost_fragments = get_all_ghost_fragment();
$all_ghost_fragments[0] = '无';

function ghost_fragment_id_render($row){
	global $all_ghost_fragments;
	return $all_ghost_fragments[$row['fragment_id']];
}

function ghost_fragment_id_editor($row){
	global $all_ghost_fragments;

	$code = '';
	$code .= '<select name="fragment_id[]" >';
	if ($row != null) {
		foreach ($all_ghost_fragments as $key => $value){
			if ($key == $row['fragment_id']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($all_ghost_fragments as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

$pconfig = array(
	'title'       => '魂侍',
	'table'       => 'ghost',
	'links'       => $ghost_links,
	'columns'     => array(
		array('name' => 'name',         'text' => '名称',     'width' => '70px'),
		array('name' => 'sign',         'text' => '资源标识', 'width' => '100px'),
		array('name' => 'music_sign',         'text' => '音乐资源标识', 'width' => '100px'),
		array('name' => 'town_id', 'text' => '出现影界', 'width' => '80px'),
		//array('name' => 'unique_key',         'text' => '唯一值',    'width' => '70px',
		//	'render' => "ghost_unique_key_render",
		//	'editor' => "ghost_unique_key_editor",
		//),
		//array('name' => 'role_id', 'text' => '使用者', 'width' => '80px',
		//	'editor' => "ghost_role_id_editor",
		//	'render' => "ghost_role_id_render",
		//),
		array('name' => 'fragment_id', 'text' => '对应碎片', 'width' => '80px',
			'editor' => "ghost_fragment_id_editor",
			'render' => "ghost_fragment_id_render",
		),
		array('name' => 'init_star', 'text' => '初始星级', 'width' => '80px',
			'editor' => "ghost_init_star_editor",
			'render' => "ghost_init_star_render",
		),
		array('name' => 'quality', 'text' => '品质', 'width' => '80px',
			'editor' => "ghost_quality_editor",
			'render' => "ghost_quality_render",
		),
		array('name' => 'health',       'text' => '基础生命',       'width' => '70px'),
		array('name' => 'attack',       'text' => '基础攻击',       'width' => '70px'),
		array('name' => 'defence',      'text' => '基础防御',       'width' => '70px'),
		array('name' => 'desc',         'text' => '魂侍简介',    'width' => '200px',
			'render' => "ghost_desc_render",
			'editor' => "ghost_desc_editor",
		),
		array('name' => 'production_info',         'text' => '产出描述',    'width' => '80px'),
	),
	'opera' => array(
		array('type' => 'link', 'render' => 'ghost_opera'),
	),
	'where'   => 'sql_where',
	'not_delete' => true,
);

?>
