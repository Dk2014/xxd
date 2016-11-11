#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class FashionInfoUp(object):
    _module = 25
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
        

class FashionInfoDown(object):
    _module = 25
    _action = 1

    def __init__(self):
        pass
        self.dress_cd_time = 0
        self.dressed_fashion_id = 0
        self.fashions = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.dress_cd_time))
        buff.extend(struct.pack("<h", self.dressed_fashion_id))
        buff.extend(struct.pack('<B', len(self.fashions)))
        for item in self.fashions:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.dress_cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.dressed_fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _fashions_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_fashions_size):
            obj = FashionInfoDownFashions()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.fashions.append(obj)

    def size(self):
        size = 11
        for item in self.fashions:
            size += item.size()
        return size


class FashionInfoDownFashions(object):
    def __init__(self):
        pass
        self.fashion_id = 0
        self.expire_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<q", self.expire_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.expire_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class DressFashionUp(object):
    _module = 25
    _action = 2

    def __init__(self):
        pass
        self.fashion_id = 0
        self.in_clubhouse = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<?", self.in_clubhouse))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.in_clubhouse = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class DressFashionDown(object):
    _module = 25
    _action = 2

    def __init__(self):
        pass
        self.dress_cd_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.dress_cd_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.dress_cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class FashionModule(interface.BaseModule):
    decoder_map = {
        1: FashionInfoDown, 
        2: DressFashionDown, 
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

    def add_fashion_info(self, callback):
        self.add_callback(1, callback)

    def add_dress_fashion(self, callback):
        self.add_callback(2, callback)
