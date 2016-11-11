#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

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


FIGHTNUM_TYPE_ALL = 0
FIGHTNUM_TYPE_ROLE_LEVEL = 1
FIGHTNUM_TYPE_SKILL = 2
FIGHTNUM_TYPE_EQUIP = 3
FIGHTNUM_TYPE_GHOST = 4
FIGHTNUM_TYPE_REALM = 5
FIGHTNUM_TYPE_FASHION = 8
FIGHTNUM_TYPE_FRIENDSHIP = 9
FIGHTNUM_TYPE_TEAMSHIP = 10
FIGHTNUM_TYPE_CLIQUE_KONGFU = 11


class PlayerInfo(object):
    def __init__(self):
        pass
        self.openid = ''
        self.pid = 0
        self.name = ''
        self.best_segment = 0
        self.best_order = 0
        self.best_record_timestamp = 0
        self.fashion_id = 0
        self.roles = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack("<h", self.best_segment))
        buff.extend(struct.pack("<b", self.best_order))
        buff.extend(struct.pack("<q", self.best_record_timestamp))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack('<B', len(self.roles)))
        for item in self.roles:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        self.best_segment = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.best_order = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.best_record_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _roles_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_roles_size):
            obj = PlayerInfoRoles()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.roles.append(obj)

    def size(self):
        size = 26
        size += len(self.openid)
        size += len(self.name)
        for item in self.roles:
            size += item.size()
        return size


class PlayerInfoRoles(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.level = 0
        self.friendship_level = 0
        self.fight_num = 0
        self.is_deploy = False
        self.status = 0
        self.coop_id = 0
        self.attack = 0
        self.defence = 0
        self.health = 0
        self.speed = 0
        self.cultivation = 0
        self.sunder = 0
        self.hit_level = 0
        self.critical_level = 0
        self.sleep_level = 0
        self.dizziness_level = 0
        self.random_level = 0
        self.disable_skill_level = 0
        self.poisoning_level = 0
        self.block_level = 0
        self.destroy_level = 0
        self.critical_hurt_level = 0
        self.tenacity_level = 0
        self.dodge_level = 0
        self.equips = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<h", self.friendship_level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<?", self.is_deploy))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<h", self.coop_id))
        buff.extend(struct.pack("<l", self.attack))
        buff.extend(struct.pack("<l", self.defence))
        buff.extend(struct.pack("<l", self.health))
        buff.extend(struct.pack("<l", self.speed))
        buff.extend(struct.pack("<l", self.cultivation))
        buff.extend(struct.pack("<l", self.sunder))
        buff.extend(struct.pack("<l", self.hit_level))
        buff.extend(struct.pack("<l", self.critical_level))
        buff.extend(struct.pack("<l", self.sleep_level))
        buff.extend(struct.pack("<l", self.dizziness_level))
        buff.extend(struct.pack("<l", self.random_level))
        buff.extend(struct.pack("<l", self.disable_skill_level))
        buff.extend(struct.pack("<l", self.poisoning_level))
        buff.extend(struct.pack("<l", self.block_level))
        buff.extend(struct.pack("<l", self.destroy_level))
        buff.extend(struct.pack("<l", self.critical_hurt_level))
        buff.extend(struct.pack("<l", self.tenacity_level))
        buff.extend(struct.pack("<l", self.dodge_level))
        buff.extend(struct.pack('<B', len(self.equips)))
        for item in self.equips:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.is_deploy = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.coop_id = struct.unpack_from("<h", raw_msg, idx)[0]
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

        self.sunder = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.hit_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.critical_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.sleep_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.dizziness_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.random_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.disable_skill_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.poisoning_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.block_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.destroy_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.critical_hurt_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.tenacity_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.dodge_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _equips_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_equips_size):
            obj = PlayerInfoRolesEquips()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.equips.append(obj)

    def size(self):
        size = 86
        for item in self.equips:
            size += item.size()
        return size


class PlayerInfoRolesEquips(object):
    def __init__(self):
        pass
        self.pos = 0
        self.item_id = 0
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
        buff.extend(struct.pack("<b", self.pos))
        buff.extend(struct.pack("<h", self.item_id))
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

        self.pos = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
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
        return 50
        

