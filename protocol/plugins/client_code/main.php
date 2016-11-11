<?php
//===============================================
// 客户端通讯协议代码生成器
//===============================================

function plugin_main()
{
	global $argv;

	$modules = parse_all($argv[2]);

	if (!isset($argv[4]) || empty($argv[4])) {
		exit("请指定客户端协议代码生成路径");
	}

	$out_dir = $argv[4];

	$jscode_path = $out_dir."/js_protocol";
	$jscode_up_path = $jscode_path."/up";
	$jscode_down_path = $jscode_path."/down";
    $jscode_json_path = $jscode_path."/json";

   if (!file_exists($jscode_path))
        mkdir($jscode_path);

    if (!file_exists($jscode_up_path))
        mkdir($jscode_up_path);

    if (!file_exists($jscode_down_path))
        mkdir($jscode_down_path);

    if (!file_exists($jscode_json_path))
        mkdir($jscode_json_path);

    generateProtocolJSON($modules, $jscode_json_path);
    
    $modulejs = '
/**
 * 本文件由脚本生成
 * 生成地址访问: http://svn.vvv.io:8765
 * PIModule文件均为脚本生成，可以修改down文件夹内容
 * 所有PIModule均包含在PIModule对象中，使用方法如下：
 *     PIModule.Player.sendLogin(xxx, xxx, ...);
 *     PIModule.{$module_name}.send{$action_name}({$param, ...})
 * 协议包含的enum使用方法如下：
 *     MOD_PLAYER.LOGIN_RESULT_SUCCEED
 *     MOD_{$module_name}.{$action_name}_{$enum_name}
 */
var PIModule = {};

';

	foreach ($modules as $k => $v) {
        $modulejs.= 'include("module/up/PIModule_' . strStartUper($v->Name) . '_Up.js");' . "\n";
        $modulejs.= 'include("module/down/PIModule_' . strStartUper($v->Name) . '_Down.js");' . "\n";
    }

    save_code($jscode_path."/PIModule.js", $modulejs);
    remove_old_files($jscode_path);
    
    //生成up和down相关的js文件
    foreach ($modules as $k => $v) {
        $module_name = $v->Name;
        $module_id = $v->ID;
        $module_str = '
// enum
var MOD_' . strtoupper($module_name) . ' = {
		    
		';

        foreach ($v->Types as $k1 => $v1) {
            $type_name = $v1->Name;
            if ($v1->Kind == 't_struct') {
                break;
            }
            foreach ($v1->Fields as $k2 => $v2) {
                $module_str.= "  " . strtoupper($type_name) . "_" . strtoupper($v2->Name) . ": " . $v2->Value . ",\n";
            }
            $module_str.= "\n";
        }
        $module_str.= "}\n\n";
        $module_str.= "PIModule." . strStartUper($module_name) . " = PIModule." . strStartUper($module_name) . " || {};\n\n";
        $module_out_str = '';
        
        foreach ($v->Actions as $k1 => $v1) {
            
            $action_name = $v1->Name;
            $action_id = $v1->ID;

            if (is_array($v1->InColumns)) {
                $action_up_name = array();
                foreach ($v1->InColumns as $k2 => $v2) {
                    $action_up_name[] = $v2->Name;
                }

                $paramCount = count($action_up_name);
                $sockTagCode = "((arguments[{$paramCount}]===undefined) ? 'def' : arguments[{$paramCount}])";

                $module_str.= 'PIModule.' . strStartUper($module_name) . '.send' . strStartUper($action_name) . ' = function(' . implode(", ", $action_up_name) . ') {' . "\n";
                $module_str.= '  PIProtocol.getInstance().send_packet( '.$sockTagCode.", " . $module_id . ',' . $action_id . ', {' . "\n";
                foreach ($action_up_name as $k2 => $v2) {
                    $module_str.= '    "' . $v2 . '":' . $v2 . ',' . "\n";
                }
                $module_str.= '  });' . "\n";
                $module_str.= '};' . "\n\n";
            }
            if (is_array($v1->OutColumns)) {
                $module_out_str.= 'PIModule.' . strStartUper($module_name) . '.proc' . strStartUper($action_name) . ' = function(dict) {' . "\n\n";
                $module_out_str.= '};' . "\n\n";
            }
        }
        save_code($jscode_up_path."/PIModule_" . strStartUper($module_name) . "_Up.js", $module_str);
        save_code($jscode_down_path."/PIModule_" . strStartUper($module_name) . "_Down.js", $module_out_str);
        remove_old_files($jscode_up_path);
        remove_old_files($jscode_down_path);
    }
}


function strStartUper($str) {
    return str_replace(array('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'), array('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'), substr($str, 0, 1)) . substr($str, 1);
}

function parseitem($v){
    $temp=array();
    $temp["name"]=$v->Name;
    $temp["type"]=substr($v->Type,2);
    if($v->Type=="t_list8"||$v->Type=="t_list16"||$v->Type=='t_struct'){
        $temp["fields"]=array();
        foreach($v->Items as $v1){
            $temp["fields"][]=parseitem($v1);
        }
    }
    return $temp;
}

function generateProtocolJSON(&$modules, $path) {
    $modules_main=array();
    foreach($modules as $k=>$v) {
        $module_id=$v->ID;
        foreach($v->Actions as $k1=>$v1){
            $protocolup=array();
            $protocoldown=array();
            $action_id=$v1->ID;

            $protocolup["action_name"]=$v1->Name;
            $protocolup["module_name"]=$v->Name;
            $protocolup["protocol"]=array();
            $protocoldown["action_name"]=$v1->Name;
            $protocoldown["module_name"]=$v->Name;
            $protocoldown["protocol"]=array();
            foreach((array)$v1->InColumns as $v2){
                $protocolup["protocol"][]=parseitem($v2);
            }
            foreach((array)$v1->OutColumns as $v2){
                $protocoldown["protocol"][]=parseitem($v2);
            }
            $modules_main[$module_id][$action_id]["up"]=$protocolup;
            $modules_main[$module_id][$action_id]["down"]=$protocoldown;

            remove_old_files($path);
            save_code("{$path}/protocol_up_".$module_id."_".$action_id.".json",json_encode($protocolup,JSON_NUMERIC_CHECK));
            save_code("{$path}/protocol_down_".$module_id."_".$action_id.".json",json_encode($protocoldown,JSON_NUMERIC_CHECK));
        }
    }
}

?>