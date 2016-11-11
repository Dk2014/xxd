<?php
db_execute($db,"

update  `func`  set  `lock` = 900 where `name` = '比武场';

"
);
?>
