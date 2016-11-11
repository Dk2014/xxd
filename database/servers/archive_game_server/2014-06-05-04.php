
<?php
db_execute($db, 
"
  insert into `func` (`name`,`sign`,`lock`,`unique_key`) values ('魂侍','FUNC_GHOST', '1000', '16'),('资源关卡','FUNC_RESOURCE_LEVEL', '2000', '48');

"

);
?>