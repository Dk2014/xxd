package battle

import "math/rand"

func createRoleSkill(id int) *Skill {
	switch id {
	case 1: return role_skill_1 // 落云
	case 2: return role_skill_2 // 裂地斩
	case 4: return role_skill_4 // 飞焰穿云
	case 6: return role_skill_6 // 甲溃
	case 8: return role_skill_8 // 大风咒
	case 9: return role_skill_9 // 咆哮利爪
	case 10: return role_skill_10 // 凶猛撕咬
	case 11: return role_skill_11 // 冰烈
	case 12: return role_skill_12 // 冰烈横向
	case 13: return role_skill_13 // 冰烈纵向
	case 14: return role_skill_14 // 冰烈全体
	case 15: return role_skill_15 // 火烈
	case 16: return role_skill_16 // 火烈横向
	case 17: return role_skill_17 // 火烈纵向
	case 18: return role_skill_18 // 火烈全体
	case 19: return role_skill_19 // 风烈
	case 20: return role_skill_20 // 风烈横向
	case 21: return role_skill_21 // 风烈纵向
	case 22: return role_skill_22 // 风烈全体
	case 23: return role_skill_23 // 雷烈
	case 24: return role_skill_24 // 雷烈横向
	case 25: return role_skill_25 // 雷烈纵向
	case 26: return role_skill_26 // 雷烈全体
	case 27: return role_skill_27 // 土烈
	case 28: return role_skill_28 // 土烈横向
	case 29: return role_skill_29 // 土烈纵向
	case 30: return role_skill_30 // 土烈全体
	case 31: return role_skill_31 // 毒烈
	case 32: return role_skill_32 // 毒烈横向
	case 33: return role_skill_33 // 毒烈纵向
	case 34: return role_skill_34 // 毒烈全体
	case 35: return role_skill_35 // 多连斩
	case 36: return role_skill_36 // 力劈华山
	case 37: return role_skill_37 // 白莲横江
	case 38: return role_skill_38 // 横扫千军
	case 39: return role_skill_39 // 乾坤刀气
	case 40: return role_skill_40 // 三千洛水剑
	case 41: return role_skill_41 // 死亡标记
	case 42: return role_skill_42 // 万箭穿心
	case 43: return role_skill_43 // 狮吼功
	case 44: return role_skill_44 // 野蛮冲撞(怪物用)
	case 45: return role_skill_45 // 如殂随行
	case 46: return role_skill_46 // 驱散
	case 49: return role_skill_49 // 青竹咒
	case 51: return role_skill_51 // 墨画影狼
	case 52: return role_skill_52 // 墨画巫雀
	case 89: return role_skill_89 // 坠星击
	case 98: return role_skill_98 // 天剑
	case 108: return role_skill_108 // 治疗
	case 109: return role_skill_109 // 增益
	case 998: return role_skill_998 // 破甲眩晕
	case 999: return role_skill_999 // 审判惊雷
	case 1000: return role_skill_1000 // 铜墙铁壁
	case 1001: return role_skill_1001 // 风之领域
	case 1002: return role_skill_1002 // 无限霸刀
	case 1003: return role_skill_1003 // 熔岩弹雨
	case 1004: return role_skill_1004 // 人鱼之歌
	case 1017: return role_skill_1017 // 振奋之击
	case 1020: return role_skill_1020 // 致胜千里
	case 1023: return role_skill_1023 // 万剑穿心
	case 1026: return role_skill_1026 // 阿修罗之怒
	case 1029: return role_skill_1029 // 命运连锁
	case 1032: return role_skill_1032 // 圣灵洛水
	case 1035: return role_skill_1035 // 死亡阻击
	case 1038: return role_skill_1038 // 紫电刀芒
	case 1039: return role_skill_1039 // 斩杀
	case 1041: return role_skill_1041 // 聚气
	case 1042: return role_skill_1042 // 灵蛇之枪
	case 1043: return role_skill_1043 // 火凤燎原
	case 1047: return role_skill_1047 // 全体治疗
	case 1048: return role_skill_1048 // 青竹咒
	case 1049: return role_skill_1049 // 雨润
	case 1050: return role_skill_1050 // 圣白莲
	case 1051: return role_skill_1051 // 神魔封禁
	case 1052: return role_skill_1052 // 墨画巫雀
	case 1053: return role_skill_1053 // 墨画影狼
	case 1054: return role_skill_1054 // 祝福
	case 1055: return role_skill_1055 // 地狱烈焰
	case 1056: return role_skill_1056 // 万剑归一
	case 1058: return role_skill_1058 // 暗影食梦
	case 1061: return role_skill_1061 // 寒霜灵泉
	case 1064: return role_skill_1064 // 阎罗试炼
	case 1067: return role_skill_1067 // 雨润
	case 1070: return role_skill_1070 // 坠星击
	case 1074: return role_skill_1074 // 治愈之风
	case 1075: return role_skill_1075 // 刀盾
	case 1076: return role_skill_1076 // 审判惊雷附属
	case 1077: return role_skill_1077 // 风之领域附属
	case 1079: return role_skill_1079 // 铜墙铁壁附属
	case 1086: return role_skill_1086 // 熔岩弹雨附属
	case 1092: return role_skill_1092 // 风镰
	case 1094: return role_skill_1094 // 重击
	case 1095: return role_skill_1095 // 提笔
	case 1096: return role_skill_1096 // 刀芒
	case 1097: return role_skill_1097 // 大风咒
	case 1098: return role_skill_1098 // 振奋之击附属
	case 1101: return role_skill_1101 // 致胜千里附属
	case 1104: return role_skill_1104 // 万剑穿心附属
	case 1107: return role_skill_1107 // 圣灵洛水附属
	case 1110: return role_skill_1110 // 命运连锁附属
	case 1113: return role_skill_1113 // 灵蛇之枪附属
	case 1116: return role_skill_1116 // 暗影食梦附属
	case 1119: return role_skill_1119 // 寒霜灵泉附属
	case 1122: return role_skill_1122 // 火凤燎原附属
	case 1125: return role_skill_1125 // 阎罗试炼附属
	case 1134: return role_skill_1134 // 怒火冲天
	case 1135: return role_skill_1135 // 剑冲阴阳
	case 1136: return role_skill_1136 // 黄粱一梦
	case 1137: return role_skill_1137 // 蹑影追风
	case 1138: return role_skill_1138 // 野蛮冲撞
	case 1139: return role_skill_1139 // 铁布衫
	case 1140: return role_skill_1140 // 裂地斩2级
	case 1141: return role_skill_1141 // 偃月
	case 1142: return role_skill_1142 // 治愈之风2级
	case 1143: return role_skill_1143 // 藤甲术
	case 1144: return role_skill_1144 // 偃月2级
	case 1145: return role_skill_1145 // 墨画魂刃
	case 1146: return role_skill_1146 // 四象皆杀
	case 1147: return role_skill_1147 // 护盾
	case 1148: return role_skill_1148 // 雨润2级
	case 1149: return role_skill_1149 // 落云2级
	case 1150: return role_skill_1150 // 纯阳罡气
	case 1151: return role_skill_1151 // 斩击2级
	case 1153: return role_skill_1153 // 顺势斩
	case 1155: return role_skill_1155 // 破釜沉舟
	case 1156: return role_skill_1156 // 纳气归元
	case 1157: return role_skill_1157 // 狂乱之刃
	case 1158: return role_skill_1158 // 护盾破坏
	case 1159: return role_skill_1159 // 刀盾2级
	case 1164: return role_skill_1164 // 魂力果实
	case 1165: return role_skill_1165 // 青竹咒2级
	case 1167: return role_skill_1167 // 墨画影狼2级
	case 1168: return role_skill_1168 // 墨画巫雀2级
	case 1169: return role_skill_1169 // 风卷尘生
	case 1170: return role_skill_1170 // 风镰2级
	case 1171: return role_skill_1171 // 妙法莲华
	case 1173: return role_skill_1173 // 魅影之爪
	case 1174: return role_skill_1174 // 魅影之爪附属
	case 1179: return role_skill_1179 // 横扫千军
	case 1185: return role_skill_1185 // 碎铁之刃
	case 1191: return role_skill_1191 // 召唤阴影
	case 1193: return role_skill_1193 // 寒冰护体
	case 1194: return role_skill_1194 // 寒冰护体附属
	case 1199: return role_skill_1199 // 瞒天过海
	case 1205: return role_skill_1205 // 召唤堕落竹筒精
	case 1206: return role_skill_1206 // 召唤治愈莲藕精
	case 1207: return role_skill_1207 // 召唤堕落燃魁
	case 1208: return role_skill_1208 // 召唤暴怒剑之守卫
	case 1209: return role_skill_1209 // 召唤暴怒败亡之剑
	case 1210: return role_skill_1210 // 召唤调皮火灵
	case 1211: return role_skill_1211 // 召唤调皮豆灵
	case 1212: return role_skill_1212 // 召唤调皮金翅飞鸾
	case 1213: return role_skill_1213 // 龙息术
	case 1214: return role_skill_1214 // 泰山压顶
	case 1215: return role_skill_1215 // 浮生入梦
	case 1216: return role_skill_1216 // 斩击
	case 1218: return role_skill_1218 // 霸拳盖天
	case 1219: return role_skill_1219 // 召唤调皮雷兽
	case 1220: return role_skill_1220 // 召唤顽皮灵剑
	case 1221: return role_skill_1221 // 召唤淘气火灵
	case 1222: return role_skill_1222 // 召唤顽皮飞鹏
	case 1223: return role_skill_1223 // 十字斩
	case 1224: return role_skill_1224 // 风咒
	case 1225: return role_skill_1225 // 聚气蓄力
	case 1226: return role_skill_1226 // 断岳
	case 1227: return role_skill_1227 // 激励
	case 1228: return role_skill_1228 // 生生不息·怪
	case 1229: return role_skill_1229 // 生生不息3级
	case 1230: return role_skill_1230 // 生生不息4级
	case 1231: return role_skill_1231 // 生生不息5级
	case 1232: return role_skill_1232 // 怒火冲天·怪
	case 1233: return role_skill_1233 // 怒火冲天3级
	case 1234: return role_skill_1234 // 怒火冲天4级
	case 1235: return role_skill_1235 // 怒火冲天5级
	case 1236: return role_skill_1236 // 野蛮冲撞·怪
	case 1237: return role_skill_1237 // 野蛮冲撞3级
	case 1238: return role_skill_1238 // 野蛮冲撞4级
	case 1239: return role_skill_1239 // 野蛮冲撞5级
	case 1240: return role_skill_1240 // 剑冲阴阳·怪
	case 1241: return role_skill_1241 // 剑冲阴阳3级
	case 1242: return role_skill_1242 // 剑冲阴阳4级
	case 1243: return role_skill_1243 // 剑冲阴阳5级
	case 1244: return role_skill_1244 // 黄粱一梦·怪
	case 1245: return role_skill_1245 // 黄粱一梦3级
	case 1246: return role_skill_1246 // 黄粱一梦4级
	case 1247: return role_skill_1247 // 黄粱一梦5级
	case 1248: return role_skill_1248 // 蹑影追风·怪
	case 1249: return role_skill_1249 // 蹑影追风3级
	case 1250: return role_skill_1250 // 蹑影追风4级
	case 1251: return role_skill_1251 // 蹑影追风5级
	case 1252: return role_skill_1252 // 雷动九天
	case 1253: return role_skill_1253 // 雷动九天·怪
	case 1254: return role_skill_1254 // 雷动九天3级
	case 1255: return role_skill_1255 // 雷动九天4级
	case 1256: return role_skill_1256 // 雷动九天5级
	case 1257: return role_skill_1257 // 妙笔生花
	case 1258: return role_skill_1258 // 妙笔生花·怪
	case 1259: return role_skill_1259 // 妙笔生花3级
	case 1260: return role_skill_1260 // 妙笔生花4级
	case 1261: return role_skill_1261 // 妙笔生花5级
	case 1262: return role_skill_1262 // 风驰云卷
	case 1263: return role_skill_1263 // 风驰云卷·怪
	case 1264: return role_skill_1264 // 风驰云卷3级
	case 1265: return role_skill_1265 // 风驰云卷4级
	case 1266: return role_skill_1266 // 风驰云卷5级
	case 1267: return role_skill_1267 // 神魅鬼目
	case 1268: return role_skill_1268 // 神魅鬼目·怪
	case 1269: return role_skill_1269 // 神魅鬼目3级
	case 1270: return role_skill_1270 // 神魅鬼目4级
	case 1271: return role_skill_1271 // 神魅鬼目5级
	case 1272: return role_skill_1272 // 生生不息
	case 1273: return role_skill_1273 // 剑十
	case 1274: return role_skill_1274 // 召唤开心莲藕精
	case 1275: return role_skill_1275 // 英勇
	case 1276: return role_skill_1276 // 霸刀
	case 1277: return role_skill_1277 // 轰击
	case 1278: return role_skill_1278 // 阿修罗之怒附属
	case 1281: return role_skill_1281 // 银光落刃
	case 1284: return role_skill_1284 // 银光落刃
	case 1287: return role_skill_1287 // 式神炎舞
	case 1288: return role_skill_1288 // 式神炎舞附属
	case 1293: return role_skill_1293 // 猛龙断空
	case 1294: return role_skill_1294 // 猛龙断空附属
	case 1299: return role_skill_1299 // 影舞
	case 1300: return role_skill_1300 // 流萤斩
	case 1301: return role_skill_1301 // 烟雾弹
	case 1302: return role_skill_1302 // 暗影火袭
	case 1303: return role_skill_1303 // 凌云斩
	case 1304: return role_skill_1304 // 风之守护
	case 1305: return role_skill_1305 // 大风咒2级
	case 1306: return role_skill_1306 // 风卷尘生2级
	case 1307: return role_skill_1307 // 妙法莲华2级
	case 1308: return role_skill_1308 // 一滴入魂
	case 1309: return role_skill_1309 // 纯阳罡气2级
	case 1310: return role_skill_1310 // 护盾破坏2级
	case 1311: return role_skill_1311 // 藤甲术2级
	case 1312: return role_skill_1312 // 灵竹咒
	case 1313: return role_skill_1313 // 神魔封禁
	case 1314: return role_skill_1314 // 墨画魂刃2级
	case 1315: return role_skill_1315 // 火炎咒
	case 1316: return role_skill_1316 // 神魅鬼目
	case 1317: return role_skill_1317 // 召唤唐家护卫
	case 1318: return role_skill_1318 // 落云3级
	case 1319: return role_skill_1319 // 落云4级
	case 1320: return role_skill_1320 // 凌云斩2级
	case 1321: return role_skill_1321 // 落云5级
	case 1322: return role_skill_1322 // 落云6级
	case 1323: return role_skill_1323 // 落云7级
	case 1324: return role_skill_1324 // 落云8级
	case 1325: return role_skill_1325 // 风镰3级
	case 1326: return role_skill_1326 // 风镰4级
	case 1327: return role_skill_1327 // 风镰5级
	case 1328: return role_skill_1328 // 风镰6级
	case 1329: return role_skill_1329 // 风镰7级
	case 1330: return role_skill_1330 // 风镰8级
	case 1331: return role_skill_1331 // 风镰9级
	case 1332: return role_skill_1332 // 风镰10级
	case 1333: return role_skill_1333 // 斩击3级
	case 1334: return role_skill_1334 // 斩击4级
	case 1335: return role_skill_1335 // 斩击5级
	case 1336: return role_skill_1336 // 斩击6级
	case 1337: return role_skill_1337 // 斩击7级
	case 1338: return role_skill_1338 // 斩击8级
	case 1339: return role_skill_1339 // 斩击9级
	case 1340: return role_skill_1340 // 斩击10级
	case 1341: return role_skill_1341 // 重击2级
	case 1342: return role_skill_1342 // 重击3级
	case 1343: return role_skill_1343 // 重击4级
	case 1344: return role_skill_1344 // 重击5级
	case 1345: return role_skill_1345 // 重击6级
	case 1346: return role_skill_1346 // 重击7级
	case 1347: return role_skill_1347 // 重击8级
	case 1348: return role_skill_1348 // 重击9级
	case 1349: return role_skill_1349 // 重击10级
	case 1350: return role_skill_1350 // 提笔2级
	case 1351: return role_skill_1351 // 提笔3级
	case 1352: return role_skill_1352 // 提笔4级
	case 1353: return role_skill_1353 // 提笔5级
	case 1354: return role_skill_1354 // 提笔6级
	case 1355: return role_skill_1355 // 提笔7级
	case 1356: return role_skill_1356 // 提笔8级
	case 1357: return role_skill_1357 // 提笔9级
	case 1358: return role_skill_1358 // 提笔10级
	case 1359: return role_skill_1359 // 影舞2级
	case 1360: return role_skill_1360 // 影舞3级
	case 1361: return role_skill_1361 // 影舞4级
	case 1362: return role_skill_1362 // 影舞5级
	case 1363: return role_skill_1363 // 影舞6级
	case 1364: return role_skill_1364 // 影舞7级
	case 1365: return role_skill_1365 // 影舞8级
	case 1366: return role_skill_1366 // 影舞9级
	case 1367: return role_skill_1367 // 影舞10级
	case 1368: return role_skill_1368 // 追骨吸元
	case 1369: return role_skill_1369 // 流萤斩2级
	case 1370: return role_skill_1370 // 烟雾弹2级
	case 1371: return role_skill_1371 // 长拳
	case 1372: return role_skill_1372 // 雷行者
	case 1373: return role_skill_1373 // 木行者
	case 1374: return role_skill_1374 // 横扫
	case 1375: return role_skill_1375 // 斩铁剑
	case 1376: return role_skill_1376 // 真元护体
	case 1377: return role_skill_1377 // 寒霜
	case 1378: return role_skill_1378 // 冰锥术
	case 1379: return role_skill_1379 // 冰心诀
	case 1380: return role_skill_1380 // 突袭
	case 1381: return role_skill_1381 // 金蛇锥
	case 1382: return role_skill_1382 // 闪光弹
	case 1383: return role_skill_1383 // 长拳2级
	case 1384: return role_skill_1384 // 长拳3级
	case 1385: return role_skill_1385 // 长拳4级
	case 1386: return role_skill_1386 // 长拳5级
	case 1387: return role_skill_1387 // 长拳6级
	case 1388: return role_skill_1388 // 长拳7级
	case 1389: return role_skill_1389 // 长拳8级
	case 1390: return role_skill_1390 // 长拳9级
	case 1391: return role_skill_1391 // 长拳10级
	case 1392: return role_skill_1392 // 横扫2级
	case 1393: return role_skill_1393 // 横扫3级
	case 1394: return role_skill_1394 // 横扫4级
	case 1395: return role_skill_1395 // 横扫5级
	case 1396: return role_skill_1396 // 横扫6级
	case 1397: return role_skill_1397 // 横扫7级
	case 1398: return role_skill_1398 // 横扫8级
	case 1399: return role_skill_1399 // 横扫9级
	case 1400: return role_skill_1400 // 横扫10级
	case 1401: return role_skill_1401 // 寒霜2级
	case 1402: return role_skill_1402 // 寒霜3级
	case 1403: return role_skill_1403 // 寒霜4级
	case 1404: return role_skill_1404 // 寒霜5级
	case 1405: return role_skill_1405 // 寒霜6级
	case 1406: return role_skill_1406 // 寒霜7级
	case 1407: return role_skill_1407 // 寒霜8级
	case 1408: return role_skill_1408 // 寒霜9级
	case 1409: return role_skill_1409 // 寒霜10级
	case 1410: return role_skill_1410 // 突袭2级
	case 1411: return role_skill_1411 // 突袭3级
	case 1412: return role_skill_1412 // 突袭4级
	case 1413: return role_skill_1413 // 突袭5级
	case 1414: return role_skill_1414 // 突袭6级
	case 1415: return role_skill_1415 // 突击7级
	case 1416: return role_skill_1416 // 突袭8级
	case 1417: return role_skill_1417 // 突袭9级
	case 1418: return role_skill_1418 // 突袭10级
	case 1419: return role_skill_1419 // 断岳
	case 1420: return role_skill_1420 // 无限霸刀附属
	case 1421: return role_skill_1421 // 人鱼之歌附属
	case 1422: return role_skill_1422 // 横扫千军附属
	case 1423: return role_skill_1423 // 碎铁之刃附属
	case 1424: return role_skill_1424 // 瞒天过海附属
	case 1425: return role_skill_1425 // 银光落刃附属
	case 1426: return role_skill_1426 // 流荧斩
	case 1427: return role_skill_1427 // 烟雾弹
	case 1428: return role_skill_1428 // 灵竹咒
	case 1429: return role_skill_1429 // 追骨吸元
	case 1430: return role_skill_1430 // 凌云斩
	case 1431: return role_skill_1431 // 顺势斩
	case 1432: return role_skill_1432 // 天剑
	case 1434: return role_skill_1434 // 召唤堕落毒蛇
	case 1435: return role_skill_1435 // 召唤堕落黑翼巨蝠
	case 1436: return role_skill_1436 // 召唤堕落燃魁
	case 1437: return role_skill_1437 // 召唤堕落狼牙血杀
	case 1438: return role_skill_1438 // 召唤堕落重刀血杀
	case 1439: return role_skill_1439 // 召唤堕落野猪、堕落穿
	case 1441: return role_skill_1441 // 召唤堕落火豹
	case 1442: return role_skill_1442 // 召唤堕落不死亡魂
	case 1443: return role_skill_1443 // 召唤堕落僵尸
	case 1444: return role_skill_1444 // 召唤堕落僵尸2
	case 1445: return role_skill_1445 // 召唤堕落拳之守卫
	case 1446: return role_skill_1446 // 召唤堕落黑狼、堕落野
	case 1447: return role_skill_1447 // 召唤堕落玉石守卫
	case 1448: return role_skill_1448 // 召唤堕落撼地兽
	case 1449: return role_skill_1449 // 召唤堕落雷兽
	case 1450: return role_skill_1450 // 召唤堕落冥界僵尸
	case 1451: return role_skill_1451 // 召唤堕落弓之守卫
	case 1452: return role_skill_1452 // 召唤堕落钩镰血杀
	case 1453: return role_skill_1453 // 召唤堕落残风
	case 1454: return role_skill_1454 // 召唤堕落毒兽、堕落毒
	case 1455: return role_skill_1455 // 召唤堕落阴毒血士、堕
	case 1456: return role_skill_1456 // 召唤堕落利爪蝙蝠
	case 1457: return role_skill_1457 // 召唤堕落血蝙蝠杀手
	case 1458: return role_skill_1458 // 召唤堕落水鬼
	case 1459: return role_skill_1459 // 召唤堕落幕府重甲武士
	case 1460: return role_skill_1460 // 召唤堕落幕府长弓武士
	case 1461: return role_skill_1461 // 召唤堕落血杀武士
	case 1462: return role_skill_1462 // 召唤堕落甲贺忍者
	case 1463: return role_skill_1463 // 召唤堕落血妖
	case 1464: return role_skill_1464 // 召唤堕落血巫
	case 1465: return role_skill_1465 // 召唤堕落唐家护卫
	case 1466: return role_skill_1466 // 召唤堕落唐家女护卫
	case 1467: return role_skill_1467 // 召唤堕落龙虎门人
	case 1468: return role_skill_1468 // 召唤堕落恶犬
	case 1469: return role_skill_1469 // 召唤堕落剑影
	case 1470: return role_skill_1470 // 召唤堕落玄音寺武僧
	case 1471: return role_skill_1471 // 召唤堕落刀殿弟子
	case 1472: return role_skill_1472 // 召唤堕落灯笼怪
	case 1473: return role_skill_1473 // 召唤堕落剑之守卫
	case 1474: return role_skill_1474 // 气疗术
	case 1475: return role_skill_1475 // 毒伤
	case 1476: return role_skill_1476 // 分裂箭
	case 1477: return role_skill_1477 // 破甲利爪
	case 1478: return role_skill_1478 // 背刺
	case 1479: return role_skill_1479 // 腐蚀攻击1
	case 1480: return role_skill_1480 // 居合
	case 1481: return role_skill_1481 // 会心一击
	case 1482: return role_skill_1482 // 达摩棍
	case 1483: return role_skill_1483 // 腐蚀剧毒
	case 1484: return role_skill_1484 // 神魅鬼目（群体）
	case 1485: return role_skill_1485 // 落云斩
	case 1486: return role_skill_1486 // 嗜血
	case 1487: return role_skill_1487 // 嘲讽
	case 1488: return role_skill_1488 // 腐蚀攻击2
	case 1489: return role_skill_1489 // 腐蚀攻击3
	case 1490: return role_skill_1490 // 人鱼之歌
	case 1491: return role_skill_1491 // 斩铁剑
	case 1492: return role_skill_1492 // 冰锥术
	case 1493: return role_skill_1493 // 金蛇锥
	case 1494: return role_skill_1494 // 闪光弹
	case 1495: return role_skill_1495 // 木行者
	case 1496: return role_skill_1496 // 真元护体
	case 1497: return role_skill_1497 // 冰心诀
	case 1498: return role_skill_1498 // 雷行者
	case 1499: return role_skill_1499 // 召唤机灵飞鹏
	case 1500: return role_skill_1500 // 血爆
	case 1501: return role_skill_1501 // 血爆连锁
	case 1502: return role_skill_1502 // 血祭
	case 1503: return role_skill_1503 // 闪惊雷
	case 1504: return role_skill_1504 // 毒爆
	case 1505: return role_skill_1505 // 铁拳
	case 1506: return role_skill_1506 // 钢拳
	case 1507: return role_skill_1507 // 钛合金拳
	case 1508: return role_skill_1508 // 暗影新星
	case 1509: return role_skill_1509 // 暗火之雨
	case 1510: return role_skill_1510 // 地突刺
	case 1511: return role_skill_1511 // 魔瞳术
	case 1512: return role_skill_1512 // 火焰冲击
	case 1513: return role_skill_1513 // 业火
	case 1514: return role_skill_1514 // 腐蚀剧毒2
	case 1515: return role_skill_1515 // 爆怒
	case 1516: return role_skill_1516 // 狂狮怒吼
	case 1517: return role_skill_1517 // 星落
	case 1518: return role_skill_1518 // 无限爆怒
	case 1519: return role_skill_1519 // 乾坤刀气X2
	case 1520: return role_skill_1520 // 召唤堕落嗜血阴毒血士
	case 1521: return role_skill_1521 // 召唤堕落嗜血阴影
	case 1522: return role_skill_1522 // 召唤堕落忆境梦魇
	case 1523: return role_skill_1523 // 铁壁
	case 1524: return role_skill_1524 // 冰咒
	case 1525: return role_skill_1525 // 暴雨梨花针
	case 1526: return role_skill_1526 // 灵旋
	case 1527: return role_skill_1527 // 伏魔伞
	case 1528: return role_skill_1528 // 神兵卷
	case 1529: return role_skill_1529 // 灵旋2级
	case 1530: return role_skill_1530 // 灵旋3级
	case 1531: return role_skill_1531 // 灵旋4级
	case 1532: return role_skill_1532 // 灵旋5级
	case 1533: return role_skill_1533 // 灵旋6级
	case 1534: return role_skill_1534 // 灵旋7级
	case 1535: return role_skill_1535 // 灵旋8级
	case 1536: return role_skill_1536 // 灵旋9级
	case 1537: return role_skill_1537 // 灵旋10级
	case 1538: return role_skill_1538 // 青
	case 1539: return role_skill_1539 // 青
	case 1540: return role_skill_1540 // 青
	case 1541: return role_skill_1541 // 兰
	case 1542: return role_skill_1542 // 兰
	case 1543: return role_skill_1543 // 兰
	case 1544: return role_skill_1544 // 莹
	case 1545: return role_skill_1545 // 莹
	case 1546: return role_skill_1546 // 莹
	case 1547: return role_skill_1547 // 赤
	case 1548: return role_skill_1548 // 赤
	case 1549: return role_skill_1549 // 赤
	case 1550: return role_skill_1550 // 白
	case 1551: return role_skill_1551 // 白
	case 1552: return role_skill_1552 // 白
	case 1556: return role_skill_1556 // 缪
	case 1557: return role_skill_1557 // 缪
	case 1558: return role_skill_1558 // 缪
	case 1559: return role_skill_1559 // 风面
	case 1560: return role_skill_1560 // 风面
	case 1561: return role_skill_1561 // 风面
	case 1562: return role_skill_1562 // 林面
	case 1563: return role_skill_1563 // 林面
	case 1564: return role_skill_1564 // 林面
	case 1565: return role_skill_1565 // 雷面
	case 1566: return role_skill_1566 // 雷面
	case 1567: return role_skill_1567 // 雷面
	case 1568: return role_skill_1568 // 山面
	case 1569: return role_skill_1569 // 山面
	case 1570: return role_skill_1570 // 山面
	case 1571: return role_skill_1571 // 火面
	case 1572: return role_skill_1572 // 火面
	case 1573: return role_skill_1573 // 火面
	case 1574: return role_skill_1574 // 兰龙
	case 1575: return role_skill_1575 // 兰龙
	case 1576: return role_skill_1576 // 兰龙
	case 1577: return role_skill_1577 // 白虎
	case 1578: return role_skill_1578 // 白虎
	case 1579: return role_skill_1579 // 白虎
	case 1580: return role_skill_1580 // 景莲
	case 1581: return role_skill_1581 // 景莲
	case 1582: return role_skill_1582 // 景莲
	case 1583: return role_skill_1583 // 苏摩之怒
	case 1584: return role_skill_1584 // 苏摩之怒附属
	case 1585: return role_skill_1585 // 伏魔伞
	case 1586: return role_skill_1586 // 神兵卷
	case 1587: return role_skill_1587 // 影狼突袭
	case 1588: return role_skill_1588 // 冰魂素魄
	case 1589: return role_skill_1589 // 追云逐电
	case 1590: return role_skill_1590 // 混沌之息
	case 1591: return role_skill_1591 // 影狼突袭·怪
	case 1592: return role_skill_1592 // 冰魂素魄·怪
	case 1593: return role_skill_1593 // 追云逐电·怪
	case 1594: return role_skill_1594 // 混沌之息·怪
	case 1595: return role_skill_1595 // 影狼成群
	case 1596: return role_skill_1596 // 巫雀成群（纵向）
	case 1597: return role_skill_1597 // 巫雀成群（横向）
	case 1598: return role_skill_1598 // 腐蚀剧毒3
	case 1599: return role_skill_1599 // 青竹咒2阶
	case 1600: return role_skill_1600 // 召唤堕落魔竹筒精
	case 1601: return role_skill_1601 // 召唤堕落画妖
	case 1602: return role_skill_1602 // 召唤堕落魔燃魁
	case 1603: return role_skill_1603 // 召唤堕落魔毒蝎
	case 1605: return role_skill_1605 // 撼地突袭
	case 1606: return role_skill_1606 // 枯木逢春
	case 1607: return role_skill_1607 // 撼地突刺
	case 1608: return role_skill_1608 // 枯木逢春
	case 1609: return role_skill_1609 // 天魔拳
	case 1610: return role_skill_1610 // 撼地
	case 1611: return role_skill_1611 // 撼地2级
	case 1612: return role_skill_1612 // 撼地3级
	case 1613: return role_skill_1613 // 撼地4级
	case 1614: return role_skill_1614 // 撼地5级
	case 1615: return role_skill_1615 // 撼地6级
	case 1616: return role_skill_1616 // 撼地7级
	case 1617: return role_skill_1617 // 撼地8级
	case 1618: return role_skill_1618 // 撼地9级
	case 1619: return role_skill_1619 // 撼地10级
	case 1620: return role_skill_1620 // 黑鹰
	case 1621: return role_skill_1621 // 黑鹰
	case 1622: return role_skill_1622 // 黑鹰
	case 1623: return role_skill_1623 // 召唤堕落血髅
	case 1624: return role_skill_1624 // 召唤堕落竹叶青
	case 1625: return role_skill_1625 // 召唤堕落盗墓贼
	case 1626: return role_skill_1626 // 召唤堕落锦衣卫
	case 1627: return role_skill_1627 // 多连斩X2
	case 1628: return role_skill_1628 // 血爆纵向
	case 1629: return role_skill_1629 // 青竹咒纵向
	case 1630: return role_skill_1630 // 青竹咒横向
	case 1631: return role_skill_1631 // 剧毒撕咬
	case 1632: return role_skill_1632 // 破甲撕咬
	case 1633: return role_skill_1633 // 致命撕咬
	case 1634: return role_skill_1634 // 乾坤一掷
	case 1635: return role_skill_1635 // 魔瞳术（单体额外）
	case 1636: return role_skill_1636 // 一闪2阶
	case 1637: return role_skill_1637 // 会心一击2阶
	case 1638: return role_skill_1638 // 暗火之雨2阶
	case 1639: return role_skill_1639 // 凶猛奇袭
	case 1644: return role_skill_1644 // 长溪
	case 1645: return role_skill_1645 // 长溪2级
	case 1646: return role_skill_1646 // 长溪3级
	case 1647: return role_skill_1647 // 长溪4级
	case 1648: return role_skill_1648 // 长溪5级
	case 1649: return role_skill_1649 // 长溪6级
	case 1650: return role_skill_1650 // 长溪7级
	case 1651: return role_skill_1651 // 长溪8级
	case 1652: return role_skill_1652 // 长溪9级
	case 1653: return role_skill_1653 // 长溪10级
	case 1654: return role_skill_1654 // 斩龙剑诀
	case 1655: return role_skill_1655 // 一剑横空
	case 1656: return role_skill_1656 // 断水斩
	case 1657: return role_skill_1657 // 百炼狂刀
	case 1658: return role_skill_1658 // 斩魄刀
	case 1659: return role_skill_1659 // 妙手回春
	case 1660: return role_skill_1660 // 风击暴袭
	case 1661: return role_skill_1661 // 怒风沙爆
	case 1662: return role_skill_1662 // 九转重阳
	case 1663: return role_skill_1663 // 回春续劲
	case 1664: return role_skill_1664 // 青竹神咒
	case 1665: return role_skill_1665 // 影狼突袭
	case 1666: return role_skill_1666 // 巫雀奇袭
	case 1667: return role_skill_1667 // 墨画灵山
	case 1668: return role_skill_1668 // 冷云烟
	case 1669: return role_skill_1669 // 落樱斩
	case 1670: return role_skill_1670 // 影舞斩
	case 1671: return role_skill_1671 // 万剑长空
	case 1672: return role_skill_1672 // 般若行龙
	case 1673: return role_skill_1673 // 收妖诀
	case 1674: return role_skill_1674 // 水行者
	case 1675: return role_skill_1675 // 金钟罩
	case 1676: return role_skill_1676 // 金刚庇护
	case 1677: return role_skill_1677 // 真灵守护
	case 1678: return role_skill_1678 // 狮子吼
	case 1679: return role_skill_1679 // 金刚怒目
	case 1680: return role_skill_1680 // 致命
	case 1681: return role_skill_1681 // 寒冰凛冽
	case 1682: return role_skill_1682 // 寒冰甲
	case 1683: return role_skill_1683 // 落星式
	case 1684: return role_skill_1684 // 冬夜之拥
	case 1685: return role_skill_1685 // 吸星锁
	case 1686: return role_skill_1686 // 龙麟盾
	case 1687: return role_skill_1687 // 千钧破
	case 1688: return role_skill_1688 // 降妖伞
	case 1689: return role_skill_1689 // 撼天狂掌
	case 1690: return role_skill_1690 // 韦驮拳
	case 1691: return role_skill_1691 // 星沉地动
	case 1692: return role_skill_1692 // 寒月剑法
	case 1693: return role_skill_1693 // 万剑藏锋
	case 1694: return role_skill_1694 // 雷动九天
	case 1696: return role_skill_1696 // 星爆光离
	case 1697: return role_skill_1697 // 召唤致命火蝎
	case 1698: return role_skill_1698 // 破甲击
	case 1699: return role_skill_1699 // 咆哮冲撞
	case 1700: return role_skill_1700 // 冲锋
	case 1701: return role_skill_1701 // 针刺
	case 1702: return role_skill_1702 // 魅惑
	case 1703: return role_skill_1703 // 暴怒之击
	case 1704: return role_skill_1704 // 破盾利爪
	case 1705: return role_skill_1705 // 横向挥击
	case 1706: return role_skill_1706 // 纵向挥击
	case 1707: return role_skill_1707 // 狂毒之咬
	case 1708: return role_skill_1708 // 奇袭利爪
	case 1709: return role_skill_1709 // 岩土冲击
	case 1710: return role_skill_1710 // 憾地突袭
	case 1711: return role_skill_1711 // 雷暴
	case 1712: return role_skill_1712 // 雷电轰击
	case 1713: return role_skill_1713 // 咆哮
	case 1714: return role_skill_1714 // 血腥狂暴
	case 1715: return role_skill_1715 // 吸魂
	case 1716: return role_skill_1716 // 毒牙撕咬
	case 1717: return role_skill_1717 // 减攻混乱
	case 1718: return role_skill_1718 // 死亡缠绕
	case 1719: return role_skill_1719 // 诅咒
	case 1720: return role_skill_1720 // 尸毒
	case 1721: return role_skill_1721 // 致命诅咒
	case 1722: return role_skill_1722 // 冰霜袭击
	case 1723: return role_skill_1723 // 冰霜奇袭
	case 1724: return role_skill_1724 // 召唤猛鬼僵尸
	case 1725: return role_skill_1725 // 利爪挥击
	case 1726: return role_skill_1726 // 妖风
	case 1727: return role_skill_1727 // 血魔之力
	case 1728: return role_skill_1728 // 梦蚀
	case 1729: return role_skill_1729 // 噩梦
	case 1730: return role_skill_1730 // 毒素
	case 1731: return role_skill_1731 // 血咒
	case 1732: return role_skill_1732 // 混乱之触
	case 1733: return role_skill_1733 // 冰霜纵向
	case 1734: return role_skill_1734 // 伤害加深（全体）
	case 1735: return role_skill_1735 // 烈焰冲击
	case 1736: return role_skill_1736 // 碾压
	case 1737: return role_skill_1737 // 绝望诅咒
	case 1738: return role_skill_1738 // 封魔剑
	case 1739: return role_skill_1739 // 爆怒之击
	case 1740: return role_skill_1740 // 血蚀纵向
	case 1741: return role_skill_1741 // 影狼纵向
	case 1742: return role_skill_1742 // 召唤魔煞僵尸
	case 1743: return role_skill_1743 // 腐蚀术
	case 1744: return role_skill_1744 // 狂乱之击
	case 1745: return role_skill_1745 // 毒烈横排穿透
	case 1746: return role_skill_1746 // 暗影之毒
	case 1747: return role_skill_1747 // 痛苦诅咒
	case 1748: return role_skill_1748 // 八象皆杀
	case 1749: return role_skill_1749 // 吸血剑
	case 1750: return role_skill_1750 // 吸魂剑
	case 1751: return role_skill_1751 // 致残
	case 1752: return role_skill_1752 // 飓风
	case 1753: return role_skill_1753 // 被动：反伤
	case 1754: return role_skill_1754 // 亡语：炎爆
	case 1755: return role_skill_1755 // 亡语：冲击
	case 1756: return role_skill_1756 // 亡语：免伤护盾
	case 1757: return role_skill_1757 // 亡语：清除增益
	case 1758: return role_skill_1758 // 被动：清除增益
	case 1759: return role_skill_1759 // 被动：化攻为守
	case 1760: return role_skill_1760 // 亡语：魂力
	case 1761: return role_skill_1761 // 亡语：爆裂
	case 1762: return role_skill_1762 // 亡语：强能
	case 1763: return role_skill_1763 // 被动：毒伤
	case 1764: return role_skill_1764 // 亡语：狂乱
	case 1765: return role_skill_1765 // 亡语：爆炸
	case 1766: return role_skill_1766 // 亡语：血爆
	case 1767: return role_skill_1767 // 亡语：毒爆
	case 1768: return role_skill_1768 // 亡语：清除减益
	case 1769: return role_skill_1769 // 亡语：狂暴之血
	case 1770: return role_skill_1770 // 亡语：强攻
	case 1771: return role_skill_1771 // 亡语：治疗敌方
	case 1772: return role_skill_1772 // 亡语：爆冲
	case 1773: return role_skill_1773 // 亡语：爆乱
	case 1774: return role_skill_1774 // 亡语：风阵
	case 1775: return role_skill_1775 // 尸王横扫
	case 1776: return role_skill_1776 // 魁拔突刺
	case 1777: return role_skill_1777 // 魔血爆纵向
	case 1778: return role_skill_1778 // 死亡挥击
	case 1779: return role_skill_1779 // 末日之拳
	case 1780: return role_skill_1780 // 尸王落刃
	case 1781: return role_skill_1781 // 亡语：蚀甲
	case 1782: return role_skill_1782 // 爆炸蓝
	case 1783: return role_skill_1783 // 爆炸红
	case 1784: return role_skill_1784 // 爆炸绿
	case 1785: return role_skill_1785 // 爆炸黄
	case 1786: return role_skill_1786 // 召唤爆裂虫
	case 1787: return role_skill_1787 // 召唤魔刃使徒
	case 1788: return role_skill_1788 // 召唤马刀使徒
	case 1789: return role_skill_1789 // 闪击
	case 1790: return role_skill_1790 // 飓风冲击
	case 1791: return role_skill_1791 // 野蛮冲击
	}
	return nil
}

