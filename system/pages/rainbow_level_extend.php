<?php


$extend_columns = array(
	'award_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$all_ghosts = get_all_ghost();
$all_ghosts[0] = '无';


function editor_award_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$award_id = $row['award_id'];

	global $all_ghosts;
	if (isset($all_ghosts[$award_id])) {
		$item_name = $all_ghosts[$award_id];
	}

	return editor_single_item($item_name, $award_id, "award_id");
}

function render_award_id($column_name, $column_val, $row, $data){
	global $all_ghosts;
	return normal_render($all_ghosts, $row['award_id']);
}

function js_function($params) {
	global $all_ghosts;

	$column_name = 'award_id';

	$html = get_items_json($all_ghosts, 'ghosts');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: ghosts, 
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
