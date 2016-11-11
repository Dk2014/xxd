<?php

$this->AddSQL("

-- INSERT INTO `vip_privilege` (`name`, `sign`, `tip`, `order`)
-- VALUES
-- 	('云海御剑特权','YunHaiYuJian','每次可购买跟多先闪御剑行动点',27);

INSERT INTO `func` (`name`, `sign`, `lock`, `level`, `unique_key`, `need_play`, `type`)
VALUES
	('云海御剑','FUNC_DRIVING_SWORD',0,45,2097152,0,2);

-- INSERT INTO `item` (`id`, `type_id`, `quality`, `name`, `level`, `desc`, `price`, `sign`, `can_use`, `panel`, `func_id`, `renew_ingot`, `use_ingot`, `valid_hours`, `equip_type_id`, `health`, `speed`, `attack`, `defence`, `show_mode`, `equip_role_id`, `appendix_num`, `appendix_level`, `music_sign`, `can_batch`, `refine_base`, `show_origin`, `act_id`)
-- VALUES
-- 	(682,12,3,'传送符',NULL,'可开启云海的传送阵，无法出售',0,'ChuanSongFu',1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL,0,0,0,0);

");

