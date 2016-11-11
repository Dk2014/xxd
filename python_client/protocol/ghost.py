#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

EQUIP_POS_POS1 = 0
EQUIP_POS_POS2 = 1
EQUIP_POS_POS3 = 2
EQUIP_POS_POS4 = 3


RESULT_TYPE_GHOST = 0
RESULT_TYPE_FRAGMENT = 1
RESULT_TYPE_FRUIT = 2


class InfoUp(object):
    _module = 9
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
        

class InfoDown(object):
    _module = 9
    _action = 0

    def __init__(self):
        pass
        self.train_times = 0
        self.flush_time = 0
        self.ghosts = []
        self.role_equip = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.train_times))
        buff.extend(struct.pack("<q", self.flush_time))
        buff.extend(struct.pack('<B', len(self.ghosts)))
        for item in self.ghosts:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.role_equip)))
        for item in self.role_equip:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.train_times = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.flush_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _ghosts_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_ghosts_size):
            obj = InfoDownGhosts()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.ghosts.append(obj)

        _role_equip_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_role_equip_size):
            obj = InfoDownRoleEquip()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.role_equip.append(obj)

    def size(self):
        size = 14
        for item in self.ghosts:
            size += item.size()
        for item in self.role_equip:
            size += item.size()
        return size


class InfoDownGhosts(object):
    def __init__(self):
        pass
        self.id = 0
        self.ghost_id = 0
        self.star = 0
        self.level = 0
        self.skill_level = 0
        self.exp = 0
        self.pos = 0
        self.health = 0
        self.attack = 0
        self.defence = 0
        self.speed = 0
        self.add_growth = 0
        self.relation_id = 0
        self.used = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.ghost_id))
        buff.extend(struct.pack("<b", self.star))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<h", self.skill_level))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.pos))
        buff.extend(struct.pack("<l", self.health))
        buff.extend(struct.pack("<l", self.attack))
        buff.extend(struct.pack("<l", self.defence))
        buff.extend(struct.pack("<l", self.speed))
        buff.extend(struct.pack("<h", self.add_growth))
        buff.extend(struct.pack("<q", self.relation_id))
        buff.extend(struct.pack("<?", self.used))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.ghost_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.star = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.health = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.attack = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.defence = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.speed = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.add_growth = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.relation_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.used = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 52
        

class InfoDownRoleEquip(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.already_use_ghost = False
        self.ghost_power = 0
        self.pos1_id = 0
        self.pos2_id = 0
        self.pos3_id = 0
        self.pos4_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<?", self.already_use_ghost))
        buff.extend(struct.pack("<l", self.ghost_power))
        buff.extend(struct.pack("<q", self.pos1_id))
        buff.extend(struct.pack("<q", self.pos2_id))
        buff.extend(struct.pack("<q", self.pos3_id))
        buff.extend(struct.pack("<q", self.pos4_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.already_use_ghost = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.ghost_power = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.pos1_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos2_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos3_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos4_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 38
        

class EquipUp(object):
    _module = 9
    _action = 1

    def __init__(self):
        pass
        self.from_id = 0
        self.role_id = 0
        self.to_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.from_id))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<B", self.to_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.from_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.to_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 12
        

class EquipDown(object):
    _module = 9
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
        return 0
        

class UnequipUp(object):
    _module = 9
    _action = 2

    def __init__(self):
        pass
        self.role_id = 0
        self.from_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.from_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.from_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 11
        

class UnequipDown(object):
    _module = 9
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
        

class SwapUp(object):
    _module = 9
    _action = 3

    def __init__(self):
        pass
        self.role_id = 0
        self.id_bag = 0
        self.id_equip = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.id_bag))
        buff.extend(struct.pack("<q", self.id_equip))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.id_bag = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.id_equip = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 19
        

class SwapDown(object):
    _module = 9
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
        

class EquipPosChangeUp(object):
    _module = 9
    _action = 5

    def __init__(self):
        pass
        self.role_id = 0
        self.from_id = 0
        self.to_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.from_id))
        buff.extend(struct.pack("<B", self.to_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.from_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.to_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 12
        

class EquipPosChangeDown(object):
    _module = 9
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
        

class TrainUp(object):
    _module = 9
    _action = 6

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class TrainDown(object):
    _module = 9
    _action = 6

    def __init__(self):
        pass
        self.add_exp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.add_exp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.add_exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class UpgradeUp(object):
    _module = 9
    _action = 7

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class UpgradeDown(object):
    _module = 9
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
        

class BaptizeUp(object):
    _module = 9
    _action = 8

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class BaptizeDown(object):
    _module = 9
    _action = 8

    def __init__(self):
        pass
        self.add_growth = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.add_growth))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.add_growth = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class ComposeUp(object):
    _module = 9
    _action = 11

    def __init__(self):
        pass
        self.ghost_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.ghost_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ghost_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class ComposeDown(object):
    _module = 9
    _action = 11

    def __init__(self):
        pass
        self.id = 0
        self.ghost_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.ghost_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.ghost_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 10
        

class TrainSkillUp(object):
    _module = 9
    _action = 12

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class TrainSkillDown(object):
    _module = 9
    _action = 12

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
        

class FlushAttrUp(object):
    _module = 9
    _action = 13

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class FlushAttrDown(object):
    _module = 9
    _action = 13

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
        

class RelationGhostUp(object):
    _module = 9
    _action = 14

    def __init__(self):
        pass
        self.master = 0
        self.slave = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.master))
        buff.extend(struct.pack("<q", self.slave))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.master = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.slave = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 18
        

class RelationGhostDown(object):
    _module = 9
    _action = 14

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
        

class GhostModule(interface.BaseModule):
    decoder_map = {
        0: InfoDown, 
        1: EquipDown, 
        2: UnequipDown, 
        3: SwapDown, 
        5: EquipPosChangeDown, 
        6: TrainDown, 
        7: UpgradeDown, 
        8: BaptizeDown, 
        11: ComposeDown, 
        12: TrainSkillDown, 
        13: FlushAttrDown, 
        14: RelationGhostDown, 
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
        self.add_callback(0, callback)

    def add_equip(self, callback):
        self.add_callback(1, callback)

    def add_unequip(self, callback):
        self.add_callback(2, callback)

    def add_swap(self, callback):
        self.add_callback(3, callback)

    def add_equip_pos_change(self, callback):
        self.add_callback(5, callback)

    def add_train(self, callback):
        self.add_callback(6, callback)

    def add_upgrade(self, callback):
        self.add_callback(7, callback)

    def add_baptize(self, callback):
        self.add_callback(8, callback)

    def add_compose(self, callback):
        self.add_callback(11, callback)

    def add_train_skill(self, callback):
        self.add_callback(12, callback)

    def add_flush_attr(self, callback):
        self.add_callback(13, callback)

    def add_relation_ghost(self, callback):
        self.add_callback(14, callback)
