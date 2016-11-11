<?php
date_default_timezone_set('Asia/Shanghai');

require_once 'lib/global.php';

// 判断是否登录
check_login();

// 参数解析
$look = isset($_GET['look']) ? $_GET['look'] : '';
$date = isset($_GET['date']) ? $_GET['date'] : date("Y-m-d");
$srv = isset($_GET['srv']) ? $_GET['srv'] : '';
$prefix = isset($_GET['srv']) ? $_GET['prefix'] : '';

?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
	<title> 错误日志 </title>
	<script src="http://cdn.staticfile.org/jquery/1.7.2/jquery.min.js"></script>
	<script src="jquery.scrolltotop.js"></script>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<style type="text/css" media="all">
	@import url("style.css");
#log_list p { padding:20px; border-bottom: solid 1px #ccc; }
	</style>
</head>
<body>

<script type="text/javascript">
	jQuery(function($){
   		$.scrolltotop({
   			controlHTML: '<div style="right: 25px; position: fixed; bottom: 25px;"><a href="javascript:;"><b>回到顶部↑</b></a></div>',
   			offsety: 0,
   		});
	});
</script>

	<div id="log_list">
	<p id="top"><a href="<?php echo "log.php?srv=$srv&prefix=$prefix&date=$date"; ?>">[全部]</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="<?php echo "log.php?look=error&srv=$srv&prefix=$prefix&date=$date"; ?>">[只看错误]</a>&nbsp;&nbsp;&nbsp;&nbsp;<a href="<?php echo "log.php?look=last&srv=$srv&prefix=$prefix&date=$date"; ?>">[最后一条]</a></p>
<?php
// 指定日志文件位置
$exist_file = true;
$log_name = '../server/'.$srv.'/'. $prefix . $date . '.log';
if (!file_exists($log_name)) {	
	$log_name = '../server/bin/'.$srv.'/' . $prefix. $date . '.log';

	if (!file_exists($log_name))
		$exist_file = false;
}

if ($exist_file) {
	// 分析日志文件	
	$contents = array(); // 存放指定类型的日志内容
	$contents[0] = '';
	$k = 0;

	if ($look) {
		if ($fh = fopen($log_name, 'r')) {
			while (!feof($fh)) {
				$line_content = fgets($fh);
				if (!$line_content)
					continue;

				$first_line_every_info = strstr($line_content, "[I]");
				if ($first_line_every_info)
					continue;

				$first_line_every_err = strstr($line_content, "[E]");
				if ($first_line_every_err) {
					$k ++;
					$contents[$k] = '';					
				}

				$contents[$k] .= $line_content;	
			}
		}
	} else {
		if ($fh = fopen($log_name, 'r')) {
			while (!feof($fh)) {
				$line_content = fgets($fh);
				if (!$line_content)
					continue;

				$first_line_every_log = preg_match('/\[[(IE)]\]/', $line_content);

				if ($first_line_every_log != 0) {
					$k ++;
					$contents[$k] = '';
				}

				$contents[$k] .= $line_content;
			}
		}
	}

	$num = count($contents);
	if ($num <= 0) {
		echo "<p>无相应日志信息</p>";
	} else {
		$last_index = $num - 1;

		// 按要求输出
		if ($look == 'last') {
				echo "<p>";
				echo str_replace("\n", "</br>", str_replace("\t", "&nbsp;&nbsp;&nbsp;&nbsp;", str_replace(" ", "&nbsp;", $contents[$last_index])));
				echo "</p>";
		} else {
			for ($i = $last_index; $i >= 0; $i--) {
				echo "<p id='$i'>";
				$pre = $i == $last_index ? 0 : $i + 1;
				$next = $i == 0 ? "top" : $i - 1;
				echo "<a href='#{$next}'>↓</a>	<a href='#{$pre}'>↑</a> <br>";
				echo preg_replace('/\[(([^\[\]]*\.go):(\d+))\]/m', '[<a href="code.php?f=\2&n=\3#\3" target="_blank">\1</a>]', str_replace("\n", "</br>", str_replace("\t", "&nbsp;&nbsp;&nbsp;&nbsp;", str_replace(" ", "&nbsp;", $contents[$i]))));
				echo "</p>";
			}
		}
	}
} else {
	echo "<p>无日志文件</p>";
}

?>
	</div>
</body>
</html>
