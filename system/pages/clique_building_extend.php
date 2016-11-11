<?php

// 所有建筑物的标识数组 id => biaoshi
$biaoshis = array(
	1 => 'base', // 总舵
	2 => 'bank', // 钱庄
	);

function descr($row){
	global $biaoshis;
	return '<a id="descr_link_'.$row['id'].'" href="javascript:;" onclick="javascript:descrEditor('.$row['id'].');" >描述配置</a>'.other_operate($biaoshis[$row['id']]);
}

function foot(){
	require dirname(__FILE__).'/clique_building_desc_editor.php';
}

function other_operate($biaoshi){
	return ' | <a href="/index.php?p=clique_building_'.$biaoshi.'">数据模板配置</a>';
}

?>