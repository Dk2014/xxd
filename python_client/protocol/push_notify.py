#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class PushInfoUp(object):
    _module = 26
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
        

class PushInfoDown(object):
    _module = 26
    _action = 1

    def __init__(self):
        pass
        self.effect_notification = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack('<B', len(self.effect_notification)))
        for item in self.effect_notification:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        _effect_notification_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_effect_notification_size):
            obj = PushInfoDownEffectNotification()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.effect_notification.append(obj)

    def size(self):
        size = 1
        for item in self.effect_notification:
            size += item.size()
        return size


class PushInfoDownEffectNotification(object):
    def __init__(self):
        pass
        self.notification_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.notification_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.notification_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4


    @staticmethod
    def size():
        return 4
        

class PushNotificationSwitchUp(object):
    _module = 26
    _action = 2

    def __init__(self):
        pass
        self.notification_id = 0
        self.turn_on = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.notification_id))
        buff.extend(struct.pack("<?", self.turn_on))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.notification_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.turn_on = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 7
        

class PushNotificationSwitchDown(object):
    _module = 26
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
        

class PushNotifyModule(interface.BaseModule):
    decoder_map = {
        1: PushInfoDown, 
        2: PushNotificationSwitchDown, 
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

    def add_push_info(self, callback):
        self.add_callback(1, callback)

    def add_push_notification_switch(self, callback):
        self.add_callback(2, callback)
