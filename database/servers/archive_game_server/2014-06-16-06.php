
<?php
db_execute($db, 
"

ALTER TABLE `mission_level_small_box` CHANGE COLUMN `id` `id` int(11) NOT NULL AUTO_INCREMENT;
ALTER TABLE `mission_level_small_box_items` CHANGE COLUMN `id` `id` int(11) NOT NULL AUTO_INCREMENT;

"
);

?>

