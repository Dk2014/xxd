<?php


function GetServerStatus() {
	$root_dir = dirname(__FILE__).'/../../';
	$building = "{$root_dir}/zbuilding"; 
	$logfile = "{$root_dir}/zbuildlog";
	if (file_exists($building)) {
		return HtmlStatus("building");
	}

	if (!file_exists($logfile)) {
		return HtmlStatus("Never");
	}


	$info = shell_exec("tail -n 1 {$logfile}");
	if (strpos($info, "Successed in") === false) {
		return HtmlStatus("Failed");
	}

	return HtmlStatus($info);
}

function HtmlStatus($info) {
	$color = "#00CC66";
	if ($info=="Failed")
		$color = "#CC4D33";
	else if ($info=="Never")
		$color = "gray";
	else if ($info=="building")
		$color = "#FFCC33";

	$file = dirname(__FILE__).'/last-hg';
	$hg = (file_exists($file)) ? file_get_contents($file) : "";
	return "
	<div style='float: right;width:280px;border:1px solid;'>
	<p style='background-color: {$color};'>Compile: {$info}</p>
	<p>服务端最近更新于:</p>
	<p> {$hg} </p>
	</div>
	";
}

?>