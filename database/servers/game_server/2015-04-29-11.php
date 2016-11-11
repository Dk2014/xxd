<?php
$this->AddSQL("
delete from world_channel_message;
");

$this->AddSQL("
	INSERT INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
	VALUES
	(1, '剑山拔剑获得', 'DrawSwordSoul', 'player,player,玩家;item,item,道具', '{0}在<font color=#c503fd >【剑山拔剑】</font>中提手拔剑获得了{1}'),
(2, '命锁宝箱获得装备', 'FateBoxEquipment', 'player,player,玩家;item,item,道具', '{0}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1}'),
(3, '命锁宝箱获得魂侍碎片', 'FateBoxGhostFrame', 'player,player,玩家;item,item,道具', '{0}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1}'),
(4, '彩虹桥获得魂侍', 'RainbowLevelGhost', 'player,player,玩家;string,level,关卡;item,item,道具', '{0}在{1}中突破挑战获得了金色魂侍{2}'),
(5, '合成金色魂侍', 'ComposeGhost', 'player,player,玩家;item,item,道具', '{0}经过了长期累积获得了金色魂侍{1}'),
(6, '召唤金色阵印', 'CallTotem', 'player,player,玩家;item,item,道具', '{0}使用<font color=#c503fd >【圣器召唤】</font>召唤出守护阵印{1}'),
(8, '帮派成立', 'FoundClique', 'player,player,玩家;clique,clique,帮派', '{0}成立了帮派【{1}】，欢迎各位侠士道友前去拜会，点击立即加入'),
(9, '帮派招募', 'CliqueRecruitment', 'clique,clique,帮派', '帮派【{0}】准备大展宏图！欢迎各位侠士道友加入帮派，点击立即申请'),
(10, '帮派动态_任命副帮主', 'CliqueAssignManger', 'player,player,玩家', '{0}被提升为副帮主'),
(11, '帮派动态_转让帮主', 'CliqueAssignOwner', 'player,old_owner,原帮主;player,new_owner,新帮主', '{0}将帮主的位置转让给了{1}'),
(12, '帮派动态_弹劾', 'CliqueElectOwner', 'player,new_owner,新班主;player,player2,原帮主', '{0}弹劾了{1}成为了新帮主'),
(13, '帮派动态_退出帮派', 'CliqueMemberLeave', 'player,player,玩家', '{0}离开了帮派'),
(14, '帮派动态_移除成员', 'CliqueKickMember', 'player,manger,管理员;player,member,成员', '{0}将{1}踢出了帮派'),
(15, '帮派动态_新公告', 'CliqueNewAnnc', '', '帮派公告已更新，请注意查看'),
(16, '聊天', 'CommonChat', 'string,content,内容', '{0}'),
(17, '帮派动态_解除副帮主', 'CliqueFireManger', 'player,manger,管理员', '{0}被解除了副帮主的职位');


");

?>
