<?php


function leveltable($row) {
	return html_get_link("关卡", "?p=extend_level_config&m=".$row['id']."&m2={$row['level_type']}", false);
}


function sql_where($params) {
	$def = 1;
	if (isset($params['m'])) {
		$def = $params['m'];
	}

	return " where `level_type`={$def}";
}


function sql_insert($params) {
	return "`level_type` = {$params['m']}";
}

?>