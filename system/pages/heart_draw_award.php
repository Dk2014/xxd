<?php

$sql = "select draw_type from heart_draw where id = {$_GET['m']}";
$res = db_query($db, $sql);
$title = ($res[0]['draw_type'] == 1) ? "大转盘" : "刮刮卡";

$pconfig = array(
	'title'   => '爱心抽奖奖品 - '.$title,
	'table'   => 'heart_draw_award',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'chance', 'text' => '抽奖概率%', 'width' => '80px'),
		array('name' => 'award_type', 'text' => '奖品类型', 'width' => '80px'),
		array('name' => 'award_num', 'text' => '奖品数量', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '道具奖品(不设请选无)', 'width' => '180px'),
	),
	
	'opera' => array(),
	

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'js'	=> 'js_function',
);
?>