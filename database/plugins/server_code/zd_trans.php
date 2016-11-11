<?php

function gen_trans_go($output_dir, $tables) {
	$code = '';
	$code .= "package mdb\n\n";
	$code .= "/*\n";
	$code .= "#include \"zd_cgo.h\"\n";
	$code .= "*/\n";
	$code .= "import \"C\"\n";
	$code .= "import \"unsafe\"\n\n";
	$code .= "import \"time\"\n\n";

	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);
		$primary_key_type_name = get_type_name($table['pri']['name']);
		$primary_key_go_type = get_go_type($table['pri']['type']);
		$primary_key_c_type = get_c_type($table['pri']['type']);

		$code .= "/* ========== {$table['name']} ========== */\n\n";

		$code .= "type {$table_type_name}InsertLog struct {\n";
		$code .= "	db *Database \n";
		$code .= "	Row *C.{$table_type_name}\n";
		$code .= "	GoRow *{$table_type_name}\n";
		$code .= "}\n\n";

		// Insert
		$code .= "func (db *Database) new{$table_type_name}InsertLog(row *C.{$table_type_name}, goRow *{$table_type_name}) *{$table_type_name}InsertLog {\n";
		$code .= "	return &{$table_type_name}InsertLog{db:db, Row:row, GoRow:goRow}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) Free() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) InvokeHook() {\n";
		$code .= "	if g_Hooks.{$table_type_name} != nil {\n";
		$code .= "		g_Hooks.{$table_type_name}.Insert(&{$table_type_name}Row{row: log.Row})\n";
		$code .= "	}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) SetEventId(time.Time) {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) CommitToFile(file *SyncFile) {\n";
		$code .= "	file.WriteUint8(1)\n";
		$code .= "	file.WriteInt32(".$table['id'].")\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$sync_write_type = get_sync_write_type($column['type']);
			$code .= "	file.{$sync_write_type}(log.GoRow.{$column_type_name})\n";
		}
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}InsertLog) CommitToMySql() error {\n";
		$code .= "	stmt := g_SQL.{$table_type_name}.Insert\n";
		$code .= "	stmt.CleanBind()\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$bind_type = get_bind_type($column['type']);
			if ($bind_type == 'BindVarchar' || $bind_type == 'BindBinary') {
				$code .= "	stmt.{$bind_type}(unsafe.Pointer(log.Row.{$column_type_name}), int(log.Row.{$column_type_name}_len))\n";
			} else {
				$code .= "	stmt.{$bind_type}(unsafe.Pointer(&(log.Row.{$column_type_name})))\n";
			}
		}
		$code .= "	return stmt.Execute()\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) CommitToTLog() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}InsertLog) CommitToXdLog() {\n";
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}InsertLog) GetTransType() int8 {\n";
		$code .= "	return TRANS_TYPE_MYSQL\n";
		$code .= "}\n\n";
		
		// Delete
		$code .= "type {$table_type_name}DeleteLog struct {\n";
		$code .= "	db *Database\n";
		$code .= "	Old *C.{$table_type_name}\n";
		$code .= "	GoOld *{$table_type_name}\n";
		$code .= "}\n\n";

		$code .= "func (db *Database) new{$table_type_name}DeleteLog(old *C.{$table_type_name}, goOld *{$table_type_name}) *{$table_type_name}DeleteLog {\n";
		$code .= "	return &{$table_type_name}DeleteLog{db:db, Old:old, GoOld:goOld}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) Free() {\n";
		$code .= "	C.Free_{$table_type_name}(log.Old)\n";
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}DeleteLog) InvokeHook() {\n";
		$code .= "	if g_Hooks.{$table_type_name} != nil {\n";
		$code .= "		g_Hooks.{$table_type_name}.Delete(&{$table_type_name}Row{row: log.Old})\n";
		$code .= "	}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) SetEventId(time.Time) {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) CommitToFile(file *SyncFile) {\n";
		$code .= "	file.WriteUint8(2)\n";
		$code .= "	file.WriteInt32(".$table['id'].")\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$sync_write_type = get_sync_write_type($column['type']);
			$code .= "	file.{$sync_write_type}(log.GoOld.{$column_type_name})\n";
		}
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}DeleteLog) CommitToMySql() error {\n";
		$code .= "	stmt := g_SQL.{$table_type_name}.Delete\n";
		$code .= "	stmt.CleanBind()\n";
		$code .= "	stmt.".get_bind_type($table['pri']['type'])."(unsafe.Pointer(&(log.Old.{$primary_key_type_name})))\n";
		$code .= "	return stmt.Execute()\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) CommitToTLog() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) CommitToXdLog() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}DeleteLog) GetTransType() int8 {\n";
		$code .= "	return TRANS_TYPE_MYSQL\n";
		$code .= "}\n\n";

		// Update
		$code .= "type {$table_type_name}UpdateLog struct {\n";
		$code .= "	db *Database \n";
		$code .= "	Row *C.{$table_type_name}\n";
		$code .= "	Old *C.{$table_type_name}\n";
		$code .= "	GoNew *{$table_type_name}\n";
		$code .= "	GoOld *{$table_type_name}\n";
		$code .= "}\n\n";

		$code .= "func (db *Database) new{$table_type_name}UpdateLog(row, old *C.{$table_type_name}, goNew, goOld *{$table_type_name}) *{$table_type_name}UpdateLog {\n";
		$code .= "	return &{$table_type_name}UpdateLog{db:db, Row:row, Old:old, GoNew:goNew, GoOld:goOld}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) Free() {\n";
		$code .= "	C.Free_{$table_type_name}(log.Old)\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) InvokeHook() {\n";
		$code .= "	if g_Hooks.{$table_type_name} != nil {\n";
		$code .= "		g_Hooks.{$table_type_name}.Update(&{$table_type_name}Row{row: log.Row}, &{$table_type_name}Row{row: log.Old})\n";
		$code .= "	}\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) SetEventId(time.Time) {\n";
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}UpdateLog) CommitToFile(file *SyncFile) {\n";
		$code .= "	file.WriteUint8(3)\n";
		$code .= "	file.WriteInt32(".$table['id'].")\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$sync_write_type = get_sync_write_type($column['type']);
			$code .= "	file.{$sync_write_type}(log.GoNew.{$column_type_name})\n";
		}
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$sync_write_type = get_sync_write_type($column['type']);
			$code .= "	file.{$sync_write_type}(log.GoOld.{$column_type_name})\n";
		}
		$code .= "}\n\n";
		
		$code .= "func (log *{$table_type_name}UpdateLog) CommitToMySql() error {\n";
		$code .= "	stmt := g_SQL.{$table_type_name}.Update\n";
		foreach ($table['cols'] as $column) {
			if ($column['name'] == $table['pri']['name'])
				continue;
			$column_type_name = get_type_name($column['name']);
			$bind_type = get_bind_type($column['type']);
			if ($bind_type == 'BindVarchar' || $bind_type == 'BindBinary') {
				$code .= "	stmt.{$bind_type}(unsafe.Pointer(log.Row.{$column_type_name}), int(log.Row.{$column_type_name}_len))\n";
			} else {
				$code .= "	stmt.{$bind_type}(unsafe.Pointer(&(log.Row.{$column_type_name})))\n";
			}
		}
		$code .= "	stmt.".get_bind_type($table['pri']['type'])."(unsafe.Pointer(&(log.Row.{$primary_key_type_name})))\n";
		$code .= "	stmt.CleanBind()\n";
		$code .= "	return stmt.Execute()\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) CommitToTLog() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) CommitToXdLog() {\n";
		$code .= "}\n\n";

		$code .= "func (log *{$table_type_name}UpdateLog) GetTransType() int8 {\n";
		$code .= "	return TRANS_TYPE_MYSQL\n";
		$code .= "}\n\n";

		if (is_hashtable($table)) {
			$target_c_table = "&C.g_GlobalTables";
			$select_table = "C.g_GlobalTables";

			if (strpos($table['name'], 'player') === 0) {
				$target_c_table = "log.db.tables";
				$select_table = "log.db.tables";
			}

			$code .= "func (log *{$table_type_name}InsertLog) Rollback() {\n";
			$code .= "	key := log.Row.{$primary_key_type_name}\n";
			$code .= "	var old, prev *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; prev, crow = crow, crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old != log.Row {\n";
			$code .= "		panic(\"insert rollback check failed '{$table_type_name}'\")\n";
			$code .= "	}\n";
			$code .= "	if prev != nil {\n";
			$code .= "		prev.next = old.next\n";
			$code .= "	} else {\n";
			$code .= "		{$select_table}.{$table_type_name} = {$select_table}.{$table_type_name}.next\n";
			$code .= "	}\n";
			$code .= "	C.Free_{$table_type_name}(log.Row)\n";
			$code .= "}\n\n";

			$code .= "func (log *{$table_type_name}DeleteLog) Rollback() {\n";
			$code .= "	key := log.Old.${primary_key_type_name}\n";
			$code .= "	var old *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; crow = crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old != nil {\n";
			$code .= "		panic(\"delete rollback check failed '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	log.Old.next = {$select_table}.{$table_type_name}\n";
			$code .= "	{$select_table}.{$table_type_name} = log.Old\n";
			$code .= "}\n\n";

			$code .= "func (log *{$table_type_name}UpdateLog) Rollback() {\n";
			$code .= "	key := log.Old.${primary_key_type_name}\n";
			$code .= "	var old, prev *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; prev, crow = crow, crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old != log.Row {\n";
			$code .= "		panic(\"update rollback check failed '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	log.Old.next = old.next\n";
			$code .= "	if prev != nil {\n";
			$code .= "		prev.next = log.Old\n";
			$code .= "	} else {\n";
			$code .= "		{$select_table}.{$table_type_name} = log.Old\n";
			$code .= "	}\n";
			$code .= "	C.Free_{$table_type_name}(log.Row)\n";
			$code .= "}\n\n";
		} else if (is_one_to_one_player_table($table)) {
			$code .= "func (log *{$table_type_name}InsertLog) Rollback() {\n";
			$code .= "	if log.db.tables.{$table_type_name} != log.Row {\n";
			$code .= "		panic(\"insert rollback check failed '{$table_type_name}'\")\n";
			$code .= "	}\n";
			$code .= "	log.db.tables.{$table_type_name} = nil\n";
			$code .= "	C.Free_{$table_type_name}(log.Row)\n";
			$code .= "}\n\n";

			$code .= "func (log *{$table_type_name}DeleteLog) Rollback() {\n";
			$code .= "	if log.db.tables.{$table_type_name} != nil {\n";
			$code .= "		panic(\"delete rollback check failed '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	log.db.tables.{$table_type_name} = log.Old\n";
			$code .= "}\n\n";

			$code .= "func (log *{$table_type_name}UpdateLog) Rollback() {\n";
			$code .= "	if log.db.tables.{$table_type_name} != log.Row {\n";
			$code .= "		panic(\"update rollback check failed '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	log.db.tables.{$table_type_name} = log.Old\n";
			$code .= "	C.Free_{$table_type_name}(log.Row)\n";
			$code .= "}\n\n";
		}
	}

	save_code($output_dir."/zd_trans.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_trans.go");
}
?>