<?php
//===============================================
// 服务端通讯协议代码生成器
//===============================================

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
	).
	"/server/src/".$server_dir."/api/protocol";

	if (!is_dir($output_dir))
		mkdir($output_dir);

	gen_main_code($output_dir, $server_dir, $modules);
	
	foreach ($modules as $module) {
		$output_dir2 = $output_dir."/{$module->Name}_api";
		gen_module_code($output_dir2, $server_dir, $module);
		remove_old_files($output_dir2."/");
	}

	remove_old_files($output_dir."/");
}

function get_type_name($name) {
	$items = explode('_', $name);
	$type_name = '';
	for ($i = 0; $i < count($items); $i ++) {
		$type_name .= ucfirst($items[$i]);
	}
	return $type_name;
}

function gen_main_code($output_dir, $server_dir, $modules) {
	$code  = "package protocol\n\n";
	
	$code .= "import \"core/net\"\n";
	$code .= "import \"core/fail\"\n";
	
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
			$code .= "	\"{$server_dir}/api/protocol/{$module->Name}_api\"\n";
		}
	}
	$code .= ")\n\n";

	$code .=<<<CODE
	
	var TheDebugConfig DebugConfig

	type DebugConfig interface {
		IsEnableDebugAPI() bool
	}

	type ProtoError struct {
		Message interface{}
	}


CODE;

	$recoverCode =<<<CODE

	defer func() {
		if err := recover(); err != nil {
			panic(ProtoError{err})
		}
	}()


CODE;

	$code .= "type Request interface {\n";
	$code .= "	Process(*net.Session)\n";
	$code .= "	TypeName() string\n";
	$code .= "	GetModuleIdAndActionId() (int8, int8)\n";
	$code .= "}\n\n";

	$code .= "func DecodeIn(message []byte) Request {\n";
	$code .= $recoverCode;
	$code .= "	var moduleId = message[0]\n";
	$code .= "	switch moduleId {\n";
	foreach ($modules as $module) {
		$code .= "	case ".$module->ID.":\n";
		if ($module->ID == 99) {
			$code .= "		fail.When(!TheDebugConfig.IsEnableDebugAPI(), \"Disable Debug API\")\n";
		}
		$code .= "		return {$module->Name}_api.DecodeIn(message[1:])\n";
	}
	$code .= "	}\n";
	$code .= "	panic(\"unsupported protocol\")\n";
	$code .= "}\n\n";

	$code .= "func DecodeOut(message []byte) Request {\n";
	$code .= $recoverCode;
	$code .= "	var moduleId = message[0]\n";
	$code .= "	switch moduleId {\n";
	foreach ($modules as $module) {
		$code .= "	case ".$module->ID.":\n";
		$code .= "		return {$module->Name}_api.DecodeOut(message[1:])\n";
	}
	$code .= "	}\n";
	$code .= "	panic(\"unsupported protocol\")\n";
	$code .= "}\n\n";

	save_code($output_dir."/api.go", $code);

	system("gofmt -w=true {$output_dir}/api.go");
}

