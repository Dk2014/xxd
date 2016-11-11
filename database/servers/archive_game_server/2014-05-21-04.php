<?php
db_execute($db, "
	delete from enemy_deploy_form where `pos1`=0 and `pos2`=0 and `pos3`=0 and `pos4`=0 and `pos5`=0 and `pos6`=0 
	and `pos7`=0 and `pos8`=0 and `pos9`=0 and `pos10`=0 and `pos11`=0 and `pos12`=0 and `pos13`=0 and `pos14`=0 and `pos15`=0;
");
?>