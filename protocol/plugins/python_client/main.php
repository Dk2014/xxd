<?php

function plugin_main() {
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
	"/python_client/protocol/";

	if (!is_dir($output_dir))
		mkdir($output_dir);

	global $saved_files;
	remove_old_files($output_dir);

	protocol_code($output_dir, $modules);

	interface_code($output_dir);

	//gen_main_code($output_dir, $server_dir, $modules);
	foreach ($modules as $module) {
		gen_module_code($output_dir, $server_dir, $module);
	}

}

function protocol_code($output_dir, $modules) {
	$code = "";

	$code .= <<<CODE
#!/usr/bin/evn python
# -*- coding: utf8 -*-



CODE;

	//import
	foreach ($modules as $module) {
		$down = false;
		foreach ($module->Actions as $action) {
			if (is_array($action->OutColumns)) {
				$down = true;
				break;
			}
		}
		if ($down) {
			$code .= "import {$module->Name}\n";
		}
	}

	//router

	$code .= <<<CODE


class BaseProtocol(object):
    def head_size(self):
        raise NotImplementedError

    def handle_head(self, session, head):
        """return body_size, receive_bytes"""
        raise NotImplementedError

    def handler_body(self, module_id, session, data):
        pass


CODE;
	foreach ($modules as $module) {
		if (is_array($action->OutColumns)) {
			$py_module_name = pyid($module->Name);

			$code .=  <<<CODE
        if module_id == {$module->ID }:
            action_id = ord(data[0])
            msg = self.{$py_module_name}.decode(data)
            callbacks = self.{$py_module_name}.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return


CODE;
		}
	}

	$code .= <<<CODE
    def __init__(self):

CODE;

	foreach($modules as $module) {
		if (is_array($action->OutColumns)) {
			$module_name = get_type_name($module->Name);
			$code .= <<<CODE
        self.{$module->Name} = {$module->Name}.{$module_name}Module()

CODE;
		}
	}
	save_code($output_dir . "/__init__.py", $code);

}

function interface_code($output_dir) {
	$code = <<<CODE
#!/usr/bin/evn python
# -*- coding: utf8 -*-


class BaseProtocol(object):
    def __init__(self, client):
        self._client = client

    def head_size(self):
        raise NotImplementedError

    def body_size(self,head):
        raise NotImplementedError

    def handle_head(self, head):
        """
        :param head:
        :return: body_size, body_handler
        """
        raise NotImplementedError

    def handler(self):
        raise NotImplementedError


class BaseModule(object):
    def decode(self, buff):
        raise NotImplementedError

    def add_callback(self, action, callback):
        raise NotImplementedError

    def clear_callback(self):
        raise NotImplementedError
CODE;
	save_code($output_dir . "/interface.py", $code);
}

function get_type_name($name) {
	$items = explode('_', $name);
	$type_name = '';
	for ($i = 0; $i < count($items); $i ++) {
		$type_name .= ucfirst($items[$i]);
	}
	return $type_name;
}

function gen_module_code($output_dir, $server_dir, $module) {
	$module_name = get_type_name($module->Name);
	$code = <<<CODE
#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface


CODE;

	$module_code = '';
	

	$module_code .= <<<CODE
class {$module_name}Module(interface.BaseModule):

CODE;

	$module_code .="    decoder_map = {\n";
	foreach ($module->Actions as $action) {
		$class_name = get_type_name($action->Name); 
		$class_name .= 'Down';
		$module_code .= "        {$action->ID}: {$class_name}, \n";
	}
	$module_code .="    }\n";

	$module_code .= <<<CODE
    receive_callback = {}

    def decode(self, message):
        action = ord(message[0])
        decoder_maker = self.decoder_map[action]
        msg = decoder_maker()
        msg.decode(message[1:])
        return msg

    def add_callback(self, action, callback):
        if self.receive_callback.has_key(action):
            self.receive_callback[action].append(callback)
        else:
            self.receive_callback[action] = [callback,]

    def clear_callback(self):
        self.receive_callback = {}

CODE;

	foreach($module->Actions as $action) {
	$module_code .= <<<CODE

    def add_{$action->Name}(self, callback):
        self.add_callback({$action->ID}, callback)

CODE;
	}

	//类型
	foreach ($module->Types as $type) {
		if (ProtoToken::IsEnumType($type->Kind)) {
			foreach ($type->Fields as $field) {
				$code .= strtoupper($type->Name)."_".strtoupper($field->Name)." = ".$field->Value."\n";
			}
			$code .= "\n\n";
		}
		
		if (ProtoToken::STRUCT == $type->Kind) {
			$type_name = get_type_name($type->Name);
			$code .= gen_struct_code($type_name, $type->Fields);
		}
	}

	//上下行协议结构体
	foreach ($module->Actions as $action) {
		if (is_array($action->InColumns)) {
			$type_name = get_type_name($action->Name) . 'Up';
			$code .= gen_struct_code($type_name, $action->InColumns, $module->ID, $action->ID, true);
		}
		if (is_array($action->OutColumns)) {
			$type_name = get_type_name($action->Name) . 'Down';
			$code .= gen_struct_code($type_name, $action->OutColumns, $module->ID, $action->ID);
		}
	}

	$code .= $module_code;
	echo $code;
	save_code($output_dir . "/{$module->Name}.py", $code);
}


