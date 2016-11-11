<?php

//神龙宝箱 9% 几率宝箱
$this->AddSQL("
insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
('10', '7', '423', '7');
");

//神龙宝箱 十连抽必得
$this->AddSQL("
insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
('-5', '4', '1', '1'),
('-5', '4', '13', '1');

");

?>
