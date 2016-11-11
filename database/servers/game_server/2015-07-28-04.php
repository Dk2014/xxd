<?php 

$this->AddSQL("
REPLACE INTO `func` (`id`, `name`, `sign`, `lock`, `level`, `unique_key`, `need_play`, `type`)
VALUES
	(29, '觉醒', 'FUNC_AWAKEN', 0, 50, 16777216, 1, 2);

");
?>
