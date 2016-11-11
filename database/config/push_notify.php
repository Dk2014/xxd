<?php
$config = array(
	array( 'const',
	'UNSCHEDULED_NOTIFY' => -1,//不定期的通知
	'PAUSED_NOTIFY' => -2, //暂停播放的通知
),

//推送通知常量
array('const',),

//默认开启的推送通知ID
array('data', 'DefaultNotificationID' => '[]int32'),
);

$db = db_connect();

$push_notifications = db_query($db, 'select * from push_notify');

foreach($push_notifications as $notifycation ) {
	$config[1][$notifycation['sign']] = intval($notifycation['id']);
	if ($notifycation['sign'] == 'AfternoonPhysical' ||
		$notifycation['sign'] == 'NightPhysical' ) {
		$config[2][] = intval($notifycation['id']);
	}
}
?>
