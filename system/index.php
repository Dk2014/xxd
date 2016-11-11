<?php
require_once 'lib/global.php';
include 'lib/sql_log.php';
date_default_timezone_set('Asia/Shanghai');
$sqlLog = new sqlLog($sql_log_config['dir']);
$sql = '';
// 判断是否登录
check_login();

if (isset($_GET['do'])) {
	if ($_GET['do'] == 'reboot') {
		file_put_contents(dirname(__FILE__) . "/reboot", '');
	} else if ($_GET['do'] == 'reload') {
		file_put_contents(dirname(__FILE__) . "/reload", '');
	} else if ($_GET['do'] == 'svn-update') {
		file_put_contents(dirname(__FILE__) . "/svn-update", dirname(dirname(__FILE__)));
	}

	header("Location:?{$_GET['ret']}");
	exit;
} else if (!isset($_GET['p'])) {
	header('Location:?p=town');
	exit;
}

$page_id = $_GET['p'];

$nav_links = array(
	array('title' => '城镇数据', 'links' => array(
		array('text' => '=基础功能=', 'type' => 'text'),
		array('text' => '功能开放', 'type' => 'link', 'id' => 'func', 'ids' => array('')),
		array('text' => '城镇', 'type' => 'link', 'id' => 'town', 'ids' => array('town', 'town_npc', 'mission', 'mission_map', 'mission_enemy', 'mission_npc')),
		array('text' => '灵宠', 'type' => 'link', 'id' => 'battle_pet', 'ids' => array('')),
		array('text' => '怪物', 'type' => 'link', 'id' => 'enemy_role', 'ids' => array('enemy_role')),
		array('text' => '绝招', 'type' => 'link', 'id' => 'skill', 'ids' => array('skill', 'skill_buddy')),
		array('text' => '角色', 'type' => 'link', 'id' => 'role', 'ids' => array('role')),
		array('text' => '队伍友情', 'type' => 'link', 'id' => 'teamship', 'ids' => array('teamship')),
		array('text' => '物品', 'type' => 'link', 'id' => 'item_type', 'ids' => array('item_type', 'item_material')),
		array('text' => '装备', 'type' => 'link', 'id' => 'item_equipment', 'ids' => array()),
		array('text' => '时装', 'type' => 'link', 'id' => 'fashion', 'ids' => array()),
		array('text' => '魂侍', 'type' => 'link', 'id' => 'ghost', 'ids' => array('')),
		array('text' => '剑心', 'type' => 'link', 'id' => 'sword_soul', 'ids' => array('sword_soul', 'sword_soul_type', 'sword_soul_level')),
		array('text' => '帮派', 'type' => 'link', 'id' => 'clique_building', 'ids' => array()),
		array('text' => '阵印', 'type' => 'link', 'id' => 'totem', 'ids' => array('')),
		array('text' => '天书', 'type' => 'link', 'id' => 'sealed_book', 'ids' => array('')),
		array('text' => '觉醒', 'type' => 'link', 'id' => 'awaken_attr', 'ids' => array('')),
	)
),
array('title' => '帮派', 'links' => array(
	array('text' => '=帮派相关=', 'type' => 'text'),
	array('text' => '建筑', 'type' => 'link', 'id' => 'clique_building', 'ids' => array()),
	array('text' => '送镖', 'type' => 'link', 'id' => 'clique_boat', 'ids' => array()),
)),
array('title' => '辅助功能', 'links' => array(
	array('text' => '=辅助功能=', 'type' => 'text'),
	array('text' => '铜币', 'type' => 'link', 'id' => 'coins_exchange', 'ids' => array()),
	array('text' => '邮件', 'type' => 'link', 'id' => 'mail', 'ids' => array('')),
	array('text' => '公告', 'type' => 'link', 'id' => 'announcement', 'ids' => array('')),
	array('text' => '随机商人', 'type' => 'link', 'id' => 'trader', 'ids' => array('')),
	array('text' => '世界消息模版', 'type' => 'link', 'id' => 'channel_message', 'ids' => array('')),
	array('text' => '内购产品信息', 'type' => 'link', 'id' => 'product_info', 'ids' => array('')),
)),

array('title' => '任务', 'links' => array(
	array('text' => '=任务相关=', 'type' => 'text'),
	array('text' => '主线任务', 'type' => 'link', 'id' => 'quest', 'ids' => array()),
	array('text' => '每日任务', 'type' => 'link', 'id' => 'daily_quest', 'ids' => array()),
	array('text' => '伙伴扩展任务', 'type' => 'link', 'id' => 'extend_quest', 'ids' => array()),
	array('text' => '支线任务', 'type' => 'link', 'id' => 'addition_quest&m=0', 'ids' => array()),
	array('text' => '每日任务星数奖励', 'type' => 'link', 'id' => 'quest_start_award', 'ids' => array()),
)),

array('title' => '关卡相关', 'links' => array(
	array('text' => '=系统关卡=', 'type' => 'text'),
	array('text' => '资源关卡', 'type' => 'link', 'id' => 'extend_level&m=1', 'ids' => array()),
	array('text' => '伙伴关卡', 'type' => 'link', 'id' => 'extend_level&m=9', 'ids' => array()),
	array('text' => '灵宠关卡', 'type' => 'link', 'id' => 'extend_level&m=10', 'ids' => array()),
	array('text' => '魂侍关卡', 'type' => 'link', 'id' => 'extend_level&m=11', 'ids' => array()),
	array('text' => '极限关卡(彩虹桥)', 'type' => 'link', 'id' => 'rainbow_level', 'ids' => array()),
	array('text' => '灵宠幻境关卡', 'type' => 'link', 'id' => 'pve_level', 'ids' => array()),
	array('text' => '云海御剑', 'type' => 'link', 'id' => 'driving_sword', 'ids' => array()),

	array('text' => '=互动=', 'type' => 'text'),
	array('text' => '声望', 'type' => 'link', 'id' => 'fame_system', 'ids' => array()),
	array('text' => '多人关卡', 'type' => 'link', 'id' => 'multi_level', 'ids' => array()),
	array('text' => '比武场', 'type' => 'link', 'id' => 'arena_award_box', 'ids' => array()),
)),

array('title' => '活动相关', 'links' => array(
	array('text' => '=抽奖=', 'type' => 'text'),
	array('text' => '爱心抽奖', 'type' => 'link', 'id' => 'heart_draw', 'ids' => array()),
	array('text' => '命锁宝箱', 'type' => 'link', 'id' => 'fate_box', 'ids' => array()),

	array('text' => '=其他=', 'type' => 'text'),
	array('text' => '游戏问答', 'type' => 'link', 'id' => 'faq', 'ids' => array()),
	array('text' => '平台好友奖励', 'type' => 'link', 'id' => 'platform_friend_award', 'ids' => array()),
	array('text' => '推送通知', 'type' => 'link', 'id' => 'push_notify', 'ids' => array()),
	array('text' => '打坐', 'type' => 'link', 'id' => 'meditation_exp', 'ids' => array()),
)),

array('title' => '运营', 'links' => array(
	array('text' => '=运营=', 'type' => 'text'),
	array('text' => '活动中心', 'type' => 'link', 'id' => 'quest_activity_center', 'ids' => array()),
	array('text' => '功能预告', 'type' => 'link', 'id' => 'func_prediction', 'ids' => array()),
	array('text' => '每日签到', 'type' => 'link', 'id' => 'daily_sign_in_award&m=0', 'ids' => array('')),
	array('text' => 'VIP', 'type' => 'link', 'id' => 'vip_level', 'ids' => array('')),
	array('text' => '购买体力设置', 'type' => 'link', 'id' => 'physical_buy_cost_config', 'ids' => array('')),
	array('text' => '购买比武场次数设置', 'type' => 'link', 'id' => 'arena_buy_cost_config', 'ids' => array('')),
	array('text' => '购买彩虹关卡扫荡次数设置', 'type' => 'link', 'id' => 'rainbow_buy_cost_config', 'ids' => array('')),
	array('text' => '购买云海御剑行动点设置', 'type' => 'link', 'id' => 'driving_sword_buy_cost_config', 'ids' => array('')),
	array('text' => '购买资源关卡次数设置', 'type' => 'link', 'id' => 'buy_resource_times_config', 'ids' => array('')),
	array('text' => '购买深渊关卡次数设置', 'type' => 'link', 'id' => 'buy_hard_level_times_config', 'ids' => array('')),
	array('text' => '购买boss关卡次数设置', 'type' => 'link', 'id' => 'buy_boss_level_times_config', 'ids' => array('')),
)),
);

