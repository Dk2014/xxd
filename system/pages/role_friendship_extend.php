<?php

$extend_columns = array(
        'favourite_item' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range'		=> array('params' => array()),
        )
);

function location($params){
    global $db;
    $html = '';
    $sql = "select `id`, `name` from role where id = {$params['m']}";
    $role = db_query($db, $sql);
    if (count($role) == 0) {
        return $html;
    }

    $html = '<a href="?p=role">'.$role[0]['name'].'</a> -> ' . $html;

    return $html;
}

function sql_where($params) {
    return "where `role_id`={$params['m']} order by `friendship_level`";
}

function sql_insert($params) {
    return "`role_id` = {$params['m']}";
}

$allItem = get_all_item();

function editor_favourite_item($column_name, $column_val, $row, $data){
    global $allItem;

    $item_name = '';
    if (isset($allItem[$column_val])) {
        $item_name = $allItem[$column_val];
    }

    return editor_single_item($item_name, $column_val, $column_name);
}

function range_favourite_item(){
    global $allItem;
    $tempitems = $allItem;
    $tempitems[0] = '无';
    return $tempitems;
}

function render_favourite_item($column_name, $column_val, $row, $data){
    global $allItem;

    return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}

function render_level_color($row){
    return '<span style="color:#'.preg_replace('/0x/i', "", $row['level_color']).'">'.$row['level_color'].'</a>';
}

function js_function($params) {
    global $allItem;

    $html = '
$("select").change( function(){
		$(this).css("color","#0033FF");
});';

    $html .= choose_single_item($allItem, 'favourite_item');
    return $html;
}
