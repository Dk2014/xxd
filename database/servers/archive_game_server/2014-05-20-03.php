
<?php
db_execute($db, "
CREATE TABLE `battle_item_config` (
  `id` smallint(6) NOT NULL COMMENT '道具物品ID',
  `target_type` tinyint(4) NOT NULL COMMENT '目标对象',
  `effect_type` tinyint(4) NOT NULL COMMENT '效果类型',
  `config` text COMMENT '产生效果配置',
  `keep` tinyint(4) NOT NULL COMMENT '持续回合',
  `max_override` tinyint(4) NOT NULL COMMENT '最大叠加数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='战斗道具';
");
?>