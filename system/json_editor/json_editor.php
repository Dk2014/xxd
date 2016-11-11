<?php // delaying calculation

$editor_config_file='editor_config.json';

require_once 'common_declare.php';
require_once 'json_editor_list.php';
if(file_exists('json_editor_list_extend.php')){
    require_once 'json_editor_list_extend.php';
}
/*
 * @Imported:
 * 
 * $editor2type[`name`];
 * 
 * function json_`name`_editor_js($init, $path, $value,
 *                                $name_mapping, $editor_mapping);
 * function json_`name`_editor_css($init, $path, $value,
 *                                 $name_mapping, $editor_mapping);
 * function json_`name`_editor_html($init, $path, $value,
 *                                  $name_mapping, $editor_mapping);
 */

session_start();

// @Guard $editor_config
if(!isset($_SESSION['editor_config'])){
    $editor_config_content=file_get_contents($editor_config_file);
    $_SESSION['editor_config']=json_decode($editor_config_content, true);
}

$editor_config=$_SESSION['editor_config'];

function get_name_mapping($editor_config){
    $name_mapping=array();
    $name_mapping_src=$editor_config['mapping'];
    foreach($name_mapping_src as $mapping_pair){
        $name_mapping[$mapping_pair['id']]=$mapping_pair['display'];
    }
    return $name_mapping;
}

function get_editor_mapping($editor_config){
    $editor_mapping=array();
    $editor_mapping_src=$editor_config['editor'];
    foreach($editor_mapping_src as $mapping_pair){
        $editor_mapping[$mapping_pair['id']]=array(
                'editor'=>$mapping_pair['editor'],
                'init'=>$mapping_pair['init']
        );
    }
    return $editor_mapping;
}

function edit_json_element_append($maybe_pair, $json, $key){
    $the_json=null;
    if(isset($key))
        $the_json=$json[$key];
    else
        $the_json=$json;
    if((empty($the_json) || is_assoc($the_json)) && isset($maybe_pair['key']))
        $the_json[$maybe_pair['key']]=$maybe_pair['value'];
    else
        array_push($the_json, $maybe_pair['value']);
    if(isset($key))
        $json[$key]=$the_json;
    else
        $json=$the_json;
    return $json;
}

function edit_json_element_update($value, $json, $key){
    $json[$key]=$value;
    return $json;
}

function edit_json_element_delete($_, $json, $key){
    if(is_assoc($json) || true) //simply synchronise with client's Javascript
        unset($json[$key]);
    else
        array_splice($json, intval($key), 1);
    return $json;
}

function iterate_json_array($arr, $f){
    foreach($arr as $key => $item)
        if(is_array($item))
            $arr[$key]=iterate_json_array($item, $f);
    return $f($arr);
}

function json_empty_check($arr){
    if(!is_assoc($arr)){
        $end=max(array_keys($arr))+1;
        for($i=$end; $i>0; $i--){
            if(!isset($arr[$i-1]))
                array_splice($arr, $i-1, 1);
        }
    }
    return $arr;
}

function tie_json_array($json){
    return iterate_json_array($json, 'json_empty_check');
}


