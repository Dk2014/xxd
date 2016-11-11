<?php

$this->AddSQL("
        alter table `level_star` drop column `two_star_score`;
        alter table `level_star` drop column `three_star_score`;
");

?>
