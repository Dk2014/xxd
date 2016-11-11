<?php

require_once '../lib/global.php';


function json_gameitem_editor_js($init, $path, $value,
                                 $name_mapping, $editor_mapping){
    return array(
            'json_gameitem_editor.js.php'=>'json_gameitem_editor.js.php',
            '../jquery.autocomplete.js'=>'../jquery.autocomplete.js'
    );
}

function json_gameitem_editor_css($init, $path, $value,
                                  $name_mapping, $editor_mapping){
    return array();
}

function json_gameitem_editor_html($init, $path, $value,
                                   $name_mapping, $editor_mapping){
    $all_items=get_all_item();
    $value=(isset($value)?$value:0);
    $display=(isset($all_items[$value])?$all_items[$value]:'');
    
    return
<<<GaMeItEm_eDiToR_HtMl
    <input class="json_gameitem_editor" type="hidden" name="type" value="number" />
    <input class="json_gameitem_editor" type="hidden" name="value" value="$value" />
    <input class="json_gameitem_editor" type="text" name="dummy_dummy_display" value="$display" />
GaMeItEm_eDiToR_HtMl;
}
