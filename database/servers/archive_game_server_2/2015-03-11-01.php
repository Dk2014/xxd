<?php
$query = $this->NewQuery("select pid from player_role where status=0 group by pid having count(*) > 9;");
$arr = array();
while($row = $query->GoNext()) {
	$arr[] = $row['pid'];
}
$this->DropQuery($query);
if(count($arr) > 0){
	foreach ($arr as $pid) {
		$query_roles = $this->NewQuery("select role_id,status from player_role where pid={$pid} and (role_id <> 1 && role_id <> 2) order by level ASC LIMIT 1");
		$row = $query_roles->GoNext();
		$final_role = $row['role_id'];
		$this->DropQuery($query_roles);
		$this->AddSQL("update player_role set status=1 where pid={$pid} and role_id={$final_role}");
	}
}
?>
