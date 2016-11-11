#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

CLIQUE_KONGFU_TRAIN_RESULT_SUCCESS = 0
CLIQUE_KONGFU_TRAIN_RESULT_LACK_CONTRIB = 1
CLIQUE_KONGFU_TRAIN_RESULT_NO_CLIQUE = 2
CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL = 3


CLIQUE_BANK_SOLD_RESULT_SUCCESS = 0
CLIQUE_BANK_SOLD_RESULT_CD = 1
CLIQUE_BANK_SOLD_RESULT_NO_CLIQUE = 2
CLIQUE_BANK_SOLD_RESULT_MAX_LEVEL = 3


CLIQUE_BUILDING_DONATE_RESULT_SUCCESS = 0
CLIQUE_BUILDING_DONATE_RESULT_FAILED = 1


class CliqueBuildingStatusBase(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<q", self.donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliqueBuildingStatusBank(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0
        self.silver_coupon_num = 0
        self.silver_coupon_timespan = 0
        self.gold_coupon_num = 0
        self.gold_coupon_timespan = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.donate_coins))
        buff.extend(struct.pack("<h", self.silver_coupon_num))
        buff.extend(struct.pack("<q", self.silver_coupon_timespan))
        buff.extend(struct.pack("<h", self.gold_coupon_num))
        buff.extend(struct.pack("<q", self.gold_coupon_timespan))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.silver_coupon_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.silver_coupon_timespan = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.gold_coupon_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.gold_coupon_timespan = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 26
        

class CliqueBuildingStatusAttack(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CliqueBuildingStatusHealth(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CliqueBuildingStatusDefence(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CliqueBuildingStatusTemple(object):
    def __init__(self):
        pass
        self.level = 0
        self.donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.level))
        buff.extend(struct.pack("<l", self.donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

ANCESTRAL_HALL_WORSHIP_WHITESANDALWOOD = 1
ANCESTRAL_HALL_WORSHIP_STORAX = 2
ANCESTRAL_HALL_WORSHIP_DAYS = 3


class CliqueBaseDonateUp(object):
    _module = 34
    _action = 1

    def __init__(self):
        pass
        self.money = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.money))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.money = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliqueBaseDonateDown(object):
    _module = 34
    _action = 1

    def __init__(self):
        pass
        self.result = 0
        self.clique_building_base_level = 0
        self.clique_building_base_donate_coins = 0
        self.player_donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<h", self.clique_building_base_level))
        buff.extend(struct.pack("<l", self.clique_building_base_donate_coins))
        buff.extend(struct.pack("<q", self.player_donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.clique_building_base_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.clique_building_base_donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.player_donate_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 15
        

class CliqueBuildingStatusUp(object):
    _module = 34
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
        

class CliqueBuildingStatusDown(object):
    _module = 34
    _action = 2

    def __init__(self):
        pass
        self.success = False
        self.daily_total_donated_coins = 0
        self.base = CliqueBuildingStatusBase()
        self.bank = CliqueBuildingStatusBank()
        self.attack_building = CliqueBuildingStatusAttack()
        self.health_building = CliqueBuildingStatusHealth()
        self.defence_building = CliqueBuildingStatusDefence()
        self.temple_building = CliqueBuildingStatusTemple()

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.success))
        buff.extend(struct.pack("<q", self.daily_total_donated_coins))
        buff.extend(self.base.encode())
        buff.extend(self.bank.encode())
        buff.extend(self.attack_building.encode())
        buff.extend(self.health_building.encode())
        buff.extend(self.defence_building.encode())
        buff.extend(self.temple_building.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.success = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.daily_total_donated_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.base.decode(raw_msg[idx:])
        idx += self.base.size()

        self.bank.decode(raw_msg[idx:])
        idx += self.bank.size()

        self.attack_building.decode(raw_msg[idx:])
        idx += self.attack_building.size()

        self.health_building.decode(raw_msg[idx:])
        idx += self.health_building.size()

        self.defence_building.decode(raw_msg[idx:])
        idx += self.defence_building.size()

        self.temple_building.decode(raw_msg[idx:])
        idx += self.temple_building.size()

    def size(self):
        size = 9
        size += self.base.size()
        size += self.bank.size()
        size += self.attack_building.size()
        size += self.health_building.size()
        size += self.defence_building.size()
        size += self.temple_building.size()
        return size


class CliqueBankDonateUp(object):
    _module = 34
    _action = 3

    def __init__(self):
        pass
        self.money = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.money))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.money = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliqueBankDonateDown(object):
    _module = 34
    _action = 3

    def __init__(self):
        pass
        self.result = 0
        self.clique_building_bank_level = 0
        self.clique_building_bank_donate_coins = 0
        self.player_donate_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<h", self.clique_building_bank_level))
        buff.extend(struct.pack("<l", self.clique_building_bank_donate_coins))
        buff.extend(struct.pack("<q", self.player_donate_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.clique_building_bank_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.clique_building_bank_donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.player_donate_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 15
        

class CliqueBankBuyUp(object):
    _module = 34
    _action = 4

    def __init__(self):
        pass
        self.kind = 0
        self.num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<b", self.kind))
        buff.extend(struct.pack("<b", self.num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.kind = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.num = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 4
        

class CliqueBankBuyDown(object):
    _module = 34
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
        

class CliqueBankSoldUp(object):
    _module = 34
    _action = 5

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
        

class CliqueBankSoldDown(object):
    _module = 34
    _action = 5

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
        

class CliqueKongfuDonateUp(object):
    _module = 34
    _action = 6

    def __init__(self):
        pass
        self.building = 0
        self.money = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.building))
        buff.extend(struct.pack("<q", self.money))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.building = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.money = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 14
        

class CliqueKongfuDonateDown(object):
    _module = 34
    _action = 6

    def __init__(self):
        pass
        self.building = 0
        self.result = 0
        self.total_donate_coins = 0
        self.player_donate_coins = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.building))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<l", self.total_donate_coins))
        buff.extend(struct.pack("<q", self.player_donate_coins))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.building = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.total_donate_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.player_donate_coins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 19
        

