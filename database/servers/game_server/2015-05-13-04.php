<?php
//帮派送镖个人信息
$this->AddSQL("
REPLACE INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
VALUES
	(20, '帮派动态_镖船被劫持', 'CliqueBoatHijacked', 'clique,clique,敌对帮派;player,hijacker,劫持者;player,boat_owner,我帮玩家;boat,dummy_link,镖船', '{0}帮派的{1}，劫持了本帮{2}的镖船，{3}'),
	(21, '帮派动态_镖船夺回', 'CliqueBoatRecover', 'player,fighter,夺回者;player,boat_owner,镖船主', '本帮{0}仗义出手，夺回{1}的镖船');

REPLACE INTO `clique_boat` (`id`, `name`, `sign`, `rate`, `escort_time`, `select_cost_ingot`, `award_coins`, `award_fame`, `award_clique_contrib`, `hijack_lose_coins`, `hijack_lose_fame`, `hijack_lose_clique_contrib`)
	VALUES
		(1, '小货船', 'XiaoHuoChuan', 30, 20, 0, 2000, 100, 100, 800, 30, 100),
	(2, '货船', 'HuoChuan', 30, 20, 0, 4000, 200, 200, 1600, 60, 200),
	(3, '商船', 'ShangChuan', 30, 20, 0, 6000, 300, 300, 2100, 90, 300),
	(4, '大商船', 'DaShangChuan', 10, 20, 0, 8000, 400, 400, 2400, 120, 400),
	(5, '宝船', 'BaoChuan', 0, 20, 100, 30000, 3000, 3000, 5000, 1000, 3000);

REPLACE INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
	VALUES
		(22, '镖船_我劫持的被夺回', 'BoatRecovered', 'player,fighter,夺回者', '很遗憾{0}击败了你，夺回了镖船'),
	(23, '镖船_我的被仗义夺回', 'BoatRecoveredByHero', 'player,fighter,夺回者', '{0}仗义出手帮你夺回了镖船'),
	(24, '镖船_劫镖成功', 'BoatHijackFinished', 'string,coins,铜钱;string,fame,声望;string,contrib,贡献', '恭喜你劫镖成功，获得 {0} 铜钱 {1}声望 {2}帮贡'),
	(25, '镖船_运送成功', 'BoatEscortFinished', 'string,coins,铜钱;string,fame,声望;string,contrib,贡献', '恭喜你护送运镖成功，获得 {0} 铜钱 {1}声望 {2}帮贡'),
	(26, '镖船_被劫持成功', 'BoatHijackingFinished', 'player,hijacker,劫持者;string,coins,铜钱;string,fame,声望;string,contrib,贡献', '{0}劫持了镖船，损失 {0} 铜钱 {1}声望 {2}帮贡');

REPLACE INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
	VALUES
		(27, '镖船_被劫持', 'BoatHijacking', 'clique,hijacker_clique,劫持者帮派;player,hijacker,劫持者', '{0}的{1}劫持了你的镖船，快去找人帮忙');




");
