<?php

/*

	备注：

	tpl_dat.php 用于从mysql模板数据结构导出成go struct的工具。
	这便于快速生成go struct，也统一了这类go struct的声明方式，因为在定义模板数据字段的时候会附上字段相应的注释

	注意：模板数据会根据业务逻辑的需要通常会把数据组织成一种高效的结构方便业务逻辑调用，
		所以使用这个工具生成的数据结构不一定是最终形式，需要你根据实际情况自行修改。


	usage:

	http://www../tpl_dat.php?p=table_name[, table_name]

 */

require_once dirname(__FILE__).'/../lib/config.php';
require_once dirname(__FILE__).'/../lib/mysql.php';

$note = <<<NOTE

注意：模板数据会根据业务逻辑的需要通常会把数据组织成一种高效的结构方便业务逻辑调用，
	 所以使用这个工具生成的数据结构不一定是最终形式，需要你根据实际情况自行修改。

NOTE;

if (isset($_GET['p'])) {

	print('<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />');
	print("<div><pre></br>");

	$db = db_connect();
	$table_list = explode(',', $_GET['p']);

	$load_fn_list = array();
	$map_var_list = array();

	$code = "";	
	foreach ($table_list as $table_name) {
		$code .= export_table_struct_to_game_dat($db, $table_name);
	}

	$load_fn_code = implode("\r\n", $load_fn_list);
	$map_var_code = implode("\r\n", $map_var_list);

	$load_code = <<<LC

package {$table_name}_dat

import (
	"core/mysql"
)


var (
	{$map_var_code}
)


func Load(db *mysql.Connection) {
	{$load_fn_code}
}

LC;

	$code = $load_code."\n\n".$code;
	print($note."</br></br></br></br>");
	print($code);
	print("</br></br>// ############### 对外接口实现 coding here ###############</br></br>");
	print("</div></pre>");
}


function export_table_struct_to_game_dat(&$db, $table_name) {
	global $db_config;
	$db_name = $db_config['name'];
	$columns_rs = db_query($db, "select COLUMN_KEY, COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE from information_schema.columns where `TABLE_SCHEMA`='{$db_name}' and table_name='{$table_name}'");
	
	if (count($columns_rs) == 0) {
		print("{$table_name} is not exists\r\n");
		return;
	}

	$go_struct_name = get_type_name($table_name);
	$code = "type {$go_struct_name} struct {\n";
	foreach ($columns_rs as $value) {
		$column_name = $value['COLUMN_NAME'];
		$comment = $value['COLUMN_COMMENT'];
		$column_type = $value['DATA_TYPE'];
		
		$code .= "	".get_type_name($column_name)." ".get_go_type($column_type)." // ".$comment."\n";
	}

	$code .= "}\n\n";

	$code .= export_go_struct_data_loadfunction($columns_rs, $table_name, $go_struct_name);

	return $code;
}


function export_go_struct_data_loadfunction(&$columns_rs, $table_name, $struct_name)
{
	global $load_fn_list, $map_var_list;

	$fn_name = get_type_name($table_name);
	list($pri_name, $pri_type) = get_pri_column_name_and_type($columns_rs);
	$idx_code = get_column_visit_idx_code($columns_rs);

	$map_var_name = "map{$fn_name} = map[".get_go_type($pri_type)."]*{$struct_name}{}";
	$defined_key_code = "var pri_{$pri_name} ".get_go_type($pri_type);
	$key_code = "pri_{$pri_name} = row.".get_go_row_type($pri_type)."(i".get_type_name($pri_name).")";
	$column_code = get_column_value_code($columns_rs, $pri_name);

	$load_fn_list[] = "load{$fn_name}(db)";
	$map_var_list[] = "map{$fn_name} map[".get_go_type($pri_type)."]*{$struct_name}";

 	$code = <<<CODE

func load{$fn_name}(db *mysql.Connection) {	
	res, err := db.ExecuteFetch([]byte("SELECT * FROM {$table_name} ORDER BY `{$pri_name}` ASC"), -1)
	if err != nil {
		panic(err)
	}

{$idx_code}

	{$defined_key_code}
	{$map_var_name}
	for _, row := range res.Rows {
		{$key_code}
		map{$fn_name}[pri_{$pri_name}] = &{$struct_name}{
{$column_code}
		}
	}
}
CODE;

	return $code;
}

function get_pri_column_name_and_type(&$columns_rs) {
	foreach ($columns_rs as $value) {
		if (!empty($value['COLUMN_KEY']))
			return array($value['COLUMN_NAME'], $value['DATA_TYPE']);
	}

	exit("没有发现表的主键字段");
}

function get_column_visit_idx_code(&$columns_rs) {
	$code = "";
	foreach ($columns_rs as $value) {
		$column_name = $value['COLUMN_NAME'];
		$comment = $value['COLUMN_COMMENT'];
		$column_type = $value['DATA_TYPE'];

		$code .= "	i".get_type_name($column_name)." := res.Map(\"{$column_name}\")\n";
	}
	return $code;
}

function get_column_value_code(&$columns_rs, $pri_name) {
	$code = "";
	foreach ($columns_rs as $value) {
		$column_name = $value['COLUMN_NAME'];
		$comment = $value['COLUMN_COMMENT'];
		$column_type = $value['DATA_TYPE'];
		if (empty($value['COLUMN_KEY'])) {
			$code .= "			".get_type_name($column_name)." : row.".get_go_row_type($column_type)."(i".get_type_name($column_name)."),\n";
		} else {
			$code .= "			".get_type_name($column_name)." : pri_{$pri_name},\n";
		}
	}
	return $code;
}

// 获取类型名（下划线分割转驼峰式）
function get_type_name($name) {
	$items = explode('_', $name);
	$type_name = '';
	for ($i = 0; $i < count($items); $i ++) {
		$type_name .= ucfirst($items[$i]);
	}
	return $type_name;
}

// mysql类型映射成go类型
function get_go_type($type) {
	switch ($type) {
		case 'tinyint':
			return 'int8';
		case 'smallint':
			return 'int16';
		case 'int':		
			return 'int32';
		case 'bigint':
			return 'int64';
		case 'text':
		case 'char':
		case 'varchar':
			return 'string';
		case 'binary':
			return '[]byte';
		case 'float':
			return 'float32';
		case 'double':
			return 'float64';
		case 'timestamp':
			return 'int64';
	}
}

// mysql类型映射成go mysql row类型
function get_go_row_type($type) {
	switch ($type) {
		case 'tinyint':
			return 'Int8';
		case 'smallint':
			return 'Int16';
		case 'int':		
			return 'Int32';
		case 'bigint':
			return 'Int64';
		case 'text':
		case 'char':
		case 'varchar':
			return 'Str';
		case 'binary':
			return 'Bin';
		case 'float':
			return 'Float32';
		case 'double':
			return 'Float64';
		case 'timestamp':
			return 'Int64';
	}
}


?>