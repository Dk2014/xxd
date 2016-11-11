<?php 

$this->AddSQL(
"delete from global_mail_attachments where `global_mail_id` not in (select id from global_mail);"
);
?>

