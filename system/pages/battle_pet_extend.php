<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'pet_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'tip_render'=> array('params' => array()),
		'range' => array('params' => array()),
	),

	'parent_pet_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),

	'quality' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('蓝','紫','金','橙'),
	),
	'star' => array(
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
	'live_pos' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('1'=>'前排','2'=>'后排','3'=>'左侧','4'=>'右侧'),
	),
	'desc' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(),
	),	
	'round_attack' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array('1'=>'1次','2'=>'2次'),
	),
	'skill' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' =>array(),
	),
);

$skill = array();
$sql = "select `id`, `name` from `skill`;";
$skillRows = db_query($db, $sql);
foreach ($skillRows as $row) {
	$skill[$row['id']] = $row['name'];
}

$all_items = get_all_item();

function range_skill(){
	global $skill;
	$skill_temp = $skill;
	$skill_temp[0] = '无';
	return $skill_temp;
}


function render_skill($column_name, $column_val, $row, $data) {
	global $skill;
	$html = '';
	if (isset($skill[$row['skill']])){
		$html .= '<a>'.$skill[$row['skill']].'</a>';
	}
	return $html;
}

function editor_skill($column_name, $column_val, $row, $data){
	global $skill;
	$name = '';
	if(isset($skill[$column_val])) {
		$name = $skill[$column_val];
	}
	return editor_single_item($name, $column_val, "skill");
}

$allEnemy = get_all_enemy();

function render_parent_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemy[$column_val])) {
		return html_get_link($allEnemy[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_parent_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy;
	return enemy_role_editor($allEnemy, $row, $column_name);
}

function render_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy;

	if ($column_val == 0) {
		return "无";
	} else if(isset($allEnemy[$column_val])) {
		return html_get_link($allEnemy[$column_val]['name'], '?p=enemy_role&m='.$column_val, true);
	}
}

function editor_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy;
	return enemy_role_editor($allEnemy, $row, $column_name);
}

function tip_render_pet_id($column_name, $column_val, $row, $data){
	global $allEnemy;

	if (isset($allEnemy[$column_val])) {
		$role = $allEnemy[$column_val];

		return mission_enemy_tip_render($role);
	}
}

function render_round_attack($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_round_attack($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_quality($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_quality($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}


function render_star($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_star($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function render_live_pos($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_live_pos($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function range_pet_id(){
	global $allEnemy;
	foreach ($allEnemy as $key => $enemy) {
		$allEnemyTemp[$key] = $enemy['name'];
	}
	$allEnemyTemp[0] = '无';
	return $allEnemyTemp;
}

function range_item_battle_pet_id(){
	global $all_items;
	$all_items_temp = $all_items;
	$all_items_temp[0] = '无';
	return $all_items_temp;
}

function range_parent_pet_id(){
	global $allEnemy;
	foreach ($allEnemy as $key => $enemy) {
		$allEnemyTemp[$key] = $enemy['name'];
	}
	$allEnemyTemp[0] = '无';
	return $allEnemyTemp;
}

function sql_where($params) {
	return " order by `quality`";
}

function jsFunction($params) {
	global $allEnemy, $allItem, $all_items, $skill;

	$html = '
		$("select").change( function(){
			$(this).css("color","#0033FF");
});';

$html .= get_items_json($skill, 'skill');
$html .= <<<EOT
	autoopt_skill = {
		lookup: skill, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};
	$(".display_skill").autocomplete(autoopt_skill);
EOT;

$html .= enemy_auto_complete($allEnemy);
$html .= choose_single_item($all_items, 'item_battle_pet_id');
return $html;
}
?>
