<?php
$this->AddSQL("
CREATE TABLE IF NOT EXISTS `events_ingot_record`(
`pid` bigint(20) NOT NULL AUTO_INCREMENT,
`ingot_in` bigint(20) COMMENT '充值元宝总数',
`ingot_in_end_time` bigint(20) COMMENT '累计充值活动结束时间戳',
`ingot_out` bigint(20) COMMENT '消耗元宝总数',
`ingot_out_end_time` bigint(20) COMMENT '消耗元宝活动结束时间戳',
primary key(`pid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家元宝充值和消耗活动记录';
");