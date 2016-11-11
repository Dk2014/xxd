<?php 

$gs_id = 0;
//计算gsid
$gsid_query = $this->NewQuery("select id>>32 as gsid from player limit 1;");
while($row =  $gsid_query->GoNext()) {
	$gs_id = intval(intval($row['gsid'])/10);
}
$this->DropQuery($gsid_query);
$checkQuery  = $this->NewQuery("select count(*) as num from player_item where refine_level>0 and price=0");
$result = $checkQuery->GoNext();
echo "{$gs_id} {$result['num']} records\n";

$this->DropQuery($checkQuery);

?>
