#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

BOAT_STATUS_CHANGE_MY_BOAT_HIJACKING = 0
BOAT_STATUS_CHANGE_MY_BOAT_HIJACKED = 1
BOAT_STATUS_CHANGE_MY_BOAT_RECOVERED = 2
BOAT_STATUS_CHANGE_HIJACKED_BOAT_RECOVERED = 3
BOAT_STATUS_CHANGE_ESCORT_FINISHED = 4
BOAT_STATUS_CHANGE_HIJACK_FINISHED = 5


RECOVER_BATTLE_WIN_RESULT_SUCCESS = 0
RECOVER_BATTLE_WIN_RESULT_BOAT_EXPIRE = 1
RECOVER_BATTLE_WIN_RESULT_NO_PERMISSION = 2


HIJACK_BATTLE_WIN_RESULT_SUCCESS = 0
HIJACK_BATTLE_WIN_RESULT_ESCORT_FINISHED = 1
HIJACK_BATTLE_WIN_RESULT_HIJACKED = 2
HIJACK_BATTLE_WIN_RESULT_NO_PERMISSION = 3


START_ESCORT_RESULT_SUCCESS = 0
START_ESCORT_RESULT_ILLEGAL_TIME = 1
START_ESCORT_RESULT_ESCORT_NOT_END = 2
START_ESCORT_RESULT_NO_CLIQUE = 3
START_ESCORT_RESULT_RUN_OUT = 4
START_ESCORT_RESULT_NO_BOAT = 5


HIJACK_BOAT_RESULT_START_BATTLE = 0
HIJACK_BOAT_RESULT_CLIQUE_MEMBER = 1
HIJACK_BOAT_RESULT_HIJACK_NOT_END = 2
HIJACK_BOAT_RESULT_NO_CLIQUE = 3
HIJACK_BOAT_RESULT_RUN_OUT = 4
HIJACK_BOAT_RESULT_NO_BOAT = 5
HIJACK_BOAT_RESULT_CAN_NOT_HIJACK = 6


RECOVER_BOAT_RESULT_START_BATTLE = 0
RECOVER_BOAT_RESULT_NO_PERMISSION = 1
RECOVER_BOAT_RESULT_RECOVERING = 2
RECOVER_BOAT_RESULT_NO_BOAT = 3
RECOVER_BOAT_RESULT_CAN_NOT_RECOVER = 4


ESCORT_STATUS_NONE = 0
ESCORT_STATUS_ESCORT = 1
ESCORT_STATUS_FINISHED = 2


BOAT_STATUS_ESCORT = 0
BOAT_STATUS_HIJACK = 1
BOAT_STATUS_RECOVER = 2
BOAT_STATUS_HIJACK_FINISH = 3


HIJACK_STATUS_NONE = 0
HIJACK_STATUS_HIJACK = 1
HIJACK_STATUS_FINISHED = 2


class CliqueBoatMessage(object):
    def __init__(self):
        pass
        self.id = 0
        self.tpl_id = 0
        self.parameters = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<h", self.tpl_id))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.tpl_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

    def size(self):
        size = 12
        size += len(self.parameters)
        return size


class CliqueBoat(object):
    def __init__(self):
        pass
        self.boat_id = 0
        self.boat_type = 0
        self.status = 0
        self.escort_time = 0
        self.start_timestamp = 0
        self.owner_pid = 0
        self.owner_nick = ''
        self.owner_level = 0
        self.hijacker_nick = ''
        self.hijacker_pid = 0
        self.clique_id = 0
        self.clique_name = ''
        self.hijack_start_timestamp = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        buff.extend(struct.pack("<h", self.boat_type))
        buff.extend(struct.pack("<B", self.status))
        buff.extend(struct.pack("<q", self.escort_time))
        buff.extend(struct.pack("<q", self.start_timestamp))
        buff.extend(struct.pack("<q", self.owner_pid))
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        buff.extend(struct.pack("<h", self.owner_level))
        buff.extend(struct.pack('<H', len(self.hijacker_nick)))
        buff.extend(self.hijacker_nick)
        buff.extend(struct.pack("<q", self.hijacker_pid))
        buff.extend(struct.pack("<q", self.clique_id))
        buff.extend(struct.pack('<H', len(self.clique_name)))
        buff.extend(self.clique_name)
        buff.extend(struct.pack("<q", self.hijack_start_timestamp))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.boat_type = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.escort_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.start_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

        self.owner_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        _hijacker_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hijacker_nick = str(raw_msg[idx:idx+_hijacker_nick_size])
        idx += _hijacker_nick_size

        self.hijacker_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.clique_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _clique_name_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.clique_name = str(raw_msg[idx:idx+_clique_name_size])
        idx += _clique_name_size

        self.hijack_start_timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

    def size(self):
        size = 67
        size += len(self.owner_nick)
        size += len(self.hijacker_nick)
        size += len(self.clique_name)
        return size


