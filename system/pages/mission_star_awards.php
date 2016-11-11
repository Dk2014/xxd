<?php
$pconfig = array(
    'title'   => '区域关卡评星总数',
    'table'   => 'mission_star_awards',
    'links'   => array(),
    
    'columns' => array(
        array('name' => 'totalstar', 'text' => '通关总星数', 'width' => '80px'),
    ),

    'where'         => 'sql_where',
    'insert'        => 'sql_insert',
);
?>
