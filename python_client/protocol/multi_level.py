#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

ROOM_STATUS_SUCCESS = 0
ROOM_STATUS_FAILED_FULL = 1
ROOM_STATUS_FAILED_NOT_EXIST = 2
ROOM_STATUS_FAILED_FIGHTING = 3
ROOM_STATUS_FAILED_REQUIRE_LEVEL = 4
ROOM_STATUS_FAILED_REQUIRE_LOCK = 5


class PartnerInfo(object):
    def __init__(self):
        pass
        self.pid = 0
        self.nick = ''
        self.role_id = 0
        self.fashion_id = 0
        self.level = 0
        self.buddy_role_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<b", self.buddy_role_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.buddy_role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 16
        size += len(self.nick)
        return size


class CreateRoomUp(object):
    _module = 16
    _action = 1

    def __init__(self):
        pass
        self.level_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.level_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CreateRoomDown(object):
    _module = 16
    _action = 1

    def __init__(self):
        pass
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class EnterRoomUp(object):
    _module = 16
    _action = 2

    def __init__(self):
        pass
        self.room_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.room_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.room_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class EnterRoomDown(object):
    _module = 16
    _action = 2

    def __init__(self):
        pass
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class AutoEnterRoomUp(object):
    _module = 16
    _action = 3

    def __init__(self):
        pass
        self.level_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.level_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class AutoEnterRoomDown(object):
    _module = 16
    _action = 3

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyRoomInfoDown(object):
    _module = 16
    _action = 4

    def __init__(self):
        pass
        self.room_id = 0
        self.leader_pid = 0
        self.level_id = 0
        self.partners = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.room_id))
        buff.extend(struct.pack("<q", self.leader_pid))
        buff.extend(struct.pack("<l", self.level_id))
        buff.extend(struct.pack('<B', len(self.partners)))
        for item in self.partners:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.room_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.leader_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.level_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _partners_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_partners_size):
            obj = NotifyRoomInfoDownPartners()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.partners.append(obj)

    def size(self):
        size = 21
        for item in self.partners:
            size += item.size()
        return size


class NotifyRoomInfoDownPartners(object):
    def __init__(self):
        pass
        self.partner = PartnerInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.partner.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.partner.decode(raw_msg[idx:])
        idx += self.partner.size()

    def size(self):
        size = 0
        size += self.partner.size()
        return size


class LeaveRoomUp(object):
    _module = 16
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
        

class LeaveRoomDown(object):
    _module = 16
    _action = 5

    def __init__(self):
        pass
        self.pid = 0
        self.leader_pid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<q", self.leader_pid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.leader_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 16
        

class NotifyJoinPartnerDown(object):
    _module = 16
    _action = 6

    def __init__(self):
        pass
        self.partner = PartnerInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.partner.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.partner.decode(raw_msg[idx:])
        idx += self.partner.size()

    def size(self):
        size = 0
        size += self.partner.size()
        return size


class ChangeBuddyUp(object):
    _module = 16
    _action = 7

    def __init__(self):
        pass
        self.buddy_role_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.buddy_role_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.buddy_role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class ChangeBuddyDown(object):
    _module = 16
    _action = 7

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class GetFormInfoUp(object):
    _module = 16
    _action = 8

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
        

class GetFormInfoDown(object):
    _module = 16
    _action = 8

    def __init__(self):
        pass
        self.tactical = 0
        self.buddy_role_id = 0
        self.buddy_role_row = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.tactical))
        buff.extend(struct.pack("<b", self.buddy_role_id))
        buff.extend(struct.pack("<b", self.buddy_role_row))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.tactical = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.buddy_role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.buddy_role_row = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class ChangeFormUp(object):
    _module = 16
    _action = 9

    def __init__(self):
        pass
        self.main_role_row = 0
        self.buddy_role_row = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.main_role_row))
        buff.extend(struct.pack("<b", self.buddy_role_row))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.main_role_row = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.buddy_role_row = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class ChangeFormDown(object):
    _module = 16
    _action = 9

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class GetInfoUp(object):
    _module = 16
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
        

class GetInfoDown(object):
    _module = 16
    _action = 10

    def __init__(self):
        pass
        self.daily_num = 0
        self.lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.daily_num))
        buff.extend(struct.pack("<l", self.lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.daily_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class CancelAutoMatchUp(object):
    _module = 16
    _action = 11

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
        

class CancelAutoMatchDown(object):
    _module = 16
    _action = 11

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
        

class GetOnlineFriendUp(object):
    _module = 16
    _action = 12

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
        

class GetOnlineFriendDown(object):
    _module = 16
    _action = 12

    def __init__(self):
        pass
        self.friends = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.friends)))
        for item in self.friends:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _friends_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_friends_size):
            obj = GetOnlineFriendDownFriends()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.friends.append(obj)

    def size(self):
        size = 1
        for item in self.friends:
            size += item.size()
        return size


