#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

GUIDE_TYPE_GUIDE_ENTER_MISSION = 0
GUIDE_TYPE_GUIDE_SKILL_USE = 1
GUIDE_TYPE_GUIDE_ADVANCED_SKILL_USE = 2
GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP = 3
GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_USE = 4
GUIDE_TYPE_GUIDE_EQUIP_REFINE = 5
GUIDE_TYPE_GUIDE_BATTLE_ITEM_ZHIXUECAO_USE = 6
GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_EQUIP = 7
GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_USE = 8
GUIDE_TYPE_GUIDE_GHOST_EQUIP = 9
GUIDE_TYPE_GUIDE_GHOST_POWER_LOOK = 10
GUIDE_TYPE_GUIDE_GHOST_BATTLE_USE = 11
GUIDE_TYPE_GUIDE_PET_EQUIP = 12
GUIDE_TYPE_GUIDE_PET_BATTLE_USE = 13
GUIDE_TYPE_GUIDE_PET_CATCH = 14
GUIDE_TYPE_GUIDE_HARD_LEVEL = 15
GUIDE_TYPE_GUIDE_TIANJIE = 16
GUIDE_TYPE_GUIDE_GOTO_ZHUBAO = 17
GUIDE_TYPE_GUIDE_SWORD_SOUL = 18
GUIDE_TYPE_GUIDE_FIRST_BATTLE = 19
GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP_FAILE = 20
GUIDE_TYPE_GUIDE_ENTER_MISSION_SECOND = 21
GUIDE_TYPE_GUIDE_BUDDY_EQUIP = 22
GUIDE_TYPE_GUIDE_BUDDY_USE_SKILL_DAODUN = 23
GUIDE_TYPE_GUIDE_EQUIP_USE_SKILL_FENGJUANCANSHENG = 24
GUIDE_TYPE_GUIDE_FRIENDSHIP = 25
GUIDE_TYPE_GUIDE_EQUIP_ROLE_3 = 26
GUIDE_TYPE_GUIDE_EQUIP_ROLE_4 = 27


GUIDE_ACTION_GUIDE_ACCEPT = 0
GUIDE_ACTION_GUIDE_FINISH = 1


class UpdateQuestUp(object):
    _module = 13
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
        

class UpdateQuestDown(object):
    _module = 13
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
        

class GetDailyInfoUp(object):
    _module = 13
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
        

class GetDailyInfoDown(object):
    _module = 13
    _action = 2

    def __init__(self):
        pass
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetDailyInfoDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 1
        for item in self.quest:
            size += item.size()
        return size


class GetDailyInfoDownQuest(object):
    def __init__(self):
        pass
        self.id = 0
        self.finish_count = 0
        self.award_state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        buff.extend(struct.pack("<h", self.finish_count))
        buff.extend(struct.pack("<b", self.award_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.finish_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.award_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class AwardDailyUp(object):
    _module = 13
    _action = 3

    def __init__(self):
        pass
        self.id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class AwardDailyDown(object):
    _module = 13
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
        

class NotifyDailyChangeDown(object):
    _module = 13
    _action = 4

    def __init__(self):
        pass
        self.id = 0
        self.finish_count = 0
        self.award_state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.id))
        buff.extend(struct.pack("<h", self.finish_count))
        buff.extend(struct.pack("<b", self.award_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.finish_count = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.award_state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 5
        

class GuideUp(object):
    _module = 13
    _action = 5

    def __init__(self):
        pass
        self.guide_type = 0
        self.action = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.guide_type))
        buff.extend(struct.pack("<B", self.action))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.guide_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.action = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class GuideDown(object):
    _module = 13
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
        

class GetExtendQuestInfoByNpcIdUp(object):
    _module = 13
    _action = 6

    def __init__(self):
        pass
        self.npc_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.npc_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.npc_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class GetExtendQuestInfoByNpcIdDown(object):
    _module = 13
    _action = 6

    def __init__(self):
        pass
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetExtendQuestInfoByNpcIdDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 1
        for item in self.quest:
            size += item.size()
        return size


