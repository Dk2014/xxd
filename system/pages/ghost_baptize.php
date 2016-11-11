<?php
include "common_links.php";

function sql_where($params) {
    return " order by `star`";
}
$pconfig = array(
    'title'   => '',
    'table'   => 'ghost_baptize',
    'links'   => array(),

    'columns' => array(
        array('name' => 'star', 'text' => '魂侍星级', 'width' => '80px'),
        array('name' => 'max_add_growth', 'text' => '最大增加成长值', 'width' => '80px'),
        array('name' => 'probablity1', 'text' => '概率1', 'width' => '80px'),
        array('name' => 'min_add_growth1', 'text' => '最小随机添加成长值1', 'width' => '80px'),
        array('name' => 'max_add_growth1', 'text' => '最大随机添加成长值1', 'width' => '80px'),
        array('name' => 'probablity2', 'text' => '概率2', 'width' => '80px'),
        array('name' => 'min_add_growth2', 'text' => '最小随机添加成长值2', 'width' => '80px'),
        array('name' => 'max_add_growth2', 'text' => '最大随机添加成长值2', 'width' => '80px'),
    ),
    'where'         => 'sql_where',
);
?>
