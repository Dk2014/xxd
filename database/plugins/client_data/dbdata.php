<?php


function dbdata_main($path) {
	$list = glob($path."/export_tables/*.php");
	$len = count($list);
	for ($i = 0; $i < $len; $i++) {
		require_once($list[$i]);
	}
}


/*
	Usage:

	$export_table = array(
		'table' 	=> 'town',	// 配置表名
		'export_order' => ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'lock',
			'blablabla',
		),
	);

	export_table($export_table);
	
*/
function export_table($export_table) {
	global $as_db, $target, $db_config;

	$db_name = $db_config[$target]['name'];
	$columns_rs = db_query($as_db, "select COLUMN_NAME from information_schema.columns where `TABLE_SCHEMA`='{$db_name}' and table_name='{$export_table['table']}'");
	$columns = array();
	foreach ($columns_rs as $value) {
		$columns[] = $value['COLUMN_NAME'];
	}

	if (count($columns) == 0) {
		print("[error] The table {$export_table['table']} is not exists\r\n");
		return;
	}
	
	$export_columns = array_diff($columns, $export_table['exclude_columns']);
	if (count($columns) != count($export_columns) + count($export_table['exclude_columns'])) {
		print("[error] check {$export_table['table']} exclude_column error\r\n");
		return;
	}

	list($cols, $records) = get_cols_and_records($export_table['table'], $export_columns, $export_table['export_order']);

	$list = array();
	$c = count($records);

	if (isset($export_table['addition'])) {
		$add = array();
		while ($key = key($export_table['addition'])) {
			$cbfunc = "get_addition_{$key}";
			$add[] = $cbfunc($as_db);
			next($export_table['addition']);
		}

		for ($i = 0; $i < $c; $i++) {
			$rs = $records[$i];
			$data = get_field_values($rs, $cols);
			for ($j = 0; $j < count($add); $j++) {
				$data[] =  strval($add[$j][$i])." /*[".strval(count($cols)+$j)."]*/";
			}
			$list[] = "[".join(", ", $data)."]";
		}
	} else {
		for ($i = 0; $i < $c; $i++) {
			$rs = $records[$i];
			$list[] = "[".join(", ", get_field_values($rs, $cols))."]";
		}
	}

	$comment = get_field_comment($cols);

	$data = join(",\n		", $list);

	if (isset($export_table['addition'])) {
		$comment2 = substr($comment, 0,strlen($comment)-3);
		reset($export_table['addition']);
		while($key = key($export_table['addition'])) {
			$current = current($export_table['addition']);
			$cols[count($cols)][0] = $key;
			$comment2 .= " * ".strval(count($cols)-1)." : ".$key.", ".$current."\n";
			next($export_table['addition']);
		}
		$comment2 .= "*/\n";
		$comment = $comment2;
	}

	$consts = generate_column_const($cols);

	$code = <<<EXPORTCODE
var {$export_table['table']}_data = {
	{$comment}

	{$consts},

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		{$data}
	],
};

EXPORTCODE;
	if (isset($export_table['filename'])) {
		file_put_contents(EXPORT_DBDATA_URL."{$export_table['filename']}_data.js", $code);
	} else {
		file_put_contents(EXPORT_DBDATA_URL."{$export_table['table']}_data.js", $code);
	}
}

function generate_column_const(&$cols) {
	$consts = array();
	
	$len = count($cols);
	for ($i = 0; $i < $len; $i++) {
		$consts[] = ucfirst($cols[$i][0])." : ".$i;
	}
	
	return join(",\n	", $consts);
}
?>