<?php
include "common_links.php";

// 品质
$all_qualitys = array();
$all_qualitys[0] = '蓝色';
$all_qualitys[1] = '紫色';
$all_qualitys[2] = '金色';
$all_qualitys[3] = '橙色';

function quality_render($row){
	global $all_qualitys;
	return $all_qualitys[$row['quality']];
}

function quality_editor($row){
	global $all_qualitys;

	$code = '';
	$code .= '<select name="quality[]" >';
	if ($row != null) {
		foreach ($all_qualitys as $key => $value){
			if ($key == $row['quality']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($all_qualitys as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}

function sql_where($params) {
	return " order by `quality`, `star`";
}

$pconfig = array(
	'title'       => '魂侍进阶',
	'table'       => 'ghost_star',
	'links'       => $ghost_links,
	'columns' => array(
		array('name' => 'quality', 'text' => '品质', 'width' => '80px',
			'editor' => "quality_editor",
			'render' => "quality_render",
		),
		array('name' => 'star', 'text' => '星级', 'width' => '80px'),
		array('name' => 'need_fragment_num', 'text' => '需要碎片数量', 'width' => '80px'),
		array('name' => 'growth', 'text' => '成长值', 'width' => '80px'),
		array('name' => 'costs', 'text' => '费用', 'width' => '80px'),
		array('name' => 'health',       'text' => '基础生命',       'width' => '70px'),
		array('name' => 'attack',       'text' => '基础攻击',       'width' => '70px'),
		array('name' => 'defence',      'text' => '基础防御',       'width' => '70px'),
	),
	'where' => 'sql_where',
);
?>
