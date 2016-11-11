
<?php
db_execute($db, "
    CREATE TABLE `player_sword_soul_only` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家物品ID',
      `pid` bigint(20) NOT NULL COMMENT '玩家ID',
      `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
      `pos` smallint(6) DEFAULT NULL COMMENT '冗余位置',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标记每个玩家只有一个的剑心相关数据';
");
?>