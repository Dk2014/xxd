<?php

function gen_cgo_h($output_dir, $tables) {
	$code = '';
	$code .= "#ifndef _cgo_types_h_\n";
	$code .= "#define _cgo_types_h_\n\n";

	$code .= "#include <stdint.h>\n";
	$code .= "#include <stdlib.h>\n";
	$code .= "#include <string.h>\n";

	$code .= "extern void Init();\n\n";

	$code .= "typedef struct GlobalTables GlobalTables;\n";
	$code .= "typedef struct PlayerTables PlayerTables;\n\n";

	$code .= "extern GlobalTables g_GlobalTables;\n";

	$code .= "extern PlayerTables* NewPlayerTables();\n";

	foreach ($tables as $table) {
		$code .= "/* ========== {$table['name']} ========== */\n\n";

		$table_type_name = get_type_name($table['name']);

		$code .= "typedef struct {$table['name']} {\n";

		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_c_type = get_c_type($column['type']);

			$code .= "	{$column_c_type} {$column_type_name};\n";
			if ($column_c_type == 'char*') {
				$code .= "	int {$column_type_name}_len;\n";
			}
		}

		if (is_hashtable($table)) {
			$code .= "	struct {$table['name']}* next;\n";
		}

		$code .= "} {$table_type_name};\n\n";

		$code .= "extern {$table_type_name}* New_{$table_type_name}();\n";
		$code .= "extern void Free_{$table_type_name}({$table_type_name}*);\n";
		$code .= "\n";
	}
	
	$code .= "/* ================================== */\n\n";

	$code .= "struct GlobalTables {\n";
	foreach ($tables as $table) {
		if (strpos($table['name'], 'global') === 0) {
			$table_type_name = get_type_name($table['name']);
			$code .= "	{$table_type_name}* {$table_type_name};\n";
		}
	}
	$code .= "};\n\n";

	$code .= "struct PlayerTables {\n";
	$code .= "	int64_t Pid;\n";
	foreach ($tables as $table) {
		if (strpos($table['name'], 'player') === 0) {
			$table_type_name = get_type_name($table['name']);
			$code .= "	{$table_type_name}* {$table_type_name};\n";
		}
	}
	$code .= "};\n\n";

	$code .= "#endif\n";

	save_code($output_dir."/zd_cgo.h", $code);
}

function gen_cgo_c($output_dir, $tables) {
	$code = '';
	$code .= "#include \"zd_cgo.h\"\n\n";

	$code .= "GlobalTables g_GlobalTables;\n";
	$code .= "PlayerTables* g_PlayerTables;\n\n";

	$code .= "void Init() {\n";
	$code .= "	bzero(&g_GlobalTables, sizeof(GlobalTables));\n";
	$code .= "}\n\n";

	$code .= "// 创建并返回新的玩家数据库切片\n";
	$code .= "PlayerTables* NewPlayerTables() {\n";
	$code .= "	return (PlayerTables*)calloc(1, sizeof(PlayerTables));\n";
	$code .= "}\n\n";

	foreach ($tables as $table) {		
		$code .= "/* ========== {$table['name']} ========== */\n\n";

		$table_type_name = get_type_name($table['name']);

		$code .= "{$table_type_name}* New_{$table_type_name}() {\n";
		$code .= "	return ({$table_type_name}*)calloc(1, sizeof({$table_type_name}));\n";
		$code .= "}\n\n";

		$code .= "void Free_{$table_type_name}({$table_type_name}* row) {\n";
		foreach ($table['cols'] as $column) {
			if (get_c_type($column['type']) == 'char*') {
				$column_type_name = get_type_name($column['name']);

				$code .= "	if (row->{$column_type_name} != NULL) {\n";
				$code .= "		free(row->{$column_type_name});\n";
				$code .= "	}\n";
			}
		}
		$code .= "	free(row);\n";
		$code .= "}\n\n";
	}

	save_code($output_dir."/zd_cgo.c", $code);
}

?>