func createTotemSkill(id int16, level int16) (*Skill, *SkillInfo) {
	switch id {
		case 1538: return role_skill_1538, &SkillInfo{SkillTrnLv: level-1}  // 青
		case 1539: return role_skill_1539, &SkillInfo{SkillTrnLv: level-1}  // 青
		case 1540: return role_skill_1540, &SkillInfo{SkillTrnLv: level-1}  // 青
		case 1541: return role_skill_1541, &SkillInfo{SkillTrnLv: level-1}  // 兰
		case 1542: return role_skill_1542, &SkillInfo{SkillTrnLv: level-1}  // 兰
		case 1543: return role_skill_1543, &SkillInfo{SkillTrnLv: level-1}  // 兰
		case 1544: return role_skill_1544, &SkillInfo{SkillTrnLv: level-1}  // 莹
		case 1545: return role_skill_1545, &SkillInfo{SkillTrnLv: level-1}  // 莹
		case 1546: return role_skill_1546, &SkillInfo{SkillTrnLv: level-1}  // 莹
		case 1547: return role_skill_1547, &SkillInfo{SkillTrnLv: level-1}  // 赤
		case 1548: return role_skill_1548, &SkillInfo{SkillTrnLv: level-1}  // 赤
		case 1549: return role_skill_1549, &SkillInfo{SkillTrnLv: level-1}  // 赤
		case 1556: return role_skill_1556, &SkillInfo{SkillTrnLv: level-1}  // 缪
		case 1557: return role_skill_1557, &SkillInfo{SkillTrnLv: level-1}  // 缪
		case 1558: return role_skill_1558, &SkillInfo{SkillTrnLv: level-1}  // 缪
		case 1550: return role_skill_1550, &SkillInfo{SkillTrnLv: level-1}  // 白
		case 1551: return role_skill_1551, &SkillInfo{SkillTrnLv: level-1}  // 白
		case 1552: return role_skill_1552, &SkillInfo{SkillTrnLv: level-1}  // 白
		case 1559: return role_skill_1559, &SkillInfo{SkillTrnLv: level-1}  // 风面
		case 1560: return role_skill_1560, &SkillInfo{SkillTrnLv: level-1}  // 风面
		case 1561: return role_skill_1561, &SkillInfo{SkillTrnLv: level-1}  // 风面
		case 1562: return role_skill_1562, &SkillInfo{SkillTrnLv: level-1}  // 林面
		case 1563: return role_skill_1563, &SkillInfo{SkillTrnLv: level-1}  // 林面
		case 1564: return role_skill_1564, &SkillInfo{SkillTrnLv: level-1}  // 林面
		case 1565: return role_skill_1565, &SkillInfo{SkillTrnLv: level-1}  // 雷面
		case 1566: return role_skill_1566, &SkillInfo{SkillTrnLv: level-1}  // 雷面
		case 1567: return role_skill_1567, &SkillInfo{SkillTrnLv: level-1}  // 雷面
		case 1568: return role_skill_1568, &SkillInfo{SkillTrnLv: level-1}  // 山面
		case 1569: return role_skill_1569, &SkillInfo{SkillTrnLv: level-1}  // 山面
		case 1570: return role_skill_1570, &SkillInfo{SkillTrnLv: level-1}  // 山面
		case 1571: return role_skill_1571, &SkillInfo{SkillTrnLv: level-1}  // 火面
		case 1572: return role_skill_1572, &SkillInfo{SkillTrnLv: level-1}  // 火面
		case 1573: return role_skill_1573, &SkillInfo{SkillTrnLv: level-1}  // 火面
		case 1574: return role_skill_1574, &SkillInfo{SkillTrnLv: level-1}  // 兰龙
		case 1575: return role_skill_1575, &SkillInfo{SkillTrnLv: level-1}  // 兰龙
		case 1576: return role_skill_1576, &SkillInfo{SkillTrnLv: level-1}  // 兰龙
		case 1577: return role_skill_1577, &SkillInfo{SkillTrnLv: level-1}  // 白虎
		case 1578: return role_skill_1578, &SkillInfo{SkillTrnLv: level-1}  // 白虎
		case 1579: return role_skill_1579, &SkillInfo{SkillTrnLv: level-1}  // 白虎
		case 1580: return role_skill_1580, &SkillInfo{SkillTrnLv: level-1}  // 景莲
		case 1581: return role_skill_1581, &SkillInfo{SkillTrnLv: level-1}  // 景莲
		case 1582: return role_skill_1582, &SkillInfo{SkillTrnLv: level-1}  // 景莲
		case 1620: return role_skill_1620, &SkillInfo{SkillTrnLv: level-1}  // 黑鹰
		case 1621: return role_skill_1621, &SkillInfo{SkillTrnLv: level-1}  // 黑鹰
		case 1622: return role_skill_1622, &SkillInfo{SkillTrnLv: level-1}  // 黑鹰
	}
	return nil, nil
}

