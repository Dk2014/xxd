<?php

require_once dirname(__FILE__).'/zd_cgo.php';
require_once dirname(__FILE__).'/zd_hooks.php';
require_once dirname(__FILE__).'/zd_load.php';
require_once dirname(__FILE__).'/zd_prepare.php';
require_once dirname(__FILE__).'/zd_rows.php';
require_once dirname(__FILE__).'/zd_sync.php';
require_once dirname(__FILE__).'/zd_tables.php';
require_once dirname(__FILE__).'/zd_trans.php';
require_once dirname(__FILE__).'/zd_trans_const.php';

//===============================================
// 服务端数据库映射代码生成器
//===============================================
function plugin_main ()
{
	global $db_config, $target;

	$output_dir = 
	/*mobile-dev*/dirname(
		/*database*/dirname(
			/*plugins*/dirname(
				/*server_code*/dirname(__FILE__)
			)
		)
	).
	"/server/src/".$db_config[$target]['server']."/mdb/";

	$db = db_connect();

	// 获取数据库结构
	$tables = db_tables(true);
	$tables_count = count($tables);
	$log_tables = db_tables(false, true);

	// 对所有表和字段做ID分配
	for ($i = 0; $i < count($tables); $i ++) {
		$tables[$i]['id'] = $i;
		for ($j = 0; $j < count($tables[$i]['cols']); $j ++ ){
			$tables[$i]['cols'][$j]['id'] = $j;
		}
	}

	gen_cgo_h($output_dir, $tables);
	gen_cgo_c($output_dir, $tables);
	gen_load_go($output_dir, $tables);
	gen_prepare_go($output_dir, $tables, $log_tables);
	gen_hooks_go($output_dir, $tables);
	gen_rows_go($output_dir, $tables);
	gen_sync_go($output_dir, $tables);
	gen_tables_go($output_dir, $tables);
	gen_trans_go($output_dir, $tables);
	GenTransConst($output_dir, $tables);
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

// 获取变量名（首字母小写）
function get_var_name($name) {
	$x = get_type_name($name);
	$x[0] = strtolower($x[0]);
	return $x;
}

function is_hashtable($table) {
	return strpos($table['name'], 'global') === 0 || (strpos($table['name'], 'player') === 0 && $table['pri']['name'] != 'pid' && $table['name'] != 'player');
}

function is_one_to_one_player_table($table) {
	return (strpos($table['name'], 'player') === 0 && $table['pri']['name'] == 'pid') || $table['name'] == 'player';
}

// mysql类型映射成sync文件头信息字段类型标识
function get_sync_head_tag($type) {
	switch ($type) {
		case 'tinyint':
			return 1;
		case 'smallint':
			return 2;
		case 'int':		
			return 3;
		case 'bigint':
			return 4;
		case 'text':
			return 5;
		case 'char':
			return 6;
		case 'varchar':
			return 7;
		case 'blob':
			return 8;
		case 'binary':
			return 9;
		case 'float':
			return 10;
		case 'double':
			return 11;
		case 'timestamp':
			return 12;
		case 'mediumblob':
			return 13;
		default:
			return "error";
	}
}

// mysql类型映射成c类型
function get_c_type($type) {
	switch ($type) {
		case 'tinyint':
			return 'int8_t';
		case 'smallint':
			return 'int16_t';
		case 'int':		
			return 'int32_t';
		case 'bigint':
			return 'int64_t';
		case 'text':
		case 'char':
		case 'varchar':
		case 'blob':
		case 'mediumblob':
		case 'binary':
			return 'char*';
		case 'float':
			return 'float';
		case 'double':
			return 'double';
		case 'timestamp':
			return 'int64_t';
	}
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
		case 'blob':
		case 'mediumblob':
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

// mysql类型映射成go类型（函数名）
function get_go_type2($type) {
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
		case 'blob':
		case 'mediumblob':
		case 'binary':
			return 'Bin';
		case 'float':
			return 'Float32';
		case 'double':
			return 'Float64';
		case 'timestamp':
			return 'LocaltimeToInt64';
	}
}

// 根据类型获取mysql绑定代码
function get_bind_type($type) {
	switch ($type) {
		case 'tinyint':
			return "BindTinyInt";
		case 'smallint':
			return "BindSmallInt";
		case 'int':		
			return "BindInt";
		case 'timestamp':
		case 'bigint':
			return "BindBigInt";
		case 'text':
		case 'char':
		case 'varchar':
			return "BindVarchar";
		case 'blob':
		case 'mediumblob':
		case 'binary':
			return "BindBinary";
		case 'float':
			return "BindFloat";
		case 'double':
			return "BindDouble";
	}
}

// 根据mysql类型获取写入函数
function get_sync_write_type($type) {
	switch ($type) {
		case 'tinyint':
			return "WriteInt8";
		case 'smallint':
			return "WriteInt16";
		case 'int':		
			return "WriteInt32";
		case 'timestamp':
		case 'bigint':
			return "WriteInt64";
		case 'text':
		case 'char':
		case 'varchar':
			return "WriteString";
		case 'blob':
		case 'mediumblob':
		case 'binary':
			return "WriteBytes";
		case 'float':
			return "WriteFloat32";
		case 'double':
			return "WriteFloat64";
	}
}

// 根据mysql类型获取读取函数
function get_unsafe_read_type($type) {
	switch ($type) {
		case 'tinyint':
			return "TinyInt(p.ReadInt8())";
		case 'smallint':
			return "SmallInt(p.ReadInt16())";
		case 'int':
			return "Int(p.ReadInt())";
		case 'bigint':
			return "BigInt(p.ReadInt64())";
		case 'text':
		case 'char':
		case 'varchar':
			return "Varchar(p.ReadByte16())";
		case 'float':
			return "Float(p.ReadFloat32())";
		case 'double':
			return "Double(p.ReadFloat64())";
	}
}

?>