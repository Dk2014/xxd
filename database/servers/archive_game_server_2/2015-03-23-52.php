<?php 
function get_cost($level, $quality, $refine_level) {
	$cost = 0;
	if( $level == 1 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 600;
		case 9:
		$cost += 540;
		case 8:
		$cost += 500;
		case 7:
		$cost += 440;
		case 6:
		$cost += 400;
		case 5:
		$cost += 360;
		case 4:
		$cost += 300;
		case 3:
		$cost += 240;
		case 2:
		$cost += 220;
		case 1:
		$cost += 200;
		}
		return $cost;
	}
	if( $level == 20 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 3000;
		case 9:
		$cost += 2700;
		case 8:
		$cost += 2500;
		case 7:
		$cost += 2200;
		case 6:
		$cost += 2000;
		case 5:
		$cost += 1800;
		case 4:
		$cost += 1500;
		case 3:
		$cost += 1200;
		case 2:
		$cost += 1100;
		case 1:
		$cost += 1000;
		}
		return $cost;
	}
	if( $level == 40 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 6000;
		case 9:
		$cost += 5400;
		case 8:
		$cost += 5000;
		case 7:
		$cost += 4400;
		case 6:
		$cost += 4000;
		case 5:
		$cost += 3600;
		case 4:
		$cost += 3000;
		case 3:
		$cost += 2400;
		case 2:
		$cost += 2200;
		case 1:
		$cost += 2000;
		}
		return $cost;
	}
	if( $level == 60 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 9000;
		case 9:
		$cost += 8100;
		case 8:
		$cost += 7500;
		case 7:
		$cost += 6600;
		case 6:
		$cost += 6000;
		case 5:
		$cost += 5400;
		case 4:
		$cost += 4500;
		case 3:
		$cost += 3600;
		case 2:
		$cost += 3300;
		case 1:
		$cost += 3000;
		}
		return $cost;
	}
	if( $level == 80 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 12000;
		case 9:
		$cost += 10800;
		case 8:
		$cost += 10000;
		case 7:
		$cost += 8800;
		case 6:
		$cost += 8000;
		case 5:
		$cost += 7200;
		case 4:
		$cost += 6000;
		case 3:
		$cost += 4800;
		case 2:
		$cost += 4400;
		case 1:
		$cost += 4000;
		}
		return $cost;
	}
	if( $level == 100 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 15000;
		case 9:
		$cost += 13500;
		case 8:
		$cost += 12500;
		case 7:
		$cost += 11000;
		case 6:
		$cost += 10000;
		case 5:
		$cost += 9000;
		case 4:
		$cost += 7500;
		case 3:
		$cost += 6000;
		case 2:
		$cost += 5500;
		case 1:
		$cost += 5000;
		}
		return $cost;
	}
	if( $level == 500 && $quality == 1) {
		switch  ($refine_level) {
		case 10:
		$cost += 15000;
		case 9:
		$cost += 13500;
		case 8:
		$cost += 12500;
		case 7:
		$cost += 11000;
		case 6:
		$cost += 10000;
		case 5:
		$cost += 9000;
		case 4:
		$cost += 7500;
		case 3:
		$cost += 6000;
		case 2:
		$cost += 5500;
		case 1:
		$cost += 5000;
		}
		return $cost;
	}
	if( $level == 1 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 15000;
		case 9:
		$cost += 13500;
		case 8:
		$cost += 12500;
		case 7:
		$cost += 11000;
		case 6:
		$cost += 10000;
		case 5:
		$cost += 9000;
		case 4:
		$cost += 7500;
		case 3:
		$cost += 6000;
		case 2:
		$cost += 5500;
		case 1:
		$cost += 5000;
		}
		return $cost;
	}
	if( $level == 10 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 30000;
		case 9:
		$cost += 27000;
		case 8:
		$cost += 25000;
		case 7:
		$cost += 22000;
		case 6:
		$cost += 20000;
		case 5:
		$cost += 18000;
		case 4:
		$cost += 15000;
		case 3:
		$cost += 12000;
		case 2:
		$cost += 11000;
		case 1:
		$cost += 10000;
		}
		return $cost;
	}
	if( $level == 40 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 60000;
		case 9:
		$cost += 54000;
		case 8:
		$cost += 50000;
		case 7:
		$cost += 44000;
		case 6:
		$cost += 40000;
		case 5:
		$cost += 36000;
		case 4:
		$cost += 30000;
		case 3:
		$cost += 24000;
		case 2:
		$cost += 22000;
		case 1:
		$cost += 20000;
		}
		return $cost;
	}
	if( $level == 60 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 90000;
		case 9:
		$cost += 81000;
		case 8:
		$cost += 75000;
		case 7:
		$cost += 66000;
		case 6:
		$cost += 60000;
		case 5:
		$cost += 54000;
		case 4:
		$cost += 45000;
		case 3:
		$cost += 36000;
		case 2:
		$cost += 33000;
		case 1:
		$cost += 30000;
		}
		return $cost;
	}
	if( $level == 80 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 120000;
		case 9:
		$cost += 108000;
		case 8:
		$cost += 100000;
		case 7:
		$cost += 88000;
		case 6:
		$cost += 80000;
		case 5:
		$cost += 72000;
		case 4:
		$cost += 60000;
		case 3:
		$cost += 48000;
		case 2:
		$cost += 44000;
		case 1:
		$cost += 40000;
		}
		return $cost;
	}
	if( $level == 100 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 150000;
		case 9:
		$cost += 135000;
		case 8:
		$cost += 125000;
		case 7:
		$cost += 110000;
		case 6:
		$cost += 100000;
		case 5:
		$cost += 90000;
		case 4:
		$cost += 75000;
		case 3:
		$cost += 60000;
		case 2:
		$cost += 55000;
		case 1:
		$cost += 50000;
		}
		return $cost;
	}
	if( $level == 500 && $quality == 2) {
		switch  ($refine_level) {
		case 10:
		$cost += 150000;
		case 9:
		$cost += 135000;
		case 8:
		$cost += 125000;
		case 7:
		$cost += 110000;
		case 6:
		$cost += 100000;
		case 5:
		$cost += 90000;
		case 4:
		$cost += 75000;
		case 3:
		$cost += 60000;
		case 2:
		$cost += 55000;
		case 1:
		$cost += 50000;
		}
		return $cost;
	}
	if( $level == 1 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 60000;
		case 9:
		$cost += 54000;
		case 8:
		$cost += 50000;
		case 7:
		$cost += 44000;
		case 6:
		$cost += 40000;
		case 5:
		$cost += 36000;
		case 4:
		$cost += 30000;
		case 3:
		$cost += 24000;
		case 2:
		$cost += 22000;
		case 1:
		$cost += 20000;
		}
		return $cost;
	}
	if( $level == 20 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 90000;
		case 9:
		$cost += 81000;
		case 8:
		$cost += 75000;
		case 7:
		$cost += 66000;
		case 6:
		$cost += 60000;
		case 5:
		$cost += 54000;
		case 4:
		$cost += 45000;
		case 3:
		$cost += 36000;
		case 2:
		$cost += 33000;
		case 1:
		$cost += 30000;
		}
		return $cost;
	}
	if( $level == 40 && $quality == 3) {
		switch ($refine_level) {
		case 10:
		$cost += 120000;
		case 9:
		$cost += 108000;
		case 8:
		$cost += 100000;
		case 7:
		$cost += 88000;
		case 6:
		$cost += 80000;
		case 5:
		$cost += 72000;
		case 4:
		$cost += 60000;
		case 3:
		$cost += 48000;
		case 2:
		$cost += 44000;
		case 1:
		$cost += 40000;
		}
		return $cost;
	}
	if( $level == 60 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 150000;
		case 9:
		$cost += 135000;
		case 8:
		$cost += 125000;
		case 7:
		$cost += 110000;
		case 6:
		$cost += 100000;
		case 5:
		$cost += 90000;
		case 4:
		$cost += 75000;
		case 3:
		$cost += 60000;
		case 2:
		$cost += 55000;
		case 1:
		$cost += 50000;
		}
		return $cost;
	}
	if( $level == 80 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 210000;
		case 9:
		$cost += 189000;
		case 8:
		$cost += 175000;
		case 7:
		$cost += 154000;
		case 6:
		$cost += 140000;
		case 5:
		$cost += 126000;
		case 4:
		$cost += 105000;
		case 3:
		$cost += 84000;
		case 2:
		$cost += 77000;
		case 1:
		$cost += 70000;
		}
		return $cost;
	}
	if( $level == 100 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 300000;
		case 9:
		$cost += 270000;
		case 8:
		$cost += 250000;
		case 7:
		$cost += 220000;
		case 6:
		$cost += 200000;
		case 5:
		$cost += 180000;
		case 4:
		$cost += 150000;
		case 3:
		$cost += 120000;
		case 2:
		$cost += 110000;
		case 1:
		$cost += 100000;
		}
		return $cost;
	}
	if( $level == 500 && $quality == 3) {
		switch  ($refine_level) {
		case 10:
		$cost += 300000;
		case 9:
		$cost += 270000;
		case 8:
		$cost += 250000;
		case 7:
		$cost += 220000;
		case 6:
		$cost += 200000;
		case 5:
		$cost += 180000;
		case 4:
		$cost += 150000;
		case 3:
		$cost += 120000;
		case 2:
		$cost += 110000;
		case 1:
		$cost += 100000;
		}
		return $cost;
	}
	if( $level == 1 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 120000;
		case 9:
		$cost += 108000;
		case 8:
		$cost += 100000;
		case 7:
		$cost += 88000;
		case 6:
		$cost += 80000;
		case 5:
		$cost += 72000;
		case 4:
		$cost += 60000;
		case 3:
		$cost += 48000;
		case 2:
		$cost += 44000;
		case 1:
		$cost += 40000;
		}
		return $cost;
	}
	if( $level == 20 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 180000;
		case 9:
		$cost += 162000;
		case 8:
		$cost += 150000;
		case 7:
		$cost += 132000;
		case 6:
		$cost += 120000;
		case 5:
		$cost += 108000;
		case 4:
		$cost += 90000;
		case 3:
		$cost += 72000;
		case 2:
		$cost += 66000;
		case 1:
		$cost += 60000;
		}
		return $cost;
	}
	if( $level == 40 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 240000;
		case 9:
		$cost += 216000;
		case 8:
		$cost += 200000;
		case 7:
		$cost += 176000;
		case 6:
		$cost += 160000;
		case 5:
		$cost += 144000;
		case 4:
		$cost += 120000;
		case 3:
		$cost += 96000;
		case 2:
		$cost += 88000;
		case 1:
		$cost += 80000;
		}
		return $cost;
	}
	if( $level == 60 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 300000;
		case 9:
		$cost += 270000;
		case 8:
		$cost += 250000;
		case 7:
		$cost += 220000;
		case 6:
		$cost += 200000;
		case 5:
		$cost += 180000;
		case 4:
		$cost += 150000;
		case 3:
		$cost += 120000;
		case 2:
		$cost += 110000;
		case 1:
		$cost += 100000;
		}
		return $cost;
	}
	if( $level == 80 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 360000;
		case 9:
		$cost += 324000;
		case 8:
		$cost += 300000;
		case 7:
		$cost += 264000;
		case 6:
		$cost += 240000;
		case 5:
		$cost += 216000;
		case 4:
		$cost += 180000;
		case 3:
		$cost += 144000;
		case 2:
		$cost += 132000;
		case 1:
		$cost += 120000;
		}
		return $cost;
	}
	if( $level == 100 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 480000;
		case 9:
		$cost += 432000;
		case 8:
		$cost += 400000;
		case 7:
		$cost += 352000;
		case 6:
		$cost += 320000;
		case 5:
		$cost += 288000;
		case 4:
		$cost += 240000;
		case 3:
		$cost += 192000;
		case 2:
		$cost += 176000;
		case 1:
		$cost += 160000;
		}
		return $cost;
	}
	if( $level == 500 && $quality == 4) {
		switch  ($refine_level) {
		case 10:
		$cost += 480000;
		case 9:
		$cost += 432000;
		case 8:
		$cost += 400000;
		case 7:
		$cost += 352000;
		case 6:
		$cost += 320000;
		case 5:
		$cost += 288000;
		case 4:
		$cost += 240000;
		case 3:
		$cost += 192000;
		case 2:
		$cost += 176000;
		case 1:
		$cost += 160000;
		}
		return $cost;
	}
	return 0;
}

$sql =  'select  pi.id as id, pi.refine_level as refine_level, item.quality as quality, item.level as level from player_item pi left join item 
	on  pi.item_id = item.id where refine_level > 0 and pi.price=0;';

$query = $this->NewQuery($sql);
while($row = $query->GoNext()) {
	$id = $row['id'];
	$quality = $row['quality'];
	$level = $row['level'];
	$refine_level = $row['refine_level'];

	$cost = get_cost($level, $quality, $refine_level);
	$this->AddSQL("update player_item set `price`={$cost} where id={$id};");
}
$this->DropQuery($query);


//$gs_id = 0;
////计算gsid
//$gsid_query = $this->NewQuery("select id>>32 as gsid from player limit 1;");
//while($row =  $gsid_query->GoNext()) {
//	$gs_id = intval(intval($row['gsid'])/10);
//}
//$this->DropQuery($gsid_query);
//sleep(10);
//$checkQuery  = $this->NewQuery("select count(*) as num from player_item where refine_level>0 and price=0");
//$result = $checkQuery->GoNext();
//echo "{$gs_id} {$result['num']} records\n";
//
//$this->DropQuery($checkQuery);

?>
