#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetTreeStatusUp(object):
    _module = 32
    _action = 0

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
        

class GetTreeStatusDown(object):
    _module = 32
    _action = 0

    def __init__(self):
        pass
        self.times = 0
        self.money = 0
        self.last_time = 0
        self.remind = 0
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.times))
        buff.extend(struct.pack("<l", self.money))
        buff.extend(struct.pack("<q", self.last_time))
        buff.extend(struct.pack("<b", self.remind))
        buff.extend(struct.pack("<b", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.times = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.money = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.last_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.remind = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 15
        

class GetTreeMoneyUp(object):
    _module = 32
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
        

class GetTreeMoneyDown(object):
    _module = 32
    _action = 1

    def __init__(self):
        pass
        self.code = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.code))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.code = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class WaveTreeUp(object):
    _module = 32
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
        

class WaveTreeDown(object):
    _module = 32
    _action = 2

    def __init__(self):
        pass
        self.status = 0
        self.money = 0
        self.remaind = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<l", self.money))
        buff.extend(struct.pack("<b", self.remaind))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.money = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.remaind = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 6
        

class MoneyTreeModule(interface.BaseModule):
    decoder_map = {
        0: GetTreeStatusDown, 
        1: GetTreeMoneyDown, 
        2: WaveTreeDown, 
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

    def add_get_tree_status(self, callback):
        self.add_callback(0, callback)

    def add_get_tree_money(self, callback):
        self.add_callback(1, callback)

    def add_wave_tree(self, callback):
        self.add_callback(2, callback)
