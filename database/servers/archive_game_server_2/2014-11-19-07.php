
<?php

$arr_levels = array("熔岩火山三","熔岩火山七","剑灵密室三","剑灵密室七","血雾岭三","血雾地牢三");

foreach ($arr_levels as $val) {
	$this->AddSQL("

delete from enemy_deploy_form where battle_type=0 and parent_id in (select id from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}'));
delete from level_battle_pet where mission_enemy_id in (select id from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}'));
delete from mission_talk where enemy_id in (select id from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}'));

delete from mission_level_box where mission_level_id = (select id from mission_level where name ='{$val}');

delete from mission_level_small_box_items where small_box_id in (select id from mission_level_small_box where mission_level_id = (select id from mission_level where name ='{$val}'));
delete from mission_level_small_box where mission_level_id = (select id from mission_level where name ='{$val}');

delete from level_star where level_id = (select id from mission_level where name ='{$val}');
delete from mission_level_recovery_meng_yao where mission_level_id = (select id from mission_level where name ='{$val}');

delete from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}');
delete from mission_level where name ='{$val}';

");
}

?>