require_once 'pages/' . $page_id . '.php';
if (file_exists("pages/{$page_id}_extend.php")) {
	require_once("pages/{$page_id}_extend.php");
} else {
	$extend_columns = array(); // 避免在下面反复判断是否存在
}

if (isset($_GET['template']) && intval($_GET['template']) == 2) {
	// 批量修改的导出配置正确判断
	if(isset($pconfig['batch_update_key']) && !empty($pconfig['batch_update_key'])){
		error_reporting(E_ERROR | E_WARNING | E_PARSE);
		if (isset($pconfig['where'])) {
			$func_where = $pconfig['where'];
			$sql_where = $func_where($_GET);
		} else {
			$sql_where = '';
		}
		//导出数据
		require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel.php";
		require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel/Writer/Excel5.php";
		$objPHPExcel = new PHPExcel();
		$objPHPExcel->setActiveSheetIndex(0);
		$objSheet = $objPHPExcel->getActiveSheet();
		$objSheet->setCellValue('A1', $pconfig['batch_update_key']);
		$objSheet->setCellValue('A2', 'KEY');
		$ind = ord('B');
		//设置标题头
		foreach ($pconfig['columns'] as $value) {
			$row = ($ind > ord('Z') ? ('A' . chr($ind - 26)) : chr($ind)) . '1';
			$objSheet->setCellValue($row, $value['name']);
			$ind++;
		}
		$ind2 = ord('B');
		foreach ($pconfig['columns'] as $value) {
			$row = ($ind2 > ord('Z') ? ('A' . chr($ind2 - 26)) : chr($ind2)) . '2';
			$objSheet->setCellValue($row, $value['text']);
			$ind2++;
		}
		//导出数据
		$sql = "SELECT * FROM `{$pconfig['table']}` $sql_where";
		// $order_pos = strpos($sql, 'order');
		// if($order_pos > 0){
		//     $sql = 
		// }
		// $sql .= " order by '{$pconfig['batch_update_key']}' ASC";
		$rows = db_query($db, $sql);
		array_unshift($pconfig['columns'], array('name' => $pconfig['batch_update_key']));
		foreach ($rows as $index => $row_value) {
			$ind3 = ord('A');
			foreach ($pconfig['columns'] as $value) {
				$row = ($ind3 > ord('Z') ? ('A' . chr($ind3 - 26)) : chr($ind3)) . ($index + 3);
				$real = $row_value[$value['name']];
				if (isset($extend_columns[$value['name']]) && count($extend_columns[$value['name']]['data']) > 0) {
					$real = $extend_columns[$value['name']]['data'][$real];
				}
				if (isset($extend_columns[$value['name']]['range'])) {
					$func = 'range_' . $value['name'];
					$arr = $func();
					if (isset($arr[$real])) {
						$real = $arr[$real];
					} else {
						$real = '无';
					}
				}
				$objSheet->setCellValue($row, $real);
				$ind3++;
			}
		}
		$objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel5');
		$fileName = $pconfig['table'] . date("Y-m-d") . ".xls";
		ob_end_clean();
		header("Pragma: public");
		header("Expires: 0");
		header("Cache-Control:must-revalidate, post-check=0, pre-check=0");
		header("Content-Type:application/force-download");
		header("Content-Type:application/vnd.ms-execl");
		header("Content-Type:application/octet-stream");
		header("Content-Type:application/download");
		header('Content-Disposition:attachment;filename="' . $fileName . '"');
		header("Content-Transfer-Encoding:binary");
		$objWriter->save('php://output');
		exit;
	}
}

if (isset($_GET['template']) && intval($_GET['template']) == 1) {
	require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel.php";
	require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel/Writer/Excel5.php";
	$objPHPExcel = new PHPExcel();
	$objPHPExcel->setActiveSheetIndex(0);
	$objSheet = $objPHPExcel->getActiveSheet();
	$ind = ord('A');
	foreach ($pconfig['columns'] as $value) {
		$row = ($ind > ord('Z') ? ('A' . chr($ind - 26)) : chr($ind)) . '1';
		$objSheet->setCellValue($row, $value['name']);
		$ind++;
	}
	$ind2 = ord('A');
	foreach ($pconfig['columns'] as $value) {
		$row = ($ind2 > ord('Z') ? ('A' . chr($ind2 - 26)) : chr($ind2)) . '2';
		$objSheet->setCellValue($row, $value['text']);
		$ind2++;
	}
	$objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel5');
	$fileName = $pconfig['table'] . "Template.xls";
	ob_end_clean();
	header("Pragma: public");
	header("Expires: 0");
	header("Cache-Control:must-revalidate, post-check=0, pre-check=0");
	header("Content-Type:application/force-download");
	header("Content-Type:application/vnd.ms-execl");
	header("Content-Type:application/octet-stream");
	header("Content-Type:application/download");
	header('Content-Disposition:attachment;filename="' . $fileName . '"');
	header("Content-Transfer-Encoding:binary");
	$objWriter->save('php://output');
	exit;
}

