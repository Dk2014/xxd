
<?php
db_execute($db, 
"
	DROP TABLE `player_sword_soul_equipped`;

	CREATE TABLE `player_sword_soul_equipment` (
		`id` bigint(20) NOT NULL COMMENT '主键ID',
		`pid` bigint(20) NOT NULL COMMENT '玩家ID',
		`role_id` tinyint(4) NOT NULL COMMENT '角色ID',
		`pos0` bigint(20) NOT NULL COMMENT '装备位置1的剑心',
		`pos1` bigint(20) NOT NULL COMMENT '装备位置2的剑心',
		`pos2` bigint(20) NOT NULL COMMENT '装备位置3的剑心',
		`pos3` bigint(20) NOT NULL COMMENT '装备位置4的剑心',
		`pos4` bigint(20) NOT NULL COMMENT '装备位置5的剑心',
		`pos5` bigint(20) NOT NULL COMMENT '装备位置6的剑心',
		`pos6` bigint(20) NOT NULL COMMENT '装备位置7的剑心',
		`pos7` bigint(20) NOT NULL COMMENT '装备位置8的剑心',
		`pos8` bigint(20) NOT NULL COMMENT '装备位置9的剑心',
		PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家剑心装备表';

"
);

?>