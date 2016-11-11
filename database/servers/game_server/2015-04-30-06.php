<?php
$this->AddSQL("
update `mail` set content = '尊敬的帮主，您管理的{0}帮派成员已达{1}人，获得每日帮派工资{2}铜钱！' where id = 37;
update `mail` set content = '尊敬的副帮主，您管理的{0}帮派成员已达{1}人，获得每日帮派工资{2}铜钱' where id = 38;
update `mail` set content = '对不起，您被踢出了{0}帮派。！' where id = 39;
update `mail` set content = '很遗憾，您的帮派{0}已解散' where id = 44;
");
?>