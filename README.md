# README #

仙侠道手机版本服务端

### 项目结构 ###

* /database     | 数据库运维脚本
* /protocol     | 通讯协议
* /server       | 主要代码
* /system       |
* /tools        |
* /scripts       | 辅助脚本
* hd.conf       | 互动服务配置
* xxd.conf      | 游戏服配置
* zautobuild.sh   | 自动更新代码编译脚本
* zbuild.sh       | 编译脚本
* zguard.sh       | 服务监控自启动脚本
* zplatform_build.sh | 编译登陆平台
* zserver.sh   | 服务管理脚本[start|stop|restart]
* ztool.sh     | 剧情编辑器数据生成脚本
* zupgrade_clientdb.sh | 更新客户端服务程序脚本

### 常用命令 ###

* ./zbuild.sh
  编译 game_server （动态链接）

* ./zbuild.sh static
  编译 game_server （静态链接 static linked）

* ./zplatform_build.sh
  编译 platform_server （动态链接）

* ./zplatform_build.sh static
  编译 platform_server （静态链接 static linked）

* ./database/zupgrade 
  升级游戏数据库和平台数据库

* php ./database/sql.php game_server 2014010101 2015010101 ~/upgrade_2014050401_to_2015010101.sql
  导出数据库增量更新的SQL文件
  
* scripts/xxds_setup.sh
  可以在基础环境中（linux amd64，仅需要 mysql 客户端，不需要golang或其他library）中快速安装部署服务端
  使用方法：./xxds_setup.sh <version|latest> <gs.conf> <hd.conf> <process_groupname> <platform_port> <db_host> <db_usrname> <db_passwd> <db_port> <dbname_gs> <dbname_platform> 
    <version|latest> 使用latest自动安装最新版本，使用数据版本号则会安装特定版本
    <gs.conf>        gs服务端配置文件，多个gs服务端可以使用逗号“,”分隔
    <hd.conf>        互动服务端配置文件
    <process_groupname>  服务器组名，务必使用与系统内其他进程不会重复的名称。例如 _XXDS_SERVER_!234_ 。例如批量关闭时会被用来识别服务进程。
    <platform_port>  平台服务器监听的端口
  例如安装最新版本单gs服务组并启动：
    ./xxds_setup.sh latest gs.conf hd.conf _xxds_server_000_ 8888 rdsyejaq3yejaq3.mysql.rds.aliyuncs.com xxd000 xxd000KJS7 3306 xxd000 xxd000_platform
  例如安装最新版本多gs服务组并启动：
    ./xxds_setup.sh latest gs1.conf,gs2.conf hd.conf _xxds_server_000_ 8888 rdsyejaq3yejaq3.mysql.rds.aliyuncs.com xxd000 xxd000KJS7 3306 xxd000 xxd000_platform
  例如安装特定版本(1000)服务组并启动：
    ./xxds_setup.sh 1000 gs.conf hd.conf _xxds_server_000_ 8888 rdsyejaq3yejaq3.mysql.rds.aliyuncs.com xxd000 xxd000KJS7 3306 xxd000 xxd000_platform