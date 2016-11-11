<?php
//羁绊
$this->AddSQL("update `func` set `sign`='FUNC_FRIENDSHIP', `level`=0, `lock`=1200, `type`=1, `need_play`=1 where `id`=24 and name='伙伴羁绊';");

?>
