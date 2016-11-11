<?php

//把所有装备位置为没有装备魂侍的状态
$this->AddSQL("
update `player_ghost_equipment` set `pos1`=0, `pos2`=0, `pos3`=0, `pos4`=0;
");

//把所有魂侍置为没有装备的状态
$this->AddSQL("
update `player_ghost` set `pos`=0;
");

