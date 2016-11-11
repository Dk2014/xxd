<?php
//
// 连接当前目标数据库
//
function db_connect()
{
	global $db_config;

	$host = $db_config/*[$target]*/['host'];
	$user = $db_config/*[$target]*/['user'];
	$pass = $db_config/*[$target]*/['pass'];
	$name = $db_config/*[$target]*/['name'];
	$port = $db_config/*[$target]*/['port'];
	
	$db = new mysqli($host, $user, $pass, $name, $port);

	if ($db->connect_error) {
		die('连接数据库失败：Error '.$db->connect_errno.', '.$db->connect_error);
	}
	
	$db->query("set names utf8mb4");
	
	return $db;
}

//
// 执行非查询操作
//
function db_execute($db, $sql) {
	if(!$db->multi_query($sql)){
		die('Query Error (' . $db->errno . ') ' . $db->error);
	}

	db_free_result($db);
}

//
// 执行数据查询，并返回字段名索引的数据
//
function db_query($db, $sql) {
	$result = $db->query($sql, MYSQLI_USE_RESULT);

	if (!$result) {
		die('Query Error (' . $db->errno . ') ' . $db->error);
	}

	$rows = array();

	while ($row = $result->fetch_array(MYSQLI_ASSOC)) {
		$rows[] = $row;
	}

	$result->close();

	return $rows;
}

//
// 执行数据库查询，并返回顺序索引的数据
//
function db_query_array($db, $sql) {
	$result = $db->query($sql, MYSQLI_USE_RESULT);

	if (!$result) {
		die('Query Error (' . $db->errno . ') ' . $db->error);
	}
	
	$rows = array();

	while ($row = $result->fetch_array(MYSQLI_NUM)) {
		$rows[] = $row;
	}

	$result->close();

	return $rows;
}

//
// 释放所有的返回结果
//
function db_free_result($db){
	do {
		if ($result = $db->store_result()) {
			$result->free();
		}

		if (! $db->more_results()) {
			break;
		}
	} while ($db->next_result());
}
function db_insert_transaction($db, $sqls){
	$db->autocommit(0);
	foreach ($sqls as $sql) {
		if(!$db->query($sql)){
			$error = sprintf("Error : %s",mysqli_error($db));
			$db->rollback();
			return $error;
		}
	}
	$db->autocommit(1);
	$db->close();
	return '';
}
?>