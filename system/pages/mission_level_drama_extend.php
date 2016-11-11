<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'quest_id' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' 		=> array(),
		'range'     => array('params' => array()),
	),
	'quest_state' => array(
		'editor' 	=> array('params' => array()),
		'render' 	=> array('params' => array()),
		'data' => array(
			'0' => '未领取',
			'1' => '未完成',
			'2' =>  '未奖励',
		),
	),
);

$allQuests = get_all_quest();
$allQuests[0] = '一直有效';

function render_quest_state($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_quest_state($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}

function editor_quest_id($column_name, $column_val, $row, $data){
	global $allQuests;
	$quest_name = '';
	if (isset($allQuests[$column_val])) {
		$quest_name = $allQuests[$column_val];
	}

	return editor_single_item($quest_name, $column_val, $column_name);
}

function render_quest_id($column_name, $column_val, $row, $data){
	global $allQuests;

	return (isset($allQuests[$column_val]) ? $allQuests[$column_val] : "0");
}

function sql_where($params) {
	if(!isset($params['m'])) {
		die("need mission_level_id");
	}
	return "where mission_level_id={$params['m']} ";
}

function sql_insert($params) {
	if(!isset($params['m'])) {
		die("need mission_level_id");
	}
	return " mission_level_id={$params['m']} ";
}

function  jsFunction($params) {
	global $allQuests;
	$html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';
	$html .= choose_single_item($allQuests, 'quest_id');
	return $html;
}

?>