function gen_module_code($output_dir, $server_dir, $module) {
	$code = "package {$module->Name}_api\n\n";

	if (!is_dir($output_dir)) {
		mkdir($output_dir);
	}

	$code .= "import \"core/net\"\n";

	$code .= "type Request interface {\n";
	$code .= "	Process(*net.Session)\n";
	$code .= "	TypeName() string\n";
	$code .= "	GetModuleIdAndActionId() (int8, int8)\n";	
	$code .= "}\n\n";
	
	$code .= "var (\n";
	$code .= "	g_InHandler  InHandler\n";
	$code .= "	g_OutHandler OutHandler\n";
	$code .= ")\n\n";

	$code .= "func SetInHandler(handler InHandler) {\n";
	$code .= "	g_InHandler = handler\n";
	$code .= "}\n\n";

	$code .= "func SetOutHandler(handler OutHandler) {\n";
	$code .= "	g_OutHandler = handler\n";
	$code .= "}\n\n";

	// In Handler
	$code .= "type InHandler interface {\n";
	foreach ($module->Actions as $action) {
		if (is_array($action->InColumns)) {
			$type_name = get_type_name($action->Name);
			$code .= "	{$type_name}(*net.Session, *{$type_name}_In)\n";
		}
	}
	$code .= "}\n\n";

	// Out Handler
	$code .= "type OutHandler interface {\n";
	foreach ($module->Actions as $action) {
		if (is_array($action->OutColumns)) {
			$type_name = get_type_name($action->Name);
			$code .= "	{$type_name}(*net.Session, *{$type_name}_Out)\n";
		}
	}
	$code .= "}\n\n";

	// Decode in
	$code .= "func DecodeIn(message []byte) Request {\n";
	$code .= "	var actionId = message[0]\n";
	$code .= "	var buffer = net.NewBuffer(message[1:])\n";
	$code .= "	switch actionId {\n";
	foreach ($module->Actions as $action) {
		if (is_array($action->InColumns)) {
			$type_name = get_type_name($action->Name);
			$code .= "	case ".$action->ID.":\n";
			$code .= "		request := new({$type_name}_In)\n";
			$code .= "		request.Decode(buffer)\n";
			$code .= "		return request\n";
		}
	}
	$code .= "	}\n";
	$code .= "	_ = buffer\n";
	$code .= "	panic(\"unsupported request\")\n";
	$code .= "}\n\n";

	// Decode out
	$code .= "func DecodeOut(message []byte) Request {\n";
	$code .= "	var actionId = message[0]\n";
	$code .= "	var buffer = net.NewBuffer(message[1:])\n";
	$code .= "	switch actionId {\n";
	foreach ($module->Actions as $action) {
		if (is_array($action->OutColumns)) {
			$type_name = get_type_name($action->Name);
			$code .= "	case ".$action->ID.":\n";
			$code .= "		request := new({$type_name}_Out)\n";
			$code .= "		request.Decode(buffer)\n";
			$code .= "		return request\n";
		}
	}
	$code .= "	}\n";
	$code .= "	_ = buffer\n";
	$code .= "	panic(\"unsupported response\")\n";
	$code .= "}\n\n";

	// Type for types
	foreach ($module->Types as $type) {
		if (ProtoToken::IsEnumType($type->Kind)) {
			$type_name = get_type_name($type->Name);
			$code .= "type $type_name ";
			switch ($type->Kind) {
			case ProtoToken::ENUM8:
				$code .= "int8";
				break;
			case ProtoToken::ENUM16:
				$code .= "int16";
				break;
			case ProtoToken::ENUM32:
				$code .= "int32";
				break;
			case ProtoToken::ENUM64:
				$code .= "int64";
				break;
			}
			$code .= "\n\n";

			$code .= "const (\n";
			foreach ($type->Fields as $field) {
				$code .= "	".strtoupper($type->Name)."_".strtoupper($field->Name)." ".$type_name." = ".$field->Value."\n";
			}
			$code .= ")\n\n";
		}

		if (ProtoToken::STRUCT == $type->Kind) {
			$type_name = get_type_name($type->Name);
			$code .= gen_struct_code($type_name, $type->Fields);
		}
	}

	// Decode/Encode/Size for types
	foreach ($module->Types as $type) {
		if (ProtoToken::STRUCT == $type->Kind) {
			$type_name = get_type_name($type->Name);
			$code .= gen_decode_code($type_name, $type->Fields);
			$code .= gen_encode_code($type_name, $type->Fields);
			$code .= gen_size_code($type_name, $type->Fields);
		}
	}

	// Type for in/out
	foreach ($module->Actions as $action) {
		$request_type_name = get_type_name($action->Name);
		if (is_array($action->InColumns)) {
			$type_name = "{$request_type_name}_In";
			$code .= gen_struct_code($type_name, $action->InColumns);
			$code .= "func (this *{$type_name}) Process(session *net.Session) {\n";
			$code .= "	g_InHandler.{$request_type_name}(session, this)\n";
			$code .= "}\n\n";
			$code .= "func (this *{$type_name}) TypeName() string {\n";
			$code .= "	return \"{$module->Name}.{$action->Name}.in\"\n";
			$code .= "}\n\n";

			$code .= "func (this *{$type_name}) GetModuleIdAndActionId() (moduleId, actionId int8) {\n";
			$code .= "	return {$module->ID}, {$action->ID}\n";
			$code .= "}\n\n";
		}
		if (is_array($action->OutColumns)) {
			$type_name = "{$request_type_name}_Out";
			$code .= gen_struct_code($type_name, $action->OutColumns);
			$code .= "func (this *{$type_name}) Process(session *net.Session) {\n";
			$code .= "	g_OutHandler.{$request_type_name}(session, this)\n";
			$code .= "}\n\n";
			$code .= "func (this *{$type_name}) TypeName() string {\n";
			$code .= "	return \"{$module->Name}.{$action->Name}.out\"\n";
			$code .= "}\n\n";

			$code .= "func (this *{$type_name}) GetModuleIdAndActionId() (moduleId, actionId int8) {\n";
			$code .= "	return {$module->ID}, {$action->ID}\n";
			$code .= "}\n\n";
		}

		$code .= "func (this *{$type_name}) Bytes() []byte {\n";
		$code .= "	data := make([]byte, this.ByteSize())\n";
		$code .= "	this.Encode(net.NewBuffer(data))\n";
		$code .= "	return data\n";
		$code .= "}\n\n";
	}

	// Decode/Encode/Size for in/out
	foreach ($module->Actions as $action) {
		if (is_array($action->InColumns)) {
			$type_name = get_type_name($action->Name)."_In";
			$code .= gen_decode_code($type_name, $action->InColumns);
			$code .= gen_encode_code($type_name, $action->InColumns, $module->ID, $action->ID);
			$code .= gen_size_code($type_name, $action->InColumns, true);
		}
		if (is_array($action->OutColumns)) {
			$type_name = get_type_name($action->Name)."_Out";
			$code .= gen_decode_code($type_name, $action->OutColumns);
			$code .= gen_encode_code($type_name, $action->OutColumns, $module->ID, $action->ID);
			$code .= gen_size_code($type_name, $action->OutColumns, true);
		}
	}

	save_code($output_dir."/{$module->Name}_api.go", $code);

	system("gofmt -w=true {$output_dir}/{$module->Name}_api.go");
}

