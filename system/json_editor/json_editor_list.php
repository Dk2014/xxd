<?php

require_once 'common_declare.php';


function json_rawjson_editor_js($init, $path, $value,
                                $name_mapping, $editor_mapping){
    return array(
            'losing_focus_means_submit.js'=>'losing_focus_means_submit.js'
    );
}

function json_rawjson_editor_css($init, $path, $value,
                                 $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_rawjson_editor_html($init, $path, $value,
                                  $name_mapping, $editor_mapping){
    $json_content=json_encode(isset($value)?$value:$init,
                              JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE);
    return

<<<RaWjSoN_eDiToR_hTmL
    <input class="json_rawjson_editor" type="hidden"
           name="type" value="rawjson" />
    <textarea class="json_rawjson_editor" name="value">$json_content</textarea>
RaWjSoN_eDiToR_hTmL;

}

function json_href_editor_js($init, $path, $value,
                             $name_mapping, $editor_mapping){
    return array('json_href_editor.js'=>'json_href_editor.js');
}

function json_href_editor_css($init, $path, $value,
                              $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_href_editor_html($init, $path, $value,
                               $name_mapping, $editor_mapping){
    $path_keys=split_str($path, '.');
    $content=map_path_id($name_mapping, $path).(isset($value)?'...':'');
    $class=(isset($value)?'json_href_editor':'json_href_editor_empty');
    return

<<<RaWjSoN_eDiToR_hTmL
    <a class="$class" href="#">$content</a>
RaWjSoN_eDiToR_hTmL;

}

$json_table_editor_js_first_call=true;
$json_table_editor_css_first_call=true;
$json_table_editor_html_first_call=true;

function json_table_editor_js($init, $path, $value,
                              $name_mapping, $editor_mapping){
    global $json_table_editor_js_first_call;
    if(!$json_table_editor_js_first_call)
        return json_href_editor_js($init, $path, $value,
                                   $name_mapping, $editor_mapping);
    else
        $json_table_editor_js_first_call=false;
    $record_rows=$value;
    $ignore_row_ids=$init['ignore_row_ids'];
    foreach($ignore_row_ids as $id ){
        unset($record_rows[$id]);
    }
    
    $column_set=array();
    $specified_columns=$init['columns'];
    $column_set
        =array_reduce($specified_columns, function($snow_ball, $item){
            $snow_ball[$item]=$item;
            return $snow_ball;
        }, $column_set);
    $column_set
        =array_reduce($record_rows, function($snow_ball, $item){
            if(!empty($item))
                foreach($item as $k => $_)
                    $snow_ball[$k]=$k;
            return $snow_ball;
        }, $column_set);
    $ignore_columns=$init['ignore_columns'];
    foreach($ignore_columns as $column){
        unset($column_set[$column]);
    }
    
    $files=array(
            'json_table_editor.js'=>'json_table_editor.js'
    );
    foreach($column_set as $column){
        $cell_path=$path.".*.$column";
        $cell_editor_pair=map_path_id($editor_mapping, $cell_path);
        $cell_editor_js_f
            ='json_'.$cell_editor_pair['editor'].'_editor_js';
        $cell_editor_js=$cell_editor_js_f($cell_editor_pair['init'],
                                          $cell_path, null,
                                          $name_mapping, $editor_mapping);
        $files=array_merge($files, $cell_editor_js);
    }
    
    return $files;
}

function json_table_editor_css($init, $path, $value,
                               $name_mapping, $editor_mapping){
    global $json_table_editor_css_first_call;
    if(!$json_table_editor_css_first_call)
        return json_href_editor_css($init, $path, $value,
                                    $name_mapping, $editor_mapping);
    else
        $json_table_editor_css_first_call=false;
    $record_rows=$value;
    $ignore_row_ids=$init['ignore_row_ids'];
    foreach($ignore_row_ids as $id ){
        unset($record_rows[$id]);
    }
    
    $column_set=array();
    $specified_columns=$init['columns'];
    $column_set
        =array_reduce($specified_columns, function($snow_ball, $item){
            $snow_ball[$item]=$item;
            return $snow_ball;
        }, $column_set);
    $column_set
        =array_reduce($record_rows, function($snow_ball, $item){
            if(!empty($item))
                foreach($item as $k => $_)
                    $snow_ball[$k]=$k;
            return $snow_ball;
        }, $column_set);
    $ignore_columns=$init['ignore_columns'];
    foreach($ignore_columns as $column){
        unset($column_set[$column]);
    }
    
    $files=array();
    foreach($column_set as $column){
        $cell_path=$path.".*.$column";
        $cell_editor_pair=map_path_id($editor_mapping, $cell_path);
        $cell_editor_css_f
            ='json_'.$cell_editor_pair['editor'].'_editor_css';
        $cell_editor_css=$cell_editor_css_f($cell_editor_pair['init'],
                                            $cell_path, null,
                                            $name_mapping, $editor_mapping);
        $files=array_merge($files, $cell_editor_css);
    }
    
    return $files;
}

function json_table_editor_html($init, $path, $value,
                                $name_mapping, $editor_mapping){
    global $json_table_editor_html_first_call;
    if(!$json_table_editor_html_first_call)
        return json_href_editor_html($init, $path, $value,
                                     $name_mapping, $editor_mapping);
    else
        $json_table_editor_html_first_call=false;
    $record_rows=$value;
    $ignore_row_ids=$init['ignore_row_ids'];
    foreach($ignore_row_ids as $id ){
        unset($record_rows[$id]);
    }
    
    $column_set=array();
    $specified_columns=$init['columns'];
    $column_set
        =array_reduce($record_rows, function($snow_ball, $item){
            if(!empty($item))
                foreach($item as $k => $_)
                    $snow_ball[$k]=$k;
            return $snow_ball;
        }, $column_set);
    $column_set
        =array_reduce($specified_columns, function($snow_ball, $item){
            $snow_ball[$item]=$item;
            return $snow_ball;
        }, $column_set);
    $ignore_columns=$init['ignore_columns'];
    foreach($ignore_columns as $column){
        unset($column_set[$column]);
    }
    
    $titles_html='<th>[id]</th>';
    foreach($column_set as $title_key){
        $maybe_title=map_path_id($name_mapping, $path.".*.$title_key");
        $titles_html.="<th><input class=\"json_table_column\" type=\"hidden\" value=\"$title_key\" />";
        $titles_html.=($maybe_title?$maybe_title:$title_key);
        $titles_html.="</th>";
    }
    $titles_html.='<th>[op]</th>';
    
    $content_html='';
    foreach($record_rows as $id => $content){
        $content_html.='<tr>';
        $content_html.="<td class=\"json_table_row_id\">$id</td>";
        foreach($column_set as $column){
            $cell_value=(isset($content[$column])?$content[$column]:null);
            $cell_path=$path.".$id.$column";
            $cell_editor_pair=map_path_id($editor_mapping, $cell_path);
            $cell_editor_html_f
                ='json_'.$cell_editor_pair['editor'].'_editor_html';
            $cell_editor_html=$cell_editor_html_f($cell_editor_pair['init'],
                                                  $cell_path, $cell_value,
                                                  $name_mapping,
                                                  $editor_mapping);
            $content_html.="<td>$cell_editor_html</td>";
        }
        $content_html.='<td><button class="delete_table_row">rm</button></td>';
        $content_html.='</tr>';
    }

    $dummy_id=null;
    if(!isset($value) || empty($value))
        $dummy_id=0;
    else
        $dummy_id=(is_assoc($value)?'':max(array_keys($value))+1);
    
    $content_html.='<tr class="table_addition_empty_row">';
    $content_html.="<td class=\"json_table_row_id\">$dummy_id</td>";
    foreach($column_set as $column){
        $cell_value=null;
        $cell_path=$path.".*.$column";
        $cell_editor_pair=map_path_id($editor_mapping, $cell_path);
        $cell_editor_html_f
        ='json_'.$cell_editor_pair['editor'].'_editor_html';
        $cell_editor_html=$cell_editor_html_f($cell_editor_pair['init'],
                                              $cell_path, $cell_value,
                                              $name_mapping,
                                              $editor_mapping);
        $content_html.="<td>$cell_editor_html</td>";
    }
    $content_html.='<td><button class="delete_table_row">rm</button></td>';
    $content_html.='</tr>';
    
    $column_amount=count($column_set)+2;
    
    return

<<<TaBlE_eDiToR_HtMl
    <table class="json_table_editor">
        <tr>$titles_html</tr>
        $content_html
        <tr><td colspan="$column_amount">
            <a class="new_table_row" href="#">New...</a>
        </td></tr>
    </table>
TaBlE_eDiToR_HtMl;

}

function json_list_editor_js($init, $path, $value,
                             $name_mapping, $editor_mapping){
    return array('json_list_editor.js'=>'json_list_editor.js');
}

function json_list_editor_css($init, $path, $value,
                              $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_list_editor_html($init, $path, $value,
                               $name_mapping, $editor_mapping){
    $exclude=array2set($init['exclude']);
    $html_content='';
    foreach($value as $key => $content){
        if(!isset($exclude[$key])){
            $json_content=(is_array($content)
                    ?json_encode($content,
                                 JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE)
                    :$content);
            $html_content.=

<<<hTmL_CoNtEnT
    <tr>
        <td><a class="json_list_item_id" href="#">$key</a></td>
        <td><textarea readonly>$json_content</textarea></td>
        <td><button class="delete_list_row">rm</button></td>
    </tr>
hTmL_CoNtEnT;

        }
    }
    return

<<<LiSt_eDiToR_HtMl
    <table class="json_list_editor">
        <tr><th>[id]</th><th>[raw json]</th><th>[op]</th></tr>
        $html_content
        <tr><td colspan="3">
            <a class="new_list_row" href="#">New...</a>
        </td></tr>
    </table>
LiSt_eDiToR_HtMl;

}

function json_number_editor_js($init, $path, $value,
                               $name_mapping, $editor_mapping){
    return array(
            'losing_focus_means_submit.js'=>'losing_focus_means_submit.js'
    );
}

function json_number_editor_css($init, $path, $value,
                                $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_number_editor_html($init, $path, $value,
                                 $name_mapping, $editor_mapping){
    $content=(isset($value)?$value:$init['default']);
    return

<<<nUmBeR_EdItOr_hTmL
    <input class="json_number_editor" type="hidden"
           name="type" value="number" />
    <input class="json_number_editor" type="number"
           name="value" value="$content" />
nUmBeR_EdItOr_hTmL;

}

function json_text_editor_js($init, $path, $value,
                             $name_mapping, $editor_mapping){
    return array(
            'losing_focus_means_submit.js'=>'losing_focus_means_submit.js'
    );
}

function json_text_editor_css($init, $path, $value,
                              $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_text_editor_html($init, $path, $value,
                               $name_mapping, $editor_mapping){
    $content=(isset($value)?$value:$init['default']);
    return

<<<tExT_EdItOr_hTmL
    <input class="json_text_editor" type="hidden" name="type" value="string" />
    <input class="json_text_editor" type="text" name="value" value="$content" />
tExT_EdItOr_hTmL;

}

function json_content_editor_js($init, $path, $value,
                                $name_mapping, $editor_mapping){
    return array(
            'losing_focus_means_submit.js'=>'losing_focus_means_submit.js'
    );
}

function json_content_editor_css($init, $path, $value,
                                 $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_content_editor_html($init, $path, $value,
                                  $name_mapping, $editor_mapping){
    $content=(isset($value)?$value:$init['default']);
    return

<<<tExT_EdItOr_hTmL
    <input class="json_content_editor" type="hidden" name="type" value="string" />
    <textarea class="json_content_editor" name="value">$content</textarea>
tExT_EdItOr_hTmL;

}

function json_bool_editor_js($init, $path, $value,
                             $name_mapping, $editor_mapping){
    return array(
            'losing_focus_means_submit.js'=>'losing_focus_means_submit.js'
    );
}

function json_bool_editor_css($init, $path, $value,
                              $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_bool_editor_html($init, $path, $value,
                               $name_mapping, $editor_mapping){
    $content=(isset($value)?$value:$init['default']);
    $is_checked=(strtoupper($content)=='TRUE'?'checked':'');
    return

<<<bOoL_EdItOr_hTmL
    <input class="json_bool_editor" type="hidden" name="type" value="boolean" />
    <input class="json_bool_editor" type="checkbox"
           name="value" value="true" $is_checked/>
bOoL_EdItOr_hTmL;

}


function json_null_editor_js($init, $path, $value,
                             $name_mapping, $editor_mapping){
    return array('dummy.js'=>'dummy.js');
}

function json_null_editor_css($init, $path, $value,
                              $name_mapping, $editor_mapping){
    return array('dummy.css'=>'dummy.css');
}

function json_null_editor_html($init, $path, $value,
                               $name_mapping, $editor_mapping){
    return

<<<NuLl_eDiToR_HtMl
    <input type="hidden" name="type" value="null" />
    <input type="text" name="value" value="null" readonly/>
NuLl_eDiToR_HtMl;

}
