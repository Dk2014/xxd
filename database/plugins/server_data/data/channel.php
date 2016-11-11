<?php

function GetImport_channel($db) {
	$messageRows = db_query($db, "SELECT * FROM world_channel_message limit 1");
	if(count($messageRows)>0) {
		return array(
			"strconv",
			"bytes",
		);
	}
	return array();
}

function GetCode_channel($db) {
	return GetCode_message_tpl($db);
}

function GetCode_message_tpl($db) {
	$messageRows = db_query($db, "SELECT * FROM world_channel_message");
	$code = '';

	foreach ($messageRows as $key => $val) {
		//结构体定义
		$struct_name = "Message".$val['sign'];
		$code .= "type $struct_name struct {\n";
		//0 -> field name, 1 -> comment, 2 -> type
		$parameters = getMessageTplParas($val['parameters']);
		foreach ($parameters as $key => $p) {
			$code .= "	{$p[0]} {$p[2]} // {$p[1]}\n";
		}
		$code .= "}\n";
		$code .= "\n";

		//GetTplId 方法定义
		$code .= "func (this $struct_name) GetTplId() int16 {\n";
		$code .= "	return {$val['id']}\n";
		$code .= "}\n";

		$code .= "\n";


		//GetParameters 方法定义
		$code .= "func (this $struct_name) GetParameters() []byte {\n";
		$code .= "	b := new(bytes.Buffer)\n";
		$code .= "	b.WriteString(\"[\")\n";
		$l = count($parameters);
		foreach ($parameters as $key => $p) {
			$code .= "	b.WriteString(strconv.Quote(this.{$p[0]}.ToString()))\n";
			if ($l > $key+1){
				$code .= "	b.WriteString(\",\")\n";
			}
		}
		$code .= "	b.WriteString(\"]\")\n";
		$code .= "	return b.Bytes()\n";
		$code .= "}\n";

		$code .= "\n";

	}


	return $code;
}


//解析 parameters 字段，返回 array(array(参数名、类型、注释), ...)
function getMessageTplParas($str) {
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

		$type = "";
		//[0] type; [1] field name; [2] comment
		$param_desc = explode(",", trim($p));
		switch (strtoupper($param_desc[0])) {
		case 'STRING':
			$type = "ParamString";
			break;
		case 'ITEM':
			$type = "ParamItem";
			break;
		case 'PLAYER':
			$type = "ParamPlayer";
			break;
		case 'CLIQUE':
			$type = "ParamClique";
			break;
		case 'BOAT':
			$type = "ParamCliqueBoat";
			break;
		//如需更多类型在此定义
		default:
			die("unknow message param type {$param_desc[0]}");
		}
		
		array_push($parameters, array(_formatName($param_desc[1]), isset($param_desc[2]) ? $param_desc[2] : '', $type, $module));
	}
	return $parameters;
}



//获取驼峰风格的变量名字符串
//TODO 把这东西移动到一个全局的lib里面
function _formatName( $name ) {
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