func createGhostSkill(fighter *Fighter) (skill *Skill, skillInfo *SkillInfo) {
	id := int(fighter.useGhostSkillId)
	switch id {
	case 999:	// 审判惊雷
		skill = role_skill_999
		skillInfo = &SkillInfo{SkillId:999, SkillId2:1076}
	case 1000:	// 铜墙铁壁
		skill = role_skill_1000
		skillInfo = &SkillInfo{SkillId:1000, SkillId2:1079}
	case 1001:	// 风之领域
		skill = role_skill_1001
		skillInfo = &SkillInfo{SkillId:1001, SkillId2:1077}
	case 1002:	// 无限霸刀
		skill = role_skill_1002
		skillInfo = &SkillInfo{SkillId:1002, SkillId2:1420}
	case 1003:	// 熔岩弹雨
		skill = role_skill_1003
		skillInfo = &SkillInfo{SkillId:1003, SkillId2:1086}
	case 1004:	// 人鱼之歌
		skill = role_skill_1004
		skillInfo = &SkillInfo{SkillId:1004, SkillId2:1421}
	case 1017:	// 振奋之击
		skill = role_skill_1017
		skillInfo = &SkillInfo{SkillId:1017, SkillId2:1098}
	case 1020:	// 致胜千里
		skill = role_skill_1020
		skillInfo = &SkillInfo{SkillId:1020, SkillId2:1101}
	case 1023:	// 万剑穿心
		skill = role_skill_1023
		skillInfo = &SkillInfo{SkillId:1023, SkillId2:1104}
	case 1026:	// 阿修罗之怒
		skill = role_skill_1026
		skillInfo = &SkillInfo{SkillId:1026, SkillId2:1278}
	case 1032:	// 圣灵洛水
		skill = role_skill_1032
		skillInfo = &SkillInfo{SkillId:1032, SkillId2:1107}
	case 1029:	// 命运连锁
		skill = role_skill_1029
		skillInfo = &SkillInfo{SkillId:1029, SkillId2:1110}
	case 1042:	// 灵蛇之枪
		skill = role_skill_1042
		skillInfo = &SkillInfo{SkillId:1042, SkillId2:1113}
	case 1058:	// 暗影食梦
		skill = role_skill_1058
		skillInfo = &SkillInfo{SkillId:1058, SkillId2:1116}
	case 1061:	// 寒霜灵泉
		skill = role_skill_1061
		skillInfo = &SkillInfo{SkillId:1061, SkillId2:1119}
	case 1043:	// 火凤燎原
		skill = role_skill_1043
		skillInfo = &SkillInfo{SkillId:1043, SkillId2:1122}
	case 1064:	// 阎罗试炼
		skill = role_skill_1064
		skillInfo = &SkillInfo{SkillId:1064, SkillId2:1125}
	case 1173:	// 魅影之爪
		skill = role_skill_1173
		skillInfo = &SkillInfo{SkillId:1173, SkillId2:1174}
	case 1179:	// 横扫千军
		skill = role_skill_1179
		skillInfo = &SkillInfo{SkillId:1179, SkillId2:1422}
	case 1185:	// 碎铁之刃
		skill = role_skill_1185
		skillInfo = &SkillInfo{SkillId:1185, SkillId2:1423}
	case 1193:	// 寒冰护体
		skill = role_skill_1193
		skillInfo = &SkillInfo{SkillId:1193, SkillId2:1194}
	case 1199:	// 瞒天过海
		skill = role_skill_1199
		skillInfo = &SkillInfo{SkillId:1199, SkillId2:1424}
	case 1284:	// 银光落刃
		skill = role_skill_1284
		skillInfo = &SkillInfo{SkillId:1284, SkillId2:1425}
	case 1287:	// 式神炎舞
		skill = role_skill_1287
		skillInfo = &SkillInfo{SkillId:1287, SkillId2:1288}
	case 1293:	// 猛龙断空
		skill = role_skill_1293
		skillInfo = &SkillInfo{SkillId:1293, SkillId2:1294}
	case 1583:	// 苏摩之怒
		skill = role_skill_1583
		skillInfo = &SkillInfo{SkillId:1583, SkillId2:1584}
	}
	ghost := fighter.GetMainGhost()
	if ghost != nil {
		skillInfo.SkillTrnLv = ghost.GhostSkillLv
		skillInfo.SkillTrnLv2 = ghost.GhostSkillLv
	}
	return 
}

// 落云
var role_skill_1 = &Skill{
	SkillId: 1,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 10,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 10 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 裂地斩
var role_skill_2 = &Skill{
	SkillId: 2,
	ChildType: 1,
	DecPower: 2,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 0 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 飞焰穿云
var role_skill_4 = &Skill{
	SkillId: 4,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 1000,
	SunderAttack: 10,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 甲溃
var role_skill_6 = &Skill{
	SkillId: 6,
	ChildType: 6,
	FixedValue: 500,
	SunderAttack: 300,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.9 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 大风咒
var role_skill_8 = &Skill{
	SkillId: 8,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 咆哮利爪
var role_skill_9 = &Skill{
	SkillId: 9,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 凶猛撕咬
var role_skill_10 = &Skill{
	SkillId: 10,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 冰烈
var role_skill_11 = &Skill{
	SkillId: 11,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 冰烈横向
var role_skill_12 = &Skill{
	SkillId: 12,
	ChildType: 1,
	FixedValue: 0,
	ReduceDefend: 0.05,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 冰烈纵向
var role_skill_13 = &Skill{
	SkillId: 13,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 冰烈全体
var role_skill_14 = &Skill{
	SkillId: 14,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 火烈
var role_skill_15 = &Skill{
	SkillId: 15,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 火烈横向
var role_skill_16 = &Skill{
	SkillId: 16,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 火烈纵向
var role_skill_17 = &Skill{
	SkillId: 17,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 火烈全体
var role_skill_18 = &Skill{
	SkillId: 18,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风烈
var role_skill_19 = &Skill{
	SkillId: 19,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风烈横向
var role_skill_20 = &Skill{
	SkillId: 20,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风烈纵向
var role_skill_21 = &Skill{
	SkillId: 21,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风烈全体
var role_skill_22 = &Skill{
	SkillId: 22,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷烈
var role_skill_23 = &Skill{
	SkillId: 23,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷烈横向
var role_skill_24 = &Skill{
	SkillId: 24,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷烈纵向
var role_skill_25 = &Skill{
	SkillId: 25,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷烈全体
var role_skill_26 = &Skill{
	SkillId: 26,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 土烈
var role_skill_27 = &Skill{
	SkillId: 27,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 土烈横向
var role_skill_28 = &Skill{
	SkillId: 28,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 土烈纵向
var role_skill_29 = &Skill{
	SkillId: 29,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 土烈全体
var role_skill_30 = &Skill{
	SkillId: 30,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 毒烈
var role_skill_31 = &Skill{
	SkillId: 31,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 2, 31, 1, false))
		return buffs
	},
}

// 毒烈横向
var role_skill_32 = &Skill{
	SkillId: 32,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 2, 32, 1, false))
		return buffs
	},
}

// 毒烈纵向
var role_skill_33 = &Skill{
	SkillId: 33,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 2, 33, 1, false))
		return buffs
	},
}

// 毒烈全体
var role_skill_34 = &Skill{
	SkillId: 34,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 3, 34, 1, false))
		return buffs
	},
}

// 多连斩
var role_skill_35 = &Skill{
	SkillId: 35,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 力劈华山
var role_skill_36 = &Skill{
	SkillId: 36,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 白莲横江
var role_skill_37 = &Skill{
	SkillId: 37,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 横扫千军
var role_skill_38 = &Skill{
	SkillId: 38,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 乾坤刀气
var role_skill_39 = &Skill{
	SkillId: 39,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 三千洛水剑
var role_skill_40 = &Skill{
	SkillId: 40,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 死亡标记
var role_skill_41 = &Skill{
	SkillId: 41,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 万箭穿心
var role_skill_42 = &Skill{
	SkillId: 42,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 狮吼功
var role_skill_43 = &Skill{
	SkillId: 43,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 野蛮冲撞(怪物用)
var role_skill_44 = &Skill{
	SkillId: 44,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 如殂随行
var role_skill_45 = &Skill{
	SkillId: 45,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 驱散
var role_skill_46 = &Skill{
	SkillId: 46,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 46, 1, false))
		return buffs
	},
}

// 青竹咒
var role_skill_49 = &Skill{
	SkillId: 49,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(500 + float64(skillTrnlv) * 10)), 2, 49, 1, false))
		return buffs
	},
}

// 墨画影狼
var role_skill_51 = &Skill{
	SkillId: 51,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(200 + float64(skillTrnlv) * 10)), 3, 51, 3, false))
		return buffs
	},
}

// 墨画巫雀
var role_skill_52 = &Skill{
	SkillId: 52,
	ChildType: 5,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(500 + float64(skillTrnlv) * 20)), 2, 52, 3, false))
		return buffs
	},
}

// 坠星击
var role_skill_89 = &Skill{
	SkillId: 89,
	ChildType: 1,
	DecPower: 2,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 天剑
var role_skill_98 = &Skill{
	SkillId: 98,
	ChildType: 1,
	DecPower: 8,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 治疗
var role_skill_108 = &Skill{
	SkillId: 108,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 108, 1, false))
			}
		}
		return buffs
	},
}

// 增益
var role_skill_109 = &Skill{
	SkillId: 109,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 2, 109, 1, false))
			}
		}
		return buffs
	},
}

// 破甲眩晕
var role_skill_998 = &Skill{
	SkillId: 998,
	ChildType: 0,
	FixedValue: 0,
}

// 审判惊雷
var role_skill_999 = &Skill{
	SkillId: 999,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 铜墙铁壁
var role_skill_1000 = &Skill{
	SkillId: 1000,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 风之领域
var role_skill_1001 = &Skill{
	SkillId: 1001,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.9 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1001, 1, false))
		return buffs
	},
}

// 无限霸刀
var role_skill_1002 = &Skill{
	SkillId: 1002,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 熔岩弹雨
var role_skill_1003 = &Skill{
	SkillId: 1003,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 人鱼之歌
var role_skill_1004 = &Skill{
	SkillId: 1004,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 振奋之击
var role_skill_1017 = &Skill{
	SkillId: 1017,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 致胜千里
var role_skill_1020 = &Skill{
	SkillId: 1020,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 万剑穿心
var role_skill_1023 = &Skill{
	SkillId: 1023,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 阿修罗之怒
var role_skill_1026 = &Skill{
	SkillId: 1026,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 3000,
	SunderAttack: 500,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 命运连锁
var role_skill_1029 = &Skill{
	SkillId: 1029,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 圣灵洛水
var role_skill_1032 = &Skill{
	SkillId: 1032,
	ChildType: 4,
	IsGhostSkill: true,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(1000 + float64(skillTrnlv) * 300 + f.Cultivation * 1.5 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1032, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1032, 1, false))
			}
		}
		return buffs
	},
}

// 死亡阻击
var role_skill_1035 = &Skill{
	SkillId: 1035,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 紫电刀芒
var role_skill_1038 = &Skill{
	SkillId: 1038,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 25,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 斩杀
var role_skill_1039 = &Skill{
	SkillId: 1039,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 聚气
var role_skill_1041 = &Skill{
	SkillId: 1041,
	ChildType: 5,
	IncPower: 2,
	FixedValue: 0,
}

// 灵蛇之枪
var role_skill_1042 = &Skill{
	SkillId: 1042,
	ChildType: 1,
	IsGhostSkill: true,
	IncPower: 2,
	FixedValue: 1000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 火凤燎原
var role_skill_1043 = &Skill{
	SkillId: 1043,
	ChildType: 1,
	IsGhostSkill: true,
	IncPower: 2,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 全体治疗
var role_skill_1047 = &Skill{
	SkillId: 1047,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1047, 1, false))
			}
		}
		return buffs
	},
}

// 青竹咒
var role_skill_1048 = &Skill{
	SkillId: 1048,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雨润
var role_skill_1049 = &Skill{
	SkillId: 1049,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 2, 1049, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1049, 1, false))
			}
		}
		return buffs
	},
}

// 圣白莲
var role_skill_1050 = &Skill{
	SkillId: 1050,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1050, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1050, 1, false))
			}
		}
		return buffs
	},
}

// 神魔封禁
var role_skill_1051 = &Skill{
	SkillId: 1051,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 1, 1051, 1, false))
		}
		return buffs
	},
}

// 墨画巫雀
var role_skill_1052 = &Skill{
	SkillId: 1052,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1052, 1, false))
		return buffs
	},
}

// 墨画影狼
var role_skill_1053 = &Skill{
	SkillId: 1053,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1 + force)), 2, 1053, 1, false))
		return buffs
	},
}

// 祝福
var role_skill_1054 = &Skill{
	SkillId: 1054,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(force)), 2, 1054, 1, false))
			}
		}
		return buffs
	},
}

// 地狱烈焰
var role_skill_1055 = &Skill{
	SkillId: 1055,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 10,
	Critial: 100,
	NotMiss: true,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 万剑归一
var role_skill_1056 = &Skill{
	SkillId: 1056,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 暗影食梦
var role_skill_1058 = &Skill{
	SkillId: 1058,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.9 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 寒霜灵泉
var role_skill_1061 = &Skill{
	SkillId: 1061,
	ChildType: 4,
	IsGhostSkill: true,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(500 + float64(skillTrnlv) * 250 + f.Cultivation * 1.2 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1061, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1061, 1, false))
			}
		}
		return buffs
	},
}

// 阎罗试炼
var role_skill_1064 = &Skill{
	SkillId: 1064,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 雨润
var role_skill_1067 = &Skill{
	SkillId: 1067,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1067, 1, false))
			}
		}
		buddies2 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HEALTH, (buff_value(3000 + float64(skillTrnlv) * 50 + f.Cultivation + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1067, 1, false))
			}
		}
		return buffs
	},
}

// 坠星击
var role_skill_1070 = &Skill{
	SkillId: 1070,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 治愈之风
var role_skill_1074 = &Skill{
	SkillId: 1074,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(2000 + float64(skillTrnlv) * 50 + f.Cultivation + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1074, 1, false))
			}
		}
		return buffs
	},
}

// 刀盾
var role_skill_1075 = &Skill{
	SkillId: 1075,
	ChildType: 3,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 3)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ABSORB_HURT, (buff_value(2000 + float64(skillTrnlv) * 50)), 2, 1075, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 2, 1075, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_REDUCE_HURT, (buff_value(50)), 2, 1075, 1, false))
		return buffs
	},
}

