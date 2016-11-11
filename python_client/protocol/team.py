#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetFormationInfoUp(object):
    _module = 2
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
        

class GetFormationInfoDown(object):
    _module = 2
    _action = 0

    def __init__(self):
        pass
        self.pos0_role = 0
        self.pos1_role = 0
        self.pos2_role = 0
        self.pos3_role = 0
        self.pos4_role = 0
        self.pos5_role = 0
        self.pos6_role = 0
        self.pos7_role = 0
        self.pos8_role = 0
        self.relationship = 0
        self.health_lv = 0
        self.attack_lv = 0
        self.defence_lv = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.pos0_role))
        buff.extend(struct.pack("<b", self.pos1_role))
        buff.extend(struct.pack("<b", self.pos2_role))
        buff.extend(struct.pack("<b", self.pos3_role))
        buff.extend(struct.pack("<b", self.pos4_role))
        buff.extend(struct.pack("<b", self.pos5_role))
        buff.extend(struct.pack("<b", self.pos6_role))
        buff.extend(struct.pack("<b", self.pos7_role))
        buff.extend(struct.pack("<b", self.pos8_role))
        buff.extend(struct.pack("<l", self.relationship))
        buff.extend(struct.pack("<h", self.health_lv))
        buff.extend(struct.pack("<h", self.attack_lv))
        buff.extend(struct.pack("<h", self.defence_lv))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pos0_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos1_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos2_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos3_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos4_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos5_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos6_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos7_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos8_role = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.relationship = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.health_lv = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.attack_lv = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.defence_lv = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 19
        

class UpFormationUp(object):
    _module = 2
    _action = 2

    def __init__(self):
        pass
        self.role_id = 0
        self.pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class UpFormationDown(object):
    _module = 2
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
        

class DownFormationUp(object):
    _module = 2
    _action = 3

    def __init__(self):
        pass
        self.pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pos = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class DownFormationDown(object):
    _module = 2
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
        

class SwapFormationUp(object):
    _module = 2
    _action = 4

    def __init__(self):
        pass
        self.pos_from = 0
        self.pos_to = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.pos_from))
        buff.extend(struct.pack("<b", self.pos_to))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pos_from = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos_to = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class SwapFormationDown(object):
    _module = 2
    _action = 4

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
        

class ReplaceFormationUp(object):
    _module = 2
    _action = 5

    def __init__(self):
        pass
        self.role_id = 0
        self.pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pos = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class ReplaceFormationDown(object):
    _module = 2
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
        

class TrainingTeamshipUp(object):
    _module = 2
    _action = 6

    def __init__(self):
        pass
        self.attr_ind = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.attr_ind))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.attr_ind = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class TrainingTeamshipDown(object):
    _module = 2
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
        return 0
        

class TeamModule(interface.BaseModule):
    decoder_map = {
        0: GetFormationInfoDown, 
        2: UpFormationDown, 
        3: DownFormationDown, 
        4: SwapFormationDown, 
        5: ReplaceFormationDown, 
        6: TrainingTeamshipDown, 
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

    def add_get_formation_info(self, callback):
        self.add_callback(0, callback)

    def add_up_formation(self, callback):
        self.add_callback(2, callback)

    def add_down_formation(self, callback):
        self.add_callback(3, callback)

    def add_swap_formation(self, callback):
        self.add_callback(4, callback)

    def add_replace_formation(self, callback):
        self.add_callback(5, callback)

    def add_training_teamship(self, callback):
        self.add_callback(6, callback)
