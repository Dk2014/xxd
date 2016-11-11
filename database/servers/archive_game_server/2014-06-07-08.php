
<?php
db_execute($db, 
"
  update `func` set `unique_key` = 1 where id = 3;
  update `func` set `unique_key` = 2 where id = 4;
  update `func` set `unique_key` = 4 where id = 5;
  insert into `func` (`name`,`sign`,`lock`,`unique_key`) values ('多人关卡','FUNC_MULTI_LEVEL', '4000', '8');
"

);
?>