if ($_POST && !empty($_FILES['file1']['name'])) {
	if($_POST['sign_form2'] != $pconfig['batch_update_password']){
		print_r('批量修改密码错误~');
		exit;
	}
	if(isset($pconfig['batch_update_key']) && !empty($pconfig['batch_update_key'])){
		$columns = array();
		$sqls = array();
		$errorMessage = array();
		$new_file_name = dirname(__FILE__) . '/uploads/' . date("y-m-d-H-i-s") . '.xls';
		$re = move_uploaded_file($_FILES['file1']['tmp_name'], $new_file_name);
		if ($re) {
			//首先取出验证用的ID数组
			if (isset($pconfig['where'])) {
				$func_where = $pconfig['where'];
				$sql_where = $func_where($_GET);
			} else {
				$sql_where = '';
			}
			$sql = "SELECT * FROM `{$pconfig['table']}` $sql_where";
			$id_rows = db_query($db, $sql);
			require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel.php';
			require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel/IOFactory.php';
			require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel/Reader/Excel5.php';
			$objReader = PHPExcel_IOFactory::createReader('Excel5');
			$objPHPExcel = $objReader->load($new_file_name);
			$sheet = $objPHPExcel->getSheet(0);
			$highestRow = $sheet->getHighestRow();
			$highestColumn = PHPExcel_Cell::columnIndexFromString($sheet->getHighestColumn());
			for ($j = 0; $j < $highestColumn; $j++) {
				array_push($columns, $sheet->getCellByColumnAndRow($j, 1)->getValue());
			}
			$diffs = get_diff($columns, $pconfig, array($pconfig['batch_update_key']));
			header("Content-type: text/html; charset=utf-8");
			if (count($diffs) > 0) {
				print_r("导入的数据缺失必要字段(" . implode(",", $diffs) . ")，可重新下载原数据<a href='#' onclick='window.open(\"" . $_SERVER['PHP_SELF'] . '/?' . $_SERVER['QUERY_STRING'] . "&template=2\")'>原数据</a>");
				exit;
			}
			for ($i = 3; $i <= $highestRow; $i++) {
				$key = intval($sheet->getCellByColumnAndRow(0, $i)->getValue());
				if($id_rows[$i-3][$pconfig['batch_update_key']] != $key){
					print_r("验证到第".$i."行有问题，KEY无法对应上，现在的是".$key.",程序期待的是".$id_rows[$i-3][$pconfig['batch_update_key']]."，请检查!");
					exit;
				}
				$sql_template = "UPDATE `{$pconfig['table']}` SET ??? WHERE `{$pconfig['batch_update_key']}` = {$key}";
				$temp_sql = '';
				$isInsert = true;
				for ($j = 1; $j < $highestColumn; $j++) {
					$content = trim($sheet->getCellByColumnAndRow($j, $i)->getValue());
					//$content = iconv('GBK','UTF-8',$content);
					$name = $columns[$j];
					//执行验证
					if (isset($extend_columns) && isset($extend_columns[$name])) {
						//特殊的值的话，单独有插入前的预处理方法
						if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['handle']) && $extend_columns[$name]['handle']) {
							$func_handle_name = 'handle_' . $columns[$j];
							$content = $func_handle_name($content);
							if ($content === false) {
								array_push($errorMessage, sprintf('第%d条第%d列数据有误', $i - 2, $j + 1));
								$isInsert = false;
							}
						} else {

							//1 范围验证
							if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['data']) && count($extend_columns[$name]['data']) > 0) {
								$search = is_in($content, $extend_columns[$name]['data']);
								if ($search === false) {
									array_push($errorMessage, sprintf('第%d条数据出错,第[%d]列数据不在范围内(%s)%s', $i - 2, $j + 1, implode(",", $extend_columns[$name]['data']), $content));
									$isInsert = false;
								} else {
									$content = $search;
								}
							}

							if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['range'])) {
								$func_range_name = 'range_' . $name;
								$ranges = $func_range_name();
								$search = is_in($content, $ranges);
								if ($search === false) {
									array_push($errorMessage, sprintf('第%d条数据出错,第[%d]列数据不在范围内', $i - 2, $j + 1));
									$isInsert = false;
								} else {
									$content = $search;
								}
							}
							//2 正则验证
							if (isset($extend_columns[$name]['regexp'])) {
								if (!preg_match($extend_columns[$name]['regexp']['reg'], $content)) {
									array_push($errorMessage, sprintf('第%d条数据出错,第[%s]列数据格式不对,%s', $i - 2, $j + 1, $extend_columns[$name]['regexp']['text']));
									$isInsert = false;
								}
							}
						}

					}
					if (!empty($name)) {
						$temp_sql .= "`$name` = '" . $db->real_escape_string($content) . "',";
					}
				}
				$temp_sql = substr($temp_sql, 0, strlen($temp_sql) - 1);
				$sql_template = str_replace('???', $temp_sql, $sql_template);
				$sql_template .= ';';
				if ($isInsert) {
					array_push($sqls, $sql_template);
				}
			}
			if (count($errorMessage) > 0) {
				foreach ($errorMessage as $message) {
					print_r($message . "<br />");
				}
				exit;
			} else {
				$error = db_insert_transaction($db, $sqls);
				if (!empty($error)) {
					print_r($error);
					exit;
				} else {
					$sqlLog->writeSql($sqls, sprintf("批量更新表%s", $pconfig['table']));
				}
			}
			//unlink($new_file_name);
			header("Location:" . $_SERVER['PHP_SELF'] . '?' . $_SERVER['QUERY_STRING']);
		} else {
			print_r("上传文件无法完成操作,请尝试重新上传");
			exit;
		}
	}
}


if ($_POST && !empty($_FILES['file']['name'])) {

	$columns = array();
	$sqls = array();
	$errorMessage = array();
	$new_file_name = dirname(__FILE__) . '/uploads/' . date("y-m-d-H-i-s") . '.xls';
	$re = move_uploaded_file($_FILES['file']['tmp_name'], $new_file_name);
	if ($re) {
		require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel.php';
		require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel/IOFactory.php';
		require_once dirname(__FILE__) . '/lib/phpExcel/PHPExcel/Reader/Excel5.php';
		$objReader = PHPExcel_IOFactory::createReader('Excel5');
		$objPHPExcel = $objReader->load($new_file_name);
		$sheet = $objPHPExcel->getSheet(0);
		$highestRow = $sheet->getHighestRow();
		$highestColumn = PHPExcel_Cell::columnIndexFromString($sheet->getHighestColumn());
		for ($j = 0; $j < $highestColumn; $j++) {
			array_push($columns, $sheet->getCellByColumnAndRow($j, 1)->getValue());
		}
		$diffs = get_diff($columns, $pconfig, array());
		header("Content-type: text/html; charset=utf-8");
		if (count($diffs) > 0) {
			print_r("模版缺失必要字段(" . implode(",", $diffs) . ")，可重新下载<a href='#' onclick='window.open(\"" . $_SERVER['PHP_SELF'] . '/?' . $_SERVER['QUERY_STRING'] . "&template=1\")'>模版</a>");
			exit;
		}

		for ($i = 3; $i <= $highestRow; $i++) {
			$sql = '';
			if (isset($pconfig['insert'])) {
				$func_insert = $pconfig['insert'];
				$insert_sql = $func_insert($_GET);
				$sql .= "INSERT INTO `{$pconfig['table']}` SET $insert_sql, ";
			} else {
				$sql .= "INSERT INTO `{$pconfig['table']}` SET ";
			}
			$isInsert = true;
			for ($j = 0; $j < $highestColumn; $j++) {
				$content = trim($sheet->getCellByColumnAndRow($j, $i)->getValue());
				//$content = iconv('GBK','UTF-8',$content);
				$name = $columns[$j];
				//执行验证
				if (isset($extend_columns) && isset($extend_columns[$name])) {
					//特殊的值的话，单独有插入前的预处理方法
					if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['handle']) && $extend_columns[$name]['handle']) {
						$func_handle_name = 'handle_' . $columns[$j];
						$content = $func_handle_name($content);
						if ($content === false) {
							array_push($errorMessage, sprintf('第%d条第%d列数据有误', $i - 2, $j + 1));
							$isInsert = false;
						}
					} else {

						//1 范围验证
						if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['data']) && count($extend_columns[$name]['data']) > 0) {
							$search = is_in($content, $extend_columns[$name]['data']);
							if ($search === false) {
								array_push($errorMessage, sprintf('第%d条数据出错,第[%d]列数据不在范围内(%s)%s', $i - 2, $j + 1, implode(",", $extend_columns[$name]['data']), $content));
								$isInsert = false;
							} else {
								$content = $search;
							}
						}

						if (is_array($extend_columns[$name]) && isset($extend_columns[$name]['range'])) {
							$func_range_name = 'range_' . $name;
							$ranges = $func_range_name();
							$search = is_in($content, $ranges);
							if ($search === false) {
								array_push($errorMessage, sprintf('第%d条数据出错,第[%d]列数据不在范围内', $i - 2, $j + 1));
								$isInsert = false;
							} else {
								$content = $search;
							}
						}
						//2 正则验证
						if (isset($extend_columns[$name]['regexp'])) {
							if (!preg_match($extend_columns[$name]['regexp']['reg'], $content)) {
								array_push($errorMessage, sprintf('第%d条数据出错,第[%s]列数据格式不对,%s', $i - 2, $j + 1, $extend_columns[$name]['regexp']['text']));
								$isInsert = false;
							}
						}
					}

				}
				if (!empty($name)) {
					$sql .= "`$name` = '" . $db->real_escape_string($content) . "',";
				}
			}
			$sql = substr($sql, 0, strlen($sql) - 1);
			$sql .= ';';
			if ($isInsert) {
				array_push($sqls, $sql);
			}
		}
		if (count($errorMessage) > 0) {
			foreach ($errorMessage as $message) {
				print_r($message . "<br />");
			}
			exit;
		} else {
			$error = db_insert_transaction($db, $sqls);
			if (!empty($error)) {
				print_r($error);
				exit;
			} else {
				$sqlLog->writeSql($sqls, sprintf("批量倒入表%s", $pconfig['table']));
			}
		}
		//unlink($new_file_name);
		header("Location:" . $_SERVER['PHP_SELF'] . '?' . $_SERVER['QUERY_STRING']);
	} else {
		print_r("上传文件无法完成操作,请尝试重新上传");
		exit;
	}
}

