#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

AWARD_TYPE_NEW_PLAYER_AWARD = 0
AWARD_TYPE_REGULAR_AWARD = 1


class InfoUp(object):
    _module = 22
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
    _module = 22
    _action = 1

    def __init__(self):
        pass
        self.index = 0
        self.records = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.index))
        buff.extend(struct.pack('<B', len(self.records)))
        for item in self.records:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.index = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _records_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_records_size):
            obj = InfoDownRecords()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.records.append(obj)

    def size(self):
        size = 2
        for item in self.records:
            size += item.size()
        return size


class InfoDownRecords(object):
    def __init__(self):
        pass
        self.award_type = 0
        self.award_index = 0
        self.signed = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.award_type))
        buff.extend(struct.pack("<b", self.award_index))
        buff.extend(struct.pack("<?", self.signed))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.award_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.award_index = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.signed = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class SignUp(object):
    _module = 22
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
        

class SignDown(object):
    _module = 22
    _action = 2

    def __init__(self):
        pass
        self.expired = False
        self.index = 0
        self.records = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.expired))
        buff.extend(struct.pack("<b", self.index))
        buff.extend(struct.pack('<B', len(self.records)))
        for item in self.records:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.expired = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.index = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _records_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_records_size):
            obj = SignDownRecords()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.records.append(obj)

    def size(self):
        size = 3
        for item in self.records:
            size += item.size()
        return size


class SignDownRecords(object):
    def __init__(self):
        pass
        self.award_type = 0
        self.award_index = 0
        self.signed = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.award_type))
        buff.extend(struct.pack("<b", self.award_index))
        buff.extend(struct.pack("<?", self.signed))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.award_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.award_index = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.signed = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class SignPastDayUp(object):
    _module = 22
    _action = 3

    def __init__(self):
        pass
        self.index = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.index))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.index = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class SignPastDayDown(object):
    _module = 22
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
        

class DailySignInModule(interface.BaseModule):
    decoder_map = {
        1: InfoDown, 
        2: SignDown, 
        3: SignPastDayDown, 
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

    def add_sign(self, callback):
        self.add_callback(2, callback)

    def add_sign_past_day(self, callback):
        self.add_callback(3, callback)
