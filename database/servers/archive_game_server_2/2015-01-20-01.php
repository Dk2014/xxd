<?php

$players = $this->NewQuery("select  a.id,a.user,a.nick,a.main_role_id,c.name,b.item_id,b.num from player a , player_item b ,item c  where a.id = b.pid and b.item_id=c.id and c.type_id =10;");

while ($player = $players->GoNext()) {
    $id = $player['id'];
    $user = $player['user'];
    $nick = $player['nick'];
    $mainid = $player['main_role_id'];
    $name = $player['name'];
    $itemid = $player['item_id'];
    $num = $player['num'];

    echo "id:{$id},user:{$user},nick:{$nick},mainid:{$mainid},name:{$name},itemid:{$itemid},num:{$num} \n";
}

$this->DropQuery($players);

$this->AddSQL(
"delete from `player_item` where exists (select * from item where id= player_item.item_id and type_id =10 );"
);

?>
