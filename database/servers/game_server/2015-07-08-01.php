 <?php

 $this->AddSQL(
	"ALTER TABLE `item_exchange` ADD COLUMN `target_item_num` smallint(6) NOT NULL DEFAULT 1 COMMENT '目标物品数量';"
 );