<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$all_items = array();
$res = db_query($db, "select `id`, `name` from `item` where `type_id` != 10 and `type_id` != 11;");
foreach ($res as $value)
{
	$all_items[$value['id']] = $value['name'];
}

function editor_item_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$item_id = $row['item_id'];
	global $all_items;
	if (isset($all_items[$item_id])) {
		$item_name = $all_items[$item_id];
	}
	return editor_single_item($item_name, $item_id, "item_id");
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_items;
	return normal_render($all_items, $row['item_id']);
}


function js_function($params) {
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

	//$nothing = array('无');
	//$html .= get_items_json($nothing, 'nothing');



return $html;
}

?>
