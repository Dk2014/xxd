#基础项目目录结构

```
├── database
├── protocol
├── python_client
├── README.md
├── scripts
├── server
├── start.sh
├── system
├── TODO.md
├── tools
├── xrestart
├── zautobuild.sh
├── zbuild.sh
├── zcmd.sh
├── zcmd.sh.orig
├── zpayment_build.sh
├── zplatform_build.sh
├── zreload.sh
├── zserver.sh
├── zserver.sh.orig
├── zstart.sh
├── ztool.sh
├── ztss_build.sh
└── zupgrade.sh
```

##几个重要的目录
1. **server** go服务器代码
2. **database** 代码生成、客户端数值导出、升级脚本、数据库数值导出等功能
3. **system** 数值后台
4. **tools** 辅助工具，客户端数值以及协议描述文件生成、API文档、素材管理等
5. **protocol** 协议描述解析以及代码生成

### database
```
├── cls_dbupgrader.php
├── config
├── config.php
├── database-game_server-upgrade.php
├── database.php
├── def_config.php
├── gamedb_upgrade.php
├── lang_rawstr.php
├── phpimport.php
├── plugins
├── servers
├── sqlimport.php
├── sql.php
├── updates
├── zcode
├── zip_dump_file.php
└── zupgrade
```

* def_config.php 开发数据库。 本地需要配置 config.php 覆盖掉 $db_config 变量
* database-game_server-upgrade.php 是升级脚本集合
* database.php 数据库操作脚本，有数据库导出 数据库升级 清理玩家数据 运行坚持脚本等功能
* lang_rawstr.php 各种语言之间字符串字面值转移
* zip_dump_file.php 把两个连续的数值导出压缩为1个
* phpimport.php 升级指定数据库
* cls_dbupgrader.php 数据库升级相关类
* gamedb_upgrade.php 生成单个php升级脚本

* confg/ 客户端以及服务端通用常量对应 game_server/dat/里面的一个包
* plugins/ 代码生成相关代码 client_dat 客户端数值；server_code 服务端mdb 代码；server_data 服务端 dat 代码 以及 战斗相关代码
* servers/ 数据库升级脚本

### protocol


## 游戏服项目
游戏服代码主要在 game_server 和 core 里面

```
game_server:
├── api
├── battle
├── config
├── dat
├── global
├── main.go
├── mdb
├── module
├── rpc
├── tencent
├── tlog
├── web
└── xdlog
```

* api 网络协议路由，基本是自动生成
* battle 战斗核心逻辑
* config 配置
* dat 数值相关
* global 全局玩家信息 cache
* main.go 入口
* mdb 内存数据库，包括 事务日志、数据持久化（写MySQL）、写运营日志等，提供玩家级别和全局级别得事务锁
* module 业务逻辑实现得地方
* rpc RPC基础框架，对标准库做了一些封装。值得注意的是，这里的RPC都分装成异步得形式。
* tencent 腾讯相关请求分装：游戏货币余额、战力提交、信鸽消息推送
* web web调试接口（开发中）
* tlog 腾讯运营日志
* xdlog 心动运营日志

```
core:
├── crontab
├── dbgutil
├── debug
├── fail
├── goid
├── gozd
├── i18l
├── log
├── mysql
├── net
├── pinideaosext
├── redis
├── syncutil
├── time
└── tprof
```
