<?php 
$this->AddSQL("
ALTER TABLE equipment_decompose MODIFY COLUMN id smallint(6);
ALTER TABLE equipment_decompose DROP PRIMARY KEY;
ALTER TABLE equipment_decompose ADD PRIMARY KEY (id);
ALTER TABLE equipment_decompose MODIFY COLUMN id smallint(6) AUTO_INCREMENT;

ALTER TABLE equipment_decompose ADD COLUMN `quality` tinyint(4) NOT NULL DEFAULT '0'  COMMENT '品质';
ALTER TABLE equipment_decompose ADD COLUMN `crystal_id` smallint(6) NOT NULL DEFAULT '0'  COMMENT '获得结晶';



ALTER TABLE equipment_refine ADD COLUMN `level6_consume_coin` bigint(20) NOT NULL DEFAULT 0 COMMENT '精练到6级消耗的铜钱';
ALTER TABLE equipment_refine ADD COLUMN `level7_consume_coin` bigint(20) NOT NULL DEFAULT 0 COMMENT '精练到7级消耗的铜钱';
ALTER TABLE equipment_refine ADD COLUMN `level8_consume_coin` bigint(20) NOT NULL DEFAULT 0 COMMENT '精练到8级消耗的铜钱';
ALTER TABLE equipment_refine ADD COLUMN `level9_consume_coin` bigint(20) NOT NULL DEFAULT 0 COMMENT '精练到9级消耗的铜钱';
ALTER TABLE equipment_refine ADD COLUMN `level10_consume_coin` bigint(20) NOT NULL DEFAULT 0 COMMENT '精练到10级消耗的铜钱';


");
?>
