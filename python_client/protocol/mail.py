#!/usr/bin/evn python
#  -*- coding: utf8 -*-

import struct

import interface

class GetListUp(object):
    _module = 12
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
    _module = 12
    _action = 0

    def __init__(self):
        pass
        self.get_heart_num = 0
        self.mails = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<l", self.get_heart_num))
        buff.extend(struct.pack('<H', len(self.mails)))
        for item in self.mails:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.get_heart_num = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        _mails_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        for i in range(_mails_size):
            obj = GetListDownMails()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.mails.append(obj)

    def size(self):
        size = 6
        for item in self.mails:
            size += item.size()
        return size


class GetListDownMails(object):
    def __init__(self):
        pass
        self.id = 0
        self.mail_id = 0
        self.state = 0
        self.priority = 0
        self.send_time = 0
        self.expire_time = 0
        self.parameters = ''
        self.title = ''
        self.content = ''
        self.attachments = []

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<l", self.mail_id))
        buff.extend(struct.pack("<b", self.state))
        buff.extend(struct.pack("<b", self.priority))
        buff.extend(struct.pack("<q", self.send_time))
        buff.extend(struct.pack("<q", self.expire_time))
        buff.extend(struct.pack('<H', len(self.parameters)))
        buff.extend(self.parameters)
        buff.extend(struct.pack('<H', len(self.title)))
        buff.extend(self.title)
        buff.extend(struct.pack('<H', len(self.content)))
        buff.extend(self.content)
        buff.extend(struct.pack('<B', len(self.attachments)))
        for item in self.attachments:
            buff.extend(item.encode())
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.mail_id = struct.unpack_from("<l", raw_msg, idx)[0]
        idx += 4

        self.state = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.priority = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.send_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.expire_time = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        _parameters_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.parameters = str(raw_msg[idx:idx+_parameters_size])
        idx += _parameters_size

        _title_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.title = str(raw_msg[idx:idx+_title_size])
        idx += _title_size

        _content_size = struct.unpack_from("<H", raw_msg, idx)[0]
        idx += 2
        self.content = str(raw_msg[idx:idx+_content_size])
        idx += _content_size

        _attachments_size = struct.unpack_from("<B", raw_msg, idx)[0]
        idx += 1
        for i in range(_attachments_size):
            obj = GetListDownMailsAttachments()
            obj.decode(raw_msg[idx:])
            idx += obj.size()
            self.attachments.append(obj)

    def size(self):
        size = 37
        size += len(self.parameters)
        size += len(self.title)
        size += len(self.content)
        for item in self.attachments:
            size += item.size()
        return size


class GetListDownMailsAttachments(object):
    def __init__(self):
        pass
        self.id = 0
        self.attachment_type = 0
        self.item_id = 0
        self.item_num = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.id))
        buff.extend(struct.pack("<b", self.attachment_type))
        buff.extend(struct.pack("<h", self.item_id))
        buff.extend(struct.pack("<q", self.item_num))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8

        self.attachment_type = struct.unpack_from("<b", raw_msg, idx)[0]
        idx += 1

        self.item_id = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.item_num = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 19
        

class ReadUp(object):
    _module = 12
    _action = 1

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
        

class ReadDown(object):
    _module = 12
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
        

class TakeAttachmentUp(object):
    _module = 12
    _action = 2

    def __init__(self):
        pass
        self.attachment_id = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<q", self.attachment_id))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.attachment_id = struct.unpack_from("<q", raw_msg, idx)[0]
        idx += 8


    @staticmethod
    def size():
        return 10
        

class TakeAttachmentDown(object):
    _module = 12
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
        

class GetInfosUp(object):
    _module = 12
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
        

class GetInfosDown(object):
    _module = 12
    _action = 3

    def __init__(self):
        pass
        self.new_mail_num = 0
        self.unread_num = 0
        self.total = 0

    def encode(self):
        buff = bytearray()
        buff.extend(struct.pack('<B', self._module))
        buff.extend(struct.pack('<B', self._action))
        buff.extend(struct.pack("<h", self.new_mail_num))
        buff.extend(struct.pack("<h", self.unread_num))
        buff.extend(struct.pack("<h", self.total))
        return buff

    def decode(self, raw_msg):
        idx = 0

        self.new_mail_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.unread_num = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2

        self.total = struct.unpack_from("<h", raw_msg, idx)[0]
        idx += 2


    @staticmethod
    def size():
        return 6
        

class RequestGlobalMailUp(object):
    _module = 12
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
        

class RequestGlobalMailDown(object):
    _module = 12
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
        

class MailModule(interface.BaseModule):
    decoder_map = {
        0: GetListDown, 
        1: ReadDown, 
        2: TakeAttachmentDown, 
        3: GetInfosDown, 
        4: RequestGlobalMailDown, 
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

    def add_read(self, callback):
        self.add_callback(1, callback)

    def add_take_attachment(self, callback):
        self.add_callback(2, callback)

    def add_get_infos(self, callback):
        self.add_callback(3, callback)

    def add_request_global_mail(self, callback):
        self.add_callback(4, callback)
