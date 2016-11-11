


<?php

$this->AddSQL("
INSERT INTO `mail` (`id`, `sign`, `type`, `title`, `parameters`, `content`, `expire_time`, `priority`)
VALUES
    (47, 'cliquejoin', 0, '帮派邮件', 'name,帮派名', '您的帮派申请已被通过，恭喜您加入【{0}】', 0, 0);
");

?>
