<?php


    $export_table = array(
        'table'             => 'clique_temple',
        'export_order'      => ' ORDER BY `worship_type` asc',
        'exclude_columns'   => array(),
    );

    export_table($export_table);

?>
