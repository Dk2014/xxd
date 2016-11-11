#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

LOGIN_STATUS_FAILED = 0
LOGIN_STATUS_SUCCEED = 1
LOGIN_STATUS_FIRST_TIME = 2
LOGIN_STATUS_WAIT_CLOSE = 3
LOGIN_STATUS_RESTORE_SUCCEED = 4
LOGIN_STATUS_BAN = 5
LOGIN_STATUS_RELOGIN = 6
LOGIN_STATUS_INVALID_PAYTOKEN = 7


PLATFORM_TYPE_MOBILE_IOS_WEIXIN = 1
PLATFORM_TYPE_MOBILE_IOS_QQ = 2
PLATFORM_TYPE_MOBILE_IOS_GUEST = 5
PLATFORM_TYPE_MOBILE_ANDROID_WEIXIN = 17
PLATFORM_TYPE_MOBILE_ANDROID_QQ = 18
PLATFORM_TYPE_MOBILE_ANDROID_GUEST = 21


CHANNEL_ID_MOBILE = 0
CHANNEL_ID_MOBILE_IOS_VN = 1
CHANNEL_ID_MOBILE_ANDROID_VN = 2
CHANNEL_ID_MOBILE_IOS_TW = 3
CHANNEL_ID_MOBILE_ANDROID_TW = 4


RANK_TYPE_FIGHTNUM = 0
RANK_TYPE_LEVEL = 1


class InfoUp(object):
    _module = 0
    _action = 3

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class InfoDown(object):
    _module = 0
    _action = 3

    def __init__(self):
        pass
        self.nickname = ''
        self.town_id = 0
        self.role_id = 0
        self.role_lv = 0
        self.role_exp = 0
        self.vip_level = 0
        self.fame = 0
        self.ingot = 0
        self.coins = 0
        self.fragments = 0
        self.heart_num = 0
        self.physical = 0
        self.physical_buy_count = 0
        self.coins_buy_count = 0
        self.batch_bought_coins = 0
        self.cornucopia_count = 0
        self.func_key = 0
        self.played_key = 0
        self.town_lock = 0
        self.mission_key = 0
        self.mission_max_order = 0
        self.mission_level_max_lock = 0
        self.mission_level_award_lock = 0
        self.hard_level_lock = 0
        self.quest_id = 0
        self.quest_state = 0
        self.arena_report_num = 0
        self.next_free_star_box_timestamp = 0
        self.next_free_moon_box_timestamp = 0
        self.next_free_sun_box_timestamp = 0
        self.next_free_hunyuan_box_timestamp = 0
        self.qq_vip_status = 0
        self.server_open_time = 0
        self.month_card_expire_timestamp = 0
        self.never_recharge = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<h", self.town_id))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.role_lv))
        buff.extend(struct.pack("<q", self.role_exp))
        buff.extend(struct.pack("<h", self.vip_level))
        buff.extend(struct.pack("<l", self.fame))
        buff.extend(struct.pack("<q", self.ingot))
        buff.extend(struct.pack("<q", self.coins))
        buff.extend(struct.pack("<q", self.fragments))
        buff.extend(struct.pack("<h", self.heart_num))
        buff.extend(struct.pack("<h", self.physical))
        buff.extend(struct.pack("<h", self.physical_buy_count))
        buff.extend(struct.pack("<h", self.coins_buy_count))
        buff.extend(struct.pack("<h", self.batch_bought_coins))
        buff.extend(struct.pack("<h", self.cornucopia_count))
        buff.extend(struct.pack("<h", self.func_key))
        buff.extend(struct.pack("<h", self.played_key))
        buff.extend(struct.pack("<l", self.town_lock))
        buff.extend(struct.pack("<l", self.mission_key))
        buff.extend(struct.pack("<b", self.mission_max_order))
        buff.extend(struct.pack("<l", self.mission_level_max_lock))
        buff.extend(struct.pack("<l", self.mission_level_award_lock))
        buff.extend(struct.pack("<l", self.hard_level_lock))
        buff.extend(struct.pack("<h", self.quest_id))
        buff.extend(struct.pack("<b", self.quest_state))
        buff.extend(struct.pack("<h", self.arena_report_num))
        buff.extend(struct.pack("<q", self.next_free_star_box_timestamp))
        buff.extend(struct.pack("<q", self.next_free_moon_box_timestamp))
        buff.extend(struct.pack("<q", self.next_free_sun_box_timestamp))
        buff.extend(struct.pack("<q", self.next_free_hunyuan_box_timestamp))
        buff.extend(struct.pack("<h", self.qq_vip_status))
        buff.extend(struct.pack("<q", self.server_open_time))
        buff.extend(struct.pack("<q", self.month_card_expire_timestamp))
        buff.extend(struct.pack("<?", self.never_recharge))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.town_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.role_lv = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.role_exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.vip_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fame = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.ingot = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.fragments = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.heart_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.physical = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.physical_buy_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.coins_buy_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.batch_bought_coins = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.cornucopia_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.func_key = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.played_key = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.town_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.mission_key = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.mission_max_order = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.mission_level_max_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.mission_level_award_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.hard_level_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.quest_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.quest_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.arena_report_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.next_free_star_box_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.next_free_moon_box_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.next_free_sun_box_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.next_free_hunyuan_box_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.qq_vip_status = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.server_open_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.month_card_expire_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.never_recharge = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 138
        size += len(self.nickname)
        return size


