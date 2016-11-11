<?php
$this->AddSQL("
alter table `totem` add column   `music_sign` varchar(30) NOT NULL DEFAULT '' COMMENT '音乐资源标识';

");

?>