function gen_struct_code($type_name, $columns) {
	$last = '';
	$code = "type {$type_name} struct {\n";
	foreach ($columns as $column) {
		$code .= "	".get_type_name($column->Name)." ";
		switch ($column->Type) {
			case ProtoToken::IBOOL:
				$code .= "bool";
				break;
			case ProtoToken::INT8:
				$code .= "int8";
				break;
			case ProtoToken::INT16:
				$code .= "int16";
				break;
			case ProtoToken::INT32:
				$code .= "int32";
				break;
			case ProtoToken::INT64:
				$code .= "int64";
				break;
			case ProtoToken::UINT8:
				$code .= "uint8";
				break;
			case ProtoToken::UINT16:
				$code .= "uint16";
				break;
			case ProtoToken::UINT32:
				$code .= "uint32";
				break;
			case ProtoToken::UINT64:
				$code .= "uint64";
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::TEXT16:
			case ProtoToken::BINARY8:
			case ProtoToken::BINARY16:
				$code .= "[]byte";
				break;
			case ProtoToken::LIST8:
			case ProtoToken::LIST16:
				$name = get_type_name($column->Name);
				$code .= "[]{$type_name}_{$name}";
				$last .= gen_struct_code("{$type_name}_{$name}", $column->Items);
				break;
			default:
				if ($column->RefType != '') {
					$code .= get_type_name($column->RefType);
				}
				break;
		}
		$code .= "`json:\"".$column->Name."\"`";
		$code .= "\n";
	}
	$code .= "}\n\n";
	$code .= $last;
	return $code;
}

