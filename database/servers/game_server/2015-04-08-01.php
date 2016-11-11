<?php 


$this->AddSQL("
alter table player_role drop column `coop_role_id`; 
alter table player_role add column `coop_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '协力组合';

");

?>




