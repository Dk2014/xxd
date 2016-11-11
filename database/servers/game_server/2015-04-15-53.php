<?php
$this->AddSQL("
delete   from player_driving_sword_event;
delete   from player_driving_sword_event_exploring;
delete   from player_driving_sword_event_treasure;
delete   from player_driving_sword_event_visiting;
delete   from player_driving_sword_map;

");
$this->AddSQL("
delete from player_item where item_id=727;
");

$this->AddSQL("
update player_driving_sword_info set current_cloud=1,highest_cloud=1,current_x=10,current_y=10,allowed_action = allowed_action+40 ;
");

?>
