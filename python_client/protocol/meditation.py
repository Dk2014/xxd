#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class MeditationInfoUp(object):
    _module = 27
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
        

class MeditationInfoDown(object):
    _module = 27
    _action = 1

    def __init__(self):
        pass
        self.exp_accumulate_time = 0
        self.key_accumulate_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.exp_accumulate_time))
        buff.extend(struct.pack("<l", self.key_accumulate_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.exp_accumulate_time = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.key_accumulate_time = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 8
        

class StartMeditationUp(object):
    _module = 27
    _action = 2

    def __init__(self):
        pass
        self.in_clubhouse = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.in_clubhouse))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.in_clubhouse = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class StartMeditationDown(object):
    _module = 27
    _action = 2

    def __init__(self):
        pass
        self.exp_accumulate_time = 0
        self.key_accumulate_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.exp_accumulate_time))
        buff.extend(struct.pack("<l", self.key_accumulate_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.exp_accumulate_time = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.key_accumulate_time = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 8
        

class StopMeditationUp(object):
    _module = 27
    _action = 3

    def __init__(self):
        pass
        self.in_clubhouse = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.in_clubhouse))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.in_clubhouse = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class StopMeditationDown(object):
    _module = 27
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
        

class MeditationModule(interface.BaseModule):
    decoder_map = {
        1: MeditationInfoDown, 
        2: StartMeditationDown, 
        3: StopMeditationDown, 
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

    def add_meditation_info(self, callback):
        self.add_callback(1, callback)

    def add_start_meditation(self, callback):
        self.add_callback(2, callback)

    def add_stop_meditation(self, callback):
        self.add_callback(3, callback)
