<?php

$this->AddSQL("
ALTER TABLE `global_clique`  add column `worship_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上香时间';
alter table `global_clique` add column `worship_cnt` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '当日上香次数';
");


$this->AddSQL("
CREATE TABLE IF NOT EXISTS `ancestral_hall`(
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '宗祠等级',
  `worship_type` tinyint(4)  NOT NULL DEFAULT '0' COMMENT  '上香类型,1:白檀香 2:苏合香 3：天木香 ',
  `fame` smallint(6) NOT NULL DEFAULT '0' COMMENT '声望',
  `contrib` int(11) NOT NULL DEFAULT '0' COMMENT '帮贡',
  `upgrade_fee` int(11) NOT NULL DEFAULT '0' COMMENT '升级费用（铜钱）',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宗祠';
");

?>
