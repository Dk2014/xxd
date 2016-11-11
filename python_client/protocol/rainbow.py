#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class InfoUp(object):
    _module = 23
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
    _module = 23
    _action = 1

    def __init__(self):
        pass
        self.segment_num = 0
        self.level_order = 0
        self.status = 0
        self.reset_num = 0
        self.max_segment_can_jump = 0
        self.max_pass_segment = 0
        self.auto_fight_num = 0
        self.buy_times = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.segment_num))
        buff.extend(struct.pack("<b", self.level_order))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<b", self.reset_num))
        buff.extend(struct.pack("<h", self.max_segment_can_jump))
        buff.extend(struct.pack("<h", self.max_pass_segment))
        buff.extend(struct.pack("<b", self.auto_fight_num))
        buff.extend(struct.pack("<h", self.buy_times))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.segment_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.level_order = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.reset_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.max_segment_can_jump = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.max_pass_segment = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.auto_fight_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.buy_times = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 12
        

class ResetUp(object):
    _module = 23
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
        

class ResetDown(object):
    _module = 23
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
        

class AwardInfoUp(object):
    _module = 23
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
        

class AwardInfoDown(object):
    _module = 23
    _action = 3

    def __init__(self):
        pass
        self.award = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.award)))
        for item in self.award:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _award_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_award_size):
            obj = AwardInfoDownAward()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.award.append(obj)

    def size(self):
        size = 1
        for item in self.award:
            size += item.size()
        return size


class AwardInfoDownAward(object):
    def __init__(self):
        pass
        self.order = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.order))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.order = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class TakeAwardUp(object):
    _module = 23
    _action = 4

    def __init__(self):
        pass
        self.pos1 = 0
        self.pos2 = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.pos1))
        buff.extend(struct.pack("<b", self.pos2))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pos1 = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos2 = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class TakeAwardDown(object):
    _module = 23
    _action = 4

    def __init__(self):
        pass
        self.next_level = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.next_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.next_level = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class JumpToSegmentUp(object):
    _module = 23
    _action = 5

    def __init__(self):
        pass
        self.segment = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.segment))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.segment = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class JumpToSegmentDown(object):
    _module = 23
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
        

class AutoFightUp(object):
    _module = 23
    _action = 6

    def __init__(self):
        pass
        self.segment = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.segment))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.segment = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class AutoFightDown(object):
    _module = 23
    _action = 6

    def __init__(self):
        pass
        self.awardCoin = 0
        self.awardExp = 0
        self.awardBoxPos1 = 0
        self.awardBoxPos2 = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.awardCoin))
        buff.extend(struct.pack("<l", self.awardExp))
        buff.extend(struct.pack("<b", self.awardBoxPos1))
        buff.extend(struct.pack("<b", self.awardBoxPos2))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.awardCoin = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.awardExp = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.awardBoxPos1 = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.awardBoxPos2 = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 10
        

class RainbowModule(interface.BaseModule):
    decoder_map = {
        1: InfoDown, 
        2: ResetDown, 
        3: AwardInfoDown, 
        4: TakeAwardDown, 
        5: JumpToSegmentDown, 
        6: AutoFightDown, 
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

    def add_reset(self, callback):
        self.add_callback(2, callback)

    def add_award_info(self, callback):
        self.add_callback(3, callback)

    def add_take_award(self, callback):
        self.add_callback(4, callback)

    def add_jump_to_segment(self, callback):
        self.add_callback(5, callback)

    def add_auto_fight(self, callback):
        self.add_callback(6, callback)
