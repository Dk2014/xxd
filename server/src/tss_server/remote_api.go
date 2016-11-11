package main

import (
	"tss_server/tss_sdk_go"
	"core/log"
	"core/debug"
)

//rpc user_info
type UserInfo struct {
	OpenId     string
	PlayerId   int64
	Plat_Id    int
	Client_Ip  string
	Client_Ver int
	Role_Name  string
}

//rpc pkg_info
type PkgInfo struct {
	OpenId    string
	PlayerId  int64
	Plat_Id   int
	Anti_Data []byte
}

//rpc operation
type RpcOper int

//user login
func (this *RpcOper) UserLogin(user_info *UserInfo, reply *int) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`tss_server.remote_api.UserLogin
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		g_syncChan <- tss_sdk_go.SyncUnit{Act: tss_sdk_go.TSS_ACTION_USERLOGIN, Data: user_info}
		g_players.Set(user_info.OpenId, user_info.PlayerId)
	}()
	return nil
}

//user logout
func (this *RpcOper) UserLogout(user_info *UserInfo, reply *int) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`tss_server.remote_api.UserLogout
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		g_syncChan <- tss_sdk_go.SyncUnit{Act: tss_sdk_go.TSS_ACTION_USERLOGOUT, Data: user_info}
	}()
	return nil
}

//receive data from client
func (this *RpcOper) RecvData(pkg_info *PkgInfo, reply *int) error {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf(`tss_server.remote_api.RecvData
Error = %v
Stack = 
%s`,
					err,
					debug.Stack(1, "    "),
				)
			}
		}()
		g_syncChan <- tss_sdk_go.SyncUnit{Act: tss_sdk_go.TSS_ACTION_RECV, Data: pkg_info}
	}()
	return nil
}
