<?php
include_once "common_links.php";

$pconfig = array(
        'title'   => '剑山拔剑概率',
        'table'   => 'sword_soul_probability',
	    'links'   => $sword_soul_links,
        'columns' => array(
                array('name' => 'box_name', 'text' => '箱子名称', 'width' => '120px'),
                array('name' => 'level_up', 'text' => '升箱概率', 'width' => '120px'),
                array('name' => 'get_exp', 'text' => '获得潜龙概率', 'width' => '120px'),
                array('name' => 'get_rubbish', 'text' => '获得杂物概率', 'width' => '90px'),
                array('name' => 'green', 'text' => '优秀剑心概率', 'width' => '90px'),
                array('name' => 'blue', 'text' => '精良剑心概率', 'width' => '90px'),
                array('name' => 'purple', 'text' => '传奇剑心概率', 'width' => '90px'),
                array('name' => 'yello', 'text' => '神器剑心概率', 'width' => '90px'),
        ),
        'where' => 'sql_where',
        'not_delete' => true,
        'not_new' => true,
);
