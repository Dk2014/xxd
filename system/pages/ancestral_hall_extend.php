<?php

$extend_columns = array(
    'worship_type' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
        'range' => array('params' => array()),
    ),
    
);

//上香类型,1:白檀香 2:苏合香 3：天木香 
$all_type = array(
    0 => '无',
    1 => '白檀香',
    2 => '苏合香',
    3 => '天木香',
);

function render_worship_type($column_name, $column_val, $row, $data) {
    global $all_type;
    return $all_type[$column_val];
}

function editor_worship_type($column_name, $column_val, $row, $data) {
    global $all_type;

    $data = $all_type;
    $field_name = 'worship_type';

    $html = "<select class={$field_name} name={$field_name}[]>";

    foreach ($data as $key => $value) {
        if (isset($row[$field_name]) && $row[$field_name] == $key) {
            $html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
        } else {
            $html .= '<option value="'.$key.'">'.$value.'</option>';
        }
    }

    $html .= '</select>';
    return $html;
}

function range_worship_type(){
    global $all_type;
    return $all_type;
}

?>
