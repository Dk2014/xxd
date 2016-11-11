<?php


require_once dirname(__FILE__).'/../lib/config.php';
require_once dirname(__FILE__).'/../lib/mysql.php';


if (isset($_GET['p'])) {
	print('<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />');
	print("<div><pre></br></br></br></br>");
	print(export_table_struct_to_pconfig($_GET['p']));
	print("</div></pre>");
}

function export_table_struct_to_pconfig($table_name) {
	global $db_config;
	$db = db_connect();


	$db_name = $db_config['name'];
	$columns_rs = db_query($db, "select COLUMN_NAME,COLUMN_COMMENT from information_schema.columns where `TABLE_SCHEMA`='{$db_name}' and table_name='{$table_name}'");
	
	if (count($columns_rs) == 0) {
		print("{$table_name} is not exists\r\n");
		return;
	}

	$column_code = "";
	foreach ($columns_rs as $value) {
		$column_name = $value['COLUMN_NAME'];
		$comment = $value['COLUMN_COMMENT'];
		$column_code .= "		array('name' => '{$column_name}', 'text' => '{$comment}', 'width' => '80px'),\n";
	}

	$code = <<<PCONFIG

\$pconfig = array(
	'title'   => '',
	'table'   => '{$table_name}',
	'links'   => array(),
	
	'columns' => array(
{$column_code}
	),
	
	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => ''),
	),
	

	'where' 		=> 'sql_where',
	'insert'  		=> 'sql_insert',
	
	'not_delete' 	=> true,
	'not_new' 		=> true,
	'not_edit' 		=> true,
);
PCONFIG;

	return $code;
}

?>