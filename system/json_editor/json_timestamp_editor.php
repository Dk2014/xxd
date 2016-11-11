<?php

function json_timestamp_editor_js($init, $path, $value,
                                  $name_mapping, $editor_mapping){
    return array(
            'json_timestamp_editor.js'=>'json_timestamp_editor.js',
            'jquery.datetimepicker.js'=>'jquery.datetimepicker.js'
    );
}

function json_timestamp_editor_css($init, $path, $value,
                                   $name_mapping, $editor_mapping){
    return array(
            'jquery.datetimepicker.css'=>'jquery.datetimepicker.css'
    );
}

function json_timestamp_editor_html($init, $path, $value,
                                    $name_mapping, $editor_mapping){
    $value=(isset($value)?$value:0);
    return
<<<tImEsTaMp_EdItOr_hTmL
    <input class="json_timestamp_editor" type="hidden" name="type" value="number" />
    <input class="json_timestamp_editor" type="text" name="value" value="$value" />
tImEsTaMp_EdItOr_hTmL;
}
