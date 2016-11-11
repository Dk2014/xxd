#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class LoginAwardInfoUp(object):
    _module = 24
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
        

class LoginAwardInfoDown(object):
    _module = 24
    _action = 1

    def __init__(self):
        pass
        self.record = 0
        self.total_login_days = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.record))
        buff.extend(struct.pack("<l", self.total_login_days))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.record = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.total_login_days = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 8
        

class TakeLoginAwardUp(object):
    _module = 24
    _action = 2

    def __init__(self):
        pass
        self.day = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.day))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.day = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TakeLoginAwardDown(object):
    _module = 24
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
        

class GetEventsUp(object):
    _module = 24
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
        

class GetEventsDown(object):
    _module = 24
    _action = 3

    def __init__(self):
        pass
        self.events = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.events)))
        for item in self.events:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _events_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_events_size):
            obj = GetEventsDownEvents()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.events.append(obj)

    def size(self):
        size = 1
        for item in self.events:
            size += item.size()
        return size


class GetEventsDownEvents(object):
    def __init__(self):
        pass
        self.event_id = 0
        self.process = 0
        self.player_process = 0
        self.page = 0
        self.is_award = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.event_id))
        buff.extend(struct.pack("<l", self.process))
        buff.extend(struct.pack("<l", self.player_process))
        buff.extend(struct.pack("<l", self.page))
        buff.extend(struct.pack("<?", self.is_award))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.event_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.process = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.player_process = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.page = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.is_award = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 15
        

class GetEventAwardUp(object):
    _module = 24
    _action = 4

    def __init__(self):
        pass
        self.event_id = 0
        self.page = 0
        self.param1 = 0
        self.param2 = 0
        self.param3 = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.event_id))
        buff.extend(struct.pack("<l", self.page))
        buff.extend(struct.pack("<l", self.param1))
        buff.extend(struct.pack("<l", self.param2))
        buff.extend(struct.pack("<l", self.param3))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.event_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.page = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.param1 = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.param2 = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.param3 = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 20
        

class GetEventAwardDown(object):
    _module = 24
    _action = 4

    def __init__(self):
        pass
        self.result = 0
        self.award = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.result))
        buff.extend(struct.pack("<l", self.award))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.award = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class PlayerEventPhysicalCostUp(object):
    _module = 24
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
        

class PlayerEventPhysicalCostDown(object):
    _module = 24
    _action = 6

    def __init__(self):
        pass
        self.value = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.value))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.value = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class PlayerMonthCardInfoUp(object):
    _module = 24
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
        return 2
        

class PlayerMonthCardInfoDown(object):
    _module = 24
    _action = 7

    def __init__(self):
        pass
        self.BeginTime = 0
        self.EndTime = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.BeginTime))
        buff.extend(struct.pack("<q", self.EndTime))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.BeginTime = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.EndTime = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 16
        

class GetSevenInfoUp(object):
    _module = 24
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
        

class GetSevenInfoDown(object):
    _module = 24
    _action = 8

    def __init__(self):
        pass
        self.status = 0
        self.day = 0
        self.schedule = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<h", self.day))
        buff.extend(struct.pack("<l", self.schedule))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.day = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.schedule = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 7
        

class GetSevenAwardUp(object):
    _module = 24
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
        

class GetSevenAwardDown(object):
    _module = 24
    _action = 9

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
        

class GetRichmanClubInfoUp(object):
    _module = 24
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
        

class GetRichmanClubInfoDown(object):
    _module = 24
    _action = 10

    def __init__(self):
        pass
        self.status = 0
        self.schedule = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.status))
        buff.extend(struct.pack("<l", self.schedule))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.schedule = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 5
        

class GetRichmanClubAwardUp(object):
    _module = 24
    _action = 11

    def __init__(self):
        pass
        self.column = 0
        self.sequence = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.column))
        buff.extend(struct.pack("<b", self.sequence))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.column = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.sequence = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class GetRichmanClubAwardDown(object):
    _module = 24
    _action = 11

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
        

class InfoShareUp(object):
    _module = 24
    _action = 12

    def __init__(self):
        pass
        self.is_share = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.is_share))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.is_share = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class InfoShareDown(object):
    _module = 24
    _action = 12

    def __init__(self):
        pass
        self.count = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.count))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class InfoGroupBuyUp(object):
    _module = 24
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
        

class InfoGroupBuyDown(object):
    _module = 24
    _action = 13

    def __init__(self):
        pass
        self.count = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.count))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.count = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class GetIngotChangeTotalUp(object):
    _module = 24
    _action = 14

    def __init__(self):
        pass
        self.is_in = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.is_in))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.is_in = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class GetIngotChangeTotalDown(object):
    _module = 24
    _action = 14

    def __init__(self):
        pass
        self.count = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.count))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.count = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 8
        

