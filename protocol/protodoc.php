<?php

$saved_files = array();

// 保存文件，并记录文件名，用于自动删除旧文件和避免重复生成
function save_code($file, $code) {
	global $saved_files;
	if (!is_file($file) || file_get_contents($file) != $code) {
		file_put_contents($file, $code);
	}
	$saved_files[] = realpath($file);
}

// 删除旧文件
function remove_old_files($dirPath){
	global $saved_files;

	if(file_exists($dirPath)){
		$dirFiles=scandir($dirPath);
		foreach ($dirFiles as $dirFile) {
			$dirFilePath = $dirPath."/".$dirFile;
			if (filetype($dirFilePath) == "file" ) {
				if (!in_array(realpath($dirFilePath), $saved_files)) {
					unlink($dirFilePath);
				}
			}
		}
	}
}

//==========================================================
// 命令行参数处理
//==========================================================

$command = isset($argv)?$argv[1]:null;    // 命令
$server_dir = isset($argv)?$argv[2]:null; // 目录名

if ($command == "invoke") {
	$plugin_name = dirname(__FILE__)."/plugins/".$argv[3]."/main.php";

	if (!is_file($plugin_name)) {
		echo "调用不存在的插件".$argv[3]."\n";
		exit;
	}

	require_once $plugin_name;

	plugin_main();
}
else {
	// echo "不支持的操作\n";
}

// var_export(load_modules());

//==========================================================
// 功能函数
//==========================================================

// 解析所有通讯协议模块
function parse_all($protodoc_dir)
{
	$dir = dirname(__FILE__)."/servers/".$protodoc_dir."/";

	if (!is_dir($dir)) {
		echo "协议目录".$protodoc_dir."不存在\n";
		exit;
	}

	$files = array();

	if ($dh = opendir($dir)) {
		while (($file = readdir($dh)) !== false) {
			if (strlen($file) - strrpos($file, ".pd") !== 3)
				continue;
			$files[] = $dir.$file;
		}
		closedir($dh);
	}

	asort($files);

	$result = array();
	
	foreach ($files as $file) {
		$text = file_get_contents($file);
		$parser = new ProtoParser($text);
		$modules = $parser->ParseModules();
		$result = array_merge($result, $modules);
	}

	verify_modules($result);

	return $result;
}

// 检查协议模块是否存在ID冲突、命名冲突、ID越界等情况
function verify_modules($modules)
{
	foreach ($modules as $module1) {
		if ($module1->ID < 0 || $module1->ID > 255)
			throw new Exception("协议模块'{$module1->Name}'的ID超出允许的0~255范围");

		foreach ($modules as $module2) {
			if ($module1 == $module2)
				continue;

			if ($module1->Name == $module2->Name)
				throw new Exception("重复声明协议模块'{$module1->Name}'");

			if ($module1->ID == $module2->ID)
				throw new Exception("协议模块'{$module1->Name}'和'{$module2->Name}'存在ID冲突");
		}

		$actions = $module1->Actions;

		foreach ($actions as $action1) {
			if ($action1->ID < 0 || $action1->ID > 255)
				throw new Exception("协议模块'{$module1->Name}'中，通讯接口'{$action1->Name}'的ID超出允许的0~255范围");

			foreach ($actions as $action2) {
				if ($action1 == $action2)
					continue;

				if ($action1->Name == $action2->Name)
					throw new Exception("协议模块'{$module1->Name}'中，重复声明通讯接口'{$action1->Name}'");

				if ($action1->ID == $action2->ID)
					throw new Exception("协议模块'{$module1->Name}'中，通讯接口'{$action1->Name}'和'{$action1->Name}'存在ID冲突");
			}
		}
	}
}

//==========================================================
// 协议描述语言
//==========================================================

// 标识符类型
class ProtoToken
{
	const EOF     = 'EOF';
	const MOD     = 'MOD';
	const API     = 'API';
	const NUM     = 'NUM';
	const ID      = 'ID';
	const COMMENT = 'COMMENT';
	const TYPE    = 'TYPE';
	const STR     = 'STR';
	
	const LB = '{';
	const RB = '}';
	const LT = '<';
	const GT = '>';
	const EQ = '=';
	const CL = ':';
	const CO = ',';
	const LS = '[';
	const RS = ']';
	const LR = '(';
	const RR = ')';
	const QT = '"';
	
