<?php
$config = array(
	array('const',),
);
$db = db_connect();
$privileges = db_query($db, 'select * from `func` where `type`=2 order by `level`');
foreach($privileges as $p) {
	$config[0][$p['sign']] = intval($p['level']);
}
?>
