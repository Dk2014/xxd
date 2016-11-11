<?php
include "vip_common.php";

function sql_where($params) {
	return "order by `level` asc";
}

$pconfig = array(
	'title'   => 'VIP等级数值',
	'table'   => 'vip_level',
	'links'  => $links, 

	'columns' => array(
		array('name' => 'level', 'text' => 'VIP等级', 'width' => '80px'),
		array('name' => 'ingot', 'text' => '累计充值元宝要求', 'width' => '80px'),
	),

	'opera' => array(
		array('type' => 'link', 'text' => '配置特权', 'render' => 'privilege_config'),
		array('type' => 'link', 'text' => 'VIP升级奖励', 'render' => 'vip_levelup_gift'),
	),

);
?>


