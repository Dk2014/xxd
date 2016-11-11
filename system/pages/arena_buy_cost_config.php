<?php
/**
 * Created by PhpStorm.
 * User: tangweichen
 * Date: 14-11-18
 * Time: 下午3:24
 */
$pconfig = array(
    'title'   => '元宝购买比武场次数所消耗元宝配置',
    'table'   => 'arena_buy_cost_config',
    'links'   => array(),

    'columns' => array(
        array('name' => 'times', 'text' => '购买次数', 'width' => '100px'),
        array('name' => 'cost', 'text' => '购买所需元宝', 'width' => '100px'),

    ),
    'where' 		=> 'sql_where',
);

?>