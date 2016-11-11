<?php 


$this->AddSQL("

CREATE TABLE IF NOT EXISTS  `player_sealedbook` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
   `sealed_record` mediumblob COMMENT '玩家天书记录',
   PRIMARY KEY (`id`),
   KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家天书记录';
");

?>