class GetAllRolesUp(object):
    _module = 3
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
        

class GetAllRolesDown(object):
    _module = 3
    _action = 0

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
            obj = GetAllRolesDownRoles()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.roles.append(obj)

    def size(self):
        size = 1
        for item in self.roles:
            size += item.size()
        return size


class GetAllRolesDownRoles(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.level = 0
        self.exp = 0
        self.friendship_level = 0
        self.in_form = False
        self.status = 0
        self.coop_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.friendship_level))
        buff.extend(struct.pack("<?", self.in_form))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<h", self.coop_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.in_form = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.coop_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 17
        

class GetRoleInfoUp(object):
    _module = 3
    _action = 1

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
        

class GetRoleInfoDown(object):
    _module = 3
    _action = 1

    def __init__(self):
        pass
        self.role_id = 0
        self.level = 0
        self.exp = 0
        self.friendship_level = 0
        self.fight_num = 0
        self.status = 0
        self.coop_id = 0
        self.attack = 0
        self.defence = 0
        self.health = 0
        self.speed = 0
        self.cultivation = 0
        self.sunder = 0
        self.ghost_power = 0
        self.hit_level = 0
        self.critical_level = 0
        self.sleep_level = 0
        self.dizziness_level = 0
        self.random_level = 0
        self.disable_skill_level = 0
        self.poisoning_level = 0
        self.block_level = 0
        self.destroy_level = 0
        self.critical_hurt_level = 0
        self.tenacity_level = 0
        self.dodge_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.exp))
        buff.extend(struct.pack("<h", self.friendship_level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<h", self.coop_id))
        buff.extend(struct.pack("<l", self.attack))
        buff.extend(struct.pack("<l", self.defence))
        buff.extend(struct.pack("<l", self.health))
        buff.extend(struct.pack("<l", self.speed))
        buff.extend(struct.pack("<l", self.cultivation))
        buff.extend(struct.pack("<l", self.sunder))
        buff.extend(struct.pack("<l", self.ghost_power))
        buff.extend(struct.pack("<l", self.hit_level))
        buff.extend(struct.pack("<l", self.critical_level))
        buff.extend(struct.pack("<l", self.sleep_level))
        buff.extend(struct.pack("<l", self.dizziness_level))
        buff.extend(struct.pack("<l", self.random_level))
        buff.extend(struct.pack("<l", self.disable_skill_level))
        buff.extend(struct.pack("<l", self.poisoning_level))
        buff.extend(struct.pack("<l", self.block_level))
        buff.extend(struct.pack("<l", self.destroy_level))
        buff.extend(struct.pack("<l", self.critical_hurt_level))
        buff.extend(struct.pack("<l", self.tenacity_level))
        buff.extend(struct.pack("<l", self.dodge_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.exp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.coop_id = struct.unpack_from("<h", raw_msg, idx)[0]
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

        self.sunder = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.ghost_power = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.hit_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.critical_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.sleep_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.dizziness_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.random_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.disable_skill_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.poisoning_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.block_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.destroy_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.critical_hurt_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.tenacity_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.dodge_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 96
        

class GetPlayerInfoUp(object):
    _module = 3
    _action = 2

    def __init__(self):
        pass
        self.pid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class GetPlayerInfoDown(object):
    _module = 3
    _action = 2

    def __init__(self):
        pass
        self.player = PlayerInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.player.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player.decode(raw_msg[idx:])
        idx += self.player.size()

    def size(self):
        size = 0
        size += self.player.size()
        return size


class GetFightNumUp(object):
    _module = 3
    _action = 3

    def __init__(self):
        pass
        self.fight_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.fight_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fight_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class GetFightNumDown(object):
    _module = 3
    _action = 3

    def __init__(self):
        pass
        self.fight_nums = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.fight_nums)))
        for item in self.fight_nums:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _fight_nums_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_fight_nums_size):
            obj = GetFightNumDownFightNums()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.fight_nums.append(obj)

    def size(self):
        size = 1
        for item in self.fight_nums:
            size += item.size()
        return size


