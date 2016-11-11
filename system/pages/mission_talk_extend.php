<?php


$extend_columns = array(
	'town_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'quest_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'content' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);




function sql_where($params) {
	return "where `enemy_id`={$params['m']}";
}

function sql_insert($params) {
	return "`enemy_id`={$params['m']}";
}

function editor_quest_id($column_name, $column_val, $row, $data){
	global $all_town;
	global $town_quest;
	$town_id = $row['town_id'];
	$quest_id = $row['quest_id'];
	$quest_name = '';
	if ($row) {
		$quest_name = $town_quest[$town_id][$quest_id];
	} 
	return editor_single_item($quest_name, $quest_id, "quest_id");
}

function render_quest_id($column_name, $column_val, $row, $data){
	global $all_town;
	global $town_quest;
	$town_id = $row['town_id'];
	$quest_id = $row['quest_id'];

	if ($town_id == 0) {
		return normal_render(array('无'), 0);
	}

	return normal_render($town_quest[$town_id], $quest_id);
}

function render_content($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_content($column_name, $column_val, $row, $data){
	return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function get_town() {
	global $db;
	$all = array();
	$tmp = db_query($db, "select * from `town`");
	foreach ($tmp as $value)
	{
		$all[$value['id']] = $value['name'];
	}

	return $all;
}

$all_town = get_town();
$all_town[0]='无';
$town_quest = array();
$town_quest[0][0]='无';
foreach($all_town as $id=>$town) {
	$sql = "select * from quest where related_town={$id}";
	$all_quest = db_query($db, $sql);
	foreach($all_quest as $quest) {
		$town_quest[$id][$quest['id']] = $quest['name'];
	}
	if(!$all_quest) {
		$town_quest[$id][0] = '无';
	}
}

function render_town_id($column_name, $column_val, $row, $data) {
	global $all_town;
	return $all_town[$column_val];
}

function editor_town_id($column_name, $column_val, $row, $data) {
	global $all_town;

	$data = $all_town;
	$field_name = 'town_id';

	$html = "<select class={$field_name} name={$field_name}[]>";

	foreach ($data as $key => $value) {
		if (isset($row[$field_name]) && $row[$field_name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}

	$html .= '</select>';
	return $html;
}

function js_function($params) {
	global $all_town;
	global $town_quest;
	$column_name = 'quest_id';

	$html = '';
	foreach($all_town as $town_id=>$town) {
		if (isset($town_quest[$town_id])) {
			$html .= get_items_json($town_quest[$town_id], "town_{$town_id}");
		}
	}

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: town_1, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;


	$html .= <<<EOT

	$(".town_id").change(function() {
		var town_id = $(this).val() ;
		autoopt_quest_id.lookup = eval('town_' + town_id);
		console.log(autoopt_quest_id.lookup);

		console.log(autoopt_quest_id);
		$(this).parent().parent().find(".display_quest_id").autocomplete(autoopt_quest_id);

	})

EOT;

	return $html;
}

?>
