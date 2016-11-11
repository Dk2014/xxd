<?php

function location($params){
	global $db;
	$html = '';
	$sql = "select `id`, `name`,`type`,`role_id` from `skill` where id = {$params['m']}";
	$role = db_query($db, $sql);
	if (count($role) == 0) {
		return $html;
	}

	if ($role[0]['type']==1){
		if ($role[0]['role_id']==-1){
			$html = '<a href="?p=skill">'.$role[0]['name'].'</a> -> ' . $html;
		}else {
			$html = '<a href="?p=skill_buddy">'.$role[0]['name'].'</a> -> ' . $html;
		}
	}else if ($role[0]['type']==3) {
		$html = '<a href="?p=skill_role_cultivation">'.$role[0]['name'].'</a> -> ' . $html;
	}else if ($role[0]['type']==4) {
		$html = '<a href="?p=skill_role_dogge">'.$role[0]['name'].'</a> -> ' . $html;
	}
	

	return $html;
}

function sql_insert($params) {
	return "`skill_id` = {$params['m']}";
}

function sql_where($params) {
	if (isset($params['m']))
		$where = " where `skill_id`={$params['m']}";
	return $where;
}

$pconfig = array(
	'title'       => '绝招数值',
	'table'       => 'skill_content',
	'location'    => 'location',
	'links'       => array(),
	'columns'     => array(
		array('name' => 'release_num', 'text' => '释放次数','width' => '100px'),
	),
	'where'   => 'sql_where',
	'insert'  => 'sql_insert',
);
?>
