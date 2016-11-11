<?php
db_execute($db, "

alter table `global_mail_attachments` modify column `item_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '数量';


");

?>
