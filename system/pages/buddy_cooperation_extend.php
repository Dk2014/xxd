<?php

$extend_columns = array(
        'role_id1' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range'		=> array('params' => array()),
        ),
        'role_id2' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range'		=> array('params' => array()),
        )
);

$all_role = get_all_role();

function sql_where($params) {
	return "  order by `role_id1` asc";
}


function editor_role_id1($column_name, $column_val, $row, $data){
    global $all_role;
    $item_name = '';
    if (isset($all_role[$column_val])) {
        $item_name = $all_role[$column_val];
    }
    return editor_single_item($item_name, $column_val, $column_name);
}

function render_role_id1($column_name, $column_val, $row, $data){
    global $all_role;

    return (isset($all_role[$column_val]) ? $all_role[$column_val] : "0");
}

function range_role_id1(){
    global $all_role;
    $tempitems = $all_role;
    $tempitems[0] = '无';
    return $tempitems;
}

function editor_role_id2($column_name, $column_val, $row, $data){
    global $all_role;
    $item_name = '';
    if (isset($all_role[$column_val])) {
        $item_name = $all_role[$column_val];
    }
    return editor_single_item($item_name, $column_val, $column_name);
}

function render_role_id2($column_name, $column_val, $row, $data){
    global $all_role;

    return (isset($all_role[$column_val]) ? $all_role[$column_val] : "0");
}

function range_role_id2(){
    global $all_role;
    $tempitems = $all_role;
    $tempitems[0] = '无';
    return $tempitems;
}

function js_function($params) {
    global $all_role;

    $html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

    $html .= choose_single_item($all_role, 'role_id1');
    $html .= choose_single_item($all_role, 'role_id2');
    return $html;
}
