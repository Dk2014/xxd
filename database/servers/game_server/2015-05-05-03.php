<?php 
$this->AddSQL("
REPLACE INTO `item` (`id`, `type_id`, `quality`, `name`, `level`, `desc`, `price`, `sign`, `can_use`, `panel`, `func_id`, `renew_ingot`, `use_ingot`, `valid_hours`, `equip_type_id`, `health`, `speed`, `attack`, `defence`, `show_mode`, `equip_role_id`, `appendix_num`, `appendix_level`, `music_sign`, `can_batch`, `refine_base`, `show_origin`, `act_id`, `can_sell`)
VALUES
	(792, 21, 3, '混云幡', NULL, '重置当前所在该层的云海', 0, 'HunYunFan', 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, NULL, 0, 0, 0, 0, 1);
");
?>
