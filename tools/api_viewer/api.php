
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
 <head>
  <title>仙侠道移动版辅助脚本</title>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /> 
 </head>
 <body>

<?php

if (!empty($_GET['sf']) && !empty($_GET['sr'])) {
	echo file_get_contents("http://debug.vvv.io:8081/unpack_file.php?filename=".$_GET['sf']."&a=".$_GET['sr']);
	exit;
} 

ini_set("memory_limit", "-1");
set_time_limit(0);

$proto_path = "../../protocol";
$cur_path = dirname(__FILE__);


require_once($proto_path."/protodoc.php");
require_once("function.php");

$modules = parse_all('game_server');

if (isset($_GET['action']) && ($_GET['action'] == "cc")) {

    var_dump(shell_exec("cd {$cur_path} && sh build.sh"));
    echo "<a href=\"genclient/data.zip\">Download(data.zip)<a/>";
    echo "<br>";
    echo "<br>";
    echo "<a href=\"genclient/data.tgz\">Download(data.tgz)<a/>";
    exit;
}

require_once("status.php");
$status = GetServerStatus();
/*======================================================================================================*/

echo "<h1>仙侠道api速查手册</h1>{$status}";
$modules_main_table = array();
$modules_main_table["key"] = array("模块id", "模块名称", "模块说明");
$modules_main_table["title"] = array("模块id： ", "模块名称： ", "模块说明： ");

foreach ($modules as $k => $v) {
    $v = (array)$v;
    
    $tableline = array();
    $tableline["模块id"] = $v["ID"];
    $tableline["模块名称"] = array('key' => "a", "href" => "#" . $v["Name"], "text" => $v["Name"],);
    $tableline["模块说明"] = array('key' => "a", "href" => "#" . $v["Name"], "text" => isset($v["Comments"][0])?$v["Comments"][0]:'',);
    $modules_main_table["table"][] = $tableline;
}
$outtable[] = html_creat_table($modules_main_table);

