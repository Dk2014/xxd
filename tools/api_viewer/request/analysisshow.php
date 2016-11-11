<html xmlns="http://www.w3.org/1999/xhtml">
 <head>
  <title>仙侠道移动版辅助脚本</title>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" /> 
 </head>


<style>
table {
	border: 2px solid black;
	width: 100%;
}

table th {
	border: 1px solid black;
	background: #CECFDA;
}

table td {
	padding: 6px;
	border: 1px solid black;
}
</style>

<?php

$dir = dirname(__FILE__).'/../../../server/bin/prof';
$file_dir = $dir . "/" . $_GET ['filename'];
$handle = fopen ( $file_dir, 'r' );
$str = file_get_contents ( $file_dir );
$file_encoding = mb_detect_encoding ( $str );
$row = 0;
$desc = array ();
while ( $data = fgetcsv ( $handle ) ) {
	if ($data ['0'] == 'name' && $data ['1'] == 'times') {
		continue;
	}
	$arr [$data ["0"]] = array (
			"name" => $data ["0"],
			"times" => $data ["1"],
			"avg" => $data ["2"],
			"min" => $data ["3"],
			"max" => $data ["4"],
			"total" => $data ["5"] 
	);
}
if ($_GET ["key"]) {
	foreach ( $arr as $v ) {
		$desc [$v ['name']] = $v [$_GET ["key"]];
	}
	arsort ( $desc );
}
if ($_GET ['filename']) {
	echo "<table>";
	echo "<tr>
			<th>name</th>
			<th><a " . (isset ( $_GET ["key"] ) && $_GET ["key"] == "times" ? "style='color:red;' " : "") . " style='text-decoration:none;' href='analysisshow.php?filename={$_GET['filename']}&key=times'>times</a></th>
			<th><a " . (isset ( $_GET ["key"] ) && $_GET ["key"] == "avg" ? "style='color:red;' " : "") . " style='text-decoration:none' href='analysisshow.php?filename={$_GET['filename']}&key=avg'>avg</a></th>
			<th><a " . (isset ( $_GET ["key"] ) && $_GET ["key"] == "min" ? "style='color:red;' " : " ") . " style='text-decoration:none' href='analysisshow.php?filename={$_GET['filename']}&key=min'>min</a></th>
			<th><a " . (isset ( $_GET ["key"] ) && $_GET ["key"] == "max" ? "style='color:red;' " : "") . " style='text-decoration:none' href='analysisshow.php?filename={$_GET['filename']}&key=max'>max</a></th>
			<th><a " . (isset ( $_GET ["key"] ) && $_GET ["key"] == "total" ? "style='color:red;' " : "") . " style='text-decoration:none' href='analysisshow.php?filename={$_GET['filename']}&key=total'>total</a></th>
		</tr>";
	if (! empty ( $desc )) {
		foreach ( $desc as $k => $v ) {
			echo "<tr>
					<td>" . $arr [$k] ['name'] . "</td>
					<td>" . formatenum ( $arr [$k] ['times'] ) . "</td>
					<td>" . formatenum ( $arr [$k] ['avg'] ) . "</td>
					<td>" . formatenum ( $arr [$k] ['min'] ) . "</td>
					<td>" . formatenum ( $arr [$k] ['max'] ) . "</td>
					<td>" . formatenum ( $arr [$k] ['total'] ) . "</td>
				</tr>";
		}
	} else {
		foreach ( $arr as $v ) {
			echo "<tr>
					<td>" . $v ['name'] . "</td>
					<td>" . formatenum ( $v ['times'] ) . "</td>
					<td>" . formatenum ( $v ['avg'] ) . "</td>
					<td>" . formatenum ( $v ['min'] ) . "</td>
					<td>" . formatenum ( $v ['max'] ) . "</td>
					<td>" . formatenum ( $v ['total'] ) . "</td>
				</tr>";
		}
	}
	echo "<table>";
}
function formatenum($num) {
	if (strlen ( $num ) > 3) {
		$formatenum = array ();
		for($i = 0; $i < strlen ( $num ) / 3; $i ++) {
			if (strlen ( $num ) - ($i + 1) * 3 >= 0) {
				$formatenum [$i] = substr ( $num, - 3 * ($i + 1), 3 );
			} else {
				$formatenum [$i] = substr ( $num, 0, strlen ( $num ) - intval(($i + 1) * 3) + 3 );
			}
		}
		krsort ( $formatenum );
		return implode ( ",", $formatenum );
	} else {
		return $num;
	}
}
?>
</html>