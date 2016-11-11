<?php
class ParseConfig {

	private $typeStr = '';
	private $dataStr = '';
	private $constStr = '';
	private $switchStr = '';
	private $typeArray = array();
	private $config = array();
	private $goType = array( 'int8', 'int', 'int32', 'int64', 'float32', 'float64', 'string', 'byte', 'char', 'int16', 'bool' );

	function __construct() {
	}

	function Parse( $config ) {
		$this->typeStr = '';
		$this->dataStr = '';
		$this->constStr = '';
		$this->switchStr = '';
		$this->typeArray = array();
		$this->config = $config;
		foreach ( $config as $key => $value ) {
			switch ( $value[0] ) {
			case 'type':
				$this->parseType( $value );
				break;
			case 'const':
				$this->parseConst( $value );
				break;
			case 'data':
				$this->parseData( $value );
				break;
			case 'switch':
				$this->parseSwitch( $value );
				break;
			case 'between':
				/*
				array('between', 'x' => 'int', "getBetween" => 'int', 'assert' => false, 'default' => 65,
					array("(,5]", 0),
					array("(5,10]", 7),
					array("(16,)", 5),
				)
				*/
				$this->parseBetween( $value );
				break;
			default:
				echo "don't parse this type ".$value[0];
				exit;
			}
		}
		return $this->constStr . "\n". $this->typeStr ."\n". $this->dataStr . "\n". $this->switchStr . "\n";
	}

	function parseConst( $value ) {
		$this->constStr .= "const( \n";
		foreach ( $value as $key => $v ) {
			if ( $key === 0 ) continue;
			if ( is_numeric( $v ) ) {
				$this->constStr.= '	'.strtoupper( $key ) . " = " .$v. "\n";
			}else {
				$this->constStr.= '	'.strtoupper( $key ) . " =  \"$v\"\n";
			}
		}
		$this->constStr .= ")\n";
	}

	function parseSwitch( $value ) {
		$param = $paramType = '';
		$func = $fucnReturn = '';
		$toAssertNull = false;
		$i = 0;
		$multiRet = false;
		$default = false;
		foreach ( $value as $key => $v ) {
			if ( $i == 1 ) {
				$param = $this->getName( $key, false );
				$paramType = $v;
			} elseif ( $i == 2 ) {
				$func = $this->getName( $key, true );
				if (!is_array($v)){
					$funcReturn = $this->getName( $v, false );
					$this->switchStr .= "func $func($param $paramType) (res $funcReturn) {\n";
				}else{
					$multiRet = true;
					$this->switchStr .= "func $func($param $paramType) (";
					foreach ($v as $ri => $ret) {
						$this->switchStr .= "r$ri ".$this->getName( $ret, false ) .",";
					}
					$this->switchStr .= ") {\n";
				}
				
				$this->switchStr .= "	switch $param {\n";
			} elseif ( $i == 3 ) {
				if ( !is_bool( $v ) ) {
					exit( "$func 第三个元素匹配错误:设置是否对取不到的配置进行抛错" );
				}
				$toAssertNull = $v;
			}elseif ($i == 4 && $key == 'default') {
				if (!$multiRet) {
					$default = "		return $v\n";
				}else{
					$default .= "		return ". implode(",", $v)."\n";
				}
			} elseif ($i != 0) {
				foreach( $v as $case => $ret) {
					$this->switchStr .= "	case $case:\n";
					if (!$multiRet) {
						$this->switchStr .= "		return $ret\n";
					}else{
						$this->switchStr .= "		return ". implode(",", $ret)."\n";
					}
				}
			}
			$i++;
		}
		if (!is_bool($default) ) {
			$this->switchStr .= "	default:\n$default\n";
		}
		$this->switchStr .= "	}\n";
		if ( $toAssertNull ) {
			$this->switchStr .= "	panic(\"$func err\")\n";
		}
		$this->switchStr .= "	return\n";
		$this->switchStr .= "}\n";
	}

