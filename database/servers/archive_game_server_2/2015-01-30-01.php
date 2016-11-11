<?php

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `player_new_year_consume_record`(
`pid` bigint(20) NOT NULL COMMENT '玩家id',
`consume_record` text COMMENT '玩家消费情况',
PRIMARY KEY(`pid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='春节玩家消费记录';
");