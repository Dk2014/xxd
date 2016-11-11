<?php
$pconfig = array(
	'title'   => '关卡梦妖',
	'table'   => 'mission_level_recovery_meng_yao',
	'links'   => array(),
	'where'	  => 'sql_where',
	'insert'  => 'sql_insert',
	'columns' => array(
		array('name' => 'my_x', 'text' => '关卡梦妖x坐标', 'width' => '150px'),
		array('name' => 'my_y', 'text' => '关卡梦妖y坐标', 'width' => '150px'),
		array('name' => 'probability', 'text' => '出现概率%', 'width' => '150px'),
		array('name' => 'my_dir', 'text' => '关卡梦妖朝向', 'width' => '60px'),
		array('name' => 'my_effect', 'text' => '关卡梦妖效果', 'width' => '150px'),
		array('name' => 'talk','text' => '关卡梦妖对话内容','width' => '150px'),

	),

);

function sql_where($params){
	if (!isset($params['m'])){
		return '';
	}
	return "WHERE `mission_level_id` = {$params['m']}";
}

function sql_insert($params) {
	return "`mission_level_id` = {$params['m']}";
}
?>