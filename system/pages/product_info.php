<?php

function is_month_card_render($row){
	$is_month_cards = array('否', '是');
	return $is_month_cards[$row['is_month_card']];
}

function is_month_card_editor($row){
	$is_month_cards = array('否', '是');

	$code = '';
	$code .= '<select name="is_month_card[]" >';
	if ($row != null) {
		foreach ($is_month_cards as $key => $value){
			if ($key == $row['is_month_card']) {
				$code .= '<option value="'.$key.'"  selected="selected">'.$value.'</option>';
			} else {
				$code .= '<option value="'.$key.'" >'.$value.'</option>';
			}
		}
	} else {
		foreach ($is_month_cards as $key => $value){
			$code .= '<option value="'.$key.'" >'.$value.'</option>';
		}
	}
	$code .= '</select>';
	return $code;
}


$pconfig = array(
	'title'       => '内购产品信息',
	'table'       => 'product_info',
	'links'       => array(),
	'columns' => array(
		array('name' => 'app_store_product_id', 'text' => 'app store 产品ID', 'width' => '80px'),
		array('name' => 'google_play_product_id', 'text' => 'google play 产品ID', 'width' => '80px'),
		array('name' => 'price', 'text' => '价格', 'width' => '80px'),
		array('name' => 'ingot', 'text' => '元宝数量', 'width' => '80px'),
		array('name' => 'unit', 'text' => '单位', 'width' => '80px'),
		array('name' => 'is_month_card', 'text' => '月卡', 'width' => '80px',
			'editor' => "is_month_card_editor",
			'render' => "is_month_card_render",
		),
		array('name' => 'first_presentation', 'text' => '首充赠送元宝数量', 'width' => '80px'),
		array('name' => 'regular_presentation', 'text' => '常规赠送元宝数量', 'width' => '80px'),

	),
);
?>
