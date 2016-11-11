 <?php
db_execute($db, "
alter table `skill` modify column `required_level` int(11) NOT NULL DEFAULT '0' COMMENT '需要等级';
 ");
?>