class GetExtendQuestInfoByNpcIdDownQuest(object):
    def __init__(self):
        pass
        self.id = 0
        self.progress = 0
        self.state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.id))
        buff.extend(struct.pack("<h", self.progress))
        buff.extend(struct.pack("<b", self.state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.progress = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 7
        

class TakeExtendQuestAwardUp(object):
    _module = 13
    _action = 7

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TakeExtendQuestAwardDown(object):
    _module = 13
    _action = 7

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class GetPannelQuestInfoUp(object):
    _module = 13
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
        

class GetPannelQuestInfoDown(object):
    _module = 13
    _action = 8

    def __init__(self):
        pass
        self.cur_stars = 0
        self.awarded = ''
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.cur_stars))
        buff.extend(struct.pack('<H', len(self.awarded)))
        buff.extend(self.awarded)
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.cur_stars = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _awarded_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.awarded = str(raw_msg[idx:idx+_awarded_size])
        idx += _awarded_size

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetPannelQuestInfoDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 7
        size += len(self.awarded)
        for item in self.quest:
            size += item.size()
        return size


class GetPannelQuestInfoDownQuest(object):
    def __init__(self):
        pass
        self.quest_class = 0
        self.id = 0
        self.progress = 0
        self.state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.quest_class))
        buff.extend(struct.pack("<l", self.id))
        buff.extend(struct.pack("<h", self.progress))
        buff.extend(struct.pack("<b", self.state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_class = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.progress = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 8
        

class GiveUpAdditionQuestUp(object):
    _module = 13
    _action = 10

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class GiveUpAdditionQuestDown(object):
    _module = 13
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
        

class TakeAdditionQuestUp(object):
    _module = 13
    _action = 11

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TakeAdditionQuestDown(object):
    _module = 13
    _action = 11

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
        

class TakeAdditionQuestAwardUp(object):
    _module = 13
    _action = 12

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TakeAdditionQuestAwardDown(object):
    _module = 13
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
        

class GetAdditionQuestUp(object):
    _module = 13
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
        

class GetAdditionQuestDown(object):
    _module = 13
    _action = 13

    def __init__(self):
        pass
        self.quest = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.quest)))
        for item in self.quest:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _quest_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_quest_size):
            obj = GetAdditionQuestDownQuest()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.quest.append(obj)

    def size(self):
        size = 1
        for item in self.quest:
            size += item.size()
        return size


class GetAdditionQuestDownQuest(object):
    def __init__(self):
        pass
        self.quest_id = 0
        self.progress = 0
        self.state = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        buff.extend(struct.pack("<h", self.progress))
        buff.extend(struct.pack("<b", self.state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.progress = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 7
        

class RefreshAdditionQuestUp(object):
    _module = 13
    _action = 14

    def __init__(self):
        pass
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.quest_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class RefreshAdditionQuestDown(object):
    _module = 13
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
        

class TakeQuestStarsAwadedUp(object):
    _module = 13
    _action = 15

    def __init__(self):
        pass
        self.stars_level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.stars_level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.stars_level = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class TakeQuestStarsAwadedDown(object):
    _module = 13
    _action = 15

    def __init__(self):
        pass
        self.result = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.result))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class QuestModule(interface.BaseModule):
    decoder_map = {
        1: UpdateQuestDown, 
        2: GetDailyInfoDown, 
        3: AwardDailyDown, 
        4: NotifyDailyChangeDown, 
        5: GuideDown, 
        6: GetExtendQuestInfoByNpcIdDown, 
        7: TakeExtendQuestAwardDown, 
        8: GetPannelQuestInfoDown, 
        10: GiveUpAdditionQuestDown, 
        11: TakeAdditionQuestDown, 
        12: TakeAdditionQuestAwardDown, 
        13: GetAdditionQuestDown, 
        14: RefreshAdditionQuestDown, 
        15: TakeQuestStarsAwadedDown, 
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

    def add_update_quest(self, callback):
        self.add_callback(1, callback)

    def add_get_daily_info(self, callback):
        self.add_callback(2, callback)

    def add_award_daily(self, callback):
        self.add_callback(3, callback)

    def add_notify_daily_change(self, callback):
        self.add_callback(4, callback)

    def add_guide(self, callback):
        self.add_callback(5, callback)

    def add_get_extend_quest_info_by_npc_id(self, callback):
        self.add_callback(6, callback)

    def add_take_extend_quest_award(self, callback):
        self.add_callback(7, callback)

    def add_get_pannel_quest_info(self, callback):
        self.add_callback(8, callback)

    def add_give_up_addition_quest(self, callback):
        self.add_callback(10, callback)

    def add_take_addition_quest(self, callback):
        self.add_callback(11, callback)

    def add_take_addition_quest_award(self, callback):
        self.add_callback(12, callback)

    def add_get_addition_quest(self, callback):
        self.add_callback(13, callback)

    def add_refresh_addition_quest(self, callback):
        self.add_callback(14, callback)

    def add_take_quest_stars_awaded(self, callback):
        self.add_callback(15, callback)
