
<?php
db_execute($db, "


ALTER  TABLE player_friend RENAME TO player_global_friend;
ALTER  TABLE player_friend_chat RENAME TO player_global_friend_chat;
ALTER  TABLE player_friend_state RENAME TO player_global_friend_state;


");
?>