class EscortInfoUp(object):
    _module = 36
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
        

class EscortInfoDown(object):
    _module = 36
    _action = 0

    def __init__(self):
        pass
        self.daily_escort_num = 0
        self.daily_hijack_num = 0
        self.escort_status = 0
        self.hijack_status = 0
        self.boat_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.daily_escort_num))
        buff.extend(struct.pack("<h", self.daily_hijack_num))
        buff.extend(struct.pack("<B", self.escort_status))
        buff.extend(struct.pack("<B", self.hijack_status))
        buff.extend(struct.pack("<h", self.boat_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.daily_escort_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.daily_hijack_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.escort_status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.hijack_status = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.boat_type = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 8
        

class GetIngotBoatUp(object):
    _module = 36
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
        

class GetIngotBoatDown(object):
    _module = 36
    _action = 1

    def __init__(self):
        pass
        self.ok = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.ok))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ok = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class StartEscortUp(object):
    _module = 36
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
        return 2
        

class StartEscortDown(object):
    _module = 36
    _action = 2

    def __init__(self):
        pass
        self.result = 0
        self.boat = CliqueBoat()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(self.boat.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.boat.decode(raw_msg[idx:])
        idx += self.boat.size()

    def size(self):
        size = 1
        size += self.boat.size()
        return size


class HijackBoatUp(object):
    _module = 36
    _action = 3

    def __init__(self):
        pass
        self.boat_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class HijackBoatDown(object):
    _module = 36
    _action = 3

    def __init__(self):
        pass
        self.result = 0
        self.boat_id = 0
        self.boat_owner_pid = 0
        self.boat_owner_nick = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<q", self.boat_id))
        buff.extend(struct.pack("<q", self.boat_owner_pid))
        buff.extend(struct.pack('<H', len(self.boat_owner_nick)))
        buff.extend(self.boat_owner_nick)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.boat_owner_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _boat_owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.boat_owner_nick = str(raw_msg[idx:idx+_boat_owner_nick_size])
        idx += _boat_owner_nick_size

    def size(self):
        size = 19
        size += len(self.boat_owner_nick)
        return size


class RecoverBoatUp(object):
    _module = 36
    _action = 4

    def __init__(self):
        pass
        self.boat_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class RecoverBoatDown(object):
    _module = 36
    _action = 4

    def __init__(self):
        pass
        self.boat_id = 0
        self.result = 0
        self.hijacker_pid = 0
        self.hijacker_nick = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<q", self.hijacker_pid))
        buff.extend(struct.pack('<H', len(self.hijacker_nick)))
        buff.extend(self.hijacker_nick)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.hijacker_pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _hijacker_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.hijacker_nick = str(raw_msg[idx:idx+_hijacker_nick_size])
        idx += _hijacker_nick_size

    def size(self):
        size = 19
        size += len(self.hijacker_nick)
        return size


class ListBoatsUp(object):
    _module = 36
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
        

class ListBoatsDown(object):
    _module = 36
    _action = 5

    def __init__(self):
        pass
        self.boats = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.boats)))
        for item in self.boats:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _boats_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_boats_size):
            obj = ListBoatsDownBoats()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.boats.append(obj)

    def size(self):
        size = 1
        for item in self.boats:
            size += item.size()
        return size


class ListBoatsDownBoats(object):
    def __init__(self):
        pass
        self.boat = CliqueBoat()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.boat.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat.decode(raw_msg[idx:])
        idx += self.boat.size()

    def size(self):
        size = 0
        size += self.boat.size()
        return size


class GetRandomBoatUp(object):
    _module = 36
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
        

class GetRandomBoatDown(object):
    _module = 36
    _action = 6

    def __init__(self):
        pass
        self.boat_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.boat_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_type = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class NotifyEscortFinishedDown(object):
    _module = 36
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
        

class NotifyHijackFinishedDown(object):
    _module = 36
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
        

class NotifyRecoverBattleWinDown(object):
    _module = 36
    _action = 9

    def __init__(self):
        pass
        self.boat_id = 0
        self.result = 0
        self.owner_nick = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

    def size(self):
        size = 11
        size += len(self.owner_nick)
        return size


