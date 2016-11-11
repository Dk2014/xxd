<?php
include_once "common_links.php";

function sword_soul_name_render($row) {
	return '<a href="?p=sword_soul_level&m='.$row['id'].'">'.$row['name'].'</a>';
}

function sword_soul_type_render($row) {
	global $allSwordSoulType;

	if (isset($allSwordSoulType[$row['type_id']])) {
		return $allSwordSoulType[$row['type_id']];
	}
}

function sword_soul_type_editor($row) {
	global $allSwordSoulType;
	$name = "type_id";
	$html = '<select name="'.$name.'[]" >';
	$html .= '<option value="0" selected="selected">无</option>';
	foreach ($allSwordSoulType as $key => $value) {
		if ( isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

function sword_soul_quality_render($row) {
	global $qualityConfig;

	return $qualityConfig[$row['quality']];
}

function sword_soul_quality_editor($row) {
	global $qualityConfig;
	$name = "quality";
	$html = '<select name="'.$name.'[]" >';
	foreach ($qualityConfig as $key => $value) {
		if ( isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

function get_all_sword_soul_fragment() {
	global $db;
	$all_sword_soul_fragments = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id` = 14");
	foreach ($tmp as $value)
	{
		$all_sword_soul_fragments[$value['id']] = $value['name'];
	}

	return $all_sword_soul_fragments;
}

$all_sword_soul_fragments = get_all_sword_soul_fragment();
$all_sword_soul_fragments[0] = '无';



function sword_soul_fragment_id_render($row) {
	global $all_sword_soul_fragments;
	return $all_sword_soul_fragments[$row['fragment_id']];
}

function sword_soul_fragment_id_editor($row) {
	global $all_sword_soul_fragments;
	$name = "fragment_id";
	$html = '<select name="'.$name.'[]" >';
	foreach ($all_sword_soul_fragments as $key => $value) {
		if ( isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}
	$html .= '</select>';
	return $html;
}

function only_exchange_render($row) {
	global $onlyExchange;

	return $onlyExchange[$row['only_exchange']];
}

function only_exchange_editor($row) {
	global $onlyExchange;

	$name = "only_exchange";
	$html = '<select name="'.$name.'[]" >';
	foreach ($onlyExchange as $key => $value) {
		if ( isset($row[$name]) && $row[$name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}
	$html .= '</select>';

	return $html;
}

function del_check($rowid) {
	global $db;
	$sql = "select * from sword_soul_level where sword_soul_id = {$rowid}";
	$datas = db_query($db, $sql);
	if (count($datas) > 0) {
		return false;
	}
	return true;
}

function sql_where() {
	return " order by `id`";
}

$allSwordSoulType = array();
$sword_soul_types = db_query($db, "select * from sword_soul_type");
foreach ($sword_soul_types as $value) {
	$allSwordSoulType[$value['id']] = $value['name'];
}

$qualityConfig = array();
$allQualityRes = db_query($db, "SELECT * FROM sword_soul_quality");
foreach ($allQualityRes as $key => $value) {
	$qualityConfig[$value['id']] = $value['name'];
}

$onlyExchange = array(0 => '否', 1 => '是');

$pconfig = array(
	'title'   => '剑心',
	'table'   => 'sword_soul',
	'links'   => $sword_soul_links,
	'columns' => array(
		array('name' => 'name', 'text' => '剑心名称', 'render' => 'sword_soul_name_render', 'width' => '120px'),
		array('name' => 'type_id', 'text' => '剑心属性', 
			'render' => 'sword_soul_type_render',
			'editor' => 'sword_soul_type_editor', 'width' => '120px',
		),
		array('name' => 'quality', 'text' => '品质',
			'render' => 'sword_soul_quality_render',
			'editor' => 'sword_soul_quality_editor', 'width' => '120px'
		),
		array('name' => 'fragment_id', 'text' => '兑换需求碎片', 'width' => '90px',
			'render' => 'sword_soul_fragment_id_render',
			'editor' => 'sword_soul_fragment_id_editor',
		),
		array('name' => 'fragment_num', 'text' => '兑换需求碎片数量', 'width' => '90px'),
		array('name' => 'only_exchange', 'text' => '只能兑换获得', 
			'render' => 'only_exchange_render',	'editor' => 'only_exchange_editor', 
			'width' => '90px'),
		array('name' => 'sign', 'text' => '资源标识', 'width' => '80px'),
		array('name' => 'desc', 'text' => '描述'),
	),
	'where' => 'sql_where',
	'del_check' => 'del_check',
);
?>
