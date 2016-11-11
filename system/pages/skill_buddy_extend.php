<?php
include "common_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'jump_attack' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否','是'),
	),
	'role_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'child_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'child_kind' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'auto_learn_level' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否','是'),
	),
	'parent_skill_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
    'cheat_id' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
        'range' =>array('params' => array()),
    ),
	'target' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(1=>'单体', 2=>'全体', 3=>'横向', 4=>'纵向'),
	),
	'warcry' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function render_jump_attack($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_jump_attack($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

function editor_warcry($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_warcry($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function render_target($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_target($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

$parentSkills = db_query($db, "select * from `skill` where `type` = 1 order by `role_id`;");
foreach ($parentSkills as $value) {
	$parentSkill[$value['id']] = $value['name'];
}

$all_item_cheats = db_query($db, "select `id`, `name` from `item` where `type_id` = 17");
foreach ($all_item_cheats as $value) {
    $all_item_cheat[$value['id']] = $value['name'];
}

$sql = "select * from `role` where id <= 50";
$roles = db_query($db,$sql);
foreach ($roles as $value) {
	$role[$value['id']] = $value['name'];
}

function render_role_id($column_name, $column_val, $row, $data) {
	global $role;
	$html = '';
	if (isset($role[$row['role_id']])){
		$html .= $role[$row['role_id']];
	}
	return $html;
}

function range_role_id(){
	global $role;
	return $role;
}

function range_child_type(){
	global $skillTypes;
	return $skillTypes;
}

function range_child_kind(){
	global $skillKind;
	$kind_temp = $skillKind;
	$king_temp[0] = '无';
	return $skillKind;
}

function range_parent_skill_id(){
	global $parentSkill;
	$temp_skills = $parentSkill;
	$temp_skills[0] = '无';
	return $temp_skills;
}

function range_quality(){
	global $qualityTypes;
	return $qualityTypes;
}

function editor_role_id($column_name, $column_val, $row, $data) {
	global $role;
	return html_get_select($column_name,$role,$column_val);
}

function render_child_kind($column_name, $column_val, $row, $data){
	global $skillKind;
	return $skillKind[$row['child_kind']];
}

function editor_child_kind($column_name, $column_val, $row, $data){
	global $skillKind;
	return html_get_select($column_name,$skillKind,$column_val);
}

function render_child_type($column_name, $column_val, $row, $data){
	global $skillTypes;
	return $skillTypes[$row['child_type']];
}

function editor_child_type($column_name, $column_val, $row, $data){
	global $skillTypes;
	return html_get_select($column_name,$skillTypes,$column_val);
}

function render_quality($column_name, $column_val, $row, $data){
	global $qualityTypes;
	return $qualityTypes[$row['quality']];
}

function editor_quality($column_name, $column_val, $row, $data){
	global $qualityTypes;
	return html_get_select($column_name,$qualityTypes,$column_val);
}

function render_auto_learn_level($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_auto_learn_level($column_name, $column_val, $row, $data){
	return html_get_select($column_name,$data,$column_val);
}

function render_parent_skill_id($column_name, $column_val, $row, $data){
	global $parentSkill;
	$html = '';
	if (isset($parentSkill[$row['parent_skill_id']])){
		$html .= $parentSkill[$row['parent_skill_id']];
	}
	return $html;
}

function editor_parent_skill_id($column_name, $column_val, $row, $data){
	global $parentSkills;
	$code = '';
	$code .= '<select name="parent_skill_id[]" >';
	$code .= '<option value="0" selected="selected">无</option>';
	if ($row != null) {
		foreach ($parentSkills as $value){
			if ($value['id'] == $row['parent_skill_id']) {
				$code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
			} else {
				$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
			}
		}
	} else {
		foreach ($parentSkills as $value){
			$code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function render_cheat_id($column_name, $column_val, $row, $data){
    global $all_item_cheat;
    $html = '';
    if (isset($all_item_cheat[$row['cheat_id']])){
        $html .= $all_item_cheat[$row['cheat_id']];
    }
    return $html;
}

function editor_cheat_id($column_name, $column_val, $row, $data){
    global $all_item_cheats;
    $code = '';
    $code .= '<select name="cheat_id[]" >';
    $code .= '<option value="0" selected="selected">无</option>';
    if ($row != null) {
        foreach ($all_item_cheats as $value){
            if ($value['id'] == $row['cheat_id']) {
                $code .= '<option value="'.$value['id'].'"  selected="selected">'.$value['name'].'</option>';
            } else {
                $code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
            }
        }
    } else {
        foreach ($all_item_cheats as $value){
            $code .= '<option value="'.$value['id'].'" >'.$value['name'].'</option>';
        }
    }
    $code .= '</select>';
    return $code;
}

function range_cheat_id(){
	global $all_item_cheat;
	return $all_item_cheat;
}

function skill_role_render($row) {
	$html = '';
	$html .= '<table border="0" width="200"><tr><td width="180">';
	$html .= '<a href="javascript:;" onclick="open_editor(\''.$row['name'].'\', \''.$row['info'].'\', \''.$row['id'].'\')">配置绝招 </a>';
	$html .= "<a href='?p=skill_buddy_data&m={$row['id']}'> 释放频率</a>";
	$html .= '<a href="javascript:;" onclick="edit_skill_info(\''.$row['id'].'\')">绝招描述</a>';
	$html .= skill_buddy_info("
select * from `skill_content` where `skill_id` = {$row['id']};
	");
	$html .= '</td></tr></table>';
	return $html;
}

function foot() {
	require dirname(__FILE__).'/skill_editor.php';
	require dirname(__FILE__).'/skill_info_editor.php';
}

function skill_opera($row) {
	return '<a href="javascript:;" onclick="open_editor(\''.$row['name'].'\', \''.$row['info'].'\', \''.$row['id'].'\', false)">配置绝招 </a>';
}

function sql_where($params) {
	if(isset($params['m'])) {
		return " where `type` =1 and `role_id` = {$params['m']} ORDER BY `role_id`,`required_level` ASC";
	}
	return " where `type` =1 and `role_id` != -2 ORDER BY `role_id`,`required_level` ASC";
}

function sql_insert($params) {
	return "`type` = 1";
}

?>
