<?php
$extend_columns = array(
	'quality' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

 
function editor_quality($column_name, $column_val, $row, $data){
	global $qualityTypes;
	$item_name = '';
	$quality = $row['quality'];
	if (isset($qualityTypes[$quality])) {
		$item_name = $qualityTypes[$quality];
	}
	return editor_single_item($item_name, $quality, "quality");
}

function render_quality($column_name, $column_val, $row, $data){
	global $qualityTypes;
	return normal_render($qualityTypes, $row['quality']);
}

function range_quality(){
	global $qualityTypes;
	return $qualityTypes;
}

function js_function($params) {
	global $qualityTypes;

	$html = get_items_json($qualityTypes, 'quality');

	$column_name = 'quality';
	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: quality, 
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