class ReloginUp(object):
    _module = 0
    _action = 4

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class ReloginDown(object):
    _module = 0
    _action = 4

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 0
        

class BuyPhysicalUp(object):
    _module = 0
    _action = 5

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class BuyPhysicalDown(object):
    _module = 0
    _action = 5

    def __init__(self):
        pass
        self.physical_buy_count = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.physical_buy_count))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.physical_buy_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class SetPlayKeyUp(object):
    _module = 0
    _action = 6

    def __init__(self):
        pass
        self.lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.lock = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class SetPlayKeyDown(object):
    _module = 0
    _action = 6

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 0
        

class GetTimeUp(object):
    _module = 0
    _action = 7

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class GetTimeDown(object):
    _module = 0
    _action = 7

    def __init__(self):
        pass
        self.unix_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.unix_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.unix_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class FromPlatformLoginUp(object):
    _module = 0
    _action = 8

    def __init__(self):
        pass
        self.platform_type = 0
        self.channel_id = 0
        self.username = ''
        self.nickname = ''
        self.role_id = 0
        self.hashcode = ''
        self.unixtime = 0
        self.restore = False
        self.recvCount = 0
        self.gsid = 0
        self.openkey = ''
        self.pay_token = ''
        self.pfkey = ''
        self.zoneid = 0
        self.pf = ''
        self.channel = 0
        self.telecom_oper = ''
        self.client_version = ''
        self.system_hardware = ''
        self.network = ''
        self.wegames_platform_uid = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.platform_type))
        buff.extend(struct.pack("<B", self.channel_id))
        buff.extend(struct.pack('<H', len(self.username)))
        buff.extend(self.username)
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack('<H', len(self.hashcode)))
        buff.extend(self.hashcode)
        buff.extend(struct.pack("<q", self.unixtime))
        buff.extend(struct.pack("<?", self.restore))
        buff.extend(struct.pack("<q", self.recvCount))
        buff.extend(struct.pack("<l", self.gsid))
        buff.extend(struct.pack('<H', len(self.openkey)))
        buff.extend(self.openkey)
        buff.extend(struct.pack('<H', len(self.pay_token)))
        buff.extend(self.pay_token)
        buff.extend(struct.pack('<H', len(self.pfkey)))
        buff.extend(self.pfkey)
        buff.extend(struct.pack("<l", self.zoneid))
        buff.extend(struct.pack('<H', len(self.pf)))
        buff.extend(self.pf)
        buff.extend(struct.pack("<l", self.channel))
        buff.extend(struct.pack('<H', len(self.telecom_oper)))
        buff.extend(self.telecom_oper)
        buff.extend(struct.pack('<H', len(self.client_version)))
        buff.extend(self.client_version)
        buff.extend(struct.pack('<H', len(self.system_hardware)))
        buff.extend(self.system_hardware)
        buff.extend(struct.pack('<H', len(self.network)))
        buff.extend(self.network)
        buff.extend(struct.pack('<H', len(self.wegames_platform_uid)))
        buff.extend(self.wegames_platform_uid)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.platform_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.channel_id = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _username_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.username = str(raw_msg[idx:idx+_username_size])
        idx += _username_size

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _hashcode_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hashcode = str(raw_msg[idx:idx+_hashcode_size])
        idx += _hashcode_size

        self.unixtime = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.restore = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.recvCount = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.gsid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _openkey_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openkey = str(raw_msg[idx:idx+_openkey_size])
        idx += _openkey_size

        _pay_token_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pay_token = str(raw_msg[idx:idx+_pay_token_size])
        idx += _pay_token_size

        _pfkey_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pfkey = str(raw_msg[idx:idx+_pfkey_size])
        idx += _pfkey_size

        self.zoneid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _pf_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pf = str(raw_msg[idx:idx+_pf_size])
        idx += _pf_size

        self.channel = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _telecom_oper_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.telecom_oper = str(raw_msg[idx:idx+_telecom_oper_size])
        idx += _telecom_oper_size

        _client_version_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.client_version = str(raw_msg[idx:idx+_client_version_size])
        idx += _client_version_size

        _system_hardware_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.system_hardware = str(raw_msg[idx:idx+_system_hardware_size])
        idx += _system_hardware_size

        _network_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.network = str(raw_msg[idx:idx+_network_size])
        idx += _network_size

        _wegames_platform_uid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.wegames_platform_uid = str(raw_msg[idx:idx+_wegames_platform_uid_size])
        idx += _wegames_platform_uid_size

    def size(self):
        size = 58
        size += len(self.username)
        size += len(self.nickname)
        size += len(self.hashcode)
        size += len(self.openkey)
        size += len(self.pay_token)
        size += len(self.pfkey)
        size += len(self.pf)
        size += len(self.telecom_oper)
        size += len(self.client_version)
        size += len(self.system_hardware)
        size += len(self.network)
        size += len(self.wegames_platform_uid)
        return size


