<?php

$pconfig = array (
        'title' => '队伍友情数据',
        'table' => 'teamship',
        'links' => array(),
        'columns' => array (
                array (
                        'name' => 'level',
                        'text' => '团队单项加成等级',
                        'width' => '30px' 
                ),
                array (
                        'name' => 'needs_relationship',
                        'text' => '升级所需友情点数',
                        'width' => '30px'
                ),
                array (
                        'name' => 'health',
                        'text' => '生命项加成',
                        'width' => '30px'
                ),
                array (
                        'name' => 'attack',
                        'text' => '攻击项加成',
                        'width' => '30px',
                ),
                array (
                        'name' => 'defence',
                        'text' => '防御项加成',
                        'width' => '30px'
                ),
        ),
        'where' => 'sql_where',
);

function sql_where($params) {
	return "order by `level`";
}

