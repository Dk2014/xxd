<?php 
$this->AddSQL("
drop table if exists `resource_origin` ;
create table `resource_origin`  (
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`origin_type` tinyint(4) NOT NULL COMMENT '产生关卡类型 0--主线关卡 1--深渊 2-彩虹 3--灵宠幻境 4--魂侍关卡 5--灵宠关卡 6--伙伴关卡 5--多人关卡',
	`origin_key` int(11) NOT NULL COMMENT '产出关卡ID',
  PRIMARY KEY (`id`)
);

alter table `item` add column show_origin tinyint(4) NOT NULL DEFAULT '0' COMMENT '显示产出 0--否 1--是';
");
?>
