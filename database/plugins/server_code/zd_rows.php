<?php

function gen_rows_go($output_dir, $tables) {
	$code = '';
	$code .= "package mdb\n\n";
	$code .= "/*\n";
	$code .= "#include \"zd_cgo.h\"\n";
	$code .= "*/\n";
	$code .= "import \"C\"\n";
	$code .= "import \"unsafe\"\n\n";

	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);

		$code .= "/* ========== {$table['name']} ========== */\n\n";
			
		$code .= "// ".$table['desc']."\n";
		$code .= "type {$table_type_name}Row struct {\n";
		$code .= "	row *C.{$table_type_name}\n";
		$code .= "	isBreak bool\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_go_type = get_go_type($column['type']);
			if ($column_go_type == 'string' || $column_go_type == '[]byte') {
				$code .= "	cached_{$column_type_name} bool\n";
				$code .= "	cacheof_{$column_type_name} $column_go_type\n";
			}
		}
		$code .= "}\n\n";

		$code .= "func (this *{$table_type_name}Row) reset(row *C.{$table_type_name}) {";
		$code .= "	this.row = row\n";
		$code .= "	this.isBreak = false\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_go_type = get_go_type($column['type']);
			if ($column_go_type == 'string' || $column_go_type == '[]byte') {
				$code .= "	this.cached_{$column_type_name} = false\n";
			}
		}
		$code .= "}\n\n";

		$code .= "func (this *{$table_type_name}Row) Break() {\n";
		$code .= "	this.isBreak = true\n";
		$code .= "}\n\n";

		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_go_type = get_go_type($column['type']);

			$code .= "// ".$column['desc']."\n";
			$code .= "func (this *{$table_type_name}Row) {$column_type_name}() {$column_go_type} {\n";
			if ($column_go_type == 'string') {
				$code .= "	if !this.cached_{$column_type_name} {\n";
				$code .= "		this.cached_{$column_type_name} = true\n";
				$code .= "		this.cacheof_{$column_type_name} = C.GoStringN(this.row.{$column_type_name}, this.row.{$column_type_name}_len)\n";
				$code .= "	}\n";
				$code .= "	return this.cacheof_{$column_type_name}\n";
			} else if ($column_go_type == '[]byte') {
				$code .= "	if !this.cached_{$column_type_name} {\n";
				$code .= "		this.cached_{$column_type_name} = true\n";
				$code .= "		this.cacheof_{$column_type_name} = C.GoBytes(unsafe.Pointer(this.row.{$column_type_name}), this.row.{$column_type_name}_len)\n";
				$code .= "	}\n";
				$code .= "	return this.cacheof_{$column_type_name}\n";
			} else {
				$code .= "	return {$column_go_type}(this.row.{$column_type_name})\n";
			}
			$code .= "}\n\n";
		}

		$code .= "func (this *{$table_type_name}Row) GoObject() *{$table_type_name} {\n";
		$code .= "	return &{$table_type_name}{\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_go_type = get_go_type($column['type']);

			$code .= "		{$column_type_name}: ";
			if ($column_go_type == 'string' || $column_go_type == '[]byte') {
				$code .= "this.{$column_type_name}(),\n";
			} else {
				$code .= "{$column_go_type}(this.row.{$column_type_name}),\n";
			}
		}
		$code .= "	}\n";
		$code .= "}\n\n";
	}

	save_code($output_dir."/zd_rows.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_rows.go");
}

?>