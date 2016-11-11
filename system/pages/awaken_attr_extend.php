<?php

$extend_columns = array(
	'is_skill' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'skill_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
);

function range_is_skill() {
	$bool = array(
		1 => '是',
		2 => '否',
	);
	return $bool;
}

function editor_is_skill($col_nam, $col_val, $row, $data) {
	$selno = "";
	$selyes = "";
	if ($row['is_skill']) {
		$selyes = "selected";
	} else {
		$selno = "selected";
	}

	return 
<<<hOoYaYaYa
	<select name="is_skill[]">
		<option value="0" {$selno}>否</option>
		<option value="1" {$selyes}>是</option>
	</select>
hOoYaYaYa;
	}

function render_is_skill($col_nam, $col_val, $row, $data) {
	return $row['is_skill']?"是":"否";
}

$skills = array();
$skills_t = db_query($db,"select `id`, `name` from `skill` where `type` = 1;");
foreach ($skills_t as $row) {
	$skills[$row['id']] = $row['name'];
}
$skills[0] = '无';

function range_skill_id(){
	global $skills;
	return $skills;
}

function render_skill_id($col_nam, $col_val, $row, $data) {
	global $skills;
	return $skills[$row['skill_id']];
}

function editor_skill_id($col_nam, $col_val, $row, $data){
	global $skills;

	$code = '<select name="skill_id[]">';
	foreach ($skills as $k => $v){
		if ($k == $row['skill_id']) {
			$code .= '<option value="'.$k.'" selected>'.$v.'</option>';
		} else {
			$code .= '<option value="'.$k.'">'.$v.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

// 属性类型定义列表
$attr_type = array(
	0 => '无',
	1 => '生命',
	2 => '攻击',
	3 => '防御',
	4 => '内力',
	5 => '速度',
	6 => '护甲',

	7 => '命中等级',
	8 => '闪避等级',
	9 => '格挡等级',
	10 => '破击等级',
	11 => '暴击等级',
	12 => '韧性等级',
	13 => '中毒抗性等级',
	14 => '封魔抗性等级',
	15 => '睡眠抗性等级',
	16 => '眩晕抗性等级',
	17 => '混乱抗性等级',
	18 => '仙灵伤害',
	19 => '人兽伤害',
	20 => '妖魔伤害',
);

function range_type(){
	global $attr_type;
	return $attr_type;
}

function render_type($col_nam, $col_val, $row, $data) {
	global $attr_type;
	return $attr_type[$row['type']];
}

function editor_type($col_nam, $col_val, $row, $data){
	global $attr_type;

	$code = '<select name="type[]">';
	foreach ($attr_type as $k => $v){
		if ($k == $row['type']) {
			$code .= '<option value="'.$k.'" selected>'.$v.'</option>';
		} else {
			$code .= '<option value="'.$k.'">'.$v.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}
