<?php

db_execute($db, "

    UPDATE `func` SET `lock`='1500' WHERE (`id`='7');

");
?>