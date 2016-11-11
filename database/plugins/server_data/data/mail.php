<?php

function GetImport_mail() {
	return array(
		"strconv",
		"bytes",
	);
}

function GetCode_mail($db) {
	return getInterfaceAndDatas($db);
}



function getInterfaceAndDatas($db) {
	$mails = db_query($db, "SELECT * FROM mail");
	$code = '';

	$code .= getAttachments($db, $mails)."\n";


	foreach ($mails as $key => $m) {
		$code .= "// {$m['title']}\n";
		$mname = "Mail".$m['sign'];
		$code .= "type $mname struct {\n";
		$parameters = getMailParas($m['parameters']);
		foreach ($parameters as $key => $p) {
			$code .= "	{$p[0]} {$p[2]} // {$p[1]}\n";
		}
		$code .= "	Attachments []*Attachment\n";


		$code .= "}\n";
		$code .= "\n";
		$code .= "func (this $mname) GetMailId() int32 {\n";
		$code .= "	return int32({$m['id']})\n";
		$code .= "}\n";

		$code .= "\n";

		$code .= "\n";
		$code .= "func (this $mname) GetSendTime() int64 {\n";
		$code .= "	return 0\n";
		$code .= "}\n";

		$code .= "\n";
		$code .= "func (this $mname) GetExpireTime() int64 {\n";
		$code .= "	return {$m['expire_time']}\n";
		$code .= "}\n";


		$code .= "func (this $mname) GetTitle() string {\n";
		$code .= "	return \"\"\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetPriority() int8 {\n";
		$code .= "	return {$m['priority']}\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetContent() string {\n";
		$code .= "	return \"\"\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetMinLevel() int16 {\n";
		$code .= "	return {$m['min_level']}\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetMaxLevel() int16 {\n";
		$code .= "	return {$m['max_level']}\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetMinVIPLevel() int16 {\n";
		$code .= "	return {$m['min_vip_level']}\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetMaxVIPLevel() int16 {\n";
		$code .= "	return {$m['max_vip_level']}\n";
		$code .= "}\n";

		$code .= "func (this $mname) GetParameters() string {\n";
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

		$code .= "func (this $mname) GetAttachments() []*Attachment {\n";
		$code .= "	if attachments, ok := g_attachments[{$m['id']}]; ok {\n";
		$code .= "		return attachments\n";
		$code .= "	}\n";
		$code .= "	return this.Attachments\n";
		$code .= "}\n\n";		
	}


	$code .= generate_test_code($db, $mails);

	return $code;
}

function getAttachments($db, $mails) {

	// var attachments = map[int64][]Attachment{
	// 	1: []Attachment{Attachment{1, 2, 3}, Attachment{2, 3, 4}},
	// }

	$code = '';
	$code .= "var g_attachments =  map[int16][]*Attachment {\n";

	foreach ($mails as $key => $mail) {
		$attachments = db_query($db, "SELECT * FROM `mail_attachments` WHERE mail_id = {$mail['id']}");
		if (count($attachments) == 0) {
			continue;
		}

		$str = "	".$mail['id'].":"." []*Attachment{";

		foreach ($attachments as $key => $attachment) {
			$str .= "&Attachment{";
			$str .= $attachment['attachment_type'].",";
			$str .= $attachment['item_id'].",";
			$str .= $attachment['item_num'];
			$str .= "},";
		}
		$str .= "},";

		$code .= $str;
		$code .= "\n";
	}

	$code .= "}";

	return $code;
}

function generate_test_code($db, $mails) {
	$code = '';

	$code .= "// 只用于调试\n";
	$code .= "func NewMailTemplete(mailId int32) Mailer {\n";
	$code .= "	switch mailId {\n";
	foreach ($mails as $key => $m) {
		$parameters = getMailParas($m['parameters']);
		$code .= "	case {$m['id']}:\n";
		$code .= "		return &Mail{$m['sign']}{";
		foreach ($parameters as $key => $p) {
			switch ($p[2]) {
				case 'int32';
					$id = db_query($db, "select max(`id`) as id from {$p[3]}");
					$code .= "{$p[0]}: {$id[0]['id']}, ";
					break;
				default:
					$code .= "{$p[0]}: \"abc\", ";
					break;	
			}
		}
		$code .= "}\n";
	}
	$code .= "	}\n";
	$code .= "	return nil\n";
	$code .= "}\n\n";

	return $code;
}


function getMailParas($str) {
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
		// if (substr($t[0], strlen($t[0])-3, 3) == '_id') {
		// 	$type = 'int32';
		// 	$module = substr($t[0], 0, strlen($t[0])-3);
		// }
		array_push($parameters, array(getName($t[0]), isset($t[1]) ? $t[1] : '', $type, $module));
	}
	return $parameters;
}

function getName( $name ) {
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
