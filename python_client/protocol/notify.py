#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

BUFF_MODE_POWER = 0
BUFF_MODE_SPEED = 1
BUFF_MODE_ATTACK = 2
BUFF_MODE_DEFEND = 3
BUFF_MODE_HEALTH = 4
BUFF_MODE_DIZZINESS = 5
BUFF_MODE_POISONING = 6
BUFF_MODE_CLEAN_BAD = 7
BUFF_MODE_CLEAN_GOOD = 8
BUFF_MODE_REDUCE_HURT = 9
BUFF_MODE_RANDOM = 10
BUFF_MODE_BLOCK = 11
BUFF_MODE_BLOCK_LEVEL = 12
BUFF_MODE_DODGE_LEVEL = 13
BUFF_MODE_CRITIAL_LEVEL = 14
BUFF_MODE_HIT_LEVEL = 15
BUFF_MODE_HURT_ADD = 16
BUFF_MODE_MAX_HEALTH = 17
BUFF_MODE_KEEPER_REDUCE_HURT = 18
BUFF_MODE_ATTRACT_FIRE = 19
BUFF_MODE_DESTROY_LEVEL = 20
BUFF_MODE_TENACITY_LEVEL = 21
BUFF_MODE_SUNDER = 22
BUFF_MODE_SLEEP = 23
BUFF_MODE_DISABLE_SKILL = 24
BUFF_MODE_DIRECT_REDUCE_HURT = 25
BUFF_MODE_ABSORB_HURT = 26
BUFF_MODE_GHOST_POWER = 27
BUFF_MODE_PET_LIVE_ROUND = 28
BUFF_MODE_BUDDY_SKILL = 29
BUFF_MODE_CLEAR_ABSORB_HURT = 30
BUFF_MODE_SLEEP_LEVEL = 31
BUFF_MODE_DIZZINESS_LEVEL = 32
BUFF_MODE_RANDOM_LEVEL = 33
BUFF_MODE_DISABLE_SKILL_LEVEL = 34
BUFF_MODE_BUFF_POISONING_LEVEL = 35
BUFF_MODE_BUFF_RECOVER_BUDDY_SKILL = 36
BUFF_MODE_BUFF_MAKE_POWER_FULL = 37
BUFF_MODE_BUFF_DOGE = 38
BUFF_MODE_BUFF_HIT = 39
BUFF_MODE_BUFF_CRITIAL = 40
BUFF_MODE_BUFF_TENACITY = 41
BUFF_MODE_BUFF_TAKE_SUNSER = 42
BUFF_MODE_BUFF_DEFEND_PERSENT = 43
BUFF_MODE_BUFF_SUNDER_STATE = 44


class BufferInfo(object):
    def __init__(self):
        pass
        self.mode = 0
        self.keep = 0
        self.value = 0
        self.skill_id = 0
        self.max_override = 0
        self.override_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.mode))
        buff.extend(struct.pack("<b", self.keep))
        buff.extend(struct.pack("<l", self.value))
        buff.extend(struct.pack("<h", self.skill_id))
        buff.extend(struct.pack("<b", self.max_override))
        buff.extend(struct.pack("<b", self.override_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.mode = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.keep = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.value = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.skill_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.max_override = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.override_num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 10
        

ATTRIBUTE_NULL = 0
ATTRIBUTE_ATTACK = 1
ATTRIBUTE_DEFENCE = 2
ATTRIBUTE_HEALTH = 3
ATTRIBUTE_SPEED = 4
ATTRIBUTE_CULTIVATION = 5
ATTRIBUTE_HIT_LEVEL = 6
ATTRIBUTE_CRITICAL_LEVEL = 7
ATTRIBUTE_BLOCK_LEVEL = 8
ATTRIBUTE_DESTROY_LEVEL = 9
ATTRIBUTE_TENACITY_LEVEL = 10
ATTRIBUTE_DODGE_LEVEL = 11
ATTRIBUTE_NUM = 11


CHEST_TYPE_COIN_FREE = 0
CHEST_TYPE_INGOT_FREE = 1


class PlayerKeyChangedUp(object):
    _module = 8
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
        

class PlayerKeyChangedDown(object):
    _module = 8
    _action = 0

    def __init__(self):
        pass
        self.key = 0
        self.max_order = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.key))
        buff.extend(struct.pack("<b", self.max_order))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.key = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.max_order = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class MissionLevelLockChangedDown(object):
    _module = 8
    _action = 1

    def __init__(self):
        pass
        self.max_lock = 0
        self.award_lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.max_lock))
        buff.extend(struct.pack("<l", self.award_lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.max_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.award_lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 8
        

class RoleExpChangeDown(object):
    _module = 8
    _action = 2

    def __init__(self):
        pass
        self.role_id = 0
        self.add_exp = 0
        self.exp = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.add_exp))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.add_exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 19
        