class FromPlatformLoginDown(object):
    _module = 0
    _action = 8

    def __init__(self):
        pass
        self.status = 0
        self.player_id = 0
        self.ban_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<q", self.ban_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.ban_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 17
        

class BuyCoinsUp(object):
    _module = 0
    _action = 9

    def __init__(self):
        pass
        self.number = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.number))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.number = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class BuyCoinsDown(object):
    _module = 0
    _action = 9

    def __init__(self):
        pass
        self.result = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.result)))
        for item in self.result:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _result_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_result_size):
            obj = BuyCoinsDownResult()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.result.append(obj)

    def size(self):
        size = 1
        for item in self.result:
            size += item.size()
        return size


class BuyCoinsDownResult(object):
    def __init__(self):
        pass
        self.ingot = 0
        self.coins = 0
        self.crit = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.ingot))
        buff.extend(struct.pack("<q", self.coins))
        buff.extend(struct.pack("<?", self.crit))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ingot = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.crit = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 17
        

class GetLoginInfoUp(object):
    _module = 0
    _action = 10

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class GetLoginInfoDown(object):
    _module = 0
    _action = 10

    def __init__(self):
        pass
        self.hash = ''
        self.time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.hash)))
        buff.extend(self.hash)
        buff.extend(struct.pack("<q", self.time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _hash_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hash = str(raw_msg[idx:idx+_hash_size])
        idx += _hash_size

        self.time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 10
        size += len(self.hash)
        return size


class CrossLoginGameServerUp(object):
    _module = 0
    _action = 20

    def __init__(self):
        pass
        self.pid = 0
        self.openid = ''
        self.nick = ''
        self.role_id = 0
        self.role_level = 0
        self.time = 0
        self.hash = ''
        self.wegames_platform_uid = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.role_level))
        buff.extend(struct.pack("<q", self.time))
        buff.extend(struct.pack('<H', len(self.hash)))
        buff.extend(self.hash)
        buff.extend(struct.pack('<H', len(self.wegames_platform_uid)))
        buff.extend(self.wegames_platform_uid)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.role_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _hash_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hash = str(raw_msg[idx:idx+_hash_size])
        idx += _hash_size

        _wegames_platform_uid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.wegames_platform_uid = str(raw_msg[idx:idx+_wegames_platform_uid_size])
        idx += _wegames_platform_uid_size

    def size(self):
        size = 29
        size += len(self.openid)
        size += len(self.nick)
        size += len(self.hash)
        size += len(self.wegames_platform_uid)
        return size


