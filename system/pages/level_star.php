<?php
$pconfig = array(
	'title'   => '关卡评星',
	'table'   => 'level_star',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'two_star_round', 'text' => '二星要求回合', 'width' => '80px'),  
		array('name' => 'three_star_round', 'text' => '三星要求回合', 'width' => '80px'),  


	),
	

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
);

function sql_where($params) {
	return "where `level_id`={$params['m']} ";
}

function sql_insert($params) {
	return "`level_id`={$params['m']} ";
}

?>