class PhysicalChangeDown(object):
    _module = 8
    _action = 3

    def __init__(self):
        pass
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.value = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class MoneyChangeDown(object):
    _module = 8
    _action = 4

    def __init__(self):
        pass
        self.moneytype = 0
        self.value = 0
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.moneytype))
        buff.extend(struct.pack("<q", self.value))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.moneytype = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.value = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 17
        

class SkillAddDown(object):
    _module = 8
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
        return 3
        

class ItemChangeDown(object):
    _module = 8
    _action = 6

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
            obj = ItemChangeDownItems()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.items.append(obj)

    def size(self):
        size = 1
        for item in self.items:
            size += item.size()
        return size


class ItemChangeDownItems(object):
    def __init__(self):
        pass
        self.id = 0
        self.item_id = 0
        self.num = 0
        self.attack = 0
        self.defence = 0
        self.health = 0
        self.speed = 0
        self.cultivation = 0
        self.hit_level = 0
        self.critical_level = 0
        self.block_level = 0
        self.destroy_level = 0
        self.tenacity_level = 0
        self.dodge_level = 0
        self.refine_level = 0
        self.recast_attr = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<h", self.num))
        buff.extend(struct.pack("<l", self.attack))
        buff.extend(struct.pack("<l", self.defence))
        buff.extend(struct.pack("<l", self.health))
        buff.extend(struct.pack("<l", self.speed))
        buff.extend(struct.pack("<l", self.cultivation))
        buff.extend(struct.pack("<l", self.hit_level))
        buff.extend(struct.pack("<l", self.critical_level))
        buff.extend(struct.pack("<l", self.block_level))
        buff.extend(struct.pack("<l", self.destroy_level))
        buff.extend(struct.pack("<l", self.tenacity_level))
        buff.extend(struct.pack("<l", self.dodge_level))
        buff.extend(struct.pack("<h", self.refine_level))
        buff.extend(struct.pack("<B", self.recast_attr))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.attack = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.defence = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.health = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.speed = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.cultivation = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.hit_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.critical_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.block_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.destroy_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.tenacity_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.dodge_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.refine_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.recast_attr = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 59
        

class RoleBattleStatusChangeDown(object):
    _module = 8
    _action = 7

    def __init__(self):
        pass
        self.roles = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.roles)))
        for item in self.roles:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _roles_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_roles_size):
            obj = RoleBattleStatusChangeDownRoles()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.roles.append(obj)

    def size(self):
        size = 1
        for item in self.roles:
            size += item.size()
        return size


class RoleBattleStatusChangeDownRoles(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.health = 0
        self.buffs = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<l", self.health))
        buff.extend(struct.pack('<B', len(self.buffs)))
        for item in self.buffs:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.health = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _buffs_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_buffs_size):
            obj = RoleBattleStatusChangeDownRolesBuffs()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.buffs.append(obj)

    def size(self):
        size = 6
        for item in self.buffs:
            size += item.size()
        return size


class RoleBattleStatusChangeDownRolesBuffs(object):
    def __init__(self):
        pass
        self.buffer = BufferInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.buffer.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.buffer.decode(raw_msg[idx:])
        idx += self.buffer.size()

    def size(self):
        size = 0
        size += self.buffer.size()
        return size


class NewMailDown(object):
    _module = 8
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
        

class HeartChangeDown(object):
    _module = 8
    _action = 9

    def __init__(self):
        pass
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.value = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class QuestChangeDown(object):
    _module = 8
    _action = 10

    def __init__(self):
        pass
        self.quest_id = 0
        self.state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.quest_id))
        buff.extend(struct.pack("<b", self.state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class TownLockChangeDown(object):
    _module = 8
    _action = 11

    def __init__(self):
        pass
        self.lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class ChatDown(object):
    _module = 8
    _action = 12

    def __init__(self):
        pass
        self.pid = 0
        self.role_id = 0
        self.nickname = ''
        self.level = 0
        self.fight_num = 0
        self.message = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack('<H', len(self.message)))
        buff.extend(self.message)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _message_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.message = str(raw_msg[idx:idx+_message_size])
        idx += _message_size

    def size(self):
        size = 19
        size += len(self.nickname)
        size += len(self.message)
        return size


class FuncKeyChangeDown(object):
    _module = 8
    _action = 13

    def __init__(self):
        pass
        self.func_key = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.func_key))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.func_key = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class ItemRecastStateRebuildDown(object):
    _module = 8
    _action = 14

    def __init__(self):
        pass
        self.id = 0
        self.selected_attr = 0
        self.attrs = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<B", self.selected_attr))
        buff.extend(struct.pack('<B', len(self.attrs)))
        for item in self.attrs:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.selected_attr = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _attrs_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_attrs_size):
            obj = ItemRecastStateRebuildDownAttrs()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.attrs.append(obj)

    def size(self):
        size = 10
        for item in self.attrs:
            size += item.size()
        return size