class CrossLoginGameServerDown(object):
    _module = 0
    _action = 20

    def __init__(self):
        pass
        self.result = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyGlobalServerInfoDown(object):
    _module = 0
    _action = 25

    def __init__(self):
        pass
        self.gsid = 0
        self.addr = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.gsid))
        buff.extend(struct.pack('<H', len(self.addr)))
        buff.extend(self.addr)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.gsid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _addr_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.addr = str(raw_msg[idx:idx+_addr_size])
        idx += _addr_size

    def size(self):
        size = 6
        size += len(self.addr)
        return size


class GlobalLoginUp(object):
    _module = 0
    _action = 26

    def __init__(self):
        pass
        self.pid = 0
        self.openid = ''
        self.nick = ''
        self.role_id = 0
        self.role_level = 0
        self.fight_num = 0
        self.time = 0
        self.hash = ''
        self.gsid = 0
        self.platform_type = 0
        self.channel_id = 0
        self.username = ''
        self.openkey = ''
        self.pay_token = ''
        self.pfkey = ''
        self.zoneid = 0
        self.pf = ''
        self.wegames_platform_uid = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.role_level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<q", self.time))
        buff.extend(struct.pack('<H', len(self.hash)))
        buff.extend(self.hash)
        buff.extend(struct.pack("<l", self.gsid))
        buff.extend(struct.pack("<B", self.platform_type))
        buff.extend(struct.pack("<B", self.channel_id))
        buff.extend(struct.pack('<H', len(self.username)))
        buff.extend(self.username)
        buff.extend(struct.pack('<H', len(self.openkey)))
        buff.extend(self.openkey)
        buff.extend(struct.pack('<H', len(self.pay_token)))
        buff.extend(self.pay_token)
        buff.extend(struct.pack('<H', len(self.pfkey)))
        buff.extend(self.pfkey)
        buff.extend(struct.pack("<l", self.zoneid))
        buff.extend(struct.pack('<H', len(self.pf)))
        buff.extend(self.pf)
        buff.extend(struct.pack('<H', len(self.wegames_platform_uid)))
        buff.extend(self.wegames_platform_uid)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.role_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _hash_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hash = str(raw_msg[idx:idx+_hash_size])
        idx += _hash_size

        self.gsid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.platform_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.channel_id = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _username_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.username = str(raw_msg[idx:idx+_username_size])
        idx += _username_size

        _openkey_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openkey = str(raw_msg[idx:idx+_openkey_size])
        idx += _openkey_size

        _pay_token_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pay_token = str(raw_msg[idx:idx+_pay_token_size])
        idx += _pay_token_size

        _pfkey_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pfkey = str(raw_msg[idx:idx+_pfkey_size])
        idx += _pfkey_size

        self.zoneid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _pf_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pf = str(raw_msg[idx:idx+_pf_size])
        idx += _pf_size

        _wegames_platform_uid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.wegames_platform_uid = str(raw_msg[idx:idx+_wegames_platform_uid_size])
        idx += _wegames_platform_uid_size

    def size(self):
        size = 53
        size += len(self.openid)
        size += len(self.nick)
        size += len(self.hash)
        size += len(self.username)
        size += len(self.openkey)
        size += len(self.pay_token)
        size += len(self.pfkey)
        size += len(self.pf)
        size += len(self.wegames_platform_uid)
        return size


