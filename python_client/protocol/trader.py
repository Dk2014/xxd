#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class InfoUp(object):
    _module = 21
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
    _module = 21
    _action = 1

    def __init__(self):
        pass
        self.during = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.during)))
        for item in self.during:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _during_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_during_size):
            obj = InfoDownDuring()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.during.append(obj)

    def size(self):
        size = 1
        for item in self.during:
            size += item.size()
        return size


class InfoDownDuring(object):
    def __init__(self):
        pass
        self.expire = 0
        self.showup = 0
        self.disappear = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.expire))
        buff.extend(struct.pack("<q", self.showup))
        buff.extend(struct.pack("<q", self.disappear))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.expire = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.showup = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.disappear = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 24
        

class StoreStateUp(object):
    _module = 21
    _action = 2

    def __init__(self):
        pass
        self.trader_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.trader_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.trader_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class StoreStateDown(object):
    _module = 21
    _action = 2

    def __init__(self):
        pass
        self.refresh_num = 0
        self.goods = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.refresh_num))
        buff.extend(struct.pack('<B', len(self.goods)))
        for item in self.goods:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.refresh_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _goods_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_goods_size):
            obj = StoreStateDownGoods()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.goods.append(obj)

    def size(self):
        size = 3
        for item in self.goods:
            size += item.size()
        return size


class StoreStateDownGoods(object):
    def __init__(self):
        pass
        self.grid_id = 0
        self.cost = 0
        self.goods_type = 0
        self.item_id = 0
        self.num = 0
        self.stock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.grid_id))
        buff.extend(struct.pack("<q", self.cost))
        buff.extend(struct.pack("<b", self.goods_type))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<h", self.num))
        buff.extend(struct.pack("<b", self.stock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.grid_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.cost = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.goods_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.stock = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 18
        

class BuyUp(object):
    _module = 21
    _action = 3

    def __init__(self):
        pass
        self.grid_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.grid_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.grid_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class BuyDown(object):
    _module = 21
    _action = 3

    def __init__(self):
        pass
        self.expired = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.expired))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.expired = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class RefreshUp(object):
    _module = 21
    _action = 4

    def __init__(self):
        pass
        self.trader_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.trader_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.trader_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class RefreshDown(object):
    _module = 21
    _action = 4

    def __init__(self):
        pass
        self.goods = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.goods)))
        for item in self.goods:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _goods_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_goods_size):
            obj = RefreshDownGoods()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.goods.append(obj)

    def size(self):
        size = 1
        for item in self.goods:
            size += item.size()
        return size


class RefreshDownGoods(object):
    def __init__(self):
        pass
        self.grid_id = 0
        self.goods_type = 0
        self.cost = 0
        self.item_id = 0
        self.num = 0
        self.stock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.grid_id))
        buff.extend(struct.pack("<b", self.goods_type))
        buff.extend(struct.pack("<q", self.cost))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<h", self.num))
        buff.extend(struct.pack("<b", self.stock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.grid_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.goods_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.cost = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.stock = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 18
        

class TraderModule(interface.BaseModule):
    decoder_map = {
        1: InfoDown, 
        2: StoreStateDown, 
        3: BuyDown, 
        4: RefreshDown, 
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

    def add_store_state(self, callback):
        self.add_callback(2, callback)

    def add_buy(self, callback):
        self.add_callback(3, callback)

    def add_refresh(self, callback):
        self.add_callback(4, callback)