function gen_decode_code($type_name, $columns) {
	$last = '';
	$code = "func (this *{$type_name}) Decode(buffer *net.Buffer) {\n";
	foreach ($columns as $column) {
		$col_type_name = get_type_name($column->Name);
		switch ($column->Type) {
			case ProtoToken::IBOOL:
				$code .= "this.".$col_type_name." = buffer.ReadUint8() == 1";
				break;
			case ProtoToken::INT8:
				$code .= "this.".$col_type_name." = int8(buffer.ReadUint8())";
				break;
			case ProtoToken::INT16:
				$code .= "this.".$col_type_name." = int16(buffer.ReadUint16LE())";
				break;
			case ProtoToken::INT32:
				$code .= "this.".$col_type_name." = int32(buffer.ReadUint32LE())";
				break;
			case ProtoToken::INT64:
				$code .= "this.".$col_type_name." = int64(buffer.ReadUint64LE())";
				break;
			case ProtoToken::UINT8:
				$code .= "this.".$col_type_name." = buffer.ReadUint8()";
				break;
			case ProtoToken::UINT16:
				$code .= "this.".$col_type_name." = buffer.ReadUint16LE()";
				break;
			case ProtoToken::UINT32:
				$code .= "this.".$col_type_name." = buffer.ReadUint32LE()";
				break;
			case ProtoToken::UINT64:
				$code .= "this.".$col_type_name." = buffer.ReadUint64LE()";
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::BINARY8:
				$code .= "this.".$col_type_name." = buffer.ReadBytes(int(buffer.ReadUint8()))";
				$break;
			case ProtoToken::TEXT16:
			case ProtoToken::BINARY16:
				$code .= "this.".$col_type_name." = buffer.ReadBytes(int(buffer.ReadUint16LE()))";
				break;
			case ProtoToken::LIST8:
				$name = get_type_name($column->Name);
				$code .= "this.".$col_type_name." = make([]{$type_name}_{$name}, buffer.ReadUint8())\n";
				$code .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
				$code .= "		this.{$col_type_name}[i].Decode(buffer)\n";
				$code .= "	}";
				$last .= gen_decode_code("{$type_name}_{$name}", $column->Items);
				break;
			case ProtoToken::LIST16:
				$name = get_type_name($column->Name);
				$code .= "this.".$col_type_name." = make([]{$type_name}_{$name}, buffer.ReadUint16LE())\n";
				$code .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
				$code .= "		this.{$col_type_name}[i].Decode(buffer)\n";
				$code .= "	}";
				$last .= gen_decode_code("{$type_name}_{$name}", $column->Items);
				break;
			case ProtoToken::ENUM8:
				$enum_type_name = get_type_name($column->RefType);
				$code .= "this.".$col_type_name." = ".$enum_type_name."(buffer.ReadUint8())";
				break;
			case ProtoToken::ENUM16:
				$enum_type_name = get_type_name($column->RefType);
				$code .= "this.".$col_type_name." = ".$enum_type_name."(buffer.ReadUint16LE())";
				break;
			case ProtoToken::ENUM32:
				$enum_type_name = get_type_name($column->RefType);
				$code .= "this.".$col_type_name." = ".$enum_type_name."(buffer.ReadUint32LE())";
				break;
			case ProtoToken::ENUM64:
				$enum_type_name = get_type_name($column->RefType);
				$code .= "this.".$col_type_name." = ".$enum_type_name."(buffer.ReadUint64LE())";
				break;
			default:
				if ($column->RefType != '') {
					$code .= "this.".$col_type_name.".Decode(buffer)";
				}
				break;
		}
		$code .= "\n";
	}
	$code .= "}\n\n";
	$code .= $last;
	return $code;
}


function gen_encode_code($type_name, $columns, $module_id = -1, $action_id = -1) {
	$last = '';
	$code = "func (this *{$type_name}) Encode(buffer *net.Buffer) {\n";
	if ($module_id >= 0) {
		$code .= "buffer.WriteUint8($module_id)\n";
		$code .= "buffer.WriteUint8($action_id)\n";
	}
	foreach ($columns as $column) {
		$col_type_name = get_type_name($column->Name);
		switch ($column->Type) {
			case ProtoToken::IBOOL:
				$code .= "if this.".$col_type_name." {\n";
				$code .= "	buffer.WriteUint8(1)";
				$code .= "} else {\n";
				$code .= "	buffer.WriteUint8(0)";
				$code .= "}";
				break;
			case ProtoToken::ENUM8:
			case ProtoToken::INT8:
				$code .= "buffer.WriteUint8(uint8(this.".$col_type_name."))";
				break;
			case ProtoToken::ENUM16:
			case ProtoToken::INT16:
				$code .= "buffer.WriteUint16LE(uint16(this.".$col_type_name."))";
				break;
			case ProtoToken::ENUM32:
			case ProtoToken::INT32:
				$code .= "buffer.WriteUint32LE(uint32(this.".$col_type_name."))";
				break;
			case ProtoToken::ENUM64:
			case ProtoToken::INT64:
				$code .= "buffer.WriteUint64LE(uint64(this.".$col_type_name."))";
				break;
			case ProtoToken::UINT8:
				$code .= "buffer.WriteUint8(this.".$col_type_name.")";
				break;
			case ProtoToken::UINT16:
				$code .= "buffer.WriteUint16LE(this.".$col_type_name.")";
				break;
			case ProtoToken::UINT32:
				$code .= "buffer.WriteUint32LE(this.".$col_type_name.")";
				break;
			case ProtoToken::UINT64:
				$code .= "buffer.WriteUint64LE(this.".$col_type_name.")";
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::BINARY8:
				$code .= "buffer.WriteUint8(uint8(len(this.".$col_type_name.")))\n";
				$code .= "buffer.WriteBytes(this.".$col_type_name.")";
				break;
			case ProtoToken::TEXT16:
			case ProtoToken::BINARY16:
				$code .= "buffer.WriteUint16LE(uint16(len(this.".$col_type_name.")))\n";
				$code .= "buffer.WriteBytes(this.".$col_type_name.")";
				break;
			case ProtoToken::LIST8:
				$code .= "buffer.WriteUint8(uint8(len(this.".$col_type_name.")))\n";
				$code .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
				$code .= "		this.{$col_type_name}[i].Encode(buffer)\n";
				$code .= "	}";
				$last .= gen_encode_code("{$type_name}_{$col_type_name}", $column->Items);
				break;
			case ProtoToken::LIST16:
				$code .= "buffer.WriteUint16LE(uint16(len(this.".$col_type_name.")))\n";
				$code .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
				$code .= "		this.{$col_type_name}[i].Encode(buffer)\n";
				$code .= "	}";
				$last .= gen_encode_code("{$type_name}_{$col_type_name}", $column->Items);
				break;
			default:
				if ($column->RefType != '') {
					$code .= "this.".$col_type_name.".Encode(buffer)";
				}
				break;
		}
		$code .= "\n";
	}
	$code .= "}\n\n";
	$code .= $last;
	return $code;
}


