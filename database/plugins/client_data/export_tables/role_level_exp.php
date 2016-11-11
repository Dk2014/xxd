<?php


$export_table = array(
  'table' 			=> 'role_level_exp',
  'export_order' 		=> ' ORDER BY `level`',
  'exclude_columns'	=> array(
        'id',
  ),
);

export_table($export_table);

?>
