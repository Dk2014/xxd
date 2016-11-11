<?php

include "sealed_links.php";

$pconfig = array(
    'title'   => '天书配置',
    'table'   => 'sealed_book',
    'links'   => $links,
    
    'columns' => array(
        array('name' => 'type', 'text' => '类型', 'width' => '80px'),
        array('name' => 'item_id', 'text' => '物品', 'width' => '80px'),
        array('name' => 'health', 'text' => '生命', 'width' => '80px'),
        array('name' => 'attack', 'text' => '攻击', 'width' => '80px'),
        array('name' => 'defense', 'text' => '防御', 'width' => '80px'),
        array('name' => 'cultivation', 'text' => '内力', 'width' => '80px'),
        array('name' => 'coins', 'text' => '消耗铜钱', 'width' => '80px'),
    ),
    
    'where'         => 'sql_where',
    'js'            => 'js_function',
);


?>
