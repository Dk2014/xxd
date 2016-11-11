<?php
function gen_load_go($output_dir, $tables) {
	$code = '';
	$code .= "package mdb\n\n";
	$code .= "/*\n";
	$code .= "#include \"zd_cgo.h\"\n";
	$code .= "*/\n";
	$code .= "import \"C\"\n";
	$code .= "import \"strconv\"\n";
	$code .= "import \"core/log\"\n";
	$code .= "import \"core/mysql\"\n";

	$code .= "
func itoa(i int) string {
	return strconv.Itoa(i)
}

func playerIdToStr(playerId int64) string {
	return strconv.FormatInt(playerId, 10)
}

func initRowId(db *mysql.Connection, id *int64, sql string) bool {
	qr, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}
	row := qr.Rows[0]
	if row[0] == nil {
		return false
	}
	*id = row.Int64(0)
	return true
}

func newPlayerTables(db *mysql.Connection) {
	sql := \"SELECT `id` FROM `player`\"

	res, err := db.ExecuteFetch([]byte(sql), -1)
	if err != nil {
		panic(err)
	}

	iId := res.Map(\"id\")

	for _, row := range res.Rows {
		playerId := C.int64_t(row.Int64(iId))

		playerDb := C.NewPlayerTables()
		playerDb.Pid = C.int64_t(playerId)

		SetPlayerTables(playerDb)
	}
}

";

	// 加载互动服务器全局表数据
	$code .= "func initGlobalTables(db *mysql.Connection) {\n";
	foreach ($tables as $table) {
		$table_name = $table['name'];
		if (strpos($table_name, 'global') === 0) {
			$table_type_name = get_type_name($table['name']);
			$code .= "	load{$table_type_name}(db)\n";
		}
	}
	$code .= "}\n\n";

	// 加载互动服务器玩家表数据
	$code .= "func initPlayerGlobalTables(db *mysql.Connection, sid, wgn, procId int) {\n";
	foreach ($tables as $table) {
		$table_name = $table['name'];
		if (strpos($table_name, 'player_global') === 0) {
			$table_type_name = get_type_name($table['name']);
			$code .= "	load{$table_type_name}(db, 0, sid, wgn, procId)\n";
		}
	}
	$code .= "}\n\n";

	// 加载游戏服玩家表数据
	$code .= "func initPlayerTables(db *mysql.Connection, sid, wgn, procId int) {\n";
	foreach ($tables as $table) {
		$table_name = $table['name'];
		// 忽略互动服务器玩家表
		if (strpos($table_name, 'player_global') === 0) {
			continue;
		}
		if (strpos($table_name, 'player') === 0) {
			$table_type_name = get_type_name($table['name']);
			$code .= "	load{$table_type_name}(db, 0, sid, wgn, procId)\n";
		}
	}
	$code .= "}\n\n";

	foreach ($tables as $table) {
		$table_name = $table['name'];
		$table_type_name = get_type_name($table['name']);
		$primary_key_type_name = get_type_name($table['pri']['name']);
		$primary_key_go_type = get_go_type($table['pri']['type']);
		$primary_key_c_type = get_c_type($table['pri']['type']);

		if (strpos($table_name, 'global') === 0) {
			$code .= "func load{$table_type_name}(db *mysql.Connection) {\n";
			$code .= "	log.Infof(\"load {$table_name}\")\n";
			$code .= "	sql := \"SELECT * FROM `{$table_name}`\"\n";
			if ($table['pri']['name'] == 'id') {
				$code .= "	if !initRowId(db, &(g_RowIds.{$table_type_name}), \"SELECT MAX(`{$table['pri']['name']}`) FROM `{$table_name}`\") {\n";
				$code .= "		return\n";
				$code .= "	}\n";
			}
			// 腾讯global_tb_xxd_onlinecnt表不加载到内存中
			if (strpos($table_name, 'global_tb_') === 0) {
				$code .= "	func(s string){}(sql)\n"; // 避免编译报错sql变量未使用
				$code .= "}\n\n";
				continue;
			}

			$code .= "	res, err := db.ExecuteFetch([]byte(sql), -1)\n";
			$code .= "	if err != nil {\n";
			$code .= "		panic(err)\n";
			$code .= "	}\n";
			foreach ($table['cols'] as $column) {
				$code .= "	i".get_type_name($column['name'])." := res.Map(\"".$column['name']."\")\n";
			}
			$code .= "	for _, row := range res.Rows {\n";
			$code .= "		crow := C.New_{$table_type_name}()\n";
			foreach ($table['cols'] as $column) {
				$column_c_type = get_c_type($column['type']);
				$column_type_name = get_type_name($column['name']);
				if ($column_c_type == "char*") {
					$code .= "		row_{$column_type_name} := row.".get_go_type2($column['type'])."(i".get_type_name($column['name']).")\n";
					$code .= "		crow.{$column_type_name} = C.CString(string(row_{$column_type_name}))\n";
					$code .= "		crow.{$column_type_name}_len = C.int(len(row_{$column_type_name}))\n";
				} else {
					$code .= "		crow.{$column_type_name} = C.{$column_c_type}(row.".get_go_type2($column['type'])."(i".get_type_name($column['name'])."))\n";
				}
			}
			$code .= "		crow.next = C.g_GlobalTables.{$table_type_name}\n";
			$code .= "		C.g_GlobalTables.{$table_type_name} = crow\n";
			$code .= "		if g_Hooks.{$table_type_name} != nil {\n";
			$code .= "			g_Hooks.{$table_type_name}.Load(&{$table_type_name}Row{row: crow})\n";
			$code .= "		}\n";			
			$code .= "	}\n";
			$code .= "}\n\n";
		}

		if (strpos($table_name, 'player') === 0) {
			$priKey = $table_name == 'player' ? 'id': 'pid';

			$isHDTable = (strpos($table_name, 'player_global_') === 0); // 是否是互动表

			$code .= "func load{$table_type_name}(db *mysql.Connection, playerId int64, sid, wgn, procId int) {\n";
			$code .= "	if playerId == 0 {\n";
			$code .= "		log.Infof(\"load {$table_name}\")\n";
			$code .= "	}\n";
			$code .= "	sql := \"SELECT * FROM `{$table_name}`\"\n";
			$code .= "	if playerId == 0 {\n";

			if ($table['pri']['name'] == 'id') {
				$where = "WHERE (".$priKey." >> 32) = \" + itoa(sid) + \"";
				
				if($isHDTable) {
					// 互动服自增长数据只有一个进程来生成，所以不按sid来区分
					$where = "";
				}

				$code .= "		if !initRowId(db, &(g_RowIds.{$table_type_name}), \"SELECT MAX(`{$table['pri']['name']}`) FROM `{$table_name}` {$where}\") {\n";
				$code .= "			return\n";
				$code .= "		}\n";
			}

			$whereBySid = "(".$priKey." >> 32) = \" + itoa(sid) + \" AND ";
			// 互动玩家表由互动服加载数据，所以不需要按sid加载
			if ($isHDTable) {
				$whereBySid = "";
			}

			$code .= "		sql += \" WHERE ".$whereBySid." ".$priKey." % \" + itoa(wgn) + \" = \" + itoa(procId)\n";
			$code .= "	} else {\n";
			$code .= "		sql += \" WHERE ".$priKey." = \" + playerIdToStr(playerId)\n";
			$code .= "	}\n";
			$code .= "	res, err := db.ExecuteFetch([]byte(sql), -1)\n";
			$code .= "	if err != nil {\n";
			$code .= "		panic(err)\n";
			$code .= "	}\n";
			foreach ($table['cols'] as $column) {
				$code .= "	i".get_type_name($column['name'])." := res.Map(\"".$column['name']."\")\n";
			}
			$code .= "	for _, row := range res.Rows {\n";
			$code .= "		crow := C.New_{$table_type_name}()\n";
			foreach ($table['cols'] as $column) {
				$column_c_type = get_c_type($column['type']);
				$column_type_name = get_type_name($column['name']);
				if ($column_c_type == "char*") {
					$code .= "		row_{$column_type_name} := row.".get_go_type2($column['type'])."(i".get_type_name($column['name']).")\n";
					$code .= "		crow.{$column_type_name} = C.CString(string(row_{$column_type_name}))\n";
					$code .= "		crow.{$column_type_name}_len = C.int(len(row_{$column_type_name}))\n";
				} else {
					$code .= "		crow.{$column_type_name} = C.{$column_c_type}(row.".get_go_type2($column['type'])."(i".get_type_name($column['name'])."))\n";
				}
			}
			if ($table_name == 'player') {
				$code .= "		playerDb := GetPlayerTables(int64(crow.Id))\n";
			} else {
				$code .= "		playerDb := GetPlayerTables(int64(crow.Pid))\n";
			}
			if (is_one_to_one_player_table($table)) {
				$code .= "		playerDb.{$table_type_name} = crow\n";
			} else {
				$code .= "		crow.next =  playerDb.{$table_type_name}\n";
				$code .= "		playerDb.{$table_type_name} = crow\n";
			}
			$code .= "		if g_Hooks.{$table_type_name} != nil {\n";
			$code .= "			g_Hooks.{$table_type_name}.Load(&{$table_type_name}Row{row: crow})\n";
			$code .= "		}\n";
			$code .= "	}\n";
			$code .= "}\n\n";
		}
	}

	save_code($output_dir."/zd_load.go", $code);
	//system("gofmt -w=true {$output_dir}/zd_load.go");
}
?>