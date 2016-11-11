<?php
    $this->AddSQL("
        alter table `clique_temple_upgrade`  add  column   `cur_temple_desc` text COMMENT '当前武功描述';
        alter table `clique_temple_upgrade`  add  column   `next_temple_desc` text COMMENT '下一等级武功描述';
    ");
?>