	const IN  = 'in';
	const OUT = 'out';
	
	const IBOOL = 't_bool';

	const INT8  = 't_int8';
	const INT16 = 't_int16';
	const INT32 = 't_int32';
	const INT64 = 't_int64';
	
	const UINT8  = 't_uint8';
	const UINT16 = 't_uint16';
	const UINT32 = 't_uint32';
	const UINT64 = 't_uint64';
	
	const ENUM8  = 't_enum8';
	const ENUM16 = 't_enum16';
	const ENUM32 = 't_enum32';
	const ENUM64 = 't_enum64';
	
	const TEXT8  = 't_text8';
	const TEXT16 = 't_text16';
	
	const LIST8  = 't_list8';
	const LIST16 = 't_list16';

	const BINARY8 = 't_binary8';
	const BINARY16 = 't_binary16';

	const STRUCT = 't_struct';
	
	public static function IsType($token)
	{
		return strpos($token, 't_') === 0;
	}
	
	public static function IsEnumType($token)
	{
		return strpos($token, 't_enum') === 0;
	}
	
	public static function IsListType($token)
	{
		return strpos($token, 't_list') === 0;
	}
	
	public static function IsTextType($token)
	{
		return strpos($token, 't_text') === 0;
	}

	public static function IsTypeDefine($token)
	{
		return strpos($token, 't_enum') === 0 || $token == 't_struct';
	}
}

// 词法分析
class ProtoLexer
{
	private $text;          // 所解析的文本
	private $cursor;        // 解析的当前位置
	private $end_of_text;   // 文本的结束位置
	private $current_token; // 当前标识符
	private $current_value; // 当前标识符对应的值
	private $current_row;   // 行号（错误信息用）
	private $current_col;   // 列号（错误信息用）
	
	function __construct($text)
	{
		$this->text = $text;
		$this->cursor = 0;
		$this->end_of_text = strlen($text);
		$this->current_token = -1;
		$this->current_value = '';
		$this->current_row = 1;
		$this->current_col = 0;
	}
	
	public function IsEnd()
	{
		return $this->current_token == ProtoToken::EOF;
	}
	
	public function CurrentToken()
	{
		return $this->current_token;
	}
	
	public function CurrentValue() 
	{
		return $this->current_value;
	}
	
	public function CurrentRow()
	{
		return $this->current_row;
	}
	
	public function CurrentCol()
	{
		return $this->current_col;
	}
	
	private function NextChar()
	{
		$this->cursor += 1;
		$this->current_col += 1;
	}
	
