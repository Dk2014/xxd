#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

DRAW_TYPE_COIN = 0
DRAW_TYPE_INGOT = 1


EQUIP_POS_POS1 = 0
EQUIP_POS_POS2 = 1
EQUIP_POS_POS3 = 2
EQUIP_POS_POS4 = 3
EQUIP_POS_POS5 = 4
EQUIP_POS_POS6 = 5
EQUIP_POS_POS7 = 6
EQUIP_POS_POS8 = 7
EQUIP_POS_POS9 = 8
EQUIP_POS_NUM = 9


BOX_A = 0
BOX_B = 1
BOX_C = 2
BOX_D = 3
BOX_E = 4


class InfoUp(object):
    _module = 10
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
    _module = 10
    _action = 0

    def __init__(self):
        pass
        self.sword_souls = []
        self.role_equip = []
        self.box_state = 0
        self.num = 0
        self.ingot_num = 0
        self.cd_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.sword_souls)))
        for item in self.sword_souls:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.role_equip)))
        for item in self.role_equip:
            buff.extend(item.encode())
        buff.extend(struct.pack("<b", self.box_state))
        buff.extend(struct.pack("<h", self.num))
        buff.extend(struct.pack("<q", self.ingot_num))
        buff.extend(struct.pack("<q", self.cd_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _sword_souls_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_sword_souls_size):
            obj = InfoDownSwordSouls()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.sword_souls.append(obj)

        _role_equip_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_role_equip_size):
            obj = InfoDownRoleEquip()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.role_equip.append(obj)

        self.box_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.ingot_num = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 21
        for item in self.sword_souls:
            size += item.size()
        for item in self.role_equip:
            size += item.size()
        return size


class InfoDownSwordSouls(object):
    def __init__(self):
        pass
        self.id = 0
        self.sword_soul_id = 0
        self.exp = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.sword_soul_id))
        buff.extend(struct.pack("<l", self.exp))
        buff.extend(struct.pack("<b", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.sword_soul_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.level = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 15
        

class InfoDownRoleEquip(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.pos1_id = 0
        self.pos2_id = 0
        self.pos3_id = 0
        self.pos4_id = 0
        self.pos5_id = 0
        self.pos6_id = 0
        self.pos7_id = 0
        self.pos8_id = 0
        self.pos9_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.pos1_id))
        buff.extend(struct.pack("<q", self.pos2_id))
        buff.extend(struct.pack("<q", self.pos3_id))
        buff.extend(struct.pack("<q", self.pos4_id))
        buff.extend(struct.pack("<q", self.pos5_id))
        buff.extend(struct.pack("<q", self.pos6_id))
        buff.extend(struct.pack("<q", self.pos7_id))
        buff.extend(struct.pack("<q", self.pos8_id))
        buff.extend(struct.pack("<q", self.pos9_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

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

        self.pos6_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos7_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos8_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.pos9_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 73
        

class DrawUp(object):
    _module = 10
    _action = 1

    def __init__(self):
        pass
        self.box = 0
        self.draw_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.box))
        buff.extend(struct.pack("<B", self.draw_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.box = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.draw_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class DrawDown(object):
    _module = 10
    _action = 1

    def __init__(self):
        pass
        self.id = 0
        self.sword_soul_id = 0
        self.coins = 0
        self.box_state = 0
        self.fragments = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.sword_soul_id))
        buff.extend(struct.pack("<q", self.coins))
        buff.extend(struct.pack("<b", self.box_state))
        buff.extend(struct.pack("<q", self.fragments))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.sword_soul_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.box_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.fragments = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 27
        

class UpgradeUp(object):
    _module = 10
    _action = 2

    def __init__(self):
        pass
        self.target_id = 0
        self.sword_souls = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.target_id))
        buff.extend(struct.pack('<B', len(self.sword_souls)))
        for item in self.sword_souls:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.target_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _sword_souls_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_sword_souls_size):
            obj = UpgradeUpSwordSouls()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.sword_souls.append(obj)

    def size(self):
        size = 11
        for item in self.sword_souls:
            size += item.size()
        return size


class UpgradeUpSwordSouls(object):
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
        return 8
        

class UpgradeDown(object):
    _module = 10
    _action = 2

    def __init__(self):
        pass
        self.id = 0
        self.sword_soul_id = 0
        self.exp = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.sword_soul_id))
        buff.extend(struct.pack("<l", self.exp))
        buff.extend(struct.pack("<b", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.sword_soul_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.level = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 15
        

class ExchangeUp(object):
    _module = 10
    _action = 3

    def __init__(self):
        pass
        self.sword_soul_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.sword_soul_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.sword_soul_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class ExchangeDown(object):
    _module = 10
    _action = 3

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
        return 8
        

class EquipUp(object):
    _module = 10
    _action = 4

    def __init__(self):
        pass
        self.from_id = 0
        self.role_id = 0
        self.equip_pos = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.from_id))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<B", self.equip_pos))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.from_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.equip_pos = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 12
        

class EquipDown(object):
    _module = 10
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
    _module = 10
    _action = 5

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
    _module = 10
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
    _module = 10
    _action = 6

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
    _module = 10
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
    _module = 10
    _action = 8

    def __init__(self):
        pass
        self.role_id = 0
        self.from_id = 0
        self.to_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.from_id))
        buff.extend(struct.pack("<q", self.to_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.from_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.to_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 19
        

class SwapDown(object):
    _module = 10
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
    _module = 10
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
    _module = 10
    _action = 9

    def __init__(self):
        pass
        self.is_full = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.is_full))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.is_full = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class EmptyPosNumUp(object):
    _module = 10
    _action = 10

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
        

class EmptyPosNumDown(object):
    _module = 10
    _action = 10

    def __init__(self):
        pass
        self.empty_pos_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.empty_pos_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.empty_pos_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class SwordSoulModule(interface.BaseModule):
    decoder_map = {
        0: InfoDown, 
        1: DrawDown, 
        2: UpgradeDown, 
        3: ExchangeDown, 
        4: EquipDown, 
        5: UnequipDown, 
        6: EquipPosChangeDown, 
        8: SwapDown, 
        9: IsBagFullDown, 
        10: EmptyPosNumDown, 
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

    def add_draw(self, callback):
        self.add_callback(1, callback)

    def add_upgrade(self, callback):
        self.add_callback(2, callback)

    def add_exchange(self, callback):
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

    def add_empty_pos_num(self, callback):
        self.add_callback(10, callback)
