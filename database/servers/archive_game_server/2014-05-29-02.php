
<?php
db_execute($db, 
"
		alter table `mail_attachments` change column `item_type` `attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型';
		alter table `player_mail_attachment` change column `item_type` `attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型';
"
);

?>

