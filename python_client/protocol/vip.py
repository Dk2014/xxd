#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

BUY_TIMES_TYPE_BIWUCHANGCISHU = 0
BUY_TIMES_TYPE_RAINBOWSAODANG = 1


class InfoUp(object):
    _module = 20
    _action = 1

    def __init__(self):
        pass
        self.player_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class InfoDown(object):
    _module = 20
    _action = 1

    def __init__(self):
        pass
        self.level = 0
        self.ingot = 0
        self.card_id = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.ingot))
        buff.extend(struct.pack('<H', len(self.card_id)))
        buff.extend(self.card_id)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.ingot = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _card_id_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.card_id = str(raw_msg[idx:idx+_card_id_size])
        idx += _card_id_size

    def size(self):
        size = 12
        size += len(self.card_id)
        return size


class GetTotalUp(object):
    _module = 20
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
        

class GetTotalDown(object):
    _module = 20
    _action = 2

    def __init__(self):
        pass
        self.total = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.total))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.total = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class VipLevelTotalUp(object):
    _module = 20
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
        

class VipLevelTotalDown(object):
    _module = 20
    _action = 3

    def __init__(self):
        pass
        self.vip_level_arr = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.vip_level_arr)))
        for item in self.vip_level_arr:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _vip_level_arr_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_vip_level_arr_size):
            obj = VipLevelTotalDownVipLevelArr()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.vip_level_arr.append(obj)

    def size(self):
        size = 1
        for item in self.vip_level_arr:
            size += item.size()
        return size


class VipLevelTotalDownVipLevelArr(object):
    def __init__(self):
        pass
        self.vip_level = 0
        self.total = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.vip_level))
        buff.extend(struct.pack("<l", self.total))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.vip_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.total = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class BuyTimesUp(object):
    _module = 20
    _action = 4

    def __init__(self):
        pass
        self.buytype = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.buytype))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.buytype = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class BuyTimesDown(object):
    _module = 20
    _action = 4

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
        

class VipModule(interface.BaseModule):
    decoder_map = {
        1: InfoDown, 
        2: GetTotalDown, 
        3: VipLevelTotalDown, 
        4: BuyTimesDown, 
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

    def add_get_total(self, callback):
        self.add_callback(2, callback)

    def add_vip_level_total(self, callback):
        self.add_callback(3, callback)

    def add_buy_times(self, callback):
        self.add_callback(4, callback)
