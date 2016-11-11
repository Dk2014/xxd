<?php
db_execute($db, "

INSERT INTO `func` (`name`, `sign`, `lock`, `level`, `unique_key`, `need_play`)
VALUES
	('比武场', 'FUNC_ARENA', 5000, 0, 128, 1);


");