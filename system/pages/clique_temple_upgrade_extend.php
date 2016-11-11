<?php

$extend_columns = array(
/*   '字段' => 配置 */

'desc' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data' => array(),
    ),
'cur_temple_desc' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data' => array(),
    ),
'next_temple_desc' => array(
        'editor'    => array('params' => array()),
        'render'    => array('params' => array()),
        'data' => array(),
    ),
);


function editor_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_cur_temple_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_cur_temple_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, false, 4, 30);
}


function editor_next_temple_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_next_temple_desc($column_name, $column_val, $row, $data){
    return html_get_textarea($column_name, $column_val, false, 4, 30);
}

?>
