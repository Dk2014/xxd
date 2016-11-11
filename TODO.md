
# platform_server

* 单元测试
* 使用redis作为数据引擎(redis+twemproxy)
* 要求极高的负载承压能力和抗攻击能力
* 无状态设计，支持同时部署多台，水平扩展
* 每个逻辑区服的配置，可以由多个物理服务器(game_server)组成
* platform_server 在 user/logon 接口中对一个openid第一次访问特定逻辑区服时，仅输出该逻辑区服下用户数（或活跃用户数）最少的 game_server 的 IP端口信息。对于不是第一次访问该逻辑区服的OpenID，必须保证每个特定逻辑区服仅返回之前已经对该OpenID提供过的 game_server 的 IP端口，永远不发生变化。换句话说，一个用户（OpenID）连入一个逻辑区服时，永远都会被分配到一个固定的 game_server。
* 从conf文件中读取配置(json格式)，例如redis服务端地址
* 一个命令行工具，可以快速执行：列出服务器配置情况，添加逻辑区服（包括game_server配置），删除逻辑区服，修改逻辑区服下的game_server配置。
* 数据结构（NoSQL）支持免维护升级
* 实现下列API，并易于扩展

    // 以下接口基于HTTP协议，使用POST方法。上传的post body必须是json encoded的字符串
    // 使用utf8编码；返回数据为JSON格式，utf8编码。接口的url是“/模块功能名”，
    // 比如：/serverall 或 user/logon
    // in 代表客户端需要POST上行到服务端的数据
    // out 代表服务端会Response下行到客户端的参数

    serverall
    {
      in {
        OpenId: string
        Type  : int8
      }
      out {
        Id:       Int32
        Name:     Str
        Status:   Int8
        IsNew:    Bool
        IsHot:    Bool
        OpenTime: Int64
      }
    }
    user/server
    {
      in {
      	OpenId string
      	Type   int8
      }
      out {
        array {
        	Id        int64  // 账号ID
        	Type      int8   // 平台类型[1-手Q,2-微信]
        	Openid    string // 玩家唯一标识
        	Sid       int32  // 游戏服ID
        	Nick      string // 玩家昵称
        	RoleId    int8   // 角色模板ID
        	RoleLevel int16  // 角色等级
        	LoginTime int64  // 登陆时间
        }...
      }
    }
    user/logon
    {
      in {
      	OpenId string
      	Type   int8
      	Sid    int32
      }
      out {
      	NeedCreate bool
      	Nick       string
      	RoleId     int8
      	LoginTime  int64
      	Hash       []byte
      	IP         string
      	Port       string
      }
    }
    user/create
    {
      in {
      	OpenId string
      	Type   int8
      	Sid    int32
      	RoleId int8
      	Nick   string
      }
      out {
      	NickExist bool
      	Nick      string
      	RoleId    int8
      	LoginTime int64
      	Hash      []byte
      	IP        string
      	Port      string
      }
    }
    user/update
    {
      in {
        OpenId   : string
        Sid      : int32
        Type     : int8
        RoleLevel: int16
      }
      out {
        
      }
    }

