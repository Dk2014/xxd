<?php

function GetImport_announcement() {
	return array(
		"strconv",
		"bytes",
	);
}



function GetCode_announcement($db) {
	$announcements = db_query($db, "SELECT * FROM announcement");
	$code = '';

	//为数据库 announcement 表里面的每一条记录生成一个对应的结构体
	foreach ($announcements as $key => $val) {
		//结构体定义
		$code .= "// {$val['name']}\n";
		$announe_name = "Announce".$val['sign'];
		$code .= "type $announe_name struct {\n";
		//0 -> field name, 1 -> comment, 2 -> type
		$parameters = getAnnounceParas($val['parameters']);
		foreach ($parameters as $key => $p) {
			$code .= "	{$p[0]} {$p[2]} // {$p[1]}\n";
		}

		$code .= "}\n";
		$code .= "\n";

		//GetAnnouncementTplId 方法定义
		$code .= "func (this $announe_name) GetAnnouncementTplId() int32 {\n";
		$code .= "	return int32({$val['id']})\n";
		$code .= "}\n";

		$code .= "\n";

		//GetDuration 方法定义
		$code .= "func (this $announe_name) GetDuration() int64 {\n";
		$code .= "	return int64({$val['duration']})\n";
		$code .= "}\n";

		$code .= "\n";

		//GetParameters 方法定义
		$code .= "func (this $announe_name) GetParameters() string {\n";
		$code .= "	b := new(bytes.Buffer)\n";
		$code .= "	b.WriteString(\"[\")\n";
		$l = count($parameters);
		foreach ($parameters as $key => $p) {
			$code .= "	b.WriteString(strconv.Quote(this.{$p[0]}))\n";
			if ($l > $key+1){
				$code .= "	b.WriteString(\",\")\n";
			}
		}
		$code .= "	b.WriteString(\"]\")\n";
		$code .= "	return string(b.Bytes())\n";
		$code .= "}\n";

		$code .= "\n";

	}


	//$code .= generate_test_code($db, $mails);

	return $code;
}


//解析 parameters 字段，返回 array(array(参数名、类型、注释), ...)
function getAnnounceParas($str) {
	if (empty($str)) { 
		return array();
	};

	$ta = explode(";", $str);
	$parameters = array();
	$module = '';
	foreach ($ta as $key => $p) {
		if (empty($p)) {
			continue;
		}

		$t = explode(",", trim($p));
		$type = 'string';
		
		array_push($parameters, array(formatName($t[0]), isset($t[1]) ? $t[1] : '', $type, $module));
	}
	return $parameters;
}



//获取驼峰风格的变量名字符串
function formatName( $name ) {
	$items = explode( '_', $name );
	$x = '';
	for ( $i = 0; $i < count( $items ); $i ++ ) {
		$y = $items[$i];
		$y[0] = strtoupper( $y[0] );
		$x .= $y;
	}
	return $x;
}
?>
