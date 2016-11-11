<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
 <head>
  <title>神仙道IOS版开发辅助脚本</title>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <script src="http://cdn.staticfile.org/jquery/2.1.1-rc2/jquery.min.js" type="text/javascript"></script>
  <script type="text/javascript">
  	function check_md5(a){
		$('#result').load("api.php?sf="+a+"&sr="+Math.random());
  	}

  
  </script>
<style>
BODY{line-height: 16pt;}
TABLE{background-color: #bbb;}
TABLE TD .title{font-size: 18px;}
TABLE TD{background-color: #FFF;padding: 5px;font-size: 14px;}
TABLE Th{background-color: #eee;padding: 5px;width:100px;}
.txt_gray{color: #aaa;}
.block{display: inline;width:20px;}
</style>
</head>
<body>
<h2>仙侠道API相关功能</h2><?php 
require_once("status.php");
echo GetServerStatus();
?>
<table cellpadding="1" cellspacing="1" bgcolor="#ccc">
	<tr>
		<th>仙侠道功能链接：</th>
		<td class="title"><a href="api.php" target="_blank">游戏API速查手册</a></td>
		<td class="title"><a href="api_code.php" target="_blank">游戏API代码生成</a></td>
		<td class="title"><a href="http://api.pinidea.com:39399/xxd_platform.txt" target="_blank">平台API速查手册</a></td>
		<td class="title"><a href="request/report.php" target="_blank">接口性能报表</a></td>
		<td class="title"><a href="api.php?action=cc" target="_blank">获取客户端数据</a></td>
		<td class="title"><a href="http://debug.vvv.io:8081" target="_blank">传送到资源管理页面</a></td>
	</tr>
	</table>
	<h3>查找dat文件中的文件</h3>
	<table>
	<tr><td>文件名</td><td><input id="filename1" type="text" name="filter"><input type=submit onclick="check_md5($('#filename1').val());" name="filter"></td></tr>
	
	</table>
    <p id=
    "result"></p>
   
</body>
