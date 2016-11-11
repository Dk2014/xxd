#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

CREATE_CLIQUE_RESULT_SUCCESS = 0
CREATE_CLIQUE_RESULT_DUP_NAME = 1
CREATE_CLIQUE_RESULT_HAVE_CLIQUE = 2


APPLY_CLIQUE_RESULT_SUCCESS = 0
APPLY_CLIQUE_RESULT_ALREADY_JOIN = 1
APPLY_CLIQUE_RESULT_NOT_EXIST = 2
APPLY_CLIQUE_RESULT_REFUSE = 3
APPLY_CLIQUE_RESULT_TOO_MUCH_APPLY = 4


CANCEL_APPLY_CLIQUE_RESULT_SUCCESS = 0
CANCEL_APPLY_CLIQUE_RESULT_EXPIRED = 1
CANCEL_APPLY_CLIQUE_RESULT_NOT_EXIST = 2


PROCESS_JOIN_APPLY_RESULT_SUCCESS = 0
PROCESS_JOIN_APPLY_RESULT_EXPIRED = 1
PROCESS_JOIN_APPLY_RESULT_NO_ROOM = 2
PROCESS_JOIN_APPLY_RESULT_NO_PERMISSION = 3
PROCESS_JOIN_APPLY_RESULT_CANCEL_APPLY = 4


MANGE_MEMBER_ACTION_SET_OWNER = 0
MANGE_MEMBER_ACTION_SET_MANGER = 1
MANGE_MEMBER_ACTION_SET_MEMBER = 2
MANGE_MEMBER_ACTION_KICKOFF = 3


MANGE_MEMBER_RESULT_SUCCESS = 0
MANGE_MEMBER_RESULT_NOT_EXIST = 1
MANGE_MEMBER_RESULT_NO_PERMISSION = 2


CLIQUE_OPERA_ERROR_SUCCESS = 0
CLIQUE_OPERA_ERROR_CLIQUE_NOT_EXIST = 1
CLIQUE_OPERA_ERROR_NO_PERMISSION = 2
CLIQUE_OPERA_ERROR_MEMBER_NOT_EXIST = 3
CLIQUE_OPERA_ERROR_ALREADY_JOIN = 4


NOTIFY_LEAVE_CLIQUE_REASON_KICKOFF = 0
NOTIFY_LEAVE_CLIQUE_REASON_COLLAPSE = 1
NOTIFY_LEAVE_CLIQUE_REASON_LEAVE = 2


NOTIFY_JOINCLIQUE_FAILED_REASON_REFUSE = 0
NOTIFY_JOINCLIQUE_FAILED_REASON_EXPIRED = 1
NOTIFY_JOINCLIQUE_FAILED_REASON_NOROOM = 2


CLIQUE_RECUITMENT_RESULT_SUCCESS = 0
CLIQUE_RECUITMENT_RESULT_NO_PERMISSION = 1
CLIQUE_RECUITMENT_RESULT_CD = 2


class Player(object):
    def __init__(self):
        pass
        self.player_id = 0
        self.nickname = ''
        self.role_id = 0
        self.at_x = 0
        self.at_y = 0
        self.fashion_id = 0
        self.in_meditation_state = False
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.at_x))
        buff.extend(struct.pack("<h", self.at_y))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<?", self.in_meditation_state))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.at_x = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.at_y = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.in_meditation_state = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 20
        size += len(self.nickname)
        return size


class CreateCliqueUp(object):
    _module = 33
    _action = 0

    def __init__(self):
        pass
        self.name = ''
        self.announce = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

    def size(self):
        size = 6
        size += len(self.name)
        size += len(self.announce)
        return size


class CreateCliqueDown(object):
    _module = 33
    _action = 0

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
        

