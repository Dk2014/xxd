<?php

$this->AddSQL("
update `player_skill` set `skill_trnlv` = 1 where `skill_trnlv` < 1;
");

