<?php
db_execute($db,"
alter table `vip_privilege` add column `order` int(11) unique;
"
);
?>
