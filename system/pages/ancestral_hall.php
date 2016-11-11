<?php

include "ancestral_hall_link.php";

$pconfig = array(
    'title'   => '宗祠',
    'table'   => 'clique_temple',  
    'links'   => $links,

    'columns' => array(
        array('name' => 'worship_type', 'text' => '上香类型', 'width' => '80px'),
        array('name' => 'level', 'text' => '宗祠等级', 'width' => '80px'),
        array('name' => 'fame', 'text' => '声望', 'width' => '80px'),
        array('name' => 'contrib', 'text' => '帮贡', 'width' => '80px'),
    ),
    
    'opera' => array(
    ),
    
    'where'         => 'sql_where',
);


function sql_where($params) {
    if(isset($_GET['m'])){
        return " where `worship_type`={$_GET['m']} ";
    }
    return " order by `id`, `worship_type` ASC";
}

?>
