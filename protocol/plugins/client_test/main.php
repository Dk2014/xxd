<?php

// 生成客户端测试代码

function plugin_main()
{
	global $argv;

	$server_dir = $argv[2];
	$modules = parse_all($server_dir);

	$output_dir = 
	/*mobile-dev*/dirname(
		/*protocol*/dirname(
			/*plugins*/dirname(
				/*server_code*/dirname(__FILE__)
			)
		)
	)."/server/src/client_test/";

	if (!is_dir($output_dir))
		mkdir($output_dir);

	gen_main_code($output_dir, $server_dir, $modules);
}

function get_type_name($name) {
	$items = explode('_', $name);
	$type_name = '';
	for ($i = 0; $i < count($items); $i ++) {
		$type_name .= ucfirst($items[$i]);
	}
	return $type_name;
}

function get_params_list($MName,$AName, $columns) {
	$list = array();
	foreach ($columns as $column) {
		$code = "";
		switch ($column->Type) {
			case ProtoToken::IBOOL:
				$code = "bool";
				break;
			case ProtoToken::INT8:
				$code = "int8";
				break;
			case ProtoToken::INT16:
				$code = "int16";
				break;
			case ProtoToken::INT32:
				$code = "int32";
				break;
			case ProtoToken::INT64:
				$code = "int64";
				break;
			case ProtoToken::UINT8:
				$code = "uint8";
				break;
			case ProtoToken::UINT16:
				$code = "uint16";
				break;
			case ProtoToken::UINT32:
				$code = "uint32";
				break;
			case ProtoToken::UINT64:
				$code = "uint64";
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::TEXT16:
				$code = "string";
				break;
			case ProtoToken::BINARY8:
			case ProtoToken::BINARY16:
				$code = "[]byte";
				break;
			case ProtoToken::LIST8:
			case ProtoToken::LIST16:
				$name = get_type_name($column->Name);
				$M = get_type_name($MName);
				$code = "[]{$MName}_api.".$AName."_In_".$name;
				break;
			default:
				if ($column->RefType != '') {
					$code = $MName.'_api.'.get_type_name($column->RefType);
				}
				break;
		}
		$var = $column->Name;
		$list[$var] = $code;
	}
	
	return $list;
}

function gen_main_code($output_dir, $server_dir, $modules) {
	$code  = "package client_test\n\n";
	$code .= "import \"core/net\"\n";
	$code .= "import \"core/fail\"\n";
	$code .= "import \"fmt\"\n";
	$code .= "import \"game_server/api/protocol\"\n";
	
	$code .= "import (\n";
	foreach ($modules as $module) {
		$has_request = false;

		foreach ($module->Actions as $action) {
			if (is_array($action->InColumns)) {
				$has_request = true;
				break;
			}
		}

		if ($has_request) {
			$code .= "	\"game_server/api/protocol/{$module->Name}_api\"\n";
		}
	}

	$code .= ")\n\n";

	// 定义数据结构
	$code .= "type Client struct {\n";
	$code .= "	Name string\n\n\n";
	$code .= "	closed bool\n";
	$code .= "	conn *net.Conn\n\n\n";

	$OutHandlersCode = "";
	$OutInnerStructCode ="";
	$OutInnerFuncsCode = "";

	foreach ($modules as $module) {
		$MName = get_type_name($module->Name);

		$OutHandlersCode .= "{$module->Name}_api.SetOutHandler(Out{$MName}{})\n";
		$OutInnerStructCode .= "type Out{$MName} struct{}\n";

		foreach ($module->Actions as $action) {
			$AName = get_type_name($action->Name);

			$code .= "Out{$MName}_{$AName} func(*{$module->Name}_api.{$AName}_Out)\n";

			$OutInnerFuncsCode.= "
				func (o Out{$MName}) {$AName}(session *net.Session, out *{$module->Name}_api.{$AName}_Out) {\n
				if session.State.(*Client).Out{$MName}_{$AName} != nil {\n
					session.State.(*Client).Out{$MName}_{$AName}(out)\n
		}\n
		}\n";
		}
	}

	$code .= "}\n\n";

	$code .="
		func init() {\n
				{$OutHandlersCode}
}\n
";

$code .= "
	func NewClient(ip string) *Client {\n
	c := &Client{}\n
	var err error
	c.conn, err = net.Dial(\"tcp\", ip, net.PacketN(8, net.LittleEndian))\n
	fail.When(err!=nil, err)\n
	\n
	c.recv()\n
	return c\n
}\n
func (c *Client) Close() {\n\n
	c.closed = true
	c.conn.Close()
}
\n
func (c *Client) RecordResponse(rsp interface{}) {}\n\n
func (c *Client) recv() {\n
go func(){\n
defer func() {
	if err := recover(); err != nil {
		fmt.Printf(\"Client %s recv() panic. %v\\n\", c.Name, err)
	}
}()
s := &net.Session{State: c}\n
for {\n
	raw, err := c.conn.Read()\n
	if err != nil {
		if c.closed {
			return
		} else {
			panic(err)
		}
	}
	protocol.DecodeOut(raw).Process(s)\n
}\n
}()\n
}\n

{$OutInnerStructCode}\n
{$OutInnerFuncsCode}\n
";

// 接口请求
foreach ($modules as $module) {
	foreach ($module->Actions as $action) {
		$MName = get_type_name($module->Name);
		$AName = get_type_name($action->Name);

		$params = (is_array($action->InColumns)) ? get_params_list($module->Name,$AName, $action->InColumns) : array();

		$p_str = "";
		$struct_str = "";
		foreach ($params as $var => $type) {
			$s = empty($p_str) ? "" : ",";
			$p_str .= $s."{$var} {$type}";

			$value = ($type=="string") ? "[]byte({$var})" : "{$var}";
			$struct_str .= get_type_name($var).':'.$value.",\n";
		}

		if (isset($action->InColumns)) {
			$code .= "
				func (c *Client) {$MName}_{$AName}({$p_str}) {\n
				in := &{$module->Name}_api.{$AName}_In{\n
		{$struct_str}\n
		}\n
		\n
		buff := net.NewBuffer(make([]byte, in.ByteSize()))\n
		in.Encode(buff)\n
		c.conn.Write(buff.Get())\n
		}\n";
		}
	}
}


save_code($output_dir."/client.go", $code);

system("gofmt -w=true {$output_dir}/client.go");
}
?>