class ItemRecastStateRebuildDownAttrs(object):
    def __init__(self):
        pass
        self.attr = 0
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.attr))
        buff.extend(struct.pack("<l", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.attr = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.value = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class SendAnnouncementDown(object):
    _module = 8
    _action = 15

    def __init__(self):
        pass
        self.id = 0
        self.tpl_id = 0
        self.expire_time = 0
        self.parameters = ''
        self.content = ''
        self.spacing_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<l", self.tpl_id))
        buff.extend(struct.pack("<q", self.expire_time))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        buff.extend(struct.pack('<H', len(self.content)))
        buff.extend(self.content)
        buff.extend(struct.pack("<l", self.spacing_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.tpl_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.expire_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

        _content_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.content = str(raw_msg[idx:idx+_content_size])
        idx += _content_size

        self.spacing_time = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 28
        size += len(self.parameters)
        size += len(self.content)
        return size


class VipLevelChangeDown(object):
    _module = 8
    _action = 16

    def __init__(self):
        pass
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class NotifyNewBuddyDown(object):
    _module = 8
    _action = 17

    def __init__(self):
        pass
        self.role_id = 0
        self.role_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.role_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.role_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 3
        

class HardLevelLockChangedDown(object):
    _module = 8
    _action = 18

    def __init__(self):
        pass
        self.lock = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.lock))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.lock = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class SendSwordSoulDrawNumChangeDown(object):
    _module = 8
    _action = 19

    def __init__(self):
        pass
        self.num = 0
        self.cd_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.num))
        buff.extend(struct.pack("<q", self.cd_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.cd_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class SendHaveNewGhostDown(object):
    _module = 8
    _action = 21

    def __init__(self):
        pass
        self.player_ghost_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_ghost_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_ghost_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class SendHeartRecoverTimeDown(object):
    _module = 8
    _action = 22

    def __init__(self):
        pass
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class SendGlobalMailDown(object):
    _module = 8
    _action = 23

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
        

class SendPhysicalRecoverTimeDown(object):
    _module = 8
    _action = 24

    def __init__(self):
        pass
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class SendFashionChangeDown(object):
    _module = 8
    _action = 25

    def __init__(self):
        pass
        self.fashion_id = 0
        self.expire_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<q", self.expire_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.expire_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class TransErrorDown(object):
    _module = 8
    _action = 26

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
        

class SendEventCenterChangeDown(object):
    _module = 8
    _action = 27

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
        

class MeditationStateDown(object):
    _module = 8
    _action = 29

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
        

class DeleteAnnouncementDown(object):
    _module = 8
    _action = 31

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
        

class SendHaveNewPetDown(object):
    _module = 8
    _action = 32

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
        return 4
        

class SendLogoutDown(object):
    _module = 8
    _action = 33

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
        

class FameChangeDown(object):
    _module = 8
    _action = 34

    def __init__(self):
        pass
        self.fame = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.fame))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fame = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class NotifyMonthCardOpenDown(object):
    _module = 8
    _action = 36

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
        

class NotifyMonthCardRenewalDown(object):
    _module = 8
    _action = 37

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
        

class NotifyNewTotemDown(object):
    _module = 8
    _action = 38

    def __init__(self):
        pass
        self.id = 0
        self.totem_id = 0
        self.skill = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.totem_id))
        buff.extend(struct.pack("<h", self.skill))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.totem_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.skill = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 12
        

