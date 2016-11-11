<?php

//神龙宝箱 十连抽必得
$this->AddSQL("
update `chest_item` set item_id=18 where id=413;
");

//神龙宝箱 9% 几率宝箱
$this->AddSQL("
update `chest_item` set `item_num`=3 where id=415;
update `chest_item` set `item_num`=3 where id=416;
");

?>

