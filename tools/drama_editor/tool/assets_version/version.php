<?php
date_default_timezone_set("Asia/Shanghai");

$is_debug = true;
$is_tw    = false;

if (isset($argc) && isset($argv) && $argc > 1) {
	if ($argv[1] == 1) $is_debug = false;
	if ($argc > 2 && $argv[2] == "tw") $is_tw = true;
}

$path = dirname(__FILE__)."/";

$dev_path    = dirname(dirname($path))."/";
$client_path = $dev_path.($is_tw ? "client_tw/" : "client_mobile/");
$assets_path = $client_path."assets/";

$version = "";
$version_list = array();

if (! $is_debug) {
	@unlink($client_path."Main(debug).swf");
	@unlink($client_path."WarPlayer(debug).swf");
	@unlink($client_path."RolePlayer(debug).swf");
	@unlink($client_path."SkillPlayer(debug).swf");
	@unlink($client_path."DramaPlayer(debug).swf");
}

$main_url = $client_path."Main.swf";
if (file_exists($main_url))
	$version_list["Main.swf"] = date("YmdHi", filemtime($main_url))."|".filesize($main_url);

$main_url = $client_path."Main(debug).swf";
if (file_exists($main_url))
	$version_list["Main(debug).swf"] = date("YmdHi", filemtime($main_url))."|".filesize($main_url);

$preface_url = $client_path."Preface.swf";
if (file_exists($preface_url))
	$version_list["Preface.swf"] = date("YmdHi", filemtime($preface_url))."|".filesize($preface_url);

$s = recursive(glob($assets_path."*"), "");
print "size: ".number_format(($s / 1024 / 1024), 2)."m, ";

foreach ($version_list as $sign => $time) {
	if ($version != "") $version .= "\n";
	$version .= $sign."|".$time;
}

#print $version;

file_put_contents($client_path."assets.txt", gzcompress($version));

#file_put_contents($client_path."assets.txt", $version);
#file_put_contents($client_path."assets_gzip.txt", gzencode($version));
#file_put_contents($client_path."assets_zlib.txt", gzcompress($version));
#file_put_contents($client_path."assets_deflate.txt", gzdeflate($version));

print "End.\n";

# 循环拷贝
function recursive ($urls, $prev) {
	global $is_debug;
	global $dev_path, $assets_path, $version, $version_date, $version_list;
	
	#var_export($urls);
	
	$s = 0;
	
	foreach ($urls as $url) {
		if (is_dir($url)) {
			$list = explode("/", $url);
			$first = substr($list[count($list) - 1], 0, 1);
			$s += recursive(glob($url.'/*'), $prev."/".$list[count($list) - 1]);
		}
		else {
			$dest = $assets_path.str_replace($dev_path, '', $url);
			
			# 发布时，删除debug版的swf
			if (preg_match("/\(debug\)\.swf/", $dest) && ! $is_debug) {
				@unlink($url);
				continue;
			}
			
			# 文件最后修改时间
			$time = filemtime($url);
			
			$sign = $prev."/".basename($url);
			$time = date("YmdHi", $time);
			
			$version_list[$sign] = $time."|".filesize($url);
			
			$s = $s + filesize($url);
		}
	}
	
	return $s;
}
?>