function gen_struct_code($type_name, $columns, $module_id=-1, $action_id=-1,$has_head = false) {
	$sub_code = "";
	$code = "class {$type_name}(object):\n";

	if ($module_id >= 0 ) {
		$code .= "    _module = {$module_id}\n";
		$code .= "    _action = {$action_id}\n";
		$code .= "\n";
	}

	// __init__
	$code .= "    def __init__(self):\n";
	$code .= "        pass\n";

	foreach ($columns as $column) {
		$py_column_name= pyid($column->Name);
		switch ($column->Type) {
			case ProtoToken::IBOOL:
				$code .= "        self.{$py_column_name} = False\n";
				break;
			case ProtoToken::INT8:
			case ProtoToken::INT16:
			case ProtoToken::INT32:
			case ProtoToken::INT64:
			case ProtoToken::UINT8:
			case ProtoToken::UINT16:
			case ProtoToken::UINT32:
			case ProtoToken::UINT64:
			case ProtoToken::ENUM8:
			case ProtoToken::ENUM16:
			case ProtoToken::ENUM32:
			case ProtoToken::ENUM64:
				$code .= "        self.{$py_column_name} = 0\n";
				break;
			case ProtoToken::TEXT8:
			case ProtoToken::TEXT16:
				$code .= "        self.{$py_column_name} = ''\n";
				break;
			case ProtoToken::BINARY8:
			case ProtoToken::BINARY16:
				$code .= "        self.{$py_column_name} = bytearray()\n";
				break;
			case ProtoToken::LIST8:
			case ProtoToken::LIST16:
				$column_type_name = get_type_name($type_name);
				$column_type_name .= get_type_name($py_column_name);
				$sub_code .= gen_struct_code($column_type_name, $column->Items);
				$code .= "        self.{$py_column_name} = []\n";
				break;
			default:
				if ($column->RefType != '') {
					$column_type_name = get_type_name($column->RefType);
					$code .= "        self.{$py_column_name} = {$column_type_name}()\n";
					break;
				}
		}
	}
	$code .= "\n";

	$code .= gen_encode_code($type_name, $columns);
	$code .= "\n";

	$code .= gen_decode_code($type_name, $columns);
	//$code .= "\n";

	$code .= gen_size_code($type_name, $columns, $has_head);
	$code .= "\n\n";

	$code .= $sub_code;
	return $code;
}

function gen_size_code($type_name, $columns, $has_head = false) {
	$const_size = 0;
	$dynamic_size = '';
	if ($has_head) {
		$const_size += 2;
	}
	$code = "    def size(self):\n";
	foreach ($columns as $column) {
		$py_column_name= pyid($column->Name);
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
				$dynamic_size .= "        size += len(self.{$py_column_name})\n";
				break;
			case ProtoToken::TEXT16:
			case ProtoToken::BINARY16:
				$const_size += 2;
				$dynamic_size .= "        size += len(self.{$py_column_name})\n";
				break;
			case ProtoToken::LIST8:
				$const_size += 1;
				$dynamic_size .= "        for item in self.{$py_column_name}:\n";
				$dynamic_size .= "            size += item.size()\n";
				break;
			case ProtoToken::LIST16:
				$const_size += 2;
				$dynamic_size .= "        for item in self.{$py_column_name}:\n";
				$dynamic_size .= "            size += item.size()\n";
				break;
			default:
				if ($column->RefType != '') {
					$dynamic_size .= "        size += self.{$py_column_name}.size()\n";
				}

		}
	}
	if ($dynamic_size == '') {
		return <<<CODE

    @staticmethod
    def size():
        return {$const_size}
        
CODE;
	}
	$code .= "        size = {$const_size}\n";
	$code .= $dynamic_size;
	$code .= "        return size\n";
	//if ($type_name == "EnterDown") {
		//echo  $code;
		//die("fo");
	//}

	return $code;
}

