<?php
/**
 * Created by PhpStorm.
 * User: tangweichen
 * Date: 14-11-18
 * Time: 下午5:51
 */
$pconfig = array(
    'title'   => '彩虹关卡购买扫荡次数所消耗元宝配置',
    'table'   => 'rainbow_buy_cost_config',
    'links'   => array(),

    'columns' => array(
        array('name' => 'times', 'text' => '购买次数', 'width' => '100px'),
        array('name' => 'cost', 'text' => '购买所需元宝', 'width' => '100px'),

    ),
    'where' 		=> 'sql_where',
);

?>