<?
//把多个连续的数据库导出文件压缩为一个


$dir = "./servers/game_server";
$php_dbdump_header = "<?php /* dump file */\n";

$all_files = array();
foreach (glob($dir . "/*.php") as $file) {
			$name              = basename($file, '.php');
			$all_files[] = $file;
}
asort($all_files);

function checkFileIsDBDump($file) {
	global $php_dbdump_header;
	$f = fopen($file, 'r');
	$is_dump = (fgets($f) == $php_dbdump_header);
	fclose($f);
	return $is_dump;	
}

for($idx = count($all_files)-1; $idx>=0; $idx--) {
	if(!checkFileIsDBDump($all_files[$idx])){
		continue;
	}
	$name = basename($all_files[$idx], ".php");
	$idx--;
	while($idx>=0 && checkFileIsDBDump($all_files[$idx])) {
		$name2 = basename($all_files[$idx], '.php');
		echo "在 $name.php 之前存在重复的导出脚 $name2.php\n";
		file_put_contents($all_files[$idx],"<?php \n////冗余导出文件\n ?>");
		$idx--;
	}
}

?>
