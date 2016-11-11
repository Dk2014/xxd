<?php

$arr = array('40','60','80','90','100');
foreach ($arr as $val) {
	
	db_execute($db, "
		delete from mission_level where parent_type = 1 and parent_id = (select id from resource_level where max_level = {$val});
		delete from resource_level where max_level = {$val};
	");

}
?>