	function parseBetween( $value ) {
		$param = $paramType = '';
		$func = $fucnReturn = '';
		$toAssertNull = false;
		$i = 0;
		$multiRet = false;
		$default = false;
		foreach ( $value as $key => $v ) {
			if ( $i == 1 ) {
				$param = $this->getName( $key, false );
				$paramType = $v;
			} elseif ( $i == 2 ) {
				$func = $this->getName( $key, true );
				if (!is_array($v)){
					$funcReturn = $this->getName( $v, false );
					$this->switchStr .= "func $func($param $paramType) (res $funcReturn) {\n";
				}else{
					$multiRet = true;
					$this->switchStr .= "func $func($param $paramType) (";
					foreach ($v as $ri => $ret) {
						$this->switchStr .= "r$ri ".$this->getName( $ret, false ) .",";
					}
					$this->switchStr .= ") {\n";
				}
			} elseif ( $i == 3 ) {
				if ( !is_bool( $v ) ) {
					exit( "$func 第三个元素匹配错误:设置是否对取不到的配置进行抛错" );
				}
				$toAssertNull = $v;
			}elseif ($i == 4 && $key == 'default') {
				if (!$multiRet) {
					$default = "	return $v\n";
				}else{
					$default .= "	return ". implode(",", $v)."\n";
				}
			} elseif ($i != 0) {
				$ret = $v[1];
				$conditions = $this->betweenCond(trim($v[0]), $param);

				if ($conditions == ''){
					continue;
				}

				$this->switchStr .= "	if $conditions {\n";
				if (!$multiRet) {
					$this->switchStr .= "		return $ret\n";
				}else{
					$this->switchStr .= "		return ". implode(",", $ret)."\n";
				}
				$this->switchStr .= "	}\n";
			}
			$i++;
		}

		if ( $toAssertNull ) {
			$this->switchStr .= "	panic(\"$func err\")\n";
		}

		if (!is_bool($default) ) {
			$this->switchStr .= "$default";
		}else{
			$this->switchStr .= "	return\n";
		}
		
		$this->switchStr .= "}\n";
	}

	function parseType( $value ) {
		$goType = $this->getName( $value[1], true );
		$this->typeStr .= "type {$goType} struct {\n";
		if ( !isset( $this->typeArray[$goType] ) ) {
			$this->typeArray[$goType] = array();
		}
		foreach ( $value as $key => $type ) {
			if ( $key === 0 || $key === 1 ) continue;
			list($ok, $isSlice, $isPointer, $type, $isDefaultType ) = $this->parseNameType( $type );
			if ($isPointer) $type = '*'.$type; 
			if ($isSlice) $type = '[]'.$type; 
			$this->typeArray[$goType][] = $type;
			
			// 如果是false的数据，不生成
			if (!$ok) continue;

			$this->typeStr .= "	".$this->getName( $key, true )."  $type\n";
		}
		$this->typeStr .= "}\n";
	}

	function parseData( $value ) {
		$i = 0;
		$isSlice = $isPointer = $isDefaultType = false;
		$type = '';
		foreach ( $value as $key => $v ) {
			if ( $i == 1 ) {
				$i++;
				// 解析名称及类型
				$goName = $this->getName( $key, true );
				$this->dataStr .= "var ".$goName." ";
				list($ok, $isSlice, $isPointer, $type, $isDefaultType ) = $this->parseNameType( $v );
				if (!$ok) continue;
				if ( $isSlice ) {
					$isSlice = true;
					$this->dataStr .= "= []".( $isPointer ? '*' : '' )."$type{";
					if (!$isDefaultType) {
						$this->dataStr .= "\n";
					}
				} else {
					$this->dataStr .= "$type = ";
				}
			}else if ( $i > 1 ) {
				// 解析数据
				$this->parseValue( $type, $v, false, $isPointer );
				if ( $isSlice && !$isDefaultType ) {
					$this->dataStr .= "\n";
				}
			}
			$i++;
		}
		if ( $isSlice ) {
			$this->dataStr .= "}\n";
		}else {
			$this->dataStr = substr( $this->dataStr, 0, strlen( $this->dataStr )-1 );
		}
		$this->dataStr .= "\n";
	}

