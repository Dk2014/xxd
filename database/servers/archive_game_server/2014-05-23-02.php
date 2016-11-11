
<?php
db_execute($db, 
	"
		update `skill` set `config`='' where `id` in(31,66,88,92,100,104,106,108);

	"
);

?>