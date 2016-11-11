<?php
$config = array(
	/*好友关系*/
	array('const',
		'FRIEND_MODE_STRANGE'      	=>0,
		'FRIEND_MODE_LISTEN_ONLY'   =>1,
		'FRIEND_MODE_LISTENED_ONLY' =>2,
		'FRIEND_MODE_FRIEND'        =>3,
		'FRIEND_MODE_NULL'          =>-1,
	),

	/*好友*/
	array( 'const',
    'FRIEND_MAX_NUM' => 50, //好友最大数量
    'FRIEND_SEND_HEART_INTERVAL' => 24*3600, //好友送爱心时间间隔
    'FRIEND_DELETE_MAX_DAY_COUNT' => 10, //每日最大删除好友数量
	),

	/*聊天*/
	array('const',
	'CHAT_MESSAGE_MODE_SEND' => 1, //消息类型,发送
	'CHAT_MESSAGE_MODE_RECEIVE' => 2, //消息类型,接收
	'CHAT_MESSAGE_STORE_TIME' => 7*24*3600, //消息记录保存时间
	),
); 
?>
