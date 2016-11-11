<?php
class Page {
	private $table;
	private $title;
	private $columns;

	function __construct($table) {
		$this->table = $table;
		$this->columns = array();
	}

	function SetTitle($title) {
		$this->title = $title;
	}

	function AddColumn($name, $text, $readonly) {
		$this->columns[] = array(
			'name' => $name,
			'text' => $text,
			'readonly' => $readonly,
		);
	}

	private function OnLoad() {
		if (isset($_GET['a'])) {
			$action = $_GET['a'];
			
			switch ($action) {
				case 'del':
					$id = $_GET['id'];
					$sql = "DELETE FROM ".$this->table." WHERE id = {$id}";
					header('Location:?m=del_done');
			}
		}
	}

	public function Rend() {
		$this->OnLoad();

		$rows = array(
			array('id' => '1', 'name' => '京城', 'lock' => 100),
			array('id' => '2', 'name' => '东瀛', 'lock' => 200),
		);
?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
	<title><?php echo $this->title; ?></title>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
	<script src="http://ajax.googleapis.com/ajax/libs/jqueryui/1.8.18/jquery-ui.min.js"></script>
	<style type="text/css" media="all">
	@import url("style.css");
	</style>
	<script type="text/javascript">
function confrim_link(text, url) {
	if (confirm(text)) {
		window.location.href = url;
	}
}

$(function(){
	$('.row_tool .edit').click(function(e){
		var div = e.target.parentNode;
		var td = div.parentNode;
		var tr = td.parentNode;
		tr.style.display = 'none';
		tr.nextElementSibling.className = 'e2';
		return false;
	});

	$('.row_tool .cancel').click(function(e){
		var div = e.target.parentNode;
		var td = div.parentNode;
		var tr = td.parentNode;
		tr.className = 'e';
		tr.previousElementSibling.style.display = '';
		return false;
	});

	$('.ckall').click(function(){
		$('.ck').attr('checked', this.checked);
	});

	$('.table_list tbody tr .ck').click(function(event){
		event.stopPropagation();
	});

	$('.table_list tbody tr').click(function(){
		var x = $('.ck', this);
		if (x.length > 0)
			x[0].checked =  !x[0].checked;
	});

	$('#apply').click(function(){
		if ($('#opers').value == $('#op-edit').value) {
			var rows = $('.table_list tbody tr');
			for (var i = 0; i < rows.length; i ++ ) {
				var x = $('.ck', rows[i]);
				if (x.length > 0 && x[0].checked)
					$('.edit', rows[i]).click();
			}
		}
	});

	$('#new').click(function(){
		$("#form1").css({
			position:"absolute",
			left: ($('#content').outerWidth() - $("#form1").outerWidth())/2+"px",
			top: "200px"
		});  
		$('#form1').fadeIn();
		$('#mark').fadeIn();
	});

	$('.form_close').click(function(e){
		$(e.target.parentNode).fadeOut();
		$('#mark').fadeOut();
		
	});
});
	</script>
</head>
<body>
	<div id="side_bar">
		<div id="nav_menu">
			<a href="#" class="current">城镇</a>
			<a href="#">副本</a>
			<a href="#">角色</a>
			<a href="#">装备</a>
			<a href="#">秘笈</a>
		</div>
	</div>
	<div id="content">
		<div class="head">
			<div><div class="title"><?php echo $this->title; ?></div><div class="navbar"><a href="#" class="current">城镇数据</a><a href="#">NPC数据</a></div></div>
			<div style="clear: both;"></div>
		</div>
		<div class="body">
			<div class="toolbar">
				<select id="opers">
					<option>批量操作</option>
					<option id="op-edit">编辑</option>
				</select>
				<input type="button" value="应用" id="apply" />
				<input type="button" value="新建" id="new" />
			</div>
			<div class="table_list">
				<table colspan="0" rowspan="0" cellpadding="0" cellspacing="0">
					<thead>
						<tr>
							<td style="width:20px"><input type="checkbox" name="ids" class="ck ckall" /></td><td style="width:50px;">ID</td>
<?php
		foreach ($this->columns as $column) {
?>
							<td><?php echo $column['text']; ?></td>
<?php
		}
?>
						</tr>
					</thead>
					<tbody>
<?php
		foreach ($rows as $row) {
?>
						<tr>
							<td><input type="checkbox" name="ids" class="ck"/></td>
							<td><?php echo $row['id']; ?></td>
<?php
			$first = true;
			foreach ($this->columns as $column) {
?>
							<td><?php echo htmlspecialchars($row[$column['name']]); ?> 
<?php
				if ($first) {
					$first = false;
?>
							<div class="row_tool"><a href="javascript:void(0)" class="edit">编辑</a><a href="javascript:confrim_link('确认删除此记录?', '?a=del&id=<?php echo $row['id']; ?>')">删除</a></div>
<?php
				}
?>
							</td>
<?php
			}
?>
						</tr>
						<tr class="e">
							<td></td>
							<td>1</td>
<?php
			$first = true;
			foreach ($this->columns as $column) {
?>
							<td><input type="text" value="<?php echo htmlspecialchars($row[$column['name']]); ?>" class="editor"/>
<?php
				if ($first) {
					$first = false;
?>
							<div class="row_tool"><a href="javascript:void(0)" class="cancel">取消</a></div>
<?php
				}
?>
							</td>
<?php
			}
?>
						</tr>
<?php
		}
?>
					</tbody>
					<tfoot>
						<tr>
							<td style="width:20px"><input type="checkbox" name="ids" class="ck ckall" /></td><td style="width:50px;">ID</td>
<?php
		foreach ($this->columns as $column) {
?>
							<td><?php echo $column['text']; ?></td>
<?php
		}
?>
						</tr>
					</tfoot>
				</table>
			</div>
		</div>
	</div>
	<div id="mark">
	</div>
	<div class="form" id="form1" style="width:600px;">
		<image class="form_close" src="close.png" />
		<div class="form_box">
			<form>
				<table style="width:100%">
					<tr><td class="label">城镇名称</td><td class="input"><input type="text" class="textbox" /></td><td class="info"></td></tr>
					<tr><td class="label">开放权值</td><td class="input"><input type="text" class="textbox" /></td><td class="info">玩家拥有的权限值大于等于城镇开放权值时方可进入城镇</td></tr>
				</table>
				<center class="buttons"><input type="submit" value="提交" /><input type="submit" value="取消" /></center>
			</form>
		</div>
	</div>
</body>
</html>
<?php
	}
}
?>