function gen_encode_code($type_name, $columns) {
	$code = "    def encode(self):\n";

	/*
	if (count($columns) > 0) {
		$code .= "        buff = bytearray()\n\n";
	} else {
		$code .= "        buff = bytearray()\n";
		$code .= "        buff.extend(struct.pack('<B', self._module))\n";
		$code .= "        buff.extend(struct.pack('<B', self._action))\n";
	}
	 */
	
	$code .= "        buff = bytearray()\n";
	$code .= "        buff.extend(struct.pack('<B', self._module))\n";
	$code .= "        buff.extend(struct.pack('<B', self._action))\n";

	foreach ($columns as $column) {
		$py_column_name= pyid($column->Name);
		$format = '';
		switch ($column->Type) {
		case ProtoToken::IBOOL:
			$format = '"<?"';
			break;
		case ProtoToken::INT8:
			$format = '"<b"';
			break;
		case ProtoToken::INT16:
			$format = '"<h"';
			break;
		case ProtoToken::INT32:
			$format = '"<l"';
			break;
		case ProtoToken::INT64:
			$format = '"<q"';
			break;
		case ProtoToken::UINT8:
		case ProtoToken::ENUM8:
			$format = '"<B"';
			break;
		case ProtoToken::UINT16:
		case ProtoToken::ENUM16:
			$format = '"<H"';
			break;
		case ProtoToken::UINT32:
		case ProtoToken::ENUM32:
			$format = '"<L"';
			break;
		case ProtoToken::UINT64:
		case ProtoToken::ENUM64:
			$format = '"<Q"';
			break;
		case ProtoToken::TEXT8:
		case ProtoToken::BINARY8:
			$code .= "        buff.extend(struct.pack('<B', len(self.{$py_column_name})))\n";
			$code .= "        buff.extend(self.{$py_column_name})\n";
			break;
		case ProtoToken::TEXT16:
		case ProtoToken::BINARY16:
			$code .= "        buff.extend(struct.pack('<H', len(self.{$py_column_name})))\n";
			$code .= "        buff.extend(self.{$py_column_name})\n";
			break;
		case ProtoToken::LIST8:
			$code .= "        buff.extend(struct.pack('<B', len(self.{$py_column_name})))\n";
			$code .= "        for item in self.{$py_column_name}:\n";
			$code .= "            buff.extend(item.encode())\n";
			break;
		case ProtoToken::LIST16:
			$code .= "        buff.extend(struct.pack('<H', len(self.{$py_column_name})))\n";
			$code .= "        for item in self.{$py_column_name}:\n";
			$code .= "            buff.extend(item.encode())\n";
			break;
		default:
			if ($column->RefType != '') {
			$code .= "        buff.extend(self.{$py_column_name}.encode())\n";
			}
			break;
		}
		if (is_simple_type($column->Type)) {
			$code .= "        buff.extend(struct.pack({$format}, self.{$py_column_name}))\n";
		}

	}
	$code .= "        return buff\n";
	return $code;
}

