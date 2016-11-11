
<?php

$arr_levels = array("青竹林二","青竹林三","青竹林五","青竹林七","黑夜森林二","黑夜森林三","暗影秘境一","暗影秘境三","莲花峰二","莲花峰三","水溶洞三");

foreach ($arr_levels as $val) {
	db_execute($db, "

delete from enemy_deploy_form where battle_type=0 and parent_id in (select id from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}'));
delete from level_battle_pet where mission_enemy_id in (select id from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}'));

delete from mission_level_box where mission_level_id = (select id from mission_level where name ='{$val}');

delete from mission_level_small_box_items where small_box_id in (select id from mission_level_small_box where mission_level_id = (select id from mission_level where name ='{$val}'));
delete from mission_level_small_box where mission_level_id = (select id from mission_level where name ='{$val}');

delete from level_star where level_id = (select id from mission_level where name ='{$val}');

delete from mission_enemy where mission_level_id = (select id from mission_level where name ='{$val}');
delete from mission_level where name ='{$val}';

");
}

?>
