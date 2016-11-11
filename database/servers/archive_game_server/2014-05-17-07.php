<?php
db_execute($db, "
	ALTER TABLE `ghost` drop column `hit_level`;
	ALTER TABLE `ghost` drop column `dodge_level`;
	ALTER TABLE `ghost` drop column `crit_level`;
	ALTER TABLE `ghost` drop column `block_level`;
	
	ALTER TABLE `ghost` drop column `max_hit_level`;
	ALTER TABLE `ghost` drop column `max_dge_level`;
	ALTER TABLE `ghost` drop column `max_cri_level`;
	ALTER TABLE `ghost` drop column `max_blk_level`;

");
?>