class ApplyJoinCliqueUp(object):
    _module = 33
    _action = 1

    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class ApplyJoinCliqueDown(object):
    _module = 33
    _action = 1

    def __init__(self):
        pass
        self.clique_id = 0
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        buff.extend(struct.pack("<B", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 9
        

class CancelApplyCliqueUp(object):
    _module = 33
    _action = 2

    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CancelApplyCliqueDown(object):
    _module = 33
    _action = 2

    def __init__(self):
        pass
        self.result = 0
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 9
        

class ProcessJoinApplyUp(object):
    _module = 33
    _action = 3

    def __init__(self):
        pass
        self.agree = False
        self.pidlist = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.agree))
        buff.extend(struct.pack('<B', len(self.pidlist)))
        for item in self.pidlist:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.agree = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        _pidlist_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_pidlist_size):
            obj = ProcessJoinApplyUpPidlist()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.pidlist.append(obj)

    def size(self):
        size = 4
        for item in self.pidlist:
            size += item.size()
        return size


class ProcessJoinApplyUpPidlist(object):
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
        return 8
        

class ProcessJoinApplyDown(object):
    _module = 33
    _action = 3

    def __init__(self):
        pass
        self.applylist = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.applylist)))
        for item in self.applylist:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _applylist_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_applylist_size):
            obj = ProcessJoinApplyDownApplylist()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.applylist.append(obj)

    def size(self):
        size = 1
        for item in self.applylist:
            size += item.size()
        return size


class ProcessJoinApplyDownApplylist(object):
    def __init__(self):
        pass
        self.pid = 0
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 9
        

class ElectOwnerUp(object):
    _module = 33
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
        return 2
        

class ElectOwnerDown(object):
    _module = 33
    _action = 4

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class MangeMemberUp(object):
    _module = 33
    _action = 5

    def __init__(self):
        pass
        self.pid = 0
        self.action = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.action))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.action = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 11
        

class MangeMemberDown(object):
    _module = 33
    _action = 5

    def __init__(self):
        pass
        self.action = 0
        self.pid = 0
        self.result = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.action = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 10
        

class DestoryCliqueUp(object):
    _module = 33
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
        

class DestoryCliqueDown(object):
    _module = 33
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
        

class UpdateAnnounceUp(object):
    _module = 33
    _action = 7

    def __init__(self):
        pass
        self.content = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.content)))
        buff.extend(self.content)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _content_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.content = str(raw_msg[idx:idx+_content_size])
        idx += _content_size

    def size(self):
        size = 4
        size += len(self.content)
        return size


class UpdateAnnounceDown(object):
    _module = 33
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
        

class LeaveCliqueUp(object):
    _module = 33
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
        

class LeaveCliqueDown(object):
    _module = 33
    _action = 8

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class ListCliqueUp(object):
    _module = 33
    _action = 9

    def __init__(self):
        pass
        self.offset = 0
        self.limit = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.offset))
        buff.extend(struct.pack("<h", self.limit))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.offset = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.limit = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class ListCliqueDown(object):
    _module = 33
    _action = 9

    def __init__(self):
        pass
        self.applied_cliques = []
        self.total = 0
        self.cliques = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.applied_cliques)))
        for item in self.applied_cliques:
            buff.extend(item.encode())
        buff.extend(struct.pack("<h", self.total))
        buff.extend(struct.pack('<H', len(self.cliques)))
        for item in self.cliques:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _applied_cliques_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_applied_cliques_size):
            obj = ListCliqueDownAppliedCliques()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.applied_cliques.append(obj)

        self.total = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _cliques_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_cliques_size):
            obj = ListCliqueDownCliques()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.cliques.append(obj)

    def size(self):
        size = 5
        for item in self.applied_cliques:
            size += item.size()
        for item in self.cliques:
            size += item.size()
        return size


class ListCliqueDownAppliedCliques(object):
    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class ListCliqueDownCliques(object):
    def __init__(self):
        pass
        self.id = 0
        self.name = ''
        self.level = 0
        self.member_num = 0
        self.owner_nick = ''
        self.owner_pid = 0
        self.announce = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<h", self.member_num))
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.member_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

    def size(self):
        size = 26
        size += len(self.name)
        size += len(self.owner_nick)
        size += len(self.announce)
        return size


