
/*
 *  the annotation is in the file lang_rawstr.php
 */
var lang_rawstr = lang_rawstr || function(lang){

	var characters_php = {
		"\e": "\\e",
		"\f": "\\f",
		"\n": "\\n",
		"\r": "\\r",
		"\t": "\\t",
		"\v": "\\v",
		"\\": "\\\\",
		"\"": "\\\"",
		"$": "\\$",
	};

	var characters_js ={
		"\b": "\\b",
		"\f": "\\f",
		"\n": "\\n",
		"\r": "\\r",
		"\t": "\\t",
		"\\": "\\\\",
		"\'": "\\\'",
		"\"": "\\\"",
	};

	switch (lang) {
		case "js":
			this.characters = characters_js;
			break;
		case "mysql":
			// TODO
			break;
		case "go":
			// TODO
			break;
		case "php":
			this.characters = characters_php;
			break;
		default:
			new Error("Wrong lang_rawstr type!");
	}

	this.suitable_character = function(ch) {
		if (this.characters[ch] != undefined) {
			return this.characters[ch];
		} else {
			return ch;
		}
	};

	this.replace = function(str) {
		var done_str = "";
		for (var i = 0; i < str.length; i++) {
			done_str += this.suitable_character(str.charAt(i));
		}
		return done_str;
	};

};

