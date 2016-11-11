<?php


$this->AddSQL("
    ALTER TABLE `level_star` add  COLUMN `two_star_round` smallint(6) DEFAULT 0 COMMENT '2星回合数';
    ALTER TABLE `level_star` add  COLUMN `three_star_round` smallint(6) DEFAULT 0 COMMENT '3星回合数';
");

?>



