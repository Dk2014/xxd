<?php

function gen_hooks_go($output_dir, $tables) {
	$code = '';
	$code .= "package mdb\n\n";

	$code .= "var g_Hooks tableHooks\n\n";

	$code .= "type tableHooks struct {\n";
	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);
		$code .= "	{$table_type_name} {$table_type_name}Hook\n";
	}
	$code .= "}\n\n";

	$code .= "type hooksOP int\n";
	$code .= "const Hook hooksOP = 0\n";

	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);
		$code .= "func (h hooksOP) {$table_type_name}(hook {$table_type_name}Hook) {\n";
		$code .= "	if g_Hooks.{$table_type_name} != nil {\n";
		$code .= "		panic(\"duplicate hook {$table_type_name}\")\n";
		$code .= "	}\n";
		$code .= "	g_Hooks.{$table_type_name} = hook\n";
		$code .= "}\n";
	}

	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);

		$code .= "/* ========== {$table['name']} ========== */\n\n";
		$code .= "// ".$table['desc']."\n";
		$code .= "type {$table_type_name}Hook interface {\n";
		$code .= "	Load(row *{$table_type_name}Row)\n";
		$code .= "	Insert(row *{$table_type_name}Row)\n";
		$code .= "	Delete(old *{$table_type_name}Row)\n";
		$code .= "	Update(row, old *{$table_type_name}Row)\n";
		$code .= "}\n\n";
	}

	save_code($output_dir."/zd_hooks.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_hooks.go");
}

?>