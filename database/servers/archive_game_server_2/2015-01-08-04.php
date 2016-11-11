<?php

$this->AddSQL("

alter table `player_use_skill` change `skill_id2` `skill_id_t` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式_临时';

alter table `player_use_skill` change `skill_id3` `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式2';

alter table `player_use_skill` change `skill_id4` `skill_id3` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式3';

alter table `player_use_skill` change `skill_id_t` `skill_id4` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式4';

");

