<?php 

$this->AddSQL("
    alter table player_mission_star_award change column `award_level` `box_type`  tinyint(4) NOT NULL DEFAULT '0' COMMENT '宝箱类型 1:铜 2:银 3:金';
        ");

?>
