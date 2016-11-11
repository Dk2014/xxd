#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetListUp(object):
    _module = 18
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
        

class GetListDown(object):
    _module = 18
    _action = 0

    def __init__(self):
        pass
        self.announcements = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.announcements)))
        for item in self.announcements:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _announcements_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_announcements_size):
            obj = GetListDownAnnouncements()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.announcements.append(obj)

    def size(self):
        size = 1
        for item in self.announcements:
            size += item.size()
        return size


class GetListDownAnnouncements(object):
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


class AnnouncementModule(interface.BaseModule):
    decoder_map = {
        0: GetListDown, 
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

    def add_get_list(self, callback):
        self.add_callback(0, callback)
