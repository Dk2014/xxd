<?php
$pconfig = array(
	'title'   => '随机物品配置',
	'table'   => 'trader_grid_config',
	'links'   => array(),

	'columns' => array(
		array('name' => 'goods_type', 'text' => '货物类型', 'width' => '80px'),
		array('name' => 'item_id', 'text' => '货物', 'width' => '80px'),
		array('name' => 'num', 'text' => '物品数量', 'width' => '80px'),
		array('name' => 'cost', 'text' => '价格', 'width' => '80px'),
		array('name' => 'probability', 'text' => '出现概率（％）', 'width' => '80px'),

	),

	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'trader_grid_config_equiement'),
	),


	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	'location'  		=> 'location',
	'js'	        => 'js_function',

);

function trader_grid_config_equiement($row) {
	return html_get_link('等级装备配置', "?p=trader_grid_config_equiement&m={$row['id']}", true);
}

?>
