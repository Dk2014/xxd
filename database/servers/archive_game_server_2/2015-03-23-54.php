<?php 

$gs_id = 0;
//计算gsid
$gsid_query = $this->NewQuery("select id>>32 as gsid from player limit 1;");
while($row =  $gsid_query->GoNext()) {
	$gs_id = intval(intval($row['gsid'])/10);
}
$this->DropQuery($gsid_query);

if($gs_id==10001) {
	//安卓QQ1区删除 老的附件
	$this->AddSQL("delete  from global_mail_attachments where id<472493647200309;");
} else {
	//其他区清理所有
	$this->AddSQL("delete  from global_mail_attachments;");
	$this->AddSQL("delete  from global_mail;");
}

?>
