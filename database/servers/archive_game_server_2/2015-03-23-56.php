<?php 
//$this->AddSQL("
//update player_totem_info set pos1=0 where pid=430660665737648;
//");
//

$gs_id = 0;
$effect_player_count = 0;
$gsid_query = $this->NewQuery("select id>>32 as gsid from player limit 1;");
while($row =  $gsid_query->GoNext()) {
	$gs_id = intval(intval($row['gsid'])/10);
}
$this->DropQuery($gsid_query);

$sql = "select * from player_totem_info";
$query = $this->NewQuery($sql);
while($row = $query->GoNext()) {
	$pos_array = array();
	$pos_array[] = $row['pos1'];
	$pos_array[] = $row['pos2'];
	$pos_array[] = $row['pos3'];
	$pos_array[] = $row['pos4'];
	$pos_array[] = $row['pos5'];
	$effect = false;
	foreach($pos_array as $idx=>$pos) {
		if($pos >0) {
			$query2 = $this->NewQuery("select * from player_totem where id={$pos}");
			if(!$query2->Have()){
				$effect = true;
				echo "totem_fix{$row['pid']} {$idx} {$pos}\n";
				$real_idx = $idx+1;
				$this->AddSQL("update player_totem_info set pos{$real_idx}=0 where pid={$row['pid']}");
			}
			$this->DropQuery($query2);
		}
	}
	if($effect) {
		$effect_player_count += 1;
	}
}

$this->DropQuery($query);

echo "gs {$gs_id} effect player {$effect_player_count}\n";

?>
