<?php

require_once dirname(__FILE__).'/../lib/global.php';
require_once dirname(__FILE__).'/../lang_rawstr.php';

// 拉取
if(isset($_GET['id'])){
    $id=$_GET['id'];
    $sql="select `info`, `info_vars` from `skill` where `id`={$id};";
    $rows=db_query($db, $sql);
	$result_json=array();
    if (count($rows)>0){
		$template=$rows[0]['info'];
		$variables=$rows[0]['info_vars'];
		if(!isset($variables) || strlen($variables)<=0){
			$variables="{}";
		}
		$result_json["template"]=$template;
		$result_json["variables"]=json_decode($variables, true);
    }else{
		$result_json["template"]="";
		$result_json["variables"]=array();
	}
	echo json_encode($result_json, JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE);
}

// 提交
if(isset($_POST['id'])){
	$rawsql_parser=new lang_rawstr('mysql');
    $id=$_POST['id'];
	$content=json_decode($_POST['content'], true);
    $variables=$content['variables'];
	$variables_sql=$rawsql_parser->replace(json_encode($variables, JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE));
    $template=$content['template'];
	$template_sql=$rawsql_parser->replace($template);
    
    $sql="update `skill` set `info`='{$template_sql}', `info_vars`='{$variables_sql}' where `id`={$id};";
    db_execute($db, $sql);
}
