<?php

$pconfig = array (
        'title' => '角色羁绊数据',
        'table' => 'role_friendship',
        'location' => 'location',
        'links' => array(),
        'columns' => array (
                array (
                        'name' => 'friendship_level',
                        'text' => '羁绊等级',
                        'width' => '30px' 
                ),
                array (
                        'name' => 'required_role_level',
                        'text' => '所需角色最小等级',
                        'width' => '30px'
                ),
                array (
                        'name' => 'favourite_item',
                        'text' => '喜好品',
                        'width' => '60px'
                ),
                array (
                        'name' => 'favourite_count',
                        'text' => '喜好品需求量',
                        'width' => '60px'
                ),
                array (
                        'name' => 'level_color',
                        'text' => '名称颜色',
                        'width' => '60px',
                        'render' => 'render_level_color'
                ),
                array (
                        'name' => 'display_graph',
                        'text' => '资源标识',
                        'width' => '60px'
                ),
                array (
                        'name' => 'relationship_name',
                        'text' => '羁绊名称',
                        'width' => '60px'
                ),
                array (
                        'name' => 'health',
                        'text' => '生命',
                        'width' => '60px' 
                ),
                array (
                        'name' => 'attack',
                        'text' => '攻击',
                        'width' => '60px' 
                ),
                array (
                        'name' => 'defend',
                        'text' => '防御',
                        'width' => '60px' 
                ),
                array (
                        'name' => 'cultivation',
                        'text' => '内力',
                        'width' => '60px' 
                )
        ),
        'where' => 'sql_where',
        'insert' => 'sql_insert',
        'js' => 'js_function',

);
