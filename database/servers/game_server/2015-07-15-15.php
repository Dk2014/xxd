<?php
$this->AddSQL("
	alter table `player` add column cid bigint(20) not null default 0 comment '渠道ID';");
?>