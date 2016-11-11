<?php
$this->AddSQL("
INSERT INTO `quest_activity_center` (`id`,`relative`, `weight`, `name`, `title`, `content`, `start`, `end`,`is_go`,`tag`,`is_mail`,`condition_template`,`dispose`,`sign`,`mail_title`,`mail_content`,`is_relative`)
VALUES
(26,21, 0, '累计消耗元宝', '累计消耗元宝', '累计消耗元宝', 0, 0, 0, 0, 0, '', 0, 'EVENT_TOTAL_CONSUME', '', '', 0);

");
?>