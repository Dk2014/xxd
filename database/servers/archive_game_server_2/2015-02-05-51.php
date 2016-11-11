<?php 

$this->AddSQL("
update `skill` set `child_type`=4, `child_kind`=2 where id in (1171,1307);

update `item` set  `price`=5000, `desc`='充满光明之力的钥匙，商人现在正在大量收购，应该能买个好价钱', `show_origin`=0 where id=262;
");


?>
