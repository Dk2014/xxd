<?php 

$this->AddSQL("
       drop table if exists  player_sealedbook;
");

$this->AddSQL("

CREATE TABLE IF NOT EXISTS  `player_sealedbook` (
 `pid` bigint(20) NOT NULL COMMENT '用户ID',
  `sealed_record` mediumblob COMMENT '玩家天书记录',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家天书记录';
");

?>




