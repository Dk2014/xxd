
<?php
db_execute($db, 
"

ALTER TABLE `quest` CHANGE `desc` `desc` varchar(240) DEFAULT '' COMMENT '简介';

"

);
?>