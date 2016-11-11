<?php

//===============================================
// 服务端常用数据映射代码生成器
//===============================================
require_once('parse_config.php');
require_once(dirname(__FILE__)."/../../../system/pages/battle_item_const.php");

function plugin_main ()
{	
	$db = db_connect();
	$cur_dir = dirname(__FILE__);

	$filelist = glob("{$cur_dir}/data/*.php");
	$c = count($filelist);
	$filename = "";
	for ($i=0; $i < $c; $i++) { 
		$filename = basename($filelist[$i], ".php");
		require_once($filelist[$i]);
		if ($filename == 'enemy_skill' ||
			$filename == 'battle_item' ||
			$filename == 'enemy_passive_skill' ||
			$filename == 'enemy_rand_skill') {
			$generate($db);
		} else {
			generate_dat($db, $filename, $filename.'_dat');
		}
	}
}

function generate_dat(&$db, $filename, $package_name) {
	$currdir = dirname(__FILE__);
	$codefn = "GetCode_{$filename}";
	$importfn = "GetImport_{$filename}";

	$import_code = generate_import_code($importfn($db));
	$body_code = $codefn($db);

  if ($import_code != "") {
    $import_code = <<<IMPC

import (
	{$import_code}
)
IMPC;
  }

	require_once($currdir."/../../config/{$filename}.php");
	$parse = new ParseConfig();
	$config_code = $parse->Parse($config);

######### generate start #########

	$code = <<<PACKAGE
package {$package_name}
{$import_code}

{$config_code}

{$body_code}

PACKAGE;

######### generate end #########
	
	$outdir = $currdir."/../../../server/src/game_server/dat";
	$out = $outdir."/{$package_name}";
	if (!is_dir($out)) {
		mkdir($out);
	}

	save_code($out."/{$filename}_auto.go", $code);
}

function generate_import_code($use_package) {
	if (!is_array($use_package)) {
		return "";
	}

	$code = "";
	foreach($use_package as $package) {
		$code .= "\"$package\"\n	";
	}
	return $code;
}
?>
