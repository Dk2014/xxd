#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetCliqueDailyQuestUp(object):
    _module = 35
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
        

class GetCliqueDailyQuestDown(object):
    _module = 35
    _action = 1

    def __init__(self):
        pass
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetCliqueDailyQuestDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 1
        for item in self.quest:
            size += item.size()
        return size


class GetCliqueDailyQuestDownQuest(object):
    def __init__(self):
        pass
        self.id = 0
        self.finish_count = 0
        self.award_state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        buff.extend(struct.pack("<h", self.finish_count))
        buff.extend(struct.pack("<b", self.award_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.finish_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.award_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class AwardCliqueDailyQuestUp(object):
    _module = 35
    _action = 2

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class AwardCliqueDailyQuestDown(object):
    _module = 35
    _action = 2

    def __init__(self):
        pass
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyCliqueDailyChangeDown(object):
    _module = 35
    _action = 3

    def __init__(self):
        pass
        self.id = 0
        self.finish_count = 0
        self.award_state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        buff.extend(struct.pack("<h", self.finish_count))
        buff.extend(struct.pack("<b", self.award_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.finish_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.award_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class GetCliqueBuildingQuestUp(object):
    _module = 35
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
        

class GetCliqueBuildingQuestDown(object):
    _module = 35
    _action = 4

    def __init__(self):
        pass
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetCliqueBuildingQuestDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 1
        for item in self.quest:
            size += item.size()
        return size


class GetCliqueBuildingQuestDownQuest(object):
    def __init__(self):
        pass
        self.id = 0
        self.award_state = 0
        self.donateCoins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        buff.extend(struct.pack("<b", self.award_state))
        buff.extend(struct.pack("<q", self.donateCoins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.award_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.donateCoins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 11
        

class AwardCliqueBuildingQuestUp(object):
    _module = 35
    _action = 5

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class AwardCliqueBuildingQuestDown(object):
    _module = 35
    _action = 5

    def __init__(self):
        pass
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class CliqueQuestModule(interface.BaseModule):
    decoder_map = {
        1: GetCliqueDailyQuestDown, 
        2: AwardCliqueDailyQuestDown, 
        3: NotifyCliqueDailyChangeDown, 
        4: GetCliqueBuildingQuestDown, 
        5: AwardCliqueBuildingQuestDown, 
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

    def add_get_clique_daily_Quest(self, callback):
        self.add_callback(1, callback)

    def add_award_clique_daily_Quest(self, callback):
        self.add_callback(2, callback)

    def add_notify_clique_daily_change(self, callback):
        self.add_callback(3, callback)

    def add_get_clique_building_quest(self, callback):
        self.add_callback(4, callback)

    def add_award_clique_building_Quest(self, callback):
        self.add_callback(5, callback)
