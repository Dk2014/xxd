#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

MESSAGE_TYPE_CHAT = 0
MESSAGE_TYPE_RARE_PROPS = 1
MESSAGE_TYPE_CLIQUE_MESSAGE = 2
MESSAGE_TYPE_CLIQUE_CHAT = 3
MESSAGE_TYPE_CLIQUE_NEWS = 4


class CliqueMessage(object):
    def __init__(self):
        pass
        self.tpl_id = 0
        self.pid = 0
        self.msg_type = 0
        self.nickname = ''
        self.timestamp = 0
        self.parameters = ''

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.tpl_id))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.msg_type))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<q", self.timestamp))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.tpl_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.msg_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

    def size(self):
        size = 23
        size += len(self.nickname)
        size += len(self.parameters)
        return size


class GetLatestWorldChannelMessageUp(object):
    _module = 29
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
        

class GetLatestWorldChannelMessageDown(object):
    _module = 29
    _action = 0

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
            obj = GetLatestWorldChannelMessageDownMessages()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.messages.append(obj)

    def size(self):
        size = 1
        for item in self.messages:
            size += item.size()
        return size


class GetLatestWorldChannelMessageDownMessages(object):
    def __init__(self):
        pass
        self.pid = 0
        self.msg_type = 0
        self.nickname = ''
        self.timestamp = 0
        self.parameters = ''
        self.tpl_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.msg_type))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<q", self.timestamp))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        buff.extend(struct.pack("<h", self.tpl_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.msg_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

        self.tpl_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 23
        size += len(self.nickname)
        size += len(self.parameters)
        return size


class AddWorldChatUp(object):
    _module = 29
    _action = 1

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


class AddWorldChatDown(object):
    _module = 29
    _action = 1

    def __init__(self):
        pass
        self.banned = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.banned))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.banned = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class WorldChannelInfoUp(object):
    _module = 29
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
        

class WorldChannelInfoDown(object):
    _module = 29
    _action = 2

    def __init__(self):
        pass
        self.timestamp = 0
        self.daily_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.timestamp))
        buff.extend(struct.pack("<h", self.daily_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.daily_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 10
        

class SendGlobalMessagesDown(object):
    _module = 29
    _action = 3

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
            obj = SendGlobalMessagesDownMessages()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.messages.append(obj)

    def size(self):
        size = 1
        for item in self.messages:
            size += item.size()
        return size


class SendGlobalMessagesDownMessages(object):
    def __init__(self):
        pass
        self.pid = 0
        self.msg_type = 0
        self.nickname = ''
        self.timestamp = 0
        self.parameters = ''
        self.tpl_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.pid))
        buff.extend(struct.pack("<B", self.msg_type))
        buff.extend(struct.pack('<H', len(self.nickname)))
        buff.extend(self.nickname)
        buff.extend(struct.pack("<q", self.timestamp))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        buff.extend(struct.pack("<h", self.tpl_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.pid = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.msg_type = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1

        _nickname_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.nickname = str(raw_msg[idx:idx+_nickname_size])
        idx += _nickname_size

        self.timestamp = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

        self.tpl_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

    def size(self):
        size = 23
        size += len(self.nickname)
        size += len(self.parameters)
        return size


class AddCliqueChatUp(object):
    _module = 29
    _action = 4

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


class AddCliqueChatDown(object):
    _module = 29
    _action = 4

    def __init__(self):
        pass
        self.banned = False

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<?", self.banned))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.banned = struct.unpack_from("<?", raw_msg, idx)[0]
        idx += 1


    @staticmethod
    def size():
        return 1
        

class GetLatestCliqueMessagesUp(object):
    _module = 29
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
        

class GetLatestCliqueMessagesDown(object):
    _module = 29
    _action = 5

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
            obj = GetLatestCliqueMessagesDownMessages()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.messages.append(obj)

    def size(self):
        size = 1
        for item in self.messages:
            size += item.size()
        return size


class GetLatestCliqueMessagesDownMessages(object):
    def __init__(self):
        pass
        self.message = CliqueMessage()

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


class SendCliqueMessageDown(object):
    _module = 29
    _action = 6

    def __init__(self):
        pass
        self.message = CliqueMessage()

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


class ChannelModule(interface.BaseModule):
    decoder_map = {
        0: GetLatestWorldChannelMessageDown, 
        1: AddWorldChatDown, 
        2: WorldChannelInfoDown, 
        3: SendGlobalMessagesDown, 
        4: AddCliqueChatDown, 
        5: GetLatestCliqueMessagesDown, 
        6: SendCliqueMessageDown, 
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

    def add_get_latest_world_channel_message(self, callback):
        self.add_callback(0, callback)

    def add_add_world_chat(self, callback):
        self.add_callback(1, callback)

    def add_world_channel_info(self, callback):
        self.add_callback(2, callback)

    def add_send_global_messages(self, callback):
        self.add_callback(3, callback)

    def add_add_clique_chat(self, callback):
        self.add_callback(4, callback)

    def add_get_latest_clique_messages(self, callback):
        self.add_callback(5, callback)

    def add_send_clique_message(self, callback):
        self.add_callback(6, callback)
