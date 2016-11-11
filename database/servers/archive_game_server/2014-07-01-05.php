<?php

db_execute($db, "
  CREATE TABLE `global_arena_rank` (
  `rank` int(11) NOT NULL COMMENT '排名',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  PRIMARY KEY (`rank`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局比武场数据';

CREATE TABLE `arena_award_box` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `max_rank` int(11) NOT NULL COMMENT '排名',
  `fame` int(11) NOT NULL COMMENT '声望',
  `coins` int(11) NOT NULL COMMENT '铜钱',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `item_num` smallint(6) NOT NULL COMMENT '物品数量',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比武场奖励宝箱';


CREATE TABLE `player_arena` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `daily_num` smallint(6) NOT NULL COMMENT '今日已挑战次数',
  `failed_cd_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '战败CD结束时间',
  `record_read_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '战报阅读时间',
  `win_times` smallint(6) NOT NULL DEFAULT '0' COMMENT '>0 连胜场次; 0 保持不变; -1 下降趋势',
  `daily_award_coin` int(11) NOT NULL DEFAULT '0' COMMENT '今日获得铜钱累计',
  `daily_award_longbi` int(11) NOT NULL DEFAULT '0' COMMENT '今日获得龙币累计',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场数据';

CREATE TABLE `player_arena_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mode` tinyint(4) NOT NULL COMMENT '记录类型，0无数据，1挑战成功，2挑战失败，3被挑战且成功，4被挑战且失败',
  `old_rank` int(11) NOT NULL COMMENT '原排位',
  `new_rank` int(11) NOT NULL COMMENT '新排位',
  `fight_num` int(11) NOT NULL COMMENT '战力',
  `target_pid` bigint(20) NOT NULL COMMENT '对手玩家ID',
  `target_old_rank` int(11) NOT NULL COMMENT '对手原排位',
  `target_new_rank` int(11) NOT NULL COMMENT '对手新排位',
  `target_fight_num` int(11) NOT NULL COMMENT '对手战力',
  `record_time` bigint(20) NOT NULL COMMENT '记录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场记录';

CREATE TABLE `player_arena_rank` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT  '昨天排名',
  `rank1` int(11) NOT NULL DEFAULT '0' COMMENT '1天前排名',
  `rank2` int(11) NOT NULL DEFAULT '0' COMMENT '2天前排名',
  `rank3` int(11) NOT NULL DEFAULT '0' COMMENT '3天前排名',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场最近排名记录';

");
?>
