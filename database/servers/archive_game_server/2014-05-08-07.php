<?php
db_execute($db, "
	ALTER TABLE `mission_level` DROP column `award_item1_type`;
	ALTER TABLE `mission_level` DROP column `award_item1_id`;
	ALTER TABLE `mission_level` DROP column `award_item1_chance`;
	ALTER TABLE `mission_level` DROP column `award_item1_num`;

	ALTER TABLE `mission_level` DROP column `award_item2_type`;
	ALTER TABLE `mission_level` DROP column `award_item2_id`;
	ALTER TABLE `mission_level` DROP column `award_item2_chance`;
	ALTER TABLE `mission_level` DROP column `award_item2_num`;


	ALTER TABLE `mission_level` DROP column `award_item3_type`;
	ALTER TABLE `mission_level` DROP column `award_item3_id`;
	ALTER TABLE `mission_level` DROP column `award_item3_chance`;
	ALTER TABLE `mission_level` DROP column `award_item3_num`;

	ALTER TABLE `mission_level` DROP column `award_item4_type`;
	ALTER TABLE `mission_level` DROP column `award_item4_id`;
	ALTER TABLE `mission_level` DROP column `award_item4_chance`;
	ALTER TABLE `mission_level` DROP column `award_item4_num`;

	ALTER TABLE `mission_level` DROP column `award_item5_type`;
	ALTER TABLE `mission_level` DROP column `award_item5_id`;
	ALTER TABLE `mission_level` DROP column `award_item5_chance`;
	ALTER TABLE `mission_level` DROP column `award_item5_num`;
");
?>