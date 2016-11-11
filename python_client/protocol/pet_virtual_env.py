#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class InfoUp(object):
    _module = 28
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
        

class InfoDown(object):
    _module = 28
    _action = 1

    def __init__(self):
        pass
        self.daily_num = 0
        self.max_floor = 0
        self.max_awarded_floor = 0
        self.unpassed_floor_enemy_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.daily_num))
        buff.extend(struct.pack("<h", self.max_floor))
        buff.extend(struct.pack("<h", self.max_awarded_floor))
        buff.extend(struct.pack("<h", self.unpassed_floor_enemy_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.daily_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.max_floor = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.max_awarded_floor = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.unpassed_floor_enemy_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 7
        

class TakeAwardUp(object):
    _module = 28
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
        

class TakeAwardDown(object):
    _module = 28
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
        return 0
        

class AutoFightUp(object):
    _module = 28
    _action = 3

    def __init__(self):
        pass
        self.floor = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.floor))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.floor = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class AutoFightDown(object):
    _module = 28
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
        return 0
        

class PveKillsDown(object):
    _module = 28
    _action = 4

    def __init__(self):
        pass
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class PetVirtualEnvModule(interface.BaseModule):
    decoder_map = {
        1: InfoDown, 
        2: TakeAwardDown, 
        3: AutoFightDown, 
        4: PveKillsDown, 
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
        self.add_callback(1, callback)

    def add_take_award(self, callback):
        self.add_callback(2, callback)

    def add_auto_fight(self, callback):
        self.add_callback(3, callback)

    def add_pve_kills(self, callback):
        self.add_callback(4, callback)
