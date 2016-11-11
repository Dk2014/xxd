<?php

$query = $this->NewQuery("select id, segment from rainbow_level order by segment desc;");

while($rainbow_level = $query->GoNext()) {
	$id = $rainbow_level['id'];
	$segment = $rainbow_level['segment'];
	$new_sgm = $segment + 1;
	$this->AddSQL("update mission_level set parent_id = {$new_sgm} where parent_id={$segment} and parent_type=12;");
	$this->AddSQL("update rainbow_level set segment = {$new_sgm} where id={$id};");
}
$this->DropQuery($query);


?>
