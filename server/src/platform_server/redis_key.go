package platform_server

import (
	"core/fail"
	"core/redis"
	"fmt"
)

/*
  data structure:
    key                                                 data type   value type                        meaning
    gameserverlist_rev_counter                                      int64                             last server list revision
    gameserverlist_key                                              int64                               current server list revision
    gameserverlist_{gameserverlist_key}                 (hash)      map[{Sid}_{Type}_{Gsid}]Server
    account_sid_nick_{sid}_{type}_{nick}                            bool                              exist
    gs_counter_{sid}_{type}_{gs_id}                                 int64                             角色数计数器
    account_openid_type_{openid}_{type}                 (hash)      map[{Sid}]User                    账户角色信息
    client_release_version_{type}                                   int64                             线上发行客户端版本
    client_audit_version_{type}                                     int64                             提交审核版本号
    client_min_version_{type}                                       int64                             允许的最小客户端版本
    client_patch_file_{type}                            (hash)      map[{client}_{server}]patch       客户端需要下载的patch文件路径
    total_resource_file_{version}_{type}                (hash)      map[resourceId]resource           客户端需要下载的资源文件路径
    client_patch_url_key                                string                                        patch下载的url前缀
    client_resource_url_key                             string                                        城镇资源下载的url前缀

    announce_revision_key_{type}                        int64                                           公告key
    announce_type_revision_{type}_{revision}            string                                          announce内容

    client_version_black_table_{type}                   (set)      set[int...]                           客户端版本黑名单
    openid_white_table_key                              (set)      set[int]                             openid白名单
    client_upgrade_url_{type}                           string                                          客户端更新地址
    action_audit_type_{type}                            int                                             活动是否开启
    disable_action_picture_{type}                       (set)      set[string...]                       客户端需要屏蔽的图片名字
    close_type_servers_{type}                           (int)      0/1                                   是否开启type平台的维护功能

    dark_launch_{type}                                  (hash)     {sid1:sid2}_{client_version}_compabile_url    灰度服配置信息

    //wegame keys
    wegame_platform_uid_{source}_{uniqueid}_key 				string                                          第三方账号对应的platform_uid

    ip_black_tables_key									(set)												ip黑名单

*/

func RedisKey_GameServerListRevCounter() string {
	return "gameserverlist_rev_counter"
}

func RedisKey_GameServerListByRevision(rev int) string {
	return fmt.Sprintf("gameserverlist_%v", rev)
}

func RedisKey_GameServerListCurrentRev() string {
	return "gameserverlist_key"
}

func RedisKey_GameServerList(c redis.Conn) string {
	// fetch from redis by key "gameserverlist"
	v, err := redis.Int(c.Do("GET", RedisKey_GameServerListCurrentRev()))
	fail.When(err != nil, "error on getting current server list revision")

	return RedisKey_GameServerListByRevision(v)
}

func RedisKey_ExistenceBySidTypeNick(sid int32, itype uint8, nick string) string {
	return fmt.Sprintf("account_sid_nick_%v_%v_%v", sid, itype, nick)
}

func RedisKey_CounterBySidGsid(sid int32, itype uint8, gsid int32) string {
	return fmt.Sprintf("gs_counter_%v_%v_%v", sid, itype, gsid)
}

func RedisKey_RoleListByOpenidType(openid string, itype uint8) string {
	return fmt.Sprintf("account_openid_type_%v_%v", openid, itype)
}

func RedisKey_ClientReleaseVersionByType(itype uint8) string {
	return fmt.Sprintf("client_release_version_%v", itype)
}

func RedisKey_ClientAuditVersionByType(itype uint8) string {
	return fmt.Sprintf("client_audit_version_%v", itype)
}

func RedisKey_ClientMinVersionByType(itype uint8) string {
	return fmt.Sprintf("client_min_version_%v", itype)
}

func RedisKey_ClientPatchHash(itype uint8) string {
	return fmt.Sprintf("client_patch_file_%v", itype)
}

func RedisKey_TotalResourceHash(version int32) string {
	return fmt.Sprintf("total_resource_file_%v", version)
}

func RedisKey_ClientPatchUrl() string {
	return "client_patch_url_key"
}

func RedisKey_ClientResourceUrl() string {
	return "client_resource_url_key"
}

func RedisKey_AnnounceCurrentRevByType(itype uint8) string {
	return fmt.Sprintf("announce_revision_key_%v", itype)
}

func RedisKey_AnnounceByTypeRevision(itype uint8, rev int32) string {
	return fmt.Sprintf("announce_type_revision_%v_%v", itype, rev)
}

func RedisKey_ClientVersionBlackTable(itype uint8) string {
	return fmt.Sprintf("client_version_black_table_%v", itype)
}

func RedisKey_OpenidWhiteTable() string {
	return "openid_white_table_key"
}

func RedisKey_ClientUpgradeUrl(itype uint8) string {
	return fmt.Sprintf("client_upgrade_url_%v", itype)
}

func RedisKey_ActionAudit(itype uint8) string {
	return fmt.Sprintf("action_audit_type_%v", itype)
}

func RedisKey_DisableActionPic(itype uint8) string {
	return fmt.Sprintf("disable_action_picture_%v", itype)
}

func RedisKey_CloseTypeServers(itype uint8) string {
	return fmt.Sprintf("close_type_servers_%v", itype)
}

func RedisKey_DarkLaunch(itype uint8) string {
	return fmt.Sprintf("dark_launch_%v", itype)
}

func RedisKey_WG_THIRD_BIND_KEY(source string, unique_id string) string {
	return fmt.Sprintf("wegame_platform_uid_%s_%s_key", source, unique_id)
}

func RedisKey_IpBlackTable() string {
	return "ip_black_tables_key"
}
