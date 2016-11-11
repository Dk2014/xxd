<?php
//include "item_links.php";
$extend_columns = array(
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);


function sql_where($params) {
	return "where config_id={$params['m']} order by `min_level`";
}

function sql_insert($params) {
	return "config_id={$params['m']}";
}


function get_all_equipment() {
	global $db;
	$all_item = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id` in ('3','4','5','6')");
	foreach ($tmp as $value)
	{
		$all_item[$value['id']] = $value['name'];
	}

	return $all_item;
}

$all_items = get_all_equipment();

function editor_item_id($column_name, $column_val, $row, $data){
	global $all_items;
	$item_id = $row['item_id'];
	$item_name = '';
	if(isset($all_items[$item_id])) {
		$item_name = $all_items[$item_id];
	} 

	return editor_single_item($item_name, $item_id, "item_id");
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_items ;
	$item_id = $row['item_id'];
	$item_name = $all_items[$item_id];
	return normal_render($all_items, $row['item_id']);
}

function jsFunction($params) {
	global $all_items;

	$column_name = 'item_id';

	$html = get_items_json($all_items, 'items');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;

	return $html;
}

?>
