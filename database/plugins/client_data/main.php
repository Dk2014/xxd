<?php
require_once(dirname(__FILE__)."/../../lang_rawstr.php");

define("DATABASE_URL", dirname(dirname(dirname(__FILE__)))."/");
define("CONFIG_CONST_URL", dirname(dirname(dirname(dirname(__FILE__))))."/tools/api_viewer/genclient/js_const/");
define("EXPORT_DBDATA_URL", dirname(dirname(dirname(dirname(__FILE__))))."/tools/api_viewer/genclient/js_dbdata/");

$is_win = PHP_OS == "WINNT";

if (!is_dir(CONFIG_CONST_URL)) {
	mkdir(CONFIG_CONST_URL, 0777, true);
}

if (!is_dir(EXPORT_DBDATA_URL)) {
	mkdir(EXPORT_DBDATA_URL);
}

$json_escaper = new lang_rawstr("js");

/**
 * 客户端数据常量代码生成器
 */
function plugin_main () {
	$path = dirname(__FILE__)."/";
	
	global $as_db;
	$as_db = db_connect();
	
	product_config($path);
	export_fromdb($path);
}

function product_config ($path) {
	require_once($path."config.php");
	remove_all(CONFIG_CONST_URL."*.js");
	config_main($path);
}

function export_fromdb($path) {
	require_once($path."dbdata.php");
	remove_all(EXPORT_DBDATA_URL."*.js");
	dbdata_main($path);
}

function remove_all($rule) {
	$list = glob($rule);
	$len = count($list);
	for ($i = 0; $i < $len; $i++) {
		unlink($list[$i]);
	}
}

$cols_to_table = array();

/**
 * 获得列信息与纪录集
 * @param $table 表名
 * @param $fields_name 字段名列表
 */
function get_cols_and_records ($table, $fields_name, $addon_sql = "") {
	global $cols_to_table;
	global $as_db, $as_table;
	
	$as_table = $table;
	
	$cols = get_table_info($table, $fields_name);
	$records = db_query($as_db, format_fields($table, $fields_name, $addon_sql));
	
	$cols_to_table[join_arr($cols)] = $table;
	
	return array($cols, $records);
}

function join_arr (&$cols) {
	$str = "";
	
	for ($i = 0; $i < count($cols); $i++) {
		if ($str != "") $str .= ",";
		$str .= join(",", $cols[$i]);
	}
	
	return sha1($str);
}

/**
 * 格式化字段列表为查询语句
 * @param $table 表名
 * @param $fields_name 查询的字段列表
 */
function format_fields ($table, $fields_name, $addon_sql = "") {
	if (count($fields_name) > 0) {
		return " select `".join("`, `", $fields_name)."` from $table ".$addon_sql;
	}
	
	return "select * from $table ".$addon_sql;
}

/**
 * 获得表字段描述
 * @param $table 表名
 * @param $fields 提取的字段的列表
 */
function get_table_info ($table, $fields) {
	static $tables;
	
	if (empty($tables)) $tables = db_tables();
	
	$cols = array();
	
	$len = count($tables);
	for ($i = 0; $i < $len; $i++) {
		if ($tables[$i]["name"] == $table) {
			$temp = $tables[$i]["cols"];
			
			for ($j = 0; $j < count($temp); $j++) {
				if (count($fields) == 0 || in_array($temp[$j]["name"], $fields)) {
					$cols[] = array($temp[$j]["name"], $temp[$j]["type"], $temp[$j]["desc"]);
				}
			}
			
			break;
		}
	}
	
	return $cols;
}

/**
 * 格式化字段及字段说明为注释
 * @param $cols get_table_info()返回的结果
 */
function get_field_comment (&$cols) {
	$names = array();
	
	$len = count($cols);
	for ($i = 0; $i < $len; $i++) {
		$names[] = "	 * ".$i." : ".$cols[$i][0].", ".$cols[$i][1].", ".$cols[$i][2];
	}
	
	return "	/**
".join("\n", $names)." 
	 */";
}

/**
 * Summary
 * @param	object	$fields	Description
 * @param	object	$cols	Description
 */
function get_field_values (&$fields, &$cols) {
	$arr = array();
	for ($j = 0; $j < count($cols); $j++) {
		switch ($cols[$j][1]) {
			case "smallint":
			case "tinyint":
			case "int":
			case "bigint":
			case "float": {
				$v = $fields[$cols[$j][0]];
				$arr[] = ($v ? $v : 0)." /*[".$j."]*/";
				break;
			}
			case "varchar": {
				global $json_escaper;
				$value = $fields[$cols[$j][0]];
				$value = $json_escaper->replace($value);
				
				$arr[] = $value." /*[".$j."]*/";
				break;
			}
			case "text": {
				global $json_escaper;
				$value = $fields[$cols[$j][0]];
				$value = $json_escaper->replace($value);
				$arr[] = $value." /*[".$j."]*/";
				//$arr[] = "\"".$value."\""." /*[".$j."]*/";
				break;
			}
		}
	}
	
	return $arr;
}

/**
 * 获得列索引
 * @param $field_name 列名
 * @param $cols 列名列表
 */
function get_field_index ($field_name, &$cols) {
	global $is_win;
	
	$len = count($cols);
	for ($i = 0; $i < $len; $i++) {
		if ($cols[$i][0] == $field_name)
			return $i." /* $field_name : ".$cols[$i][1].", ".$cols[$i][2]." */";
	}
	
	$table = $cols_to_table[join_arr($cols)];
	
	$error = "table.".$table." 找不到字段 $as_table.$field_name。";
	if ($is_win) $error = iconv("utf-8", "gbk", $error);
	
	throw new Exception($error);
	exit(1);
}

/**
 * 输出as代码——获取指定id的纪录
 * @param $records 纪录集
 * @param $cols 列名列表
 */
function get_list_item (&$records, &$cols) {
	global $cols_to_table;
	
	$table = $cols_to_table[join_arr($cols)];
	
	return 'private static var _itemCache : Object = {};
	
	/**
	 * 获取指定id的记录
	 */
	private static function getListItem (id : int) : Array
	{
		if (_itemCache[id]) return _itemCache[id];
		
		for (var i : int = 0; i < '.count($records).'; i++)
		{
			if (_list[i]['.get_field_index("id", $cols).'] == id)
				return (_itemCache[id] = _list[i]);
		}
		
		throw new Error("table.'.$table.' 查询不到对应的记录：id=" + id);
		return [];
	}';
}
?>
