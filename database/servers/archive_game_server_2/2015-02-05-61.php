<?php

$this->AddSQL("
update mission_enemy set enter_x=1853,enter_y=697 where id=1031;
update mission_enemy set enter_y=623 where id=478;
update mission_enemy set enter_y=757 where id=376;

");


$this->AddSQL("
update enemy_role set health=5000 where id=115;
");

$this->AddSQL("
	INSERT INTO `purchase_limit` (`id`, `item_id`, `num`)
	VALUES
	        (23,214,1);
");

?>