foreach ($modules as $k => $v) {
    $v = (array)$v;
    $outtable[] = html_new_node("hr", array());
    $module_tile_line = html_new_node('h2', array());
    $module_tile = html_new_node("a", array('name' => $v["Name"]));
    html_add_text($module_tile, isset($v["Comments"][0])?$v["Comments"][0]:'');
    html_add_text($module_tile, "(" . $v["ID"] . ")");
    html_add_node($module_tile_line, $module_tile);
    $outtable[] = $module_tile_line;
    $module_table = html_new_node("table", array("cellspacing" => 0, "cellpadding" => 5));
    $tr = html_new_node("tr");
    
    html_add_new_node($tr, "td", array(), "action_id");
    html_add_new_node($tr, "td", array("colspan" => 3), "借口名称");
    html_add_new_node($tr, "td", array("colspan" => 2), "接口说明");
    html_add_new_node($tr, "td", array(), "格式");
    html_add_node($module_table, $tr);
    foreach ($v["Actions"] as $k1 => $v1) {
        $desid = $v['Name'] . "_" . $v1->Name;
        $tr = html_new_node("tr");
        html_add_new_node($tr, "td", array(), $v1->ID);
        html_add_new_node($tr, "td", array("colspan" => 3), $v1->Name);
        $i = 0;
        $comments = "";
        while ($i < sizeof($v1->Comments)) {
            if (trim($v1->Comments[$i]) != "") {
                $comments = $v1->Comments[$i];
                break;
            }
            $i++;
        }
        html_add_new_node($tr, "td", array("colspan" => 2), $comments);
        $td = html_new_node("td", array("onclick" => "doc = document.getElementById(\"" . $desid . "\");if(doc.style.display==\"none\"){ doc.style.display=\"table-row\";}else {doc.style.display=\"none\";}"));
        html_add_new_node($td, "a", array("href" => "javascript:;"), "查看");
        html_add_node($tr, $td);
        html_add_node($module_table, $tr);
        $tr = html_new_node("tr", array("style" => "display:none;background-color:#ccc;", "id" => $desid));
        $td = html_new_node("td", array("colspan" => 3));
        $intable = html_new_node("table", array("cellspacing" => 0, "cellpadding" => 2, "width" => "100%"));
        $intr = html_new_node("tr");
        $intd = html_new_node("td", array("colspan" => 3), "输入\n");
        html_add_node($intr, $intd);
        html_add_node($intable, $intr);
        $intr = html_new_node("tr");
        
        html_add_node($intable, $intr);
        
        $intr = html_new_node("tr");
        $intd = html_new_node("td", array());
        $alltxt = "";
        foreach ((array)$v1->InColumns as $k2 => $v2) {
            $alltxt.= build_item_table($v2, $intable);
        }
        html_add_new_node($intd, "pre", array(), $alltxt);
        html_add_node($intr, $intd);
        html_add_node($intable, $intr);
        html_add_node($td, $intable);
        html_add_node($tr, $td);
        $td = html_new_node("td", array("width" => "10%"));
        html_add_node($tr, $td);
        $td = html_new_node("td", array("colspan" => 4));
        $intable = html_new_node("table", array("cellspacing" => 0, "cellpadding" => 2, "width" => "100%"));
        $intr = html_new_node("tr");
        $intd = html_new_node("td", array("colspan" => 3), "输出\n");
        html_add_node($intr, $intd);
        html_add_node($intable, $intr);
        $intr = html_new_node("tr");
        $intd = html_new_node("td", array());
        $alltxt = "";
        if (is_array($v1->OutColumns)) {
            foreach ((array)$v1->OutColumns as $k2 => $v2) {
                $alltxt.= build_item_table($v2, $intable);
            }
        }
        html_add_new_node($intd, "pre", array(), $alltxt);
        html_add_node($intr, $intd);
        html_add_node($intable, $intr);
        html_add_node($td, $intable);
        html_add_node($tr, $td);
        html_add_node($module_table, $tr);
    }
    
    $outtable[] = $module_table;
}

html_creat($outtable);

function parseitem($v) {
    $temp = array();
    $temp["name"] = $v->Name;
    $temp["type"] = substr($v->Type, 2);
    if ($v->Type == "t_list8" || $v->Type == "t_list16") {
        $temp["fields"] = array();
        foreach ($v->Items as $v1) {
            $temp["fields"][] = parseitem($v1);
        }
    }
    return $temp;
}
function build_item_table($v2, $intable, $pre = "") {
    
    $alltext = $pre . $v2->Name . " ";
    // echo "<div><pre>";
    // print_r($v2);
    // echo "</div></pre>";
    $cmtstr = "";
    foreach ($v2->Comments as $v3) {
        $cmtstr.= $v3;
    }
    if (isset($v2->Type)) {
        $type_str = substr($v2->Type, 2);
        if ($v2->Type == "t_list8" || $v2->Type == "t_list16") {
            $cmtstr.= substr($v2->Type, 6) . "位长度的数组\n";
        }
        if ($v2->Type == "t_enum8" || $v2->Type == "t_enum16") {
            $type_str.= "-" . $v2->RefType;
        }

    } else {
        $type_str = "-" . $v2->Value;
    }

    $add_cmt_str = '';
    if (isset($v2->Type) && ($v2->Type == 't_struct')) {
        $add_cmt_str = '结构体类型';
    }
    $alltext.= $type_str . " ($add_cmt_str" . trim($cmtstr, "\n") . ")\n";
    
    if (isset($v2->Type) && ($v2->Type == "t_list8" || $v2->Type == "t_list16" || $v2->Type == "t_enum8" || $v2->Type == "t_enum16" || $v2->Type == 't_struct')) {
        $alltext.= $pre . "{\n";
        foreach ($v2->Items as $v3) {
            $alltext.= build_item_table($v3, $intable, $pre . "  ");
        }
        $alltext.= $pre . "}\n";

    }
    return $alltext;
}


