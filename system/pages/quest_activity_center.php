<?php
	
$pconfig = array(
	'title'   => '活动中心',
	'table'   => 'quest_activity_center',
	'links'   => array(),
	
	'columns' => array(
		array('name' => 'relative', 'text' => '关联的活动', 'width' => '80px'),
		array('name' => 'sign', 'text' => '活动标识', 'width' => '80px'),
		array('name' => 'weight', 'text' => '活动权值', 'width' => '80px'),
		array('name' => 'name', 'text' => '活动名称(列表左侧)', 'width' => '120px'),
		array('name' => 'title', 'text' => '活动标题(列表右侧)', 'width' => '120px'),
		array('name' => 'content', 'text' => '活动描述', 'width' => '160px'),
		array('name' => 'start', 'text' => '活动开始时间戳', 'width' => '80px'),
		array('name' => 'end', 'text' => '活动结束时间戳', 'width' => '80px'),
		array('name' => 'dispose', 'text' => '活动过期时间戳', 'width' => '80px'),
		array('name' => 'is_go', 'text' => '是否前往', 'width' => '60px'),
		array('name' => 'tag', 'text' => '活动标签', 'width' => '60px'),
		array('name' => 'is_relative', 'text' => '相对时间', 'width' => '80px'),
		array('name' => 'is_mail', 'text' => '活动结束是否补发奖励', 'width' => '100px'),
		array('name' => 'condition_template', 'text' => '条件标题模板,{val}替换权值', 'width' => '80px'),
		array('name' => 'mail_title', 'text' => '补发奖励邮件标题', 'width' => '80px'),
		array('name' => 'mail_content', 'text' => '补发奖励邮件内容,{val}对应权值', 'width' => '80px'),
	),
	'opera' => array(
		array('type' => 'link', 'text' => '', 'render' => 'event_config_index'),
		array('type' => 'link', 'text' => '', 'render' => 'event_status'),
	),
	'where' => 'sql_where',
);

?>