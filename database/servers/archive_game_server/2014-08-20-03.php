<?php
db_execute($db, "
insert into `func`(`name`, `sign`, `lock`, `level`, `unique_key`, `need_Play`) 
values('神龙宝箱', 'FUNC_CHEST_DRAW', '1200', '0', '256', '0');

");

?>