class GetFightNumDownFightNums(object):
    def __init__(self):
        pass
        self.fight_type = 0
        self.fight_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.fight_type))
        buff.extend(struct.pack("<l", self.fight_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.fight_type = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class GetPlayerInfoWithOpenidUp(object):
    _module = 3
    _action = 4

    def __init__(self):
        pass
        self.openid = ''
        self.game_server_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        buff.extend(struct.pack("<l", self.game_server_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

        self.game_server_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

    def size(self):
        size = 8
        size += len(self.openid)
        return size


class GetPlayerInfoWithOpenidDown(object):
    _module = 3
    _action = 4

    def __init__(self):
        pass
        self.player = PlayerInfo()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.player.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player.decode(raw_msg[idx:])
        idx += self.player.size()

    def size(self):
        size = 0
        size += self.player.size()
        return size


class LevelupRoleFriendshipUp(object):
    _module = 3
    _action = 5

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
        

class LevelupRoleFriendshipDown(object):
    _module = 3
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
        

class RecruitBuddyUp(object):
    _module = 3
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
        

class RecruitBuddyDown(object):
    _module = 3
    _action = 6

    def __init__(self):
        pass
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class GetRoleFightNumUp(object):
    _module = 3
    _action = 7

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
        

class GetRoleFightNumDown(object):
    _module = 3
    _action = 7

    def __init__(self):
        pass
        self.fight_nums = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.fight_nums)))
        for item in self.fight_nums:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _fight_nums_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_fight_nums_size):
            obj = GetRoleFightNumDownFightNums()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.fight_nums.append(obj)

    def size(self):
        size = 1
        for item in self.fight_nums:
            size += item.size()
        return size


class GetRoleFightNumDownFightNums(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.fight_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<l", self.fight_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class ChangeRoleStatusUp(object):
    _module = 3
    _action = 8

    def __init__(self):
        pass
        self.role_id = 0
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class ChangeRoleStatusDown(object):
    _module = 3
    _action = 8

    def __init__(self):
        pass
        self.result = False
        self.role_id = 0
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.result))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<b", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class GetInnRoleListUp(object):
    _module = 3
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
        

class GetInnRoleListDown(object):
    _module = 3
    _action = 9

    def __init__(self):
        pass
        self.role_list = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.role_list)))
        for item in self.role_list:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _role_list_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_role_list_size):
            obj = GetInnRoleListDownRoleList()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.role_list.append(obj)

    def size(self):
        size = 1
        for item in self.role_list:
            size += item.size()
        return size


class GetInnRoleListDownRoleList(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.friendship_level = 0
        self.fight_num = 0
        self.role_level = 0
        self.operate = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.friendship_level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<h", self.role_level))
        buff.extend(struct.pack("<b", self.operate))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.role_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.operate = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 10
        

class BuddyCoopUp(object):
    _module = 3
    _action = 10

    def __init__(self):
        pass
        self.coop_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.coop_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.coop_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class BuddyCoopDown(object):
    _module = 3
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
        return 0
        

class RoleModule(interface.BaseModule):
    decoder_map = {
        0: GetAllRolesDown, 
        1: GetRoleInfoDown, 
        2: GetPlayerInfoDown, 
        3: GetFightNumDown, 
        4: GetPlayerInfoWithOpenidDown, 
        5: LevelupRoleFriendshipDown, 
        6: RecruitBuddyDown, 
        7: GetRoleFightNumDown, 
        8: ChangeRoleStatusDown, 
        9: GetInnRoleListDown, 
        10: BuddyCoopDown, 
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

    def add_get_all_roles(self, callback):
        self.add_callback(0, callback)

    def add_get_role_info(self, callback):
        self.add_callback(1, callback)

    def add_get_player_info(self, callback):
        self.add_callback(2, callback)

    def add_get_fight_num(self, callback):
        self.add_callback(3, callback)

    def add_get_player_info_with_openid(self, callback):
        self.add_callback(4, callback)

    def add_levelup_role_friendship(self, callback):
        self.add_callback(5, callback)

    def add_recruit_buddy(self, callback):
        self.add_callback(6, callback)

    def add_get_role_fight_num(self, callback):
        self.add_callback(7, callback)

    def add_change_role_status(self, callback):
        self.add_callback(8, callback)

    def add_get_inn_role_list(self, callback):
        self.add_callback(9, callback)

    def add_buddy_coop(self, callback):
        self.add_callback(10, callback)
