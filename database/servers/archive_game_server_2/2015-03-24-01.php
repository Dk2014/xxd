<?php
$this->AddSQL("
INSERT INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
VALUES
	(1, '剑山拔剑获得', 'DrawSwordSoul', 'name,剑心名', '{nick}在【剑山拔剑】中提神运气，提手拔剑获得了神剑{0}'),
	(2, '命锁宝箱获得装备', 'FateBoxEquipment', 'name,装备名', '{nick}在【命锁】中运气爆棚获得了金色装备{0}'),
	(3, '命锁宝箱获得魂侍碎片', 'FateBoxGhostFrame', 'name,魂侍名', '{nick}在【命锁】中运气爆棚获得了金色魂侍碎片0}'),
	(4, '彩虹桥获得魂侍', 'RainbowLevelGhost', 'level,关卡;name,魂侍名', '{nick}{0}中突破挑战获得了金色魂侍{1}'),
	(5, '合成金色魂侍', 'ComposeGhost', 'name,魂侍名', '{nick}经过了长期累积获得了金色魂侍{0}');
");
?>


