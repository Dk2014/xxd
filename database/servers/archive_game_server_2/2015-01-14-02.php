<?php
$this->AddSQL("ALTER TABLE `events_level_up` CHANGE `coin` `coins` int(11) DEFAULT 0 COMMENT '奖励铜钱';");

?>
