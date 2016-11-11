<?php

function gen_sync_go($output_dir, $tables) {
	$code = '';

	$code .= "package mdb\n\n";

	$code .= "func (file *SyncFile) initDB() error {\n";
	$code .= "	file.BeginTrans()\n";
	foreach ($tables as $table) {
		$code .= "	file.WriteBytes([]byte(\"".$table['name'].",";
		$cols = array();
		foreach ($table['cols'] as $column) {
			$cols[] = $column['name'];
		}
		$code .= implode($cols, ',');

		$code .= "\"))\n";

		$code .= "	file.WriteBytes([]byte{";
		$cols = array();
		foreach ($table['cols'] as $column) {
			$cols[] = get_sync_head_tag($column['type'])."/*{$column['name']} - {$column['type']}*/";
		}
		$code .= implode($cols, ', ');
		$code .= "})\n\n";
	}

	$code .= "	return file.EndTrans()\n";
	$code .= "}\n\n";

	save_code($output_dir."/zd_sync.go", $code);
}