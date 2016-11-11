# /usr/bin/evn python
# -*- coding: utf8 -*-
__author__ = 'su21'

import socket
from functools import partial
import binascii

from tornado.ioloop import IOLoop
from tornado.iostream import IOStream

import tornado.gen


class Client(object):

    def __init__(self, host, port, protocol, name, session, dial=True):
        self._host = host
        self._port = port
        self._protocol = protocol
        self._name = name
        self._session = session
        self._running = False
        self._socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM, 0)
        self._connection = IOStream(self._socket)
        if dial:
            self._connection.connect((host, port), self.on_dial)

    def dial(self):
        if self._running:
            return
        self._connection.connect((self._host, self._port), self.on_dial)
        self._read_loop()

    def on_dial(self):
        """
        subclass should implement this method, example:
        def on_dial(self):
            self._connection.send(msg)
            self._read_loop()
        :return: None
        """
        self._running = True

    def read_loop(self):
        #self._connection.read_bytes(self._protocol.head_size(), callback=self._debug)
        self._connection.read_bytes(self._protocol.head_size(), callback=self._handle_head)

    def _handle_head(self, head):
        receive_bytes, body_size = self._protocol.handle_head(self._session, head)
        # if self._session.receive_bytes >= receive_bytes:
        #     # TODO it is an error
        #     pass
        #     #return
        self._session.receive_bytes = receive_bytes

        def handler(body):
            self._protocol.handle(self._session, body)
            self.read_loop()

        if body_size > 0:
            self._connection.read_bytes(body_size, callback=handler)

    def send_raw(self, data):
        self._connection.write_to_fd(self._protocol.encode(self._session, data))

    def send(self, msg):
        raw_data = self._protocol.encode(self._session, msg.encode())
        self._connection.write_to_fd(raw_data)

