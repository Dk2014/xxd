<?php

$this->AddSQL("
alter table player_info add column `sword_soul_fragment` bigint(20) not null default '0' comment '剑心碎片数量';
update item set type_id ='9' where type_id = '14';
");

