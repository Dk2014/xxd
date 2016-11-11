<?php
$this->AddSQL("
delete from `town` where id = -1;
insert into `town` (`id`,`lock`,`name`,`sign`,`music`,`start_x`,`start_y`,`exit_x`,`exit_y`) values (-1,0,'帮派集会所','JiHuiSuo','YouQu',393,317,305,256);
");
?>
