#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

CHEAT_RESULT_SUCCESS = 0
CHEAT_RESULT_NO_DEPAND_SKILL = 1
CHEAT_RESULT_ALREADY_STUDY = 2
CHEAT_RESULT_CAN_NOT_STUDY_BEFORE = 3
CHEAT_RESULT_NOT_ROLE_SKILL = 4
CHEAT_RESULT_NOT_CHEAT_TYPE = 5
CHEAT_RESULT_ROLE_DOES_NOT_EXISTS = 6
CHEAT_RESULT_SKILL_NOT_MATCH_ROLE = 7
CHEAT_RESULT_LEVEL_NOT_REACHED = 8


class GetAllSkillsInfoUp(object):
    _module = 5
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
        

class GetAllSkillsInfoDown(object):
    _module = 5
    _action = 1

    def __init__(self):
        pass
        self.flush_time = 0
        self.roles = []
        self.skills = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.flush_time))
        buff.extend(struct.pack('<B', len(self.roles)))
        for item in self.roles:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.skills)))
        for item in self.skills:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.flush_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _roles_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_roles_size):
            obj = GetAllSkillsInfoDownRoles()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.roles.append(obj)

        _skills_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_skills_size):
            obj = GetAllSkillsInfoDownSkills()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.skills.append(obj)

    def size(self):
        size = 10
        for item in self.roles:
            size += item.size()
        for item in self.skills:
            size += item.size()
        return size


class GetAllSkillsInfoDownRoles(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.status = 0
        self.skill_id1 = 0
        self.skill_id2 = 0
        self.skill_id3 = 0
        self.skill_id4 = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<h", self.skill_id1))
        buff.extend(struct.pack("<h", self.skill_id2))
        buff.extend(struct.pack("<h", self.skill_id3))
        buff.extend(struct.pack("<h", self.skill_id4))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.skill_id1 = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill_id2 = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill_id3 = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill_id4 = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 10
        

class GetAllSkillsInfoDownSkills(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.skill_id = 0
        self.training_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.skill_id))
        buff.extend(struct.pack("<h", self.training_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.skill_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.training_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 5
        

class EquipSkillUp(object):
    _module = 5
    _action = 2

    def __init__(self):
        pass
        self.role_id = 0
        self.order_number = 0
        self.skill_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.order_number))
        buff.extend(struct.pack("<h", self.skill_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.order_number = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.skill_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class EquipSkillDown(object):
    _module = 5
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
        

class UnequipSkillUp(object):
    _module = 5
    _action = 3

    def __init__(self):
        pass
        self.role_id = 0
        self.order_number = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.order_number))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.order_number = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class UnequipSkillDown(object):
    _module = 5
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
        

class StudySkillByCheatUp(object):
    _module = 5
    _action = 4

    def __init__(self):
        pass
        self.role_id = 0
        self.item_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.item_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 5
        

class StudySkillByCheatDown(object):
    _module = 5
    _action = 4

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
        

class TrainSkillUp(object):
    _module = 5
    _action = 5

    def __init__(self):
        pass
        self.role_id = 0
        self.skill_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.skill_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.skill_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 5
        

class TrainSkillDown(object):
    _module = 5
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
        

class FlushSkillUp(object):
    _module = 5
    _action = 6

    def __init__(self):
        pass
        self.role_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class FlushSkillDown(object):
    _module = 5
    _action = 6

    def __init__(self):
        pass
        self.flush_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.flush_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.flush_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class SkillModule(interface.BaseModule):
    decoder_map = {
        1: GetAllSkillsInfoDown, 
        2: EquipSkillDown, 
        3: UnequipSkillDown, 
        4: StudySkillByCheatDown, 
        5: TrainSkillDown, 
        6: FlushSkillDown, 
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

    def add_get_all_skills_info(self, callback):
        self.add_callback(1, callback)

    def add_equip_skill(self, callback):
        self.add_callback(2, callback)

    def add_unequip_skill(self, callback):
        self.add_callback(3, callback)

    def add_study_skill_by_cheat(self, callback):
        self.add_callback(4, callback)

    def add_train_skill(self, callback):
        self.add_callback(5, callback)

    def add_flush_skill(self, callback):
        self.add_callback(6, callback)
