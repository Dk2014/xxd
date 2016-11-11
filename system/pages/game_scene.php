<?php

$pconfig = array(
        'title'   => '剧情',
        'table'   => 'game_scene',
        'links'   => array(),
        'columns' => array(
                array('name' => 'order', 'text' => '排序', 'width' => '120px',),
                array('name' => 'name', 'text' => '名称', 'width' => '150px',),
                array('name' => 'quest_order', 'text' => '关联任务', 'width' => '150px',),
        ),
        'opera'   => array(),
        'where'	  => 'sql_where',
        'js'      => 'js_function',
        
);
