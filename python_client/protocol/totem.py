#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

CALL_TYPE_BONE = 0
CALL_TYPE_HALLOW = 1
CALL_TYPE_INGOT = 2


EQUIP_POS_POS1 = 0
EQUIP_POS_POS2 = 1
EQUIP_POS_POS3 = 2
EQUIP_POS_POS4 = 3
EQUIP_POS_POS5 = 4


class InfoUp(object):
    _module = 31
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
    _module = 31
    _action = 0

    def __init__(self):
        pass
        self.totem = []
        self.rock_rune_num = 0
        self.jade_rune_num = 0
        self.call_num = 0
        self.pos1_id = 0
        self.pos2_id = 0
        self.pos3_id = 0
        self.pos4_id = 0
        self.pos5_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.totem)))
        for item in self.totem:
            buff.extend(item.encode())
        buff.extend(struct.pack("<l", self.rock_rune_num))
        buff.extend(struct.pack("<l", self.jade_rune_num))
        buff.extend(struct.pack("<h", self.call_num))
        buff.extend(struct.pack("<q", self.pos1_id))
        buff.extend(struct.pack("<q", self.pos2_id))
        buff.extend(struct.pack("<q", self.pos3_id))
        buff.extend(struct.pack("<q", self.pos4_id))
        buff.extend(struct.pack("<q", self.pos5_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _totem_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_totem_size):
            obj = InfoDownTotem()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.totem.append(obj)

        self.rock_rune_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.jade_rune_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.call_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.pos1_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos2_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos3_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos4_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos5_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 51
        for item in self.totem:
            size += item.size()
        return size


class InfoDownTotem(object):
    def __init__(self):
        pass
        self.id = 0
        self.totem_id = 0
        self.skill = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.totem_id))
        buff.extend(struct.pack("<h", self.skill))
        buff.extend(struct.pack("<b", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.totem_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.level = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 13
        

class CallTotemUp(object):
    _module = 31
    _action = 1

    def __init__(self):
        pass
        self.call_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.call_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.call_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class CallTotemDown(object):
    _module = 31
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
        

class UpgradeUp(object):
    _module = 31
    _action = 2

    def __init__(self):
        pass
        self.target_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.target_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.target_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class UpgradeDown(object):
    _module = 31
    _action = 2

    def __init__(self):
        pass
        self.ok = False
        self.id = 0
        self.skill = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.ok))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.skill))
        buff.extend(struct.pack("<b", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ok = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.skill = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.level = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 12
        

class DecomposeUp(object):
    _module = 31
    _action = 3

    def __init__(self):
        pass
        self.totem_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.totem_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.totem_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class DecomposeDown(object):
    _module = 31
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
        

class EquipUp(object):
    _module = 31
    _action = 4

    def __init__(self):
        pass
        self.totem_id = 0
        self.equip_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.totem_id))
        buff.extend(struct.pack("<B", self.equip_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.totem_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.equip_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 11
        

class EquipDown(object):
    _module = 31
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
        

class UnequipUp(object):
    _module = 31
    _action = 5

    def __init__(self):
        pass
        self.equip_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.equip_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.equip_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class UnequipDown(object):
    _module = 31
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
        

class EquipPosChangeUp(object):
    _module = 31
    _action = 6

    def __init__(self):
        pass
        self.from_pos = 0
        self.to_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.from_pos))
        buff.extend(struct.pack("<B", self.to_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.from_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.to_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class EquipPosChangeDown(object):
    _module = 31
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
        

class SwapUp(object):
    _module = 31
    _action = 8

    def __init__(self):
        pass
        self.equiped_id = 0
        self.inbag_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.equiped_id))
        buff.extend(struct.pack("<q", self.inbag_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.equiped_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.inbag_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 18
        

class SwapDown(object):
    _module = 31
    _action = 8

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
        

class IsBagFullUp(object):
    _module = 31
    _action = 9

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
        

class IsBagFullDown(object):
    _module = 31
    _action = 9

    def __init__(self):
        pass
        self.full = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.full))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.full = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class TotemModule(interface.BaseModule):
    decoder_map = {
        0: InfoDown, 
        1: CallTotemDown, 
        2: UpgradeDown, 
        3: DecomposeDown, 
        4: EquipDown, 
        5: UnequipDown, 
        6: EquipPosChangeDown, 
        8: SwapDown, 
        9: IsBagFullDown, 
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

    def add_call_totem(self, callback):
        self.add_callback(1, callback)

    def add_upgrade(self, callback):
        self.add_callback(2, callback)

    def add_decompose(self, callback):
        self.add_callback(3, callback)

    def add_equip(self, callback):
        self.add_callback(4, callback)

    def add_unequip(self, callback):
        self.add_callback(5, callback)

    def add_equip_pos_change(self, callback):
        self.add_callback(6, callback)

    def add_swap(self, callback):
        self.add_callback(8, callback)

    def add_is_bag_full(self, callback):
        self.add_callback(9, callback)