class CliquePublicInfoUp(object):
    _module = 33
    _action = 10

    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliquePublicInfoDown(object):
    _module = 33
    _action = 10

    def __init__(self):
        pass
        self.clique_id = 0
        self.exist = False
        self.name = ''
        self.owner_nick = ''
        self.owner_pid = 0
        self.member_num = 0
        self.level = 0
        self.announce = ''
        self.center_building_level = 0
        self.temple_building_level = 0
        self.bank_building_level = 0
        self.health_building_level = 0
        self.attack_building_level = 0
        self.defense_building_level = 0
        self.applied_cliques = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        buff.extend(struct.pack("<?", self.exist))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack("<h", self.member_num))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        buff.extend(struct.pack("<h", self.center_building_level))
        buff.extend(struct.pack("<h", self.temple_building_level))
        buff.extend(struct.pack("<h", self.bank_building_level))
        buff.extend(struct.pack("<h", self.health_building_level))
        buff.extend(struct.pack("<h", self.attack_building_level))
        buff.extend(struct.pack("<h", self.defense_building_level))
        buff.extend(struct.pack('<B', len(self.applied_cliques)))
        for item in self.applied_cliques:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.exist = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.member_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

        self.center_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.temple_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.bank_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.health_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.attack_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.defense_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _applied_cliques_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_applied_cliques_size):
            obj = CliquePublicInfoDownAppliedCliques()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.applied_cliques.append(obj)

    def size(self):
        size = 40
        size += len(self.name)
        size += len(self.owner_nick)
        size += len(self.announce)
        for item in self.applied_cliques:
            size += item.size()
        return size


class CliquePublicInfoDownAppliedCliques(object):
    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class CliqueInfoUp(object):
    _module = 33
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
        

class CliqueInfoDown(object):
    _module = 33
    _action = 11

    def __init__(self):
        pass
        self.clique_id = 0
        self.name = ''
        self.announce = ''
        self.total_donate_coins = 0
        self.contrib = 0
        self.owner_login_time = 0
        self.owner_pid = 0
        self.manger_pid1 = 0
        self.manger_pid2 = 0
        self.center_building_coins = 0
        self.temple_building_coins = 0
        self.bank_building_coins = 0
        self.health_building_coins = 0
        self.attack_building_coins = 0
        self.defense_building_coins = 0
        self.center_building_level = 0
        self.temple_building_level = 0
        self.bank_building_level = 0
        self.health_building_level = 0
        self.attack_building_level = 0
        self.defense_building_level = 0
        self.recruit_timestamp = 0
        self.members = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        buff.extend(struct.pack("<q", self.total_donate_coins))
        buff.extend(struct.pack("<q", self.contrib))
        buff.extend(struct.pack("<q", self.owner_login_time))
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack("<q", self.manger_pid1))
        buff.extend(struct.pack("<q", self.manger_pid2))
        buff.extend(struct.pack("<q", self.center_building_coins))
        buff.extend(struct.pack("<q", self.temple_building_coins))
        buff.extend(struct.pack("<q", self.bank_building_coins))
        buff.extend(struct.pack("<q", self.health_building_coins))
        buff.extend(struct.pack("<q", self.attack_building_coins))
        buff.extend(struct.pack("<q", self.defense_building_coins))
        buff.extend(struct.pack("<h", self.center_building_level))
        buff.extend(struct.pack("<h", self.temple_building_level))
        buff.extend(struct.pack("<h", self.bank_building_level))
        buff.extend(struct.pack("<h", self.health_building_level))
        buff.extend(struct.pack("<h", self.attack_building_level))
        buff.extend(struct.pack("<h", self.defense_building_level))
        buff.extend(struct.pack("<q", self.recruit_timestamp))
        buff.extend(struct.pack('<H', len(self.members)))
        for item in self.members:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

        self.total_donate_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.contrib = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.owner_login_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.manger_pid1 = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.manger_pid2 = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.center_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.temple_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.bank_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.health_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.attack_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.defense_building_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.center_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.temple_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.bank_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.health_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.attack_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.defense_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.recruit_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _members_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_members_size):
            obj = CliqueInfoDownMembers()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.members.append(obj)

    def size(self):
        size = 130
        size += len(self.name)
        size += len(self.announce)
        for item in self.members:
            size += item.size()
        return size


class CliqueInfoDownMembers(object):
    def __init__(self):
        pass
        self.pid = 0
        self.role_id = 0
        self.level = 0
        self.nick = ''
        self.contrib = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<q", self.contrib))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.contrib = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 21
        size += len(self.nick)
        return size


