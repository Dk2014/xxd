<?php
db_execute($db, "

create table `player_global_mail_state`(
	  `pid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家ID',
	  `max_timestamp` bigint(20) NOT NULL COMMENT '发送时间戳',
PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家全局邮件记录';

drop table `global_mail_pending_queue`;
drop table `player_mail_state`;

alter table `global_mail_attachments` change column `attach_ref` `global_mail_id` bigint(20) NOT NULL COMMENT '全局邮件表主键';
alter table `global_mail` drop column `attach_ref`;


");

?>
