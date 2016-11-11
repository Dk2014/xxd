<?php

header ( "Content-type:text/html;charset=utf-8" );
date_default_timezone_set('PRC');
$dir = dirname(__FILE__).'/../../../server/bin/prof/*.csv';

foreach (glob($dir) as $v) {
	$v = basename($v);
	$filename[$v] = @date("Y-m-d H:i:s",substr($v, strrpos($v, "-")+1,-4)).substr($v,strrpos($v, "."));
}
krsort($filename, SORT_STRING);
if (!empty($filename)) {
	echo "<table>";
	foreach ($filename as $k=>$v) {
		echo "<tr><td><a style='color:black;' href='analysisshow.php?filename={$k}' target='showframe'>{$v}</a><br></td></tr>";
	}
	echo "</table>";
}
?>