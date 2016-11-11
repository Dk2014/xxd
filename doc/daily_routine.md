# 日常任务

## 数值导出
一般在结版的时候，需要导出一份数值。或者在开发中，需要将一些数值部署到其他机器时。
```
ssh staff #登录到 staff 机器
cd /Users/staff/Project/xxd-server/database
php database.php export local_game
hg ci -A ./servers/game_server/<PHP_SCRIPT_NAME>
hg push
```

## 数值相关代码生成
在 tools 下面有个网页能根据数据库表结构生成 game_server/dat/XXX 代码。实际中可能需要做一些修改以满足需求。URL 如下：
*http://127.0.0.1:7000/tools/tpl_dat.php?p=TABLE_NAME*

## 网络协议接口实现代码生成
*http://127.0.0.1:7000/tools/protocol.php?p=debug*

## 导出客户端所需得数值和协议文件
```
cd tools/api_viewer
sh build.sh
```
