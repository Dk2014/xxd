<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
	<title><?php echo $_GET['f']; ?></title>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<style>
html, body { padding:0; margin:0; }
#editor { 
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
	font-size:13px;
}
	</style>
</head>
<body>
<?php
$file = $_GET['f'];
$target_line = $_GET['n'];
$code = file_get_contents($file);
?>
<div id="editor"><?php echo htmlspecialchars($code); ?></div>
<script src="ace-builds/src-noconflict/ace.js" type="text/javascript" charset="utf-8"></script>
<script>
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/github");
    editor.getSession().setMode("ace/mode/golang");
    var timer = setInterval("goto()", 500);
    function goto() {
        clearInterval(timer);
        editor.scrollToLine(<?php echo $target_line; ?>, true, false);
        editor.gotoLine(<?php echo $target_line; ?>);
    }
    editor.setReadOnly(true);
</script>
</body>
</html>
