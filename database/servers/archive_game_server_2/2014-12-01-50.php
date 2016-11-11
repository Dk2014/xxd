<?php


$this->AddSQL("
update `player_mail` set `title`='CDKey兑换礼包', `content`='CDKey兑换礼包' where mail_id=0 and `content`='' and `title` like 'CDKey兑换礼包%';
");

?>