// 审判惊雷附属
var role_skill_1076 = &Skill{
	SkillId: 1076,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(500 + float64(skillTrnlv) * 25)), 1, 1076, 1, false))
			}
		}
		return buffs
	},
}

// 风之领域附属
var role_skill_1077 = &Skill{
	SkillId: 1077,
	ChildType: 4,
	IsGhostSkill: true,
	FixedValue: 500,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(500 + float64(skillTrnlv) * 150 + f.Cultivation + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1077, 1, false))
			}
		}
		return buffs
	},
}

// 铜墙铁壁附属
var role_skill_1079 = &Skill{
	SkillId: 1079,
	ChildType: 3,
	IsGhostSkill: true,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(50)), 1, 1079, 1, false))
			}
		}
		return buffs
	},
}

// 熔岩弹雨附属
var role_skill_1086 = &Skill{
	SkillId: 1086,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000 + float64(skillTrnlv) * 30)), 2, 1086, 1, false))
		return buffs
	},
}

// 风镰
var role_skill_1092 = &Skill{
	SkillId: 1092,
	ChildType: 1,
	FixedValue: 10,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 10 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击
var role_skill_1094 = &Skill{
	SkillId: 1094,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔
var role_skill_1095 = &Skill{
	SkillId: 1095,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 刀芒
var role_skill_1096 = &Skill{
	SkillId: 1096,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 大风咒
var role_skill_1097 = &Skill{
	SkillId: 1097,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 0 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 振奋之击附属
var role_skill_1098 = &Skill{
	SkillId: 1098,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(2000 + float64(skillTrnlv) * 50)), 3, 1098, 1, false))
			}
		}
		return buffs
	},
}

// 致胜千里附属
var role_skill_1101 = &Skill{
	SkillId: 1101,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(2000 + float64(skillTrnlv) * 40)), 2, 1101, 1, false))
			}
		}
		return buffs
	},
}

// 万剑穿心附属
var role_skill_1104 = &Skill{
	SkillId: 1104,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 15)), 2, 1104, 1, false))
			}
		}
		return buffs
	},
}

// 圣灵洛水附属
var role_skill_1107 = &Skill{
	SkillId: 1107,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 命运连锁附属
var role_skill_1110 = &Skill{
	SkillId: 1110,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.9 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1110, 1, false))
		}
		return buffs
	},
}

// 灵蛇之枪附属
var role_skill_1113 = &Skill{
	SkillId: 1113,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 暗影食梦附属
var role_skill_1116 = &Skill{
	SkillId: 1116,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.9 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1116, 1, false))
		}
		return buffs
	},
}

// 寒霜灵泉附属
var role_skill_1119 = &Skill{
	SkillId: 1119,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 火凤燎原附属
var role_skill_1122 = &Skill{
	SkillId: 1122,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1000 + float64(skillTrnlv) * 30)), 2, 1122, 1, false))
			}
		}
		return buffs
	},
}

// 阎罗试炼附属
var role_skill_1125 = &Skill{
	SkillId: 1125,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(1000 + float64(skillTrnlv) * 50)), 2, 1125, 1, false))
			}
		}
		return buffs
	},
}

// 怒火冲天
var role_skill_1134 = &Skill{
	SkillId: 1134,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(100 + float64(skillTrnlv) * 25)), 2, 1134, 2, false))
			}
		}
		return buffs
	},
}

// 剑冲阴阳
var role_skill_1135 = &Skill{
	SkillId: 1135,
	ChildType: 6,
	FixedValue: 3000,
	SunderAttack: 40,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
}

// 黄粱一梦
var role_skill_1136 = &Skill{
	SkillId: 1136,
	ChildType: 5,
	FixedValue: 3000,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.6 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 3, 1136, 1, false))
		}
		return buffs
	},
}

// 蹑影追风
var role_skill_1137 = &Skill{
	SkillId: 1137,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 1, 1137, 1, false))
			}
		}
		return buffs
	},
}

// 野蛮冲撞
var role_skill_1138 = &Skill{
	SkillId: 1138,
	ChildType: 6,
	FixedValue: 2500,
	SunderAttack: 40,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2500 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
}

// 铁布衫
var role_skill_1139 = &Skill{
	SkillId: 1139,
	ChildType: 3,
	DecPower: 2,
	FixedValue: 1000,
	SunderAttack: 5,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 1, 1139, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_DEFEND, (buff_value(1000 + float64(skillTrnlv) * 30)), 1, 1139, 1, false))
		return buffs
	},
}

// 裂地斩2级
var role_skill_1140 = &Skill{
	SkillId: 1140,
	ChildType: 1,
	DecPower: 2,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 偃月
var role_skill_1141 = &Skill{
	SkillId: 1141,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 10,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 治愈之风2级
var role_skill_1142 = &Skill{
	SkillId: 1142,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(6000 + float64(skillTrnlv) * 50 + f.Cultivation + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1142, 1, false))
			}
		}
		return buffs
	},
}

// 藤甲术
var role_skill_1143 = &Skill{
	SkillId: 1143,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(500 + float64(skillTrnlv) * 20)), 1, 1143, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH && buddy2.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_SUNDER, (buff_value(100)), 0, 1143, 1, false))
			}
		}
		return buffs
	},
}

// 偃月2级
var role_skill_1144 = &Skill{
	SkillId: 1144,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 10,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 墨画魂刃
var role_skill_1145 = &Skill{
	SkillId: 1145,
	ChildType: 5,
	FixedValue: 500,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 四象皆杀
var role_skill_1146 = &Skill{
	SkillId: 1146,
	ChildType: 5,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_BLOCK_LEVEL, -(buff_value(500 + float64(skillTrnlv) * 10)), 3, 1146, 2, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DODGE_LEVEL, -(buff_value(500 + float64(skillTrnlv) * 10)), 3, 1146, 2, false))
		return buffs
	},
}

// 护盾
var role_skill_1147 = &Skill{
	SkillId: 1147,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ABSORB_HURT, (buff_value(1 + force)), 3, 1147, 1, false))
			}
		}
		return buffs
	},
}

// 雨润2级
var role_skill_1148 = &Skill{
	SkillId: 1148,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1148, 1, false))
			}
		}
		buddies2 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HEALTH, (buff_value(6000 + float64(skillTrnlv) * 50 + f.Cultivation + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1148, 1, false))
			}
		}
		return buffs
	},
}

// 落云2级
var role_skill_1149 = &Skill{
	SkillId: 1149,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 纯阳罡气
var role_skill_1150 = &Skill{
	SkillId: 1150,
	ChildType: 3,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ABSORB_HURT, (buff_value(5000 + float64(skillTrnlv) * 50)), 2, 1150, 1, false))
			}
		}
		return buffs
	},
}

// 斩击2级
var role_skill_1151 = &Skill{
	SkillId: 1151,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 顺势斩
var role_skill_1153 = &Skill{
	SkillId: 1153,
	ChildType: 1,
	DecPower: 3,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 破釜沉舟
var role_skill_1155 = &Skill{
	SkillId: 1155,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 1500,
	SunderAttack: 150,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1500 + f.Cultivation * 0.1 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 纳气归元
var role_skill_1156 = &Skill{
	SkillId: 1156,
	ChildType: 5,
	DecPower: 4,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HEALTH, (buff_value(8000 + float64(skillTrnlv) * 30 + f.Cultivation + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1156, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HURT_ADD, (buff_value(50)), 2, 1156, 1, false))
		return buffs
	},
}

// 狂乱之刃
var role_skill_1157 = &Skill{
	SkillId: 1157,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findMostHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 1.2 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffCount = t.GetBuffCount(1)
		if rand.Float64() < 0.8 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01 - (float64(buffCount) * 20 / 100.0)  {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1157, 1, false))
			t.AddBuffCount(1)
		}
		return buffs
	},
}

// 护盾破坏
var role_skill_1158 = &Skill{
	SkillId: 1158,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 10,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1158, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1158, 1, false))
		return buffs
	},
}

// 刀盾2级
var role_skill_1159 = &Skill{
	SkillId: 1159,
	ChildType: 3,
	FixedValue: 3000,
	SunderAttack: 10,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 3)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ABSORB_HURT, (buff_value(10000 + float64(skillTrnlv) * 50)), 2, 1159, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 2, 1159, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_REDUCE_HURT, (buff_value(50)), 2, 1159, 1, false))
		return buffs
	},
}

// 魂力果实
var role_skill_1164 = &Skill{
	SkillId: 1164,
	ChildType: 5,
	FixedValue: 2000,
	SunderAttack: 5,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_GHOST_POWER, (buff_value(100)), 0, 1164, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HIT_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 1, 1164, 1, false))
			}
		}
		return buffs
	},
}

// 青竹咒2级
var role_skill_1165 = &Skill{
	SkillId: 1165,
	ChildType: 1,
	FixedValue: 3000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1165, 1, false))
		return buffs
	},
}

// 墨画影狼2级
var role_skill_1167 = &Skill{
	SkillId: 1167,
	ChildType: 5,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(500 + float64(skillTrnlv) * 10)), 3, 1167, 3, false))
		return buffs
	},
}

// 墨画巫雀2级
var role_skill_1168 = &Skill{
	SkillId: 1168,
	ChildType: 5,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1000 + float64(skillTrnlv) * 20)), 2, 1168, 3, false))
		return buffs
	},
}

// 风卷尘生
var role_skill_1169 = &Skill{
	SkillId: 1169,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰2级
var role_skill_1170 = &Skill{
	SkillId: 1170,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 妙法莲华
var role_skill_1171 = &Skill{
	SkillId: 1171,
	ChildType: 4,
	FixedValue: 2000,
	SunderAttack: 5,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffCount = buddy1.GetBuffCount(2)
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(6000 + float64(skillTrnlv) * 50 + f.Cultivation - buffCount * 4000 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1171, 1, false))
					buddy1.AddBuffCount(2)
			}
		}
		return buffs
	},
}

// 魅影之爪
var role_skill_1173 = &Skill{
	SkillId: 1173,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 魅影之爪附属
var role_skill_1174 = &Skill{
	SkillId: 1174,
	ChildType: 5,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 4, 1174, 1, false))
		}
		return buffs
	},
}

// 横扫千军
var role_skill_1179 = &Skill{
	SkillId: 1179,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 碎铁之刃
var role_skill_1185 = &Skill{
	SkillId: 1185,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 召唤阴影
var role_skill_1191 = &Skill{
	SkillId: 1191,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 24,3},CallInfo{ 24,1}},
}

// 寒冰护体
var role_skill_1193 = &Skill{
	SkillId: 1193,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 寒冰护体附属
var role_skill_1194 = &Skill{
	SkillId: 1194,
	ChildType: 5,
	IsGhostSkill: true,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ABSORB_HURT, (buff_value(1000 + float64(skillTrnlv) * 200)), 3, 1194, 1, false))
			}
		}
		return buffs
	},
}

// 瞒天过海
var role_skill_1199 = &Skill{
	SkillId: 1199,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 召唤堕落竹筒精
var role_skill_1205 = &Skill{
	SkillId: 1205,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 436,13},CallInfo{ 436,11},CallInfo{ 436,3},CallInfo{ 436,1}},
}

// 召唤治愈莲藕精
var role_skill_1206 = &Skill{
	SkillId: 1206,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 152,3},CallInfo{ 152,1}},
}

// 召唤堕落燃魁
var role_skill_1207 = &Skill{
	SkillId: 1207,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 16,8},CallInfo{ 16,6},CallInfo{ 16,2}},
}

// 召唤暴怒剑之守卫
var role_skill_1208 = &Skill{
	SkillId: 1208,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 194,3},CallInfo{ 194,1}},
}

// 召唤暴怒败亡之剑
var role_skill_1209 = &Skill{
	SkillId: 1209,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 192,8},CallInfo{ 192,6},CallInfo{ 192,2}},
}

// 召唤调皮火灵
var role_skill_1210 = &Skill{
	SkillId: 1210,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 379,3},CallInfo{ 379,1}},
}

// 召唤调皮豆灵
var role_skill_1211 = &Skill{
	SkillId: 1211,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 386,8},CallInfo{ 386,6},CallInfo{ 385,3},CallInfo{ 385,2},CallInfo{ 385,1}},
}

// 召唤调皮金翅飞鸾
var role_skill_1212 = &Skill{
	SkillId: 1212,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 393,3},CallInfo{ 393,1}},
}

// 龙息术
var role_skill_1213 = &Skill{
	SkillId: 1213,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 20,
	Critial: 80,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 泰山压顶
var role_skill_1214 = &Skill{
	SkillId: 1214,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 10,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 浮生入梦
var role_skill_1215 = &Skill{
	SkillId: 1215,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 10,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(50 + float64(hurt) * 0.3)), 3, 1215, 1, false))
		return buffs
	},
}

// 斩击
var role_skill_1216 = &Skill{
	SkillId: 1216,
	ChildType: 1,
	FixedValue: 10,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 10 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 霸拳盖天
var role_skill_1218 = &Skill{
	SkillId: 1218,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 10,
	Critial: 50,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 召唤调皮雷兽
var role_skill_1219 = &Skill{
	SkillId: 1219,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 411,3},CallInfo{ 411,1}},
}

// 召唤顽皮灵剑
var role_skill_1220 = &Skill{
	SkillId: 1220,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 418,12},CallInfo{ 418,8},CallInfo{ 418,6},CallInfo{ 418,2}},
}

// 召唤淘气火灵
var role_skill_1221 = &Skill{
	SkillId: 1221,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 425,8},CallInfo{ 425,6}},
}

// 召唤顽皮飞鹏
var role_skill_1222 = &Skill{
	SkillId: 1222,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 432,8},CallInfo{ 432,6},CallInfo{ 431,2}},
}

// 十字斩
var role_skill_1223 = &Skill{
	SkillId: 1223,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风咒
var role_skill_1224 = &Skill{
	SkillId: 1224,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 聚气蓄力
var role_skill_1225 = &Skill{
	SkillId: 1225,
	ChildType: 5,
	FixedValue: 500,
	SunderAttack: 5,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_POWER, (buff_value(2)), 0, 1225, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_TENACITY_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 1, 1225, 1, false))
		return buffs
	},
}

// 断岳
var role_skill_1226 = &Skill{
	SkillId: 1226,
	ChildType: 1,
	DecPower: 2,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 激励
var role_skill_1227 = &Skill{
	SkillId: 1227,
	ChildType: 5,
	DecPower: 8,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 3)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurBuddies(f)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BUDDY_SKILL, (buff_value(1)), 0, 1227, 1, false))
			}
		}
		buddies2 := findBattlePetBuddy(f)
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_PET_LIVE_ROUND, (buff_value(1)), 0, 1227, 1, false))
			}
		}
		buddies3 := f.getBuddies()
		for _, buddy3 := range buddies3 {
			if buddy3 != nil && ( false /*dead target*/ || (buddy3.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy3.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 1, 1227, 1, false))
			}
		}
		return buffs
	},
}

// 生生不息·怪
var role_skill_1228 = &Skill{
	SkillId: 1228,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(1 + force + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1228, 1, false))
			}
		}
		return buffs
	},
}

// 生生不息3级
var role_skill_1229 = &Skill{
	SkillId: 1229,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1229, 1, false))
			}
		}
		return buffs
	},
}

// 生生不息4级
var role_skill_1230 = &Skill{
	SkillId: 1230,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1230, 1, false))
			}
		}
		return buffs
	},
}

// 生生不息5级
var role_skill_1231 = &Skill{
	SkillId: 1231,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1231, 1, false))
			}
		}
		return buffs
	},
}

// 怒火冲天·怪
var role_skill_1232 = &Skill{
	SkillId: 1232,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(1 + force)), 2, 1232, 2, false))
			}
		}
		return buffs
	},
}

// 怒火冲天3级
var role_skill_1233 = &Skill{
	SkillId: 1233,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 2, 1233, 2, false))
			}
		}
		return buffs
	},
}

// 怒火冲天4级
var role_skill_1234 = &Skill{
	SkillId: 1234,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 2, 1234, 2, false))
			}
		}
		return buffs
	},
}

// 怒火冲天5级
var role_skill_1235 = &Skill{
	SkillId: 1235,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(force)), 2, 1235, 2, false))
			}
		}
		return buffs
	},
}

// 野蛮冲撞·怪
var role_skill_1236 = &Skill{
	SkillId: 1236,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 野蛮冲撞3级
var role_skill_1237 = &Skill{
	SkillId: 1237,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 60,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 野蛮冲撞4级
var role_skill_1238 = &Skill{
	SkillId: 1238,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 80,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 野蛮冲撞5级
var role_skill_1239 = &Skill{
	SkillId: 1239,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 100,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 剑冲阴阳·怪
var role_skill_1240 = &Skill{
	SkillId: 1240,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 40,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 剑冲阴阳3级
var role_skill_1241 = &Skill{
	SkillId: 1241,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 剑冲阴阳4级
var role_skill_1242 = &Skill{
	SkillId: 1242,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 70,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 剑冲阴阳5级
var role_skill_1243 = &Skill{
	SkillId: 1243,
	ChildType: 6,
	FixedValue: 0,
	SunderAttack: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 黄粱一梦·怪
var role_skill_1244 = &Skill{
	SkillId: 1244,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.7 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 3, 1244, 1, false))
		}
		return buffs
	},
}

// 黄粱一梦3级
var role_skill_1245 = &Skill{
	SkillId: 1245,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 3, 1245, 1, false))
		}
		return buffs
	},
}

// 黄粱一梦4级
var role_skill_1246 = &Skill{
	SkillId: 1246,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 4, 1246, 1, false))
		}
		return buffs
	},
}

// 黄粱一梦5级
var role_skill_1247 = &Skill{
	SkillId: 1247,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 5, 1247, 1, false))
		}
		return buffs
	},
}

// 蹑影追风·怪
var role_skill_1248 = &Skill{
	SkillId: 1248,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 1, 1248, 1, false))
			}
		}
		return buffs
	},
}

// 蹑影追风3级
var role_skill_1249 = &Skill{
	SkillId: 1249,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 1, 1249, 1, false))
			}
		}
		return buffs
	},
}

// 蹑影追风4级
var role_skill_1250 = &Skill{
	SkillId: 1250,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 1, 1250, 1, false))
			}
		}
		return buffs
	},
}

// 蹑影追风5级
var role_skill_1251 = &Skill{
	SkillId: 1251,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 1, 1251, 1, false))
			}
		}
		return buffs
	},
}

// 雷动九天
var role_skill_1252 = &Skill{
	SkillId: 1252,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
}

// 雷动九天·怪
var role_skill_1253 = &Skill{
	SkillId: 1253,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷动九天3级
var role_skill_1254 = &Skill{
	SkillId: 1254,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷动九天4级
var role_skill_1255 = &Skill{
	SkillId: 1255,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷动九天5级
var role_skill_1256 = &Skill{
	SkillId: 1256,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 妙笔生花
var role_skill_1257 = &Skill{
	SkillId: 1257,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
}

// 妙笔生花·怪
var role_skill_1258 = &Skill{
	SkillId: 1258,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 妙笔生花3级
var role_skill_1259 = &Skill{
	SkillId: 1259,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 妙笔生花4级
var role_skill_1260 = &Skill{
	SkillId: 1260,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 妙笔生花5级
var role_skill_1261 = &Skill{
	SkillId: 1261,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风驰云卷
var role_skill_1262 = &Skill{
	SkillId: 1262,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
}

// 风驰云卷·怪
var role_skill_1263 = &Skill{
	SkillId: 1263,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风驰云卷3级
var role_skill_1264 = &Skill{
	SkillId: 1264,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风驰云卷4级
var role_skill_1265 = &Skill{
	SkillId: 1265,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 风驰云卷5级
var role_skill_1266 = &Skill{
	SkillId: 1266,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 神魅鬼目
var role_skill_1267 = &Skill{
	SkillId: 1267,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1267, 1, false))
		}
		return buffs
	},
}

// 神魅鬼目·怪
var role_skill_1268 = &Skill{
	SkillId: 1268,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1268, 1, false))
		}
		return buffs
	},
}

// 神魅鬼目3级
var role_skill_1269 = &Skill{
	SkillId: 1269,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.6 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1269, 1, false))
		}
		return buffs
	},
}

// 神魅鬼目4级
var role_skill_1270 = &Skill{
	SkillId: 1270,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.7 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1270, 1, false))
		}
		return buffs
	},
}

