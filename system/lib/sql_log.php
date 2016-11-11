<?php
/*
 * web sql 操作记录，注： 只记录update，insert，delete操作
 *默认文件路径保存再当前目录下得sql_log文件夹中，也可以在构造函数中传入自定义得路径
 */
class sqlLog{
	private $logDir = './sql_log';
	private $sqlNote = '-- ';
	function __construct($dir){
		if(isset($dir) && !empty($dir)){
			$this->logDir = $dir;
		}
	}

	//递归创建文件夹
	private function createFolder($path)
 	{
  	if (!file_exists($path))
  	{
   		$this->createFolder(dirname($path));
   		mkdir($path, 0777);
  	}
 	}


	//根据小时获取当前要写入得文件路径
	private function getPath(){
		$dirFullPath = $this->logDir . date("/Y/m/d/");
		$fileName = date("H"). '.log';
		if(!is_dir($dirFullPath)){
				$this->createFolder($dirFullPath);
		}
		return fopen($dirFullPath.$fileName, "a+");
	}

	//写入sql语句
	function writeSql($sql, $comment){
		//1.$sql判断 2. getpath 3. 写入时间点和注释 4. $sql isarray 5.逐一再写入操作得SQL语句 6.写入空行 7.关闭文件句柄
		if(!isset($sql) || empty($sql)){
			return;
		}
		$fhandle = $this->getPath();
		$noteStr = "-- ------------------------------------------------------".PHP_EOL;
		$noteStr .= $this->sqlNote.date('m/d/Y H:i',time()).' '. $comment.' IP:'.$_SERVER["REMOTE_ADDR"].PHP_EOL;
		if(is_array($sql)){
			foreach ($sql as $value) {
				$noteStr .= $value.PHP_EOL;
			}
		}else{
			$noteStr .= $sql.PHP_EOL;
		}
		$noteStr .= PHP_EOL;
		fwrite($fhandle, $noteStr);
		fclose($fhandle);
	}
}


?>