class ListApplyUp(object):
    _module = 33
    _action = 12

    def __init__(self):
        pass
        self.offset = 0
        self.limit = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.offset))
        buff.extend(struct.pack("<h", self.limit))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.offset = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.limit = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class ListApplyDown(object):
    _module = 33
    _action = 12

    def __init__(self):
        pass
        self.auto_audit = False
        self.level = 0
        self.players = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.auto_audit))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack('<H', len(self.players)))
        for item in self.players:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.auto_audit = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _players_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_players_size):
            obj = ListApplyDownPlayers()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.players.append(obj)

    def size(self):
        size = 5
        for item in self.players:
            size += item.size()
        return size


class ListApplyDownPlayers(object):
    def __init__(self):
        pass
        self.pid = 0
        self.nick = ''
        self.level = 0
        self.arena_rank = 0
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.arena_rank))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.arena_rank = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 28
        size += len(self.nick)
        return size


class OperaErrorNotifyDown(object):
    _module = 33
    _action = 13

    def __init__(self):
        pass
        self.resutl = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.resutl))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.resutl = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class EnterClubhouseUp(object):
    _module = 33
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
        return 2
        

class EnterClubhouseDown(object):
    _module = 33
    _action = 14

    def __init__(self):
        pass
        self.ok = False
        self.player = Player()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.ok))
        buff.extend(self.player.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ok = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.player.decode(raw_msg[idx:])
        idx += self.player.size()

    def size(self):
        size = 1
        size += self.player.size()
        return size


class LeaveClubhouseUp(object):
    _module = 33
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
        

class LeaveClubhouseDown(object):
    _module = 33
    _action = 15

    def __init__(self):
        pass
        self.player_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class ClubMoveUp(object):
    _module = 33
    _action = 16

    def __init__(self):
        pass
        self.to_x = 0
        self.to_y = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.to_x))
        buff.extend(struct.pack("<h", self.to_y))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.to_x = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.to_y = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class ClubMoveDown(object):
    _module = 33
    _action = 16

    def __init__(self):
        pass
        self.player_id = 0
        self.to_x = 0
        self.to_y = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<h", self.to_x))
        buff.extend(struct.pack("<h", self.to_y))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.to_x = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.to_y = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 12
        

class NotifyClubhousePlayersDown(object):
    _module = 33
    _action = 17

    def __init__(self):
        pass
        self.players = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.players)))
        for item in self.players:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _players_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_players_size):
            obj = NotifyClubhousePlayersDownPlayers()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.players.append(obj)

    def size(self):
        size = 1
        for item in self.players:
            size += item.size()
        return size


class NotifyClubhousePlayersDownPlayers(object):
    def __init__(self):
        pass
        self.player = Player()

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


class NotifyNewPlayerDown(object):
    _module = 33
    _action = 18

    def __init__(self):
        pass
        self.player = Player()

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


class NotifyUpdatePlayerDown(object):
    _module = 33
    _action = 19

    def __init__(self):
        pass
        self.player_id = 0
        self.fashion_id = 0
        self.in_meditation_state = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<h", self.fashion_id))
        buff.extend(struct.pack("<?", self.in_meditation_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.in_meditation_state = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 11
        

class CliquePublicInfoSummaryUp(object):
    _module = 33
    _action = 20

    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliquePublicInfoSummaryDown(object):
    _module = 33
    _action = 20

    def __init__(self):
        pass
        self.CliqueId = 0
        self.name = ''
        self.level = 0
        self.member_num = 0
        self.owner_nick = ''
        self.owner_pid = 0
        self.announce = ''
        self.cliques = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.CliqueId))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<h", self.member_num))
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        buff.extend(struct.pack('<H', len(self.cliques)))
        for item in self.cliques:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.CliqueId = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.member_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

        _cliques_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_cliques_size):
            obj = CliquePublicInfoSummaryDownCliques()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.cliques.append(obj)

    def size(self):
        size = 28
        size += len(self.name)
        size += len(self.owner_nick)
        size += len(self.announce)
        for item in self.cliques:
            size += item.size()
        return size


