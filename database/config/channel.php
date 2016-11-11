<?php
$config = array(
	//!!!消息类型常量 请勿移动!!!
	array( 'const', ),
	//!!!消息类型常量 请勿移动!!!
	
	array('const',
		'WORLD_CHAT_CD_TIME' => 5, //世界聊天频道CD时间
		'WORLD_CHAT_COST' => 5, //付费的世界聊天单次价格
		'WORLD_CHAT_MAX_CONTENT_LEN' => 60, //最大长度
		'WORLD_CHAT_DAILY_FREE_NUM' => 5, //世界聊天每日免费次数
		'WORLD_CHAT_SERVER_OPEN_LEVEL' => 24, //可以发送世界聊天消息的等级，因为互动服的等级数据不及时，所以允许24级的时候发送。
		'WORLD_CHAT_CLIENT_BUFF_SIZE' => 50, //客户端最大保留聊天数量
		'WORLD_CHAT_CLIENT_PAGE_SIZE' => 20, //客户端每页显示聊天数量
		'WORLD_CHAT_CLIENT_BUFF_EXPIRE_TIME' => 60 * 10, //客户端在进入后台10分钟后原来的聊天记录过期
	),

	array('const',
		'PARAM_TYPE_STRING' => 1,
		'PARAM_TYPE_ITEM' => 2,
		'PARAM_TYPE_PLAYER' => 3,
		'PARAM_TYPE_CLIQUE' => 4,
		'PARAM_TYPE_CLIQUE_BOAT' => 5,
	),
);

$db = db_connect();
$messages = db_query($db, 'select * from world_channel_message ');
foreach($messages as $msg) {
	 $config[0][ 'MESSAGE_TPL_' . $msg['sign']] = intval($msg['id']);
}

?>