	// 获取下一个标识符
	public function NextToken()
	{
		if ($this->cursor == $this->end_of_text) {
			$this->current_token = ProtoToken::EOF;
			return;
		}
		
		// 忽略空白符
		while (true) {
			$c = $this->text[$this->cursor];
			
			if ($c != " " && $c != "\t" && $c != "\r" && $c != "\n")
				break;
			
			if ($c == "\r") {
				// Mac平台的换回行
				$this->current_row += 1;
				$this->current_col = 0;
				$this->NextChar();
				
				// Windows平台的换行
				if ($c == "\n")
					$this->NextChar();
			}
			else if ($c == "\n") {
				// Linux平台的换行
				$this->current_row += 1;
				$this->current_col = 0;
				$this->NextChar();
			}
			else {
				$this->NextChar();
			}
			
			if ($this->cursor == $this->end_of_text) {
				$this->curr_tok = ProtoToken::EOF;
				return;
			}
		}
		
		// 特殊符号
		switch ($c) {
			case '<' :
			case '>' :
			case '{' : 
			case '}' : 
			case '=' : 
			case ':' : 
			case ',' :
			case '[' :
			case ']' :
			case '(' :
			case ')':
				$this->current_token = $c;
				$this->current_value = $c;
				$this->NextChar(); 
				return;
		}
		
		// 标识
		if ($c == '_' || ctype_alpha($c)) {
			$id_begin = $this->cursor;
			$id_length = 0;
		
			while ($this->cursor != $this->end_of_text) {
				$c = $this->text[$this->cursor];
		
				if ($c != '_' && !ctype_alpha($c) && !ctype_digit($c))
					break;
				
				$this->NextChar();
				
				$id_length += 1;
			}
		
			$this->current_value = substr($this->text, $id_begin, $id_length);
			
			if (array_key_exists($this->current_value, self::$keywords))
				$this->current_token = self::$keywords[$this->current_value];
			else
				$this->current_token = ProtoToken::ID;
			
			return;
		}
		
		// 数值
		if (ctype_digit($c) || $c == '+' || $c == '-') {
			$num_begin = $this->cursor;
			$num_length = 0;
			$ignore_sign = $c == '+' || $c == '-';
		
			while ($this->cursor != $this->end_of_text) {
				$c = $this->text[$this->cursor];
		
				if (!ctype_digit($c)) {
					if ($ignore_sign)
						$ignore_sign = false;
					else
						break;
				}
				
				$this->NextChar();
				
				$num_length += 1;
			}
		
			$this->current_value = substr($this->text, $num_begin, $num_length);
			$this->current_token = ProtoToken::NUM;
		
			return;
		}

		// 字符串
		if ($c == '"') {
			$this->NextChar();
				
			$content = '';
			$escape = false;
			
			// 截取字符串内容
			while ($this->cursor != $this->end_of_text) {
				$c = $this->text[$this->cursor];
		
				$this->NextChar();
				
				if ($c == '"' && !$escape) {
					break;
				}

				if ($escape) {
					switch ($c) {
					case 'n':
						$content .= "\n";
						break;
					case 't':
						$content .= "\t";
						break;
					case '"':
						$content .= "\"";
						break;
					default:
						$content .= $c;
						break;
					}

					$escape = fasle;
				} else {
					// 非转义状态遇到转义符
					if ($c == '\\') {
						$escape = true;
					} else {
						$content .= $c;
					}
				}
			}
			
			$this->current_value = $content;
			$this->current_token = ProtoToken::STR;
			
			return;
		}
		
		// 注释
		if ($c == '/') {
			// 忽略连续的'/'
			while ($this->cursor != $this->end_of_text) {
				$c = $this->text[$this->cursor];
				
				if ($c != '/')
					break;
				
				$this->NextChar();
			}
			
			$comment_begin = $this->cursor;
			$comment_length = 0;
			
			// 截取注释的内容
			while ($this->cursor != $this->end_of_text) {
				$c = $this->text[$this->cursor];
		
				if ($c == "\r" || $c == "\n") {
					break;
				}
		
				$this->NextChar();
				
				$comment_length += 1;
			}
			
			$this->current_value = trim(substr($this->text, $comment_begin, $comment_length));
			$this->current_token = ProtoToken::COMMENT;
			
			return;
		}
		
		throw new Exception("在第{$this->row_num}行，第{$this->col_num}列：遇到无法识别的表示符'$c'");
	}
	
	// 关键字表
	static $keywords = array(
		'mod'    => ProtoToken::MOD,
		'api'    => ProtoToken::API,
		'in'     => ProtoToken::IN,
		'out'    => ProtoToken::OUT,
		'type'   => ProtoToken::TYPE,
		'struct' => ProtoToken::STRUCT,
		
		'bool'   => ProtoToken::IBOOL,

		'int8'   => ProtoToken::INT8,
		'int16'  => ProtoToken::INT16,
		'int32'  => ProtoToken::INT32,
		'int64'  => ProtoToken::INT64,
			
		'uint8'  => ProtoToken::UINT8,
		'uint16' => ProtoToken::UINT16,
		'uint32' => ProtoToken::UINT32,
		'uint64' => ProtoToken::UINT64,
			
		'enum8'  => ProtoToken::ENUM8,
		'enum16' => ProtoToken::ENUM16,
		'enum32' => ProtoToken::ENUM32,
		'enum64' => ProtoToken::ENUM64,
			
		'text'   => ProtoToken::TEXT16,
		'text8'  => ProtoToken::TEXT8,
		'text16' => ProtoToken::TEXT16,
			
		'list8'  => ProtoToken::LIST8,
		'list16' => ProtoToken::LIST16,

		'binary'   => ProtoToken::BINARY8,
		'binary8'  => ProtoToken::BINARY8,
		'binary16' => ProtoToken::BINARY16,
			
		'byte'   => ProtoToken::INT8,
		'char'   => ProtoToken::INT8,
		'int'    => ProtoToken::INT32,
		'uint'   => ProtoToken::UINT32,
		'enum'   => ProtoToken::ENUM8,
		'list'   => ProtoToken::LIST8,
		'string' => ProtoToken::TEXT16,
	);
}

