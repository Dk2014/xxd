#!/usr/bin/evn python
# -*- coding: utf8 -*-


import player
import town
import team
import role
import mission
import skill
import battle
import item
import notify
import ghost
import sword_soul
import mail
import quest
import friend
import tower
import multi_level
import battle_pet
import announcement
import arena
import vip
import trader
import daily_sign_in
import rainbow
import event
import fashion
import push_notify
import meditation
import pet_virtual_env
import channel
import driving_sword
import totem
import money_tree
import clique
import clique_building
import clique_quest
import clique_escort
import awaken
import draw
import server_info
import debug


class BaseProtocol(object):
    def head_size(self):
        raise NotImplementedError

    def handle_head(self, session, head):
        """return body_size, receive_bytes"""
        raise NotImplementedError

    def handler_body(self, module_id, session, data):
        pass

        if module_id == 0:
            action_id = ord(data[0])
            msg = self.player.decode(data)
            callbacks = self.player.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 1:
            action_id = ord(data[0])
            msg = self.town.decode(data)
            callbacks = self.town.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 2:
            action_id = ord(data[0])
            msg = self.team.decode(data)
            callbacks = self.team.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 3:
            action_id = ord(data[0])
            msg = self.role.decode(data)
            callbacks = self.role.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 4:
            action_id = ord(data[0])
            msg = self.mission.decode(data)
            callbacks = self.mission.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 5:
            action_id = ord(data[0])
            msg = self.skill.decode(data)
            callbacks = self.skill.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 6:
            action_id = ord(data[0])
            msg = self.battle.decode(data)
            callbacks = self.battle.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 7:
            action_id = ord(data[0])
            msg = self.item.decode(data)
            callbacks = self.item.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 8:
            action_id = ord(data[0])
            msg = self.notify.decode(data)
            callbacks = self.notify.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 9:
            action_id = ord(data[0])
            msg = self.ghost.decode(data)
            callbacks = self.ghost.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 10:
            action_id = ord(data[0])
            msg = self.sword_soul.decode(data)
            callbacks = self.sword_soul.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 12:
            action_id = ord(data[0])
            msg = self.mail.decode(data)
            callbacks = self.mail.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 13:
            action_id = ord(data[0])
            msg = self.quest.decode(data)
            callbacks = self.quest.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 14:
            action_id = ord(data[0])
            msg = self.friend.decode(data)
            callbacks = self.friend.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 15:
            action_id = ord(data[0])
            msg = self.tower.decode(data)
            callbacks = self.tower.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 16:
            action_id = ord(data[0])
            msg = self.multi_level.decode(data)
            callbacks = self.multi_level.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 17:
            action_id = ord(data[0])
            msg = self.battle_pet.decode(data)
            callbacks = self.battle_pet.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 18:
            action_id = ord(data[0])
            msg = self.announcement.decode(data)
            callbacks = self.announcement.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 19:
            action_id = ord(data[0])
            msg = self.arena.decode(data)
            callbacks = self.arena.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 20:
            action_id = ord(data[0])
            msg = self.vip.decode(data)
            callbacks = self.vip.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 21:
            action_id = ord(data[0])
            msg = self.trader.decode(data)
            callbacks = self.trader.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 22:
            action_id = ord(data[0])
            msg = self.daily_sign_in.decode(data)
            callbacks = self.daily_sign_in.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 23:
            action_id = ord(data[0])
            msg = self.rainbow.decode(data)
            callbacks = self.rainbow.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 24:
            action_id = ord(data[0])
            msg = self.event.decode(data)
            callbacks = self.event.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 25:
            action_id = ord(data[0])
            msg = self.fashion.decode(data)
            callbacks = self.fashion.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 26:
            action_id = ord(data[0])
            msg = self.push_notify.decode(data)
            callbacks = self.push_notify.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 27:
            action_id = ord(data[0])
            msg = self.meditation.decode(data)
            callbacks = self.meditation.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 28:
            action_id = ord(data[0])
            msg = self.pet_virtual_env.decode(data)
            callbacks = self.pet_virtual_env.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 29:
            action_id = ord(data[0])
            msg = self.channel.decode(data)
            callbacks = self.channel.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 30:
            action_id = ord(data[0])
            msg = self.driving_sword.decode(data)
            callbacks = self.driving_sword.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 31:
            action_id = ord(data[0])
            msg = self.totem.decode(data)
            callbacks = self.totem.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 32:
            action_id = ord(data[0])
            msg = self.money_tree.decode(data)
            callbacks = self.money_tree.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 33:
            action_id = ord(data[0])
            msg = self.clique.decode(data)
            callbacks = self.clique.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 34:
            action_id = ord(data[0])
            msg = self.clique_building.decode(data)
            callbacks = self.clique_building.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 35:
            action_id = ord(data[0])
            msg = self.clique_quest.decode(data)
            callbacks = self.clique_quest.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 36:
            action_id = ord(data[0])
            msg = self.clique_escort.decode(data)
            callbacks = self.clique_escort.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 38:
            action_id = ord(data[0])
            msg = self.awaken.decode(data)
            callbacks = self.awaken.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 50:
            action_id = ord(data[0])
            msg = self.draw.decode(data)
            callbacks = self.draw.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 88:
            action_id = ord(data[0])
            msg = self.server_info.decode(data)
            callbacks = self.server_info.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

        if module_id == 99:
            action_id = ord(data[0])
            msg = self.debug.decode(data)
            callbacks = self.debug.receive_callback.get(action_id, [])
            for cb in callbacks:
                cb(session, msg)
            return

    def __init__(self):
        self.player = player.PlayerModule()
        self.town = town.TownModule()
        self.team = team.TeamModule()
        self.role = role.RoleModule()
        self.mission = mission.MissionModule()
        self.skill = skill.SkillModule()
        self.battle = battle.BattleModule()
        self.item = item.ItemModule()
        self.notify = notify.NotifyModule()
        self.ghost = ghost.GhostModule()
        self.sword_soul = sword_soul.SwordSoulModule()
        self.mail = mail.MailModule()
        self.quest = quest.QuestModule()
        self.friend = friend.FriendModule()
        self.tower = tower.TowerModule()
        self.multi_level = multi_level.MultiLevelModule()
        self.battle_pet = battle_pet.BattlePetModule()
        self.announcement = announcement.AnnouncementModule()
        self.arena = arena.ArenaModule()
        self.vip = vip.VipModule()
        self.trader = trader.TraderModule()
        self.daily_sign_in = daily_sign_in.DailySignInModule()
        self.rainbow = rainbow.RainbowModule()
        self.event = event.EventModule()
        self.fashion = fashion.FashionModule()
        self.push_notify = push_notify.PushNotifyModule()
        self.meditation = meditation.MeditationModule()
        self.pet_virtual_env = pet_virtual_env.PetVirtualEnvModule()
        self.channel = channel.ChannelModule()
        self.driving_sword = driving_sword.DrivingSwordModule()
        self.totem = totem.TotemModule()
        self.money_tree = money_tree.MoneyTreeModule()
        self.clique = clique.CliqueModule()
        self.clique_building = clique_building.CliqueBuildingModule()
        self.clique_quest = clique_quest.CliqueQuestModule()
        self.clique_escort = clique_escort.CliqueEscortModule()
        self.awaken = awaken.AwakenModule()
        self.draw = draw.DrawModule()
        self.server_info = server_info.ServerInfoModule()
        self.debug = debug.DebugModule()
