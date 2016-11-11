<?php


function sql_where($params) {
    return "where `mission_id`={$params['m']}   ";
}

function sql_insert($params) {
    return "`mission_id` = {$params['m']}";
}

?>
