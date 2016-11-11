#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetVersionUp(object):
    _module = 88
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
        

class GetVersionDown(object):
    _module = 88
    _action = 0

    def __init__(self):
        pass
        self.version = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.version)))
        buff.extend(self.version)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _version_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.version = str(raw_msg[idx:idx+_version_size])
        idx += _version_size

    def size(self):
        size = 2
        size += len(self.version)
        return size


class GetApiCountUp(object):
    _module = 88
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
        

class GetApiCountDown(object):
    _module = 88
    _action = 1

    def __init__(self):
        pass
        self.count_in = 0
        self.count_out = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.count_in))
        buff.extend(struct.pack("<q", self.count_out))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.count_in = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.count_out = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 16
        

class SearchPlayerRoleUp(object):
    _module = 88
    _action = 2

    def __init__(self):
        pass
        self.openid = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.openid)))
        buff.extend(self.openid)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _openid_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.openid = str(raw_msg[idx:idx+_openid_size])
        idx += _openid_size

    def size(self):
        size = 4
        size += len(self.openid)
        return size


class SearchPlayerRoleDown(object):
    _module = 88
    _action = 2

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
        

class UpdateAccessTokenUp(object):
    _module = 88
    _action = 3

    def __init__(self):
        pass
        self.token = ''
        self.pfkey = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.token)))
        buff.extend(self.token)
        buff.extend(struct.pack('<H', len(self.pfkey)))
        buff.extend(self.pfkey)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _token_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.token = str(raw_msg[idx:idx+_token_size])
        idx += _token_size

        _pfkey_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.pfkey = str(raw_msg[idx:idx+_pfkey_size])
        idx += _pfkey_size

    def size(self):
        size = 6
        size += len(self.token)
        size += len(self.pfkey)
        return size


class UpdateAccessTokenDown(object):
    _module = 88
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
        

class UpdateEventDataUp(object):
    _module = 88
    _action = 4

    def __init__(self):
        pass
        self.version = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.version))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.version = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 6
        

class UpdateEventDataDown(object):
    _module = 88
    _action = 4

    def __init__(self):
        pass
        self.json = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.json)))
        buff.extend(self.json)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _json_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.json = str(raw_msg[idx:idx+_json_size])
        idx += _json_size

    def size(self):
        size = 2
        size += len(self.json)
        return size


class TssDataUp(object):
    _module = 88
    _action = 5

    def __init__(self):
        pass
        self.data = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.data)))
        buff.extend(self.data)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _data_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.data = str(raw_msg[idx:idx+_data_size])
        idx += _data_size

    def size(self):
        size = 4
        size += len(self.data)
        return size


class TssDataDown(object):
    _module = 88
    _action = 5

    def __init__(self):
        pass
        self.data = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<H', len(self.data)))
        buff.extend(self.data)
        return buff

    def decode(self, raw_msg):
        idx = 0

        _data_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.data = str(raw_msg[idx:idx+_data_size])
        idx += _data_size

    def size(self):
        size = 2
        size += len(self.data)
        return size


class ServerInfoModule(interface.BaseModule):
    decoder_map = {
        0: GetVersionDown, 
        1: GetApiCountDown, 
        2: SearchPlayerRoleDown, 
        3: UpdateAccessTokenDown, 
        4: UpdateEventDataDown, 
        5: TssDataDown, 
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

    def add_get_version(self, callback):
        self.add_callback(0, callback)

    def add_get_api_count(self, callback):
        self.add_callback(1, callback)

    def add_search_player_role(self, callback):
        self.add_callback(2, callback)

    def add_update_access_token(self, callback):
        self.add_callback(3, callback)

    def add_update_event_data(self, callback):
        self.add_callback(4, callback)

    def add_tss_data(self, callback):
        self.add_callback(5, callback)
