<?php

include "clique_buildings_link.php";

$pconfig = array(
    'title'   => '宗祠建筑升级',
    'table'   => 'clique_temple_upgrade',  
    'links'   => $clique_buildings_link,

    'columns' => array(
        array('name' => 'level', 'text' => '宗祠等级', 'width' => '80px'),
        array('name' => 'upgrade_fee', 'text' => '升级费用(铜钱)', 'width' => '80px'),
        array('name' => 'desc', 'text' => '宗祠描述', 'width' => '200px'),
        array('name' => 'cur_temple_desc', 'text' => '当前宗祠等级描述', 'width' => '200px'),
        array('name' => 'next_temple_desc', 'text' => '下一宗祠等级描述', 'width' => '200px'),
    ),
    
    'opera' => array(
    ),
);


?>
