
<?php
db_execute($db, 
"
		alter table `player_mail` change column `hava_attachment` `have_attachment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有附件';
"
);

?>

