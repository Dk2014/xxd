<?php

$links = array(
	array('text' => '邮件', 'id' => 'mail'),
	array('text' => '模块公告', 'id' => 'mail_announcement'),
);

function parameters_render( $row ) {
	return '<span >'. preg_replace("/\\n/", ", ", $row['parameters']) .'</span>';
}

function parameters_editor( $row ) {
	$html = '<textarea name="parameters[]" rows="3" cols="20">'.$row['parameters'].'</textarea>';
	return $html;
}

function content_render( $row ) {
	return '<span>'.$row['content'].'</span>';
}

function content_editor( $row ) {
	return '<textarea name="content[]" rows="3" cols="40">'.$row['content'].'</textarea>';
}

function item_id_render($row) {
	global $allItem;
	return item_render($row, 'item_id');
}

function item_id_editor($row) {
	global $allItem;
	return item_editor($row, 'item_id');
}

function item_render($row, $name) {
	global $allItem;

	if (empty($row[$name])) {
		return '';
	}
	return $allItem[$row[$name]];
}



function item_editor($row, $name) {
	global $allItem;
	// var_dump($name,$allItem);
	$html = '<select name="'.$name.'[]" >';
	$html .= '<option value="0" selected="selected">无</option>';
	foreach ($allItem as $key => $value) {
		if ( isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

$sql = "select `id`, `name` from `item` where `type_id`!='10' and `type_id`!='11'";
$allItemTmp = db_query($db, $sql);
foreach ($allItemTmp as $value) {
	$allItem[$value['id']] = $value['name'];
}

$expire_time = array(
	0 => "过期删除",
	1 => "无附件已阅读",
);
function mail_expire_time_render($row){
	global $expire_time;
	return $expire_time[$row['expire_time']];
}

function mail_expire_time_editor($row) {
	global $expire_time;

	$code = '';
	if ($row != null) {
		$code .= '<select name="expire_time[]" >';
	foreach ($expire_time as $key => $value){

		if ($key == $row['expire_time']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="expire_time[]" >';
		$code .= '<option value="0" selected="selected">过期删除</option>';
		$code .= '<option value="1" >无附件已阅读删除</option>';
		$code .= '</select>';
	}
	return $code;
}
	

function mail_type_render($row){
	global $types;
	return $types[$row['type']];
}

function mail_type_editor($row){
	global $types;

	$code = '';
	if ($row != null) {
		$code .= '<select name="type[]" >';
	foreach ($types as $key => $value){

		if ($key == $row['type']) {
			$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
		} else {
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
		$code .= '</select>';
	} else {
		$code .= '<select name="type[]" >';
		$code .= '<option value="0" selected="selected">邮件</option>';
		$code .= '<option value="1" >模块公告</option>';
		$code .= '</select>';
	}
	
	return $code;
}

// 任务类型
$types[0] = '邮件';
$types[1] = '模块公告';

?>
