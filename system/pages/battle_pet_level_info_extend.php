<?php

function sql_where($params) {
	if(!isset($params['m'])) {
		die("need params `m`");
	}
	return " where `pet_id`={$params['m']} order by `level` asc";
}

function sql_insert($params) {
    if(!isset($params['m'])) {
        die("need params `m`");
    }
    return " `pet_id`={$params['m']} ";
}