class CliquePublicInfoSummaryDownCliques(object):
    def __init__(self):
        pass
        self.clique_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.clique_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class CliqueAutoAuditUp(object):
    _module = 33
    _action = 21

    def __init__(self):
        pass
        self.level = 0
        self.enable = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<?", self.enable))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.enable = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class CliqueAutoAuditDown(object):
    _module = 33
    _action = 21

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
        

class CliqueBaseDonateUp(object):
    _module = 33
    _action = 22

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
        

class CliqueBaseDonateDown(object):
    _module = 33
    _action = 22

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
        

class NotifyLeaveCliqueDown(object):
    _module = 33
    _action = 23

    def __init__(self):
        pass
        self.pid = 0
        self.reason = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.reason))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.reason = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 9
        

class NotifyJoincliqueSuccessDown(object):
    _module = 33
    _action = 24

    def __init__(self):
        pass
        self.pidlist = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.pidlist)))
        for item in self.pidlist:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _pidlist_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_pidlist_size):
            obj = NotifyJoincliqueSuccessDownPidlist()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.pidlist.append(obj)

    def size(self):
        size = 1
        for item in self.pidlist:
            size += item.size()
        return size


class NotifyJoincliqueSuccessDownPidlist(object):
    def __init__(self):
        pass
        self.pid = 0
        self.role_id = 0
        self.level = 0
        self.nick = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<b", self.role_id))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack('<H', len(self.nick)))
        buff.extend(self.nick)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.role_id = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nick = str(raw_msg[idx:idx+_nick_size])
        idx += _nick_size

    def size(self):
        size = 13
        size += len(self.nick)
        return size


class NotifyCliqueMangerActionDown(object):
    _module = 33
    _action = 25

    def __init__(self):
        pass
        self.actiontype = 0
        self.pid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.actiontype))
        buff.extend(struct.pack("<q", self.pid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.actiontype = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 9
        

class CliqueRecruitmentUp(object):
    _module = 33
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
        return 2
        

class CliqueRecruitmentDown(object):
    _module = 33
    _action = 26

    def __init__(self):
        pass
        self.result = 0
        self.timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<q", self.timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 9
        

class NotifyCliqueAnnounceDown(object):
    _module = 33
    _action = 27

    def __init__(self):
        pass
        self.announce = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.announce)))
        buff.extend(self.announce)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _announce_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.announce = str(raw_msg[idx:idx+_announce_size])
        idx += _announce_size

    def size(self):
        size = 2
        size += len(self.announce)
        return size


class NotifyCliqueElectownerDown(object):
    _module = 33
    _action = 28

    def __init__(self):
        pass
        self.ownerid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.ownerid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ownerid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class QuickApplyUp(object):
    _module = 33
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
        return 2
        

class QuickApplyDown(object):
    _module = 33
    _action = 29

    def __init__(self):
        pass
        self.success = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class NotifyContribChangeDown(object):
    _module = 33
    _action = 30

    def __init__(self):
        pass
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.value = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class OtherCliqueUp(object):
    _module = 33
    _action = 31

    def __init__(self):
        pass
        self.page = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.page))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.page = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class OtherCliqueDown(object):
    _module = 33
    _action = 31

    def __init__(self):
        pass
        self.total_num = 0
        self.clique_list = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.total_num))
        buff.extend(struct.pack('<B', len(self.clique_list)))
        for item in self.clique_list:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.total_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _clique_list_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_clique_list_size):
            obj = OtherCliqueDownCliqueList()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.clique_list.append(obj)

    def size(self):
        size = 3
        for item in self.clique_list:
            size += item.size()
        return size


