<?php 

$this->AddSQL("
	alter table player_item add column `price` int(11) not null default '0' comment '装备精炼价格';
");
?>
