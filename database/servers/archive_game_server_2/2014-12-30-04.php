<?php 
$this->AddSQL("

DROP TABLE ghost_skill_force;
CREATE TABLE  ghost_skill_force (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`ghost_id` smallint(6) NOT NULL COMMENT '魂侍ID',
	`level` smallint(6) NOT NULL COMMENT '等级',
	`force` int(11) NOT NULL COMMENT '威力',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍技能威力';


");
?>
