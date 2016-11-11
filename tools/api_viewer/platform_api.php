<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
 <head>
  <title>平台API接口</title>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <style>
 li {list-style-type:none;}
</style>
</head>
<body>
<?php

$APIConfigs = array(
	array(
		'name' => "serverall",
		'in'   => array(
				array('OpenId','string','平台唯一标识'),
				array('Type','int8','平台类型（1 -- 手Q；2 -- 微信）'),
		),
		'out'	=> array(
				array('Id','int','服务器ID'),
				array('Name','string','服务器名称'),
				array('Status','int8','服务器状态[0-维护,1-通畅,2-繁忙,3-拥挤]'),
				array('IsNew','bool','是否是新服'),
				array('IsHot','bool','是否是推荐服'),
				array('OpenTime','int64','开服时间'),
		),
		'out_list' => true,

		'comment'  => '获取游戏服',
	),


	array(
		'name' => "user/server",
		'in'   => array(
				array('OpenId','string','平台唯一标识'),
				array('Type','int','平台类型（1 -- 手Q；2 -- 微信）'),
		),
		'out'	=> array(
				array('Sid','int','服务器ID'),
				array('Nick','string','角色昵称'),
				array('RoleId','int8','角色ID'),
				array('RoleLevel','int16','角色等级'),
				array('LoginTime','int64','登陆时间'),
		),
		'out_list' => true,

		'comment'  => '获取玩家对应服务器角色',
	),


	array(
		'name' => "user/logon",
		'in'   => array(
				array('OpenId','string','平台唯一标识'),
				array('Type','int','平台类型（1 -- 手Q；2 -- 微信）'),
				array('Sid','int','服务器ID'),
		),		
		'out'   => array(
				array('NeedCreate','bool','需要创建角色'),
				array('Nick','string','角色昵称'),
				array('RoleId','int8','角色ID'),
				array('LoginTime','int64','登陆时间'),
				array('Hash','string','hash串（被base64-encoded）'),
				array('IP','string','游戏服IP'),
				array('Port','string','游戏服端口'),
		),
		'comment'  => '登陆游戏',
	),

	array(
		'name' => "user/create",
		'in'   => array(
				array('OpenId','string','平台唯一标识'),
				array('Type','int','平台类型（1 -- 手Q；2 -- 微信）'),
				array('Sid','int','服务器ID'),
				array('RoleId','int8','角色ID'),
				array('Nick','int','角色昵称'),
		),
		'out'   => array(
				array('NickExist','bool','昵称已存在'),
				array('Nick','string','角色昵称'),
				array('RoleId','int8','角色ID'),
				array('LoginTime','int64','登陆时间'),
				array('Hash','string','hash串（被base64-encoded）'),
				array('IP','string','游戏服IP'),
				array('Port','string','游戏服端口'),
		),		
		'comment'  => '创建角色',
	),

	array(
		'name' => "user/update",
		'in'   => array(
				array('OpenId','string','平台唯一标识'),
				array('Type','int','平台类型（1 -- 手Q；2 -- 微信）'),
				array('Sid','int','服务器ID'),
				array('RoleLevel','int16','角色等级'),
		),
		'out'   => array(
		),		
		'comment'  => '更新角色信息',
	),
);

echo "平台接口列表<br/><br/>";
foreach ($APIConfigs as $key => $api) {
	echo "<pre>-------------------------------------------------------------</pre>";
	echo "<ul>";
	echo "<li><b> {$api['name']} - {$api['comment']} </b></li>";
		echo "<br/>IN {<ul>";
			foreach ($api['in'] as $key => $rs) {
				echo "<li>{$rs[0]} ： {$rs[1]} （{$rs[2]}）</li>";
			}
		echo "</ul>}";
		$outList = $api['out_list'] ? "list" : "";
		echo "<br/><br/>OUT<br/>{$outList} {<ul>";
			foreach ($api['out'] as $key => $rs) {
				echo "<li>{$rs[0]} ： {$rs[1]} （{$rs[2]}）</li>";
			}
		echo "</ul>}";
	echo "</ul>";
}
echo "<pre>-------------------------------------------------------------</pre>";

// function 
?>
</body>
</html>