class NotifyRuneChangeDown(object):
    _module = 8
    _action = 39

    def __init__(self):
        pass
        self.rock_rune_num = 0
        self.jade_rune_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.rock_rune_num))
        buff.extend(struct.pack("<l", self.jade_rune_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.rock_rune_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.jade_rune_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 8
        

class NotifyMonthCardChangeDown(object):
    _module = 8
    _action = 40

    def __init__(self):
        pass
        self.expire_at = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.expire_at))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.expire_at = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class NotifyFirstRechargeStateDown(object):
    _module = 8
    _action = 41

    def __init__(self):
        pass
        self.never_recharge = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.never_recharge))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.never_recharge = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyModule(interface.BaseModule):
    decoder_map = {
        0: PlayerKeyChangedDown, 
        1: MissionLevelLockChangedDown, 
        2: RoleExpChangeDown, 
        3: PhysicalChangeDown, 
        4: MoneyChangeDown, 
        5: SkillAddDown, 
        6: ItemChangeDown, 
        7: RoleBattleStatusChangeDown, 
        8: NewMailDown, 
        9: HeartChangeDown, 
        10: QuestChangeDown, 
        11: TownLockChangeDown, 
        12: ChatDown, 
        13: FuncKeyChangeDown, 
        14: ItemRecastStateRebuildDown, 
        15: SendAnnouncementDown, 
        16: VipLevelChangeDown, 
        17: NotifyNewBuddyDown, 
        18: HardLevelLockChangedDown, 
        19: SendSwordSoulDrawNumChangeDown, 
        21: SendHaveNewGhostDown, 
        22: SendHeartRecoverTimeDown, 
        23: SendGlobalMailDown, 
        24: SendPhysicalRecoverTimeDown, 
        25: SendFashionChangeDown, 
        26: TransErrorDown, 
        27: SendEventCenterChangeDown, 
        29: MeditationStateDown, 
        31: DeleteAnnouncementDown, 
        32: SendHaveNewPetDown, 
        33: SendLogoutDown, 
        34: FameChangeDown, 
        36: NotifyMonthCardOpenDown, 
        37: NotifyMonthCardRenewalDown, 
        38: NotifyNewTotemDown, 
        39: NotifyRuneChangeDown, 
        40: NotifyMonthCardChangeDown, 
        41: NotifyFirstRechargeStateDown, 
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

    def add_player_key_changed(self, callback):
        self.add_callback(0, callback)

    def add_mission_level_lock_changed(self, callback):
        self.add_callback(1, callback)

    def add_role_exp_change(self, callback):
        self.add_callback(2, callback)

    def add_physical_change(self, callback):
        self.add_callback(3, callback)

    def add_money_change(self, callback):
        self.add_callback(4, callback)

    def add_skill_add(self, callback):
        self.add_callback(5, callback)

    def add_item_change(self, callback):
        self.add_callback(6, callback)

    def add_role_battle_status_change(self, callback):
        self.add_callback(7, callback)

    def add_new_mail(self, callback):
        self.add_callback(8, callback)

    def add_heart_change(self, callback):
        self.add_callback(9, callback)

    def add_quest_change(self, callback):
        self.add_callback(10, callback)

    def add_townLock_change(self, callback):
        self.add_callback(11, callback)

    def add_chat(self, callback):
        self.add_callback(12, callback)

    def add_func_key_change(self, callback):
        self.add_callback(13, callback)

    def add_item_recast_state_rebuild(self, callback):
        self.add_callback(14, callback)

    def add_send_announcement(self, callback):
        self.add_callback(15, callback)

    def add_vip_level_change(self, callback):
        self.add_callback(16, callback)

    def add_notify_new_buddy(self, callback):
        self.add_callback(17, callback)

    def add_hard_level_lock_changed(self, callback):
        self.add_callback(18, callback)

    def add_send_sword_soul_draw_num_change(self, callback):
        self.add_callback(19, callback)

    def add_send_have_new_ghost(self, callback):
        self.add_callback(21, callback)

    def add_send_heart_recover_time(self, callback):
        self.add_callback(22, callback)

    def add_send_global_mail(self, callback):
        self.add_callback(23, callback)

    def add_send_physical_recover_time(self, callback):
        self.add_callback(24, callback)

    def add_send_fashion_change(self, callback):
        self.add_callback(25, callback)

    def add_trans_error(self, callback):
        self.add_callback(26, callback)

    def add_send_event_center_change(self, callback):
        self.add_callback(27, callback)

    def add_meditation_state(self, callback):
        self.add_callback(29, callback)

    def add_delete_announcement(self, callback):
        self.add_callback(31, callback)

    def add_send_have_new_pet(self, callback):
        self.add_callback(32, callback)

    def add_send_logout(self, callback):
        self.add_callback(33, callback)

    def add_fame_change(self, callback):
        self.add_callback(34, callback)

    def add_notify_month_card_open(self, callback):
        self.add_callback(36, callback)

    def add_notify_month_card_renewal(self, callback):
        self.add_callback(37, callback)

    def add_notify_new_totem(self, callback):
        self.add_callback(38, callback)

    def add_notify_rune_change(self, callback):
        self.add_callback(39, callback)

    def add_notify_month_card_change(self, callback):
        self.add_callback(40, callback)

    def add_notify_first_recharge_state(self, callback):
        self.add_callback(41, callback)
