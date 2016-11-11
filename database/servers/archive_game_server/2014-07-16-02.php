<?php
db_execute($db,"


delete from `func` where `id` in(12,13);
update `func` set `name` = '侠之试炼', `sign` = 'FUNC_XIA_ZHI_SHI_LIAN'  where `id` = 11;

"
);
?>
