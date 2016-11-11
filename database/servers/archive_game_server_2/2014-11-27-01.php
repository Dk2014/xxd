<?php

$this->AddSQL("
delete from trader_grid_config where grid_id=18;

insert into `trader_grid_config` (`grid_id`, `item_id`, `num`, `probability`, `cost`, `goods_type`)  values
	('18', '256', '10', '8', '498', '0'),
	('18', '257', '10', '8', '498', '0'),
	('18', '258', '10', '8', '498', '0'),
	('18', '338', '10', '8', '498', '0'),
	('18', '260', '10', '8', '498', '0'),
	('18', '312', '10', '8', '498', '0'),
	('18', '314', '10', '8', '498', '0'),
	('18', '316', '10', '8', '498', '0'),
	('18', '313', '10', '8', '498', '0'),
	('18', '315', '10', '8', '498', '0'),
	('18', '424', '10', '10', '498', '0'),
	('18', '335', '10', '10', '498', '0');

insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
	('12', '4', '26', '1');

update `quest_activity_center` set `start`='0', `end`='604800', `dispose`='604800', `is_relative`='1' where `id`=24;

update `server_info` set `version`='8317';

");

?>
