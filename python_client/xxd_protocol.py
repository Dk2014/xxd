# /usr/bin/evn python
# -*- coding: utf8 -*-
__author__ = 'su21'
15 / 6 / 22

import struct

import protocol


class Protocol(protocol.BaseProtocol):
    _head_size = 8

    def head_size(self):
        return self._head_size

    @staticmethod
    def handle_head(session, head):
        head_size, receive_bytes  = struct.unpack("<II", str(head))
        return receive_bytes, head_size

    def handle(self, session, body):
        module_id = ord(body[0])
        self.handler_body(module_id, session, body[1:])

    @staticmethod
    def encode(session, data):
        session.send_bytes += len(data)
        code = session.send_bytes  # send_bytes to bytes
        buff = bytearray()
        #第4字节存的是发送字节数
        buff.extend(struct.pack("<I", code))
        #高4字节存包长度 包长度于与 send_bytes 做异或操作
        size = len(data) ^ code
        buff.extend(struct.pack("<I", size))
        data2 = bytearray()
        for x in data:
            data2.append(x^(code&0xff))
        buff.extend(data2)
        return buff



