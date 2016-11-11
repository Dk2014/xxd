<?php
db_execute($db, "

INSERT INTO `ghost` (`name`, `sign`, `town_id`, `role_id`, `fragment_id`, `unique_key`, `init_star`, `health`, `attack`, `defence`, `speed`, `desc`) VALUES
('1', '2', 1, 3, 200, 1, 1, 3, 3, 3, 3, 'descr');

");
?>