if(!isset($_SESSION['editing_json'])){/*======================================*/


if(isset($_POST['edit'])){
    $json_content=$_POST['json_content'];
    $_SESSION['editing_json']=json_decode($json_content, true);
    header('Location:'.$_SERVER['PHP_SELF']);
    exit;
}else

/*Shit stuff*/{
    if(isset($_POST['load'])){
        $json_content=file_get_contents('../../event.json');
        $_SESSION['editing_json']=json_decode($json_content, true);
        header('Location:'.$_SERVER['PHP_SELF']);
        exit;
    }else
    
    if(isset($_POST['export'])){
        $json_content=$_POST['json_content'];
        
        //now this realy fucking shit stuff begin
        $version=json_decode($json_content, true)['version'];
        file_put_contents('../events/'.$version.'.json', $json_content);
        
        header('Location:'.$_SERVER['PHP_SELF']);
        exit;
    }
/*Shit stuff*/}

$json_content
=(isset($_POST['json_content'])
        ?$_POST['json_content']
        :(isset($_SESSION['json_content'])
                ?$_SESSION['json_content']
                :''));
unset($_SESSION['json_content']);

?>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>Keeping in mind that I'm your master!</title>
</head>
<body>
<div>
    <form action="<?=$_SERVER['PHP_SELF'];?>" method="post">
        <div>
            <!-- Standard terminal size is my lucky numbers! -->
            <textarea rows="40"
                      cols="80"
                      name="json_content"
             ><?=$json_content;?></textarea>
        </div>
        <div>
            <input type="submit" name="edit" value="Edit" />
            
            <!-- Shit stuff -->
            <input type="submit" name="load" value="Load" />
            <?=(isset($json_content) && strlen($json_content)>0)
                ?'<input type="submit" name="export" value="Export" />'
                :'';?>
            <!-- Shit stuff -->
            
        </div>
    </form>
</div>
</body>
</html>
<?php


}else{/*======================================================================*/


$editing_json=$_SESSION['editing_json'];

$name_mapping=get_name_mapping($editor_config);
$editor_mapping=get_editor_mapping($editor_config);

// using GET for specifying through url
$path
=(isset($_POST['path'])
        ?$_POST['path']
        :(isset($_GET['path'])
                ?$_GET['path']
                :""));

$path_keys=split_str($path, '.');

if(isset($_POST['done'])){
    $editing_json=tie_json_array($editing_json);
    if(isset($after_json_editing_done))
        $editing_json=$after_json_editing_done($editing_json);
    $_SESSION['json_content']=json_encode($editing_json,
                                          JSON_PRETTY_PRINT
                                          |JSON_UNESCAPED_UNICODE);
    unset($_SESSION['editing_json']);
    header('Location:'.$_SERVER['PHP_SELF']);
    exit;
}else

if(isset($_POST['append'])){
    $maybe_pair=array();
    if(isset($_POST['value'])){
        $type=(isset($_POST['type'])?$_POST['type']:'string');
        $maybe_pair['value']=json_type_convert($type, $_POST['value']);
    }else{
        $maybe_pair['value']=array();
    }
    if(isset($_POST['key'])){
        $maybe_pair['key']=$_POST['key'];
    }
    $editing_json=set_json_element($editing_json,
                                   $path_keys, 1, count($path_keys)-1,
                                   'edit_json_element_append', $maybe_pair);
    $_SESSION['editing_json']=$editing_json;
    exit;
}else

if(isset($_POST['update'])){
    $value=json_type_convert($_POST['type'], $_POST['value']);
    $editing_json=set_json_element($editing_json,
                                   $path_keys, 1, count($path_keys)-1,
                                   'edit_json_element_update', $value);
    $_SESSION['editing_json']=$editing_json;
    exit;
}else

if(isset($_POST['delete'])){
    $editing_json=set_json_element($editing_json,
                                   $path_keys, 1, count($path_keys)-1,
                                   'edit_json_element_delete', 'whatever');
    $_SESSION['editing_json']=$editing_json;
    exit;
}

$editing_json_part=get_json_element($editing_json,
                                    $path_keys, 1, count($path_keys)-1);

$page_name=map_path_id($name_mapping, $path);
$page_editor_pair=map_path_id($editor_mapping, $path);

?>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>Worship me, you puny!</title>

<!-- Public Javascript Library -->
<script src="../jquery.min.js"></script>
<link rel="stylesheet" href="../jquery-ui.css" />
<link rel="stylesheet" href="../style.css" />
<script src="../jquery-ui.min.js"></script>
<script src="common_function.js"></script>

<?php {
    $page_css4editor_f='json_'.$page_editor_pair['editor'].'_editor_css';
    $page_js4editor_f='json_'.$page_editor_pair['editor'].'_editor_js';
    $page_css4editor=$page_css4editor_f($page_editor_pair['init'],
                                        $path, $editing_json_part,
                                        $name_mapping, $editor_mapping);
    $page_js4editor=$page_js4editor_f($page_editor_pair['init'],
                                      $path, $editing_json_part,
                                      $name_mapping, $editor_mapping);
    foreach($page_css4editor as $css_src){?>
        <link rel="stylesheet" type="text/css" href="<?=$css_src;?>">
    <?php }
    foreach($page_js4editor as $js_src){?>
        <script src="<?=$js_src;?>"></script>
    <?php }
}?>

<script type="text/javascript">

$(document).ready(function(){
    $('body > .editor_operation > .done_editing').click(function(){
        $('<form>', {
            'action': '<?=$_SERVER['PHP_SELF'];?>',
            'method': 'post'
        }).append($('<input>', {
            'type': 'hidden',
            'name': 'done',
            'value': 'true'
        })).appendTo(document.body).submit();
    });
});

</script>

</head>
<body>
<div class="editor_title">
    <?php {
        $path_recombine="";
        $path_key="";
        $path_name=map_path_id($name_mapping, $path_recombine);
        ?><?=$path_name
            ?'<a href="?path='.$path_recombine.'">'.$path_name.'</a>'
            :$path_key;?><?php
        for($i=1; $i<count($path_keys); $i++){
            $path_key=$path_keys[$i];
            $path_recombine.='.'.$path_key;
            $path_name=map_path_id($name_mapping, $path_recombine);
            ?><?='.'.($path_name
                ?'<a href="?path='.$path_recombine.'">'.$path_name.'</a>'
                :$path_key);?><?php
        }
    }?>
</div>
<div class="editor_content">
    <?php {
        $page_html4editor_f='json_'.$page_editor_pair['editor'].'_editor_html';
        ?><?=$page_html4editor_f($page_editor_pair['init'],
                                 $path, $editing_json_part,
                                 $name_mapping, $editor_mapping);?><?php
    }?>
</div>
<div class="editor_operation">
    <button class="done_editing">Done</button>
</div>
</body>
</html>
<?php


}/*===========================================================================*/

?>