class CliqueKongfuInfoUp(object):
    _module = 34
    _action = 7

    def __init__(self):
        pass
        self.building = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.building))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.building = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CliqueKongfuInfoDown(object):
    _module = 34
    _action = 7

    def __init__(self):
        pass
        self.kongfu_list = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.kongfu_list)))
        for item in self.kongfu_list:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _kongfu_list_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_kongfu_list_size):
            obj = CliqueKongfuInfoDownKongfuList()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.kongfu_list.append(obj)

    def size(self):
        size = 1
        for item in self.kongfu_list:
            size += item.size()
        return size


class CliqueKongfuInfoDownKongfuList(object):
    def __init__(self):
        pass
        self.kongfu_id = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.kongfu_id))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.kongfu_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class CliqueKongfuTrainUp(object):
    _module = 34
    _action = 8

    def __init__(self):
        pass
        self.kongfu_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.kongfu_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.kongfu_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class CliqueKongfuTrainDown(object):
    _module = 34
    _action = 8

    def __init__(self):
        pass
        self.kongfu_id = 0
        self.result = 0
        self.level = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.kongfu_id))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<h", self.level))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.kongfu_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 7
        

class CliqueTempleWorshipUp(object):
    _module = 34
    _action = 9

    def __init__(self):
        pass
        self.worship_type = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.worship_type))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.worship_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 3
        

class CliqueTempleWorshipDown(object):
    _module = 34
    _action = 9

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
        

class CliqueTempleDonateUp(object):
    _module = 34
    _action = 10

    def __init__(self):
        pass
        self.money = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.money))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.money = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class CliqueTempleDonateDown(object):
    _module = 34
    _action = 10

    def __init__(self):
        pass
        self.result = 0
        self.temple_building_level = 0
        self.temple_building_coins = 0
        self.totaldonatecoins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<B", self.result))
        buff.extend(struct.pack("<h", self.temple_building_level))
        buff.extend(struct.pack("<l", self.temple_building_coins))
        buff.extend(struct.pack("<q", self.totaldonatecoins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.result = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        self.temple_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.temple_building_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.totaldonatecoins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 15
        

class CliqueTempleInfoUp(object):
    _module = 34
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
        

class CliqueTempleInfoDown(object):
    _module = 34
    _action = 11

    def __init__(self):
        pass
        self.totaldonatecoins = 0
        self.isworship = False
        self.worship_type = 0
        self.worshipcnt = 0
        self.temple_building_level = 0
        self.temple_building_coins = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.totaldonatecoins))
        buff.extend(struct.pack("<?", self.isworship))
        buff.extend(struct.pack("<b", self.worship_type))
        buff.extend(struct.pack("<b", self.worshipcnt))
        buff.extend(struct.pack("<h", self.temple_building_level))
        buff.extend(struct.pack("<l", self.temple_building_coins))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.totaldonatecoins = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.isworship = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1

        self.worship_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.worshipcnt = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.temple_building_level = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.temple_building_coins = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 17
        

class CliqueBuildingModule(interface.BaseModule):
    decoder_map = {
        1: CliqueBaseDonateDown, 
        2: CliqueBuildingStatusDown, 
        3: CliqueBankDonateDown, 
        4: CliqueBankBuyDown, 
        5: CliqueBankSoldDown, 
        6: CliqueKongfuDonateDown, 
        7: CliqueKongfuInfoDown, 
        8: CliqueKongfuTrainDown, 
        9: CliqueTempleWorshipDown, 
        10: CliqueTempleDonateDown, 
        11: CliqueTempleInfoDown, 
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

    def add_clique_base_donate(self, callback):
        self.add_callback(1, callback)

    def add_clique_building_status(self, callback):
        self.add_callback(2, callback)

    def add_clique_bank_donate(self, callback):
        self.add_callback(3, callback)

    def add_clique_bank_buy(self, callback):
        self.add_callback(4, callback)

    def add_clique_bank_sold(self, callback):
        self.add_callback(5, callback)

    def add_clique_kongfu_donate(self, callback):
        self.add_callback(6, callback)

    def add_clique_kongfu_info(self, callback):
        self.add_callback(7, callback)

    def add_clique_kongfu_train(self, callback):
        self.add_callback(8, callback)

    def add_clique_temple_worship(self, callback):
        self.add_callback(9, callback)

    def add_clique_temple_donate(self, callback):
        self.add_callback(10, callback)

    def add_clique_temple_info(self, callback):
        self.add_callback(11, callback)
