<?php
$this->AddSQL("
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquesalaryowner',0,'帮派邮件','name,帮派名;num,帮派人数;coins,奖励铜钱','尊敬的帮主，您管理的【{0}】帮派成员已达{1}人，获得每日帮派工资{2}铜钱！',14,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquesalarymanager',0,'帮派邮件','name,帮派名;num,帮派人数;coins,奖励铜钱','尊敬的副帮主，您管理的【{0}】帮派成员已达{1}人，获得每日帮派工资{2}铜钱！',14,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliqueleave',0,'帮派邮件','name,帮派名','对不起，您被踢出了【{0}】帮派。！',0,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquebemanger',0,'帮派邮件','','恭喜您，您被提升为副帮主，每日可获得帮派工资（帮众人数*100）。',0,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquecancelmanager',0,'帮派邮件','','对不起，您的副帮主职务已被撤除',0,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquecancelowner',0,'帮派邮件','name,新任帮主','对不起，由于您长时间没有上线，【{0}】弹劾了您，成为了帮派的新帮主。',0,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquechangeowner',0,'帮派邮件','name,老帮主','经过慎重考虑，【{0}】决定退位让贤，将帮主之位让给您，恭喜您成为新帮主',0,0);
insert into `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) values ('cliquedestory',0,'帮派邮件','name,帮派名','很遗憾，您的帮派【{0}】已解散',0,0);
");
?>