// 神魅鬼目5级
var role_skill_1271 = &Skill{
	SkillId: 1271,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1271, 1, false))
		}
		return buffs
	},
}

// 生生不息
var role_skill_1272 = &Skill{
	SkillId: 1272,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(1000 + float64(skillTrnlv) * 150 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1272, 1, false))
			}
		}
		return buffs
	},
}

// 剑十
var role_skill_1273 = &Skill{
	SkillId: 1273,
	ChildType: 1,
	DecPower: 3,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findCrossTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 召唤开心莲藕精
var role_skill_1274 = &Skill{
	SkillId: 1274,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 550,4},CallInfo{ 550,3},CallInfo{ 550,1},CallInfo{ 550,0}},
}

// 英勇
var role_skill_1275 = &Skill{
	SkillId: 1275,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 霸刀
var role_skill_1276 = &Skill{
	SkillId: 1276,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 轰击
var role_skill_1277 = &Skill{
	SkillId: 1277,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 阿修罗之怒附属
var role_skill_1278 = &Skill{
	SkillId: 1278,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 3000,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 银光落刃
var role_skill_1281 = &Skill{
	SkillId: 1281,
	ChildType: 0,
	FixedValue: 0,
}

// 银光落刃
var role_skill_1284 = &Skill{
	SkillId: 1284,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 式神炎舞
var role_skill_1287 = &Skill{
	SkillId: 1287,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 式神炎舞附属
var role_skill_1288 = &Skill{
	SkillId: 1288,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1000 + float64(skillTrnlv) * 50)), 2, 1288, 1, false))
		return buffs
	},
}

// 猛龙断空
var role_skill_1293 = &Skill{
	SkillId: 1293,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	Critial: 200,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 猛龙断空附属
var role_skill_1294 = &Skill{
	SkillId: 1294,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 2000,
	SunderAttack: 5,
	Critial: 200,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1294, 1, false))
		}
		return buffs
	},
}

// 影舞
var role_skill_1299 = &Skill{
	SkillId: 1299,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 流萤斩
var role_skill_1300 = &Skill{
	SkillId: 1300,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 50,
	Critial: 300,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 烟雾弹
var role_skill_1301 = &Skill{
	SkillId: 1301,
	ChildType: 1,
	FixedValue: 500,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.9 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(500 + float64(skillTrnlv) * 10)), 2, 1301, 1, false))
		return buffs
	},
}

// 暗影火袭
var role_skill_1302 = &Skill{
	SkillId: 1302,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 4000,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 0.9 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 凌云斩
var role_skill_1303 = &Skill{
	SkillId: 1303,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 2000,
	SunderAttack: 5,
	Critial: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风之守护
var role_skill_1304 = &Skill{
	SkillId: 1304,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 1, 1304, 1, false))
			}
		}
		return buffs
	},
}

// 大风咒2级
var role_skill_1305 = &Skill{
	SkillId: 1305,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风卷尘生2级
var role_skill_1306 = &Skill{
	SkillId: 1306,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 妙法莲华2级
var role_skill_1307 = &Skill{
	SkillId: 1307,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffCount = buddy1.GetBuffCount(3)
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(15000 + float64(skillTrnlv) * 50 + f.Cultivation - buffCount * 10000 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1307, 1, false))
					buddy1.AddBuffCount(3)
			}
		}
		return buffs
	},
}

// 一滴入魂
var role_skill_1308 = &Skill{
	SkillId: 1308,
	ChildType: 3,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1000 + float64(skillTrnlv) * 10)), 3, 1308, 2, false))
			}
		}
		return buffs
	},
}

// 纯阳罡气2级
var role_skill_1309 = &Skill{
	SkillId: 1309,
	ChildType: 3,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ABSORB_HURT, (buff_value(20000 + float64(skillTrnlv) * 50)), 2, 1309, 1, false))
			}
		}
		return buffs
	},
}

// 护盾破坏2级
var role_skill_1310 = &Skill{
	SkillId: 1310,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 10,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1310, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1310, 1, false))
		return buffs
	},
}

// 藤甲术2级
var role_skill_1311 = &Skill{
	SkillId: 1311,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(1000 + float64(skillTrnlv) * 20)), 1, 1311, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH && buddy2.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_SUNDER, (buff_value(200)), 0, 1311, 1, false))
			}
		}
		return buffs
	},
}

// 灵竹咒
var role_skill_1312 = &Skill{
	SkillId: 1312,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 50,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1000 + float64(skillTrnlv) * 10 + float64(hurt) * 0.2)), 2, 1312, 1, false))
		return buffs
	},
}

// 神魔封禁
var role_skill_1313 = &Skill{
	SkillId: 1313,
	ChildType: 5,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.7 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1313, 1, false))
		}
		return buffs
	},
}

// 墨画魂刃2级
var role_skill_1314 = &Skill{
	SkillId: 1314,
	ChildType: 1,
	FixedValue: 3000,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 火炎咒
var role_skill_1315 = &Skill{
	SkillId: 1315,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 神魅鬼目
var role_skill_1316 = &Skill{
	SkillId: 1316,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1316, 1, false))
		}
		return buffs
	},
}

// 召唤唐家护卫
var role_skill_1317 = &Skill{
	SkillId: 1317,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 567,12},CallInfo{ 567,8},CallInfo{ 567,6},CallInfo{ 567,2}},
}

// 落云3级
var role_skill_1318 = &Skill{
	SkillId: 1318,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 落云4级
var role_skill_1319 = &Skill{
	SkillId: 1319,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 凌云斩2级
var role_skill_1320 = &Skill{
	SkillId: 1320,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 4000,
	SunderAttack: 5,
	Critial: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 落云5级
var role_skill_1321 = &Skill{
	SkillId: 1321,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 落云6级
var role_skill_1322 = &Skill{
	SkillId: 1322,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 落云7级
var role_skill_1323 = &Skill{
	SkillId: 1323,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 落云8级
var role_skill_1324 = &Skill{
	SkillId: 1324,
	ChildType: 1,
	DecPower: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰3级
var role_skill_1325 = &Skill{
	SkillId: 1325,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰4级
var role_skill_1326 = &Skill{
	SkillId: 1326,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰5级
var role_skill_1327 = &Skill{
	SkillId: 1327,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰6级
var role_skill_1328 = &Skill{
	SkillId: 1328,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰7级
var role_skill_1329 = &Skill{
	SkillId: 1329,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰8级
var role_skill_1330 = &Skill{
	SkillId: 1330,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰9级
var role_skill_1331 = &Skill{
	SkillId: 1331,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 风镰10级
var role_skill_1332 = &Skill{
	SkillId: 1332,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击3级
var role_skill_1333 = &Skill{
	SkillId: 1333,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击4级
var role_skill_1334 = &Skill{
	SkillId: 1334,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击5级
var role_skill_1335 = &Skill{
	SkillId: 1335,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击6级
var role_skill_1336 = &Skill{
	SkillId: 1336,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击7级
var role_skill_1337 = &Skill{
	SkillId: 1337,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击8级
var role_skill_1338 = &Skill{
	SkillId: 1338,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击9级
var role_skill_1339 = &Skill{
	SkillId: 1339,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩击10级
var role_skill_1340 = &Skill{
	SkillId: 1340,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击2级
var role_skill_1341 = &Skill{
	SkillId: 1341,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击3级
var role_skill_1342 = &Skill{
	SkillId: 1342,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击4级
var role_skill_1343 = &Skill{
	SkillId: 1343,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击5级
var role_skill_1344 = &Skill{
	SkillId: 1344,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击6级
var role_skill_1345 = &Skill{
	SkillId: 1345,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击7级
var role_skill_1346 = &Skill{
	SkillId: 1346,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击8级
var role_skill_1347 = &Skill{
	SkillId: 1347,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击9级
var role_skill_1348 = &Skill{
	SkillId: 1348,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 重击10级
var role_skill_1349 = &Skill{
	SkillId: 1349,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔2级
var role_skill_1350 = &Skill{
	SkillId: 1350,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔3级
var role_skill_1351 = &Skill{
	SkillId: 1351,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔4级
var role_skill_1352 = &Skill{
	SkillId: 1352,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔5级
var role_skill_1353 = &Skill{
	SkillId: 1353,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔6级
var role_skill_1354 = &Skill{
	SkillId: 1354,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔7级
var role_skill_1355 = &Skill{
	SkillId: 1355,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔8级
var role_skill_1356 = &Skill{
	SkillId: 1356,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔9级
var role_skill_1357 = &Skill{
	SkillId: 1357,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 提笔10级
var role_skill_1358 = &Skill{
	SkillId: 1358,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞2级
var role_skill_1359 = &Skill{
	SkillId: 1359,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞3级
var role_skill_1360 = &Skill{
	SkillId: 1360,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞4级
var role_skill_1361 = &Skill{
	SkillId: 1361,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞5级
var role_skill_1362 = &Skill{
	SkillId: 1362,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞6级
var role_skill_1363 = &Skill{
	SkillId: 1363,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞7级
var role_skill_1364 = &Skill{
	SkillId: 1364,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞8级
var role_skill_1365 = &Skill{
	SkillId: 1365,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞9级
var role_skill_1366 = &Skill{
	SkillId: 1366,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 影舞10级
var role_skill_1367 = &Skill{
	SkillId: 1367,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 追骨吸元
var role_skill_1368 = &Skill{
	SkillId: 1368,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.9 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(float64(hurt) * 0.2 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1368, 1, false))
			}
		}
		return buffs
	},
}

// 流萤斩2级
var role_skill_1369 = &Skill{
	SkillId: 1369,
	ChildType: 1,
	FixedValue: 3000,
	SunderAttack: 50,
	Critial: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 烟雾弹2级
var role_skill_1370 = &Skill{
	SkillId: 1370,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 0.9 + float64(skillTrnlv) * 15 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1370, 1, false))
		return buffs
	},
}

// 长拳
var role_skill_1371 = &Skill{
	SkillId: 1371,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 雷行者
var role_skill_1372 = &Skill{
	SkillId: 1372,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1372, 1, false))
		}
		return buffs
	},
}

// 木行者
var role_skill_1373 = &Skill{
	SkillId: 1373,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(5000 + float64(skillTrnlv) * 50 + f.Cultivation + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1373, 1, false))
			}
		}
		return buffs
	},
}

// 横扫
var role_skill_1374 = &Skill{
	SkillId: 1374,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩铁剑
var role_skill_1375 = &Skill{
	SkillId: 1375,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1375, 2, false))
		return buffs
	},
}

// 真元护体
var role_skill_1376 = &Skill{
	SkillId: 1376,
	ChildType: 1,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(2000 + float64(skillTrnlv) * 20)), 2, 1376, 1, false))
			}
		}
		return buffs
	},
}

// 寒霜
var role_skill_1377 = &Skill{
	SkillId: 1377,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 冰锥术
var role_skill_1378 = &Skill{
	SkillId: 1378,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1378, 2, false))
		return buffs
	},
}

// 冰心诀
var role_skill_1379 = &Skill{
	SkillId: 1379,
	ChildType: 1,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1379, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_DEFEND, (buff_value(1000 + float64(skillTrnlv) * 30)), 1, 1379, 1, false))
			}
		}
		return buffs
	},
}

// 突袭
var role_skill_1380 = &Skill{
	SkillId: 1380,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 金蛇锥
var role_skill_1381 = &Skill{
	SkillId: 1381,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 50,
	ReduceDefend: 0.3,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 闪光弹
var role_skill_1382 = &Skill{
	SkillId: 1382,
	ChildType: 1,
	FixedValue: 4000,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 0.9 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DODGE_LEVEL, -(buff_value(500 + float64(skillTrnlv) * 10)), 2, 1382, 1, false))
		return buffs
	},
}

// 长拳2级
var role_skill_1383 = &Skill{
	SkillId: 1383,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳3级
var role_skill_1384 = &Skill{
	SkillId: 1384,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳4级
var role_skill_1385 = &Skill{
	SkillId: 1385,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳5级
var role_skill_1386 = &Skill{
	SkillId: 1386,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳6级
var role_skill_1387 = &Skill{
	SkillId: 1387,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳7级
var role_skill_1388 = &Skill{
	SkillId: 1388,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳8级
var role_skill_1389 = &Skill{
	SkillId: 1389,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳9级
var role_skill_1390 = &Skill{
	SkillId: 1390,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长拳10级
var role_skill_1391 = &Skill{
	SkillId: 1391,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫2级
var role_skill_1392 = &Skill{
	SkillId: 1392,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫3级
var role_skill_1393 = &Skill{
	SkillId: 1393,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫4级
var role_skill_1394 = &Skill{
	SkillId: 1394,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫5级
var role_skill_1395 = &Skill{
	SkillId: 1395,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫6级
var role_skill_1396 = &Skill{
	SkillId: 1396,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫7级
var role_skill_1397 = &Skill{
	SkillId: 1397,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫8级
var role_skill_1398 = &Skill{
	SkillId: 1398,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫9级
var role_skill_1399 = &Skill{
	SkillId: 1399,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 横扫10级
var role_skill_1400 = &Skill{
	SkillId: 1400,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜2级
var role_skill_1401 = &Skill{
	SkillId: 1401,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜3级
var role_skill_1402 = &Skill{
	SkillId: 1402,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜4级
var role_skill_1403 = &Skill{
	SkillId: 1403,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜5级
var role_skill_1404 = &Skill{
	SkillId: 1404,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜6级
var role_skill_1405 = &Skill{
	SkillId: 1405,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜7级
var role_skill_1406 = &Skill{
	SkillId: 1406,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜8级
var role_skill_1407 = &Skill{
	SkillId: 1407,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜9级
var role_skill_1408 = &Skill{
	SkillId: 1408,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒霜10级
var role_skill_1409 = &Skill{
	SkillId: 1409,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭2级
var role_skill_1410 = &Skill{
	SkillId: 1410,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭3级
var role_skill_1411 = &Skill{
	SkillId: 1411,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭4级
var role_skill_1412 = &Skill{
	SkillId: 1412,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭5级
var role_skill_1413 = &Skill{
	SkillId: 1413,
	ChildType: 1,
	FixedValue: 500,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭6级
var role_skill_1414 = &Skill{
	SkillId: 1414,
	ChildType: 1,
	FixedValue: 1000,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突击7级
var role_skill_1415 = &Skill{
	SkillId: 1415,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭8级
var role_skill_1416 = &Skill{
	SkillId: 1416,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭9级
var role_skill_1417 = &Skill{
	SkillId: 1417,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 突袭10级
var role_skill_1418 = &Skill{
	SkillId: 1418,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 断岳
var role_skill_1419 = &Skill{
	SkillId: 1419,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 无限霸刀附属
var role_skill_1420 = &Skill{
	SkillId: 1420,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 80,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 人鱼之歌附属
var role_skill_1421 = &Skill{
	SkillId: 1421,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 500,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1421, 1, false))
			}
		}
		return buffs
	},
}

// 横扫千军附属
var role_skill_1422 = &Skill{
	SkillId: 1422,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 90,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 碎铁之刃附属
var role_skill_1423 = &Skill{
	SkillId: 1423,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 100,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 瞒天过海附属
var role_skill_1424 = &Skill{
	SkillId: 1424,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 90,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 银光落刃附属
var role_skill_1425 = &Skill{
	SkillId: 1425,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 2, 1425, 1, false))
			}
		}
		return buffs
	},
}

// 流荧斩
var role_skill_1426 = &Skill{
	SkillId: 1426,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	Critial: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 烟雾弹
var role_skill_1427 = &Skill{
	SkillId: 1427,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(500)), 1, 1427, 1, false))
		return buffs
	},
}

// 灵竹咒
var role_skill_1428 = &Skill{
	SkillId: 1428,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 2, 1428, 1, false))
		return buffs
	},
}

// 追骨吸元
var role_skill_1429 = &Skill{
	SkillId: 1429,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(float64(hurt) * 0.2 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1429, 1, false))
			}
		}
		return buffs
	},
}

// 凌云斩
var role_skill_1430 = &Skill{
	SkillId: 1430,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 顺势斩
var role_skill_1431 = &Skill{
	SkillId: 1431,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 天剑
var role_skill_1432 = &Skill{
	SkillId: 1432,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 召唤堕落毒蛇
var role_skill_1434 = &Skill{
	SkillId: 1434,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 666,3},CallInfo{ 666,2},CallInfo{ 666,1}},
}

// 召唤堕落黑翼巨蝠
var role_skill_1435 = &Skill{
	SkillId: 1435,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 27,13},CallInfo{ 27,11}},
}

// 召唤堕落燃魁
var role_skill_1436 = &Skill{
	SkillId: 1436,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 667,9},CallInfo{ 667,6}},
}

// 召唤堕落狼牙血杀
var role_skill_1437 = &Skill{
	SkillId: 1437,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 669,8},CallInfo{ 669,6}},
}

// 召唤堕落重刀血杀
var role_skill_1438 = &Skill{
	SkillId: 1438,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 670,3},CallInfo{ 670,1}},
}

// 召唤堕落野猪、堕落穿
var role_skill_1439 = &Skill{
	SkillId: 1439,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 671,8},CallInfo{ 671,6},CallInfo{ 672,4},CallInfo{ 672,0}},
}

// 召唤堕落火豹
var role_skill_1441 = &Skill{
	SkillId: 1441,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 673,3},CallInfo{ 673,1}},
}

// 召唤堕落不死亡魂
var role_skill_1442 = &Skill{
	SkillId: 1442,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 674,8},CallInfo{ 674,6},CallInfo{ 674,2}},
}

// 召唤堕落僵尸
var role_skill_1443 = &Skill{
	SkillId: 1443,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 675,9},CallInfo{ 675,7},CallInfo{ 675,3}},
}

// 召唤堕落僵尸2
var role_skill_1444 = &Skill{
	SkillId: 1444,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 675,4},CallInfo{ 675,2}},
}

// 召唤堕落拳之守卫
var role_skill_1445 = &Skill{
	SkillId: 1445,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 676,9},CallInfo{ 676,8},CallInfo{ 676,6},CallInfo{ 676,5}},
}

// 召唤堕落黑狼、堕落野
var role_skill_1446 = &Skill{
	SkillId: 1446,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 677,13},CallInfo{ 677,11},CallInfo{ 677,9},CallInfo{ 677,5}},
}

// 召唤堕落玉石守卫
var role_skill_1447 = &Skill{
	SkillId: 1447,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 678,4},CallInfo{ 678,1}},
}

// 召唤堕落撼地兽
var role_skill_1448 = &Skill{
	SkillId: 1448,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 679,14},CallInfo{ 679,11}},
}

// 召唤堕落雷兽
var role_skill_1449 = &Skill{
	SkillId: 1449,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 680,8},CallInfo{ 680,6}},
}

// 召唤堕落冥界僵尸
var role_skill_1450 = &Skill{
	SkillId: 1450,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 681,9},CallInfo{ 681,5},CallInfo{ 681,3},CallInfo{ 681,1}},
}

// 召唤堕落弓之守卫
var role_skill_1451 = &Skill{
	SkillId: 1451,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 682,13},CallInfo{ 682,11},CallInfo{ 682,9},CallInfo{ 682,5}},
}

// 召唤堕落钩镰血杀
var role_skill_1452 = &Skill{
	SkillId: 1452,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 330,8},CallInfo{ 330,6},CallInfo{ 330,3},CallInfo{ 330,1}},
}

// 召唤堕落残风
var role_skill_1453 = &Skill{
	SkillId: 1453,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 683,8},CallInfo{ 683,6}},
}

// 召唤堕落毒兽、堕落毒
var role_skill_1454 = &Skill{
	SkillId: 1454,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 685,13},CallInfo{ 685,11},CallInfo{ 684,3},CallInfo{ 684,1}},
}

// 召唤堕落阴毒血士、堕
var role_skill_1455 = &Skill{
	SkillId: 1455,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 686,9},CallInfo{ 686,5},CallInfo{ 687,3},CallInfo{ 687,1}},
}

// 召唤堕落利爪蝙蝠
var role_skill_1456 = &Skill{
	SkillId: 1456,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 688,3},CallInfo{ 688,1}},
}

// 召唤堕落血蝙蝠杀手
var role_skill_1457 = &Skill{
	SkillId: 1457,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 689,8},CallInfo{ 689,6}},
}

