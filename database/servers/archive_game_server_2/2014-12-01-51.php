<?php


$this->AddSQL("
delete from `mail_attachments` where `mail_id`='29';
insert `mail_attachments` (`mail_id`, `item_id`, `attachment_type`, `item_num`)  values
(29, 0, 2, 128),
(29, 0, 1, 188888),
(29, 232, 0, 10),
(29, 3, 6, 1);

	
");

?>
