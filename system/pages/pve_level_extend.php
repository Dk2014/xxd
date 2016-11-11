<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'award_item' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);



 
function editor_award_item($column_name, $column_val, $row, $data){
	global $all_items;
	$item_name = '';
	$award_item = $row['award_item'];
	if (isset($all_items[$award_item])) {
		$item_name = $all_items[$award_item];
	}
	return editor_single_item($item_name, $award_item, "award_item");
}

function render_award_item($column_name, $column_val, $row, $data){
	global $all_items;
	return normal_render($all_items, $row['award_item']);
}

$all_items = get_all_item();

function js_function($params) {
	global $all_items;

	$html = get_items_json($all_items, 'items');

	$column_name = 'award_item';
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

function leveltable($row) {
	return html_get_link("怪物配置", "?p=pve_level_config&m=".$row['id'], true);
}

?>
