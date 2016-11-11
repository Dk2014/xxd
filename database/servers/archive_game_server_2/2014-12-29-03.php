<?php

//数据修复

$equal_sql=/*$this->AddSQL*/("

-- 踢掉玩家等级未开启的宠物格中的宠物
UPDATE `player_battle_pet_grid`
SET `battle_pet_id` = 0
WHERE `id` IN (

		SELECT `id`
		FROM (

			-- Attach level column to player_battle_pet_grid
			SELECT `id`
				,`grid_id`
				,`battle_pet_id`
				,`lt`.`level`
			FROM `player_battle_pet_grid`
			LEFT JOIN (

				-- Player main role level
				SELECT `pid`
					,`level`
				FROM `player_role`
				WHERE `role_id` = 1
					OR `role_id` = 2

				) `lt` ON `lt`.`pid` = `player_battle_pet_grid`.`pid`

			) `atable`
		WHERE `battle_pet_id` != 0
			AND (
				(
					`grid_id` = 2
					AND `level` < 30
					)
				OR (
					`grid_id` = 3
					AND `level` < 40
					)
				OR (
					`grid_id` = 4
					AND `level` < 50
					)
				OR (
					`grid_id` = 5
					AND `level` < 60
					)
				)

		);

");

{/*below is a replacing of previous sql, their behaviors should be same*/

$player_battle_pet_grid=$this->NewQuery("
select `id`, `pid`, `grid_id`, `battle_pet_id` from `player_battle_pet_grid`;
");
$player_role=$this->NewQuery("
select `pid`, `level` from `player_role` where `role_id`=1 or `role_id`=2;
");
$mapPid2Level=array();
while($row_iter=$player_role->GoNext()){
	$pid=$row_iter["pid"];
	$level=$row_iter["level"];
	$mapPid2Level[$pid]=$level;
}
$this->DropQuery($player_role);
$ids="";
while($row_iter=$player_battle_pet_grid->GoNext()){
	if($row_iter['battle_pet_id']!=0
		&& (($row_iter['grid_id']==2 && $mapPid2Level[$row_iter['pid']]<30)
			|| ($row_iter['grid_id']==3 && $mapPid2Level[$row_iter['pid']]<40)
			|| ($row_iter['grid_id']==4 && $mapPid2Level[$row_iter['pid']]<50)
			|| ($row_iter['grid_id']==5 && $mapPid2Level[$row_iter['pid']]<60)
			)
		){
		$ids.=$row_iter['id'].',';
	}
}
$this->DropQuery($player_battle_pet_grid);
if(strlen($ids)>1){
	$ids=substr($ids, 0, strlen($ids)-1);
	$this->AddSQL("
update `player_battle_pet_grid` set `battle_pet_id`=0 where `id` in ({$ids});
	");
}

}/*===========================================================================*/

$t_pid2soul_num=$this->NewQuery("

-- Soul number
SELECT `result`.`pid`
	,round(sum(`result`.`soul_num`)) `soul_num`
FROM (

	-- Get each grid level of a player needs soul number
	SELECT `player_battle_pet_grid`.`pid`
		,`level_soul`.`soul_num`
	FROM `player_battle_pet_grid`
	INNER JOIN (

		-- Calculate the soul number needed for each level to levelup
		SELECT `level`
			,`exp` * 2 / (`min_add_exp` + `max_add_exp`) * `cost_soul_num` * 1.05 `soul_num`
		FROM `battle_pet_grid_level`

		) `level_soul` ON `player_battle_pet_grid`.`level` > `level_soul`.`level`

	UNION ALL

	-- Calculate the soul number by players' grid exp
	SELECT `players`.`pid`
		,`players`.`exp` * 2 / (`systems`.`min_add_exp` + `systems`.`max_add_exp`) * `systems`.`cost_soul_num` * 1.05 `soul_num`
	FROM `player_battle_pet_grid` `players`
	LEFT JOIN `battle_pet_grid_level` `systems` ON `players`.`level` = `systems`.`level`
	WHERE `players`.`exp` > 0

	) `result`
GROUP BY `result`.`pid`;

");

while($row_iter=$t_pid2soul_num->GoNext()){
	$pid = $row_iter["pid"];
	$count = $row_iter["soul_num"];
	$mail_auto_id = $this->GetAutoID($pid, 'player_mail');
	$mail_attach_auto_id = $this->GetAutoID($pid, 'player_mail_attachment');
	$timestamp = time();

	$email_title="灵宠补偿邮件";
	$email_content="亲爱的大侠，小梦妖代表灵宠们表示不喜欢对于它们的培养系统，所以我们对灵宠系统进行了修改~\n我们取消了灵宠契约之力的培养，大侠之前用在培养契约之力所用的灵魄将返还给您，并会给与返还总数的5%作为补偿。\n我们同时取消了灵宠的星级，将星级兑换成了相应等级，请各位大侠放心\n除此之外，现在增加了灵宠格子的等级开启，灵魄可用于提升每只灵宠的等级。特别的是，我们增加了铜币训练灵宠的绝招。\n大侠赶快去体验一下新的灵宠系统吧。";

	$insert_sql="

insert into player_mail (id, pid, mail_id, state, send_time, parameters, have_attachment, title, content, expire_time, priority) values({$mail_auto_id}, {$pid}, 0, 0, {$timestamp}, '', 1, '{$email_title}', '{$email_content}', 1519574400, 1);

insert into player_mail_attachment (id, pid, player_mail_id, attachment_type, item_id, item_num) values({$mail_attach_auto_id}, {$pid}, {$mail_auto_id}, 0, 355, {$count});

	";

	$this->AddSQL($insert_sql);
}

$this->DropQuery($t_pid2soul_num);

/*============================================================================*/

//$player_battle_pet_grid=$this->NewQuery("
//select `player_battle_pet_grid`.*, `player_battle_pet`.`parent_pet_id` from `player_battle_pet_grid` left join `player_battle_pet` on `player_battle_pet_grid`.`pid`=`player_battle_pet`.`pid` and `player_battle_pet_grid`.`battle_pet_id`=`player_battle_pet`.`battle_pet_id`;
//");
//
//while($row_iter=$player_battle_pet_grid->GoNext()){
//	if(isset($row_iter['parent_pet_id'])&&$row_iter['parent_pet_id']>0)
//		$this->AddSQL("
//update `player_battle_pet_grid` set `battle_pet_id`={$row_iter['parent_pet_id']} where `id`={$row_iter['id']};
//		");
//}
//
//$this->DropQuery($player_battle_pet_grid);

/*============================================================================*/

$this->AddSQL("

-- 取消灵宠阶数
update `player_battle_pet` set `battle_pet_id`=`parent_pet_id`;

-- 星级等级转换
update `player_battle_pet` set `level`=20 where `star`=1;
update `player_battle_pet` set `level`=30 where `star`=2;
update `player_battle_pet` set `level`=45 where `star`=3;
update `player_battle_pet` set `level`=60 where `star`=4;
update `player_battle_pet` set `level`=75 where `star`=5;

");