// 语法解析
class ProtoParser
{
	private $lexer;
	private $current_token;
	private $current_value;
	private $prev_row;
	private $prev_col;
	private $current_module;
	private $current_action;
	
	function __construct($text)
	{
		$this->lexer = new ProtoLexer($text);
	}
	
	private function NextToken($catch_comment = false)
	{
		for (;;) {
			// 保存前一个标识符的行号和列号
			$this->prev_row = $this->lexer->CurrentRow();
			$this->prev_col = $this->lexer->CurrentCol();
			
			$this->lexer->NextToken();
			
			// 缓存表示符信息
			$this->current_token = $this->lexer->CurrentToken();
			$this->current_value = $this->lexer->CurrentValue();
			
			// 需要捕获注释的时候，返回所有token
			if ($catch_comment)
				return;
			
			// 不需要捕获注释的时候，需要忽略注释
			if ($this->current_token == ProtoToken::COMMENT)
				continue;
			
			// 获取到不是注释的token，需要返回
			return;
		}
	}
	
	// 解析模块声明
	function ParseModules()
	{
		$modules = array();
		
		for (;;) {
			$module_comments = array();
			
			$this->NextToken(true);
			
			if ($this->current_token == ProtoToken::EOF)
				break;
			
			// 捕捉模块注释
			while ($this->current_token == ProtoToken::COMMENT) {
				if ($this->current_value != '')
					$module_comments[] = $this->current_value;
				$this->NextToken(true);
			}
			
			if ($this->current_token == ProtoToken::EOF)
				break;
			
			if ($this->current_token != ProtoToken::MOD)
				$this->Error("模块声明必须以mod关键字起始");
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::ID)
				$this->Error("模块声明起始位置缺少模块名称");
			
			$module_name = $this->current_value;
			
			$this->current_module = $module_name;
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::EQ)
				$this->Error("缺少必要的'='");
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::NUM)
				$this->Error("缺少模块编号");
			
			$module_id = $this->current_value;
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::LB)
				$this->Error("缺少起始的'{'");
			
			$types = array();
			$actions = array();
			
			$this->NextToken();

			for (;;) {
				$sub_comments = array();

				while ($this->current_token == ProtoToken::COMMENT) {
					if ($this->current_value != '')
						$sub_comments[] = $this->current_value;
					$this->NextToken(true);
				}
				
				$flags = $this->ParseFlags();

				if ($this->current_token == ProtoToken::TYPE) {
					$type = $this->ParseType($types);
					$types[$type->Name] = $type;
				} else if ($this->current_token == ProtoToken::API) {
					$action = $this->ParseAction($sub_comments, $flags, $types);
					$actions[] = $action;
				} else {
					break;
				}

				$this->NextToken(true);
			}

			if ($this->current_token != ProtoToken::RB)
				$this->Error("缺少结尾的'}'");
			
			$this->current_module  = "";
			
			$modules[] = new ProtoModule($module_name, $module_id, $module_comments, $types, $actions);
			
			$this->NextToken();
			
			if ($this->current_token == ProtoToken::EOF)
				break;
		}
		
		return $modules;
	}

	private function ParseType($types)
	{
		$this->NextToken();
		
		if ($this->current_token != ProtoToken::ID)
			$this->Error("类型定义缺少类型名");
		
		$type_name = $this->current_value;

		$this->NextToken();

		if (!ProtoToken::IsTypeDefine($this->current_token))
			$this->Error("类型定义缺少必要的类别标识");

		$type_kind = $this->current_token;

		$this->NextToken();
		
		if ($this->current_token != ProtoToken::LB)
			$this->Error("类型定义起始处缺少'{'");
		
		$this->NextToken();
		
		if (ProtoToken::IsEnumType($type_kind)) {
			$fields = $this->ParseEnumFields();
		} else if ($type_kind == ProtoToken::STRUCT) {
			$fields = $this->ParseColumns($types);
		}

		if ($this->current_token != ProtoToken::RB)
			$this->Error("类型定义结尾处缺少'}'");

		return new ProtoType($type_name, $type_kind, $fields);
	}
	
	// 解析枚举字段
	private function ParseEnumFields()
	{
		$fields = array();
		
		for (;;) {
			$field_comments = array();
			
			if ($this->current_token != ProtoToken::ID)
				$this->Error("枚举字段缺少名称");
			
			$field_name = $this->current_value;
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::EQ)
				$this->Error("枚举字段缺少'='");
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::NUM)
				$this->Error("枚举字段缺少值定义");
			
			$field_value = $this->current_value;
			
			$this->NextToken(true);
			
			// 捕捉字段注释
			while ($this->current_token == ProtoToken::COMMENT) {
				if ($this->current_value != '')
					$field_comments[] = $this->current_value;
				$this->NextToken(true);
			}
			
			$fields[] = new ProtoEnumField($field_name, $field_value, $field_comments);
			
			if ($this->current_token == ProtoToken::RB || $this->current_token == ProtoToken::EOF)
				break;
		}
		
		return $fields;
	}

	private function ParseFlags()
	{
		$flags = array();

		for (;;) {
			if ($this->current_token != ProtoToken::LS)
				break;

			$this->NextToken();

			if ($this->current_token != ProtoToken::ID)
				$this->Error("标签缺少名称");

			$flag_name = $this->current_value;

			$this->NextToken();

			$flag_content = '';
			if ($this->current_token == ProtoToken::LR) {
				$this->NextToken();
				
				if ($this->current_token != ProtoToken::STR)
					$this->Error("标签缺少内容");

				$flag_content = $this->current_value;

				$this->NextToken();

				if ($this->current_token != ProtoToken::RR)
					$this->Error("标签结尾缺少')'");

				$this->NextToken();
			}

			if ($this->current_token != ProtoToken::RS)
				$this->Error("标签结尾缺少']'");

			$flags[$flag_name] = $flag_content;

			$this->NextToken();
		}

		return $flags;
	}
	
	// 解析接口定义
	private function ParseAction($action_comments, $flags, $types)
	{
		$this->NextToken();
		
		if ($this->current_token != ProtoToken::ID)
			$this->Error("接口定义起始位置缺少接口名称");
		
		$action_name = $this->current_value;
		
		$this->current_action = $action_name;
		
		$this->NextToken();
		
		if ($this->current_token != ProtoToken::EQ)
			$this->Error("接口定义缺少必要的'='");
		
		$this->NextToken();
		
		if ($this->current_token != ProtoToken::NUM)
			$this->Error("接口定义缺少接口编号");
		
		$action_id = $this->current_value;
		
		$this->NextToken();
		
		if ($this->current_token != ProtoToken::LB)
			$this->Error("缺少起始的'{'");
		
		$this->NextToken();
		
		if ($this->current_token == ProtoToken::IN) {
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::LB)
				$this->Error("上行数据格式缺少起始的'{'");
			
			$this->NextToken();
			
			$in_columns = $this->ParseColumns($types);
			
			if ($this->current_token != ProtoToken::RB)
				$this->Error("上行数据格式缺少结尾的'}'");
			
			$this->NextToken();
		} else {
			$in_columns = null;
		}

		if ($this->current_token == ProtoToken::OUT) {
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::LB)
				$this->Error("下行数据格式缺少起始的'{'");
			
			$this->NextToken();
			
			$out_columns = $this->ParseColumns($types);
			
			if ($this->current_token != ProtoToken::RB)
				$this->Error("下行数据格式缺少结尾的'}'");
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::RB)
				$this->Error("缺少结尾的'}'");
		} else{
			$out_columns = null;
		}

		$this->current_action = "";
		
		return new ProtoAction($action_name, $action_id, $action_comments, $flags, $in_columns, $out_columns);
	}
	
	// 解析协议字段
	private function ParseColumns($types)
	{
		$columns = array();
		
		for (;;) {
			$column_comments = array();
			
			if ($this->current_token == ProtoToken::RB)
				break;
			
			if ($this->current_token != ProtoToken::ID)
				$this->Error("协议字段缺少名称");
			
			$column_name = $this->current_value;
			
			$this->NextToken();
			
			if ($this->current_token != ProtoToken::CL)
				$this->Error("协议字段名称后缺少':'");
			
			$this->NextToken();

			if (!ProtoToken::IsType($this->current_token) && $this->current_token != ProtoToken::ID)
				$this->Error("协议字段缺少类型");
			
			$column_type = $this->current_token;
			
			$items = null;

			if ($this->current_token == ProtoToken::ID) {
				if (!array_key_exists($this->current_value, $types)) {
					$this->Error("无法找到字段类型“".$this->current_value."”的定义");
				}

				$target_type = $types[$this->current_value];
				$column_type = $target_type->Kind;
				$items = $target_type->Fields;
				$reftype = $this->current_value;
			}
			else if (ProtoToken::IsListType($this->current_token)) {
				$this->NextToken();
				
				if ($this->current_token == ProtoToken::LT) {
					$this->NextToken();

					if (!ProtoToken::IsType($this->current_token) && $this->current_token != ProtoToken::ID)
						$this->Error("缺少列表类型声明");

					$target_type = $types[$this->current_value];
					$items = $target_type->Fields;
					$reftype = $this->current_value;

					$this->NextToken();

					if ($this->current_token != ProtoToken::GT)
						$this->Error("列表类型声明结尾处缺少'>'");
				} else {
					if ($this->current_token != ProtoToken::LB)
						$this->Error("列表定义起始处缺少'{'");
					
					$this->NextToken();

					$items = $this->ParseColumns($types);
				
					if ($this->current_token != ProtoToken::RB)
						$this->Error("列表定义结尾处缺少'}'");
				}
			}
			
			$this->NextToken(true);
			
			// 捕捉字段注释
			while ($this->current_token == ProtoToken::COMMENT) {
				if ($this->current_value != '')
					$column_comments[] = $this->current_value;
				$this->NextToken(true);
			}
			
			$columns[] = new ProtoColumn($column_name, $column_type, $column_comments, isset($reftype) ? $reftype : NULL, $items);
			
			if ($this->current_token == ProtoToken::RB || $this->current_token == ProtoToken::EOF)
				break;
		}
		
		return $columns;
	}
	
	// 错误处理
	private function Error($msg) {
		if ($this->current_module == "") {
			throw new Exception("在第".$this->prev_row."行，第".$this->prev_col."列：".$msg);
		}
		else {
			if ($this->current_action == "") {
				throw new Exception("在第".$this->prev_row."行，第".$this->prev_col."列：模块“".$this->current_module."”，".$msg);
			}
			else {
				throw new Exception("在第".$this->prev_row."行，第".$this->prev_col."列：接口“".$this->current_module.":".$this->current_action."”，".$msg);
			}
		}
	}
}

