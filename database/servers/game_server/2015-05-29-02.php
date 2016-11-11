<?php

$this->AddSQL("
	alter table mail add column  min_level smallint(6) not null default '0' comment '要求最低等级';
	alter table mail add column max_level smallint(6) not null default '0' comment '要求最高等级';
	alter table mail add column min_vip_level smallint(6) not null default '0' comment '要求VIP最低等级';
	alter table mail add column max_vip_level smallint(6) not null default '0' comment '要求VIP最等级';

	alter table global_mail add column min_level smallint(6) not null default '0' comment '要求最低等级';
	alter table global_mail add column max_level smallint(6) not null default '0' comment '要求最高等级';
	alter table global_mail add column min_vip_level smallint(6) not null default '0' comment '要求VIP最低等级';
	alter table global_mail add column max_vip_level smallint(6) not null default '0' comment '要求VIP最等级';
");


?>
