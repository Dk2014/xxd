
<?php
db_execute($db, 
"

ALTER TABLE `player_item_appendix`
DROP COLUMN `refine_level`;

ALTER TABLE `player_item`
ADD COLUMN `refine_level` tinyint(4) NOT NULL DEFAULT 0 COMMENT '精练等级' AFTER `appendix_id`;

"
);

?>

