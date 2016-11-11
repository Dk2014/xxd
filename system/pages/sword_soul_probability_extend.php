<?php

function sql_where($params) {
    return ' WHERE `id` in (SELECT max(`id`) FROM `sword_soul_probability` GROUP BY `order`) ORDER BY `order`';
}