// 召唤堕落水鬼
var role_skill_1458 = &Skill{
	SkillId: 1458,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 690,4},CallInfo{ 690,3},CallInfo{ 690,1},CallInfo{ 690,0}},
}

// 召唤堕落幕府重甲武士
var role_skill_1459 = &Skill{
	SkillId: 1459,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 691,3},CallInfo{ 691,1}},
}

// 召唤堕落幕府长弓武士
var role_skill_1460 = &Skill{
	SkillId: 1460,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 692,13},CallInfo{ 692,11}},
}

// 召唤堕落血杀武士
var role_skill_1461 = &Skill{
	SkillId: 1461,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 693,8},CallInfo{ 693,6},CallInfo{ 693,2}},
}

// 召唤堕落甲贺忍者
var role_skill_1462 = &Skill{
	SkillId: 1462,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 526,13},CallInfo{ 526,11}},
}

// 召唤堕落血妖
var role_skill_1463 = &Skill{
	SkillId: 1463,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 694,9},CallInfo{ 694,5},CallInfo{ 694,3},CallInfo{ 694,1}},
}

// 召唤堕落血巫
var role_skill_1464 = &Skill{
	SkillId: 1464,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 536,13},CallInfo{ 536,11}},
}

// 召唤堕落唐家护卫
var role_skill_1465 = &Skill{
	SkillId: 1465,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 696,3},CallInfo{ 696,1}},
}

// 召唤堕落唐家女护卫
var role_skill_1466 = &Skill{
	SkillId: 1466,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 572,12},CallInfo{ 572,8},CallInfo{ 572,6}},
}

// 召唤堕落龙虎门人
var role_skill_1467 = &Skill{
	SkillId: 1467,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 697,8},CallInfo{ 697,6}},
}

// 召唤堕落恶犬
var role_skill_1468 = &Skill{
	SkillId: 1468,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 698,-2}},
}

// 召唤堕落剑影
var role_skill_1469 = &Skill{
	SkillId: 1469,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 699,9},CallInfo{ 699,5},CallInfo{ 699,3},CallInfo{ 699,1}},
}

// 召唤堕落玄音寺武僧
var role_skill_1470 = &Skill{
	SkillId: 1470,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 700,8},CallInfo{ 700,6},CallInfo{ 700,4},CallInfo{ 700,0}},
}

// 召唤堕落刀殿弟子
var role_skill_1471 = &Skill{
	SkillId: 1471,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 701,4},CallInfo{ 701,3},CallInfo{ 701,1},CallInfo{ 701,0}},
}

// 召唤堕落灯笼怪
var role_skill_1472 = &Skill{
	SkillId: 1472,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 665,3},CallInfo{ 665,1}},
}

// 召唤堕落剑之守卫
var role_skill_1473 = &Skill{
	SkillId: 1473,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 668,3},CallInfo{ 668,1}},
}

// 气疗术
var role_skill_1474 = &Skill{
	SkillId: 1474,
	ChildType: 4,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1474, 1, false))
		return buffs
	},
}

// 毒伤
var role_skill_1475 = &Skill{
	SkillId: 1475,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 0, 1475, 0, false))
		return buffs
	},
}

// 分裂箭
var role_skill_1476 = &Skill{
	SkillId: 1476,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 破甲利爪
var role_skill_1477 = &Skill{
	SkillId: 1477,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 背刺
var role_skill_1478 = &Skill{
	SkillId: 1478,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 腐蚀攻击1
var role_skill_1479 = &Skill{
	SkillId: 1479,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(250)), 4, 1479, 3, false))
		return buffs
	},
}

// 居合
var role_skill_1480 = &Skill{
	SkillId: 1480,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 200,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000)), 2, 1480, 2, false))
		return buffs
	},
}

// 会心一击
var role_skill_1481 = &Skill{
	SkillId: 1481,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(200)), 4, 1481, 0, false))
			}
		}
		return buffs
	},
}

// 达摩棍
var role_skill_1482 = &Skill{
	SkillId: 1482,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.1 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1482, 1, false))
		}
		return buffs
	},
}

// 腐蚀剧毒
var role_skill_1483 = &Skill{
	SkillId: 1483,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 3, 1483, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(500)), 3, 1483, 1, false))
		return buffs
	},
}

// 神魅鬼目（群体）
var role_skill_1484 = &Skill{
	SkillId: 1484,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1484, 1, false))
		}
		return buffs
	},
}

// 落云斩
var role_skill_1485 = &Skill{
	SkillId: 1485,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 嗜血
var role_skill_1486 = &Skill{
	SkillId: 1486,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTACK, (buff_value(1 + force * 1.3)), 3, 1486, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1486, 1, false))
		return buffs
	},
}

// 嘲讽
var role_skill_1487 = &Skill{
	SkillId: 1487,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ABSORB_HURT, (buff_value(1 + force)), 2, 1487, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 2, 1487, 1, false))
		return buffs
	},
}

// 腐蚀攻击2
var role_skill_1488 = &Skill{
	SkillId: 1488,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(500)), 4, 1488, 3, false))
		return buffs
	},
}

// 腐蚀攻击3
var role_skill_1489 = &Skill{
	SkillId: 1489,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(800)), 4, 1489, 3, false))
		return buffs
	},
}

// 人鱼之歌
var role_skill_1490 = &Skill{
	SkillId: 1490,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 3, 1490, 1, false))
		return buffs
	},
}

// 斩铁剑
var role_skill_1491 = &Skill{
	SkillId: 1491,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 冰锥术
var role_skill_1492 = &Skill{
	SkillId: 1492,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1000)), 2, 1492, 2, false))
		return buffs
	},
}

// 金蛇锥
var role_skill_1493 = &Skill{
	SkillId: 1493,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	ReduceDefend: 0.3,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 闪光弹
var role_skill_1494 = &Skill{
	SkillId: 1494,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DODGE_LEVEL, -(buff_value(500)), 1, 1494, 1, false))
		return buffs
	},
}

// 木行者
var role_skill_1495 = &Skill{
	SkillId: 1495,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1495, 1, false))
			}
		}
		return buffs
	},
}

// 真元护体
var role_skill_1496 = &Skill{
	SkillId: 1496,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(1000)), 2, 1496, 2, false))
			}
		}
		return buffs
	},
}

// 冰心诀
var role_skill_1497 = &Skill{
	SkillId: 1497,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1497, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_DEFEND, (buff_value(1 + force)), 1, 1497, 1, false))
			}
		}
		return buffs
	},
}

// 雷行者
var role_skill_1498 = &Skill{
	SkillId: 1498,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1498, 1, false))
		}
		return buffs
	},
}

// 召唤机灵飞鹏
var role_skill_1499 = &Skill{
	SkillId: 1499,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 712,8},CallInfo{ 712,6},CallInfo{ 711,4},CallInfo{ 711,0}},
}

// 血爆
var role_skill_1500 = &Skill{
	SkillId: 1500,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 2, 1500, 1, false))
		return buffs
	},
}

// 血爆连锁
var role_skill_1501 = &Skill{
	SkillId: 1501,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 60,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000)), 2, 1501, 1, false))
		return buffs
	},
}

// 血祭
var role_skill_1502 = &Skill{
	SkillId: 1502,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_GHOST_POWER, -(buff_value(50)), 0, 1502, 1, false))
		return buffs
	},
}

// 闪惊雷
var role_skill_1503 = &Skill{
	SkillId: 1503,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 毒爆
var role_skill_1504 = &Skill{
	SkillId: 1504,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 150,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 2, 1504, 1, false))
		return buffs
	},
}

// 铁拳
var role_skill_1505 = &Skill{
	SkillId: 1505,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1505, 1, false))
		}
		return buffs
	},
}

// 钢拳
var role_skill_1506 = &Skill{
	SkillId: 1506,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1506, 1, false))
		}
		return buffs
	},
}

// 钛合金拳
var role_skill_1507 = &Skill{
	SkillId: 1507,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 150,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1507, 1, false))
		}
		return buffs
	},
}

// 暗影新星
var role_skill_1508 = &Skill{
	SkillId: 1508,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HEALTH, (buff_value(float64(hurt) * 0.2 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1508, 1, false))
		return buffs
	},
}

// 暗火之雨
var role_skill_1509 = &Skill{
	SkillId: 1509,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 3, 1509, 2, false))
		return buffs
	},
}

// 地突刺
var role_skill_1510 = &Skill{
	SkillId: 1510,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 魔瞳术
var role_skill_1511 = &Skill{
	SkillId: 1511,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1511, 1, false))
		}
		return buffs
	},
}

// 火焰冲击
var role_skill_1512 = &Skill{
	SkillId: 1512,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1512, 1, false))
		}
		return buffs
	},
}

// 业火
var role_skill_1513 = &Skill{
	SkillId: 1513,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 2, 1513, 2, false))
		return buffs
	},
}

// 腐蚀剧毒2
var role_skill_1514 = &Skill{
	SkillId: 1514,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(float64(hurt) * 0.3)), 3, 1514, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000)), 3, 1514, 2, false))
		return buffs
	},
}

// 爆怒
var role_skill_1515 = &Skill{
	SkillId: 1515,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTACK, (buff_value(force)), 9, 1515, 3, false))
		return buffs
	},
}

// 狂狮怒吼
var role_skill_1516 = &Skill{
	SkillId: 1516,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 星落
var role_skill_1517 = &Skill{
	SkillId: 1517,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 3, 1517, 2, false))
		return buffs
	},
}

// 无限爆怒
var role_skill_1518 = &Skill{
	SkillId: 1518,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTACK, (buff_value(force)), 0, 1518, 0, false))
		return buffs
	},
}

// 乾坤刀气X2
var role_skill_1519 = &Skill{
	SkillId: 1519,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 召唤堕落嗜血阴毒血士
var role_skill_1520 = &Skill{
	SkillId: 1520,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 720,13},CallInfo{ 720,9},CallInfo{ 720,6}},
}

// 召唤堕落嗜血阴影
var role_skill_1521 = &Skill{
	SkillId: 1521,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 721,13},CallInfo{ 721,3}},
}

// 召唤堕落忆境梦魇
var role_skill_1522 = &Skill{
	SkillId: 1522,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 722,13},CallInfo{ 722,11}},
}

// 铁壁
var role_skill_1523 = &Skill{
	SkillId: 1523,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_BLOCK, (buff_value(1000 + float64(skillTrnlv) * 6)), 2, 1523, 1, false))
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 2, 1523, 1, false))
		return buffs
	},
}

// 冰咒
var role_skill_1524 = &Skill{
	SkillId: 1524,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAN_GOOD, 1, 0, 1524, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1524, 2, false))
		return buffs
	},
}

// 暴雨梨花针
var role_skill_1525 = &Skill{
	SkillId: 1525,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋
var role_skill_1526 = &Skill{
	SkillId: 1526,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 伏魔伞
var role_skill_1527 = &Skill{
	SkillId: 1527,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 2, 1527, 1, false))
		}
		return buffs
	},
}

// 神兵卷
var role_skill_1528 = &Skill{
	SkillId: 1528,
	ChildType: 1,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1000 + float64(skillTrnlv) * 20)), 2, 1528, 2, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 2, 1528, 2, false))
			}
		}
		return buffs
	},
}

// 灵旋2级
var role_skill_1529 = &Skill{
	SkillId: 1529,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋3级
var role_skill_1530 = &Skill{
	SkillId: 1530,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋4级
var role_skill_1531 = &Skill{
	SkillId: 1531,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋5级
var role_skill_1532 = &Skill{
	SkillId: 1532,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋6级
var role_skill_1533 = &Skill{
	SkillId: 1533,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋7级
var role_skill_1534 = &Skill{
	SkillId: 1534,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋8级
var role_skill_1535 = &Skill{
	SkillId: 1535,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋9级
var role_skill_1536 = &Skill{
	SkillId: 1536,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 灵旋10级
var role_skill_1537 = &Skill{
	SkillId: 1537,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 青
var role_skill_1538 = &Skill{
	SkillId: 1538,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(50 + float64(skillTrnlv) * 5)), 0, 1538, 1, false))
			}
		}
		return buffs
	},
}

// 青
var role_skill_1539 = &Skill{
	SkillId: 1539,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(50 + float64(skillTrnlv) * 5)), 0, 1539, 1, false))
			}
		}
		return buffs
	},
}

// 青
var role_skill_1540 = &Skill{
	SkillId: 1540,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(50 + float64(skillTrnlv) * 5)), 0, 1540, 1, false))
			}
		}
		return buffs
	},
}

// 兰
var role_skill_1541 = &Skill{
	SkillId: 1541,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(20000 + float64(skillTrnlv) * 2500 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1541, 1, false))
			}
		}
		return buffs
	},
}

// 兰
var role_skill_1542 = &Skill{
	SkillId: 1542,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(20000 + float64(skillTrnlv) * 2500 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1542, 1, false))
			}
		}
		return buffs
	},
}

// 兰
var role_skill_1543 = &Skill{
	SkillId: 1543,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(20000 + float64(skillTrnlv) * 2500 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1543, 1, false))
			}
		}
		return buffs
	},
}

// 莹
var role_skill_1544 = &Skill{
	SkillId: 1544,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1544, 1, false))
			}
		}
		return buffs
	},
}

// 莹
var role_skill_1545 = &Skill{
	SkillId: 1545,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1545, 1, false))
			}
		}
		return buffs
	},
}

// 莹
var role_skill_1546 = &Skill{
	SkillId: 1546,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1546, 1, false))
			}
		}
		return buffs
	},
}

// 赤
var role_skill_1547 = &Skill{
	SkillId: 1547,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1547, 1, false))
			}
		}
		return buffs
	},
}

// 赤
var role_skill_1548 = &Skill{
	SkillId: 1548,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1548, 1, false))
			}
		}
		return buffs
	},
}

// 赤
var role_skill_1549 = &Skill{
	SkillId: 1549,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1549, 1, false))
			}
		}
		return buffs
	},
}

// 白
var role_skill_1550 = &Skill{
	SkillId: 1550,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(20 + float64(skillTrnlv) * 2)), 1, 1550, 1, false))
			}
		}
		return buffs
	},
}

// 白
var role_skill_1551 = &Skill{
	SkillId: 1551,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(20 + float64(skillTrnlv) * 2)), 1, 1551, 1, false))
			}
		}
		return buffs
	},
}

// 白
var role_skill_1552 = &Skill{
	SkillId: 1552,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(20 + float64(skillTrnlv) * 2)), 1, 1552, 1, false))
			}
		}
		return buffs
	},
}

// 缪
var role_skill_1556 = &Skill{
	SkillId: 1556,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1556, 1, false))
			}
		}
		return buffs
	},
}

// 缪
var role_skill_1557 = &Skill{
	SkillId: 1557,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1557, 1, false))
			}
		}
		return buffs
	},
}

// 缪
var role_skill_1558 = &Skill{
	SkillId: 1558,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 50)), 1, 1558, 1, false))
			}
		}
		return buffs
	},
}

// 风面
var role_skill_1559 = &Skill{
	SkillId: 1559,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(100 + float64(skillTrnlv) * 25)), 0, 1559, 1, false))
			}
		}
		return buffs
	},
}

// 风面
var role_skill_1560 = &Skill{
	SkillId: 1560,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(100 + float64(skillTrnlv) * 25)), 0, 1560, 1, false))
			}
		}
		return buffs
	},
}

// 风面
var role_skill_1561 = &Skill{
	SkillId: 1561,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(100 + float64(skillTrnlv) * 25)), 0, 1561, 1, false))
			}
		}
		return buffs
	},
}

// 林面
var role_skill_1562 = &Skill{
	SkillId: 1562,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(40 + float64(skillTrnlv) * 5)), 1, 1562, 1, false))
			}
		}
		return buffs
	},
}

// 林面
var role_skill_1563 = &Skill{
	SkillId: 1563,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(40 + float64(skillTrnlv) * 5)), 1, 1563, 1, false))
			}
		}
		return buffs
	},
}

// 林面
var role_skill_1564 = &Skill{
	SkillId: 1564,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(40 + float64(skillTrnlv) * 5)), 1, 1564, 1, false))
			}
		}
		return buffs
	},
}

// 雷面
var role_skill_1565 = &Skill{
	SkillId: 1565,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1565, 1, false))
			}
		}
		return buffs
	},
}

// 雷面
var role_skill_1566 = &Skill{
	SkillId: 1566,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1566, 1, false))
			}
		}
		return buffs
	},
}

// 雷面
var role_skill_1567 = &Skill{
	SkillId: 1567,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1567, 1, false))
			}
		}
		return buffs
	},
}

// 山面
var role_skill_1568 = &Skill{
	SkillId: 1568,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1568, 1, false))
			}
		}
		return buffs
	},
}

// 山面
var role_skill_1569 = &Skill{
	SkillId: 1569,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1569, 1, false))
			}
		}
		return buffs
	},
}

// 山面
var role_skill_1570 = &Skill{
	SkillId: 1570,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_BLOCK, (buff_value(1000 + float64(skillTrnlv) * 250)), 1, 1570, 1, false))
			}
		}
		return buffs
	},
}

// 火面
var role_skill_1571 = &Skill{
	SkillId: 1571,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(50 + float64(skillTrnlv) * 10)), 1, 1571, 1, false))
			}
		}
		return buffs
	},
}

// 火面
var role_skill_1572 = &Skill{
	SkillId: 1572,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(50 + float64(skillTrnlv) * 10)), 1, 1572, 1, false))
			}
		}
		return buffs
	},
}

// 火面
var role_skill_1573 = &Skill{
	SkillId: 1573,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(50 + float64(skillTrnlv) * 10)), 1, 1573, 1, false))
			}
		}
		return buffs
	},
}

// 兰龙
var role_skill_1574 = &Skill{
	SkillId: 1574,
	ChildType: 5,
	FixedValue: 0,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findFixFrontendRow,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return  0   },
	IsTotemSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1574, 1, false))
		if skillTrnlv >= 1 {
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND_PERSENT, -(buff_value(0 + float64(skillTrnlv) * 10)), 1, 1574, 1, false))
		}
		return buffs
	},
}

// 兰龙
var role_skill_1575 = &Skill{
	SkillId: 1575,
	ChildType: 5,
	FixedValue: 0,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findFixMiddleRow,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return  0   },
	IsTotemSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1575, 1, false))
		if skillTrnlv >= 1 {
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND_PERSENT, -(buff_value(0 + float64(skillTrnlv) * 10)), 1, 1575, 1, false))
		}
		return buffs
	},
}

// 兰龙
var role_skill_1576 = &Skill{
	SkillId: 1576,
	ChildType: 5,
	FixedValue: 0,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findFixBackendRow,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return  0   },
	IsTotemSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1576, 1, false))
		if skillTrnlv >= 1 {
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND_PERSENT, -(buff_value(0 + float64(skillTrnlv) * 10)), 1, 1576, 1, false))
		}
		return buffs
	},
}

// 白虎
var role_skill_1577 = &Skill{
	SkillId: 1577,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(80 + float64(skillTrnlv) * 20)), 1, 1577, 1, false))
			}
		}
		return buffs
	},
}

// 白虎
var role_skill_1578 = &Skill{
	SkillId: 1578,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(80 + float64(skillTrnlv) * 20)), 1, 1578, 1, false))
			}
		}
		return buffs
	},
}

// 白虎
var role_skill_1579 = &Skill{
	SkillId: 1579,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_TAKE_SUNDER, (buff_value(80 + float64(skillTrnlv) * 20)), 1, 1579, 1, false))
			}
		}
		return buffs
	},
}

// 景莲
var role_skill_1580 = &Skill{
	SkillId: 1580,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(60000 + float64(skillTrnlv) * 20000 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1580, 1, false))
			}
		}
		return buffs
	},
}

// 景莲
var role_skill_1581 = &Skill{
	SkillId: 1581,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(60000 + float64(skillTrnlv) * 20000 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1581, 1, false))
			}
		}
		return buffs
	},
}

// 景莲
var role_skill_1582 = &Skill{
	SkillId: 1582,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(60000 + float64(skillTrnlv) * 20000 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1582, 1, false))
			}
		}
		return buffs
	},
}

// 苏摩之怒
var role_skill_1583 = &Skill{
	SkillId: 1583,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 3000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
}