function gen_size_code($type_name, $columns, $has_head = false, $for_list = false) {
	$last = '';
	$code = "func (this *{$type_name}) ByteSize() int {\n";
	if ($has_head) {
		$const_size = 2;
		$dynamic_size = '';
	} else {
		$const_size = 0;
		$dynamic_size = '';
	}
	foreach ($columns as $column) {
		$col_type_name = get_type_name($column->Name);
		switch ($column->Type) {
			case ProtoToken::IBOOL:
			case ProtoToken::ENUM8:
			case ProtoToken::UINT8:
			case ProtoToken::INT8:
				$const_size += 1;
				break;
			case ProtoToken::ENUM16:
			case ProtoToken::INT16:
			case ProtoToken::UINT16:
				$const_size += 2;
				break;
			case ProtoToken::ENUM32:
			case ProtoToken::UINT32:
			case ProtoToken::INT32:
				$const_size += 4;
				break;
			case ProtoToken::ENUM64:
			case ProtoToken::UINT64:
			case ProtoToken::INT64:
				$const_size += 8;
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::BINARY8:
				$const_size += 1;
				$dynamic_size .= "size += len(this.".$col_type_name.")\n";
				break;
			case ProtoToken::TEXT16:
			case ProtoToken::BINARY16:
				$const_size += 2;
				$dynamic_size .= "size += len(this.".$col_type_name.")\n";
				break;
			case ProtoToken::LIST8:
				$sub_code = gen_size_code("{$type_name}_{$col_type_name}", $column->Items, false, true);
				$const_size += 1;
				if (is_int($sub_code)) {
					$dynamic_size .= "size += len(this.".$col_type_name.") * $sub_code\n";
				} else {
					$last .= $sub_code;
					$dynamic_size .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
					$dynamic_size .= "		size += this.{$col_type_name}[i].ByteSize()\n";
					$dynamic_size .= "	}\n";
				}
				break;
			case ProtoToken::LIST16:
				$sub_code = gen_size_code("{$type_name}_{$col_type_name}", $column->Items, false, true);
				$const_size += 2;
				if (is_int($sub_code)) {
					$dynamic_size .= "size += len(this.".$col_type_name.") * $sub_code\n";
				} else {
					$last .= $sub_code;
					//$const_size += 2;
					$dynamic_size .= "	for i := 0; i < len(this.{$col_type_name}); i ++ {\n";
					$dynamic_size .= "		size += this.{$col_type_name}[i].ByteSize()\n";
					$dynamic_size .= "	}\n";
				}
				break;
			default:
				if ($column->RefType != '') {
					$dynamic_size .= "size += this.".$col_type_name.".ByteSize()\n";
				}
				break;
		}
	}
	$code .= "	size := $const_size\n";
	$code .= $dynamic_size;
	$code .= "	return size\n";
	$code .= "}\n\n";
	$code .= $last;

	if ($for_list && $dynamic_size == '') {
		return $const_size;
	}

	return $code;
}
?>
