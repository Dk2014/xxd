提供一个接入腾讯tlog的golang版本。

####Status: under developing

#主要代码结构

* src/client/ 主要的 tlog package 代码
* src/parser/ parser工具的源代码
* xml/        可用于放置日志xml规范文件。tlog-parser根据目录下的xml文件生成tlog-ext.go
* bin/        放置parser工具

#使用方法
1. 开发和使用时将 src/tlog_client 做软连接到 $GOPATH/src
2. 对parser修改后可以使用 ./build.sh 脚本编译
3. 可以使用 ./gen.sh 生成tlog客户端代码（常量要另外生成)
4. xml/query.py 脚本可用来查找xml里面的使用到的数据类型



##首先使用parser工具生成tlog-ext.go

bin/tlog-parser (源码位于 src/parser 下)。该工具可以分析 xml 目录下的日志xml规范，生成 src/client 下的 tlog-ext.go 。生成的 tlog-ext.go 提供严格的 tlog struct 和 Packet() 的实现。

如：
    
    // 将Log内容设计成结构的意义是在业务场景可以提前创建，但仅在事务提交时才发送
    // 此结构体由工具按XML描述自动生成
    type PlayerLoginLog struct {
      PlayerId  int64
      ClientIP  string
      LoginTime time.Time
    }

    // 每个结构体都实现Log接口要求的Packet方法
    func (log *PlayerLoginLog) Packet() []byte {
      // 可以加入必要的校验
      struct_name = "PlayerLoginLog"
      buff := make([]byte, 0, 1024 /* TODO: 代码生成器可以根据字段类型计算出最合适buff大小，避免重复申请内存和复制数据 */)
      
      buff = append(buff, []byte(struct_name))
      buff = append(buff, '|')
      buff = strconv.AppendInt(buff, log.PlayerId, 10)
      buff = append(buff, '|')
      // 加入必要的校验和修正
      if (len(log.ClientIP) > 16) {
        buff = append(buff, []byte(log.ClientIP[:16])...)
      } else {
        buff = append(buff, []byte(log.ClientIP)...)
      }
      buff = append(buff, []byte(log.ClientIP)...)
      buff = append(buff, '|')
      buff = strconv.AppendInt(buff, log.LoginTime.Unix(), 10)
      
      return buff
    }
    
    // 严格的 NewStucture functions 防止松散的创建方式带来未初始化的字段
    tlog.NewPlayerLoginLog(PlayerId  int64, ClientIP string, LoginTime time.Time) *PlayerLoginLog 
    
    // *关于“必要的校验和修正”：是指确实存在传入参数超过xml规范约定允许值时才需要进行的校验和修正操作。
    // 例如，如果struct的member如果定义为string格式时，需要先确保其长度不超过tlog规范xml中所定义的长度。
    // 如果是可以在编译时或更早避免的数据问题，则不必再进行校验。以减少CPU开销

parser 同时也会生成 tlog-ext_test.go 对所有 struct 的 Packet() function 进行单元测试。

##初始化tlog

    // server 配置tlog服务器地址与端口
    // path 日志xml规范的存储路径，也可以为空（但不推荐为空）
    tlog.Init(server string, path string)

##发送Log

    // 所有Log数据均可调用Send来发送
    tlog.Send(log tlog.Log)

    
## 常量变化
在两个地方会用到常量：
1. xml/tlog_template.xml
2. 另外一个项目：game_server/tlog/const.go

诸如 XXX原因、是人工归纳，然后写在 `xml/tlog_template.xml` 里面的。
物品累类型、物品ID 这些常量可通过 `xml/gen_item_const.php` 产生，然后更新到`xml/tlog_template.xml`里面。

`game_server/tlog/const.go` 文件可以使用 `xml/go_const.py log_template.xml` 产生，然后更新到 game_server 项目里面。

`go_const.py` 依赖一些第三方的包：

	sudo brew install python
	sudo brew install pip
	sudo pip install jinja2 MySQL-python

