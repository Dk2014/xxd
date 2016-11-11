<?php

$this->AddSQL("

update item_box_content set min_num=6500 where id=78;

INSERT INTO `item_box_content` (`id`, `item_id`, `type`, `mode`, `get_item_id`, `item_id_set`, `item_desc`, `min_num`, `max_num`, `probability`)
	VALUES
		(81,233,0,0,0,'','',1500,0,0);

INSERT INTO `item_box_content` (`id`, `item_id`, `type`, `mode`, `get_item_id`, `item_id_set`, `item_desc`, `min_num`, `max_num`, `probability`)
	VALUES
		(82,234,0,0,0,'','',3000,0,0);

");

