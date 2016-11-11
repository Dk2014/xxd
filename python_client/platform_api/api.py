# -*- coding: utf8 -*-
__author__ = 'supei'

import requests
import json


class PlatformAPI:
    _token = '_SANDBOX_'

    PLATFORM_TYPE_IWX = 1
    PLATFORM_TYPE_IQQ = 2
    PLATFORM_TYPE_AWX = 17
    PLATFORM_TYPE_AQQ = 18

    API_SERVER_LIST = "/server/list"
    API_CREATE_USER = "/user/create"
    API_USER_GS_INFO = "/user/gsinfo"

    def __init__(self, host, port, version=99999):
        self._host = host
        self._port = str(port)
        self._version = version

    def _url(self, api_path):
        return "http://" + self._host + ":" + self._port + api_path

    def server_list(self, open_id, _type=PLATFORM_TYPE_IWX):
        '''
        :param open_id: 用户登录名
        :param _type: 平台类型[1-ios微信,2-ios手Q,17-安卓微信,18-安卓手Q]
        :param version: 客户端版本
        :return:
         out {
        List: [
          {
            Id        int64  // 账号ID
            Type      uint8  //
            Openid    string // 玩家唯一标识
            Sid       int32  // 游戏服ID
            Nick      string // 玩家昵称
            RoleId    int8   // 角色模板ID
            RoleLevel int16  // 角色等级
            LoginTime int64  // 登陆时间
            Gsid      int32
          }...
        ]
      出错：
        {"msg":"访问出错：xxx", "error":400}
      }
        '''
        data = dict(
            OpenId=open_id,
            Type=_type,
            AreaId=0,
            PlatId=0,
            Version=self._version,
        )
        url = self._url(self.API_SERVER_LIST)
        rsp = requests.post(url, json=data)
        return rsp.json()

    def create_user(self, open_id, sid, role_id, nick, _type=PLATFORM_TYPE_IWX):
        """
        :param open_id:
        :param sid:
        :param role_id:
        :param nick:
        :param _type:
        :return:
           out {
        NickExist bool
        Nick      string
        RoleId    int8
        LoginTime int64
        Hash      []byte
        IP        string
        Port      string
        Gsid      int32
      }
      出错：
      {"TokenUnavaiable": 1, "error": 405}
      {"NickExist": 1, "error": 304}
      {"NickTooShort": 1, "error": 1325}
      {"NickTooLong": 1, "error": 413}
      {"NickIllegal": 1, "error": 2925}
      {"msg":"访问出错：xxx", "error":400}
        """
        data = dict(
            OpenId=open_id,
            Type=_type,
            Sid=sid,
            RoleId=role_id,
            Nick=nick,
            Version=self._version,
        )

        url = self._url(self.API_CREATE_USER)
        rsp = requests.post(url, json=data)
        return rsp.json()

    def game_server_info(self, open_id, sid, _type=PLATFORM_TYPE_IWX):
        data = dict(
            OpenId=open_id,
            Type=_type,
            AreaId=0,
            PlatId=0,
            Version=self._version,
        )
        url = self._url(self.API_USER_GS_INFO)
        rsp = requests.post(url, json=data)
        return rsp.json()


if __name__ == '__main__':
    pf = PlatformAPI('platform.xxd.io', 8888)
    data = pf.server_list("wyb1")
    if len(data['List']) > 0:
        for server in data['List']:
            print json.dumps(server, indent=4, ensure_ascii=False).encode('utf8')