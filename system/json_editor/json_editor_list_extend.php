<?php

require_once 'json_timestamp_editor.php';
require_once 'json_gameitem_editor.php';
require_once 'IsRelative_editor.php';

function version_increment($json){
    $json['version']+=1;
    return $json;
}

$after_json_editing_done='version_increment';
