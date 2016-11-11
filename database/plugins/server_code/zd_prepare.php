<?php

function gen_prepare_go($output_dir, $tables, $log_tables) {
	$code = '';
	$code .= "package mdb\n\n";

	$code .= "import \"core/mysql\"\n";

	$code .= "
func prepare(db *mysql.Connection, sql string) *mysql.Stmt {
	stmt, err := db.Prepare([]byte(sql))
	if err != nil {
		panic(err)
	}
	return stmt
}
";

	$code .= "var g_SQL tableSQL\n\n";

	$code .= "type tableSQL struct {\n";
	foreach ($tables as $table) {
		$code .= "	".get_type_name($table['name']).' '.get_var_name($table['name'])."SQL\n";
	}
	foreach ($log_tables as $table) {
		$code .= "	".get_type_name($table['name']).' '.get_var_name($table['name'])."SQL\n";
	}
	$code .= "}\n\n";

	foreach ($tables as $table) {
		$var_name = get_var_name($table['name']);
		$code .= "type {$var_name}SQL struct {\n";
		$code .= "	Insert *mysql.Stmt\n";
		$code .= "	Delete *mysql.Stmt\n";
		$code .= "	Update *mysql.Stmt\n";
		$code .= "}\n\n";
	}

	foreach ($log_tables as $table) {
		$var_name = get_var_name($table['name']);
		$code .= "type {$var_name}SQL struct {\n";
		$code .= "	Insert *mysql.Stmt\n";
		$code .= "}\n\n";
	}

	$code .= "func prepareSQL(db *mysql.Connection) {\n";
	foreach ($tables as $table) {
		$table_name = $table['name'];
		$var_name = get_var_name($table_name);
		$type_name = get_type_name($table_name);

		$code .= "	g_SQL.{$type_name}.Insert = prepare(db, \"INSERT INTO `$table_name` SET";
		foreach ($table['cols'] as $column) {
			$code .= " `{$column['name']}`=?,";
		}
		$code = rtrim($code, ',');
		$code .= "\")\n";
		$code .= "	g_SQL.{$type_name}.Update = prepare(db, \"UPDATE `$table_name` SET";
		foreach ($table['cols'] as $column) {
			if ($column['name'] == $table['pri']['name']) 
				continue;
			$code .= " `{$column['name']}`=?,";
		}
		$code = rtrim($code, ',');
		$code .= " WHERE `{$table['pri']['name']}`=?";
		$code .= "\")\n";
		$code .= "	g_SQL.{$type_name}.Delete = prepare(db, \"DELETE FROM `$table_name` WHERE `{$table['pri']['name']}`=?\")\n";
	}

	foreach ($log_tables as $table) {
		$table_name = $table['name'];
		$var_name = get_var_name($table_name);

		$code .= "	g_SQL.{$type_name}.Insert = prepare(db, \"REPLACE INTO `$table_name` SET";
		foreach ($table['cols'] as $column) {
			$code .= " `{$column['name']}`=?,";
		}
		$code = rtrim($code, ',');
		$code .= "\")\n";
	}
	$code .= "}\n\n";


	$code .= "func prepareSQLClean() {\n";
	foreach ($tables as $table) {
		$table_name = $table['name'];
		$var_name = get_var_name($table_name);
		$type_name = get_type_name($table_name);

		$code .= "	g_SQL.{$type_name}.Insert.Close()\n";
		$code .= "	g_SQL.{$type_name}.Update.Close()\n";
		$code .= "	g_SQL.{$type_name}.Delete.Close()\n";
	}

	foreach ($log_tables as $table) {
		$table_name = $table['name'];
		$var_name = get_var_name($table_name);

		$code .= "	g_SQL.{$type_name}.Insert.Close()\n";
	}
	$code .= "}\n\n";	

	save_code($output_dir."/zd_prepare.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_prepare.go");
}

?>