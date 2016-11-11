<?php

$sql = 'select `pid`,`ghost_id` from `player_ghost` where `star`>1;';

$query = $this->NewQuery($sql);

$result = array();

$row = $query->GoNext();

while(isset($row)){
	if(isset($row['pid']) && isset($row['ghost_id'])){
		if(!isset($result[$row['pid']])){
			$result[$row['pid']] = array();
		}
		array_push($result[$row['pid']], $row['ghost_id']);
	}
	
	$row = $query->GoNext();
}
$this->DropQuery($query);
// 开始发邮件逻辑
foreach ($result as $pid => $ghost_ids) {
	$mail_auto_id = $this->GetAutoID($pid, 'player_mail');
	$timestamp = time();
	$this->AddSQL("insert into player_mail (id,pid,mail_id, state, send_time, parameters, have_attachment, title, content, expire_time, priority)" . 
						       " values({$mail_auto_id}, {$pid}, 0, 0, {$timestamp},  '', 1, '魂侍碎片补偿邮件', '在新的版本中，我们降低了魂侍1星升2星所需求的魂侍碎片数量，在此补偿各位大侠拥有的每个2星及以上的魂侍10个此魂侍的碎片。于此同时魂侍4星升5星的所需求的魂侍碎片数量增加了20个，所以拥有5星魂侍的玩家我们就不给于补偿了~但这样算下来大侠还净赚了10个碎片哟！', 1451577600, 1);");
	foreach ($ghost_ids as $id) {
		// 根据魂侍id去查找魂侍碎片id
		$now_sql = 'select fragment_id from `ghost` where id='.$id;
		$new_query = $this->NewQuery($now_sql);
		$row_fragment_id = $new_query->GoNext();
		if(isset($row_fragment_id)){
			$fragment_id = $row_fragment_id['fragment_id'];
			$mail_attach_auto_id = $this->GetAutoID($pid, 'player_mail_attachment');
			$this->AddSQL("insert into player_mail_attachment (id, pid, player_mail_id, attachment_type, item_id, item_num)" .
				  " values( {$mail_attach_auto_id}, {$pid}, {$mail_auto_id}, 0, {$fragment_id}, 10);");
		}
		$this->DropQuery($new_query);
	}
}

?>