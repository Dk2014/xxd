<?php
$this->AddSQL("
INSERT INTO `mail` (`sign`,`type`,`title`,`parameters`,`content`,`expire_time`,`priority`) VALUES ('cliqueleavebank',0,'帮派邮件','','您持有的金券、银券已中止兑换，现将您的投资原价退还。',0,0);
");
?>