<?php

function gen_tables_go($output_dir, $tables) {
	$code = '';
	$code .= "package mdb\n\n";
	$code .= "/*\n";
	$code .= "#include \"zd_cgo.h\"\n";
	$code .= "*/\n";
	$code .= "import \"C\"\n";
	$code .= "import \"sync/atomic\"\n";
	$code .= "import \"fmt\"\n";

	$code .= "
type insertOP struct{ db *Database }
type deleteOP struct{ db *Database }
type updateOP struct{ db *Database }
type lookupOP struct{ db *Database }
type selectOP struct{ db *Database }
type allOP    struct{ db *Database }


type Database struct {
	playerId int64
	cId int32
	tables *C.PlayerTables
	Insert insertOP
	Delete deleteOP
	Update updateOP
	Lookup lookupOP
	Select selectOP
	All    allOP
	transController *transController
}

func NewDatabase() *Database {
	return newDatabase(newTransController(g_LockManager.Get(0), g_FileCommitter))
}

func newDatabase(transController *transController) *Database {
	db := new(Database)
	db.Insert.db = db
	db.Delete.db = db
	db.Update.db = db
	db.Lookup.db = db
	db.Select.db = db
	db.All.db    = db
	db.transController = transController
	return db
}

func (db *Database) Init() {
	db.tables = C.NewPlayerTables()
}

func (db *Database) Mount(playerId int64,cId int32) {
	db.mountTables2(playerId,cId)
	db.transController.lock = g_LockManager.Get(playerId)
}

func (db *Database) AgentExecute(pid int64, cb func(*Database)) {
	agentDB := newDatabase(db.transController)
	agentDB.mountTables(pid)
	cb(agentDB)
}

func (db *Database) Rollback() {
	        db.transController.Rollback()
}

func (db *Database) mountTables2(playerId int64,cId int32) {
	db.playerId = playerId
	db.cId = cId
	db.tables = GetPlayerTables(playerId)
	if db.tables == nil {
		panic(\"tables is nil\")
	}
}

func (db *Database) mountTables(playerId int64) {
	db.playerId = playerId
	db.tables = GetPlayerTables(playerId)
	if db.tables == nil {
		panic(\"tables is nil\")
	}
}

func (db *Database) PlayerId() int64 {
	return db.playerId
}

func (db *Database) CId() int32 {
	return db.cId
}

func (db *Database) Transaction(transId uint32, work func()) {
	if db != g_Database {
		g_GlobalLock.RLock()
		defer g_GlobalLock.RUnlock()
	}

	db.transController.Transaction(transId, work)
}

func (db *Database) AddTLog(log TransLog) {
	db.transController.AddTransLog(log)
}

func (db *Database) AddXdLog(log TransLog) {
	db.transController.AddTransLog(log)
}

func (db *Database) addTransLog(log TransLog) {
	db.transController.AddTransLog(log)
}

\n";

	$code .= "type rowIds struct {\n";
	foreach ($tables as $table) {
		if ($table['pri']['name'] == 'id') {
			$code .= "	".get_type_name($table['name'])." int64\n";
		}
	}
	/*
	foreach ($log_tables as $table) {
		if ($table['pri']['name'] != 'id') {
			continue;
		}
		$code .= "	".get_type_name($table['name'])." int64\n";
	}
	*/
	$code .= "}\n\n";

	$code .= "var g_RowIds rowIds\n\n";
	$code .= "func initRowIdsWithServerId(serverId int64) {\n";
		foreach ($tables as $table) {
		if ($table['name'] == 'global_clique') {
			//FIXME 策划需求：帮派ID从1000开始增加，需要考虑一个机制 处理这些特殊情况
			$code .= "	g_RowIds.".get_type_name($table['name'])." = 1000\n";
			continue;
		}
		if ($table['pri']['name'] == 'id') {
			$code .= "	g_RowIds.".get_type_name($table['name'])." = serverId\n";
		}
	}
	$code .= "}\n\n";


	foreach ($tables as $table) {
		$table_type_name = get_type_name($table['name']);
		$primary_key_type_name = get_type_name($table['pri']['name']);
		$primary_key_go_type = get_go_type($table['pri']['type']);
		$primary_key_c_type = get_c_type($table['pri']['type']);

		$code .= "/* ========== {$table['name']} ========== */\n\n";

		$code .= "// ".$table['desc']."\n";
		$code .= "type {$table_type_name} struct {\n";
		foreach ($table['cols'] as $column) {
			$code .= "	".get_type_name($column['name'])." ".get_go_type($column['type'])." // ".$column['desc']."\n";
		}
		$code .= "}\n\n";

		$code .= "func (this *{$table_type_name}) CObject() *C.{$table_type_name} {\n";
		$code .= "	object := C.New_{$table_type_name}()\n";
		foreach ($table['cols'] as $column) {
			$column_type_name = get_type_name($column['name']);
			$column_c_type = get_c_type($column['type']);
			if ($column_c_type == "char*") {
				$code .= "	object.{$column_type_name} = C.CString(string(this.{$column_type_name}))\n";
				$code .= "	object.{$column_type_name}_len = C.int(len(this.{$column_type_name}))\n";
			} else {
				$code .= "	object.{$column_type_name} = C.{$column_c_type}(this.{$column_type_name})\n";
			}
		}
		$code .= "	return object\n";
		$code .= "}\n\n";

		if (is_hashtable($table)) {
			$target_c_table = "&C.g_GlobalTables";
			$select_table = "C.g_GlobalTables";

			if (strpos($table['name'], 'player') === 0) {
				$target_c_table = "this.db.tables";
				$select_table = "this.db.tables";
			}

			$code .= "func (this insertOP) {$table_type_name}(row *{$table_type_name}) {\n";
			if ($table['pri']['name'] == 'id') {
				$code .= "	row.{$primary_key_type_name} = atomic.AddInt64(&g_RowIds.{$table_type_name}, 1)\n";
				//$code .= "	row.{$primary_key_type_name} = g_RowIds.{$table_type_name}\n";
			}
			$code .= "	key := C.{$primary_key_c_type}(row.${primary_key_type_name})\n";
			$code .= "	var old *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; crow = crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old != nil {\n";
			$code .= "		panic(\"duplicate insert '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	newRow := row.CObject()\n";
			$code .= "	newRow.next = {$select_table}.{$table_type_name}\n";
			$code .= "	{$select_table}.{$table_type_name} = newRow\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}InsertLog(newRow, row))\n";
			$code .= "}\n\n";

			$code .= "func (this deleteOP) {$table_type_name}(row *{$table_type_name}) {\n";
			$code .= "	key := C.{$primary_key_c_type}(row.${primary_key_type_name})\n";
			$code .= "	var old, prev *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; prev, crow = crow, crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old == nil {\n";
			$code .= "		panic(\"delete not exists '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	if prev != nil {\n";
			$code .= "		prev.next = old.next\n";
			$code .= "	} else {\n";
			$code .= "		{$select_table}.{$table_type_name} = {$select_table}.{$table_type_name}.next\n";
			$code .= "	}\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}DeleteLog(old, row))\n";
			$code .= "}\n\n";

			$code .= "func (this updateOP) {$table_type_name}(row *{$table_type_name}) {\n";
			$code .= "	key := C.{$primary_key_c_type}(row.${primary_key_type_name})\n";
			$code .= "	var old, prev *C.{$table_type_name}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; prev, crow = crow, crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == key {\n";
			$code .= "			old = crow\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	if old == nil {\n";
			$code .= "		panic(\"update not exists '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	newRow := row.CObject()\n";
			$code .= "	newRow.next = old.next\n";
			$code .= "	if prev != nil {\n";
			$code .= "		prev.next = newRow\n";
			$code .= "	} else {\n";
			$code .= "		{$select_table}.{$table_type_name} = newRow\n";
			$code .= "	}\n";
			$code .= "	tmpRow := {$table_type_name}Row{row:old}\n";
			$code .= "	goOld := tmpRow.GoObject()\n";
			$code .= "	tmpRow.row = nil\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}UpdateLog(newRow, old, row, goOld))\n";
			$code .= "}\n\n";

			$code .= "func (this lookupOP) {$table_type_name}(key {$primary_key_go_type}) *{$table_type_name} {\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; crow = crow.next {\n";
			$code .= "		if crow.{$primary_key_type_name} == C.{$primary_key_c_type}(key) {\n";
			$code .= "			tmpRow := {$table_type_name}Row{row:crow}\n";
			$code .= "			goRow := tmpRow.GoObject()\n";
			$code .= "			tmpRow.row = nil\n";
			$code .= "			return goRow\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	return nil\n";
			$code .= "}\n\n";

			$code .= "func (this selectOP) {$table_type_name}(callback func(*{$table_type_name}Row)) {\n";
			$code .= "	row := &{$table_type_name}Row{}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; crow = crow.next {\n";
			$code .= "		row.reset(crow)\n";
			$code .= "		callback(row)\n";
			$code .= "		if row.isBreak {\n";
			$code .= "			break\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "	row.row = nil\n";
			$code .= "}\n\n";

			$code .= "func (this allOP) {$table_type_name}() (rows []interface{}) {\n";
			$code .= "	row := &{$table_type_name}Row{}\n";
			$code .= "	for crow := {$select_table}.{$table_type_name}; crow != nil; crow = crow.next {\n";
			$code .= "		row.reset(crow)\n";
			$code .= "		rows = append(rows, row.GoObject())\n";
			$code .= "		fmt.Println(\"{$table_type_name}\", row.GoObject())\n";
			$code .= "	}\n";
			$code .= "	row.row = nil\n";
			$code .= "	fmt.Println(\"{$table_type_name}\", len(rows))\n";
			$code .= "	return rows\n";
			$code .= "}\n\n";

		} else if (is_one_to_one_player_table($table)) {
			$code .= "func (this insertOP) {$table_type_name}(row *{$table_type_name}) {\n";
			$code .= "	if this.db.tables.{$table_type_name} != nil {\n";
			$code .= "		panic(\"duplicate insert '{$table['name']}'\")\n";
			$code .= "	}\n";
			if ($table['pri']['name'] == 'id' || $table['name'] == 'player') {
				$code .= "	row.{$primary_key_type_name} = atomic.AddInt64(&g_RowIds.{$table_type_name}, 1)\n";
			}
			if ($table['name'] == 'player') {
				$code .= "	this.db.playerId = row.{$primary_key_type_name}\n";
				$code .= "  this.db.transController.lock = g_LockManager.Get(this.db.playerId)\n";
				$code .= "	this.db.tables.Pid = C.int64_t(row.{$primary_key_type_name})\n";
				$code .= "	SetPlayerTables(this.db.tables)\n";
			}
			$code .= "	newRow := row.CObject()\n";
			$code .= "	this.db.tables.{$table_type_name} = newRow\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}InsertLog(newRow, row))\n";
			$code .= "}\n\n";

			$code .= "func (this deleteOP) {$table_type_name}(row *{$table_type_name}) {\n";
			$code .= "	if this.db.tables.{$table_type_name} == nil {\n";
			$code .= "		panic(\"delete not exists '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	old := this.db.tables.{$table_type_name}\n";
			$code .= "	this.db.tables.{$table_type_name} = nil\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}DeleteLog(old, row))\n";
			$code .= "}\n\n";

			$code .= "func (this updateOP) {$table_type_name}(row *{$table_type_name}) {\n";
			$code .= "	if this.db.tables.{$table_type_name} == nil {\n";
			$code .= "		panic(\"update not exists '{$table['name']}'\")\n";
			$code .= "	}\n";
			$code .= "	newRow := row.CObject()\n";
			$code .= "	old := this.db.tables.{$table_type_name}\n";
			$code .= "	this.db.tables.{$table_type_name} = newRow\n";
			$code .= "	tmpRow := {$table_type_name}Row{row:old}\n";
			$code .= "	goOld := tmpRow.GoObject()\n";
			$code .= "	tmpRow.row = nil\n";
			$code .= "	this.db.addTransLog(this.db.new{$table_type_name}UpdateLog(newRow, old, row, goOld))\n";
			$code .= "}\n\n";

			$code .= "func (this lookupOP) {$table_type_name}(key {$primary_key_go_type}) *{$table_type_name} {\n";
			$code .= "	if this.db.tables.{$table_type_name} == nil {\n";
			$code .= "		return nil\n";
			$code .= "	}\n";
			$code .= "	tmpRow := {$table_type_name}Row{row:this.db.tables.{$table_type_name}}\n";
			$code .= "	goRow := tmpRow.GoObject()\n";
			$code .= "	tmpRow.row = nil\n";
			$code .= "	return goRow\n";
			$code .= "}\n\n";
		}
	}

	save_code($output_dir."/zd_tables.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_tables.go");
}

?>
