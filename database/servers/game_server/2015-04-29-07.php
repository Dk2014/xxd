<?php

$this->AddSQL("
INSERT INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
VALUES
	(8, '帮派成立', 'FoundClique', 'player,player,玩家;clique,clique,帮派', '{0}成立了帮派【{1}】，欢迎各位侠士道友前去拜会，点击立即加入'),
	(9, '帮派招募', 'CliqueRecruitment', 'clique,clique,帮派', '帮派【{0}】准备大展宏图！欢迎各位侠士道友加入帮派，点击立即申请');

	");

?>
