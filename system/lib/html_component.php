<?php

function hava_extend_column($column_name, $action) {
	global $extend_columns;
	return isset($extend_columns[$column_name]);
}

function execute_extend_column($action, &$column, &$row) {
	global $extend_columns;
	$extend_data = $extend_columns[$column['name']];
	
	$i = 0;
	$point_column = $column['name'];
	while (!is_array($extend_data)) { //复用扩展数据
		$point_column = $extend_data;
		$extend_data = $extend_columns[$extend_data];
		if ($i < 9999)
			$i++;
		else 
			break; // 防止配置不当,死循环
	}

	if (is_array($extend_data) && isset($extend_data[$action])) {
		$fn = "{$action}_{$point_column}";
		return $fn($column['name'], $row[$column['name']], $row, $extend_data['data']);
	}
	return "";
}

function format_big_number($value) {
	// 101,000 6位
	$str = (string)$value;
	$arr = array();
	for ($i = 0; $i < strlen($str); $i++){
		if ($i % 3 == 2 ){
			array_push($arr, $str[$i-2].$str[$i-1].$str[$i]);
		}
	}
	return count($arr) ? implode(',', $arr) : $str;
}

function html_get_select($column, $data_list, $current_select_value) {
	$code = "<select name=\"{$column}[]\" >";
	foreach ($data_list as $key => $value){
		if ($key == $current_select_value) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}

	$code .= '</select>';
	return $code;
}

function html_get_link($name, $uri, $is_blank) {
	$blank = "";
	if ($is_blank) {
		$blank = 'target="_blank"';
	}
	return "<a href=\"{$uri}\" {$blank}>{$name}</a>";
}

function html_get_input($name, $value, $width = 60) {
	return '<input name="'.$name.'[]" type="text" value="'.$value.'" class="editor" style="overflow: hidden;width:'.$width.'px"/>';
}

function html_get_showFrame($tag, $path, $w, $h) {
	return "<iframe frameborder='0' scrolling='no' src='tip_iframe.html?tag=".$tag."&path=".$path."' width='".$w."' height='".$h."'></iframe>";
}

function html_get_textarea($name, $content, $enable, $rows, $cols) {
	$disabled = $enable ? "" : 'disabled="disabled"';
	$content = trim($content);
	return <<<C
<textarea {$disabled} name="{$name}[]" rows="{$rows}" cols="{$cols}">{$content}</textarea>
C;
}

function html_expand_arrs_json_str($js_name, $values, $className){
	if(!isset($values[0])){
		$values[0] = '无';
	}
	$content = "<script>";
	$content .= $js_name.'=['.PHP_EOL;
	foreach ($values as $key => $value) {
		$content .= "{value:'".$value."', id:'".$key."'},".PHP_EOL;
	}
	$content .= '];'.PHP_EOL;
	$content .= <<<C
	$('.{$className}').autocomplete({
	lookup: {$js_name}, 
	minChars: 0,
	onSelect: function(s){
		$('#' + this.id + '_hidden')[0].setAttribute('value', s.id);
	}
});
C;
$content .= "</script>";
	return $content;
}

function html_get_auto_completed_input($id, $className, $values, $now = 0, $width = 60){
	$content = '';
	if(!isset($values[$now])){
		$now = 0;
		$values[0] = '无';
	}
	$content .= "<input id='{$id}' type='text' value='{$values[$now]}' class='{$className}' style='width:{$width}px' />";
	$content .= "<input name='{$className}[]' id='{$id}_hidden' type='hidden' value='{$now}' />";
	return $content;
}

?>