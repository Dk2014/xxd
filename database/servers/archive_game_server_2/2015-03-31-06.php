<?php 

$this->AddSQL("
        alter table town_star_awards add column `box_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '宝箱类型 1:铜 2:银 3:金';
        ");
?>
