#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetInfoUp(object):
    _module = 15
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
        

class GetInfoDown(object):
    _module = 15
    _action = 1

    def __init__(self):
        pass
        self.floor_num = 0
        self.friends = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.floor_num))
        buff.extend(struct.pack('<B', len(self.friends)))
        for item in self.friends:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.floor_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _friends_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_friends_size):
            obj = GetInfoDownFriends()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.friends.append(obj)

    def size(self):
        size = 3
        for item in self.friends:
            size += item.size()
        return size


class GetInfoDownFriends(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.nickname = ''
        self.level = 0
        self.floor = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<h", self.floor))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.floor = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 7
        size += len(self.nickname)
        return size


class UseLadderUp(object):
    _module = 15
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
        

class UseLadderDown(object):
    _module = 15
    _action = 2

    def __init__(self):
        pass
        self.floor_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.floor_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.floor_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class TowerModule(interface.BaseModule):
    decoder_map = {
        1: GetInfoDown, 
        2: UseLadderDown, 
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

    def add_get_info(self, callback):
        self.add_callback(1, callback)

    def add_use_ladder(self, callback):
        self.add_callback(2, callback)
