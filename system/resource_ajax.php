<?php

$tag = $_GET['tag'];
$path = $_GET['path'];

if (!isset($tag) || !isset($path)) {
	exit;
}

require_once("lib/global.php");

$config = getResourceConfig($path, $tag);
echo json_encode($config);


function getResourceConfig($res_path, $tag) {
	if(!file_exists($res_path)) {
		exit;
	}

	// 找$tag对应的资源目录
	$realPath = findPath($res_path, $tag);

	// 解析资源配置文件
	$cfg = $realPath.'/right_bottom.txt';
	if (!file_exists($cfg)) {
		exit;
	}

	$cfg_txt = file_get_contents($cfg);

	$config = array('realPath'=>$realPath.'/right_bottom.png');
	$lastToken = '';

	$items = array('width', 'height', 'frames', 'interval');
	foreach (explode("\r\n", $cfg_txt) as $line) {
		preg_match('/\[(\w+)\]/', $line, $matchs);
		if (count($matchs)> 0 && in_array($matchs[1], $items)) {
			$lastToken = $matchs[1];
			continue;
		}

		if (!empty($lastToken)) {
			$config[$lastToken] = $line;
		}

		$lastToken = '';
	}

	return $config;
}

function findPath($res_path, $tag) {
	global $db;
	$rows = db_query($db, "select * from enemy_role");
	$map_sign_to_name = array();
	foreach ($rows as $key => $value) {
		$map_sign_to_name[$value['sign']][] = $value['name'];
	}

	$fileList = glob("{$res_path}/*");
	foreach ($map_sign_to_name[$tag] as $key => $name) {
		foreach ($fileList as $key => $file) {
			if (strpos($file, $name) === false)
				continue;

			return $file.'/standby';
		}
	}
}

?>