class GetOnlineFriendDownFriends(object):
    def __init__(self):
        pass
        self.pid = 0
        self.role_id = 0
        self.nickname = ''
        self.level = 0
        self.fight_num = 0
        self.daily_num = 0
        self.lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<b", self.daily_num))
        buff.extend(struct.pack("<l", self.lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.daily_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 22
        size += len(self.nickname)
        return size


class InviteIntoRoomUp(object):
    _module = 16
    _action = 13

    def __init__(self):
        pass
        self.pid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class InviteIntoRoomDown(object):
    _module = 16
    _action = 13

    def __init__(self):
        pass
        self.isOffline = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.isOffline))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.isOffline = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyRoomInviteDown(object):
    _module = 16
    _action = 14

    def __init__(self):
        pass
        self.room_id = 0
        self.level_id = 0
        self.nickname = ''
        self.inviter_id = 0
        self.daily_num = 0
        self.game_server_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.room_id))
        buff.extend(struct.pack("<l", self.level_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<q", self.inviter_id))
        buff.extend(struct.pack("<b", self.daily_num))
        buff.extend(struct.pack("<q", self.game_server_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.room_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.level_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.inviter_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.daily_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.game_server_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 31
        size += len(self.nickname)
        return size


class NotifyPlayersInfoDown(object):
    _module = 16
    _action = 15

    def __init__(self):
        pass
        self.players = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.players)))
        for item in self.players:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _players_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_players_size):
            obj = NotifyPlayersInfoDownPlayers()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.players.append(obj)

    def size(self):
        size = 1
        for item in self.players:
            size += item.size()
        return size


class NotifyPlayersInfoDownPlayers(object):
    def __init__(self):
        pass
        self.player_id = 0
        self.nickname = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

    def size(self):
        size = 10
        size += len(self.nickname)
        return size


class RefuseRoomInviteUp(object):
    _module = 16
    _action = 16

    def __init__(self):
        pass
        self.room_id = 0
        self.inviter_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.room_id))
        buff.extend(struct.pack("<q", self.inviter_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.room_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.inviter_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 18
        

class RefuseRoomInviteDown(object):
    _module = 16
    _action = 16

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
        

class NotifyRoomInviteRefuseDown(object):
    _module = 16
    _action = 17

    def __init__(self):
        pass
        self.nickname = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

    def size(self):
        size = 2
        size += len(self.nickname)
        return size


class NotifyUpdatePartnerDown(object):
    _module = 16
    _action = 18

    def __init__(self):
        pass
        self.partner = PartnerInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.partner.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.partner.decode(raw_msg[idx:])
        idx += self.partner.size()

    def size(self):
        size = 0
        size += self.partner.size()
        return size


class MultiLevelModule(interface.BaseModule):
    decoder_map = {
        1: CreateRoomDown, 
        2: EnterRoomDown, 
        3: AutoEnterRoomDown, 
        4: NotifyRoomInfoDown, 
        5: LeaveRoomDown, 
        6: NotifyJoinPartnerDown, 
        7: ChangeBuddyDown, 
        8: GetFormInfoDown, 
        9: ChangeFormDown, 
        10: GetInfoDown, 
        11: CancelAutoMatchDown, 
        12: GetOnlineFriendDown, 
        13: InviteIntoRoomDown, 
        14: NotifyRoomInviteDown, 
        15: NotifyPlayersInfoDown, 
        16: RefuseRoomInviteDown, 
        17: NotifyRoomInviteRefuseDown, 
        18: NotifyUpdatePartnerDown, 
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

    def add_create_room(self, callback):
        self.add_callback(1, callback)

    def add_enter_room(self, callback):
        self.add_callback(2, callback)

    def add_auto_enter_room(self, callback):
        self.add_callback(3, callback)

    def add_notify_room_info(self, callback):
        self.add_callback(4, callback)

    def add_leave_room(self, callback):
        self.add_callback(5, callback)

    def add_notify_join_partner(self, callback):
        self.add_callback(6, callback)

    def add_change_buddy(self, callback):
        self.add_callback(7, callback)

    def add_get_form_info(self, callback):
        self.add_callback(8, callback)

    def add_change_form(self, callback):
        self.add_callback(9, callback)

    def add_get_info(self, callback):
        self.add_callback(10, callback)

    def add_cancel_auto_match(self, callback):
        self.add_callback(11, callback)

    def add_get_online_friend(self, callback):
        self.add_callback(12, callback)

    def add_invite_into_room(self, callback):
        self.add_callback(13, callback)

    def add_notify_room_invite(self, callback):
        self.add_callback(14, callback)

    def add_notify_players_info(self, callback):
        self.add_callback(15, callback)

    def add_refuse_room_invite(self, callback):
        self.add_callback(16, callback)

    def add_notify_room_invite_refuse(self, callback):
        self.add_callback(17, callback)

    def add_notify_update_partner(self, callback):
        self.add_callback(18, callback)
