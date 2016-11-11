<?php

$this->AddSQL("
alter table `player_use_skill` add index `pid`(`pid`);
");

