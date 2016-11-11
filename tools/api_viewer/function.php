<?php

function strStartUper($str) {
    return str_replace(array('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'), array('A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'), substr($str, 0, 1)) . substr($str, 1);
}

function emptydir($dirPath) {
	$dirFiles=scandir($dirPath);
	foreach ($dirFiles as $dirFile) {
		$dirFilePath = $dirPath."/".$dirFile;
		if (filetype($dirFilePath) == "file" ) {
			unlink($dirFilePath);
		}
	}
}

//html页面相关函数
function html_creat_table($array) {
    $outarray = html_new_node("table", array());
    if (isset($array["key"])) {
        $tr = html_new_node("tr", array());
        foreach ($array['key'] as $k => $v) {
            
            $td = html_new_node("td", array());
            html_add_text($td, isset($array['title'][$k]) ? $array['title'][$k] : $v);
            html_add_node($tr, $td);
            
            //printval($tr);
            
        }
        html_add_node($outarray, $tr);
        foreach ($array['table'] as $v) {
            $tr = html_new_node("tr", array());
            foreach ($array['key'] as $v1) {
                $td = html_new_node("td", array());
                if (is_array($v[$v1])) {
                    
                    switch ($v[$v1]['key']) {
                        case "img":
                            html_add_new_node($td, "img", array("src" => $v[$v1]["src"]));
                            break;

                        case "a":
                            html_add_new_node($td, "a", array("href" => $v[$v1]["href"]));
                            html_add_text($td, $v[$v1]["text"]);
                            break;

                        case "node":
                            html_add_node($td, $v[$v1]["node"]);
                            break;
                    }
                } else {
                    
                    //printval($v[$v1]);
                    html_add_text($td, $v[$v1]);
                }
                html_add_node($tr, $td);
            }
            html_add_node($outarray, $tr);
        }
    }
    return $outarray;
}
function html_new_node($key, $attr = array(), $text = "") {
    $temp = array("key" => $key);
    foreach ($attr as $k => $v) {
        $temp['attr'][$k] = $v;
    }
    if ($text != "") {
        $temp["subnode"][] = array("key" => "text", "text" => $text);
    }
    return $temp;
}
function html_add_node(&$node, $subnode) {
    
    $out = & $node['subnode'][];
    $out = $subnode;
    return $out;
}

function html_add_new_node(&$node, $key, $attr = array(), $text = "") {
    $temp = array("key" => $key);
    foreach ($attr as $k => $v) {
        $temp['attr'][$k] = $v;
    }
    if ($text != "") {
        $temp["subnode"][] = array("key" => "text", "text" => $text);
    }
    $temp2 = & $node['subnode'][];
    $temp2 = $temp;
    return $temp2;
}
function html_add_text(&$node, $text) {
    while (strpos($text, "\n") > 0) {
        $txt = substr($text, 0, strpos($text, "\n"));
        $text = substr($text, strpos($text, "\n") + 1);
        $node["subnode"][] = array("key" => "text", "text" => $txt);
        $node["subnode"][] = array("key" => "br");
    }
    $node["subnode"][] = array("key" => "text", "text" => $text);
}
function html_creat_file($file, $html) {
    $xml = new XMLWriter();
    $xml->openURI($file);
    $xml->setIndentString('  ');
    $xml->setIndent(true);
    
    //开始创建文档并设置DTD
    $xml->startDocument('1.0', 'UTF-8');
    $xml->startDtd('html', '-//W3C//DTD XHTML 1.0 Strict//EN', 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd');
    $xml->endDtd();
    
    //创建HTML文档
    $xml->startElement('html');
    $xml->writeAttribute('xmlns', 'http://www.w3.org/1999/xhtml');
    $xml->writeAttribute('xml:lang', 'en');
    $xml->writeAttribute('lang', 'en');
    $xml->startElement('head');
    $xml->writeElement('title', '神仙道场景汇总');
    $xml->endElement();
    $xml->startElement('body');
    foreach ($html as $v) {
        creat_tag($xml, $v);
    }
    $xml->endElement();
}
function html_creat($html) {
    $xml = new XMLWriter();
    $xml->openURI("php://output");
    $xml->setIndentString('  ');
    $xml->setIndent(true);
    
    //开始创建文档并设置DTD
    $xml->startDocument('1.0', 'UTF-8');
    foreach ($html as $v) {
        creat_tag($xml, $v);
    }
    $xml->endElement();
}
function creat_tag(&$xml, $val) {
    switch ($val["key"]) {
        case "text":
            $xml->text($val['text']);
            break;

        default:
            if (isset($val['key'])) {
                $xml->startElement($val['key']);
                if (isset($val['attr']) && is_array($val['attr'])) {
                    foreach ($val['attr'] as $k => $v) {
                        $xml->writeAttribute($k, $v);
                    }
                }
                if (isset($val['subnode']) && is_array($val['subnode'])) {
                    foreach ($val['subnode'] as $v) {
                        creat_tag($xml, $v);
                    }
                }
                $xml->endElement();
            }
            break;
    }
}
