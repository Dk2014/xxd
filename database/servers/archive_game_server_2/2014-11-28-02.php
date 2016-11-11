<?php

//神龙宝箱 1% 几率宝箱
$this->AddSQL("
insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
('12', '7', '335', '7'),
('12', '7', '424', '7');
");

//神龙宝箱 十连抽必得
$this->AddSQL("
insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
('-5', '4', '24', '1'),
('-5', '4', '25', '1');

");

//神龙宝箱 9% 几率宝箱
$this->AddSQL("
insert into `chest_item` (`chest_id`, `type`, `item_id`, `item_num`) values
('10', '7', '335', '7'),
('10', '7', '424', '7');
");

?>
