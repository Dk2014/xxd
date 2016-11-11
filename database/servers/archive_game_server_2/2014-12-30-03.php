<?php 
$this->AddSQL("
ALTER TABLE ghost_star ADD COLUMN `quality` tinyint(4) NOT NULL DEFAULT '0'  COMMENT '品质';
ALTER TABLE ghost_star drop column `color`; 

ALTER TABLE ghost_star ADD COLUMN `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命';
ALTER TABLE ghost_star ADD COLUMN `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击';
ALTER TABLE ghost_star ADD COLUMN `defence` int(11) NOT NULL DEFAULT '0' COMMENT '防御';

ALTER TABLE ghost_passive_skill ADD COLUMN `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级';
ALTER TABLE ghost_passive_skill DROP COLUMN `level`;

ALTER TABLE skill DROP COLUMN `require_ghost_star`;

ALTER TABLE ghost drop column `speed`; 

ALTER TABLE player_ghost add column skill_level smallint(6) NOT NULL DEFAULT '1' COMMENT '技能等级';

CREATE TABLE  ghost_skill_force (
	`id` int(11) NOT NULL COMMENT 'ID',
	`ghost_id` smallint(6) NOT NULL COMMENT '魂侍ID',
	`level` smallint(6) NOT NULL COMMENT '等级',
	`force` int(11) NOT NULL COMMENT '威力',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍技能威力';

CREATE TABLE ghost_skill_train_price(
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
	`level` smallint(6) NOT NULL COMMENT '等级',
	`cost` int(11) NOT NULL COMMENT '费用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍技能训练价格';

");
?>
