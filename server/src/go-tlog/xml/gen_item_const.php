<?php
$db = mysql_connect ( '42.120.22.64', "rlms", "xxdrlms" );
$ok = mysql_query('use xxd_dev20140919;' , $db);
if(!$ok) {
	die(mysql_error());
}

$sql = 'select * from item_type';
$result = mysql_query($sql, $db);

echo "<macrosgroup name=\"ItemType\">\n";
echo "	<macro name=\"IT_GHOST\" value=\"-1\" desc=\"魂侍(完整)\"/>\n	<macro name=\"IT_SWORD\" value=\"-2\" desc=\"剑心(完整)\"/>\n	<macro name=\"IT_PETF\" value=\"-3\" desc=\"灵宠(完整)\"/>\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign);
	$value = intval($row->id);
	$desc = ($row->name);
	echo "	<macro name=\"IT_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";


echo "\n\n";
$sql = 'select * from item';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Item\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign) . '_' . $row->id;
	$value = intval($row->id);
	$desc = ($row->name);
	echo "	<macro name=\"ITEM_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select * from ghost';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Ghost\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign) . '_' . $row->id;
	$value = intval($row->id);
	$desc = ($row->name);
	echo "	<macro name=\"GHOST_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select * from sword_soul';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Sword\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign) . '_' . $row->id;
	$value = intval($row->id);
	$desc = ($row->name);
	echo "	<macro name=\"SWORD_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select * from mission_level';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Mission\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign) . '_' . $row->id;
	$value = intval($row->id);
	$desc = ($row->name);
	$parent_type = ($row->parent_type);
	$parent_id = ($row->parent_id);
	if ($parent_type == 12) {
		echo "	<macro name=\"Mission_{$name}\" value=\"{$value}\" desc=\"{$desc}_{$parent_id}\"/>\n";	
	} else {
		echo "	<macro name=\"Mission_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
	}
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select * from mission_level  where `lock` >0 order  by `lock` asc';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"MissionOrder\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->sign) . '_' . $row->id;
	$value = intval($row->lock);
	$desc = ($row->id);
	echo "	<macro name=\"MissionOrder_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select id,name,`order` from quest order by `order` asc';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Quest\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = strtoupper($row->order);
	$value = intval($row->id);
	$desc = ($row->name);
	echo "	<macro name=\"QUEST_{$name}_{$value}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select id,name from daily_quest';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"DailyQuest\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = $row->name;
	$value = intval($row->id);
	echo "	<macro name=\"QUEST_{$value}\" value=\"{$value}\" desc=\"{$name}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select p.pet_id,e.name from battle_pet as p left join enemy_role as e on (p.pet_id = e.id)';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"Pet\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name = $row->name;
	$value = intval($row->pet_id);
	echo "	<macro name=\"Pet_{$value}\" value=\"{$value}\" desc=\"{$name}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

echo "\n\n";
$sql = 'select m.id,h.hard_level_lock from hard_level as h left join mission_level as m on (h.id = m.parent_id) where m.parent_type = 8 order by h.hard_level_lock asc';
$result = mysql_query($sql, $db);
echo "<macrosgroup name=\"HardMissionOrder\">\n";
do {
	$row = mysql_fetch_object($result);
	if(!$row) {
		break;
	}
	$name =  $row->id;
	$value = intval($row->hard_level_lock);
	$desc = ($row->id);
	echo "	<macro name=\"HardMissionOrder_{$name}\" value=\"{$value}\" desc=\"{$desc}\"/>\n";
} while($row);
echo "</macrosgroup>\n";

?>
