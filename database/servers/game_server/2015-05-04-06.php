
<?php

$this->AddSQL("
        delete from world_channel_message where id = 19;
");

$this->AddSQL("
INSERT INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
VALUES
    (19, '帮派动态_建筑升级', 'CliqueBuildingLevelUP', 'clique,clique,帮派;string,build_type,帮派建筑名称;string,clique_level,帮派等级', '帮派{0} {1}已升级至 {2} 级');
");

?>
