<?php
$this->AddSQL("
update `server_info` set `version` = 0;
update `server_info` set `event_version` = 0;
");

?>
