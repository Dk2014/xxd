<?php
$extend_columns = array();

function sql_where($params) {
	return "where trader_id={$params['m']}";
}

function sql_insert($params) {
	return "trader_id={$params['m']}";
}

?>