//数据导出
if (isset($_GET['export'])) {
	error_reporting(E_ERROR | E_WARNING | E_PARSE);
	if (isset($pconfig['where'])) {
		$func_where = $pconfig['where'];
		$sql_where = $func_where($_GET);
	} else {
		$sql_where = '';
	}
	//导出数据
	require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel.php";
	require_once dirname(__FILE__) . "/lib/phpExcel/PHPExcel/Writer/Excel5.php";
	$objPHPExcel = new PHPExcel();
	$objPHPExcel->setActiveSheetIndex(0);
	$objSheet = $objPHPExcel->getActiveSheet();
	$ind = ord('A');
	//设置标题头
	foreach ($pconfig['columns'] as $value) {
		$row = ($ind > ord('Z') ? ('A' . chr($ind - 26)) : chr($ind)) . '1';
		$objSheet->setCellValue($row, $value['name']);
		$ind++;
	}
	$ind2 = ord('A');
	foreach ($pconfig['columns'] as $value) {
		$row = ($ind2 > ord('Z') ? ('A' . chr($ind2 - 26)) : chr($ind2)) . '2';
		$objSheet->setCellValue($row, $value['text']);
		$ind2++;
	}
	//导出数据
	$sql = "SELECT * FROM `{$pconfig['table']}` $sql_where";
	$rows = db_query($db, $sql);
	foreach ($rows as $index => $row_value) {
		$ind3 = ord('A');
		foreach ($pconfig['columns'] as $value) {
			$row = ($ind3 > ord('Z') ? ('A' . chr($ind3 - 26)) : chr($ind3)) . ($index + 3);
			$real = $row_value[$value['name']];
			if (isset($extend_columns[$value['name']]) && count($extend_columns[$value['name']]['data']) > 0) {
				$real = $extend_columns[$value['name']]['data'][$real];
			}
			if (isset($extend_columns[$value['name']]['range'])) {
				$func = 'range_' . $value['name'];
				$arr = $func();
				if (isset($arr[$real])) {
					$real = $arr[$real];
				} else {
					$real = '无';
				}
			}
			$objSheet->setCellValue($row, $real);
			$ind3++;
		}
	}
	$objWriter = PHPExcel_IOFactory::createWriter($objPHPExcel, 'Excel5');
	$fileName = $pconfig['table'] . date("Y-m-d") . ".xls";
	ob_end_clean();
	header("Pragma: public");
	header("Expires: 0");
	header("Cache-Control:must-revalidate, post-check=0, pre-check=0");
	header("Content-Type:application/force-download");
	header("Content-Type:application/vnd.ms-execl");
	header("Content-Type:application/octet-stream");
	header("Content-Type:application/download");
	header('Content-Disposition:attachment;filename="' . $fileName . '"');
	header("Content-Transfer-Encoding:binary");
	$objWriter->save('php://output');
	exit;
}

if (isset($_POST['save'])) {
	if (isset($pconfig['eventsjson']))  {      
		$josn_string = file_get_contents($pconfig['eventsjson']);
		$obj=json_decode($josn_string);
		$data=$obj->JsonEvents;
		$row_col1s = $_POST[$pconfig['columns'][0]['name']];
		for ($i = 0; $i < count($row_col1s); $i++) {
			$type = $pconfig['columns'][0]['name'];
			$page = $pconfig['columns'][4]['name'];
			$flag=1;
			foreach ( $data as $unit ) {
				if ($_POST[$type][$i] == $unit->Type  && $_POST[$page][$i] == $unit->Page) { 
					$flag=2;
				} 
			}

			if ($flag == 1) { //新纪录，需要添加                    
				$StartUnixTime = $pconfig['columns'][1]['name'];
				$EndUnixTime = $pconfig['columns'][2]['name'];
				$DisposeUnixTime = $pconfig['columns'][3]['name'];
				$IsRelative = $pconfig['columns'][5]['name'];
				$Weight = $pconfig['columns'][6]['name'];
				$Tag = $pconfig['columns'][7]['name'];
				$LTitle = $pconfig['columns'][8]['name'];
				$RTitle = $pconfig['columns'][9]['name'];
				$Content = $pconfig['columns'][10]['name'];

				$var01=$_POST[$type][$i];
				$var02=$_POST[$page][$i];
				$var03=$_POST[$StartUnixTime][$i];
				$var04=$_POST[$EndUnixTime][$i];
				$var05=$_POST[$DisposeUnixTime][$i];
				$var06=$_POST[$IsRelative][$i];
				$var07=$_POST[$Weight][$i];
				$var08=$_POST[$Tag][$i];
				$var09=$_POST[$LTitle][$i];
				$var10=$_POST[$RTitle][$i];
				$var11=$_POST[$Content][$i];

				$test = "{
					\"Type\":{$var01},
						\"Page\":{$var02},
						\"StartUnixTime\":{$var03},
						\"EndUnixTime\":{$var04},
						\"DisposeUnixTime\":{$var05},
						\"IsRelative\":{$var06},
						\"Weight\": {$var07},
						\"Tag\": {$var08},
						\"LTitle\": \"{$var09}\",
						\"RTitle\": \"{$var10}\",
						\"Content\": \"{$var11}\",
						\"AwardList\":[] }";
					$test = json_decode($test);

					array_push($data,$test);             
			}     
		}
		$obj->JsonEvents= $data;
		file_put_contents($pconfig['eventsjson'], json_encode($obj,JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE));      
	} else {
		$sql = '';
		if ($_GET['p'] == 'enemy_deploy_form' && !isset($_POST[$pconfig['columns'][0]['name']])) {
			$func_insert = $pconfig['insert'];
			$insert_sql = $func_insert($_GET);
			$sql .= "INSERT INTO `{$pconfig['table']}` SET $insert_sql;";
		} else {
			$row_ids = isset($_POST['id']) ? $_POST['id'] : array();
			$row_col1s = $_POST[$pconfig['columns'][0]['name']];

			for ($i = 0; $i < count($row_col1s); $i++) {
				if ($i < count($row_ids))
					$sql .= "UPDATE `{$pconfig['table']}` SET ";
				else {
					if (isset($pconfig['insert'])) {
						$func_insert = $pconfig['insert'];
						$insert_sql = $func_insert($_GET);

						$sql .= "INSERT INTO `{$pconfig['table']}` SET $insert_sql, ";
					} else
						$sql .= "INSERT INTO `{$pconfig['table']}` SET ";
				}

				$totalColumn = count($pconfig['columns']);
				$remainColumn = $totalColumn;
				for ($j = 0; $j < $totalColumn; $j++) {
					if (isset($pconfig['columns'][$j]['readonly'])) {
						$readonly = $pconfig['columns'][$j]['readonly'];
						if($readonly) {
							$remainColumn--;
						}
					}
				}

				for ($j = 0; $j < count($pconfig['columns']); $j++) {
					if (isset($pconfig['columns'][$j]['readonly'])) {
						$readonly = $pconfig['columns'][$j]['readonly'];
						if ($readonly) {
							continue;
						}
					}
					$remainColumn--;
					$name = $pconfig['columns'][$j]['name'];
					$sql .= "`$name` = '" . $db->real_escape_string(trim($_POST[$name][$i])) . "'";
					if ($remainColumn > 0) 
						$sql .= ', ';
					else
						$sql .= ' ';
				}

				if ($i < count($row_ids))
					$sql .= "WHERE `id` = " . $row_ids[$i];

				$sql .= ';';
			}
		}
	  //print($sql);
		//exit();

		db_execute($db, $sql);
		$sqlLog->writeSql($sql, sprintf("插入或更新发生在表%s", $pconfig['table']));
	}
	$addr_m = '';
	if (isset($_GET['m'])) {
		$addr_m .= '&m=' . $_GET['m'];
	}
	if (isset($_GET['m2'])) {
		$addr_m .= '&m2=' . $_GET['m2'];
	}
	if (isset($_GET['bt'])) {
		$addr_m .= '&bt=' . $_GET['bt'];
	}

	$addr_extra = '';

	if (isset($pconfig['addr'])) {
		$func_addr = $pconfig['addr'];
		$addr_extra = $func_addr($_GET);
	}

	header('Location:?p=' . $page_id . $addr_m . $addr_extra);
	exit;
}

