<?php
require_once 'lib/global.php';

session_start();

if (isset($_SESSION['user_id'])) {
	header('Location:index.php');
	exit;
}

if (isset($_POST['login'])) {
	$admin = $admins[$_POST['username']];

	if ($admin && $admin['password'] == md5($_POST['password'])) {
		$_SESSION['user_id'] = $admin['id'];
		$_SESSION['username'] = $admin['username'];
		header('Location:index.php');
		exit;
	} else {
		echo "<script>alert('用户名或者密码错误');</script>";
	}
}

?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<title>《猎命师》游戏数据录入工具</title>
	<script src="http://cdn.staticfile.org/jquery/1.7.2/jquery.min.js"></script>
	<style type="text/css">
html, body { margin:0; padding:0; }
body { font-size:13px; font-family:"微软雅黑","Lucida Grande","Lucida Sans Unicode",Calibri,Arial,Helvetica,Sans,FreeSans,Jamrul,Garuda,Kalimati; }

#login_form_wrap { padding:0 16px; }
#login_form { width:320px; border:1px solid #e5e5e5; background:whiteSmoke; }
#login_form form { margin:0; padding:0; }
#login_form .label { color:#333; margin:16px 0px 8px 0px; font-size:14px; }
#login_form .textbox { border:1px solid #e5e5e5; background:#FFF; padding:2px; }
#login_form .textbox input { width:100%; font-size:28px; border:0; background:#FFF; outline:0; }
#login_form #button_bar { text-align:right; padding:16px 0; }
#login_button { font-size:14px; width:60px; height:34px; }
#login_form_title { text-align:center; }
#login_error_msg { padding:16px; border:1px solid #CCC; background:#FFF0F0; margin-bottom:16px; }
	</style>
	<script type="text/javascript">
$(window).resize(function(){
	$("#login_form").css({
		position:"absolute",
		left: ($(window).width() - $("#login_form").outerWidth())/2+"px",
		top: ($(window).height() - $("#login_form").outerHeight())/2+"px"
	});  
});

$(function(){
	$(window).resize();
});
	</script>
</head>
<body>
	<div id="login_form">
		<div id="login_form_wrap">
			<form action="<?php echo $_SERVER['PHP_SELF']; ?>" method="post">
			<div class="label">帐号</div>
			<div class="textbox"><input type="text" name="username" /></div>
			<div class="label">密码</div>
			<div class="textbox"><input type="password" name="password" /></div>
			<div id="button_bar"><input type="submit" value="登录" name="login" id="login_button" /></div>
			</form>					
		</div>
	</div>
</body>
</html>