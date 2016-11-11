<?php

$extend_columns = array(

        'award_item1_id' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range'		=> array('params' => array()),
        ),
        'award_item2_id' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range' =>array('params' => array()),
        ),
        'award_item3_id' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range' =>array('params' => array()),
        ),
        'award_item4_id' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range' =>array('params' => array()),
        ),
        
);

$allItem = get_all_item();

function editor_award_item1_id($column_name, $column_val, $row, $data){
    global $allItem;

    $item_name = '';
    if (isset($allItem[$column_val])) {
        $item_name = $allItem[$column_val];
    }

    return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item1_id($column_name, $column_val, $row, $data){
    global $allItem;

    return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}
function editor_award_item2_id($column_name, $column_val, $row, $data){
    global $allItem;

    $item_name = '';
    if (isset($allItem[$column_val])) {
        $item_name = $allItem[$column_val];
    }

    return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item2_id($column_name, $column_val, $row, $data){
    global $allItem;

    return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}

function editor_award_item3_id($column_name, $column_val, $row, $data){
    global $allItem;

    $item_name = '';
    if (isset($allItem[$column_val])) {
        $item_name = $allItem[$column_val];
    }

    return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item3_id($column_name, $column_val, $row, $data){
    global $allItem;

    return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function editor_award_item4_id($column_name, $column_val, $row, $data){
    global $allItem;

    $item_name = '';
    if (isset($allItem[$column_val])) {
        $item_name = $allItem[$column_val];
    }

    return editor_single_item($item_name, $column_val, $column_name);
}

function render_award_item4_id($column_name, $column_val, $row, $data){
    global $allItem;

    return (isset($allItem[$column_val]) ? $allItem[$column_val] : "无");
}
function range_award_item1_id(){
    global $allItem;
    $all_item_temp = $allItem;
    $all_item_temp[0] = '无';
    return $all_item_temp;
}
function range_award_item2_id(){
    global $allItem;
    $all_item_temp = $allItem;
    $all_item_temp[0] = '无';
    return $all_item_temp;
}
function range_award_item3_id(){
    global $allItem;
    $all_item_temp = $allItem;
    $all_item_temp[0] = '无';
    return $all_item_temp;
}
function range_award_item4_id(){
    global $allItem;
    $all_item_temp = $allItem;
    $all_item_temp[0] = '无';
    return $all_item_temp;
}

function jsFunction($params) {
    global $allItem;

    $html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

    $html .= choose_single_item($allItem, 'award_item1_id');
    $html .= choose_single_item($allItem, 'award_item2_id');
    $html .= choose_single_item($allItem, 'award_item3_id');
    $html .= choose_single_item($allItem, 'award_item4_id');
    return $html;
}

function sql_where($params) {
    return "WHERE `town_id`={$params['m']} ORDER BY id";
}

function sql_insert($params) {
    return "`town_id` = {$params['m']}";
}

function location($params){
    global $db;
    $html = '';

    $sql = "select `id`, `name` from town where id = {$params['m']}";
    $town = db_query($db, $sql);
    if (count($town) == 0) {
        return $html;
    }
    $html .= '<a href="?p=town">所有城镇</a> -> '.$town[0]['name'];

    return $html;
}
