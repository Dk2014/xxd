<?php

   $this->AddSQL(" 
            drop table if exists ancestral_hall;
    ");

    $this->AddSQL(" 
    CREATE TABLE IF NOT EXISTS `clique_temple` (
      `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
      `worship_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上香类型,1:白檀香 2:苏合香 3：天木香 ',
      `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '宗祠等级',
      `fame` smallint(6) NOT NULL DEFAULT '0' COMMENT '声望',
      `contrib` int(11) NOT NULL DEFAULT '0' COMMENT '帮贡',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='宗祠';
    ");

    $this->AddSQL("
    CREATE TABLE IF NOT EXISTS `clique_temple_upgrade` (
      `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
     `upgrade_fee` int(11) NOT NULL DEFAULT '0' COMMENT '升级费用（铜钱）',
     `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '宗祠等级',
     `desc` text COMMENT '对应等级描述',
       PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COMMENT='宗祠升级';
    ");

?>