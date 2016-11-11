<?php

$pid_list = $this->NewQuery("
select `pid` from `player_vip` where convert(`card_id`, unsigned integer) > 0 and convert(`card_id`, unsigned integer) <= 20150213;
");

while($row_iter=$pid_list->GoNext()){
        $pid = $row_iter["pid"];
        $mail_auto_id = $this->GetAutoID($pid, 'player_mail');
        $mail_attach_auto_id = $this->GetAutoID($pid, 'player_mail_attachment');
        $timestamp = time();

        $email_title="新版本仙尊玩家补偿";
        $email_content="感谢您一直以来对《仙侠道》的支持和热爱，在此我们给与特殊奖励“聚宝盆”（永久有效）作为2月13日前仙尊玩家的补偿。请注意查收！";

        $insert_sql="

insert into player_mail (id, pid, mail_id, state, send_time, parameters, have_attachment, title, content, expire_time, priority) values({$mail_auto_id}, {$pid}, 0, 0, {$timestamp}, '', 1, '{$email_title}', '{$email_content}', 1519574400, 1);

insert into player_mail_attachment (id, pid, player_mail_id, attachment_type, item_id, item_num) values({$mail_attach_auto_id}, {$pid}, {$mail_auto_id}, 0, 681, 1);

        ";

        $this->AddSQL($insert_sql);
}

$this->DropQuery($pid_list);

