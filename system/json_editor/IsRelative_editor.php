<?php

function json_IsRelative_editor_css($init, $path, $value,
                                    $name_mapping, $editor_mapping){
    return array();
}

function json_IsRelative_editor_js($init, $path, $value,
                                   $name_mapping, $editor_mapping){
    return array('IsRelative_editor.js'=>'IsRelative_editor.js');
}

function json_IsRelative_editor_html($init, $path, $value,
                                     $name_mapping, $editor_mapping){
    $selectedNegative="";
    $selectedTrue="";
    $selectedFalse="";
    if(!isset($value) || $value==-1)
        $selectedNegative="selected";
    else if($value==0)
        $selectedFalse="selected";
    else
        $selectedTrue="selected";
    return

<<<iSrElAtIvE_EdItOr_hTmL
    <input class="json_IsRelative_editor" type="hidden"
           name="type" value="number" />
    <select class="json_IsRelative_editor" name="value">
            <option value="-1" $selectedNegative>-1</option>
            <option value="0" $selectedFalse>0</option>
            <option value="1" $selectedTrue>1</option>
    </select>
iSrElAtIvE_EdItOr_hTmL;

}
