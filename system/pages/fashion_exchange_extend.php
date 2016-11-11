<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function sql_where($params) {
	return " where `fashion_id` = {$params['m']} ";
}

function sql_insert($params) {
	return " `fashion_id` = {$params['m']} ";
}


$all_fashion = get_all_fashion_item();

 
function editor_item_id($column_name, $column_val, $row, $data){
	global $all_fashion;
	$item_name = '';
	$item_id = $row['item_id'];
	if (isset($all_fashion[$item_id])) {
		$item_name = $all_fashion[$item_id];
	}
	return editor_single_item($item_name, $item_id, "item_id");
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_fashion;
	return normal_render($all_fashion, $row['item_id']);
}


function js_function($params) {
	global $all_fashion;

	$html = get_items_json($all_fashion, 'items');

	$column_name = 'item_id';
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