if (isset($_POST['delete'])) {
	$sql = '';
	$row_ids = $_POST['del_id'];

	$del_failed = 0;

	for ($i = 0; $i < count($row_ids); $i++) {
		if (!isset($pconfig['del_check']) || $pconfig['del_check']($row_ids[$i])) {
			$sql .= "DELETE FROM `{$pconfig['table']}` WHERE `id` = " . $row_ids[$i] . ';';
		} else {
			$del_failed = 1;
		}
	}
	if ($sql != '') {
		db_execute($db, $sql);
		$sqlLog->writeSql($sql, sprintf("删除操作发生在表%s", $pconfig['table']));
	}

	$addr_m = '';
	if (isset($_GET['m'])) {
		$addr_m .= '&m=' . $_GET['m'];
	}
	if (isset($_GET['m2'])) {
		$addr_m .= '&m2=' . $_GET['m2'];
	}

	if (isset($_GET['bt'])) {
		$addr_m .= '&bt=' . $_GET['bt'];
	}

	$func_addr = $addr_extra = "";
	if (isset($pconfig['addr'])) {
		$func_addr = $pconfig['addr'];
		$addr_extra = $func_addr($_GET);
	}

	header('Location:?del_failed=' . $del_failed . '&p=' . $page_id . $addr_m . $addr_extra);
	exit;

}

// 增加条件查询 
if (isset($pconfig['where'])) {
	$func_where = $pconfig['where'];
	$sql_where = $func_where($_GET);
} else {
	$sql_where = '';
}

//添加是否需要分页选择
$pagesize = 20; //默认每页数据量
$ispaged = false;
$pagenow = 1;
$pagecount = 1;
if (isset($pconfig['pagesize'])) {
	$pagesize = $pconfig['pagesize'];
}

$result = 0;
if (isset($pconfig['eventsjson'])) {
	$josn_string = file_get_contents($pconfig['eventsjson']);
	$obj=json_decode($josn_string);
	$data=$obj->JsonEvents;
	foreach ( $data as $unit ){
		$result++;
	}
} else {
	$count_all_sql = "SELECT COUNT(*) from `{$pconfig['table']}` $sql_where";
	$result = db_query_array($db, $count_all_sql);
}
$size = $result[0][0];
if ($size > $pagesize) {
	$ispaged = true;
	$pagenow = (isset($_GET['page']) && intval($_GET['page']) > 0) ? intval($_GET['page']) : 1;
	$pagecount = ($size % $pagesize == 0 ? intval($size / $pagesize) : (intval($size / $pagesize) + 1));
}
if (isset($pconfig['eventsjson'])) {
	$filename = $pconfig['eventsjson'];
	$josn_string = file_get_contents($filename);
	$obj=json_decode($josn_string);
	$data=$obj->JsonEvents;
	$i=0;
	foreach ( $data as $unit ){
		$i++;
		$rows[$i]['Type']=$unit->Type;
		$rows[$i]['Page']=$unit->Page;
		$rows[$i]['StartUnixTime']=$unit->StartUnixTime;
		$rows[$i]['EndUnixTime']=$unit->EndUnixTime;
		$rows[$i]['DisposeUnixTime']=$unit->DisposeUnixTime;
		$rows[$i]['IsRelative']=$unit->IsRelative;
		$rows[$i]['Weight']=$unit->Weight;
		$rows[$i]['Tag']=$unit->Tag;
		$rows[$i]['LTitle']=$unit->LTitle;
		$rows[$i]['RTitle']=$unit->RTitle;
		$rows[$i]['Content']=$unit->Content;
	}
} else {
	$sel_all_sql = "select * from `{$pconfig['table']}` $sql_where" . ($ispaged ? (' LIMIT ' . (($pagenow - 1) * $pagesize) . ',' . $pagesize) : '');
	$rows = db_query($db, $sel_all_sql);
}
?>
<?php
function make_tip($column, $row)
{
?>
<?
	if (isset($column['tip_render']) || hava_extend_column($column['name'], 'tip_render')) {
?>
	<span style="display:none" id="tip_<?= $row['id']; ?>_<?= $column['name'] ?>">
<?php
		if (isset($column['tip_render'])) {
			print($column['tip_render']($row));
		} else {
			print(execute_extend_column('tip_render', $column, $row));
		}

?></span>
	<script>
		$(function () {
			$('#row_<?= $row['id']; ?>_<?= $column['name'] ?>').qtip({
				content: document.getElementById('tip_<?= $row['id']; ?>_<?= $column['name'] ?>').innerHTML,
					position: { adjust: {screen: true} },
					style: { name: 'light', tip: 'topLeft', border: { width: 3, radius: 8 }, width: { max: 1000 } }
			})
		});
		</script>
<?
	}
?>
<?php
}

?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title><?= $pconfig['title']; ?></title>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
<script src="jquery.min.js"></script>
<script src="jquery-ui.min.js"></script>
<script src="jquery.qtip-1.0.0-rc3.js"></script>
<script src="jquery.autocomplete.js"></script>
<link rel="stylesheet" href="jquery-ui.css"/>
<style type="text/css" media="all">
    @import url("style.css");
</style>
<?
if (isset($style_text)) {
?>
    <style type="text/css">
	<?= $style_text; ?>
    </style>
<?
}
?>
<script type="text/javascript">
function resize_role_player(id, width, height) {
	//var swf = document.getElementById('role_player'); 
	//swf.style.width = width + 'px'; 
	//    swf.style.height = height + 'px'; 
}

function del_new_row(e) {
	e.parentNode.parentNode.parentNode.removeChild(e.parentNode.parentNode);
}

function set_table_event() {
	// 点击行的时候自动选取行

	$('.table_list tbody tr').click(function () {
		var x = $('.ck', this);
		if (x.length > 0)
			x[0].checked = !x[0].checked;
	});

	// 禁用此事件的冒泡，否则会跟上一个事件处理有冲突
	$('.table_list tbody tr .ck').click(function (event) {
		event.stopPropagation();
	});

	$('.table_list tbody tr').hover(
		function () {
			$(this).css("background", "#f0f0f0");
		},
			function () {
				$(this).removeAttr("style");
			}
	);
}

var changes = [];

function addToChanges(t) {
	had = false;
	for (var ii in changes) {
		if (changes[ii].html() == t.html()) {
			had = true;
			break;
		}
	}
	if (!had) {
		changes.push(t);
	}
}

// 新行的临时ID
var new_row_id = 0;
var enemy_roles = [];

