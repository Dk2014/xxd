<?php

include "item_links.php";

$pconfig = array(
	'title'   => '物品类型',
	'table'   => 'item_type',
	'links'   => $links,
	'columns' => array(
		array('name' => 'order', 'text' => '客户端排序权重', 'width' => '120px'),
		array('name' => 'name', 'text' => '名称', 'width' => '150px'),
		array('name' => 'max_num_in_pos', 'text' => '堆叠数量', 'width' => '120px'),
		array('name' => 'sign', 'text' => '类型标志'),
	),
	'not_delete' => true,
	
	'where' 		=> 'sql_where',
);

function sql_where($params) {
	return "order by `order` ASC";
}
?>
