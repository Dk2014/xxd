#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

EXCHANGE_GIFT_RESULT_SUCCESS = 0
EXCHANGE_GIFT_RESULT_EXPIRE = 1
EXCHANGE_GIFT_RESULT_DUP_EXCHANGE = 2


CHEST_TYPE_COIN = 0
CHEST_TYPE_COIN_FREE = 1
CHEST_TYPE_COIN_TEN = 2
CHEST_TYPE_INGOT = 3
CHEST_TYPE_INGOT_FREE = 4
CHEST_TYPE_INGOT_TEN = 5
CHEST_TYPE_PET = 6
CHEST_TYPE_PET_FREE = 7
CHEST_TYPE_PET_TEN = 8


ITEM_TYPE_COIN = 1
ITEM_TYPE_INGOT = 2
ITEM_TYPE_ITEM = 3
ITEM_TYPE_GHOST = 4
ITEM_TYPE_SWORD_SOUL = 5
ITEM_TYPE_PET = 6
ITEM_TYPE_GHOST_FRAGMENT = 7
ITEM_TYPE_PREFERENCE = 8
ITEM_TYPE_EQUIPMENT = 9


class AwardInfo(object):
    def __init__(self):
        pass
        self.award_type = 0
        self.award_num = 0
        self.item_id = 0
        self.draw_time = 0
        self.award_index = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.award_type))
        buff.extend(struct.pack("<h", self.award_num))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<q", self.draw_time))
        buff.extend(struct.pack("<h", self.award_index))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.award_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.award_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.draw_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.award_index = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 15
        

class GetHeartDrawInfoUp(object):
    _module = 50
    _action = 1

    def __init__(self):
        pass
        self.draw_type = 0
        self.award_record = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.draw_type))
        buff.extend(struct.pack("<?", self.award_record))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.draw_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.award_record = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class GetHeartDrawInfoDown(object):
    _module = 50
    _action = 1

    def __init__(self):
        pass
        self.hearts = 0
        self.daily_num = 0
        self.award_record = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.hearts))
        buff.extend(struct.pack("<b", self.daily_num))
        buff.extend(struct.pack('<B', len(self.award_record)))
        for item in self.award_record:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.hearts = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.daily_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _award_record_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_award_record_size):
            obj = GetHeartDrawInfoDownAwardRecord()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.award_record.append(obj)

    def size(self):
        size = 4
        for item in self.award_record:
            size += item.size()
        return size


class GetHeartDrawInfoDownAwardRecord(object):
    def __init__(self):
        pass
        self.award = AwardInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.award.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.award.decode(raw_msg[idx:])
        idx += self.award.size()

    def size(self):
        size = 0
        size += self.award.size()
        return size


class HeartDrawUp(object):
    _module = 50
    _action = 2

    def __init__(self):
        pass
        self.draw_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.draw_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.draw_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class HeartDrawDown(object):
    _module = 50
    _action = 2

    def __init__(self):
        pass
        self.award = AwardInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.award.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.award.decode(raw_msg[idx:])
        idx += self.award.size()

    def size(self):
        size = 0
        size += self.award.size()
        return size


class GetChestInfoUp(object):
    _module = 50
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
        

class GetChestInfoDown(object):
    _module = 50
    _action = 3

    def __init__(self):
        pass
        self.free_coin_num = 0
        self.next_free_coin_left = 0
        self.next_free_ingot_left = 0
        self.next_free_pet_left = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.free_coin_num))
        buff.extend(struct.pack("<q", self.next_free_coin_left))
        buff.extend(struct.pack("<q", self.next_free_ingot_left))
        buff.extend(struct.pack("<q", self.next_free_pet_left))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.free_coin_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.next_free_coin_left = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.next_free_ingot_left = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.next_free_pet_left = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 28
        

class DrawChestUp(object):
    _module = 50
    _action = 4

    def __init__(self):
        pass
        self.chest_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.chest_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.chest_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class DrawChestDown(object):
    _module = 50
    _action = 4

    def __init__(self):
        pass
        self.items = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.items)))
        for item in self.items:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _items_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_items_size):
            obj = DrawChestDownItems()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.items.append(obj)

    def size(self):
        size = 1
        for item in self.items:
            size += item.size()
        return size


class DrawChestDownItems(object):
    def __init__(self):
        pass
        self.item_type = 0
        self.item_id = 0
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.item_type))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<l", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.item_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 7
        

class HeartInfoUp(object):
    _module = 50
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
        

class HeartInfoDown(object):
    _module = 50
    _action = 5

    def __init__(self):
        pass
        self.recover_today = 0
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.recover_today))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.recover_today = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class GetFateBoxInfoUp(object):
    _module = 50
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
        return 2
        

class GetFateBoxInfoDown(object):
    _module = 50
    _action = 6

    def __init__(self):
        pass
        self.lock = 0
        self.next_free_star_box = 0
        self.next_free_moon_box = 0
        self.next_free_sun_box = 0
        self.next_free_hunyuan_box = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.lock))
        buff.extend(struct.pack("<l", self.next_free_star_box))
        buff.extend(struct.pack("<l", self.next_free_moon_box))
        buff.extend(struct.pack("<l", self.next_free_sun_box))
        buff.extend(struct.pack("<l", self.next_free_hunyuan_box))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.next_free_star_box = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.next_free_moon_box = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.next_free_sun_box = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.next_free_hunyuan_box = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 20
        

class OpenFateBoxUp(object):
    _module = 50
    _action = 7

    def __init__(self):
        pass
        self.box_type = 0
        self.times = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.box_type))
        buff.extend(struct.pack("<b", self.times))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.box_type = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.times = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 7
        

class OpenFateBoxDown(object):
    _module = 50
    _action = 7

    def __init__(self):
        pass
        self.items = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.items)))
        for item in self.items:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _items_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_items_size):
            obj = OpenFateBoxDownItems()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.items.append(obj)

    def size(self):
        size = 1
        for item in self.items:
            size += item.size()
        return size


class OpenFateBoxDownItems(object):
    def __init__(self):
        pass
        self.item_type = 0
        self.item_id = 0
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.item_type))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<l", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.item_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 7
        

class ExchangeGiftCodeUp(object):
    _module = 50
    _action = 8

    def __init__(self):
        pass
        self.code = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.code)))
        buff.extend(self.code)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _code_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.code = str(raw_msg[idx:idx+_code_size])
        idx += _code_size

    def size(self):
        size = 4
        size += len(self.code)
        return size


class ExchangeGiftCodeDown(object):
    _module = 50
    _action = 8

    def __init__(self):
        pass
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class DrawModule(interface.BaseModule):
    decoder_map = {
        1: GetHeartDrawInfoDown, 
        2: HeartDrawDown, 
        3: GetChestInfoDown, 
        4: DrawChestDown, 
        5: HeartInfoDown, 
        6: GetFateBoxInfoDown, 
        7: OpenFateBoxDown, 
        8: ExchangeGiftCodeDown, 
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

    def add_get_heart_draw_info(self, callback):
        self.add_callback(1, callback)

    def add_heart_draw(self, callback):
        self.add_callback(2, callback)

    def add_get_chest_info(self, callback):
        self.add_callback(3, callback)

    def add_draw_chest(self, callback):
        self.add_callback(4, callback)

    def add_heart_info(self, callback):
        self.add_callback(5, callback)

    def add_get_fate_box_info(self, callback):
        self.add_callback(6, callback)

    def add_open_fate_box(self, callback):
        self.add_callback(7, callback)

    def add_exchange_gift_code(self, callback):
        self.add_callback(8, callback)