$(function () {
	set_table_event();

	// 点击全选按钮
	$('.ckall').click(function () {
		$('.ck').attr('checked', this.checked);
	});

	// 全选的情况下，取消任何一条记录，全选记号删除
	$('.ck').click(function () {
		if ($(this).attr('checked') == undefined && $('.ckall').attr('checked') == 'checked') {
			$('.ckall').attr('checked', false);
		}
	})

		// 点击删除按钮，需要判断是否选中某条记录
		$('.del').click(function () {
			var cb = $('.table_list tbody').find('input.ck:checked');
			if (cb.length < 1) {
				alert("请选择删除项");
				return false;
			}

			if (confirm("确定删除?") == true)
				return true;
			else
				return false;
		})

			// 点击编辑按钮
			$('.edit').click(function () {
				var rows = $('.table_list tbody tr');
				for (var i = 0; i < rows.length; i++) {
					var x = $('.ck', rows[i]);
					if (x.length) {
						rows[i].style.display = 'none';
						rows[i].nextElementSibling.className = 'e2';
					}
				}
				$('.normal_buttons').css({display: 'none'});
				$('.editing_buttons').fadeIn();
				$('.ck').css({display: 'none'});

				InitTableHead();
			});

	// 点击取消编辑按钮
	$('.cancel').click(function () {
		var rows = $('.table_list tbody tr');
		for (var i = 0; i < rows.length; i++) {
			var x = $('.ck', rows[i]);
			if (x.length == 0) {
				if (rows[i].className == 'new_row') {
					rows[i].parentNode.removeChild(rows[i]);
				} else if (rows[i].className == 'e2') {
					rows[i].className = 'e';
					rows[i].previousElementSibling.style.display = '';
				}
			}
		}
		$('.normal_buttons').fadeIn();
		$('.editing_buttons').css({display: 'none'});
		$('.ck').css({display: ''});

		InitTableHead();
	});

	// 点击新建按钮
	$('.new').click(function () {
		new_row_id += 1;

		a = $('#new_row_tpl')[0].textContent.replace(/##id##/g, new_row_id);

		$('#table_body').append(a.substring(10, a.length - 4));

		$('.normal_buttons').css({display: 'none'});
		$('.editing_buttons').fadeIn();
		$('.ck').css({display: 'none'});

		register_events();

		InitTableHead();

		$body = (window.opera) ? (document.compatMode == "CSS1Compat" ? $('html') : $('body')) : $('html,body');
		$body.animate({scrollTop: $('.new_row').offset().top}, 500);
	});

    /*
	    $("#form").submit(function(){
     if (changes.length == 0) {
     $('.cancel').click();
     return false;
     }

     $('#table_body tr:odd').each(function(i){
     isIn = false;
     for (var ii in changes) {
     if ( $(this).html() == changes[ii].html() ){
     isIn = true;
     break;
     }
     }
     if (!isIn) {
     $(this).remove();
     }
     })
     })
     */
	// 点击浮动窗口关闭按钮
	$('.form_close').click(function (e) {
		$(e.target.parentNode).fadeOut();
		$('#mark').fadeOut();
	});

	$('.ajax_tip').qtip({
		content: {
			url: function (self) {
				return 'ajax.php' + self.options.show.when.target.context.search;
			}
		}, position: {
			adjust: {screen: true}
		}, style: {
			name: 'light',
				tip: 'topLeft',
				border: { width: 3, radius: 8 },
				width: {max: 1000}
		}
	});

    /*
    <?
    foreach ($pconfig['columns'] as $column) {
	if (isset($column['hide']))
	    continue;
    ?>


    <?
	    if (!isset($column['tip_render'])) {
	?>
     $('.data_table_col_
    <?= $column['name']; ?>').qtip({
     content: '
    <?= $column['text']; ?>',
     position:{ corner: {target: 'bottomLeft'} },
     style: { name:'light', tip:'topLeft', border:{ width:3, radius:8 }, width:200 }
     });
    <?
	}
    ?>


    <?
	}
	?>
     */

    /*
    <?
    if (isset($pconfig['js'])) {
    ?>


    <?= $pconfig['js']($_GET); ?>


    <?
	}
	?>
     */

<?
if (isset($_GET['del_failed']) && $_GET['del_failed'] == 1) {
?>
    alert('数据被其他地方引用，无法直接删除，请先删除引用再删除数据。');
<?
}
?>
})

function register_events() {
<?
if (isset($pconfig['js'])) {
?>
    <?= $pconfig['js']($_GET); ?>
<?
}
?>
    //$('#data_table').fixHeader({height:700});
    /*
     $('input').change(function(e){
     if (!$(this).getAttribute('event_registed')) {
     $(this).setAttribute('event_registed', true)
     addToChanges($(this).parent().parent())
     }
     });

     $('select').change(function(e){
     if (!$(this).getAttribute('event_registed')) {
     $(this).setAttribute('event_registed', true)
     addToChanges($(this).parent().parent())
     }
     });

     $('textarea').change(function(e){
     if (!$(this).getAttribute('event_registed')) {
     $(this).setAttribute('event_registed', true)
     addToChanges($(this).parent().parent())
     }
     });
     */
}

$(register_events);
</script>
</head>
<body onclick="$('#side_bar').css('display','none');">
<form id="form" action="<?= $_SERVER['PHP_SELF'] . '?' . $_SERVER['QUERY_STRING']; ?>" method="post"
      enctype="multipart/form-data">
<script>
var side_bar_tigger_timeout;

function side_bar_tigger_mouseover() {
	side_bar_tigger_timeout = setTimeout("side_bar_tigger_timeout = null; $('#side_bar').css('display','');", 250);
}

function side_bar_tigger_mouseout() {
	if (side_bar_tigger_timeout) {
		clearTimeout(side_bar_tigger_timeout);
	}
}
</script>
<div onmouseover="side_bar_tigger_mouseover()" onmouseout="side_bar_tigger_mouseout()"
     style="background:#e0e0e0;height:14px;z-index:2;position:fixed;width:100%;cursor:pointer;"></div>
<div id="side_bar" style="display:none">
<?php
    /*
    <div class="admin">用户名: <?= $_SESSION['username'] ?></div>
    <div class="logout"><a href="logout.php">退出系统</a></div>
     */
?>
    <div id="nav_menu">
	<table style="width:100%">
	    <tr>
<?
foreach ($nav_links as $group) {
?>
		    <td valign="top" style="border-right:solid 1px #e0e0e0;">
<?
	foreach ($group['links'] as $link) {
?>
<?
		if ($link['type'] == "link") {
?>
	<div><a href="?p=<?= $link['id'];  ?> "<?
	if (in_array($page_id, $link['ids']))
		echo 'class="current"';
?>><?= $link['text']; ?></a></div>
<?
		} else {
?>
				<a style="color:#BC0000;"><?= $link['text']; ?></a>
<?
		}
?>
<?
	}
?>
		    </td>
<?
}
?>
		<td valign="top">
		    <a href="?do=reload&ret=<? echo urlencode($_SERVER['QUERY_STRING']); ?>" style="color:#BC0000;">[重载数据]</a>
		    <a href="?do=reboot&ret=<? echo urlencode($_SERVER['QUERY_STRING']); ?>" style="color:#BC0000;">[重启游戏]</a>
		    <!-- <a href="?do=svn-update&ret=<? echo urlencode($_SERVER['QUERY_STRING']); ?>" style="color:#BC0000;">[更新文件]</a> -->
<?
if (exec('hg branch') == 'default') {
?>
		    <a href="log.php?srv=log/srv1&prefix=gs" style="color:#BC0000;">[srv1(gs)]</a>
		    <a href="log.php?srv=log/srv10&prefix=gs" style="color:#BC0000;">[srv10(gs)]</a>
		    <a href="log.php?srv=log/srv2&prefix=hd" style="color:#BC0000;">[srv2(互动)]</a>
<?
}
?>

<?
if (exec('hg branch') == 'soha') {
?>
		    <a href="log.php?srv=gslog&prefix=gs" style="color:#BC0000;">[越南内网开发(gs)]</a>
		    <a href="log.php?srv=hdlog&prefix=gs" style="color:#BC0000;">[越南内网开发(互动)]</a>
<?
}
?>
		    <a href="json_editor/json_editor.php" style="color:#BC0000;">[活动配置编辑器]</a>
		</td>
	    </tr>
	</table>
    </div>
</div>
<div id="content">
<div class="head">
    <div>
<?
if (isset($pconfig['location']) && $pconfig['location'] != '') {
?>
	    <div class="location"><?= $pconfig['location']($_GET); ?></div>
<?
}
?>
	<div class="title"><?= $pconfig['title']; ?></div>
	<div class="navbar">
<?
foreach ($pconfig['links'] as $link) {
?>
	<a href="?p=<?= $link['id']; ?>"<?
	if ("p=" . $link['id'] == $_SERVER['QUERY_STRING'])
		echo 'class="current"';
?>><?= $link['text']; ?></a>
<?
}
?>
	</div>
    </div>
    <div style="clear: both;"></div>
</div>

<?php
if (isset($pconfig['nested_page'])) {
	require_once('pages/' . $pconfig['nested_page'] . '.php');
} else {
?>
<div class="body">
<div class="toolbar" id="real_table_head">
    <div class="normal_buttons">
<?
	if (!isset($pconfig['not_delete'])) {
?>
	    <input type="submit" value="删除" class="del" name="delete"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_edit'])) {
?>
	    <input type="button" value="编辑" class="edit"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_new'])) {
?>
	    <input type="button" value="新建" class="new"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_template'])) {
?>
	    <span style="margin-left:25px">导入：</span>
	    <a href="#" class="template" onclick="javascript:window.open('/?<?
		echo $_SERVER['QUERY_STRING'];
?>&template=1')">1.下载模版</a> > 2.<input id="template" type="file" name="file" class="import"/> > 3.<input
		type="submit" value="上传" name="import"
		onclick="if(!document.getElementById('template').value){alert('please choose template first!');return false;}document.getElementById('sign_form').value=1;"/>
	    <input type="hidden" id="sign_form" value="0"/>
	    <font color="red"> // </font>
<?
	}
?>
	<?php if(isset($pconfig['batch_update_key']) && !empty($pconfig['batch_update_key'])){ ?>
	 <span style="margin-left:15px">批量修改:</span>
	 <a href="#" class="template" onclick="javascript:window.open('/?<?
	echo $_SERVER['QUERY_STRING'];
?>&template=2')">1.下载原数据</a> > 2.<input id="template2" type="file" name="file1" class="import"/> > 3.<input
		type="submit" value="上传" name="import"
		onclick="if(!document.getElementById('template2').value){alert('please choose data file first!');return false;}document.getElementById('sign_form2').value=prompt('输入批量修改的密码');"/>
	    <input type="hidden" id="sign_form2" name="sign_form2" value="0"/>
	    <font color="red"> // </font>
	<?php } ?>

	<?php if (!isset($pconfig['allow_export']) || !$pconfig['dis_export']) { ?> <input type="button"
												       value="导出数据"
												       onclick="javascript:window.open('./?<?
	echo $_SERVER['QUERY_STRING'];
	?>&export=1')" /><?php } ?>
<?
	if (isset($pconfig['remark'])) {
?>
	<span style=\"display:block;color:#666;margin-top:10px\"> <?
		echo($pconfig['remark']);
?> </span>
<?
	}
?>
    </div>
    <div class="editing_buttons" style="display: none">
	<input type="submit" value="提交" class="save" name="save"/>
	<input type="button" value="取消" class="cancel"/>
<?
	if (!isset($pconfig['not_new'])) {
?>
	    <input type="button" value="新建" class="new"/>
<?
	}
?>
<?
	if (isset($pconfig['remark'])) {
?>
	<span style=\"display:block;color:#666;margin-top:10px\"> <?
		echo($pconfig['remark']);
?> </span>
<?
	}
?>
    </div>
</div>
<div class="table_list">
<table id="data_table" colspan="0" rowspan="0" cellpadding="0" cellspacing="0" style="width:100%">
<thead>
<tr>
    <td style="width:20px"><input type="checkbox" class="ck ckall"/></td>
<?
	if (!empty($pconfig['show_order'])) {
?>
	<td style="width:50px;">编号</td>
<?
	}
?>
<?
	$row_index_mess_wtf = 1;
?>
<?
	foreach ($pconfig['columns'] as $column) {
		if (isset($column['hide']))
			continue;
?>
	<td class="data_table_col" style="<?
		if (isset($column['width']))
			echo 'width:' . $column['width'];
		?>"><?= $column['text']; ?></td>
<?
	}
?>
<?
	if (isset($pconfig['opera'])) {
?>
	<td>操作</td>
<?
	}
?>
</tr>
</thead>
<tbody id="table_body">
<?
	foreach ($rows as $key => $row) {
?>
	<tr <?
		echo isset($pconfig['row_attr_render']) ? $pconfig['row_attr_render']($row) : '';
?>

<?
		if (isset($pconfig['row_highlight']) && $pconfig['row_highlight']($row)) {
			?> class="highlight" <?
		}
?>
	>
	<td
<?
		if (isset($pconfig['highlight_special_quest']) && $pconfig['highlight_special_quest']($row)) {
			?> class="highlight_special_quest" <?
		}
?>
	><input type="checkbox" name="del_id[]" class="ck" value="<?= $row['id']; ?>"/> <?
	echo $row_index_mess_wtf;;
?>  </td>
<?
		if (!empty($pconfig['show_order'])) {
?>
	<td><?
			echo $key + 1;
?></td>
<?
		}
?>
<?
		foreach ($pconfig['columns'] as $column) {
			if (isset($column['hide']))
				continue;
?>
<?
			if (isset($column['render']) || hava_extend_column($column['name'], 'render')) {
?>
		<td class="data_table_col data_table_col_<?= $column['name'] ?>"
		    id="row_<?= $row['id']; ?>_<?= $column['name'] ?>">
<?php
				if (isset($column['render'])) {
					print($column['render']($row));
				} else {
					print(execute_extend_column('render', $column, $row));
				}
?>
<?
				make_tip($column, $row);
?>
		</td>
<?
			} else {
?>
		<td class="data_table_col data_table_col_<?= $column['name'] ?>"
		    id="row_<?= $row['id']; ?>_<?= $column['name'] ?>">
		    <?= htmlspecialchars($row[$column['name']]); ?>
<?
				make_tip($column, $row);
?>
		</td>
<?
			}
?>
<?
		}
?>
<?
		if (isset($pconfig['opera'])) {
?>
	    <td>
<?
			foreach ($pconfig['opera'] as $opera) {
				echo $opera['render']($row);
			}
?>
	    </td>
<?
		}
?>
    </tr>
    <tr class="e">
<?
		$have_id = false;
		foreach ($pconfig['columns'] as $column) {
			if ($column['name'] == 'id') {
				$have_id = true;
				break;
			}
		}
?>
	<td><?
		if (!$have_id) {
			?><input type="hidden" name="id[]" value="<?= $row['id']; ?>" /><?
		}
		?> <?
		echo $row_index_mess_wtf++;
?></td>
<?
		if (!empty($pconfig['show_order'])) {
?>
	    <td>编号</td>
<?
		}
?>
<?
		foreach ($pconfig['columns'] as $column) {
			if (isset($column['hide']))
				continue;
?>
<?
			if (isset($column['editor']) || hava_extend_column($column['name'], 'editor')) {
?>
		<td class="data_table_col data_table_col_<?= $column['name'] ?>"
		    id="row_<?= $row['id']; ?>_<?= $column['name'] ?>">
<?php
				if (isset($column['editor'])) {
					print($column['editor']($row));
				} else {
					print(execute_extend_column('editor', $column, $row));
				}
?>
<?
				make_tip($column, $row);
?>
		</td>
<?
			} else {
?>
		<td class="data_table_col data_table_col_<?= $column['name'] ?>"
		    id="row_<?= $row['id']; ?>_<?= $column['name'] ?>">
		    <input name="<?= $column['name']; ?>[]" type="text"
		    value="<?= htmlspecialchars($row[$column['name']]); ?>" class="editor" style="<?
		    if (isset($column['width']))
			    echo 'overflow: hidden;width:' . $column['width'];
?>"/>
<?
make_tip($column, $row);
?>
		</td>
<?
			}
?>
<?
		}
?>
<?
		if (isset($pconfig['opera'])) {
?>
	    <td>
<?
			foreach ($pconfig['opera'] as $opera) {
				echo $opera['render']($row, true);
			}
?>
	    </td>
<?
		}
?>
    </tr>
<?
	}
?>
</tbody>
<tfoot>
<tr>
    <td style="width:20px"><input type="checkbox" class="ck ckall"/>
<?
	if (!empty($pconfig['show_order'])) {
?>
    <td>编号</td>
<?
	}
?>
<?
	foreach ($pconfig['columns'] as $column) {
		if (isset($column['hide']))
			continue;
?>
	<td><?= $column['text']; ?></td>
<?
	}
?>
<?
	if (isset($pconfig['opera'])) {
?>
	<td>操作</td>
<?
	}
?>
</tr>
</tfoot>
</table>
<div id="table_head" style="display:none; position:absolute; top:0;padding-top:8px">
    <div class="toolbar">
	<div class="normal_buttons">
<?
	if (!isset($pconfig['not_delete'])) {
?>
		<input type="submit" value="删除" class="del" name="delete"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_edit'])) {
?>
		<input type="button" value="编辑" class="edit"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_new'])) {
?>
		<input type="button" value="新建" class="new"/>
<?
	}
?>


	</div>
	<div class="editing_buttons" style="display: none">
	    <input type="submit" value="提交" class="save" name="save"/>
	    <input type="button" value="取消" class="cancel"/>
<?
	if (!isset($pconfig['not_new'])) {
?>
		<input type="button" value="新建" class="new"/>
<?
	}
?>
	</div>
    </div>
    <div class="item"><span><input type="checkbox" class="ck ckall"/></span></div>
<?
	if (!empty($pconfig['show_order'])) {
?>
	<div class="item"><span>编号</span></div>
<?
	}
?>
<?
	foreach ($pconfig['columns'] as $column) {
		if (isset($column['hide']))
			continue;
?>
	<div class="item"><span><?= $column['text']; ?></span></div>
<?
	}
?>
<?
	if (isset($pconfig['opera'])) {
?>
	<div class="item"><span>操作</span></div>
<?
	}
?>
</div>
	<script>
	function InitTableHead() {
		var tds1 = $('#data_table thead td');
		var tds2 = $('#table_head .item');
		var x = 0;

		for (var i = 0; i < tds1.length; i++) {
			tds2[i].style.width = (tds1[i].offsetWidth - 17) + "px";
			tds2[i].style.height = (tds1[i].offsetHeight - 17) + "px";
			tds2[i].style.left = x + "px";
			x += tds1[i].offsetWidth;
		}

		$('#table_head').css('width', x);
	}
	;

	$(document).ready(InitTableHead);
	$(window).resize(InitTableHead);

	$(window).scroll(function () {
		var $window = $(window);

		var window_left = $window.scrollLeft();
		var window_top = $window.scrollTop();
		var offset = $('#real_table_head').offset();
		var left = offset.left;
		var top = offset.top;
		var element = $('#table_head')[0];

		if (top >= window_top && top <= window_top + $window.height()) {
			element.style.display = 'none';
			return;
		}

		element.style.display = 'block';

		if (window.XMLHttpRequest) {
			element.style.position = "fixed";
			element.style.top = "0px";
		} else {
			element.style.top = window_top + "px";
		}

		element.style.left = ($('#data_table').offset().left - $window.scrollLeft()) + "px";
	});
	</script>
</div>
<div class="toolbar" style="padding-bottom:8px">
    <div class="normal_buttons">
<?
	if (!isset($pconfig['not_delete'])) {
?>
	    <input type="submit" value="删除" class="del" name="delete"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_edit'])) {
?>
	    <input type="button" value="编辑" class="edit"/>
<?
	}
?>
<?
	if (!isset($pconfig['not_new'])) {
?>
	    <input type="button" value="新建" class="new"/>
<?
	}
?>
<?
	if (isset($pconfig['remark'])) {
?>
	<span style=\"display:block;color:#666;margin-top:10px\"> <?
		echo($pconfig['remark']);
?> </span>
<?
	}
?>
    </div>
    <div class="editing_buttons" style="display: none">
	<input type="submit" value="提交" class="save" name="save"/>
	<input type="button" value="取消" class="cancel"/>
<?
	if (!isset($pconfig['not_new'])) {
?>
	    <input type="button" value="新建" class="new"/>
<?
	}
?>

<?
	if (isset($pconfig['remark'])) {
?>
	<span style=\"display:block;color:#666;margin-top:10px\"> <?
		echo($pconfig['remark']);
?> </span>
<?
	}
?>

    </div>
    <?php if ($pagecount > 1) { ?>
	<div class="page<?php echo $pagecount ?>" style="text-align : right">
	    <?php if ($pagenow > 1) { ?> <a
		href="<?php echo $_SERVER['PHP_SELF'] . '?' . (isset($_GET['page']) ? preg_replace("/&page=\d+$|&page=\d+&/i", ('&page=' . ($pagenow - 1)), $_SERVER['QUERY_STRING']) : $_SERVER['QUERY_STRING'] . '&page=' . ($pagenow - 1)) ?>">
		    < </a><?php } ?>
<?php for ($i = 1; $i <= $pagecount; $i++) {
	echo ($pagenow == $i ? $i : '<a href="' . $_SERVER['PHP_SELF'] . '?' . (isset($_GET['page']) ? preg_replace("/&page=\d+$|&page=\d+&/i", ('&page=' . $i), $_SERVER['QUERY_STRING']) : $_SERVER['QUERY_STRING'] . '&page=' . $i) . '" >' . $i . '</a>') . " ";
} ?>
	    <?php if ($pagenow < $pagecount) { ?> <a
		href="<?php echo $_SERVER['PHP_SELF'] . '?' . (isset($_GET['page']) ? preg_replace("/&page=\d+$|&page=\d+&/i", ('&page=' . ($pagenow + 1)), $_SERVER['QUERY_STRING']) : $_SERVER['QUERY_STRING'] . '&page=' . ($pagenow + 1)) ?>">
		    > </a><?php } ?>
	</div>
    <?php } ?>
</div>
</div>
</div>
<div id="mark">
</div>
<div class="form" id="form1" style="width:600px;">
    <image class="form_close" src="close.png"/>
    <div class="form_box">
	<form>
	    <table style="width:100%">
		<tr>
		    <td class="label">城镇名称</td>
		    <td class="input"><input type="text" class="textbox"/></td>
		    <td class="info"></td>
		</tr>
		<tr>
		    <td class="label">开放权值</td>
		    <td class="input"><input type="text" class="textbox"/></td>
		    <td class="info">玩家拥有的权限值大于等于城镇开放权值时方可进入城镇</td>
		</tr>
	    </table>
	    <center class="buttons"><input type="submit" value="提交"/><input type="submit" value="取消"/></center>
	</form>
    </div>
</div>
<?php
	if (isset($pconfig['foot']))
		$pconfig['foot']();
?>
<textarea id="new_row_tpl"><![CDATA[
    <tr class="new_row">
	<td>
	    <image class="del_row" src="close.png" onclick="del_new_row(this)"/>
	</td>
<?
        if (!empty($pconfig['show_order'])) {
            ?>
            <td>编号</td>
        <?
        }
        ?>
        <?
        foreach ($pconfig['columns'] as $column) {
            if (isset($column['hide']))
                continue;
            ?>
            <?
            if (isset($column['editor']) || hava_extend_column($column['name'], 'editor')) {
                ?>
                <td><?php
                    if (isset($column['editor'])) {
                        print(htmlspecialchars($column['editor'](null), ENT_NOQUOTES, "utf-8"));
                    } else {
                        $row = null;
                        print(htmlspecialchars(execute_extend_column('editor', $column, $row), ENT_NOQUOTES, "utf-8"));
                    }

                    ?></td>
            <?
            } else {
                ?>
                <td><input name="<?= $column['name']; ?>[]" type="text" class="editor" style="<?
                    if (isset($column['width']))
                        echo 'overflow: hidden;width:' . $column['width'];
                    ?>" <?
                    echo(isset($column['default']) ? ' value="' . $column['default'] . '"' : '');
                    ?>/></td>
            <?
            }
            ?>
        <?
        }
        ?>
    </tr>
    ]]></textarea>
</form>
<?php
}
?>
</body>
</html>
<?php
function reload()
{
    $request = array(
        'method' => "Reload",
        'params' => array(
            (object)array()
        ),
        'id' => 1
    );
    $jsonRequest = json_encode($request);
    $ctx = stream_context_create(array(
        'http' => array(
            'method' => 'POST',
            'timeout' => 10,
            'header' => 'Content-Type: application/json\r\n',
            'content' => $jsonRequest
        )
    ));

    $fp = fopen('http://127.0.0.1:7002', 'r', false, $ctx);
    fpassthru($fp);
    $buffer = fgets($fp, 4096);

    fclose($fp);
}

?>
