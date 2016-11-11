<?php
db_execute($db, "

INSERT INTO `ghost_star` (`star`, `need_fragment_num`, `growth`, `color`) VALUES
(1, 10, 10, 0),
(2, 30, 30, 1),
(3, 50, 50, 2),
(4, 80, 80, 3),
(5, 100, 100, 3);

");
?>