function gen_decode_code($type_name, $columns) {
	//echo "gen_decode_code\n";

	$code = "    def decode(self, raw_msg):\n";
	if (count($columns)>0) {
		$code .= "        idx = 0\n\n";
	} else {
		$code .= "        pass\n";
	}

	foreach ($columns as $column) {
		$py_column_name= pyid($column->Name);
		$format = '';
		$size = 0;
		switch ($column->Type) {
		case ProtoToken::IBOOL:
			$format = '"<?"';
			$size = 1;
			break;
		case ProtoToken::INT8:
			$format = '"<b"';
			$size = 1;
			break;
		case ProtoToken::INT16:
			$format = '"<h"';
			$size = 2;
			break;
		case ProtoToken::INT32:
			$format = '"<l"';
			$size = 4;
			break;
		case ProtoToken::INT64:
			$format = '"<q"';
			$size = 8;
			break;
		case ProtoToken::UINT8:
		case ProtoToken::ENUM8:
			$format = '"<B"';
			$size = 1;
			break;
		case ProtoToken::UINT16:
		case ProtoToken::ENUM16:
			$format = '"<H"';
			$size = 2;
			break;
		case ProtoToken::UINT32:
		case ProtoToken::ENUM32:
			$format = '"<L"';
			$size = 4;
			break;
		case ProtoToken::UINT64:
		case ProtoToken::ENUM64:
			$format = '"<Q"';
			$size = 8;
			break;
		case ProtoToken::TEXT8:
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<B\", raw_msg, idx)[0]\n";
			$code .= "        idx += 1\n";
			$code .= "        self.{$py_column_name} = str(raw_msg[idx:idx+_{$py_column_name}_size])\n";
			$code .= "        idx += _{$py_column_name}_size\n";
			break;
		case ProtoToken::BINARY8:
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<B\", raw_msg, idx)[0]\n";
			$code .= "        idx += 1\n";
			$code .= "        self.{$py_column_name} = raw_msg[idx:idx+_{$py_column_name}_size]\n";
			$code .= "        idx += _{$py_column_name}_size\n";
			break;
		case ProtoToken::TEXT16:
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<H\", raw_msg, idx)[0]\n";
			$code .= "        idx += 2\n";
			$code .= "        self.{$py_column_name} = str(raw_msg[idx:idx+_{$py_column_name}_size])\n";
			$code .= "        idx += _{$py_column_name}_size\n";
			break;
		case ProtoToken::BINARY16:
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<H\", raw_msg, idx)[0]\n";
			$code .= "        idx += 2\n";
			$code .= "        self.{$py_column_name} = raw_msg[idx:idx+_{$py_column_name}_size]\n";
			$code .= "        idx += _{$py_column_name}_size\n";
			break;
		case ProtoToken::LIST8:
			$column_type_name = get_type_name($type_name);
			$column_type_name .= get_type_name($column->Name);
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<B\", raw_msg, idx)[0]\n";
			$code .= "        idx += 1\n";
			$code .= "        for i in range(_{$py_column_name}_size):\n";
			$code .= "            obj = {$column_type_name}()\n";
			$code .= "            obj.decode(raw_msg[idx:])\n";
			$code .= "            idx += obj.size()\n";
			$code .= "            self.{$py_column_name}.append(obj)\n";
			break;
		case ProtoToken::LIST16:
			$column_type_name = get_type_name($type_name);
			$column_type_name .= get_type_name($column->Name);
			$code .= "        _{$py_column_name}_size = struct.unpack_from(\"<H\", raw_msg, idx)[0]\n";
			$code .= "        idx += 2\n";
			$code .= "        for i in range(_{$py_column_name}_size):\n";
			$code .= "            obj = {$column_type_name}()\n";
			$code .= "            obj.decode(raw_msg[idx:])\n";
			$code .= "            idx += obj.size()\n";
			$code .= "            self.{$py_column_name}.append(obj)\n";
			break;
		default:
			if ($column->RefType != '') {
				//$code .= "        _{$py_column_name}_obj = {$column_name
				$column_type_name = get_type_name($column->RefType);
				$code .= "        self.".$py_column_name.".decode(raw_msg[idx:])\n";
				$code .= "        idx += self.".$py_column_name.".size()\n";
			}
			break;
		}
		if (is_simple_type($column->Type)) {
			$code .= "        self.{$py_column_name} = struct.unpack_from({$format}, raw_msg, idx)[0]\n";
			$code .= "        idx += {$size}\n";
		}
		$code .= "\n";
	}
	return $code;
}

function is_simple_type($type) {
	switch ($type) {
	case ProtoToken::IBOOL:
	case ProtoToken::INT8:
	case ProtoToken::INT16:
	case ProtoToken::INT32:
	case ProtoToken::INT64:
	case ProtoToken::UINT8:
	case ProtoToken::UINT16:
	case ProtoToken::UINT32:
	case ProtoToken::UINT64:
	case ProtoToken::ENUM8:
	case ProtoToken::ENUM16:
	case ProtoToken::ENUM32:
	case ProtoToken::ENUM64:
		return true;
	case ProtoToken::TEXT16:
	case ProtoToken::BINARY16:
	case ProtoToken::LIST8:
	case ProtoToken::LIST16:
		return false;
	default:
		return false;
	}
}

//python identifier wraper `class` to `class_`
function pyid($keyword) {
	$kw_list = array(
		'and',
		'as',
		'assert',
		'break',
		'class',
		'continue',
		'def',
		'del',
		'elif',
		'else',
		'except',
		'exec',
		'finally',
		'for',
		'from',
		'global',
		'if',
		'import',
		'in',
		'is',
		'lambda',
		'not',
		'or',
		'pass',
		'print',
		'raise',
		'return',
		'try',
		'while',
		'with',
		'yield'
	);
	if( in_array($keyword, $kw_list)) {
		return $keyword . "_";
	}
	return $keyword;

}
