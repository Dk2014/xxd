<?php 
$this->AddSQL("
REPLACE INTO `equipment_decompose` (`id`, `level`, `fragment_num`, `crystal_num`, `quality`, `crystal_id`, `dragon_ball`, `dragon_ball_num`)
VALUES
	(29, 120, 12, 0, 2, 0, 0, 0);

REPLACE INTO `mail` (`id`, `sign`, `type`, `title`, `parameters`, `content`, `expire_time`, `priority`)
	VALUES
	(45, 'zongcishangxiangqifujiang', 0, '上香祈福', 'num,铜钱数量', '今日帮派祈福灯已点满，您获得{0}铜钱奖励', 0, 0);

REPLACE INTO `mail` (`id`, `sign`, `type`, `title`, `parameters`, `content`, `expire_time`, `priority`)
	VALUES
	(42, 'cliquecancelowner', 0, '帮派邮件', 'name,新任帮主', '对不起，由于您长时间没有上线，{0}弹劾了您，成为了帮派的新帮主。', 0, 0);

REPLACE INTO `mail` (`id`, `sign`, `type`, `title`, `parameters`, `content`, `expire_time`, `priority`)
	VALUES
	(43, 'cliquechangeowner', 0, '帮派邮件', 'name,老帮主', '经过慎重考虑，{0}决定退位让贤，将帮主之位让给您，恭喜您成为新帮主', 0, 0);
REPLACE INTO `mail` (`id`, `sign`, `type`, `title`, `parameters`, `content`, `expire_time`, `priority`)
	VALUES
	(47, 'cliquejoin', 0, '帮派邮件', 'name,帮派名', '您的帮派申请已被通过，恭喜您加入{0}', 0, 0);

update `server_info` set `version`=11945;
");
?>
