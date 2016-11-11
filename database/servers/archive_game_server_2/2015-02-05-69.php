<?php

$this->AddSQL("
update item set type_id=9,can_use=1,can_batch=1 where id in (262, 233, 234);

INSERT INTO `item_box_content` (`id`, `item_id`, `type`, `mode`, `get_item_id`, `item_id_set`, `item_desc`, `min_num`, `max_num`, `probability`)
	VALUES
		(78,262,0,0,0,'','',5000,0,0);

INSERT INTO `item_box_content` (`id`, `item_id`, `type`, `mode`, `get_item_id`, `item_id_set`, `item_desc`, `min_num`, `max_num`, `probability`)
	VALUES
		(79,233,1,0,0,'','',50,0,0);

INSERT INTO `item_box_content` (`id`, `item_id`, `type`, `mode`, `get_item_id`, `item_id_set`, `item_desc`, `min_num`, `max_num`, `probability`)
	VALUES
		(80,234,1,0,0,'','',100,0,0);

");