class NotifyHijackBattleWinDown(object):
    _module = 36
    _action = 10

    def __init__(self):
        pass
        self.boat_id = 0
        self.result = 0
        self.owner_nick = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.boat_id))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack('<H', len(self.owner_nick)))
        buff.extend(self.owner_nick)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _owner_nick_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.owner_nick = str(raw_msg[idx:idx+_owner_nick_size])
        idx += _owner_nick_size

    def size(self):
        size = 11
        size += len(self.owner_nick)
        return size


class TakeHijackAwardUp(object):
    _module = 36
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
        

class TakeHijackAwardDown(object):
    _module = 36
    _action = 11

    def __init__(self):
        pass
        self.ok = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.ok))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ok = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class TakeEscortAwardUp(object):
    _module = 36
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
        

class TakeEscortAwardDown(object):
    _module = 36
    _action = 12

    def __init__(self):
        pass
        self.ok = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.ok))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.ok = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class GetCliqueBoatMessagesUp(object):
    _module = 36
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
        return 2
        

class GetCliqueBoatMessagesDown(object):
    _module = 36
    _action = 13

    def __init__(self):
        pass
        self.messages = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.messages)))
        for item in self.messages:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _messages_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_messages_size):
            obj = GetCliqueBoatMessagesDownMessages()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.messages.append(obj)

    def size(self):
        size = 1
        for item in self.messages:
            size += item.size()
        return size


class GetCliqueBoatMessagesDownMessages(object):
    def __init__(self):
        pass
        self.message = CliqueBoatMessage()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.message.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.message.decode(raw_msg[idx:])
        idx += self.message.size()

    def size(self):
        size = 0
        size += self.message.size()
        return size


class SendCliqueBoatMessageDown(object):
    _module = 36
    _action = 14

    def __init__(self):
        pass
        self.message = CliqueBoatMessage()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.message.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.message.decode(raw_msg[idx:])
        idx += self.message.size()

    def size(self):
        size = 0
        size += self.message.size()
        return size


class ReadCliqueBoatMessageUp(object):
    _module = 36
    _action = 15

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
        

class ReadCliqueBoatMessageDown(object):
    _module = 36
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
        

class NotifyBoatStatusChangeDown(object):
    _module = 36
    _action = 16

    def __init__(self):
        pass
        self.boat = CliqueBoat()
        self.change = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(self.boat.encode())
        buff.extend(struct.pack("<B", self.change))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.boat.decode(raw_msg[idx:])
        idx += self.boat.size()

        self.change = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

    def size(self):
        size = 1
        size += self.boat.size()
        return size


class CliqueEscortModule(interface.BaseModule):
    decoder_map = {
        0: EscortInfoDown, 
        1: GetIngotBoatDown, 
        2: StartEscortDown, 
        3: HijackBoatDown, 
        4: RecoverBoatDown, 
        5: ListBoatsDown, 
        6: GetRandomBoatDown, 
        7: NotifyEscortFinishedDown, 
        8: NotifyHijackFinishedDown, 
        9: NotifyRecoverBattleWinDown, 
        10: NotifyHijackBattleWinDown, 
        11: TakeHijackAwardDown, 
        12: TakeEscortAwardDown, 
        13: GetCliqueBoatMessagesDown, 
        14: SendCliqueBoatMessageDown, 
        15: ReadCliqueBoatMessageDown, 
        16: NotifyBoatStatusChangeDown, 
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

    def add_escort_info(self, callback):
        self.add_callback(0, callback)

    def add_get_ingot_boat(self, callback):
        self.add_callback(1, callback)

    def add_start_escort(self, callback):
        self.add_callback(2, callback)

    def add_hijack_boat(self, callback):
        self.add_callback(3, callback)

    def add_recover_boat(self, callback):
        self.add_callback(4, callback)

    def add_list_boats(self, callback):
        self.add_callback(5, callback)

    def add_get_random_boat(self, callback):
        self.add_callback(6, callback)

    def add_notify_escort_finished(self, callback):
        self.add_callback(7, callback)

    def add_notify_hijack_finished(self, callback):
        self.add_callback(8, callback)

    def add_notify_recover_battle_win(self, callback):
        self.add_callback(9, callback)

    def add_notify_hijack_battle_win(self, callback):
        self.add_callback(10, callback)

    def add_take_hijack_award(self, callback):
        self.add_callback(11, callback)

    def add_take_escort_award(self, callback):
        self.add_callback(12, callback)

    def add_get_clique_boat_messages(self, callback):
        self.add_callback(13, callback)

    def add_send_clique_boat_message(self, callback):
        self.add_callback(14, callback)

    def add_read_clique_boat_message(self, callback):
        self.add_callback(15, callback)

    def add_notify_boat_status_change(self, callback):
        self.add_callback(16, callback)
