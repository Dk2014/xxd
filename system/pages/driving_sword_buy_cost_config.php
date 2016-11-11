<?php

$pconfig = array(
    'title'   => '云海御剑行动点购买次数所消耗元宝配置',
    'table'   => 'driving_sword_buy_cost_config',
    'links'   => array(),

    'columns' => array(
        array('name' => 'times', 'text' => '购买次数', 'width' => '100px'),
        array('name' => 'cost', 'text' => '购买所需元宝', 'width' => '100px'),

    ),
    'where' 		=> 'sql_where',
);

?>