class GlobalLoginDown(object):
    _module = 0
    _action = 26

    def __init__(self):
        pass
        self.paytoken = False
        self.result = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.paytoken))
        buff.extend(struct.pack("<?", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.paytoken = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.result = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 2
        

class GetIngotUp(object):
    _module = 0
    _action = 27

    def __init__(self):
        pass

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        return buff

    def decode(self, raw_msg):
        pass

    @staticmethod
    def size():
        return 2
        

class GetIngotDown(object):
    _module = 0
    _action = 27

    def __init__(self):
        pass
        self.ingot = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.ingot))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ingot = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class SystemFameUp(object):
    _module = 0
    _action = 28

    def __init__(self):
        pass
        self.system = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.system))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.system = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class SystemFameDown(object):
    _module = 0
    _action = 28

    def __init__(self):
        pass
        self.fame = 0
        self.system = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.fame))
        buff.extend(struct.pack("<h", self.system))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fame = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.system = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class GetRanksUp(object):
    _module = 0
    _action = 29

    def __init__(self):
        pass
        self.flag = 0
        self.page_index = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.flag))
        buff.extend(struct.pack("<q", self.page_index))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.flag = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.page_index = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 11
        

class GetRanksDown(object):
    _module = 0
    _action = 29

    def __init__(self):
        pass
        self.ranks = []
        self.has_next = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.ranks)))
        for item in self.ranks:
            buff.extend(item.encode())
        buff.extend(struct.pack("<?", self.has_next))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _ranks_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_ranks_size):
            obj = GetRanksDownRanks()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.ranks.append(obj)

        self.has_next = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 2
        for item in self.ranks:
            size += item.size()
        return size


class GetRanksDownRanks(object):
    def __init__(self):
        pass
        self.pid = 0
        self.nickname = ''
        self.rank = 0
        self.values = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<q", self.rank))
        buff.extend(struct.pack('<B', len(self.values)))
        for item in self.values:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.rank = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _values_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_values_size):
            obj = GetRanksDownRanksValues()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.values.append(obj)

    def size(self):
        size = 19
        size += len(self.nickname)
        for item in self.values:
            size += item.size()
        return size


class GetRanksDownRanksValues(object):
    def __init__(self):
        pass
        self.flag = 0
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.flag))
        buff.extend(struct.pack("<q", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.flag = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.value = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 9
        

class PlayerModule(interface.BaseModule):
    decoder_map = {
        3: InfoDown, 
        4: ReloginDown, 
        5: BuyPhysicalDown, 
        6: SetPlayKeyDown, 
        7: GetTimeDown, 
        8: FromPlatformLoginDown, 
        9: BuyCoinsDown, 
        10: GetLoginInfoDown, 
        20: CrossLoginGameServerDown, 
        25: NotifyGlobalServerInfoDown, 
        26: GlobalLoginDown, 
        27: GetIngotDown, 
        28: SystemFameDown, 
        29: GetRanksDown, 
    }
    receive_callback = {}

    def decode(self, message):
        action = ord(message[0])
        decoder_maker = self.decoder_map[action]
        msg = decoder_maker()
        msg.decode(message[1:])
        return msg

    def add_callback(self, action, callback):
        if self.receive_callback.has_key(action):
            self.receive_callback[action].append(callback)
        else:
            self.receive_callback[action] = [callback,]

    def clear_callback(self):
        self.receive_callback = {}

    def add_info(self, callback):
        self.add_callback(3, callback)

    def add_relogin(self, callback):
        self.add_callback(4, callback)

    def add_buy_physical(self, callback):
        self.add_callback(5, callback)

    def add_set_play_key(self, callback):
        self.add_callback(6, callback)

    def add_get_time(self, callback):
        self.add_callback(7, callback)

    def add_from_platform_login(self, callback):
        self.add_callback(8, callback)

    def add_buy_coins(self, callback):
        self.add_callback(9, callback)

    def add_get_login_info(self, callback):
        self.add_callback(10, callback)

    def add_cross_login_game_server(self, callback):
        self.add_callback(20, callback)

    def add_notify_global_server_info(self, callback):
        self.add_callback(25, callback)

    def add_global_login(self, callback):
        self.add_callback(26, callback)

    def add_get_ingot(self, callback):
        self.add_callback(27, callback)

    def add_system_fame(self, callback):
        self.add_callback(28, callback)

    def add_get_ranks(self, callback):
        self.add_callback(29, callback)
