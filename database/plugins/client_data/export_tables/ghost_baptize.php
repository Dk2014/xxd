<?php


$export_table = array(
    'table' 			=> 'ghost_baptize',
    'export_order' 		=> ' ORDER BY `star` ASC',
    'exclude_columns'	=> array(
        'probablity1',
        'min_add_growth1',
        'max_add_growth1',
        'probablity2' ,
        'min_add_growth2',
        'max_add_growth2',
    ),
);

export_table($export_table);

?>
