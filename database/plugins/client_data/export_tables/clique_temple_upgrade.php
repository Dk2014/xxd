<?php


    $export_table = array(
        'table'             => 'clique_temple_upgrade',
        'export_order'      => ' ORDER BY `level` asc',
        'exclude_columns'   => array(),
    );

    export_table($export_table);

?>
