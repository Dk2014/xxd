<?php
$this->AddSQL("
update `item` set `type_id`='18' where `type_id`='2' and `name` like '%喜好品%';

");
?>
