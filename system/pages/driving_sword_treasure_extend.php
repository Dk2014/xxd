<?php

function sql_where($params) {	
	return " where `cloud_id`={$params['m']} order by `event_id` asc ";
}

function sql_insert($params) {
	return " `cloud_id`={$params['m']}";
}

