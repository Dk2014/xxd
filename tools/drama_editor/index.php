<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>仙侠道剧情编辑器(手机版)</title>
    <script type="text/javascript" src="http://sxd.xdcdn.net/swfobject/swfobject.js"></script>
    <style type="text/css">	
	html, body { background:#000; padding:0; margin:0; width:100%; height:100%; text-align:center; overflow:hidden;}
	#webgame {
		position:absolute;
		left:0;
		top:0;
		width:100%;
		height:100%;
	}
    </style>
</head>
<body>
<div id="webgame"></div>
    <script type="text/javascript">
	var swf = './editor.swf?r=20140724_2';
	var flashVars = 'client_url=http://<?php echo $_SERVER['HTTP_HOST'].$_SERVER['REQUEST_URI']; ?>client_mobile/';
	
	var params = {
		flashvars : flashVars,
		quality   : "high",
		wmode     : "direct",
		menu      : "false",
		bgcolor   : "#ffffff",
		allowscriptaccess : "always",
	};
	
	var attributes = {
		id   : "lms",
		name : "lms"
	};
	
	swfobject.embedSWF(
		swf, "webgame", "100%", "100%", 
		"10.0.0", "swfobject/expressInstall.swf",
		false, params, attributes
	);
    </script>
</body>
</html>