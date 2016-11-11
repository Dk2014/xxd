<?php
$config = array(
	array( 'const',
		'MODULE_ANNOUNCEMENT' => 0, //模块公告，滚动展示
		'BACKSTAGE_ANNOUNCEMENT' => 1, //后台公告，选择角色时展示
	),

	// 静态公告数据类型
	array( 'type', 'announcement_config',
		'tpl_id'  => 'int64',    // 公告模版ID
		'params'   => '[]string',   // 参数
		'show_timing' => 'int64' // 定期展示时间描述一天内第X秒开始展示这条公告，如果X小于零则登录展示一次
	),

	// 静态公告数据
	array( 'data', 'static_announcements' => '[]announcement_config',
		//array(1, array("测试公告下午五点半滚动播放"), 63000),
	)
);
?>
