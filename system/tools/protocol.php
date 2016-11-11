<?php

$mod = $_GET['p']; // 要处理的模块

define('PPATH', '../../protocol/servers/game_server/'); // protocol文件路劲
if(isset($mod) && !empty($mod)){
$dir = dir(PPATH);
while(($file = $dir->read())!== false){
	if ($file != '.' && $file !='..' && preg_match('/^\d+\_'.$mod.'\.pd$/i', $file)) {
		$real_file = PPATH.$file; // 此次要操作的协议文件真实路劲
		echo 'type '.ucfirst($mod).'API struct{}<br /><br />';
		$lines = file($real_file);
		foreach ($lines as $line) {
			$matchs = array();
			if (preg_match('/api\s+([^\s]+)\s*=\s*\d+\s*{/i', $line, $matchs)) {
				$api = $matchs[1];
				$parts = explode('_', $api);
				$api_func_name = '';
				foreach ($parts as $part) {
					$api_func_name .= ucfirst($part);
				}
				echo '<br />func (this '.ucfirst($mod).'API) '.$api_func_name.'(session *net.Session, in *'.$mod.'_api.'.$api_func_name.'_In) {<br /><br />}<br />';
			}
		}
	}
}

}

?>