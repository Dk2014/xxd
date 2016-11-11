<?php
db_execute($db, "
alter table `global_tb_xxd_onlinecnt` add index timekey(timekey);

");

?>
