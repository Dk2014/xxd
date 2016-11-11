<?php

function is_assoc($array){
    return (bool)count(array_filter(array_keys($array), 'is_string'));
}

function swapKV($array){
    $result=array();
    foreach($array as $k => $v)
        $result[$v]=$k;
    return $result;
}

function array2set($array){
    return swapKV($array);
}

function split_str($str, $delim){
    $result=explode($delim, $str);
    if(count($result)<1)
        $result=array($str);
    return $result;
}

function map_path_id($editor_mapping, $path_id){
    $keys=split_str($path_id, '.');
    array_splice($keys, 0, 1);
    $bit_size=count($keys);
    $matcher=$path_id;
    for($mask=0;
        $mask<(1<<$bit_size) && !isset($editor_mapping[$matcher]);
        $mask++){
            $matcher
            =array_reduce($keys, function($snow_ball, $key){
                $snow_ball['mask_iter']>>=1;
                $snow_ball['result'].='.'
                        .(($snow_ball['mask'] & $snow_ball['mask_iter'])
                                ?'*'
                                :$key);
                return $snow_ball;
            }, array(
                    'mask_iter'=>(1<<$bit_size),
                    'mask'=>$mask,
                    'result'=>''
            ))['result'];
    }
    if(isset($editor_mapping[$matcher]))
        return $editor_mapping[$matcher];
    else
        return false;
}

function get_json_element($json, $keys, $offset, $length){
    $json_recurse=$json;
    for($i=$offset; $i<$offset+$length; $i++){
        $json_recurse=$json_recurse[$keys[$i]];
    }
    return $json_recurse;
}

function set_json_element($json, $keys, $offset, $length, $edit_f, $args){
    if($length==0)
        return $edit_f($args, $json, null);
    if($length==1){
        return $edit_f($args, $json, $keys[$offset]);
    }else{
        $json[$keys[$offset]]=set_json_element($json[$keys[$offset]],
                                               $keys,
                                               $offset+1,
                                               $length-1,
                                               $edit_f, $args);
        return $json;
    }
}

function json_type_convert($type, $value){
    switch($type){
        case 'null':
            return null;
        case 'number':
            return floatval($value);
        case 'boolean':
            return strtoupper($typed_value['value'])=='TRUE';
        case 'rawjson':
            return json_decode($value, true);
        default: //string
            return $value;
    }
}
