<?php

$this->AddSQL("
ALTER TABLE `player_formation`
	ADD `pos6` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位6',
	ADD `pos7` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位7',
	ADD `pos8` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位8';
");

