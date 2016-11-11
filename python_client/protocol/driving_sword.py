#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

EXPLORING_MOUNTAIN_STATUS_UNEXPLORED = 0
EXPLORING_MOUNTAIN_STATUS_TREASURE_EMPTY = 1
EXPLORING_MOUNTAIN_STATUS_IN_GARRISON = 2
EXPLORING_MOUNTAIN_STATUS_BROKEN = 3


VISITING_MOUNTAIN_STATUS_UNWIN = 0
VISITING_MOUNTAIN_STATUS_WIN = 1
VISITING_MOUNTAIN_STATUS_AWARDED = 2


COMMON_EVENT_HOLE = 0
COMMON_EVENT_TELEPORT = 1
COMMON_EVENT_OBSTACLE = 2
COMMON_EVENT_UNKNOW_TELEPORT = 3


MOVING_DIRECTION_NORTH = 0
MOVING_DIRECTION_SOUTH = 1
MOVING_DIRECTION_WEST = 2
MOVING_DIRECTION_EAST = 3


class EventAreas(object):
    def __init__(self):
        pass
        self.common = []
        self.exploring_status = []
        self.visiting_status = []
        self.treasure_progress = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.common)))
        for item in self.common:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.exploring_status)))
        for item in self.exploring_status:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.visiting_status)))
        for item in self.visiting_status:
            buff.extend(item.encode())
        buff.extend(struct.pack('<B', len(self.treasure_progress)))
        for item in self.treasure_progress:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _common_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_common_size):
            obj = EventAreasCommon()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.common.append(obj)

        _exploring_status_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_exploring_status_size):
            obj = EventAreasExploringStatus()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.exploring_status.append(obj)

        _visiting_status_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_visiting_status_size):
            obj = EventAreasVisitingStatus()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.visiting_status.append(obj)

        _treasure_progress_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_treasure_progress_size):
            obj = EventAreasTreasureProgress()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.treasure_progress.append(obj)

    def size(self):
        size = 4
        for item in self.common:
            size += item.size()
        for item in self.exploring_status:
            size += item.size()
        for item in self.visiting_status:
            size += item.size()
        for item in self.treasure_progress:
            size += item.size()
        return size


class EventAreasCommon(object):
    def __init__(self):
        pass
        self.x = 0
        self.y = 0
        self.id = 0
        self.event = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.x))
        buff.extend(struct.pack("<B", self.y))
        buff.extend(struct.pack("<b", self.id))
        buff.extend(struct.pack("<B", self.event))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.event = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class EventAreasExploringStatus(object):
    def __init__(self):
        pass
        self.x = 0
        self.y = 0
        self.id = 0
        self.status = 0
        self.garrison_time = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.x))
        buff.extend(struct.pack("<B", self.y))
        buff.extend(struct.pack("<b", self.id))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<q", self.garrison_time))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.garrison_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 12
        

class EventAreasVisitingStatus(object):
    def __init__(self):
        pass
        self.x = 0
        self.y = 0
        self.id = 0
        self.status = 0
        self.pid = 0
        self.nick = ''
        self.role_id = 0
        self.level = 0
        self.fight_num = 0
        self.fashion_id = 0
        self.friendship_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.x))
        buff.extend(struct.pack("<B", self.y))
        buff.extend(struct.pack("<b", self.id))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<h", self.friendship_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 25
        size += len(self.nick)
        return size


