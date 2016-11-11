<?php
$sql = 'select * from `player_month_card_info` where endtime-starttime < 30*3600*24;';
$query = $this->NewQuery($sql);
$row =  $query->GoNext();
while(isset($row)){
	$new_sql = 'update `player_month_card_info` set `endtime` = (`starttime` + 30*24*3600) where pid = '.$row['pid'];
	echo $row['pid'] + ' | ' + $row['endtime']+"\n";
	$this->AddSQL($new_sql);
	$row =  $query->GoNext();
}
$this->DropQuery($query);
?>