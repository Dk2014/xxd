<?php

$this->AddSQL("
delete from player_daily_quest where `class` in (18, 10);
");