class EventAreasTreasureProgress(object):
    def __init__(self):
        pass
        self.x = 0
        self.y = 0
        self.id = 0
        self.progress = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.x))
        buff.extend(struct.pack("<B", self.y))
        buff.extend(struct.pack("<b", self.id))
        buff.extend(struct.pack("<b", self.progress))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.progress = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class CloudMap(object):
    def __init__(self):
        pass
        self.shadows = bytearray()
        self.events = EventAreas()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.shadows)))
        buff.extend(self.shadows)
        buff.extend(self.events.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _shadows_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.shadows = raw_msg[idx:idx+_shadows_size]
        idx += _shadows_size

        self.events.decode(raw_msg[idx:])
        idx += self.events.size()

    def size(self):
        size = 2
        size += len(self.shadows)
        size += self.events.size()
        return size


class CloudMapInfoUp(object):
    _module = 30
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
        

class CloudMapInfoDown(object):
    _module = 30
    _action = 0

    def __init__(self):
        pass
        self.current_cloud = 0
        self.highest_cloud = 0
        self.current_x = 0
        self.current_y = 0
        self.allowed_action = 0
        self.daily_action_bought = 0
        self.map = CloudMap()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.current_cloud))
        buff.extend(struct.pack("<h", self.highest_cloud))
        buff.extend(struct.pack("<B", self.current_x))
        buff.extend(struct.pack("<B", self.current_y))
        buff.extend(struct.pack("<h", self.allowed_action))
        buff.extend(struct.pack("<b", self.daily_action_bought))
        buff.extend(self.map.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.current_cloud = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.highest_cloud = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.current_x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.current_y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.allowed_action = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.daily_action_bought = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.map.decode(raw_msg[idx:])
        idx += self.map.size()

    def size(self):
        size = 9
        size += self.map.size()
        return size


class CloudClimbUp(object):
    _module = 30
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
        

class CloudClimbDown(object):
    _module = 30
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
        

class CloudTeleportUp(object):
    _module = 30
    _action = 2

    def __init__(self):
        pass
        self.cloud = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.cloud))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.cloud = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class CloudTeleportDown(object):
    _module = 30
    _action = 2

    def __init__(self):
        pass
        self.map = CloudMap()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.map.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.map.decode(raw_msg[idx:])
        idx += self.map.size()

    def size(self):
        size = 0
        size += self.map.size()
        return size


class AreaTeleportUp(object):
    _module = 30
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
        return 2
        

class AreaTeleportDown(object):
    _module = 30
    _action = 3

    def __init__(self):
        pass
        self.events = EventAreas()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.events.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.events.decode(raw_msg[idx:])
        idx += self.events.size()

    def size(self):
        size = 0
        size += self.events.size()
        return size


class AreaMoveUp(object):
    _module = 30
    _action = 4

    def __init__(self):
        pass
        self.direction = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.direction))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.direction = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class AreaMoveDown(object):
    _module = 30
    _action = 4

    def __init__(self):
        pass
        self.events = EventAreas()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.events.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.events.decode(raw_msg[idx:])
        idx += self.events.size()

    def size(self):
        size = 0
        size += self.events.size()
        return size


class ExplorerStartBattleUp(object):
    _module = 30
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
        return 2
        

class ExplorerStartBattleDown(object):
    _module = 30
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
        

class ExplorerAwardUp(object):
    _module = 30
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
        return 2
        

class ExplorerAwardDown(object):
    _module = 30
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
        

class ExplorerGarrisonUp(object):
    _module = 30
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
        

class ExplorerGarrisonDown(object):
    _module = 30
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
        

class VisitMountainUp(object):
    _module = 30
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
        return 2
        

class VisitMountainDown(object):
    _module = 30
    _action = 8

    def __init__(self):
        pass
        self.status = 0
        self.pid = 0
        self.nick = ''
        self.role_id = 0
        self.level = 0
        self.fight_num = 0
        self.fashion_id = 0
        self.friendship_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.fight_num))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<h", self.friendship_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fight_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.friendship_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 22
        size += len(self.nick)
        return size


class VisiterStartBattleUp(object):
    _module = 30
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
        

class VisiterStartBattleDown(object):
    _module = 30
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
        return 0
        

class VisiterAwardUp(object):
    _module = 30
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
        

class VisiterAwardDown(object):
    _module = 30
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
        