class GetEventTotalAwardUp(object):
    _module = 24
    _action = 15

    def __init__(self):
        pass
        self.order = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.order))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.order = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class GetEventTotalAwardDown(object):
    _module = 24
    _action = 15

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
        

class GetEventArenaRankUp(object):
    _module = 24
    _action = 16

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
        

class GetEventArenaRankDown(object):
    _module = 24
    _action = 16

    def __init__(self):
        pass
        self.rank = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.rank))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.rank = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class GetEventTenDrawTimesUp(object):
    _module = 24
    _action = 17

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
        

class GetEventTenDrawTimesDown(object):
    _module = 24
    _action = 17

    def __init__(self):
        pass
        self.times = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.times))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.times = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class GetEventRechargeAwardUp(object):
    _module = 24
    _action = 18

    def __init__(self):
        pass
        self.page = 0
        self.requireid = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.page))
        buff.extend(struct.pack("<l", self.requireid))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.page = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.requireid = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 10
        

class GetEventRechargeAwardDown(object):
    _module = 24
    _action = 18

    def __init__(self):
        pass
        self.is_rechage = False
        self.is_award = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.is_rechage))
        buff.extend(struct.pack("<?", self.is_award))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.is_rechage = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.is_award = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 2
        

class GetEventNewYearUp(object):
    _module = 24
    _action = 19

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
        

class GetEventNewYearDown(object):
    _module = 24
    _action = 19

    def __init__(self):
        pass
        self.processes = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.processes)))
        for item in self.processes:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _processes_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_processes_size):
            obj = GetEventNewYearDownProcesses()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.processes.append(obj)

    def size(self):
        size = 1
        for item in self.processes:
            size += item.size()
        return size


class GetEventNewYearDownProcesses(object):
    def __init__(self):
        pass
        self.day = 0
        self.ingot = 0
        self.is_awarded = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.day))
        buff.extend(struct.pack("<l", self.ingot))
        buff.extend(struct.pack("<?", self.is_awarded))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.day = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.ingot = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.is_awarded = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 6
        

class QqVipContinueUp(object):
    _module = 24
    _action = 20

    def __init__(self):
        pass
        self.kind = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.kind))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.kind = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class QqVipContinueDown(object):
    _module = 24
    _action = 20

    def __init__(self):
        pass
        self.status = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.status))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.status = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class EventModule(interface.BaseModule):
    decoder_map = {
        1: LoginAwardInfoDown, 
        2: TakeLoginAwardDown, 
        3: GetEventsDown, 
        4: GetEventAwardDown, 
        6: PlayerEventPhysicalCostDown, 
        7: PlayerMonthCardInfoDown, 
        8: GetSevenInfoDown, 
        9: GetSevenAwardDown, 
        10: GetRichmanClubInfoDown, 
        11: GetRichmanClubAwardDown, 
        12: InfoShareDown, 
        13: InfoGroupBuyDown, 
        14: GetIngotChangeTotalDown, 
        15: GetEventTotalAwardDown, 
        16: GetEventArenaRankDown, 
        17: GetEventTenDrawTimesDown, 
        18: GetEventRechargeAwardDown, 
        19: GetEventNewYearDown, 
        20: QqVipContinueDown, 
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

    def add_login_award_info(self, callback):
        self.add_callback(1, callback)

    def add_take_login_award(self, callback):
        self.add_callback(2, callback)

    def add_get_events(self, callback):
        self.add_callback(3, callback)

    def add_get_event_award(self, callback):
        self.add_callback(4, callback)

    def add_player_event_physical_cost(self, callback):
        self.add_callback(6, callback)

    def add_player_month_card_info(self, callback):
        self.add_callback(7, callback)

    def add_get_seven_info(self, callback):
        self.add_callback(8, callback)

    def add_get_seven_award(self, callback):
        self.add_callback(9, callback)

    def add_get_richman_club_info(self, callback):
        self.add_callback(10, callback)

    def add_get_richman_club_award(self, callback):
        self.add_callback(11, callback)

    def add_info_share(self, callback):
        self.add_callback(12, callback)

    def add_info_group_buy(self, callback):
        self.add_callback(13, callback)

    def add_get_ingot_change_total(self, callback):
        self.add_callback(14, callback)

    def add_get_event_total_award(self, callback):
        self.add_callback(15, callback)

    def add_get_event_arena_rank(self, callback):
        self.add_callback(16, callback)

    def add_get_event_ten_draw_times(self, callback):
        self.add_callback(17, callback)

    def add_get_event_recharge_award(self, callback):
        self.add_callback(18, callback)

    def add_get_event_new_year(self, callback):
        self.add_callback(19, callback)

    def add_qq_vip_continue(self, callback):
        self.add_callback(20, callback)
