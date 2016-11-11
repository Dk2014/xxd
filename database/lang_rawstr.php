<?php

class lang_rawstr {

	/*
	 *  The characters variable format:
	 *  left side is one character in the string of running language,
	 *  right side is what should be outputted while combining the target
	 *  language plain raw code
	 *
	 *  RIGHT SIDE IS THE ONLY REASON DRIVING THIS CONVERTION!
	 */

	// javascript string
	public static $characters_js = array(
		"\b" => "\\b",
		"\f" => "\\f",
		"\n" => "\\n",
		"\r" => "\\r",
		"\t" => "\\t",
		"\\" => "\\\\",
		"\'" => "\\\'",
		"\"" => "\\\"",
	);

	// mysql string
	public static $characters_mysql = array(
		"\b" => "\\b",
		"\Z" => "\\Z",
		"\n" => "\\n",
		"\r" => "\\r",
		"\t" => "\\t",
		"\\" => "\\\\",
		"'" => "\\'",
		"\"" => "\\\"",
		"\0" => "\\0",
		//"%" => "\\%",
		"_" => "\\_",
	);

	// php double quote string
	public static $characters_php_dq = array(
		"\e" => "\\e",
		"\f" => "\\f",
		//"\n" => "\\n",
		//"\r" => "\\r",
		//"\t" => "\\t",
		"\v" => "\\v",
		"\\" => "\\\\",
		"\"" => "\\\"",
		"\$" => "\\\$",
	);

	var $characters;
	var $char_type;

	/**
	 *  @param "js":     javascript string;
	 *         "mysql":  mysql string;
	 *         "go":     go string -- TODO;
	 *         "php_dq": the double quote string in php;
	 */
	function __construct($lang){
		switch ($lang) {
			case "js":
				$this->characters = self::$characters_js;
				$this->char_type = "js";
				break;
			case "mysql":
				$this->characters = self::$characters_mysql;
				$this->char_type = "mysql";
				break;
			case "go":
				// TODO
				break;
			case "php_dq":
				$this->char_type = "php";
				$this->characters = self::$characters_php_dq;
				break;
			default:
				die("Wrong lang_rawstr type!");
		}
	}

	function suitable_character($char) {
		if (isset($this->characters[$char])) {
			return $this->characters[$char];
		} else {
			return $char;
		}
	}

	/**
	 *  the user interface
	 *
	 *  @param the string of running language.
	 *  @return the string you can use combining into the string of target
	 *          language while generic another language from plain text
	 *          combination.
	 */
	function replace($str) {
		switch ($this->char_type) {
		case "js":
			return str_replace("\\r", "",  json_encode($str, JSON_UNESCAPED_UNICODE|JSON_UNESCAPED_SLASHES));
		}
		$done_str = "";
		for ($i = 0; $i < strlen($str); $i++) {
			$done_str .= $this->suitable_character($str[$i]);
		}
		return $done_str;
	}

};