// 苏摩之怒附属
var role_skill_1584 = &Skill{
	SkillId: 1584,
	ChildType: 1,
	IsGhostSkill: true,
	FixedValue: 3000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 3000 + f.Cultivation * 1.5 + float64(skillTrnlv) * 50 },
	IsTotemSkill: false, 
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(3000 + float64(skillTrnlv) * 40)), 2, 1584, 1, false))
			}
		}
		return buffs
	},
}

// 伏魔伞
var role_skill_1585 = &Skill{
	SkillId: 1585,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.1 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 2, 1585, 1, false))
		}
		return buffs
	},
}

// 神兵卷
var role_skill_1586 = &Skill{
	SkillId: 1586,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 2, 1586, 2, false))
			}
		}
		return buffs
	},
}

// 影狼突袭
var role_skill_1587 = &Skill{
	SkillId: 1587,
	ChildType: 1,
	FixedValue: 6000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.4 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1587, 1, false))
		}
		return buffs
	},
}

// 冰魂素魄
var role_skill_1588 = &Skill{
	SkillId: 1588,
	ChildType: 1,
	FixedValue: 6000,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DODGE_LEVEL, -(buff_value(1000 + float64(skillTrnlv) * 20)), 2, 1588, 1, false))
		return buffs
	},
}

// 追云逐电
var role_skill_1589 = &Skill{
	SkillId: 1589,
	ChildType: 1,
	FixedValue: 6000,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTargetFromBack,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + float64(f.NaturalSkillLv) * 50 },
	IsTotemSkill: false, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1589, 1, false))
		return buffs
	},
}

// 混沌之息
var role_skill_1590 = &Skill{
	SkillId: 1590,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(10000 + float64(skillTrnlv) * 200 + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1590, 1, false))
			}
		}
		return buffs
	},
}

// 影狼突袭·怪
var role_skill_1591 = &Skill{
	SkillId: 1591,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.4 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1591, 1, false))
		}
		return buffs
	},
}

// 冰魂素魄·怪
var role_skill_1592 = &Skill{
	SkillId: 1592,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DOGE, -(buff_value(60)), 2, 1592, 1, false))
		return buffs
	},
}

// 追云逐电·怪
var role_skill_1593 = &Skill{
	SkillId: 1593,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTargetFromBack,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1593, 1, false))
		return buffs
	},
}

// 混沌之息·怪
var role_skill_1594 = &Skill{
	SkillId: 1594,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_HEALTH, (buff_value(force + ghost.CureAdd * 1)) * (1 + ghost.CureAddRate), 0, 1594, 1, false))
			}
		}
		return buffs
	},
}

// 影狼成群
var role_skill_1595 = &Skill{
	SkillId: 1595,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(2000)), 2, 1595, 1, false))
		return buffs
	},
}

// 巫雀成群（纵向）
var role_skill_1596 = &Skill{
	SkillId: 1596,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 2, 1596, 1, false))
		return buffs
	},
}

// 巫雀成群（横向）
var role_skill_1597 = &Skill{
	SkillId: 1597,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 0, 1597, 1, false))
		return buffs
	},
}

// 腐蚀剧毒3
var role_skill_1598 = &Skill{
	SkillId: 1598,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 3, 1598, 2, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 0, 1598, 1, false))
		return buffs
	},
}

// 青竹咒2阶
var role_skill_1599 = &Skill{
	SkillId: 1599,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 3, 1599, 1, false))
		return buffs
	},
}

// 召唤堕落魔竹筒精
var role_skill_1600 = &Skill{
	SkillId: 1600,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 796,9},CallInfo{ 796,5},CallInfo{ 796,3},CallInfo{ 796,1}},
}

// 召唤堕落画妖
var role_skill_1601 = &Skill{
	SkillId: 1601,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 798,13},CallInfo{ 798,9},CallInfo{ 798,6}},
}

// 召唤堕落魔燃魁
var role_skill_1602 = &Skill{
	SkillId: 1602,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 800,13},CallInfo{ 800,6},CallInfo{ 800,3}},
}

// 召唤堕落魔毒蝎
var role_skill_1603 = &Skill{
	SkillId: 1603,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 802,9},CallInfo{ 802,3},CallInfo{ 802,1}},
}

// 撼地突袭
var role_skill_1605 = &Skill{
	SkillId: 1605,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1605, 1, false))
		}
		return buffs
	},
}

// 枯木逢春
var role_skill_1606 = &Skill{
	SkillId: 1606,
	ChildType: 5,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 撼地突刺
var role_skill_1607 = &Skill{
	SkillId: 1607,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1607, 1, false))
		}
		return buffs
	},
}

// 枯木逢春
var role_skill_1608 = &Skill{
	SkillId: 1608,
	ChildType: 1,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ALL_RESIST, (buff_value(1500 + float64(skillTrnlv) * 15)), 2, 1608, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_REDUCE_HURT, (buff_value(10)), 2, 1608, 1, false))
			}
		}
		return buffs
	},
}

// 天魔拳
var role_skill_1609 = &Skill{
	SkillId: 1609,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 0.9 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT, -(buff_value(10 + float64(skillTrnlv) * 0.2)), 1, 1609, 1, false))
		return buffs
	},
}

// 撼地
var role_skill_1610 = &Skill{
	SkillId: 1610,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地2级
var role_skill_1611 = &Skill{
	SkillId: 1611,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地3级
var role_skill_1612 = &Skill{
	SkillId: 1612,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地4级
var role_skill_1613 = &Skill{
	SkillId: 1613,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地5级
var role_skill_1614 = &Skill{
	SkillId: 1614,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地6级
var role_skill_1615 = &Skill{
	SkillId: 1615,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地7级
var role_skill_1616 = &Skill{
	SkillId: 1616,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地8级
var role_skill_1617 = &Skill{
	SkillId: 1617,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地9级
var role_skill_1618 = &Skill{
	SkillId: 1618,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼地10级
var role_skill_1619 = &Skill{
	SkillId: 1619,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 黑鹰
var role_skill_1620 = &Skill{
	SkillId: 1620,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixFrontendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DISABLE_SKILL_LEVEL, (buff_value(2000 + float64(skillTrnlv) * 500)), 1, 1620, 1, false))
			}
		}
		return buffs
	},
}

// 黑鹰
var role_skill_1621 = &Skill{
	SkillId: 1621,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixMiddleRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DISABLE_SKILL_LEVEL, (buff_value(2000 + float64(skillTrnlv) * 500)), 1, 1621, 1, false))
			}
		}
		return buffs
	},
}

// 黑鹰
var role_skill_1622 = &Skill{
	SkillId: 1622,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := findOurRoleByFixRow(f, fixBackendRowIndex)
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DISABLE_SKILL_LEVEL, (buff_value(2000 + float64(skillTrnlv) * 500)), 1, 1622, 1, false))
			}
		}
		return buffs
	},
}

// 召唤堕落血髅
var role_skill_1623 = &Skill{
	SkillId: 1623,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 844,7},CallInfo{ 844,4},CallInfo{ 844,1}},
}

// 召唤堕落竹叶青
var role_skill_1624 = &Skill{
	SkillId: 1624,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 846,13},CallInfo{ 846,10},CallInfo{ 846,9},CallInfo{ 846,6}},
}

// 召唤堕落盗墓贼
var role_skill_1625 = &Skill{
	SkillId: 1625,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 848,13},CallInfo{ 848,6},CallInfo{ 848,3}},
}

// 召唤堕落锦衣卫
var role_skill_1626 = &Skill{
	SkillId: 1626,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 850,9},CallInfo{ 850,3},CallInfo{ 850,1}},
}

// 多连斩X2
var role_skill_1627 = &Skill{
	SkillId: 1627,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 血爆纵向
var role_skill_1628 = &Skill{
	SkillId: 1628,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(2000)), 2, 1628, 1, false))
		return buffs
	},
}

// 青竹咒纵向
var role_skill_1629 = &Skill{
	SkillId: 1629,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(float64(hurt) * 0.3)), 2, 1629, 1, false))
		return buffs
	},
}

// 青竹咒横向
var role_skill_1630 = &Skill{
	SkillId: 1630,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(float64(hurt) * 0.3)), 2, 1630, 1, false))
		return buffs
	},
}

// 剧毒撕咬
var role_skill_1631 = &Skill{
	SkillId: 1631,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 3, 1631, 1, false))
		return buffs
	},
}

// 破甲撕咬
var role_skill_1632 = &Skill{
	SkillId: 1632,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 200,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1632, 1, false))
		return buffs
	},
}

// 致命撕咬
var role_skill_1633 = &Skill{
	SkillId: 1633,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 1000,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTACK, (buff_value(1 + force * 0.1)), 2, 1633, 1, false))
		return buffs
	},
}

// 乾坤一掷
var role_skill_1634 = &Skill{
	SkillId: 1634,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 50,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 魔瞳术（单体额外）
var role_skill_1635 = &Skill{
	SkillId: 1635,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1635, 1, false))
		}
		return buffs
	},
}

// 一闪2阶
var role_skill_1636 = &Skill{
	SkillId: 1636,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 200,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, (buff_value(3000)), 2, 1636, 2, false))
		return buffs
	},
}

// 会心一击2阶
var role_skill_1637 = &Skill{
	SkillId: 1637,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CRITIAL, (buff_value(1000)), 4, 1637, 0, false))
			}
		}
		return buffs
	},
}

// 暗火之雨2阶
var role_skill_1638 = &Skill{
	SkillId: 1638,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(3000)), 3, 1638, 2, false))
		return buffs
	},
}

// 凶猛奇袭
var role_skill_1639 = &Skill{
	SkillId: 1639,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	Critial: 200,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 长溪
var role_skill_1644 = &Skill{
	SkillId: 1644,
	ChildType: 1,
	FixedValue: 100,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 100 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪2级
var role_skill_1645 = &Skill{
	SkillId: 1645,
	ChildType: 1,
	FixedValue: 200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪3级
var role_skill_1646 = &Skill{
	SkillId: 1646,
	ChildType: 1,
	FixedValue: 300,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 300 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪4级
var role_skill_1647 = &Skill{
	SkillId: 1647,
	ChildType: 1,
	FixedValue: 400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪5级
var role_skill_1648 = &Skill{
	SkillId: 1648,
	ChildType: 1,
	FixedValue: 500,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 500 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪6级
var role_skill_1649 = &Skill{
	SkillId: 1649,
	ChildType: 1,
	FixedValue: 1000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1000 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪7级
var role_skill_1650 = &Skill{
	SkillId: 1650,
	ChildType: 1,
	FixedValue: 1200,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1200 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪8级
var role_skill_1651 = &Skill{
	SkillId: 1651,
	ChildType: 1,
	FixedValue: 1400,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1400 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪9级
var role_skill_1652 = &Skill{
	SkillId: 1652,
	ChildType: 1,
	FixedValue: 1600,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1600 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 长溪10级
var role_skill_1653 = &Skill{
	SkillId: 1653,
	ChildType: 1,
	FixedValue: 1800,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1800 + f.Cultivation * 0.6 + float64(skillTrnlv) * 10 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩龙剑诀
var role_skill_1654 = &Skill{
	SkillId: 1654,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1654, 1, false))
		return buffs
	},
}

// 一剑横空
var role_skill_1655 = &Skill{
	SkillId: 1655,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1655, 1, false))
		return buffs
	},
}

// 断水斩
var role_skill_1656 = &Skill{
	SkillId: 1656,
	ChildType: 6,
	FixedValue: 6000,
	SunderAttack: 300,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 百炼狂刀
var role_skill_1657 = &Skill{
	SkillId: 1657,
	ChildType: 6,
	FixedValue: 4000,
	SunderAttack: 100,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 斩魄刀
var role_skill_1658 = &Skill{
	SkillId: 1658,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTargetWithOneColRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 妙手回春
var role_skill_1659 = &Skill{
	SkillId: 1659,
	ChildType: 4,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(500 + float64(skillTrnlv) * 10)), 2, 1659, 1, false))
			}
		}
		buddies2 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HEALTH_PERCENT, (buff_value(100)), 0, 1659, 1, false))
			}
		}
		return buffs
	},
}

// 风击暴袭
var role_skill_1660 = &Skill{
	SkillId: 1660,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1660, 1, false))
		return buffs
	},
}

// 怒风沙爆
var role_skill_1661 = &Skill{
	SkillId: 1661,
	ChildType: 1,
	FixedValue: 5000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 3)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1661, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DESTROY_LEVEL, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1661, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CRITIAL_LEVEL, -(buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1661, 1, false))
		return buffs
	},
}

// 九转重阳
var role_skill_1662 = &Skill{
	SkillId: 1662,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findDeadFighter(f.side.Fighters)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( true /*buff to death*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REBOTH_HEALTH_PERCENT, (buff_value(50)), 0, 1662, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_POISONING_LEVEL, (buff_value(500 + float64(skillTrnlv) * 20)), 2, 1662, 1, false))
			}
		}
		return buffs
	},
}

// 回春续劲
var role_skill_1663 = &Skill{
	SkillId: 1663,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(200 + float64(skillTrnlv) * 3)), 0, 1663, 1, false))
			}
		}
		return buffs
	},
}

// 青竹神咒
var role_skill_1664 = &Skill{
	SkillId: 1664,
	ChildType: 6,
	FixedValue: 5000,
	SunderAttack: 100,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 5000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(2000 + float64(skillTrnlv) * 20)), 2, 1664, 1, false))
		return buffs
	},
}

// 影狼突袭
var role_skill_1665 = &Skill{
	SkillId: 1665,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1665, 1, false))
		}
		return buffs
	},
}

// 巫雀奇袭
var role_skill_1666 = &Skill{
	SkillId: 1666,
	ChildType: 1,
	FixedValue: 2000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.2 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1666, 1, false))
		}
		return buffs
	},
}

// 墨画灵山
var role_skill_1667 = &Skill{
	SkillId: 1667,
	ChildType: 6,
	FixedValue: 2000,
	SunderAttack: 100,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 2000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 冷云烟
var role_skill_1668 = &Skill{
	SkillId: 1668,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.4 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 1, 1668, 1, false))
		}
		return buffs
	},
}

// 落樱斩
var role_skill_1669 = &Skill{
	SkillId: 1669,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 10)), 3, 1669, 1, false))
		return buffs
	},
}

// 影舞斩
var role_skill_1670 = &Skill{
	SkillId: 1670,
	ChildType: 6,
	FixedValue: 6000,
	SunderAttack: 300,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1670, 1, false))
		return buffs
	},
}

// 万剑长空
var role_skill_1671 = &Skill{
	SkillId: 1671,
	ChildType: 6,
	DecPower: 8,
	FixedValue: 8000,
	SunderAttack: 100,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 般若行龙
var role_skill_1672 = &Skill{
	SkillId: 1672,
	ChildType: 6,
	DecPower: 8,
	FixedValue: 6000,
	SunderAttack: 400,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 30 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 收妖诀
var role_skill_1673 = &Skill{
	SkillId: 1673,
	ChildType: 1,
	DecPower: 6,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 水行者
var role_skill_1674 = &Skill{
	SkillId: 1674,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1674, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_DEFEND, (buff_value(3000 + float64(skillTrnlv) * 20)), 2, 1674, 1, false))
			}
		}
		return buffs
	},
}

// 金钟罩
var role_skill_1675 = &Skill{
	SkillId: 1675,
	ChildType: 3,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ABSORB_HURT, (buff_value(20000 + float64(skillTrnlv) * 100)), 2, 1675, 1, false))
			}
		}
		return buffs
	},
}

// 金刚庇护
var role_skill_1676 = &Skill{
	SkillId: 1676,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH && buddy1.sunderValue > 0)) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_SUNDER, (buff_value(100 + float64(skillTrnlv) * 2)), 0, 1676, 1, false))
			}
		}
		return buffs
	},
}

// 真灵守护
var role_skill_1677 = &Skill{
	SkillId: 1677,
	ChildType: 3,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_REDUCE_HURT, (buff_value(80)), 2, 1677, 1, false))
			}
		}
		buddies2 := []*Fighter{findLeastHealthBuddy(f)}
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_DEFEND, (buff_value(1000 + float64(skillTrnlv) * 10)), 2, 1677, 1, false))
			}
		}
		return buffs
	},
}

// 狮子吼
var role_skill_1678 = &Skill{
	SkillId: 1678,
	ChildType: 5,
	FixedValue: 1,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 1 + f.Cultivation * 0 + float64(skillTrnlv) * 0 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.4 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 2, 1678, 1, false))
		}
		return buffs
	},
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(2000 + float64(skillTrnlv) * 10)), 2, 1678, 1, false))
			}
		}
		return buffs
	},
}

// 金刚怒目
var role_skill_1679 = &Skill{
	SkillId: 1679,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTRACT_FIRE, (buff_value(1)), 2, 1679, 1, false))
		return buffs
	},
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(3000 + float64(skillTrnlv) * 30)), 2, 1679, 1, false))
			}
		}
		return buffs
	},
}

// 致命
var role_skill_1680 = &Skill{
	SkillId: 1680,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 3)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1680, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HURT_ADD, (buff_value(100)), 2, 1680, 1, false))
			}
		}
		buddies3 := f.getBuddies()
		for _, buddy3 := range buddies3 {
			if buddy3 != nil && ( false /*dead target*/ || (buddy3.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy3.addBuff(f, BUFF_HIT_LEVEL, (buff_value(1000 + float64(skillTrnlv) * 20)), 2, 1680, 1, false))
			}
		}
		return buffs
	},
}

// 寒冰凛冽
var role_skill_1681 = &Skill{
	SkillId: 1681,
	ChildType: 1,
	FixedValue: 8000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 寒冰甲
var role_skill_1682 = &Skill{
	SkillId: 1682,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DEFEND, (buff_value(2000 + float64(skillTrnlv) * 20)), 2, 1682, 1, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_HURT_ADD, (buff_value(60)), 2, 1682, 1, false))
			}
		}
		return buffs
	},
}

// 落星式
var role_skill_1683 = &Skill{
	SkillId: 1683,
	ChildType: 1,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.6 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 3, 1683, 1, false))
		}
		return buffs
	},
}

// 冬夜之拥
var role_skill_1684 = &Skill{
	SkillId: 1684,
	ChildType: 1,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 吸星锁
var role_skill_1685 = &Skill{
	SkillId: 1685,
	ChildType: 1,
	FixedValue: 6000,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.7 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1685, 1, false))
		}
		return buffs
	},
}

// 龙麟盾
var role_skill_1686 = &Skill{
	SkillId: 1686,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_SUNDER, (buff_value(200 + float64(skillTrnlv) * 3)), 0, 1686, 1, false))
		return buffs
	},
}

// 千钧破
var role_skill_1687 = &Skill{
	SkillId: 1687,
	ChildType: 1,
	FixedValue: 4000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 4000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 降妖伞
var role_skill_1688 = &Skill{
	SkillId: 1688,
	ChildType: 1,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 撼天狂掌
var role_skill_1689 = &Skill{
	SkillId: 1689,
	ChildType: 1,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findOneTargetWithOneRandom,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 韦驮拳
var role_skill_1690 = &Skill{
	SkillId: 1690,
	ChildType: 6,
	FixedValue: 8000,
	SunderAttack: 100,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
}

// 星沉地动
var role_skill_1691 = &Skill{
	SkillId: 1691,
	ChildType: 1,
	FixedValue: 6000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 6000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1691, 1, false))
		}
		return buffs
	},
}

// 寒月剑法
var role_skill_1692 = &Skill{
	SkillId: 1692,
	ChildType: 1,
	FixedValue: 8000,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1692, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 2, 1692, 1, false))
		return buffs
	},
}

// 万剑藏锋
var role_skill_1693 = &Skill{
	SkillId: 1693,
	ChildType: 1,
	FixedValue: 8000,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
	TriggerTargetBuff: true ,
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1693, 1, false))
		return buffs
	},
}

// 雷动九天
var role_skill_1694 = &Skill{
	SkillId: 1694,
	ChildType: 1,
	FixedValue: 8000,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + 8000 + f.Cultivation * 1.2 + float64(skillTrnlv) * 20 },
	IsTotemSkill: false, 
	IsRoleSkill: true, 
	_EventTrigger: func(f, t *Fighter) {
		if t.HasBuff(BUFF_ABSORB_HURT) {
			t.triggerEvent = FE_CRIT
		}
	},
}