class MountainTreasureOpenUp(object):
    _module = 30
    _action = 11

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
        

class MountainTreasureOpenDown(object):
    _module = 30
    _action = 11

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
        

class ListGarrisonsUp(object):
    _module = 30
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
        return 2
        

class ListGarrisonsDown(object):
    _module = 30
    _action = 12

    def __init__(self):
        pass
        self.garrisons = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.garrisons)))
        for item in self.garrisons:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _garrisons_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_garrisons_size):
            obj = ListGarrisonsDownGarrisons()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.garrisons.append(obj)

    def size(self):
        size = 1
        for item in self.garrisons:
            size += item.size()
        return size


class ListGarrisonsDownGarrisons(object):
    def __init__(self):
        pass
        self.role_id = 0
        self.garrison_time = 0
        self.awarded_time = 0
        self.cloud = 0
        self.event_id = 0
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<q", self.garrison_time))
        buff.extend(struct.pack("<q", self.awarded_time))
        buff.extend(struct.pack("<h", self.cloud))
        buff.extend(struct.pack("<b", self.event_id))
        buff.extend(struct.pack("<B", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.garrison_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.awarded_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.cloud = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.event_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 21
        

class AwardGarrisonUp(object):
    _module = 30
    _action = 13

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
        

class AwardGarrisonDown(object):
    _module = 30
    _action = 13

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
        

class EndGarrisonUp(object):
    _module = 30
    _action = 14

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
        

class EndGarrisonDown(object):
    _module = 30
    _action = 14

    def __init__(self):
        pass
        self.x = 0
        self.y = 0
        self.status = 0
        self.cloud_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.x))
        buff.extend(struct.pack("<B", self.y))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<h", self.cloud_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.x = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.y = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.cloud_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 5
        

class BuyAllowedActionUp(object):
    _module = 30
    _action = 15

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
        

class BuyAllowedActionDown(object):
    _module = 30
    _action = 15

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
        

class DrivingSwordModule(interface.BaseModule):
    decoder_map = {
        0: CloudMapInfoDown, 
        1: CloudClimbDown, 
        2: CloudTeleportDown, 
        3: AreaTeleportDown, 
        4: AreaMoveDown, 
        5: ExplorerStartBattleDown, 
        6: ExplorerAwardDown, 
        7: ExplorerGarrisonDown, 
        8: VisitMountainDown, 
        9: VisiterStartBattleDown, 
        10: VisiterAwardDown, 
        11: MountainTreasureOpenDown, 
        12: ListGarrisonsDown, 
        13: AwardGarrisonDown, 
        14: EndGarrisonDown, 
        15: BuyAllowedActionDown, 
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

    def add_cloud_map_info(self, callback):
        self.add_callback(0, callback)

    def add_cloud_climb(self, callback):
        self.add_callback(1, callback)

    def add_cloud_teleport(self, callback):
        self.add_callback(2, callback)

    def add_area_teleport(self, callback):
        self.add_callback(3, callback)

    def add_area_move(self, callback):
        self.add_callback(4, callback)

    def add_explorer_start_battle(self, callback):
        self.add_callback(5, callback)

    def add_explorer_award(self, callback):
        self.add_callback(6, callback)

    def add_explorer_garrison(self, callback):
        self.add_callback(7, callback)

    def add_visit_mountain(self, callback):
        self.add_callback(8, callback)

    def add_visiter_start_battle(self, callback):
        self.add_callback(9, callback)

    def add_visiter_award(self, callback):
        self.add_callback(10, callback)

    def add_mountain_treasure_open(self, callback):
        self.add_callback(11, callback)

    def add_list_garrisons(self, callback):
        self.add_callback(12, callback)

    def add_award_garrison(self, callback):
        self.add_callback(13, callback)

    def add_end_garrison(self, callback):
        self.add_callback(14, callback)

    def add_buy_allowed_action(self, callback):
        self.add_callback(15, callback)
