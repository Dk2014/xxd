<?php
$config = array(
	array('const', 
		'MAX_VIP_LEVEL' => 15,
	),

	//VIP特权
    array('const',
),
);
$db = db_connect();
$privileges = db_query($db, 'select * from vip_privilege order by `order`');
foreach($privileges as $p) {
	$config[1][$p['sign']] = intval($p['id']);
}
?>
