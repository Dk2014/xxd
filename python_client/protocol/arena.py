#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

NOTIFY_ARENA_MODE_ATTACKED_SUCC = 3
NOTIFY_ARENA_MODE_ATTACKED_FAILED = 4


class EnterUp(object):
    _module = 19
    _action = 1

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
        

class EnterDown(object):
    _module = 19
    _action = 1

    def __init__(self):
        pass
        self.rank = 0
        self.award_box_time = 0
        self.daily_num = 0
        self.buy_times = 0
        self.win_times = 0
        self.new_record_num = 0
        self.failed_cd_time = 0
        self.daily_award_item = 0
        self.daily_fame = 0
        self.award_box = []
        self.ranks = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.rank))
        buff.extend(struct.pack("<q", self.award_box_time))
        buff.extend(struct.pack("<h", self.daily_num))
        buff.extend(struct.pack("<h", self.buy_times))
        buff.extend(struct.pack("<h", self.win_times))
        buff.extend(struct.pack("<h", self.new_record_num))
        buff.extend(struct.pack("<q", self.failed_cd_time))
        buff.extend(struct.pack("<l", self.daily_award_item))
        buff.extend(struct.pack("<l", self.daily_fame))
        buff.extend(struct.pack('<B', len(self.award_box)))
        for item in self.award_box:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.ranks)))
        for item in self.ranks:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.award_box_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.daily_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.buy_times = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.win_times = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.new_record_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.failed_cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.daily_award_item = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.daily_fame = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _award_box_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_award_box_size):
            obj = EnterDownAwardBox()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.award_box.append(obj)

        _ranks_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_ranks_size):
            obj = EnterDownRanks()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.ranks.append(obj)

    def size(self):
        size = 38
        for item in self.award_box:
            size += item.size()
        for item in self.ranks:
            size += item.size()
        return size


class EnterDownAwardBox(object):
    def __init__(self):
        pass
        self.num = 0
        self.rank = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.num))
        buff.extend(struct.pack("<l", self.rank))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class EnterDownRanks(object):
    def __init__(self):
        pass
        self.pid = 0
        self.nick = ''
        self.role_id = 0
        self.rank = 0
        self.level = 0
        self.fight_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<l", self.rank))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.fight_num))
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

        self.rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 21
        size += len(self.nick)
        return size


class GetRecordsUp(object):
    _module = 19
    _action = 2

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
        

class GetRecordsDown(object):
    _module = 19
    _action = 2

    def __init__(self):
        pass
        self.records = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.records)))
        for item in self.records:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _records_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_records_size):
            obj = GetRecordsDownRecords()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.records.append(obj)

    def size(self):
        size = 1
        for item in self.records:
            size += item.size()
        return size


class GetRecordsDownRecords(object):
    def __init__(self):
        pass
        self.mode = 0
        self.time = 0
        self.old_rank = 0
        self.new_rank = 0
        self.fight_num = 0
        self.target_pid = 0
        self.target_nick = ''
        self.target_old_rank = 0
        self.target_new_rank = 0
        self.target_fight_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.mode))
        buff.extend(struct.pack("<q", self.time))
        buff.extend(struct.pack("<l", self.old_rank))
        buff.extend(struct.pack("<l", self.new_rank))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<q", self.target_pid))
        buff.extend(struct.pack('<H', len(self.target_nick)))
        buff.extend(self.target_nick)
        buff.extend(struct.pack("<l", self.target_old_rank))
        buff.extend(struct.pack("<l", self.target_new_rank))
        buff.extend(struct.pack("<l", self.target_fight_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.mode = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.old_rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.new_rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.target_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _target_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.target_nick = str(raw_msg[idx:idx+_target_nick_size])
        idx += _target_nick_size

        self.target_old_rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.target_new_rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.target_fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 43
        size += len(self.target_nick)
        return size


class AwardBoxUp(object):
    _module = 19
    _action = 3

    def __init__(self):
        pass
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class AwardBoxDown(object):
    _module = 19
    _action = 3

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
        

class NotifyFailedCdTimeDown(object):
    _module = 19
    _action = 4

    def __init__(self):
        pass
        self.cd_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.cd_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class StartBattleUp(object):
    _module = 19
    _action = 5

    def __init__(self):
        pass
        self.player_id = 0
        self.player_rank = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<l", self.player_rank))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.player_rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 14
        

class StartBattleDown(object):
    _module = 19
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
        return 0
        

class NotifyLoseRankDown(object):
    _module = 19
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
        

class NotifyArenaInfoDown(object):
    _module = 19
    _action = 7

    def __init__(self):
        pass
        self.notify_type = 0
        self.pid = 0
        self.nick = ''
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.notify_type))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<l", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.notify_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 15
        size += len(self.nick)
        return size


class GetTopRankUp(object):
    _module = 19
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
        

class GetTopRankDown(object):
    _module = 19
    _action = 8

    def __init__(self):
        pass
        self.top3 = []
        self.ranks = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.top3)))
        for item in self.top3:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.ranks)))
        for item in self.ranks:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _top3_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_top3_size):
            obj = GetTopRankDownTop3()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.top3.append(obj)

        _ranks_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_ranks_size):
            obj = GetTopRankDownRanks()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.ranks.append(obj)

    def size(self):
        size = 2
        for item in self.top3:
            size += item.size()
        for item in self.ranks:
            size += item.size()
        return size


class GetTopRankDownTop3(object):
    def __init__(self):
        pass
        self.openid = ''
        self.role_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        buff.extend(struct.pack("<b", self.role_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 3
        size += len(self.openid)
        return size


class GetTopRankDownRanks(object):
    def __init__(self):
        pass
        self.pid = 0
        self.rank = 0
        self.nick = ''
        self.level = 0
        self.trend = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<l", self.rank))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<b", self.trend))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.trend = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 17
        size += len(self.nick)
        return size


class CleanFailedCdTimeUp(object):
    _module = 19
    _action = 9

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
        

class CleanFailedCdTimeDown(object):
    _module = 19
    _action = 9

    def __init__(self):
        pass
        self.failed_cd_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.failed_cd_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.failed_cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class ArenaModule(interface.BaseModule):
    decoder_map = {
        1: EnterDown, 
        2: GetRecordsDown, 
        3: AwardBoxDown, 
        4: NotifyFailedCdTimeDown, 
        5: StartBattleDown, 
        6: NotifyLoseRankDown, 
        7: NotifyArenaInfoDown, 
        8: GetTopRankDown, 
        9: CleanFailedCdTimeDown, 
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

    def add_enter(self, callback):
        self.add_callback(1, callback)

    def add_get_records(self, callback):
        self.add_callback(2, callback)

    def add_award_box(self, callback):
        self.add_callback(3, callback)

    def add_notify_failed_cd_time(self, callback):
        self.add_callback(4, callback)

    def add_start_battle(self, callback):
        self.add_callback(5, callback)

    def add_notify_lose_rank(self, callback):
        self.add_callback(6, callback)

    def add_notify_arena_info(self, callback):
        self.add_callback(7, callback)

    def add_get_top_rank(self, callback):
        self.add_callback(8, callback)

    def add_clean_failed_cd_time(self, callback):
        self.add_callback(9, callback)
