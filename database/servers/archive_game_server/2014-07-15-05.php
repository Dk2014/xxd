<?php
db_execute($db,"
alter table `func` modify column `need_play` tinyint(4) DEFAULT '0' COMMENT '是否需要播放 0不需要 1需要';
"
);
?>