// AST: 协议模块
class ProtoModule
{
	public $Name;
	public $ID;
	public $Comments;
	public $Types;
	public $Actions;
	
	function __construct($name, $id, $comments, $types, $actions)
	{
		$this->Name     = $name;
		$this->ID       = $id;
		$this->Comments = $comments;
		$this->Types    = $types;
		$this->Actions  = $actions;
	}
}

// AST: 类型
class ProtoType
{
	public $Name;
	public $Kind;
	public $Fields;

	function __construct($name, $kind, $fields)
	{
		$this->Name = $name;
		$this->Kind = $kind;
		$this->Fields = $fields;
	}
}

// AST: 通讯接口
class ProtoAction
{
	public $Name;
	public $ID;
	public $Comments;
	public $Flags;
	public $InColumns;
	public $OutColumns;
	
	function __construct($name, $id, $comments, $flags, $in_columns, $out_columns)
	{
		$this->Name       = $name;
		$this->ID         = $id;
		$this->Comments   = $comments;
		$this->Flags      = $flags;
		$this->InColumns  = $in_columns;
		$this->OutColumns = $out_columns;
	}
}

// AST: 协议字段
class ProtoColumn
{
	public $Name;
	public $Type;
	public $Comments;
	public $RefType;
	public $Items;
	
	function __construct($name, $type, $comments, $reftype, $items)
	{
		$this->Name     = $name;
		$this->Type     = $type;
		$this->Comments = $comments;
		$this->RefType  = $reftype;
		$this->Items    = $items;
	}
}

// AST: 枚举字段
class ProtoEnumField
{
	public $Name;
	public $Value;
	public $Comments;
	
	function __construct($name, $value, $comments)
	{
		$this->Name     = $name;
		$this->Value    = $value;
		$this->Comments = $comments;
	}
}
?>
