<?php
    
$pconfig = array(
    'title'   => '活动配置Json文件',
    'eventsjson' => '../json_events.json',
    'links'   => array(),
    
    'columns' => array(
        array('name' => 'Type', 'text' => '活动类型', 'width' => '80px'),
        array('name' => 'StartUnixTime', 'text' => '活动开始时间戳', 'width' => '80px'),
        array('name' => 'EndUnixTime', 'text' => '活动结束时间戳', 'width' => '80px'),
        array('name' => 'DisposeUnixTime', 'text' => '活动过期时间戳', 'width' => '80px'),
       
        array('name' => 'Page', 'text' => '活动期数', 'width' => '80px'),
        
        array('name' => 'IsRelative', 'text' => '关联的活动', 'width' => '80px'),
        array('name' => 'Weight', 'text' => '活动权值', 'width' => '80px'),
        array('name' => 'Tag', 'text' => '活动标识', 'width' => '80px'),        
        array('name' => 'LTitle', 'text' => '活动名称(列表左侧)', 'width' => '160px'),
        array('name' => 'RTitle', 'text' => '活动标题(列表右侧)', 'width' => '160px'),
        array('name' => 'Content', 'text' => '活动描述', 'width' => '160px'),
    ),
    
    'opera' => array(
        array('type' => 'link', 'text' => '参数配置', 'render' => 'version_opera'),
    ),

    'foot'          => 'foot',
    'not_delete'    => true,
);

?>