// 星爆光离
var role_skill_1696 = &Skill{
	SkillId: 1696,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 10,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 召唤致命火蝎
var role_skill_1697 = &Skill{
	SkillId: 1697,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 1110,-5},CallInfo{ 1110,-5},CallInfo{ 1110,-5}},
}

// 破甲击
var role_skill_1698 = &Skill{
	SkillId: 1698,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(force)), 2, 1698, 2, false))
		return buffs
	},
}

// 咆哮冲撞
var role_skill_1699 = &Skill{
	SkillId: 1699,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 1, 1699, 2, false))
		return buffs
	},
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5  {
			buffs = append_buff(buffs, f.addBuff(f, BUFF_DIZZINESS, 1, 1, 1699, 2, false))
		}
		return buffs
	},
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(force * 0.2)), 2, 1699, 1, false))
			}
		}
		return buffs
	},
}

// 冲锋
var role_skill_1700 = &Skill{
	SkillId: 1700,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 针刺
var role_skill_1701 = &Skill{
	SkillId: 1701,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 魅惑
var role_skill_1702 = &Skill{
	SkillId: 1702,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.8 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 2, 1702, 1, false))
		}
		return buffs
	},
}

// 暴怒之击
var role_skill_1703 = &Skill{
	SkillId: 1703,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 0, 1703, 9, false))
			}
		}
		return buffs
	},
}

// 破盾利爪
var role_skill_1704 = &Skill{
	SkillId: 1704,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_CLEAR_ABSORB_HURT, 1, 0, 1704, 1, false))
		return buffs
	},
}

// 横向挥击
var role_skill_1705 = &Skill{
	SkillId: 1705,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 纵向挥击
var role_skill_1706 = &Skill{
	SkillId: 1706,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 狂毒之咬
var role_skill_1707 = &Skill{
	SkillId: 1707,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(force)), 3, 1707, 2, false))
		return buffs
	},
}

// 奇袭利爪
var role_skill_1708 = &Skill{
	SkillId: 1708,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 岩土冲击
var role_skill_1709 = &Skill{
	SkillId: 1709,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 憾地突袭
var role_skill_1710 = &Skill{
	SkillId: 1710,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 雷暴
var role_skill_1711 = &Skill{
	SkillId: 1711,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(4000)), 1, 1711, 1, false))
		return buffs
	},
}

// 雷电轰击
var role_skill_1712 = &Skill{
	SkillId: 1712,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1712, 1, false))
		}
		return buffs
	},
}

// 咆哮
var role_skill_1713 = &Skill{
	SkillId: 1713,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1713, 1, false))
		}
		return buffs
	},
}

// 血腥狂暴
var role_skill_1714 = &Skill{
	SkillId: 1714,
	ChildType: 5,
	FixedValue: 0,
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_ATTACK, (buff_value(1 + f.raw.Attack * 0.15)), 0, 1714, 9, true))
		return buffs
	},
}

// 吸魂
var role_skill_1715 = &Skill{
	SkillId: 1715,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_GHOST_POWER, -(buff_value(force)), 0, 1715, 1, false))
		return buffs
	},
}

// 毒牙撕咬
var role_skill_1716 = &Skill{
	SkillId: 1716,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + f.Attack * 0.3)), 3, 1716, 3, false))
		return buffs
	},
}

// 减攻混乱
var role_skill_1717 = &Skill{
	SkillId: 1717,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.4 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 2, 1717, 1, false))
		}
		return buffs
	},
}

// 死亡缠绕
var role_skill_1718 = &Skill{
	SkillId: 1718,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(1 + t.raw.HitLevel * 0.5)), 0, 1718, 1, false))
		return buffs
	},
}

// 诅咒
var role_skill_1719 = &Skill{
	SkillId: 1719,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_HIT_LEVEL, -(buff_value(1 + force)), 3, 1719, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1719, 1, false))
		return buffs
	},
}

// 尸毒
var role_skill_1720 = &Skill{
	SkillId: 1720,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 9, 1720, 9, false))
		return buffs
	},
}

// 致命诅咒
var role_skill_1721 = &Skill{
	SkillId: 1721,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_REDUCE_HURT, -(buff_value(10)), 9, 1721, 9, false))
		return buffs
	},
}

// 冰霜袭击
var role_skill_1722 = &Skill{
	SkillId: 1722,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1722, 2, false))
		return buffs
	},
}

// 冰霜奇袭
var role_skill_1723 = &Skill{
	SkillId: 1723,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findLeastHealthTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 召唤猛鬼僵尸
var role_skill_1724 = &Skill{
	SkillId: 1724,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 997,-5},CallInfo{ 997,-5}},
}

// 利爪挥击
var role_skill_1725 = &Skill{
	SkillId: 1725,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 妖风
var role_skill_1726 = &Skill{
	SkillId: 1726,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 3, 1726, 1, false))
			}
		}
		return buffs
	},
}

// 血魔之力
var role_skill_1727 = &Skill{
	SkillId: 1727,
	ChildType: 5,
	FixedValue: 0,
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 3, 1727, 1, false))
			}
		}
		return buffs
	},
}

// 梦蚀
var role_skill_1728 = &Skill{
	SkillId: 1728,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.6 - (t.Sleep + (t.SleepLevel * SLEEP_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_SLEEP, (buff_value(1)), 2, 1728, 1, false))
		}
		return buffs
	},
}

// 噩梦
var role_skill_1729 = &Skill{
	SkillId: 1729,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1729, 2, false))
		return buffs
	},
}

// 毒素
var role_skill_1730 = &Skill{
	SkillId: 1730,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 3, 1730, 1, false))
		return buffs
	},
}

// 血咒
var role_skill_1731 = &Skill{
	SkillId: 1731,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1 + force)), 3, 1731, 1, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1731, 1, false))
		return buffs
	},
}

// 混乱之触
var role_skill_1732 = &Skill{
	SkillId: 1732,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 2, 1732, 1, false))
		return buffs
	},
}

// 冰霜纵向
var role_skill_1733 = &Skill{
	SkillId: 1733,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1733, 3, false))
		return buffs
	},
}

// 伤害加深（全体）
var role_skill_1734 = &Skill{
	SkillId: 1734,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_REDUCE_HURT, -(buff_value(1 + force)), 2, 1734, 1, false))
		return buffs
	},
}

// 烈焰冲击
var role_skill_1735 = &Skill{
	SkillId: 1735,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 300,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1735, 1, false))
		}
		return buffs
	},
}

// 碾压
var role_skill_1736 = &Skill{
	SkillId: 1736,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.Dizziness + (t.DizzinessLevel * DIZZINESS_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DIZZINESS, 1, 2, 1736, 1, false))
		}
		return buffs
	},
}

// 绝望诅咒
var role_skill_1737 = &Skill{
	SkillId: 1737,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ATTACK, -(buff_value(1 + force)), 9, 1737, 0, false))
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 9, 1737, 0, false))
		return buffs
	},
}

// 封魔剑
var role_skill_1738 = &Skill{
	SkillId: 1738,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.3 - (t.DisableSkill + (t.DisableSkillLevel * DISABLE_SKILL_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_DISABLE_SKILL, 1, 2, 1738, 1, false))
		}
		return buffs
	},
}

// 爆怒之击
var role_skill_1739 = &Skill{
	SkillId: 1739,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToBuddy: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 2)
		buffCount := 0.0
		_ = buffCount
		buddies1 := f.getBuddies()
		for _, buddy1 := range buddies1 {
			if buddy1 != nil && ( false /*dead target*/ || (buddy1.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy1.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 9, 1739, 9, false))
			}
		}
		buddies2 := f.getBuddies()
		for _, buddy2 := range buddies2 {
			if buddy2 != nil && ( false /*dead target*/ || (buddy2.Health > MIN_HEALTH )) {
				buffs = append_buff(buffs, buddy2.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(300)), 9, 1739, 9, false))
			}
		}
		return buffs
	},
}

// 血蚀纵向
var role_skill_1740 = &Skill{
	SkillId: 1740,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1740, 3, false))
		return buffs
	},
}

// 影狼纵向
var role_skill_1741 = &Skill{
	SkillId: 1741,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1741, 2, false))
		return buffs
	},
}

// 召唤魔煞僵尸
var role_skill_1742 = &Skill{
	SkillId: 1742,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 1053,-5},CallInfo{ 1053,-5}},
}

// 腐蚀术
var role_skill_1743 = &Skill{
	SkillId: 1743,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1743, 5, false))
		return buffs
	},
}

// 狂乱之击
var role_skill_1744 = &Skill{
	SkillId: 1744,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5 - (t.Random + (t.RandomLevel * RANDOM_LEVEL_ARG / f.probLevel)) * 0.01   {
			buffs = append_buff(buffs, t.addBuff(f, BUFF_RANDOM, 1, 2, 1744, 1, false))
		}
		return buffs
	},
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		if rand.Float64() < 0.5  {
			buffs = append_buff(buffs, f.addBuff(f, BUFF_RANDOM, 1, 2, 1744, 1, false))
		}
		return buffs
	},
}

// 毒烈横排穿透
var role_skill_1745 = &Skill{
	SkillId: 1745,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.3)), 3, 1745, 1, false))
		return buffs
	},
}

// 暗影之毒
var role_skill_1746 = &Skill{
	SkillId: 1746,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_POISONING, (buff_value(1 + float64(hurt) * 0.2)), 3, 1746, 1, false))
		return buffs
	},
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HEALTH, (buff_value(1 + float64(hurt) * 0.2 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1746, 1, false))
		return buffs
	},
}

// 痛苦诅咒
var role_skill_1747 = &Skill{
	SkillId: 1747,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_ALL_RESIST, -(buff_value(1 + force)), 3, 1747, 1, false))
		return buffs
	},
}

// 八象皆杀
var role_skill_1748 = &Skill{
	SkillId: 1748,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1748, 3, false))
		return buffs
	},
}

// 吸血剑
var role_skill_1749 = &Skill{
	SkillId: 1749,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToSelf: func(f *Fighter, hurt int, force float64, skillTrnlv int16, ghost *GhostSkill) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, f.addBuff(f, BUFF_HEALTH, (buff_value(1 + float64(hurt) * 0.1 + ghost.CureAdd * 5)) * (1 + ghost.CureAddRate), 0, 1749, 1, false))
		return buffs
	},
}

// 吸魂剑
var role_skill_1750 = &Skill{
	SkillId: 1750,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_GHOST_POWER, -(buff_value(force)), 0, 1750, 1, false))
		return buffs
	},
}

// 致残
var role_skill_1751 = &Skill{
	SkillId: 1751,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
	_BuffToTarget: func(f, t *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		buffCount := 0.0
		_ = buffCount
		buffs = append_buff(buffs, t.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1751, 2, false))
		return buffs
	},
}

// 飓风
var role_skill_1752 = &Skill{
	SkillId: 1752,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 被动：反伤
var role_skill_1753 = &Skill{
	SkillId: 1753,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		targets0 := []*Fighter{f}
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_HEALTH, -(buff_value(1 + force)), 0, 1753, 1, false))
			}
		}
		return buffs
	},
}

// 亡语：炎爆
var role_skill_1754 = &Skill{
	SkillId: 1754,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 3)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_HEALTH, -(buff_value(1 + force)), 0, 1754, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force * 0.3)), 3, 1754, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets2 := t.side.Fighters
		for _, target2 := range targets2 {
			if target2 != nil && ( false /*对死亡目标有效*/ || (target2.Health > MIN_HEALTH && target2.sunderValue > 0)) {
			buffs = append_buff(buffs, target2.addBuff(f, BUFF_SUNDER, -(buff_value(300)), 0, 1754, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：冲击
var role_skill_1755 = &Skill{
	SkillId: 1755,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 3)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1755, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_DIZZINESS, 1, 2, 1755, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets2 := t.side.Fighters
		for _, target2 := range targets2 {
			if target2 != nil && ( false /*对死亡目标有效*/ || (target2.Health > MIN_HEALTH && target2.sunderValue > 0)) {
			buffs = append_buff(buffs, target2.addBuff(f, BUFF_SUNDER, -(buff_value(300)), 0, 1755, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：免伤护盾
var role_skill_1756 = &Skill{
	SkillId: 1756,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_REDUCE_HURT, (buff_value(90)), 2, 1756, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := f.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_HEALTH, (buff_value(1 + force)), 0, 1756, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：清除增益
var role_skill_1757 = &Skill{
	SkillId: 1757,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_CLEAN_GOOD, 1, 0, 1757, 1, false))
			}
		}
		}
		return buffs
	},
}

// 被动：清除增益
var role_skill_1758 = &Skill{
	SkillId: 1758,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_CLEAN_GOOD, 1, 0, 1758, 1, false))
			}
		}
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1758, 1, false))
			}
		}
		return buffs
	},
}

// 被动：化攻为守
var role_skill_1759 = &Skill{
	SkillId: 1759,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_REDUCE_HURT, (buff_value(1 + force)), 0, 1759, 1, false))
			}
		}
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_ATTACK, -(buff_value(1 + target1.raw.Attack * 0.9)), 0, 1759, 1, false))
			}
		}
		return buffs
	},
}

// 亡语：魂力
var role_skill_1760 = &Skill{
	SkillId: 1760,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_GHOST_POWER, (buff_value(force)), 0, 1760, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：爆裂
var role_skill_1761 = &Skill{
	SkillId: 1761,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 3)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1761, 0, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := f.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1761, 0, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets2 := t.side.Fighters
		for _, target2 := range targets2 {
			if target2 != nil && ( false /*对死亡目标有效*/ || (target2.Health > MIN_HEALTH && target2.sunderValue > 0)) {
			buffs = append_buff(buffs, target2.addBuff(f, BUFF_HEALTH, -(buff_value(1 + force * 3)), 0, 1761, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：强能
var role_skill_1762 = &Skill{
	SkillId: 1762,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 3, 1762, 9, false))
			}
		}
		}
		return buffs
	},
}

// 被动：毒伤
var role_skill_1763 = &Skill{
	SkillId: 1763,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		targets0 := []*Fighter{f}
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 3, 1763, 1, false))
			}
		}
		return buffs
	},
}

// 亡语：狂乱
var role_skill_1764 = &Skill{
	SkillId: 1764,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 3, 1764, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_RANDOM, 1, 2, 1764, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：爆炸
var role_skill_1765 = &Skill{
	SkillId: 1765,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_HEALTH, -(buff_value(1 + force)), 0, 1765, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：血爆
var role_skill_1766 = &Skill{
	SkillId: 1766,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_HEALTH, -(buff_value(1 + force)), 0, 1766, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1000)), 3, 1766, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：毒爆
var role_skill_1767 = &Skill{
	SkillId: 1767,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 3, 1767, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := f.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_POISONING, (buff_value(1 + force)), 3, 1767, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：清除减益
var role_skill_1768 = &Skill{
	SkillId: 1768,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_CLEAN_BAD, 1, 0, 1768, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：狂暴之血
var role_skill_1769 = &Skill{
	SkillId: 1769,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_CRITIAL_LEVEL, (buff_value(1 + force)), 4, 1769, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：强攻
var role_skill_1770 = &Skill{
	SkillId: 1770,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_ATTACK, (buff_value(1 + force)), 0, 1770, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：治疗敌方
var role_skill_1771 = &Skill{
	SkillId: 1771,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := f.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_HEALTH, (buff_value(1 + force)), 0, 1771, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：爆冲
var role_skill_1772 = &Skill{
	SkillId: 1772,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 2)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1772, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := t.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_DIZZINESS, 1, 2, 1772, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：爆乱
var role_skill_1773 = &Skill{
	SkillId: 1773,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 4)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_RANDOM, 1, 2, 1773, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets1 := f.side.Fighters
		for _, target1 := range targets1 {
			if target1 != nil && ( false /*对死亡目标有效*/ || (target1.Health > MIN_HEALTH && target1.sunderValue > 0)) {
			buffs = append_buff(buffs, target1.addBuff(f, BUFF_RANDOM, 1, 2, 1773, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets2 := t.side.Fighters
		for _, target2 := range targets2 {
			if target2 != nil && ( false /*对死亡目标有效*/ || (target2.Health > MIN_HEALTH && target2.sunderValue > 0)) {
			buffs = append_buff(buffs, target2.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1773, 1, false))
			}
		}
		}
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets3 := f.side.Fighters
		for _, target3 := range targets3 {
			if target3 != nil && ( false /*对死亡目标有效*/ || (target3.Health > MIN_HEALTH && target3.sunderValue > 0)) {
			buffs = append_buff(buffs, target3.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 2, 1773, 1, false))
			}
		}
		}
		return buffs
	},
}

// 亡语：风阵
var role_skill_1774 = &Skill{
	SkillId: 1774,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DODGE_LEVEL, (buff_value(1 + force)), 2, 1774, 1, false))
			}
		}
		}
		return buffs
	},
}

// 尸王横扫
var role_skill_1775 = &Skill{
	SkillId: 1775,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 魁拔突刺
var role_skill_1776 = &Skill{
	SkillId: 1776,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findLastRowTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 魔血爆纵向
var role_skill_1777 = &Skill{
	SkillId: 1777,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findColTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 死亡挥击
var role_skill_1778 = &Skill{
	SkillId: 1778,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 末日之拳
var role_skill_1779 = &Skill{
	SkillId: 1779,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 尸王落刃
var role_skill_1780 = &Skill{
	SkillId: 1780,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 9999,
	NotMiss: true,
	GhostAddRate: 1,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.6,
	_FindTargets: findAllTargets,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 亡语：蚀甲
var role_skill_1781 = &Skill{
	SkillId: 1781,
	ChildType: 1,
	FixedValue: 0,
	_PassiveBuff: func(f/*攻击者*/, t/*被动buff使用者*/ *Fighter, hurt int, force float64, skillTrnlv int16) []*Buff {
		buffs := make([]*Buff, 0, 1)
		if float64(t.Health) / float64(t.MaxHealth)  <= float64(0)/100 {
		targets0 := t.side.Fighters
		for _, target0 := range targets0 {
			if target0 != nil && ( false /*对死亡目标有效*/ || (target0.Health > MIN_HEALTH && target0.sunderValue > 0)) {
			buffs = append_buff(buffs, target0.addBuff(f, BUFF_DEFEND, -(buff_value(1 + force)), 0, 1781, 1, false))
			}
		}
		}
		return buffs
	},
}

// 爆炸蓝
var role_skill_1782 = &Skill{
	SkillId: 1782,
	ChildType: 1,
	FixedValue: 0,
}

// 爆炸红
var role_skill_1783 = &Skill{
	SkillId: 1783,
	ChildType: 1,
	FixedValue: 0,
}

// 爆炸绿
var role_skill_1784 = &Skill{
	SkillId: 1784,
	ChildType: 1,
	FixedValue: 0,
}

// 爆炸黄
var role_skill_1785 = &Skill{
	SkillId: 1785,
	ChildType: 1,
	FixedValue: 0,
}

// 召唤爆裂虫
var role_skill_1786 = &Skill{
	SkillId: 1786,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 1120,-5},CallInfo{ 1120,-5},CallInfo{ 1120,-5}},
}

// 召唤魔刃使徒
var role_skill_1787 = &Skill{
	SkillId: 1787,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 1066,9},CallInfo{ 1066,1}},
}

// 召唤马刀使徒
var role_skill_1788 = &Skill{
	SkillId: 1788,
	ChildType: 5,
	FixedValue: 0,
	CallEnemys:[]CallInfo{CallInfo{ 1067,3}},
}

// 闪击
var role_skill_1789 = &Skill{
	SkillId: 1789,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 5,
	AttackMode: SKILL_ATTACK_MODE_SINGLE,
	AttackRangeRatio: 1,
	_FindTargets: findOneTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 飓风冲击
var role_skill_1790 = &Skill{
	SkillId: 1790,
	ChildType: 1,
	FixedValue: 0,
	SunderAttack: 5,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

// 野蛮冲击
var role_skill_1791 = &Skill{
	SkillId: 1791,
	ChildType: 1,
	FixedValue: 0,
	GhostAddRate: 2,
	AttackMode: SKILL_ATTACK_MODE_AOE,
	AttackRangeRatio: 0.75,
	_FindTargets: findRowPenetrateTarget,
	_GetAttack: func(f *Fighter, force float64, skillTrnlv int16) float64 { return f.Attack + force },
}

