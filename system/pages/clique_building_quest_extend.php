<?php

$extend_columns = array(
/*   '字段' => 配置 */
    'desc' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
    ),
    
    'require_open_day' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
    ),
    
    'class' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array( '1'=>'宗祠','2'=>'回春堂','3'=>'神兵堂', '4'=>'金刚堂','5'=>'总舵','6'=>'钱庄', ),
    ),
  

    'award_item1_id' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
        'range'     => array('params' => array()),
    ),
    'award_item2_id' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
        'range' =>array('params' => array()),
    ),
    'award_item3_id' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
        'range' =>array('params' => array()),
    ),
    'award_item4_id' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data'      => array(),
        'range' =>array('params' => array()),
    ),

);

$allItem = get_all_item();

function render_class($column_name, $column_val, $row, $data){
    return $data[$column_val];
}

function editor_class($column_name, $column_val, $row, $data){
    return html_get_select($column_name, $data, $column_val);
}

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

function editor_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, false, 4, 30);
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
    if(isset($_GET['m'])){
        return " where `class`={$_GET['m']} ";
    }
    return " order by `id`, `class` ASC";
}

?>
