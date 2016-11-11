<?php


function plugin_main()
{

	global $argv;

	$server_dir = $argv[2];

	$output_dir = 
	/*mobile-dev*/dirname(
		/*protocol*/dirname(
			/*plugins*/dirname(
				/*server_code*/dirname(__FILE__)
			)
		)
	).
	"/server/src/".$server_dir."/mdb";

	$modules = parse_all($server_dir);

	$code = "package mdb\n\n";

	$code .= "func (file *SyncFile) initProtocol() error {\n";
	$code .= "	file.BeginTrans()\n";
	foreach ($modules as $module) {
		$cols = array();
		$cols[] = $module->ID;
		$cols[] = $module->Name;

		foreach ($module->Actions as $action) {
			$cols[] = $action->ID;
			$cols[] = $action->Name;
		}

		$code .= "	file.WriteBytes([]byte(\"".implode($cols, ',')."\"))\n";
	}

	$code .= "	return file.EndTrans()\n";
	$code .= "}\n\n";

	save_code("{$output_dir}/zd_sync_protocol.go", $code);
}
?>