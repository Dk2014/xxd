<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'keys' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function editor_keys($column_name, $column_val, $row, $data){
	return html_get_input($column_name, format_big_number($column_val));
}

function render_keys($column_name, $column_val, $row, $data){
	return format_big_number($column_val);
}

function location($params){
	global $db;
	$html = '';
	
	$sql = "select `id`, `name` from town where id = {$params['m']}";
	$town = db_query($db, $sql);
	if (count($town) == 0) {
		return $html;
	}

	return html_get_link("所有城镇", "?p=town", false).' -> '.$town[0]['name'];
}

function sql_where($params) {
	global $allTown;
	
	if (!isset($params['m'])){
		return '';
	}
	
	return "where `town_id`={$params['m']} order by `id` asc";
}

function sql_insert($params) {
	return "`town_id` = {$params['m']}";
}


function js_function($params){
	global $itemId;
	$html = '';

	$html .= <<<EOT

	$("#form").submit(function(){
		isOk = true;
		$("#table_body  input[name='keys[]']").each(function(i){
			newValue = $(this).val().replace(/,/g, '');
			if (newValue < 0 || newValue > 999999) {
				isOk = false;
				return;
			}
			$(this).val( newValue );
		})
		if (!isOk) {
			alert("权值位数应该设为6位")
			return false;
		} 
	})
EOT;
	$html .= '
	$(".new").click(function(){
		$("#table_body  input[name=\'keys[]\']").each(function(i){
			if ($(this).val() == "" ) {
				$(this).val( "100" );
			}
		})
	});
	';
	return $html;
}

function mission_enemy_render($row) {
	return html_get_link("关卡", "?p=mission_level&m=".$row['id'], false). ' | '.
		html_get_link("区域评星总数配置", "?p=mission_star_awards&m=".$row['id'], false);
}


?>
