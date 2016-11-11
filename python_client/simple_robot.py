# /usr/bin/evn python
# -*- coding: utf8 -*-
__author__ = 'su21'
15 / 6 / 22

from tornado.ioloop import IOLoop
import tornado.gen
import tornado.options

tornado.options.parse_command_line()


from net.client import Client
from protocol import player, town
from xxd_protocol import Protocol

import time
import json
import random


class GameSrvClient(Client):
    def on_dial(self):
        msg = player.FromPlatformLoginUp()
        msg.nickname = self._session.name
        msg.username = self._session.name
        msg.gsid = 91
        msg.role_id = 1
        msg.unixtime = int(time.time())
        msg.platform_type = 1
        # TODO set more attribute
        self.send(msg)
        self._running = True
        self.read_loop()


class HdSrvClient(Client):

    def on_dial(self, session):
        msg = player.GlobalLoginUp()
        msg.username = session.name
        msg.openid = session.name
        self.send(msg)


class Session(object):
    def __init__(self, name, proto):
        self.send_bytes = 0
        self.receive_bytes = 20141021
        self.name = name
        self.client = GameSrvClient('127.0.0.1', 8080, proto, name, self)

# some test below:
# 1. create protocol instance
# 2. register message handler
# 3. create session
# 4. start main loop

protocol = Protocol()


def town_move(session):
    info = town.MoveUp()
    info.to_x = random.randint(0, 700)+1200
    info.to_y = random.randint(0, 400) + 700
    print session.name, "move to", info.to_x, info.to_y
    session.client.send(info)
    IOLoop.instance().add_timeout(time.time()+2, town_move, session)


def town_enter_cb(session, msg):
    town_move(session)

protocol.town.add_enter(town_enter_cb)


def player_info_cb(session, msg):
    print session.name, "receive info", msg, json.dumps(msg.__dict__)
    enter_town_up = town.EnterUp()
    enter_town_up.town_id = 1
    session.client.send(enter_town_up)

protocol.player.add_info(player_info_cb)


def player_login_cb(session, msg):
    info = player.InfoUp()
    session.client.send(info)
    print session.name, "receive login rsp", msg, json.dumps(msg.__dict__)

protocol.player.add_from_platform_login(player_login_cb)

for x in range(1000):
    Session('pyrobot' + str(x), protocol)


IOLoop.instance().start()