	function parseValue( $type, $value, $isSlice, $isPointer ) {
		if ( is_array( $value ) ) {
			if ( $isSlice && $isPointer ) {
				$this->dataStr .= "[]*$type{";
			}else {
				if ( $isSlice ) {
					$this->dataStr .= "[]";
				}elseif ( $isPointer ) {
					$this->dataStr .= "&";
				}
				$this->dataStr .= "\t$type{";
			}

			foreach ( $value as $k=> $v ) {
				if ( $isSlice ) {
					$this->parseValue( $type, $v, false, $isPointer );
				}else {
					if ( isset( $this->typeArray[$type][$k] ) ) {
						$kt = $this->typeArray[$type][$k] ;
					}else {
						$kt = '';
					}
					list($ok, $isSliceTemp, $isPointerTemp, $t, $isDefaultType ) = $this->parseNameType( $kt );
					if (!$ok) continue;
					$this->parseValue( $t, $v, $isSliceTemp, $isPointerTemp );
				}
			}
			$this->dataStr .= "}," ;
		}else {
			switch ( $type ) {
			case 'string':
				$this->dataStr .= "\"$value\",";
				break;
			case 'byte':
				$this->dataStr .= "'$value',";
				break;
			case 'bool':
				$this->dataStr .= (($value ) ? 'true,' : 'false,');
				break;
			default:
				$this->dataStr .= "$value,";
				break;
			}
		}
	}

	function getName( $name, $isPublic ) {
		if ( is_array( $name ) ) return $name;
		$items = explode( '_', $name );
		$x = '';
		for ( $i = 0; $i < count( $items ); $i ++ ) {
			$y = $items[$i];
			if ( $isPublic || $i > 0 ) {
				$y[0] = strtoupper( $y[0] );
			}
			$x .= $y;
		}
		return $x;
	}

	function parseNameType( $type ) {
		$isSlice = false;
		$isPointer = $isDefaultType = false;
		$realType = $type;
		$ok = true;
		if ($type == 'client') {
			return array(false, $isSlice, $isPointer, $realType, $isDefaultType);
		}
		if ( substr( $realType, 0, 2 ) == '[]' ) {
			$isSlice = true;
			$realType = substr( $realType, 2 );
		}
		if ( substr( $realType, 0, 1 ) == '*' ) {
			$isPointer = true;
			$realType = substr( $realType, 1 );
		}
		$realType = $this->getName( $realType, true );
		if ( in_array( strtolower( $realType ), $this->goType ) ){
			$isDefaultType = true;
			$realType = strtolower( $realType );
		}
		return array($ok, $isSlice, $isPointer, $realType, $isDefaultType);
	}

	function betweenCond($cond, $param) {
		$ok = preg_match("/([\(|\[])([0-9]*),([0-9]*)([\]|\)])/", $cond, $matchs);
		if ($ok === 0){
			throw new Exception("between match err", 1);
		}

		$carr = array();
		if ($matchs[2] != '' ){
			switch ($matchs[1]) {
				case '(':
					array_push($carr, "$param > {$matchs[2]}");
					break;
				case '[':
					array_push($carr, "$param >= {$matchs[2]}");
					break;
			}
		}

		if ($matchs[3] != '' ) {
			switch ($matchs[4]) {
				case ')':
					array_push($carr, "$param < {$matchs[3]}");
					break;
				case ']':
					array_push($carr, "$param <= {$matchs[3]}");
					break;
			}
		}

		return implode($carr, " && ");
	}
}

?>