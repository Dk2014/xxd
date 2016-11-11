<?php

$links = array(
  // => array( '1'=>'宗祠','2'=>'回春堂','3'=>'神兵堂', '4'=>'金刚堂','5'=>'总舵','6'=>'钱庄', ),
    array('text' => '全部建筑任务', 'id' => 'clique_building_quest'),
    array('text' => '宗祠', 'id' => 'clique_building_quest&m=1'),
    array('text' => '回春堂', 'id' => 'clique_building_quest&m=2'),
    array('text' => '神兵堂', 'id' => 'clique_building_quest&m=3'),
    array('text' => '金刚堂', 'id' => 'clique_building_quest&m=4'),    
    array('text' => '总舵', 'id' => 'clique_building_quest&m=5'),
    array('text' => '钱庄', 'id' => 'clique_building_quest&m=6'),
    array('text' => '每日任务', 'id' => 'clique_daily_quest'),
    );
$pconfig = array(
    'title'   => '帮派建筑任务',
    'table'   => 'clique_building_quest',
    'links'   =>$links,
    
    'columns' => array(
        array('name' => 'order', 'text' => '显示顺序', 'width' => '80px'),
        array('name' => 'name', 'text' => '任务标题', 'width' => '80px'),
       
        array('name' => 'require_building_level', 'text' => '要求最低等级', 'width' => '50px'),
    
        array('name' => 'class', 'text' => '类别', 'width' => '50px'),
     
        array('name' => 'award_clique_contri', 'text' => '奖励帮贡','width' => '50px'),
        array('name' => 'award_exp', 'text' => '奖励经验', 'width' => '80px'),
        array('name' => 'award_ingot', 'text' => '奖励元宝', 'width' => '80px'),
        array('name' => 'award_coins', 'text' => '奖励铜钱', 'width' => '80px'),
        array('name' => 'award_physical', 'text' => '奖励体力', 'width' => '80px'),
        array('name' => 'award_item1_id', 'text' => '奖励物品1', 'width' => '80px'),
        array('name' => 'award_item1_num', 'text' => '奖励物品1数量', 'width' => '80px'),
        array('name' => 'award_item2_id', 'text' => '奖励物品2', 'width' => '80px'),
        array('name' => 'award_item2_num', 'text' => '奖励物品2数量', 'width' => '80px'),
        array('name' => 'award_item3_id', 'text' => '奖励物品3', 'width' => '80px'),
        array('name' => 'award_item3_num', 'text' => '奖励物品3数量', 'width' => '80px'),
        array('name' => 'award_item4_id', 'text' => '奖励物品4', 'width' => '80px'),
        array('name' => 'award_item4_num', 'text' => '奖励物品4数量', 'width' => '80px'),
        array('name' => 'desc', 'text' => '简介', 'width' => '180px'),
    ),


    'where'         => 'sql_where',
    'js'            => 'jsFunction',

);

?>