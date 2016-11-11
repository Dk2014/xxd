#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetPetInfoUp(object):
    _module = 17
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
        

class GetPetInfoDown(object):
    _module = 17
    _action = 1

    def __init__(self):
        pass
        self.pet = []
        self.set = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.pet)))
        for item in self.pet:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.set)))
        for item in self.set:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _pet_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_pet_size):
            obj = GetPetInfoDownPet()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.pet.append(obj)

        _set_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_set_size):
            obj = GetPetInfoDownSet()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.set.append(obj)

    def size(self):
        size = 2
        for item in self.pet:
            size += item.size()
        for item in self.set:
            size += item.size()
        return size


class GetPetInfoDownPet(object):
    def __init__(self):
        pass
        self.pet_id = 0
        self.level = 0
        self.exp = 0
        self.skill_level = 0
        self.called = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.pet_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.skill_level))
        buff.extend(struct.pack("<?", self.called))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pet_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.skill_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.called = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 17
        

class GetPetInfoDownSet(object):
    def __init__(self):
        pass
        self.grid_num = 0
        self.pet_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.grid_num))
        buff.extend(struct.pack("<l", self.pet_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.grid_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pet_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class SetPetUp(object):
    _module = 17
    _action = 3

    def __init__(self):
        pass
        self.grid_num = 0
        self.pet_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.grid_num))
        buff.extend(struct.pack("<l", self.pet_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.grid_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.pet_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 7
        

class SetPetDown(object):
    _module = 17
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
        

class SetPetSwapUp(object):
    _module = 17
    _action = 4

    def __init__(self):
        pass
        self.from_grid_num = 0
        self.to_grid_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.from_grid_num))
        buff.extend(struct.pack("<b", self.to_grid_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.from_grid_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.to_grid_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class SetPetSwapDown(object):
    _module = 17
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
        

class UpgradePetUp(object):
    _module = 17
    _action = 6

    def __init__(self):
        pass
        self.pet_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.pet_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pet_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class UpgradePetDown(object):
    _module = 17
    _action = 6

    def __init__(self):
        pass
        self.exp = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 10
        

class TrainingPetSkillUp(object):
    _module = 17
    _action = 7

    def __init__(self):
        pass
        self.pet_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.pet_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pet_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TrainingPetSkillDown(object):
    _module = 17
    _action = 7

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
        

class BattlePetModule(interface.BaseModule):
    decoder_map = {
        1: GetPetInfoDown, 
        3: SetPetDown, 
        4: SetPetSwapDown, 
        6: UpgradePetDown, 
        7: TrainingPetSkillDown, 
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

    def add_get_pet_info(self, callback):
        self.add_callback(1, callback)

    def add_set_pet(self, callback):
        self.add_callback(3, callback)

    def add_set_pet_swap(self, callback):
        self.add_callback(4, callback)

    def add_upgrade_pet(self, callback):
        self.add_callback(6, callback)

    def add_training_pet_skill(self, callback):
        self.add_callback(7, callback)
