<?php
    $export_table = array(
        'table'             => 'random_award_box',
        'export_order'      => ' ORDER BY `mission_level_id`',
        'exclude_columns'   => array(
            'order',
            'award_chance',
            'must_in_first',
        ),
    );

    export_table($export_table);

?>
