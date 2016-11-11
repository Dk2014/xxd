
<?php
db_execute($db, 
	"
 DROP TABLE IF EXISTS `player_realm`;
CREATE TABLE `player_realm` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `realm_level` smallint(6) NOT NULL COMMENT '角色境界等级',
  `realm_exp` bigint(20) NOT NULL COMMENT '角色境界经验',
  `realm_class` smallint(6) NOT NULL COMMENT '角色境界阶级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色境界表';
	"
);
?>