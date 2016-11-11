#!/usr/bin/evn python
# -*- coding: utf8 -*-


class BaseProtocol(object):
    def __init__(self, client):
        self._client = client

    def head_size(self):
        raise NotImplementedError

    def body_size(self,head):
        raise NotImplementedError

    def handle_head(self, head):
        """
        :param head:
        :return: body_size, body_handler
        """
        raise NotImplementedError

    def handler(self):
        raise NotImplementedError


class BaseModule(object):
    def decode(self, buff):
        raise NotImplementedError

    def add_callback(self, action, callback):
        raise NotImplementedError

    def clear_callback(self):
        raise NotImplementedError