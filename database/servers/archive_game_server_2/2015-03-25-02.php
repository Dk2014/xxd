<?php

$this->AddSQL("
    alter table `player_extend_level` add `max_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '通关了的最大等级';
");