class OtherCliqueDownCliqueList(object):
    def __init__(self):
        pass
        self.rank = 0
        self.name = ''
        self.clique_id = 0
        self.clique_level = 0
        self.owner_name = ''
        self.owner_pid = 0
        self.clique_member_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.rank))
        buff.extend(struct.pack('<H', len(self.name)))
        buff.extend(self.name)
        buff.extend(struct.pack("<q", self.clique_id))
        buff.extend(struct.pack("<h", self.clique_level))
        buff.extend(struct.pack('<H', len(self.owner_name)))
        buff.extend(self.owner_name)
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack("<h", self.clique_member_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.rank = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.name = str(raw_msg[idx:idx+_name_size])
        idx += _name_size

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.clique_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _owner_name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_name = str(raw_msg[idx:idx+_owner_name_size])
        idx += _owner_name_size

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.clique_member_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 26
        size += len(self.name)
        size += len(self.owner_name)
        return size


class CliqueModule(interface.BaseModule):
    decoder_map = {
        0: CreateCliqueDown, 
        1: ApplyJoinCliqueDown, 
        2: CancelApplyCliqueDown, 
        3: ProcessJoinApplyDown, 
        4: ElectOwnerDown, 
        5: MangeMemberDown, 
        6: DestoryCliqueDown, 
        7: UpdateAnnounceDown, 
        8: LeaveCliqueDown, 
        9: ListCliqueDown, 
        10: CliquePublicInfoDown, 
        11: CliqueInfoDown, 
        12: ListApplyDown, 
        13: OperaErrorNotifyDown, 
        14: EnterClubhouseDown, 
        15: LeaveClubhouseDown, 
        16: ClubMoveDown, 
        17: NotifyClubhousePlayersDown, 
        18: NotifyNewPlayerDown, 
        19: NotifyUpdatePlayerDown, 
        20: CliquePublicInfoSummaryDown, 
        21: CliqueAutoAuditDown, 
        22: CliqueBaseDonateDown, 
        23: NotifyLeaveCliqueDown, 
        24: NotifyJoincliqueSuccessDown, 
        25: NotifyCliqueMangerActionDown, 
        26: CliqueRecruitmentDown, 
        27: NotifyCliqueAnnounceDown, 
        28: NotifyCliqueElectownerDown, 
        29: QuickApplyDown, 
        30: NotifyContribChangeDown, 
        31: OtherCliqueDown, 
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

    def add_create_clique(self, callback):
        self.add_callback(0, callback)

    def add_apply_join_clique(self, callback):
        self.add_callback(1, callback)

    def add_cancel_apply_clique(self, callback):
        self.add_callback(2, callback)

    def add_process_join_apply(self, callback):
        self.add_callback(3, callback)

    def add_elect_owner(self, callback):
        self.add_callback(4, callback)

    def add_mange_member(self, callback):
        self.add_callback(5, callback)

    def add_destory_clique(self, callback):
        self.add_callback(6, callback)

    def add_update_announce(self, callback):
        self.add_callback(7, callback)

    def add_leave_clique(self, callback):
        self.add_callback(8, callback)

    def add_list_clique(self, callback):
        self.add_callback(9, callback)

    def add_clique_public_info(self, callback):
        self.add_callback(10, callback)

    def add_clique_info(self, callback):
        self.add_callback(11, callback)

    def add_list_apply(self, callback):
        self.add_callback(12, callback)

    def add_opera_error_notify(self, callback):
        self.add_callback(13, callback)

    def add_enter_clubhouse(self, callback):
        self.add_callback(14, callback)

    def add_leave_clubhouse(self, callback):
        self.add_callback(15, callback)

    def add_club_move(self, callback):
        self.add_callback(16, callback)

    def add_notify_clubhouse_players(self, callback):
        self.add_callback(17, callback)

    def add_notify_new_player(self, callback):
        self.add_callback(18, callback)

    def add_notify_update_player(self, callback):
        self.add_callback(19, callback)

    def add_clique_public_info_summary(self, callback):
        self.add_callback(20, callback)

    def add_clique_auto_audit(self, callback):
        self.add_callback(21, callback)

    def add_clique_base_donate(self, callback):
        self.add_callback(22, callback)

    def add_notify_leave_clique(self, callback):
        self.add_callback(23, callback)

    def add_notify_joinclique_success(self, callback):
        self.add_callback(24, callback)

    def add_notify_clique_manger_action(self, callback):
        self.add_callback(25, callback)

    def add_clique_recruitment(self, callback):
        self.add_callback(26, callback)

    def add_notify_clique_announce(self, callback):
        self.add_callback(27, callback)

    def add_notify_clique_electowner(self, callback):
        self.add_callback(28, callback)

    def add_quick_apply(self, callback):
        self.add_callback(29, callback)

    def add_notify_contrib_change(self, callback):
        self.add_callback(30, callback)

    def add_other_clique(self, callback):
        self.add_callback(31, callback)
