#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class TownPlayer(object):
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


class EnterUp(object):
    _module = 1
    _action = 0

    def __init__(self):
        pass
        self.town_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.town_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.town_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class EnterDown(object):
    _module = 1
    _action = 0

    def __init__(self):
        pass
        self.player = TownPlayer()

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


class LeaveUp(object):
    _module = 1
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
        

class LeaveDown(object):
    _module = 1
    _action = 1

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
        

class MoveUp(object):
    _module = 1
    _action = 2

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
        

class MoveDown(object):
    _module = 1
    _action = 2

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
        

class TalkedNpcListUp(object):
    _module = 1
    _action = 3

    def __init__(self):
        pass
        self.town_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.town_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.town_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class TalkedNpcListDown(object):
    _module = 1
    _action = 3

    def __init__(self):
        pass
        self.npc_list = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.npc_list)))
        for item in self.npc_list:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _npc_list_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_npc_list_size):
            obj = TalkedNpcListDownNpcList()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.npc_list.append(obj)

    def size(self):
        size = 1
        for item in self.npc_list:
            size += item.size()
        return size


class TalkedNpcListDownNpcList(object):
    def __init__(self):
        pass
        self.npc_id = 0
        self.quest_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.npc_id))
        buff.extend(struct.pack("<h", self.quest_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.npc_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.quest_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class NpcTalkUp(object):
    _module = 1
    _action = 4

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
        

class NpcTalkDown(object):
    _module = 1
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
        

class NotifyTownPlayersDown(object):
    _module = 1
    _action = 5

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
            obj = NotifyTownPlayersDownPlayers()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.players.append(obj)

    def size(self):
        size = 1
        for item in self.players:
            size += item.size()
        return size


class NotifyTownPlayersDownPlayers(object):
    def __init__(self):
        pass
        self.player = TownPlayer()

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


class UpdateTownPlayerDown(object):
    _module = 1
    _action = 6

    def __init__(self):
        pass
        self.player_id = 0
        self.fashion_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<h", self.fashion_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.fashion_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 10
        

class UpdateTownPlayerMeditationStateDown(object):
    _module = 1
    _action = 7

    def __init__(self):
        pass
        self.player_id = 0
        self.meditation_state = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.player_id))
        buff.extend(struct.pack("<?", self.meditation_state))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.player_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.meditation_state = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 9
        

class ListOpenedTownTreasuresUp(object):
    _module = 1
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
        

class ListOpenedTownTreasuresDown(object):
    _module = 1
    _action = 8

    def __init__(self):
        pass
        self.treasures = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.treasures)))
        for item in self.treasures:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _treasures_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_treasures_size):
            obj = ListOpenedTownTreasuresDownTreasures()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.treasures.append(obj)

    def size(self):
        size = 2
        for item in self.treasures:
            size += item.size()
        return size


class ListOpenedTownTreasuresDownTreasures(object):
    def __init__(self):
        pass
        self.town_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.town_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.town_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 2
        

class TakeTownTreasuresUp(object):
    _module = 1
    _action = 9

    def __init__(self):
        pass
        self.town_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.town_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.town_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 4
        

class TakeTownTreasuresDown(object):
    _module = 1
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
        

class TownModule(interface.BaseModule):
    decoder_map = {
        0: EnterDown, 
        1: LeaveDown, 
        2: MoveDown, 
        3: TalkedNpcListDown, 
        4: NpcTalkDown, 
        5: NotifyTownPlayersDown, 
        6: UpdateTownPlayerDown, 
        7: UpdateTownPlayerMeditationStateDown, 
        8: ListOpenedTownTreasuresDown, 
        9: TakeTownTreasuresDown, 
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

    def add_enter(self, callback):
        self.add_callback(0, callback)

    def add_leave(self, callback):
        self.add_callback(1, callback)

    def add_move(self, callback):
        self.add_callback(2, callback)

    def add_talked_npc_list(self, callback):
        self.add_callback(3, callback)

    def add_npc_talk(self, callback):
        self.add_callback(4, callback)

    def add_notify_town_players(self, callback):
        self.add_callback(5, callback)

    def add_update_town_player(self, callback):
        self.add_callback(6, callback)

    def add_update_town_player_meditation_state(self, callback):
        self.add_callback(7, callback)

    def add_list_opened_town_treasures(self, callback):
        self.add_callback(8, callback)

    def add_take_town_treasures(self, callback):
        self.add_callback(9, callback)
