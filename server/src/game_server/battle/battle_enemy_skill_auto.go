package battle


import (
	"math/rand"
)

func createEnemySkill(f *Fighter) (int, /*force for enemy and training level for battle pet*/ int) {
	//玩家的灵宠
	if f.IsBattlePet && f.PlayerId > 0 {
		switch f.RoleId {
	case 91:
		return 1272, int(f.NaturalSkillLv)
	case 92:
		return 1134, int(f.NaturalSkillLv)
	case 93:
		return 1135, int(f.NaturalSkillLv)
	case 116:
		return 1267, int(f.NaturalSkillLv)
	case 117:
		return 1262, int(f.NaturalSkillLv)
	case 118:
		return 1257, int(f.NaturalSkillLv)
	case 119:
		return 1136, int(f.NaturalSkillLv)
	case 296:
		return 1137, int(f.NaturalSkillLv)
	case 297:
		return 1138, int(f.NaturalSkillLv)
	case 299:
		return 1252, int(f.NaturalSkillLv)
	case 771:
		return 1587, int(f.NaturalSkillLv)
	case 772:
		return 1588, int(f.NaturalSkillLv)
	case 773:
		return 1589, int(f.NaturalSkillLv)
	case 774:
		return 1590, int(f.NaturalSkillLv)

			default:
				return 0,0
		}
	}

	switch f.RoleId {
	case 1:
		return randSkillWithEnemy_1(f)
	case 2:
		return randSkillWithEnemy_2(f)
	case 3:
		return randSkillWithEnemy_3(f)
	case 4:
		return randSkillWithEnemy_4(f)
	case 5:
		return randSkillWithEnemy_5(f)
	case 8:
		return randSkillWithEnemy_8(f)
	case 10:
		return randSkillWithEnemy_10(f)
	case 11:
		return randSkillWithEnemy_11(f)
	case 12:
		return randSkillWithEnemy_12(f)
	case 14:
		return randSkillWithEnemy_14(f)
	case 15:
		return randSkillWithEnemy_15(f)
	case 16:
		return randSkillWithEnemy_16(f)
	case 17:
		return randSkillWithEnemy_17(f)
	case 18:
		return randSkillWithEnemy_18(f)
	case 19:
		return randSkillWithEnemy_19(f)
	case 21:
		return randSkillWithEnemy_21(f)
	case 22:
		return randSkillWithEnemy_22(f)
	case 23:
		return randSkillWithEnemy_23(f)
	case 25:
		return randSkillWithEnemy_25(f)
	case 26:
		return randSkillWithEnemy_26(f)
	case 37:
		return randSkillWithEnemy_37(f)
	case 38:
		return randSkillWithEnemy_38(f)
	case 43:
		return randSkillWithEnemy_43(f)
	case 44:
		return randSkillWithEnemy_44(f)
	case 45:
		return randSkillWithEnemy_45(f)
	case 48:
		return randSkillWithEnemy_48(f)
	case 49:
		return randSkillWithEnemy_49(f)
	case 50:
		return randSkillWithEnemy_50(f)
	case 51:
		return randSkillWithEnemy_51(f)
	case 52:
		return randSkillWithEnemy_52(f)
	case 53:
		return randSkillWithEnemy_53(f)
	case 54:
		return randSkillWithEnemy_54(f)
	case 55:
		return randSkillWithEnemy_55(f)
	case 56:
		return randSkillWithEnemy_56(f)
	case 57:
		return randSkillWithEnemy_57(f)
	case 58:
		return randSkillWithEnemy_58(f)
	case 59:
		return randSkillWithEnemy_59(f)
	case 60:
		return randSkillWithEnemy_60(f)
	case 61:
		return randSkillWithEnemy_61(f)
	case 62:
		return randSkillWithEnemy_62(f)
	case 63:
		return randSkillWithEnemy_63(f)
	case 64:
		return randSkillWithEnemy_64(f)
	case 65:
		return randSkillWithEnemy_65(f)
	case 66:
		return randSkillWithEnemy_66(f)
	case 67:
		return randSkillWithEnemy_67(f)
	case 68:
		return randSkillWithEnemy_68(f)
	case 69:
		return randSkillWithEnemy_69(f)
	case 70:
		return randSkillWithEnemy_70(f)
	case 71:
		return randSkillWithEnemy_71(f)
	case 72:
		return randSkillWithEnemy_72(f)
	case 73:
		return randSkillWithEnemy_73(f)
	case 74:
		return randSkillWithEnemy_74(f)
	case 75:
		return randSkillWithEnemy_75(f)
	case 76:
		return randSkillWithEnemy_76(f)
	case 77:
		return randSkillWithEnemy_77(f)
	case 78:
		return randSkillWithEnemy_78(f)
	case 79:
		return randSkillWithEnemy_79(f)
	case 80:
		return randSkillWithEnemy_80(f)
	case 81:
		return randSkillWithEnemy_81(f)
	case 82:
		return randSkillWithEnemy_82(f)
	case 83:
		return randSkillWithEnemy_83(f)
	case 84:
		return randSkillWithEnemy_84(f)
	case 85:
		return randSkillWithEnemy_85(f)
	case 86:
		return randSkillWithEnemy_86(f)
	case 87:
		return randSkillWithEnemy_87(f)
	case 91:
		return randSkillWithEnemy_91(f)
	case 92:
		return randSkillWithEnemy_92(f)
	case 93:
		return randSkillWithEnemy_93(f)
	case 94:
		return randSkillWithEnemy_94(f)
	case 96:
		return randSkillWithEnemy_96(f)
	case 97:
		return randSkillWithEnemy_97(f)
	case 98:
		return randSkillWithEnemy_98(f)
	case 100:
		return randSkillWithEnemy_100(f)
	case 101:
		return randSkillWithEnemy_101(f)
	case 102:
		return randSkillWithEnemy_102(f)
	case 103:
		return randSkillWithEnemy_103(f)
	case 104:
		return randSkillWithEnemy_104(f)
	case 105:
		return randSkillWithEnemy_105(f)
	case 106:
		return randSkillWithEnemy_106(f)
	case 107:
		return randSkillWithEnemy_107(f)
	case 110:
		return randSkillWithEnemy_110(f)
	case 111:
		return randSkillWithEnemy_111(f)
	case 112:
		return randSkillWithEnemy_112(f)
	case 113:
		return randSkillWithEnemy_113(f)
	case 114:
		return randSkillWithEnemy_114(f)
	case 115:
		return randSkillWithEnemy_115(f)
	case 116:
		return randSkillWithEnemy_116(f)
	case 117:
		return randSkillWithEnemy_117(f)
	case 118:
		return randSkillWithEnemy_118(f)
	case 119:
		return randSkillWithEnemy_119(f)
	case 120:
		return randSkillWithEnemy_120(f)
	case 121:
		return randSkillWithEnemy_121(f)
	case 122:
		return randSkillWithEnemy_122(f)
	case 124:
		return randSkillWithEnemy_124(f)
	case 125:
		return randSkillWithEnemy_125(f)
	case 126:
		return randSkillWithEnemy_126(f)
	case 127:
		return randSkillWithEnemy_127(f)
	case 129:
		return randSkillWithEnemy_129(f)
	case 130:
		return randSkillWithEnemy_130(f)
	case 131:
		return randSkillWithEnemy_131(f)
	case 132:
		return randSkillWithEnemy_132(f)
	case 133:
		return randSkillWithEnemy_133(f)
	case 134:
		return randSkillWithEnemy_134(f)
	case 135:
		return randSkillWithEnemy_135(f)
	case 136:
		return randSkillWithEnemy_136(f)
	case 137:
		return randSkillWithEnemy_137(f)
	case 138:
		return randSkillWithEnemy_138(f)
	case 139:
		return randSkillWithEnemy_139(f)
	case 140:
		return randSkillWithEnemy_140(f)
	case 141:
		return randSkillWithEnemy_141(f)
	case 142:
		return randSkillWithEnemy_142(f)
	case 143:
		return randSkillWithEnemy_143(f)
	case 144:
		return randSkillWithEnemy_144(f)
	case 145:
		return randSkillWithEnemy_145(f)
	case 146:
		return randSkillWithEnemy_146(f)
	case 147:
		return randSkillWithEnemy_147(f)
	case 148:
		return randSkillWithEnemy_148(f)
	case 149:
		return randSkillWithEnemy_149(f)
	case 150:
		return randSkillWithEnemy_150(f)
	case 151:
		return randSkillWithEnemy_151(f)
	case 152:
		return randSkillWithEnemy_152(f)
	case 153:
		return randSkillWithEnemy_153(f)
	case 154:
		return randSkillWithEnemy_154(f)
	case 155:
		return randSkillWithEnemy_155(f)
	case 156:
		return randSkillWithEnemy_156(f)
	case 157:
		return randSkillWithEnemy_157(f)
	case 158:
		return randSkillWithEnemy_158(f)
	case 159:
		return randSkillWithEnemy_159(f)
	case 160:
		return randSkillWithEnemy_160(f)
	case 161:
		return randSkillWithEnemy_161(f)
	case 162:
		return randSkillWithEnemy_162(f)
	case 163:
		return randSkillWithEnemy_163(f)
	case 164:
		return randSkillWithEnemy_164(f)
	case 165:
		return randSkillWithEnemy_165(f)
	case 166:
		return randSkillWithEnemy_166(f)
	case 167:
		return randSkillWithEnemy_167(f)
	case 168:
		return randSkillWithEnemy_168(f)
	case 169:
		return randSkillWithEnemy_169(f)
	case 170:
		return randSkillWithEnemy_170(f)
	case 171:
		return randSkillWithEnemy_171(f)
	case 172:
		return randSkillWithEnemy_172(f)
	case 173:
		return randSkillWithEnemy_173(f)
	case 174:
		return randSkillWithEnemy_174(f)
	case 175:
		return randSkillWithEnemy_175(f)
	case 176:
		return randSkillWithEnemy_176(f)
	case 177:
		return randSkillWithEnemy_177(f)
	case 178:
		return randSkillWithEnemy_178(f)
	case 179:
		return randSkillWithEnemy_179(f)
	case 180:
		return randSkillWithEnemy_180(f)
	case 181:
		return randSkillWithEnemy_181(f)
	case 182:
		return randSkillWithEnemy_182(f)
	case 183:
		return randSkillWithEnemy_183(f)
	case 184:
		return randSkillWithEnemy_184(f)
	case 185:
		return randSkillWithEnemy_185(f)
	case 186:
		return randSkillWithEnemy_186(f)
	case 187:
		return randSkillWithEnemy_187(f)
	case 188:
		return randSkillWithEnemy_188(f)
	case 189:
		return randSkillWithEnemy_189(f)
	case 191:
		return randSkillWithEnemy_191(f)
	case 192:
		return randSkillWithEnemy_192(f)
	case 194:
		return randSkillWithEnemy_194(f)
	case 195:
		return randSkillWithEnemy_195(f)
	case 196:
		return randSkillWithEnemy_196(f)
	case 197:
		return randSkillWithEnemy_197(f)
	case 198:
		return randSkillWithEnemy_198(f)
	case 204:
		return randSkillWithEnemy_204(f)
	case 205:
		return randSkillWithEnemy_205(f)
	case 206:
		return randSkillWithEnemy_206(f)
	case 208:
		return randSkillWithEnemy_208(f)
	case 210:
		return randSkillWithEnemy_210(f)
	case 214:
		return randSkillWithEnemy_214(f)
	case 215:
		return randSkillWithEnemy_215(f)
	case 216:
		return randSkillWithEnemy_216(f)
	case 217:
		return randSkillWithEnemy_217(f)
	case 218:
		return randSkillWithEnemy_218(f)
	case 219:
		return randSkillWithEnemy_219(f)
	case 220:
		return randSkillWithEnemy_220(f)
	case 221:
		return randSkillWithEnemy_221(f)
	case 222:
		return randSkillWithEnemy_222(f)
	case 223:
		return randSkillWithEnemy_223(f)
	case 224:
		return randSkillWithEnemy_224(f)
	case 225:
		return randSkillWithEnemy_225(f)
	case 226:
		return randSkillWithEnemy_226(f)
	case 227:
		return randSkillWithEnemy_227(f)
	case 228:
		return randSkillWithEnemy_228(f)
	case 229:
		return randSkillWithEnemy_229(f)
	case 230:
		return randSkillWithEnemy_230(f)
	case 231:
		return randSkillWithEnemy_231(f)
	case 232:
		return randSkillWithEnemy_232(f)
	case 233:
		return randSkillWithEnemy_233(f)
	case 234:
		return randSkillWithEnemy_234(f)
	case 235:
		return randSkillWithEnemy_235(f)
	case 236:
		return randSkillWithEnemy_236(f)
	case 237:
		return randSkillWithEnemy_237(f)
	case 238:
		return randSkillWithEnemy_238(f)
	case 239:
		return randSkillWithEnemy_239(f)
	case 240:
		return randSkillWithEnemy_240(f)
	case 241:
		return randSkillWithEnemy_241(f)
	case 242:
		return randSkillWithEnemy_242(f)
	case 243:
		return randSkillWithEnemy_243(f)
	case 244:
		return randSkillWithEnemy_244(f)
	case 245:
		return randSkillWithEnemy_245(f)
	case 246:
		return randSkillWithEnemy_246(f)
	case 247:
		return randSkillWithEnemy_247(f)
	case 248:
		return randSkillWithEnemy_248(f)
	case 249:
		return randSkillWithEnemy_249(f)
	case 250:
		return randSkillWithEnemy_250(f)
	case 251:
		return randSkillWithEnemy_251(f)
	case 252:
		return randSkillWithEnemy_252(f)
	case 253:
		return randSkillWithEnemy_253(f)
	case 254:
		return randSkillWithEnemy_254(f)
	case 255:
		return randSkillWithEnemy_255(f)
	case 256:
		return randSkillWithEnemy_256(f)
	case 257:
		return randSkillWithEnemy_257(f)
	case 258:
		return randSkillWithEnemy_258(f)
	case 259:
		return randSkillWithEnemy_259(f)
	case 260:
		return randSkillWithEnemy_260(f)
	case 261:
		return randSkillWithEnemy_261(f)
	case 262:
		return randSkillWithEnemy_262(f)
	case 263:
		return randSkillWithEnemy_263(f)
	case 264:
		return randSkillWithEnemy_264(f)
	case 265:
		return randSkillWithEnemy_265(f)
	case 266:
		return randSkillWithEnemy_266(f)
	case 267:
		return randSkillWithEnemy_267(f)
	case 268:
		return randSkillWithEnemy_268(f)
	case 269:
		return randSkillWithEnemy_269(f)
	case 270:
		return randSkillWithEnemy_270(f)
	case 272:
		return randSkillWithEnemy_272(f)
	case 273:
		return randSkillWithEnemy_273(f)
	case 274:
		return randSkillWithEnemy_274(f)
	case 275:
		return randSkillWithEnemy_275(f)
	case 276:
		return randSkillWithEnemy_276(f)
	case 281:
		return randSkillWithEnemy_281(f)
	case 282:
		return randSkillWithEnemy_282(f)
	case 289:
		return randSkillWithEnemy_289(f)
	case 290:
		return randSkillWithEnemy_290(f)
	case 291:
		return randSkillWithEnemy_291(f)
	case 292:
		return randSkillWithEnemy_292(f)
	case 296:
		return randSkillWithEnemy_296(f)
	case 297:
		return randSkillWithEnemy_297(f)
	case 299:
		return randSkillWithEnemy_299(f)
	case 308:
		return randSkillWithEnemy_308(f)
	case 310:
		return randSkillWithEnemy_310(f)
	case 311:
		return randSkillWithEnemy_311(f)
	case 312:
		return randSkillWithEnemy_312(f)
	case 313:
		return randSkillWithEnemy_313(f)
	case 314:
		return randSkillWithEnemy_314(f)
	case 315:
		return randSkillWithEnemy_315(f)
	case 316:
		return randSkillWithEnemy_316(f)
	case 317:
		return randSkillWithEnemy_317(f)
	case 318:
		return randSkillWithEnemy_318(f)
	case 319:
		return randSkillWithEnemy_319(f)
	case 320:
		return randSkillWithEnemy_320(f)
	case 321:
		return randSkillWithEnemy_321(f)
	case 322:
		return randSkillWithEnemy_322(f)
	case 323:
		return randSkillWithEnemy_323(f)
	case 324:
		return randSkillWithEnemy_324(f)
	case 325:
		return randSkillWithEnemy_325(f)
	case 326:
		return randSkillWithEnemy_326(f)
	case 327:
		return randSkillWithEnemy_327(f)
	case 328:
		return randSkillWithEnemy_328(f)
	case 329:
		return randSkillWithEnemy_329(f)
	case 330:
		return randSkillWithEnemy_330(f)
	case 331:
		return randSkillWithEnemy_331(f)
	case 332:
		return randSkillWithEnemy_332(f)
	case 333:
		return randSkillWithEnemy_333(f)
	case 334:
		return randSkillWithEnemy_334(f)
	case 335:
		return randSkillWithEnemy_335(f)
	case 336:
		return randSkillWithEnemy_336(f)
	case 337:
		return randSkillWithEnemy_337(f)
	case 338:
		return randSkillWithEnemy_338(f)
	case 339:
		return randSkillWithEnemy_339(f)
	case 340:
		return randSkillWithEnemy_340(f)
	case 341:
		return randSkillWithEnemy_341(f)
	case 342:
		return randSkillWithEnemy_342(f)
	case 343:
		return randSkillWithEnemy_343(f)
	case 344:
		return randSkillWithEnemy_344(f)
	case 345:
		return randSkillWithEnemy_345(f)
	case 346:
		return randSkillWithEnemy_346(f)
	case 347:
		return randSkillWithEnemy_347(f)
	case 348:
		return randSkillWithEnemy_348(f)
	case 349:
		return randSkillWithEnemy_349(f)
	case 350:
		return randSkillWithEnemy_350(f)
	case 351:
		return randSkillWithEnemy_351(f)
	case 352:
		return randSkillWithEnemy_352(f)
	case 353:
		return randSkillWithEnemy_353(f)
	case 354:
		return randSkillWithEnemy_354(f)
	case 355:
		return randSkillWithEnemy_355(f)
	case 356:
		return randSkillWithEnemy_356(f)
	case 357:
		return randSkillWithEnemy_357(f)
	case 358:
		return randSkillWithEnemy_358(f)
	case 359:
		return randSkillWithEnemy_359(f)
	case 360:
		return randSkillWithEnemy_360(f)
	case 361:
		return randSkillWithEnemy_361(f)
	case 362:
		return randSkillWithEnemy_362(f)
	case 363:
		return randSkillWithEnemy_363(f)
	case 364:
		return randSkillWithEnemy_364(f)
	case 365:
		return randSkillWithEnemy_365(f)
	case 366:
		return randSkillWithEnemy_366(f)
	case 367:
		return randSkillWithEnemy_367(f)
	case 368:
		return randSkillWithEnemy_368(f)
	case 369:
		return randSkillWithEnemy_369(f)
	case 370:
		return randSkillWithEnemy_370(f)
	case 371:
		return randSkillWithEnemy_371(f)
	case 372:
		return randSkillWithEnemy_372(f)
	case 373:
		return randSkillWithEnemy_373(f)
	case 380:
		return randSkillWithEnemy_380(f)
	case 387:
		return randSkillWithEnemy_387(f)
	case 394:
		return randSkillWithEnemy_394(f)
	case 395:
		return randSkillWithEnemy_395(f)
	case 396:
		return randSkillWithEnemy_396(f)
	case 397:
		return randSkillWithEnemy_397(f)
	case 398:
		return randSkillWithEnemy_398(f)
	case 399:
		return randSkillWithEnemy_399(f)
	case 400:
		return randSkillWithEnemy_400(f)
	case 401:
		return randSkillWithEnemy_401(f)
	case 402:
		return randSkillWithEnemy_402(f)
	case 403:
		return randSkillWithEnemy_403(f)
	case 412:
		return randSkillWithEnemy_412(f)
	case 419:
		return randSkillWithEnemy_419(f)
	case 426:
		return randSkillWithEnemy_426(f)
	case 433:
		return randSkillWithEnemy_433(f)
	case 437:
		return randSkillWithEnemy_437(f)
	case 438:
		return randSkillWithEnemy_438(f)
	case 439:
		return randSkillWithEnemy_439(f)
	case 440:
		return randSkillWithEnemy_440(f)
	case 441:
		return randSkillWithEnemy_441(f)
	case 442:
		return randSkillWithEnemy_442(f)
	case 443:
		return randSkillWithEnemy_443(f)
	case 444:
		return randSkillWithEnemy_444(f)
	case 445:
		return randSkillWithEnemy_445(f)
	case 446:
		return randSkillWithEnemy_446(f)
	case 447:
		return randSkillWithEnemy_447(f)
	case 492:
		return randSkillWithEnemy_492(f)
	case 493:
		return randSkillWithEnemy_493(f)
	case 494:
		return randSkillWithEnemy_494(f)
	case 495:
		return randSkillWithEnemy_495(f)
	case 496:
		return randSkillWithEnemy_496(f)
	case 497:
		return randSkillWithEnemy_497(f)
	case 498:
		return randSkillWithEnemy_498(f)
	case 499:
		return randSkillWithEnemy_499(f)
	case 500:
		return randSkillWithEnemy_500(f)
	case 501:
		return randSkillWithEnemy_501(f)
	case 502:
		return randSkillWithEnemy_502(f)
	case 503:
		return randSkillWithEnemy_503(f)
	case 504:
		return randSkillWithEnemy_504(f)
	case 505:
		return randSkillWithEnemy_505(f)
	case 506:
		return randSkillWithEnemy_506(f)
	case 507:
		return randSkillWithEnemy_507(f)
	case 508:
		return randSkillWithEnemy_508(f)
	case 509:
		return randSkillWithEnemy_509(f)
	case 510:
		return randSkillWithEnemy_510(f)
	case 511:
		return randSkillWithEnemy_511(f)
	case 512:
		return randSkillWithEnemy_512(f)
	case 513:
		return randSkillWithEnemy_513(f)
	case 514:
		return randSkillWithEnemy_514(f)
	case 515:
		return randSkillWithEnemy_515(f)
	case 516:
		return randSkillWithEnemy_516(f)
	case 517:
		return randSkillWithEnemy_517(f)
	case 518:
		return randSkillWithEnemy_518(f)
	case 519:
		return randSkillWithEnemy_519(f)
	case 520:
		return randSkillWithEnemy_520(f)
	case 521:
		return randSkillWithEnemy_521(f)
	case 522:
		return randSkillWithEnemy_522(f)
	case 523:
		return randSkillWithEnemy_523(f)
	case 524:
		return randSkillWithEnemy_524(f)
	case 525:
		return randSkillWithEnemy_525(f)
	case 526:
		return randSkillWithEnemy_526(f)
	case 527:
		return randSkillWithEnemy_527(f)
	case 528:
		return randSkillWithEnemy_528(f)
	case 529:
		return randSkillWithEnemy_529(f)
	case 530:
		return randSkillWithEnemy_530(f)
	case 531:
		return randSkillWithEnemy_531(f)
	case 532:
		return randSkillWithEnemy_532(f)
	case 533:
		return randSkillWithEnemy_533(f)
	case 534:
		return randSkillWithEnemy_534(f)
	case 535:
		return randSkillWithEnemy_535(f)
	case 536:
		return randSkillWithEnemy_536(f)
	case 537:
		return randSkillWithEnemy_537(f)
	case 538:
		return randSkillWithEnemy_538(f)
	case 539:
		return randSkillWithEnemy_539(f)
	case 540:
		return randSkillWithEnemy_540(f)
	case 541:
		return randSkillWithEnemy_541(f)
	case 542:
		return randSkillWithEnemy_542(f)
	case 543:
		return randSkillWithEnemy_543(f)
	case 544:
		return randSkillWithEnemy_544(f)
	case 545:
		return randSkillWithEnemy_545(f)
	case 546:
		return randSkillWithEnemy_546(f)
	case 565:
		return randSkillWithEnemy_565(f)
	case 566:
		return randSkillWithEnemy_566(f)
	case 567:
		return randSkillWithEnemy_567(f)
	case 568:
		return randSkillWithEnemy_568(f)
	case 569:
		return randSkillWithEnemy_569(f)
	case 570:
		return randSkillWithEnemy_570(f)
	case 571:
		return randSkillWithEnemy_571(f)
	case 572:
		return randSkillWithEnemy_572(f)
	case 573:
		return randSkillWithEnemy_573(f)
	case 574:
		return randSkillWithEnemy_574(f)
	case 575:
		return randSkillWithEnemy_575(f)
	case 576:
		return randSkillWithEnemy_576(f)
	case 577:
		return randSkillWithEnemy_577(f)
	case 578:
		return randSkillWithEnemy_578(f)
	case 579:
		return randSkillWithEnemy_579(f)
	case 580:
		return randSkillWithEnemy_580(f)
	case 581:
		return randSkillWithEnemy_581(f)
	case 582:
		return randSkillWithEnemy_582(f)
	case 583:
		return randSkillWithEnemy_583(f)
	case 584:
		return randSkillWithEnemy_584(f)
	case 585:
		return randSkillWithEnemy_585(f)
	case 586:
		return randSkillWithEnemy_586(f)
	case 587:
		return randSkillWithEnemy_587(f)
	case 588:
		return randSkillWithEnemy_588(f)
	case 589:
		return randSkillWithEnemy_589(f)
	case 590:
		return randSkillWithEnemy_590(f)
	case 591:
		return randSkillWithEnemy_591(f)
	case 592:
		return randSkillWithEnemy_592(f)
	case 593:
		return randSkillWithEnemy_593(f)
	case 594:
		return randSkillWithEnemy_594(f)
	case 595:
		return randSkillWithEnemy_595(f)
	case 596:
		return randSkillWithEnemy_596(f)
	case 597:
		return randSkillWithEnemy_597(f)
	case 598:
		return randSkillWithEnemy_598(f)
	case 599:
		return randSkillWithEnemy_599(f)
	case 600:
		return randSkillWithEnemy_600(f)
	case 601:
		return randSkillWithEnemy_601(f)
	case 602:
		return randSkillWithEnemy_602(f)
	case 603:
		return randSkillWithEnemy_603(f)
	case 604:
		return randSkillWithEnemy_604(f)
	case 605:
		return randSkillWithEnemy_605(f)
	case 606:
		return randSkillWithEnemy_606(f)
	case 607:
		return randSkillWithEnemy_607(f)
	case 608:
		return randSkillWithEnemy_608(f)
	case 609:
		return randSkillWithEnemy_609(f)
	case 610:
		return randSkillWithEnemy_610(f)
	case 611:
		return randSkillWithEnemy_611(f)
	case 612:
		return randSkillWithEnemy_612(f)
	case 613:
		return randSkillWithEnemy_613(f)
	case 614:
		return randSkillWithEnemy_614(f)
	case 616:
		return randSkillWithEnemy_616(f)
	case 617:
		return randSkillWithEnemy_617(f)
	case 618:
		return randSkillWithEnemy_618(f)
	case 619:
		return randSkillWithEnemy_619(f)
	case 620:
		return randSkillWithEnemy_620(f)
	case 623:
		return randSkillWithEnemy_623(f)
	case 626:
		return randSkillWithEnemy_626(f)
	case 629:
		return randSkillWithEnemy_629(f)
	case 632:
		return randSkillWithEnemy_632(f)
	case 635:
		return randSkillWithEnemy_635(f)
	case 636:
		return randSkillWithEnemy_636(f)
	case 637:
		return randSkillWithEnemy_637(f)
	case 638:
		return randSkillWithEnemy_638(f)
	case 639:
		return randSkillWithEnemy_639(f)
	case 640:
		return randSkillWithEnemy_640(f)
	case 641:
		return randSkillWithEnemy_641(f)
	case 642:
		return randSkillWithEnemy_642(f)
	case 643:
		return randSkillWithEnemy_643(f)
	case 644:
		return randSkillWithEnemy_644(f)
	case 645:
		return randSkillWithEnemy_645(f)
	case 646:
		return randSkillWithEnemy_646(f)
	case 647:
		return randSkillWithEnemy_647(f)
	case 648:
		return randSkillWithEnemy_648(f)
	case 669:
		return randSkillWithEnemy_669(f)
	case 670:
		return randSkillWithEnemy_670(f)
	case 671:
		return randSkillWithEnemy_671(f)
	case 675:
		return randSkillWithEnemy_675(f)
	case 676:
		return randSkillWithEnemy_676(f)
	case 677:
		return randSkillWithEnemy_677(f)
	case 679:
		return randSkillWithEnemy_679(f)
	case 680:
		return randSkillWithEnemy_680(f)
	case 681:
		return randSkillWithEnemy_681(f)
	case 682:
		return randSkillWithEnemy_682(f)
	case 684:
		return randSkillWithEnemy_684(f)
	case 685:
		return randSkillWithEnemy_685(f)
	case 686:
		return randSkillWithEnemy_686(f)
	case 687:
		return randSkillWithEnemy_687(f)
	case 688:
		return randSkillWithEnemy_688(f)
	case 689:
		return randSkillWithEnemy_689(f)
	case 690:
		return randSkillWithEnemy_690(f)
	case 693:
		return randSkillWithEnemy_693(f)
	case 695:
		return randSkillWithEnemy_695(f)
	case 696:
		return randSkillWithEnemy_696(f)
	case 697:
		return randSkillWithEnemy_697(f)
	case 699:
		return randSkillWithEnemy_699(f)
	case 700:
		return randSkillWithEnemy_700(f)
	case 701:
		return randSkillWithEnemy_701(f)
	case 702:
		return randSkillWithEnemy_702(f)
	case 703:
		return randSkillWithEnemy_703(f)
	case 704:
		return randSkillWithEnemy_704(f)
	case 706:
		return randSkillWithEnemy_706(f)
	case 713:
		return randSkillWithEnemy_713(f)
	case 714:
		return randSkillWithEnemy_714(f)
	case 715:
		return randSkillWithEnemy_715(f)
	case 716:
		return randSkillWithEnemy_716(f)
	case 717:
		return randSkillWithEnemy_717(f)
	case 718:
		return randSkillWithEnemy_718(f)
	case 719:
		return randSkillWithEnemy_719(f)
	case 720:
		return randSkillWithEnemy_720(f)
	case 721:
		return randSkillWithEnemy_721(f)
	case 722:
		return randSkillWithEnemy_722(f)
	case 723:
		return randSkillWithEnemy_723(f)
	case 724:
		return randSkillWithEnemy_724(f)
	case 725:
		return randSkillWithEnemy_725(f)
	case 726:
		return randSkillWithEnemy_726(f)
	case 727:
		return randSkillWithEnemy_727(f)
	case 728:
		return randSkillWithEnemy_728(f)
	case 729:
		return randSkillWithEnemy_729(f)
	case 730:
		return randSkillWithEnemy_730(f)
	case 731:
		return randSkillWithEnemy_731(f)
	case 732:
		return randSkillWithEnemy_732(f)
	case 733:
		return randSkillWithEnemy_733(f)
	case 734:
		return randSkillWithEnemy_734(f)
	case 735:
		return randSkillWithEnemy_735(f)
	case 736:
		return randSkillWithEnemy_736(f)
	case 737:
		return randSkillWithEnemy_737(f)
	case 738:
		return randSkillWithEnemy_738(f)
	case 739:
		return randSkillWithEnemy_739(f)
	case 740:
		return randSkillWithEnemy_740(f)
	case 741:
		return randSkillWithEnemy_741(f)
	case 742:
		return randSkillWithEnemy_742(f)
	case 743:
		return randSkillWithEnemy_743(f)
	case 744:
		return randSkillWithEnemy_744(f)
	case 745:
		return randSkillWithEnemy_745(f)
	case 746:
		return randSkillWithEnemy_746(f)
	case 747:
		return randSkillWithEnemy_747(f)
	case 748:
		return randSkillWithEnemy_748(f)
	case 749:
		return randSkillWithEnemy_749(f)
	case 750:
		return randSkillWithEnemy_750(f)
	case 751:
		return randSkillWithEnemy_751(f)
	case 752:
		return randSkillWithEnemy_752(f)
	case 753:
		return randSkillWithEnemy_753(f)
	case 754:
		return randSkillWithEnemy_754(f)
	case 755:
		return randSkillWithEnemy_755(f)
	case 758:
		return randSkillWithEnemy_758(f)
	case 761:
		return randSkillWithEnemy_761(f)
	case 766:
		return randSkillWithEnemy_766(f)
	case 767:
		return randSkillWithEnemy_767(f)
	case 768:
		return randSkillWithEnemy_768(f)
	case 769:
		return randSkillWithEnemy_769(f)
	case 770:
		return randSkillWithEnemy_770(f)
	case 771:
		return randSkillWithEnemy_771(f)
	case 772:
		return randSkillWithEnemy_772(f)
	case 773:
		return randSkillWithEnemy_773(f)
	case 774:
		return randSkillWithEnemy_774(f)
	case 775:
		return randSkillWithEnemy_775(f)
	case 776:
		return randSkillWithEnemy_776(f)
	case 777:
		return randSkillWithEnemy_777(f)
	case 778:
		return randSkillWithEnemy_778(f)
	case 779:
		return randSkillWithEnemy_779(f)
	case 780:
		return randSkillWithEnemy_780(f)
	case 795:
		return randSkillWithEnemy_795(f)
	case 796:
		return randSkillWithEnemy_796(f)
	case 797:
		return randSkillWithEnemy_797(f)
	case 798:
		return randSkillWithEnemy_798(f)
	case 799:
		return randSkillWithEnemy_799(f)
	case 800:
		return randSkillWithEnemy_800(f)
	case 801:
		return randSkillWithEnemy_801(f)
	case 802:
		return randSkillWithEnemy_802(f)
	case 803:
		return randSkillWithEnemy_803(f)
	case 804:
		return randSkillWithEnemy_804(f)
	case 805:
		return randSkillWithEnemy_805(f)
	case 806:
		return randSkillWithEnemy_806(f)
	case 807:
		return randSkillWithEnemy_807(f)
	case 808:
		return randSkillWithEnemy_808(f)
	case 809:
		return randSkillWithEnemy_809(f)
	case 810:
		return randSkillWithEnemy_810(f)
	case 811:
		return randSkillWithEnemy_811(f)
	case 812:
		return randSkillWithEnemy_812(f)
	case 813:
		return randSkillWithEnemy_813(f)
	case 814:
		return randSkillWithEnemy_814(f)
	case 815:
		return randSkillWithEnemy_815(f)
	case 816:
		return randSkillWithEnemy_816(f)
	case 817:
		return randSkillWithEnemy_817(f)
	case 818:
		return randSkillWithEnemy_818(f)
	case 819:
		return randSkillWithEnemy_819(f)
	case 820:
		return randSkillWithEnemy_820(f)
	case 821:
		return randSkillWithEnemy_821(f)
	case 822:
		return randSkillWithEnemy_822(f)
	case 823:
		return randSkillWithEnemy_823(f)
	case 824:
		return randSkillWithEnemy_824(f)
	case 825:
		return randSkillWithEnemy_825(f)
	case 826:
		return randSkillWithEnemy_826(f)
	case 827:
		return randSkillWithEnemy_827(f)
	case 828:
		return randSkillWithEnemy_828(f)
	case 829:
		return randSkillWithEnemy_829(f)
	case 830:
		return randSkillWithEnemy_830(f)
	case 831:
		return randSkillWithEnemy_831(f)
	case 832:
		return randSkillWithEnemy_832(f)
	case 833:
		return randSkillWithEnemy_833(f)
	case 836:
		return randSkillWithEnemy_836(f)
	case 839:
		return randSkillWithEnemy_839(f)
	case 842:
		return randSkillWithEnemy_842(f)
	case 843:
		return randSkillWithEnemy_843(f)
	case 844:
		return randSkillWithEnemy_844(f)
	case 845:
		return randSkillWithEnemy_845(f)
	case 846:
		return randSkillWithEnemy_846(f)
	case 847:
		return randSkillWithEnemy_847(f)
	case 848:
		return randSkillWithEnemy_848(f)
	case 849:
		return randSkillWithEnemy_849(f)
	case 850:
		return randSkillWithEnemy_850(f)
	case 851:
		return randSkillWithEnemy_851(f)
	case 859:
		return randSkillWithEnemy_859(f)
	case 860:
		return randSkillWithEnemy_860(f)
	case 861:
		return randSkillWithEnemy_861(f)
	case 862:
		return randSkillWithEnemy_862(f)
	case 865:
		return randSkillWithEnemy_865(f)
	case 868:
		return randSkillWithEnemy_868(f)
	case 947:
		return randSkillWithEnemy_947(f)
	case 948:
		return randSkillWithEnemy_948(f)
	case 949:
		return randSkillWithEnemy_949(f)
	case 950:
		return randSkillWithEnemy_950(f)
	case 951:
		return randSkillWithEnemy_951(f)
	case 952:
		return randSkillWithEnemy_952(f)
	case 953:
		return randSkillWithEnemy_953(f)
	case 954:
		return randSkillWithEnemy_954(f)
	case 955:
		return randSkillWithEnemy_955(f)
	case 956:
		return randSkillWithEnemy_956(f)
	case 957:
		return randSkillWithEnemy_957(f)
	case 958:
		return randSkillWithEnemy_958(f)
	case 959:
		return randSkillWithEnemy_959(f)
	case 960:
		return randSkillWithEnemy_960(f)
	case 961:
		return randSkillWithEnemy_961(f)
	case 962:
		return randSkillWithEnemy_962(f)
	case 963:
		return randSkillWithEnemy_963(f)
	case 964:
		return randSkillWithEnemy_964(f)
	case 965:
		return randSkillWithEnemy_965(f)
	case 966:
		return randSkillWithEnemy_966(f)
	case 967:
		return randSkillWithEnemy_967(f)
	case 968:
		return randSkillWithEnemy_968(f)
	case 969:
		return randSkillWithEnemy_969(f)
	case 970:
		return randSkillWithEnemy_970(f)
	case 971:
		return randSkillWithEnemy_971(f)
	case 972:
		return randSkillWithEnemy_972(f)
	case 973:
		return randSkillWithEnemy_973(f)
	case 974:
		return randSkillWithEnemy_974(f)
	case 975:
		return randSkillWithEnemy_975(f)
	case 976:
		return randSkillWithEnemy_976(f)
	case 979:
		return randSkillWithEnemy_979(f)
	case 980:
		return randSkillWithEnemy_980(f)
	case 982:
		return randSkillWithEnemy_982(f)
	case 983:
		return randSkillWithEnemy_983(f)
	case 984:
		return randSkillWithEnemy_984(f)
	case 987:
		return randSkillWithEnemy_987(f)
	case 988:
		return randSkillWithEnemy_988(f)
	case 989:
		return randSkillWithEnemy_989(f)
	case 990:
		return randSkillWithEnemy_990(f)
	case 991:
		return randSkillWithEnemy_991(f)
	case 992:
		return randSkillWithEnemy_992(f)
	case 993:
		return randSkillWithEnemy_993(f)
	case 994:
		return randSkillWithEnemy_994(f)
	case 995:
		return randSkillWithEnemy_995(f)
	case 996:
		return randSkillWithEnemy_996(f)
	case 997:
		return randSkillWithEnemy_997(f)
	case 998:
		return randSkillWithEnemy_998(f)
	case 999:
		return randSkillWithEnemy_999(f)
	case 1000:
		return randSkillWithEnemy_1000(f)
	case 1002:
		return randSkillWithEnemy_1002(f)
	case 1005:
		return randSkillWithEnemy_1005(f)
	case 1007:
		return randSkillWithEnemy_1007(f)
	case 1008:
		return randSkillWithEnemy_1008(f)
	case 1009:
		return randSkillWithEnemy_1009(f)
	case 1010:
		return randSkillWithEnemy_1010(f)
	case 1011:
		return randSkillWithEnemy_1011(f)
	case 1012:
		return randSkillWithEnemy_1012(f)
	case 1014:
		return randSkillWithEnemy_1014(f)
	case 1015:
		return randSkillWithEnemy_1015(f)
	case 1016:
		return randSkillWithEnemy_1016(f)
	case 1017:
		return randSkillWithEnemy_1017(f)
	case 1018:
		return randSkillWithEnemy_1018(f)
	case 1019:
		return randSkillWithEnemy_1019(f)
	case 1020:
		return randSkillWithEnemy_1020(f)
	case 1021:
		return randSkillWithEnemy_1021(f)
	case 1022:
		return randSkillWithEnemy_1022(f)
	case 1023:
		return randSkillWithEnemy_1023(f)
	case 1025:
		return randSkillWithEnemy_1025(f)
	case 1026:
		return randSkillWithEnemy_1026(f)
	case 1028:
		return randSkillWithEnemy_1028(f)
	case 1029:
		return randSkillWithEnemy_1029(f)
	case 1030:
		return randSkillWithEnemy_1030(f)
	case 1031:
		return randSkillWithEnemy_1031(f)
	case 1032:
		return randSkillWithEnemy_1032(f)
	case 1034:
		return randSkillWithEnemy_1034(f)
	case 1035:
		return randSkillWithEnemy_1035(f)
	case 1036:
		return randSkillWithEnemy_1036(f)
	case 1037:
		return randSkillWithEnemy_1037(f)
	case 1038:
		return randSkillWithEnemy_1038(f)
	case 1039:
		return randSkillWithEnemy_1039(f)
	case 1040:
		return randSkillWithEnemy_1040(f)
	case 1041:
		return randSkillWithEnemy_1041(f)
	case 1042:
		return randSkillWithEnemy_1042(f)
	case 1045:
		return randSkillWithEnemy_1045(f)
	case 1046:
		return randSkillWithEnemy_1046(f)
	case 1048:
		return randSkillWithEnemy_1048(f)
	case 1049:
		return randSkillWithEnemy_1049(f)
	case 1050:
		return randSkillWithEnemy_1050(f)
	case 1051:
		return randSkillWithEnemy_1051(f)
	case 1052:
		return randSkillWithEnemy_1052(f)
	case 1054:
		return randSkillWithEnemy_1054(f)
	case 1055:
		return randSkillWithEnemy_1055(f)
	case 1056:
		return randSkillWithEnemy_1056(f)
	case 1057:
		return randSkillWithEnemy_1057(f)
	case 1058:
		return randSkillWithEnemy_1058(f)
	case 1059:
		return randSkillWithEnemy_1059(f)
	case 1060:
		return randSkillWithEnemy_1060(f)
	case 1061:
		return randSkillWithEnemy_1061(f)
	case 1062:
		return randSkillWithEnemy_1062(f)
	case 1063:
		return randSkillWithEnemy_1063(f)
	case 1064:
		return randSkillWithEnemy_1064(f)
	case 1065:
		return randSkillWithEnemy_1065(f)
	case 1066:
		return randSkillWithEnemy_1066(f)
	case 1067:
		return randSkillWithEnemy_1067(f)
	case 1069:
		return randSkillWithEnemy_1069(f)
	case 1070:
		return randSkillWithEnemy_1070(f)
	case 1071:
		return randSkillWithEnemy_1071(f)
	case 1072:
		return randSkillWithEnemy_1072(f)
	case 1073:
		return randSkillWithEnemy_1073(f)
	case 1074:
		return randSkillWithEnemy_1074(f)
	case 1075:
		return randSkillWithEnemy_1075(f)
	case 1076:
		return randSkillWithEnemy_1076(f)
	case 1077:
		return randSkillWithEnemy_1077(f)
	case 1082:
		return randSkillWithEnemy_1082(f)
	case 1083:
		return randSkillWithEnemy_1083(f)
	case 1084:
		return randSkillWithEnemy_1084(f)
	case 1086:
		return randSkillWithEnemy_1086(f)
	case 1087:
		return randSkillWithEnemy_1087(f)
	case 1088:
		return randSkillWithEnemy_1088(f)
	case 1089:
		return randSkillWithEnemy_1089(f)
	case 1090:
		return randSkillWithEnemy_1090(f)
	case 1091:
		return randSkillWithEnemy_1091(f)
	case 1093:
		return randSkillWithEnemy_1093(f)
	case 1095:
		return randSkillWithEnemy_1095(f)
	case 1099:
		return randSkillWithEnemy_1099(f)
	case 1100:
		return randSkillWithEnemy_1100(f)
	case 1101:
		return randSkillWithEnemy_1101(f)
	case 1102:
		return randSkillWithEnemy_1102(f)
	case 1103:
		return randSkillWithEnemy_1103(f)
	case 1104:
		return randSkillWithEnemy_1104(f)
	case 1105:
		return randSkillWithEnemy_1105(f)
	case 1106:
		return randSkillWithEnemy_1106(f)
	case 1109:
		return randSkillWithEnemy_1109(f)
	case 1110:
		return randSkillWithEnemy_1110(f)
	case 1111:
		return randSkillWithEnemy_1111(f)
	case 1112:
		return randSkillWithEnemy_1112(f)
	case 1113:
		return randSkillWithEnemy_1113(f)
	case 1114:
		return randSkillWithEnemy_1114(f)
	case 1115:
		return randSkillWithEnemy_1115(f)
	case 1116:
		return randSkillWithEnemy_1116(f)
	case 1117:
		return randSkillWithEnemy_1117(f)
	case 1118:
		return randSkillWithEnemy_1118(f)
	case 1119:
		return randSkillWithEnemy_1119(f)
	case 1121:
		return randSkillWithEnemy_1121(f)
	case 1122:
		return randSkillWithEnemy_1122(f)
	case 1123:
		return randSkillWithEnemy_1123(f)
	case 1124:
		return randSkillWithEnemy_1124(f)
	case 1125:
		return randSkillWithEnemy_1125(f)
	case 1126:
		return randSkillWithEnemy_1126(f)
	case 1127:
		return randSkillWithEnemy_1127(f)
	case 1128:
		return randSkillWithEnemy_1128(f)
	case 1129:
		return randSkillWithEnemy_1129(f)
	case 1130:
		return randSkillWithEnemy_1130(f)
	case 1131:
		return randSkillWithEnemy_1131(f)
	case 1132:
		return randSkillWithEnemy_1132(f)
	case 1133:
		return randSkillWithEnemy_1133(f)
	case 1134:
		return randSkillWithEnemy_1134(f)
	case 1135:
		return randSkillWithEnemy_1135(f)
	case 1136:
		return randSkillWithEnemy_1136(f)
	case 1137:
		return randSkillWithEnemy_1137(f)
	case 1138:
		return randSkillWithEnemy_1138(f)
	case 1140:
		return randSkillWithEnemy_1140(f)
	case 1141:
		return randSkillWithEnemy_1141(f)
	case 1142:
		return randSkillWithEnemy_1142(f)
	case 1144:
		return randSkillWithEnemy_1144(f)
	case 1146:
		return randSkillWithEnemy_1146(f)
	case 1150:
		return randSkillWithEnemy_1150(f)
	case 1161:
		return randSkillWithEnemy_1161(f)
	case 1162:
		return randSkillWithEnemy_1162(f)
	case 1163:
		return randSkillWithEnemy_1163(f)
	case 1164:
		return randSkillWithEnemy_1164(f)
	case 1166:
		return randSkillWithEnemy_1166(f)
	case 1167:
		return randSkillWithEnemy_1167(f)
	case 1168:
		return randSkillWithEnemy_1168(f)
	case 1169:
		return randSkillWithEnemy_1169(f)
	case 1170:
		return randSkillWithEnemy_1170(f)
	case 1172:
		return randSkillWithEnemy_1172(f)
	case 1173:
		return randSkillWithEnemy_1173(f)
	case 1174:
		return randSkillWithEnemy_1174(f)
	case 1175:
		return randSkillWithEnemy_1175(f)
	case 1176:
		return randSkillWithEnemy_1176(f)
		
	}
	return 0, 0
}


func randSkillWithEnemy_1(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 480
	} else if randNum > 30 && randNum <= 80 {
		skillId = 43
		skillForce = 480
	}	


	return
}

func randSkillWithEnemy_2(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 31
		skillForce = 1320
	} else if randNum > 40 && randNum <= 80 {
		skillId = 29
		skillForce = 1320
	}	


	return
}

func randSkillWithEnemy_3(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 1290
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1035
		skillForce = 1290
	} else if randNum > 60 && randNum <= 80 {
		skillId = 41
		skillForce = 1290
	}	


	return
}

func randSkillWithEnemy_4(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 1260
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1038
		skillForce = 1260
	}	


	return
}

func randSkillWithEnemy_5(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 1560
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 1560
	} else if randNum > 40 && randNum <= 60 {
		skillId = 24
		skillForce = 1560
	} else if randNum > 60 && randNum <= 80 {
		skillId = 25
		skillForce = 1560
	}	


	return
}

func randSkillWithEnemy_8(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 24
		skillForce = 1590
	} else if randNum > 40 && randNum <= 80 {
		skillId = 25
		skillForce = 1590
	}	


	return
}

func randSkillWithEnemy_10(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 1950
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 1950
	} else if randNum > 40 && randNum <= 60 {
		skillId = 24
		skillForce = 1950
	} else if randNum > 60 && randNum <= 80 {
		skillId = 25
		skillForce = 1950
	}	


	return
}

func randSkillWithEnemy_11(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 110
	} else if randNum > 30 && randNum <= 80 {
		skillId = 43
		skillForce = 110
	}	


	return
}

func randSkillWithEnemy_12(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_14(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 100
	} else if randNum > 30 && randNum <= 60 {
		skillId = 9
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_15(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_16(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1315
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_17(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_18(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1315
		skillForce = 100
	} else if randNum > 20 && randNum <= 35 {
		skillId = 15
		skillForce = 100
	} else if randNum > 35 && randNum <= 50 {
		skillId = 27
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_19(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 15
		skillForce = 100
	} else if randNum > 20 && randNum <= 35 {
		skillId = 28
		skillForce = 100
	} else if randNum > 35 && randNum <= 50 {
		skillId = 43
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_21(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_22(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_23(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 220
	} else if randNum > 30 && randNum <= 60 {
		skillId = 42
		skillForce = 220
	} else if randNum > 60 && randNum <= 90 {
		skillId = 38
		skillForce = 220
	}	


	return
}

func randSkillWithEnemy_25(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 1230
	} else if randNum > 20 && randNum <= 40 {
		skillId = 25
		skillForce = 1230
	} else if randNum > 40 && randNum <= 60 {
		skillId = 21
		skillForce = 1230
	} else if randNum > 60 && randNum <= 80 {
		skillId = 24
		skillForce = 1230
	}	


	return
}

func randSkillWithEnemy_26(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 13
		skillForce = 1500
	} else if randNum > 40 && randNum <= 80 {
		skillId = 28
		skillForce = 1500
	}	


	return
}

func randSkillWithEnemy_37(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_38(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 24
		skillForce = 1830
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1039
		skillForce = 1830
	}	


	return
}

func randSkillWithEnemy_43(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_44(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1315
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_45(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 50
	}	


	return
}

func randSkillWithEnemy_48(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_49(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 10
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_50(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 11
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_51(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 15
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_52(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 19
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_53(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 23
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_54(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 27
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_55(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 31
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_56(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 35
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_57(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 36
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_58(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 41
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_59(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 42
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_60(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 43
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_61(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 46
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_62(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 109
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_63(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 37
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_64(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 38
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_65(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 39
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_66(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 44
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_67(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 108
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_68(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 40
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_69(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 45
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_70(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 12
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_71(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 16
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_72(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 20
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_73(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 24
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_74(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 28
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_75(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 32
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_76(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 13
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_77(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 17
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_78(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 21
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_79(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 25
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_80(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 29
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_81(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 33
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_82(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 14
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_83(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 18
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_84(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 22
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_85(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 26
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_86(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 30
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_87(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 34
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_91(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1228
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_92(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1232
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_93(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1240
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_94(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 34
		skillForce = 0
	}	


	return
}

func randSkillWithEnemy_96(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 45
		skillForce = 1680
	} else if randNum > 50 && randNum <= 80 {
		skillId = 37
		skillForce = 1680
	}	


	return
}

func randSkillWithEnemy_97(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1048
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_98(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 19
		skillForce = 0
	}	


	return
}

func randSkillWithEnemy_100(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 600
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 600
	} else if randNum > 60 && randNum <= 90 {
		skillId = 42
		skillForce = 600
	}	


	return
}

func randSkillWithEnemy_101(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 10 {
		skillId = 1038
		skillForce = 200
	} else if randNum > 10 && randNum <= 20 {
		skillId = 1039
		skillForce = 200
	} else if randNum > 20 && randNum <= 50 {
		skillId = 35
		skillForce = 200
	}	


	return
}

func randSkillWithEnemy_102(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1054
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_103(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 1350
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 1350
	} else if randNum > 60 && randNum <= 90 {
		skillId = 42
		skillForce = 1350
	}	


	return
}

func randSkillWithEnemy_104(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 24
		skillForce = 200
	} else if randNum > 30 && randNum <= 50 {
		skillId = 25
		skillForce = 200
	} else if randNum > 50 && randNum <= 80 {
		skillId = 20
		skillForce = 200
	}	


	return
}

func randSkillWithEnemy_105(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 12
		skillForce = 480
	} else if randNum > 30 && randNum <= 60 {
		skillId = 13
		skillForce = 480
	}	


	return
}

func randSkillWithEnemy_106(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 28
		skillForce = 640
	} else if randNum > 30 && randNum <= 60 {
		skillId = 29
		skillForce = 640
	}	


	return
}

func randSkillWithEnemy_107(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 1200
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 1200
	} else if randNum > 60 && randNum <= 90 {
		skillId = 42
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_110(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 7000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_111(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 19
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_112(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_113(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 45
		skillForce = 1650
	} else if randNum > 50 && randNum <= 80 {
		skillId = 37
		skillForce = 1650
	}	


	return
}

func randSkillWithEnemy_114(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 4080
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 4080
	} else if randNum > 40 && randNum <= 60 {
		skillId = 24
		skillForce = 4080
	} else if randNum > 60 && randNum <= 80 {
		skillId = 25
		skillForce = 4080
	}	


	return
}

func randSkillWithEnemy_115(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 19
		skillForce = 30
	} else if randNum > 20 && randNum <= 40 {
		skillId = 23
		skillForce = 30
	}	


	return
}

func randSkillWithEnemy_116(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1268
		skillForce = 2600
	}	


	return
}

func randSkillWithEnemy_117(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1263
		skillForce = 600
	}	


	return
}

func randSkillWithEnemy_118(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1258
		skillForce = 600
	}	


	return
}

func randSkillWithEnemy_119(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1244
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_120(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_121(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_122(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_124(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_125(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_126(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_127(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_129(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_130(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_131(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_132(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_133(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 250
	}	


	return
}

func randSkillWithEnemy_134(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 250
	}	


	return
}

func randSkillWithEnemy_135(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 270
	}	


	return
}

func randSkillWithEnemy_136(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 270
	}	


	return
}

func randSkillWithEnemy_137(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_138(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_139(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1052
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_140(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_141(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_142(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_143(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_144(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_145(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_146(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_147(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_148(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_149(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_150(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_151(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 420
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_152(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 108
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_153(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_154(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_155(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 130
	}	


	return
}

func randSkillWithEnemy_156(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_157(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_158(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_159(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 150
	}	


	return
}

func randSkillWithEnemy_160(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_161(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 28
		skillForce = 180
	} else if randNum > 25 && randNum <= 50 {
		skillId = 29
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_162(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 34
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_163(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 32
		skillForce = 210
	} else if randNum > 25 && randNum <= 50 {
		skillId = 33
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_164(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 250
	}	


	return
}

func randSkillWithEnemy_165(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 29
		skillForce = 250
	}	


	return
}

func randSkillWithEnemy_166(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 270
	}	


	return
}

func randSkillWithEnemy_167(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 17
		skillForce = 270
	}	


	return
}

func randSkillWithEnemy_168(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_169(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 38
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_170(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1053
		skillForce = 420
	}	


	return
}

func randSkillWithEnemy_171(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_172(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 45
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_173(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_174(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 465
	}	


	return
}

func randSkillWithEnemy_175(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 24
		skillForce = 510
	} else if randNum > 25 && randNum <= 50 {
		skillId = 25
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_176(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 12
		skillForce = 510
	} else if randNum > 25 && randNum <= 50 {
		skillId = 13
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_177(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 510
	}	


	return
}

func randSkillWithEnemy_178(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 33
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_179(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 28
		skillForce = 555
	} else if randNum > 25 && randNum <= 50 {
		skillId = 29
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_180(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 20
		skillForce = 555
	} else if randNum > 25 && randNum <= 50 {
		skillId = 21
		skillForce = 555
	}	


	return
}

func randSkillWithEnemy_181(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 36
		skillForce = 450
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1038
		skillForce = 450
	}	


	return
}

func randSkillWithEnemy_182(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1048
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_183(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 630
	} else if randNum > 50 && randNum <= 60 {
		skillId = 43
		skillForce = 630
	} else if randNum > 60 && randNum <= 80 {
		skillId = 29
		skillForce = 630
	}	


	return
}

func randSkillWithEnemy_184(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 19
		skillForce = 810
	} else if randNum > 25 && randNum <= 50 {
		skillId = 15
		skillForce = 810
	} else if randNum > 50 && randNum <= 65 {
		skillId = 17
		skillForce = 810
	} else if randNum > 65 && randNum <= 80 {
		skillId = 20
		skillForce = 810
	}	


	return
}

func randSkillWithEnemy_185(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 15
		skillForce = 870
	} else if randNum > 20 && randNum <= 40 {
		skillId = 28
		skillForce = 870
	} else if randNum > 40 && randNum <= 55 {
		skillId = 44
		skillForce = 870
	} else if randNum > 55 && randNum <= 70 {
		skillId = 43
		skillForce = 870
	}	


	return
}

func randSkillWithEnemy_186(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1053
		skillForce = 960
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1052
		skillForce = 960
	}	


	return
}

func randSkillWithEnemy_187(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 45
		skillForce = 1020
	} else if randNum > 50 && randNum <= 80 {
		skillId = 37
		skillForce = 1020
	}	


	return
}

func randSkillWithEnemy_188(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1038
		skillForce = 1170
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1039
		skillForce = 1170
	}	


	return
}

func randSkillWithEnemy_189(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_191(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_192(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_194(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_195(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_196(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_197(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 180
	}	


	return
}

func randSkillWithEnemy_198(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 32
		skillForce = 210
	} else if randNum > 25 && randNum <= 50 {
		skillId = 33
		skillForce = 210
	}	


	return
}

func randSkillWithEnemy_204(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 35
		skillForce = 3500
	} else if randNum > 20 && randNum <= 50 {
		skillId = 1431
		skillForce = 3500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 40
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_205(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 21
		skillForce = 800
	} else if randNum > 20 && randNum <= 40 {
		skillId = 20
		skillForce = 800
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1070
		skillForce = 800
	} else if randNum > 60 && randNum <= 80 {
		skillId = 22
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_206(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 36
		skillForce = 3000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 44
		skillForce = 3000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 43
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_208(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 36
		skillForce = 400
	} else if randNum > 20 && randNum <= 40 {
		skillId = 44
		skillForce = 400
	} else if randNum > 40 && randNum <= 60 {
		skillId = 43
		skillForce = 400
	}	


	return
}

func randSkillWithEnemy_210(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 39
		skillForce = 7000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 38
		skillForce = 7000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 42
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_214(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 1200
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_215(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 36
		skillForce = 1080
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1038
		skillForce = 1080
	}	


	return
}

func randSkillWithEnemy_216(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 1240
	} else if randNum > 50 && randNum <= 60 {
		skillId = 43
		skillForce = 1240
	} else if randNum > 60 && randNum <= 80 {
		skillId = 29
		skillForce = 1240
	}	


	return
}

func randSkillWithEnemy_217(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 15
		skillForce = 1520
	} else if randNum > 20 && randNum <= 40 {
		skillId = 28
		skillForce = 1520
	} else if randNum > 40 && randNum <= 55 {
		skillId = 44
		skillForce = 1520
	} else if randNum > 55 && randNum <= 70 {
		skillId = 43
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_218(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 45
		skillForce = 1800
	} else if randNum > 50 && randNum <= 80 {
		skillId = 37
		skillForce = 1800
	}	


	return
}

func randSkillWithEnemy_219(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 24
		skillForce = 1920
	} else if randNum > 20 && randNum <= 40 {
		skillId = 25
		skillForce = 1920
	} else if randNum > 40 && randNum <= 60 {
		skillId = 20
		skillForce = 1920
	} else if randNum > 60 && randNum <= 80 {
		skillId = 21
		skillForce = 1920
	}	


	return
}

func randSkillWithEnemy_220(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 43
		skillForce = 320
	}	


	return
}

func randSkillWithEnemy_221(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1504
		skillForce = 380
	} else if randNum > 20 && randNum <= 50 {
		skillId = 31
		skillForce = 380
	}	


	return
}

func randSkillWithEnemy_222(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 43
		skillForce = 500
	} else if randNum > 25 && randNum <= 50 {
		skillId = 15
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_223(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 800
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_224(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 880
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 880
	}	


	return
}

func randSkillWithEnemy_225(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1039
		skillForce = 1120
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1698
		skillForce = 1120
	}	


	return
}

func randSkillWithEnemy_226(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 44
		skillForce = 1350
	} else if randNum > 40 && randNum <= 70 {
		skillId = 43
		skillForce = 1350
	}	


	return
}

func randSkillWithEnemy_227(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 19
		skillForce = 1600
	} else if randNum > 25 && randNum <= 50 {
		skillId = 15
		skillForce = 1600
	} else if randNum > 50 && randNum <= 75 {
		skillId = 17
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_228(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1070
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_229(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_230(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_231(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_232(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_233(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 13
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_234(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 800
	}	


	return
}

func randSkillWithEnemy_235(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_236(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_237(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 43
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_238(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_239(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_240(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 860
	}	


	return
}

func randSkillWithEnemy_241(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 860
	}	


	return
}

func randSkillWithEnemy_242(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_243(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_244(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 28
		skillForce = 820
	}	


	return
}

func randSkillWithEnemy_245(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 860
	}	


	return
}

func randSkillWithEnemy_246(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 860
	}	


	return
}

func randSkillWithEnemy_247(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 920
	}	


	return
}

func randSkillWithEnemy_248(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 920
	}	


	return
}

func randSkillWithEnemy_249(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 980
	}	


	return
}

func randSkillWithEnemy_250(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 980
	}	


	return
}

func randSkillWithEnemy_251(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 28
		skillForce = 920
	}	


	return
}

func randSkillWithEnemy_252(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 13
		skillForce = 920
	}	


	return
}

func randSkillWithEnemy_253(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 28
		skillForce = 980
	}	


	return
}

func randSkillWithEnemy_254(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 13
		skillForce = 980
	}	


	return
}

func randSkillWithEnemy_255(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_256(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_257(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 1090
	}	


	return
}

func randSkillWithEnemy_258(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 38
		skillForce = 1090
	}	


	return
}

func randSkillWithEnemy_259(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 24
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_260(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 13
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_261(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 1090
	}	


	return
}

func randSkillWithEnemy_262(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 38
		skillForce = 1090
	}	


	return
}

func randSkillWithEnemy_263(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 1140
	}	


	return
}

func randSkillWithEnemy_264(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1140
	}	


	return
}

func randSkillWithEnemy_265(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_266(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_267(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 28
		skillForce = 1140
	}	


	return
}

func randSkillWithEnemy_268(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 1140
	}	


	return
}

func randSkillWithEnemy_269(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_270(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 39
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_272(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 2000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1035
		skillForce = 2000
	} else if randNum > 60 && randNum <= 80 {
		skillId = 41
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_273(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 31
		skillForce = 2320
	} else if randNum > 40 && randNum <= 80 {
		skillId = 29
		skillForce = 2320
	}	


	return
}

func randSkillWithEnemy_274(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 13
		skillForce = 2480
	} else if randNum > 40 && randNum <= 80 {
		skillId = 28
		skillForce = 2480
	}	


	return
}

func randSkillWithEnemy_275(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 24
		skillForce = 2840
	} else if randNum > 40 && randNum <= 80 {
		skillId = 25
		skillForce = 2840
	}	


	return
}

func randSkillWithEnemy_276(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 17
		skillForce = 3200
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1039
		skillForce = 3200
	}	


	return
}

func randSkillWithEnemy_281(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 36
		skillForce = 1300
	} else if randNum > 20 && randNum <= 40 {
		skillId = 1038
		skillForce = 1300
	} else if randNum > 40 && randNum <= 60 {
		skillId = 41
		skillForce = 1300
	} else if randNum > 60 && randNum <= 80 {
		skillId = 30
		skillForce = 1300
	}	


	return
}

func randSkillWithEnemy_282(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 43
		skillForce = 3900
	} else if randNum > 20 && randNum <= 50 {
		skillId = 1236
		skillForce = 3900
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1218
		skillForce = 3900
	}	


	return
}

func randSkillWithEnemy_289(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 39
		skillForce = 2000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 38
		skillForce = 2000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 42
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_290(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 38
		skillForce = 5000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 39
		skillForce = 5000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 42
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_291(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 2000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_292(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 37
		skillForce = 5000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 35
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_296(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1248
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_297(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1236
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_299(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1000 {
		skillId = 1253
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_308(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 36
		skillForce = 1000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 44
		skillForce = 1000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 43
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_310(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 36
		skillForce = 2000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 44
		skillForce = 2000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 43
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_311(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1147
		skillForce = 10000
	}	


	return
}

func randSkillWithEnemy_312(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_313(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_314(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_315(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_316(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 1360
	}	


	return
}

func randSkillWithEnemy_317(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_318(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_319(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 1280
	}	


	return
}

func randSkillWithEnemy_320(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 1360
	}	


	return
}

func randSkillWithEnemy_321(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 14
		skillForce = 1360
	}	


	return
}

func randSkillWithEnemy_322(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 42
		skillForce = 1360
	}	


	return
}

func randSkillWithEnemy_323(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 1440
	}	


	return
}

func randSkillWithEnemy_324(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 1440
	}	


	return
}

func randSkillWithEnemy_325(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_326(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_327(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_328(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 1440
	}	


	return
}

func randSkillWithEnemy_329(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 1440
	}	


	return
}

func randSkillWithEnemy_330(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_331(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_332(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_333(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 1980
	}	


	return
}

func randSkillWithEnemy_334(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1980
	}	


	return
}

func randSkillWithEnemy_335(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 32
		skillForce = 2160
	}	


	return
}

func randSkillWithEnemy_336(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 33
		skillForce = 2160
	}	


	return
}

func randSkillWithEnemy_337(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 21
		skillForce = 1980
	}	


	return
}

func randSkillWithEnemy_338(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 24
		skillForce = 1980
	}	


	return
}

func randSkillWithEnemy_339(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 32
		skillForce = 2160
	}	


	return
}

func randSkillWithEnemy_340(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 33
		skillForce = 2160
	}	


	return
}

func randSkillWithEnemy_341(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_342(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_343(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 1740
	}	


	return
}

func randSkillWithEnemy_344(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 1740
	}	


	return
}

func randSkillWithEnemy_345(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_346(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_347(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 43
		skillForce = 1740
	}	


	return
}

func randSkillWithEnemy_348(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 1740
	}	


	return
}

func randSkillWithEnemy_349(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 13
		skillForce = 1740
	}	


	return
}

func randSkillWithEnemy_350(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_351(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_352(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_353(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_354(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_355(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_356(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_357(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 33
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_358(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_359(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 24
		skillForce = 2460
	}	


	return
}

func randSkillWithEnemy_360(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 9
		skillForce = 2100
	} else if randNum > 40 && randNum <= 80 {
		skillId = 19
		skillForce = 2100
	}	


	return
}

func randSkillWithEnemy_361(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 42
		skillForce = 2190
	} else if randNum > 40 && randNum <= 70 {
		skillId = 41
		skillForce = 2190
	}	


	return
}

func randSkillWithEnemy_362(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1035
		skillForce = 2370
	} else if randNum > 30 && randNum <= 50 {
		skillId = 27
		skillForce = 2370
	} else if randNum > 50 && randNum <= 80 {
		skillId = 37
		skillForce = 2370
	}	


	return
}

func randSkillWithEnemy_363(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1039
		skillForce = 2490
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 2490
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1038
		skillForce = 2490
	}	


	return
}

func randSkillWithEnemy_364(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 3270
	} else if randNum > 40 && randNum <= 60 {
		skillId = 38
		skillForce = 3270
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1039
		skillForce = 3270
	}	


	return
}

func randSkillWithEnemy_365(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 27
		skillForce = 3600
	} else if randNum > 40 && randNum <= 60 {
		skillId = 32
		skillForce = 3600
	} else if randNum > 60 && randNum <= 80 {
		skillId = 33
		skillForce = 3600
	}	


	return
}

func randSkillWithEnemy_366(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 20
		skillForce = 2790
	} else if randNum > 40 && randNum <= 80 {
		skillId = 21
		skillForce = 2790
	}	


	return
}

func randSkillWithEnemy_367(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 11
		skillForce = 3000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 42
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_368(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1039
		skillForce = 3870
	} else if randNum > 40 && randNum <= 80 {
		skillId = 37
		skillForce = 3870
	}	


	return
}

func randSkillWithEnemy_369(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 42
		skillForce = 3640
	} else if randNum > 40 && randNum <= 60 {
		skillId = 41
		skillForce = 3640
	}	


	return
}

func randSkillWithEnemy_370(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1039
		skillForce = 4400
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1315
		skillForce = 4400
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1484
		skillForce = 4400
	}	


	return
}

func randSkillWithEnemy_371(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 11
		skillForce = 5280
	} else if randNum > 40 && randNum <= 80 {
		skillId = 42
		skillForce = 5280
	}	


	return
}

func randSkillWithEnemy_372(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 4500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 34
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_373(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 6640
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 6640
	} else if randNum > 40 && randNum <= 60 {
		skillId = 24
		skillForce = 6640
	} else if randNum > 60 && randNum <= 80 {
		skillId = 25
		skillForce = 6640
	}	


	return
}

func randSkillWithEnemy_380(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 19
		skillForce = 1200
	} else if randNum > 25 && randNum <= 50 {
		skillId = 15
		skillForce = 1200
	} else if randNum > 50 && randNum <= 65 {
		skillId = 17
		skillForce = 1200
	} else if randNum > 65 && randNum <= 80 {
		skillId = 20
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_387(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 2400
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1038
		skillForce = 2400
	}	


	return
}

func randSkillWithEnemy_394(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 19
		skillForce = 3500
	} else if randNum > 40 && randNum <= 80 {
		skillId = 9
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_395(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 10 {
		skillId = 45
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_396(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_397(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_398(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_399(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_400(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_401(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 1
	}	


	return
}

func randSkillWithEnemy_402(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 1
	}	


	return
}

func randSkillWithEnemy_403(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 9
		skillForce = 1
	}	


	return
}

func randSkillWithEnemy_412(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 24
		skillForce = 6000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 25
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_419(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 8000
	} else if randNum > 30 && randNum <= 60 {
		skillId = 38
		skillForce = 8000
	} else if randNum > 60 && randNum <= 90 {
		skillId = 42
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_426(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 43
		skillForce = 10000
	} else if randNum > 50 && randNum <= 80 {
		skillId = 41
		skillForce = 10000
	}	


	return
}

func randSkillWithEnemy_433(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 42
		skillForce = 15000
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1035
		skillForce = 15000
	}	


	return
}

func randSkillWithEnemy_437(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1070
		skillForce = 2000
	} else if randNum > 30 && randNum <= 70 {
		skillId = 1512
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_438(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 45
		skillForce = 2000
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1747
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_439(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 2300
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1038
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_440(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1507
		skillForce = 2450
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1505
		skillForce = 2450
	} else if randNum > 60 && randNum <= 100 {
		skillId = 1506
		skillForce = 2450
	}	


	return
}

func randSkillWithEnemy_441(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 31
		skillForce = 2710
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1504
		skillForce = 2710
	}	


	return
}

func randSkillWithEnemy_442(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1712
		skillForce = 3000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 25
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_443(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1719
		skillForce = 3310
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1721
		skillForce = 3310
	}	


	return
}

func randSkillWithEnemy_444(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 108
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_445(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1047
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_446(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1047
		skillForce = 10000
	}	


	return
}

func randSkillWithEnemy_447(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1047
		skillForce = 15000
	}	


	return
}

func randSkillWithEnemy_492(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 2680
	}	


	return
}

func randSkillWithEnemy_493(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2680
	}	


	return
}

func randSkillWithEnemy_494(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_495(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_496(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 2680
	}	


	return
}

func randSkillWithEnemy_497(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2680
	}	


	return
}

func randSkillWithEnemy_498(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_499(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_500(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1039
		skillForce = 4470
	} else if randNum > 40 && randNum <= 80 {
		skillId = 38
		skillForce = 4470
	}	


	return
}

func randSkillWithEnemy_501(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1039
		skillForce = 4530
	} else if randNum > 40 && randNum <= 80 {
		skillId = 38
		skillForce = 4530
	}	


	return
}

func randSkillWithEnemy_502(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 2780
	}	


	return
}

func randSkillWithEnemy_503(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2780
	}	


	return
}

func randSkillWithEnemy_504(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1035
		skillForce = 2820
	}	


	return
}

func randSkillWithEnemy_505(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 2820
	}	


	return
}

func randSkillWithEnemy_506(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 2780
	}	


	return
}

func randSkillWithEnemy_507(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 37
		skillForce = 2780
	}	


	return
}

func randSkillWithEnemy_508(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1035
		skillForce = 2820
	}	


	return
}

func randSkillWithEnemy_509(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1419
		skillForce = 2820
	}	


	return
}

func randSkillWithEnemy_510(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 9
		skillForce = 4620
	} else if randNum > 40 && randNum <= 80 {
		skillId = 10
		skillForce = 4620
	}	


	return
}

func randSkillWithEnemy_511(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 42
		skillForce = 4680
	} else if randNum > 40 && randNum <= 80 {
		skillId = 41
		skillForce = 4680
	}	


	return
}

func randSkillWithEnemy_512(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 2880
	}	


	return
}

func randSkillWithEnemy_513(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 2880
	}	


	return
}

func randSkillWithEnemy_514(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 2920
	}	


	return
}

func randSkillWithEnemy_515(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1035
		skillForce = 2920
	}	


	return
}

func randSkillWithEnemy_516(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 2880
	}	


	return
}

func randSkillWithEnemy_517(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 24
		skillForce = 2880
	}	


	return
}

func randSkillWithEnemy_518(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 2920
	}	


	return
}

func randSkillWithEnemy_519(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1035
		skillForce = 2920
	}	


	return
}

func randSkillWithEnemy_520(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 39
		skillForce = 4770
	} else if randNum > 40 && randNum <= 80 {
		skillId = 38
		skillForce = 4770
	}	


	return
}

func randSkillWithEnemy_521(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1038
		skillForce = 4830
	} else if randNum > 40 && randNum <= 80 {
		skillId = 39
		skillForce = 4830
	}	


	return
}

func randSkillWithEnemy_522(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2980
	}	


	return
}

func randSkillWithEnemy_523(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 2980
	}	


	return
}

func randSkillWithEnemy_524(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 3040
	}	


	return
}

func randSkillWithEnemy_525(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 3040
	}	


	return
}

func randSkillWithEnemy_526(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2980
	}	


	return
}

func randSkillWithEnemy_527(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 2980
	}	


	return
}

func randSkillWithEnemy_528(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 29
		skillForce = 3040
	}	


	return
}

func randSkillWithEnemy_529(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 3040
	}	


	return
}

func randSkillWithEnemy_530(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 20
		skillForce = 4890
	} else if randNum > 40 && randNum <= 80 {
		skillId = 21
		skillForce = 4890
	}	


	return
}

func randSkillWithEnemy_531(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 41
		skillForce = 4950
	} else if randNum > 40 && randNum <= 80 {
		skillId = 24
		skillForce = 4950
	}	


	return
}

func randSkillWithEnemy_532(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_533(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_534(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 3220
	}	


	return
}

func randSkillWithEnemy_535(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 3220
	}	


	return
}

func randSkillWithEnemy_536(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 32
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_537(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_538(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 3220
	}	


	return
}

func randSkillWithEnemy_539(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 17
		skillForce = 3220
	}	


	return
}

func randSkillWithEnemy_540(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 45
		skillForce = 5250
	} else if randNum > 40 && randNum <= 80 {
		skillId = 39
		skillForce = 5250
	}	


	return
}

func randSkillWithEnemy_541(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 16
		skillForce = 5400
	} else if randNum > 40 && randNum <= 80 {
		skillId = 17
		skillForce = 5400
	}	


	return
}

func randSkillWithEnemy_542(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1039
		skillForce = 6840
	} else if randNum > 40 && randNum <= 80 {
		skillId = 38
		skillForce = 6840
	}	


	return
}

func randSkillWithEnemy_543(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 9
		skillForce = 7080
	} else if randNum > 40 && randNum <= 80 {
		skillId = 10
		skillForce = 7080
	}	


	return
}

func randSkillWithEnemy_544(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1480
		skillForce = 7320
	}	


	return
}

func randSkillWithEnemy_545(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1215
		skillForce = 5940
	} else if randNum > 40 && randNum <= 80 {
		skillId = 21
		skillForce = 7920
	}	


	return
}

func randSkillWithEnemy_546(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1485
		skillForce = 8360
	} else if randNum > 30 && randNum <= 70 {
		skillId = 39
		skillForce = 8360
	}	


	return
}

func randSkillWithEnemy_565(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_566(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_567(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 3240
	}	


	return
}

func randSkillWithEnemy_568(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 3240
	}	


	return
}

func randSkillWithEnemy_569(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 20
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_570(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 3160
	}	


	return
}

func randSkillWithEnemy_571(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1070
		skillForce = 3240
	}	


	return
}

func randSkillWithEnemy_572(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 3240
	}	


	return
}

func randSkillWithEnemy_573(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 41
		skillForce = 5340
	} else if randNum > 40 && randNum <= 60 {
		skillId = 42
		skillForce = 5340
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1035
		skillForce = 5340
	}	


	return
}

func randSkillWithEnemy_574(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 38
		skillForce = 5460
	} else if randNum > 40 && randNum <= 70 {
		skillId = 39
		skillForce = 5460
	}	


	return
}

func randSkillWithEnemy_575(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 3320
	}	


	return
}

func randSkillWithEnemy_576(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 3320
	}	


	return
}

func randSkillWithEnemy_577(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 3400
	}	


	return
}

func randSkillWithEnemy_578(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3400
	}	


	return
}

func randSkillWithEnemy_579(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 3320
	}	


	return
}

func randSkillWithEnemy_580(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 3320
	}	


	return
}

func randSkillWithEnemy_581(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 26
		skillForce = 3400
	}	


	return
}

func randSkillWithEnemy_582(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3400
	}	


	return
}

func randSkillWithEnemy_583(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 43
		skillForce = 5580
	} else if randNum > 40 && randNum <= 70 {
		skillId = 16
		skillForce = 5580
	}	


	return
}

func randSkillWithEnemy_584(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 5700
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 5700
	} else if randNum > 40 && randNum <= 60 {
		skillId = 25
		skillForce = 5700
	} else if randNum > 60 && randNum <= 80 {
		skillId = 25
		skillForce = 5700
	}	


	return
}

func randSkillWithEnemy_585(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_586(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_587(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3600
	}	


	return
}

func randSkillWithEnemy_588(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 39
		skillForce = 3600
	}	


	return
}

func randSkillWithEnemy_589(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_590(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_591(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 3600
	}	


	return
}

func randSkillWithEnemy_592(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 39
		skillForce = 3600
	}	


	return
}

func randSkillWithEnemy_593(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 5850
	} else if randNum > 30 && randNum <= 60 {
		skillId = 37
		skillForce = 5850
	}	


	return
}

func randSkillWithEnemy_594(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 6000
	} else if randNum > 30 && randNum <= 70 {
		skillId = 39
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_595(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 3700
	}	


	return
}

func randSkillWithEnemy_596(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 3700
	}	


	return
}

func randSkillWithEnemy_597(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_598(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_599(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 3700
	}	


	return
}

func randSkillWithEnemy_600(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 29
		skillForce = 3700
	}	


	return
}

func randSkillWithEnemy_601(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 29
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_602(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_603(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 6150
	} else if randNum > 40 && randNum <= 70 {
		skillId = 28
		skillForce = 6150
	}	


	return
}

func randSkillWithEnemy_604(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 44
		skillForce = 6300
	} else if randNum > 35 && randNum <= 70 {
		skillId = 43
		skillForce = 6300
	}	


	return
}

func randSkillWithEnemy_605(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 3900
	}	


	return
}

func randSkillWithEnemy_606(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 27
		skillForce = 3900
	}	


	return
}

func randSkillWithEnemy_607(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 11
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_608(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_609(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 9
		skillForce = 3900
	}	


	return
}

func randSkillWithEnemy_610(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 29
		skillForce = 3900
	}	


	return
}

func randSkillWithEnemy_611(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 12
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_612(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 25
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_613(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 28
		skillForce = 6450
	} else if randNum > 35 && randNum <= 70 {
		skillId = 29
		skillForce = 6450
	}	


	return
}

func randSkillWithEnemy_614(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 38
		skillForce = 6600
	} else if randNum > 35 && randNum <= 70 {
		skillId = 39
		skillForce = 6600
	}	


	return
}

func randSkillWithEnemy_616(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 38
		skillForce = 9000
	} else if randNum > 40 && randNum <= 70 {
		skillId = 39
		skillForce = 9000
	}	


	return
}

func randSkillWithEnemy_617(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 43
		skillForce = 9600
	} else if randNum > 30 && randNum <= 60 {
		skillId = 16
		skillForce = 9600
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1468
		skillForce = 9600
	}	


	return
}

func randSkillWithEnemy_618(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 39
		skillForce = 10000
	}	


	return
}

func randSkillWithEnemy_619(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1482
		skillForce = 10400
	} else if randNum > 40 && randNum <= 70 {
		skillId = 29
		skillForce = 4700
	}	


	return
}

func randSkillWithEnemy_620(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 39
		skillForce = 10800
	} else if randNum > 35 && randNum <= 70 {
		skillId = 38
		skillForce = 10800
	}	


	return
}

func randSkillWithEnemy_623(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 16
		skillForce = 1000
	} else if randNum > 20 && randNum <= 50 {
		skillId = 17
		skillForce = 1000
	} else if randNum > 50 && randNum <= 80 {
		skillId = 18
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_626(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 24
		skillForce = 2340
	} else if randNum > 20 && randNum <= 50 {
		skillId = 25
		skillForce = 2340
	} else if randNum > 50 && randNum <= 80 {
		skillId = 26
		skillForce = 2340
	}	


	return
}

func randSkillWithEnemy_629(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 16
		skillForce = 2000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 17
		skillForce = 2000
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1070
		skillForce = 2000
	} else if randNum > 60 && randNum <= 80 {
		skillId = 22
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_632(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 41
		skillForce = 4200
	} else if randNum > 20 && randNum <= 50 {
		skillId = 42
		skillForce = 4200
	} else if randNum > 50 && randNum <= 80 {
		skillId = 40
		skillForce = 4200
	}	


	return
}

func randSkillWithEnemy_635(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 39
		skillForce = 3180
	} else if randNum > 20 && randNum <= 40 {
		skillId = 35
		skillForce = 3180
	} else if randNum > 40 && randNum <= 60 {
		skillId = 41
		skillForce = 3180
	} else if randNum > 60 && randNum <= 80 {
		skillId = 40
		skillForce = 3180
	}	


	return
}

func randSkillWithEnemy_636(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 20
		skillForce = 300
	} else if randNum > 60 && randNum <= 90 {
		skillId = 22
		skillForce = 300
	} else if randNum > 90 && randNum <= 100 {
		skillId = 108
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_637(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1038
		skillForce = 200
	} else if randNum > 40 && randNum <= 80 {
		skillId = 39
		skillForce = 200
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1147
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_638(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 1048
		skillForce = 2000
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1049
		skillForce = 1000
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1428
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_639(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1052
		skillForce = 3000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1053
		skillForce = 3000
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1051
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_640(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1426
		skillForce = 2000
	} else if randNum > 20 && randNum <= 70 {
		skillId = 1427
		skillForce = 2000
	} else if randNum > 70 && randNum <= 100 {
		skillId = 1429
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_641(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1430
		skillForce = 200
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1431
		skillForce = 200
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1432
		skillForce = 300
	}	


	return
}

func randSkillWithEnemy_642(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1430
		skillForce = 200
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1431
		skillForce = 200
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1432
		skillForce = 300
	}	


	return
}

func randSkillWithEnemy_643(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 17
		skillForce = 620
	} else if randNum > 30 && randNum <= 50 {
		skillId = 16
		skillForce = 620
	} else if randNum > 50 && randNum <= 80 {
		skillId = 18
		skillForce = 620
	}	


	return
}

func randSkillWithEnemy_644(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 45
		skillForce = 1180
	} else if randNum > 30 && randNum <= 50 {
		skillId = 37
		skillForce = 1180
	} else if randNum > 50 && randNum <= 70 {
		skillId = 13
		skillForce = 1150
	}	


	return
}

func randSkillWithEnemy_645(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 20
		skillForce = 1400
	} else if randNum > 20 && randNum <= 40 {
		skillId = 21
		skillForce = 1400
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1070
		skillForce = 1400
	} else if randNum > 60 && randNum <= 80 {
		skillId = 22
		skillForce = 1400
	}	


	return
}

func randSkillWithEnemy_646(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 13
		skillForce = 1520
	} else if randNum > 30 && randNum <= 50 {
		skillId = 28
		skillForce = 1520
	} else if randNum > 50 && randNum <= 80 {
		skillId = 14
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_647(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 25
		skillForce = 1760
	} else if randNum > 20 && randNum <= 50 {
		skillId = 24
		skillForce = 1760
	} else if randNum > 50 && randNum <= 80 {
		skillId = 26
		skillForce = 1760
	}	


	return
}

func randSkillWithEnemy_648(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 12
		skillForce = 2760
	} else if randNum > 20 && randNum <= 40 {
		skillId = 13
		skillForce = 2760
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1268
		skillForce = 2760
	} else if randNum > 60 && randNum <= 80 {
		skillId = 14
		skillForce = 2760
	}	


	return
}

func randSkillWithEnemy_669(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 36
		skillForce = 720
	}	


	return
}

func randSkillWithEnemy_670(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1038
		skillForce = 720
	}	


	return
}

func randSkillWithEnemy_671(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_675(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1479
		skillForce = 615
	}	


	return
}

func randSkillWithEnemy_676(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1487
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_677(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 10
		skillForce = 980
	}	


	return
}

func randSkillWithEnemy_679(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 29
		skillForce = 1100
	}	


	return
}

func randSkillWithEnemy_680(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 26
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_681(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1488
		skillForce = 1360
	}	


	return
}

func randSkillWithEnemy_682(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1476
		skillForce = 1520
	}	


	return
}

func randSkillWithEnemy_684(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 31
		skillForce = 1845
	}	


	return
}

func randSkillWithEnemy_685(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 31
		skillForce = 1845
	}	


	return
}

func randSkillWithEnemy_686(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 17
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_687(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 16
		skillForce = 2720
	}	


	return
}

func randSkillWithEnemy_688(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1477
		skillForce = 2115
	}	


	return
}

func randSkillWithEnemy_689(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1478
		skillForce = 2115
	}	


	return
}

func randSkillWithEnemy_690(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1489
		skillForce = 2190
	}	


	return
}

func randSkillWithEnemy_693(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 39
		skillForce = 3220
	}	


	return
}

func randSkillWithEnemy_695(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 108
		skillForce = 15000
	}	


	return
}

func randSkillWithEnemy_696(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 38
		skillForce = 3700
	}	


	return
}

func randSkillWithEnemy_697(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 43
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_699(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1481
		skillForce = 4400
	}	


	return
}

func randSkillWithEnemy_700(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1482
		skillForce = 4700
	}	


	return
}

func randSkillWithEnemy_701(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1487
		skillForce = 30000
	}	


	return
}

func randSkillWithEnemy_702(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1498
		skillForce = 3500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1495
		skillForce = 3500
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1218
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_703(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1491
		skillForce = 3500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1496
		skillForce = 3500
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1485
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_704(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1492
		skillForce = 3500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1497
		skillForce = 1000
	} else if randNum > 80 && randNum <= 100 {
		skillId = 14
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_706(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1493
		skillForce = 3500
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1494
		skillForce = 3500
	} else if randNum > 80 && randNum <= 100 {
		skillId = 41
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_713(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 11
		skillForce = 18000
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1492
		skillForce = 18000
	}	


	return
}

func randSkillWithEnemy_714(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 35
		skillForce = 4160
	}	


	return
}

func randSkillWithEnemy_715(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 4160
	}	


	return
}

func randSkillWithEnemy_716(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 4160
	}	


	return
}

func randSkillWithEnemy_717(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 37
		skillForce = 4160
	}	


	return
}

func randSkillWithEnemy_718(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1505
		skillForce = 6750
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1506
		skillForce = 6750
	} else if randNum > 70 && randNum <= 100 {
		skillId = 1507
		skillForce = 6750
	}	


	return
}

func randSkillWithEnemy_719(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 6900
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1509
		skillForce = 6900
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1039
		skillForce = 6900
	}	


	return
}

func randSkillWithEnemy_720(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1504
		skillForce = 5400
	}	


	return
}

func randSkillWithEnemy_721(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 5400
	}	


	return
}

func randSkillWithEnemy_722(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 5600
	}	


	return
}

func randSkillWithEnemy_723(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 11200
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1039
		skillForce = 11200
	}	


	return
}

func randSkillWithEnemy_724(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1510
		skillForce = 11600
	} else if randNum > 30 && randNum <= 50 {
		skillId = 1508
		skillForce = 11600
	} else if randNum > 50 && randNum <= 70 {
		skillId = 1500
		skillForce = 11600
	}	


	return
}

func randSkillWithEnemy_725(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1511
		skillForce = 12000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1512
		skillForce = 12000
	}	


	return
}

func randSkillWithEnemy_726(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1510
		skillForce = 12400
	} else if randNum > 40 && randNum <= 60 {
		skillId = 43
		skillForce = 12400
	} else if randNum > 60 && randNum <= 70 {
		skillId = 1518
		skillForce = 1700
	}	


	return
}

func randSkillWithEnemy_727(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 1519
		skillForce = 12800
	} else if randNum > 35 && randNum <= 70 {
		skillId = 39
		skillForce = 12800
	}	


	return
}

func randSkillWithEnemy_728(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 35
		skillForce = 4240
	}	


	return
}

func randSkillWithEnemy_729(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 4240
	}	


	return
}

func randSkillWithEnemy_730(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1477
		skillForce = 4320
	}	


	return
}

func randSkillWithEnemy_731(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1504
		skillForce = 4320
	}	


	return
}

func randSkillWithEnemy_732(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1510
		skillForce = 7050
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1508
		skillForce = 7050
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1500
		skillForce = 7050
	}	


	return
}

func randSkillWithEnemy_733(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 38
		skillForce = 7200
	} else if randNum > 40 && randNum <= 70 {
		skillId = 39
		skillForce = 7200
	}	


	return
}

func randSkillWithEnemy_734(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 23
		skillForce = 4400
	}	


	return
}

func randSkillWithEnemy_735(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 37
		skillForce = 4400
	}	


	return
}

func randSkillWithEnemy_736(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 23
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_737(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 10
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_738(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 7200
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1511
		skillForce = 7200
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1512
		skillForce = 7200
	}	


	return
}

func randSkillWithEnemy_739(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1511
		skillForce = 7350
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1512
		skillForce = 7350
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1513
		skillForce = 7350
	}	


	return
}

func randSkillWithEnemy_740(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 23
		skillForce = 4600
	}	


	return
}

func randSkillWithEnemy_741(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 4600
	}	


	return
}

func randSkillWithEnemy_742(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 4700
	}	


	return
}

func randSkillWithEnemy_743(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 10
		skillForce = 4700
	}	


	return
}

func randSkillWithEnemy_744(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 33
		skillForce = 7500
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1514
		skillForce = 7500
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1504
		skillForce = 7500
	}	


	return
}

func randSkillWithEnemy_745(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1510
		skillForce = 7650
	} else if randNum > 40 && randNum <= 70 {
		skillId = 43
		skillForce = 7650
	}	


	return
}

func randSkillWithEnemy_746(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 4800
	}	


	return
}

func randSkillWithEnemy_747(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 4800
	}	


	return
}

func randSkillWithEnemy_748(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1477
		skillForce = 4900
	}	


	return
}

func randSkillWithEnemy_749(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1477
		skillForce = 4900
	}	


	return
}

func randSkillWithEnemy_750(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 39
		skillForce = 7800
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1512
		skillForce = 7800
	}	


	return
}

func randSkillWithEnemy_751(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1503
		skillForce = 7950
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1500
		skillForce = 7950
	}	


	return
}

func randSkillWithEnemy_752(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 42
		skillForce = 3450
	} else if randNum > 40 && randNum <= 80 {
		skillId = 41
		skillForce = 3450
	}	


	return
}

func randSkillWithEnemy_753(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1476
		skillForce = 3450
	}	


	return
}

func randSkillWithEnemy_754(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 20
		skillForce = 3630
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1224
		skillForce = 3630
	}	


	return
}

func randSkillWithEnemy_755(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 3630
	}	


	return
}

func randSkillWithEnemy_758(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 43
		skillForce = 4500
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1516
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_761(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1512
		skillForce = 4800
	} else if randNum > 20 && randNum <= 50 {
		skillId = 17
		skillForce = 4800
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1513
		skillForce = 4800
	}	


	return
}

func randSkillWithEnemy_766(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1585
		skillForce = 8000
	} else if randNum > 70 && randNum <= 100 {
		skillId = 1586
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_767(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1492
		skillForce = 8000
	} else if randNum > 50 && randNum <= 80 {
		skillId = 1497
		skillForce = 3000
	} else if randNum > 80 && randNum <= 100 {
		skillId = 14
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_768(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1426
		skillForce = 8000
	} else if randNum > 20 && randNum <= 70 {
		skillId = 1427
		skillForce = 8000
	} else if randNum > 70 && randNum <= 100 {
		skillId = 1429
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_769(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 1048
		skillForce = 8000
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1049
		skillForce = 4000
	} else if randNum > 80 && randNum <= 100 {
		skillId = 1428
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_770(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 20
		skillForce = 8000
	} else if randNum > 60 && randNum <= 90 {
		skillId = 22
		skillForce = 8000
	} else if randNum > 90 && randNum <= 100 {
		skillId = 1050
		skillForce = 50000
	}	


	return
}

func randSkillWithEnemy_771(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1591
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_772(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1592
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_773(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1593
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_774(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1594
		skillForce = 30000
	}	


	return
}

func randSkillWithEnemy_775(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1598
		skillForce = 4000
	} else if randNum > 20 && randNum <= 60 {
		skillId = 33
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_776(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1504
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_777(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 4450
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1429
		skillForce = 4450
	}	


	return
}

func randSkillWithEnemy_778(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1478
		skillForce = 4450
	}	


	return
}

func randSkillWithEnemy_779(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 12
		skillForce = 4800
	} else if randNum > 40 && randNum <= 80 {
		skillId = 10
		skillForce = 4800
	}	


	return
}

func randSkillWithEnemy_780(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 4800
	}	


	return
}

func randSkillWithEnemy_795(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1599
		skillForce = 13200
	}	


	return
}

func randSkillWithEnemy_796(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1048
		skillForce = 11200
	}	


	return
}

func randSkillWithEnemy_797(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1597
		skillForce = 13680
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1596
		skillForce = 13680
	}	


	return
}

func randSkillWithEnemy_798(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 21
		skillForce = 5840
	}	


	return
}

func randSkillWithEnemy_799(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1513
		skillForce = 14160
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1512
		skillForce = 14160
	}	


	return
}

func randSkillWithEnemy_800(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1513
		skillForce = 6080
	}	


	return
}

func randSkillWithEnemy_801(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 14640
	} else if randNum > 40 && randNum <= 60 {
		skillId = 33
		skillForce = 14640
	} else if randNum > 60 && randNum <= 70 {
		skillId = 32
		skillForce = 14640
	}	


	return
}

func randSkillWithEnemy_802(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1504
		skillForce = 6320
	}	


	return
}

func randSkillWithEnemy_803(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 1510
		skillForce = 15120
	} else if randNum > 35 && randNum <= 70 {
		skillId = 1515
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_804(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1048
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_805(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 108
		skillForce = 30000
	}	


	return
}

func randSkillWithEnemy_806(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1048
		skillForce = 5120
	}	


	return
}

func randSkillWithEnemy_807(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 108
		skillForce = 30000
	}	


	return
}

func randSkillWithEnemy_808(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1599
		skillForce = 8100
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1428
		skillForce = 8100
	}	


	return
}

func randSkillWithEnemy_809(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1605
		skillForce = 8280
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1510
		skillForce = 8280
	}	


	return
}

func randSkillWithEnemy_810(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 5240
	}	


	return
}

func randSkillWithEnemy_811(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 5240
	}	


	return
}

func randSkillWithEnemy_812(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 5360
	}	


	return
}

func randSkillWithEnemy_813(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 15
		skillForce = 5360
	}	


	return
}

func randSkillWithEnemy_814(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 17
		skillForce = 8460
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1513
		skillForce = 8460
	}	


	return
}

func randSkillWithEnemy_815(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1512
		skillForce = 8640
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1513
		skillForce = 8640
	}	


	return
}

func randSkillWithEnemy_816(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1147
		skillForce = 20000
	}	


	return
}

func randSkillWithEnemy_817(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1481
		skillForce = 5480
	}	


	return
}

func randSkillWithEnemy_818(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 5600
	}	


	return
}

func randSkillWithEnemy_819(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 38
		skillForce = 5600
	}	


	return
}

func randSkillWithEnemy_820(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 36
		skillForce = 8820
	} else if randNum > 40 && randNum <= 60 {
		skillId = 35
		skillForce = 8820
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1038
		skillForce = 8820
	}	


	return
}

func randSkillWithEnemy_821(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1597
		skillForce = 9000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1596
		skillForce = 9000
	}	


	return
}

func randSkillWithEnemy_822(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 23
		skillForce = 5720
	}	


	return
}

func randSkillWithEnemy_823(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 35
		skillForce = 5720
	}	


	return
}

func randSkillWithEnemy_824(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 5840
	}	


	return
}

func randSkillWithEnemy_825(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 35
		skillForce = 5840
	}	


	return
}

func randSkillWithEnemy_826(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 33
		skillForce = 9180
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1514
		skillForce = 9180
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1504
		skillForce = 9180
	}	


	return
}

func randSkillWithEnemy_827(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 9360
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1598
		skillForce = 9360
	}	


	return
}

func randSkillWithEnemy_828(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 10
		skillForce = 5960
	}	


	return
}

func randSkillWithEnemy_829(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1477
		skillForce = 5960
	}	


	return
}

func randSkillWithEnemy_830(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1138
		skillForce = 6080
	}	


	return
}

func randSkillWithEnemy_831(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1478
		skillForce = 6080
	}	


	return
}

func randSkillWithEnemy_832(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 9540
	} else if randNum > 40 && randNum <= 70 {
		skillId = 33
		skillForce = 9540
	}	


	return
}

func randSkillWithEnemy_833(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1510
		skillForce = 9720
	} else if randNum > 40 && randNum <= 70 {
		skillId = 29
		skillForce = 9720
	}	


	return
}

func randSkillWithEnemy_836(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 21
		skillForce = 5100
	} else if randNum > 20 && randNum <= 40 {
		skillId = 20
		skillForce = 5100
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1070
		skillForce = 5100
	} else if randNum > 60 && randNum <= 80 {
		skillId = 22
		skillForce = 5100
	}	


	return
}

func randSkillWithEnemy_839(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 17
		skillForce = 5400
	} else if randNum > 20 && randNum <= 40 {
		skillId = 16
		skillForce = 5400
	} else if randNum > 40 && randNum <= 70 {
		skillId = 18
		skillForce = 5400
	}	


	return
}

func randSkillWithEnemy_842(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 45
		skillForce = 5700
	} else if randNum > 30 && randNum <= 60 {
		skillId = 37
		skillForce = 5700
	} else if randNum > 60 && randNum <= 80 {
		skillId = 13
		skillForce = 5700
	}	


	return
}

func randSkillWithEnemy_843(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1505
		skillForce = 15600
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1506
		skillForce = 15600
	} else if randNum > 60 && randNum <= 100 {
		skillId = 1507
		skillForce = 15600
	}	


	return
}

func randSkillWithEnemy_844(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1500
		skillForce = 6800
	}	


	return
}

func randSkillWithEnemy_845(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1599
		skillForce = 16000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1629
		skillForce = 16000
	}	


	return
}

func randSkillWithEnemy_846(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1631
		skillForce = 7000
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1639
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_847(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1513
		skillForce = 16400
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1512
		skillForce = 16400
	}	


	return
}

func randSkillWithEnemy_848(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 7200
	}	


	return
}

func randSkillWithEnemy_849(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1430
		skillForce = 16800
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1431
		skillForce = 16800
	}	


	return
}

func randSkillWithEnemy_850(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 39
		skillForce = 7400
	}	


	return
}

func randSkillWithEnemy_851(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1627
		skillForce = 17200
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1637
		skillForce = 17200
	}	


	return
}

func randSkillWithEnemy_859(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1636
		skillForce = 5100
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1478
		skillForce = 5100
	}	


	return
}

func randSkillWithEnemy_860(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1039
		skillForce = 5100
	}	


	return
}

func randSkillWithEnemy_861(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 5570
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1504
		skillForce = 5570
	}	


	return
}

func randSkillWithEnemy_862(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1513
		skillForce = 5570
	}	


	return
}

func randSkillWithEnemy_865(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1512
		skillForce = 6800
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1513
		skillForce = 6800
	}	


	return
}

func randSkillWithEnemy_868(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 21
		skillForce = 7400
	} else if randNum > 30 && randNum <= 60 {
		skillId = 20
		skillForce = 7400
	} else if randNum > 60 && randNum <= 80 {
		skillId = 19
		skillForce = 7400
	}	


	return
}

func randSkillWithEnemy_947(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 6200
	}	


	return
}

func randSkillWithEnemy_948(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1627
		skillForce = 6200
	}	


	return
}

func randSkillWithEnemy_949(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 6260
	}	


	return
}

func randSkillWithEnemy_950(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1500
		skillForce = 6260
	}	


	return
}

func randSkillWithEnemy_951(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1500
		skillForce = 10050
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1628
		skillForce = 10050
	}	


	return
}

func randSkillWithEnemy_952(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1505
		skillForce = 10140
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1506
		skillForce = 10140
	} else if randNum > 60 && randNum <= 100 {
		skillId = 1507
		skillForce = 10140
	}	


	return
}

func randSkillWithEnemy_953(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 6320
	}	


	return
}

func randSkillWithEnemy_954(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1630
		skillForce = 6320
	}	


	return
}

func randSkillWithEnemy_955(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1629
		skillForce = 6380
	}	


	return
}

func randSkillWithEnemy_956(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1631
		skillForce = 6380
	}	


	return
}

func randSkillWithEnemy_957(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1599
		skillForce = 10320
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1629
		skillForce = 10320
	}	


	return
}

func randSkillWithEnemy_958(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1597
		skillForce = 10230
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1596
		skillForce = 10230
	}	


	return
}

func randSkillWithEnemy_959(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 6440
	}	


	return
}

func randSkillWithEnemy_960(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1511
		skillForce = 6440
	}	


	return
}

func randSkillWithEnemy_961(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1513
		skillForce = 6500
	}	


	return
}

func randSkillWithEnemy_962(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1512
		skillForce = 6500
	}	


	return
}

func randSkillWithEnemy_963(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1513
		skillForce = 10410
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1512
		skillForce = 10410
	}	


	return
}

func randSkillWithEnemy_964(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1511
		skillForce = 10500
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1635
		skillForce = 10500
	}	


	return
}

func randSkillWithEnemy_965(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 6560
	}	


	return
}

func randSkillWithEnemy_966(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1482
		skillForce = 6560
	}	


	return
}

func randSkillWithEnemy_967(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 6620
	}	


	return
}

func randSkillWithEnemy_968(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1039
		skillForce = 6620
	}	


	return
}

func randSkillWithEnemy_969(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1519
		skillForce = 10590
	} else if randNum > 40 && randNum <= 60 {
		skillId = 1627
		skillForce = 10590
	} else if randNum > 60 && randNum <= 80 {
		skillId = 1039
		skillForce = 10590
	}	


	return
}

func randSkillWithEnemy_970(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1430
		skillForce = 10680
	} else if randNum > 40 && randNum <= 80 {
		skillId = 1431
		skillForce = 10680
	}	


	return
}

func randSkillWithEnemy_971(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1639
		skillForce = 6680
	}	


	return
}

func randSkillWithEnemy_972(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1513
		skillForce = 6680
	}	


	return
}

func randSkillWithEnemy_973(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 25
		skillForce = 6740
	}	


	return
}

func randSkillWithEnemy_974(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 6740
	}	


	return
}

func randSkillWithEnemy_975(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1627
		skillForce = 10770
	} else if randNum > 40 && randNum <= 70 {
		skillId = 1637
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_976(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1515
		skillForce = 5000
	} else if randNum > 30 && randNum <= 70 {
		skillId = 1508
		skillForce = 10860
	}	


	return
}

func randSkillWithEnemy_979(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 10
		skillForce = 610
	}	


	return
}

func randSkillWithEnemy_980(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1639
		skillForce = 900
	}	


	return
}

func randSkillWithEnemy_982(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1715
		skillForce = 10
	}	


	return
}

func randSkillWithEnemy_983(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1717
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_984(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 1716
		skillForce = 2280
	}	


	return
}

func randSkillWithEnemy_987(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 1718
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_988(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 1047
		skillForce = 15000
	}	


	return
}

func randSkillWithEnemy_989(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1718
		skillForce = 1300
	}	


	return
}

func randSkillWithEnemy_990(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 1719
		skillForce = 1200
	}	


	return
}

func randSkillWithEnemy_991(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1720
		skillForce = 1500
	}	


	return
}

func randSkillWithEnemy_992(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1721
		skillForce = 1700
	}	


	return
}

func randSkillWithEnemy_993(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_994(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1722
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_995(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1722
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_996(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1719
		skillForce = 2200
	}	


	return
}

func randSkillWithEnemy_997(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1725
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_998(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 38
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_999(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1726
		skillForce = 1000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 1727
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1000(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1704
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_1002(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 19
		skillForce = 2600
	}	


	return
}

func randSkillWithEnemy_1005(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_1007(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 31
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1008(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1720
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1009(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1500
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1010(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 33
		skillForce = 2700
	} else if randNum > 50 && randNum <= 100 {
		skillId = 1730
		skillForce = 2700
	}	


	return
}

func randSkillWithEnemy_1011(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 32
		skillForce = 2700
	} else if randNum > 50 && randNum <= 100 {
		skillId = 1730
		skillForce = 2700
	}	


	return
}

func randSkillWithEnemy_1012(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1731
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1014(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 108
		skillForce = 20000
	} else if randNum > 35 && randNum <= 70 {
		skillId = 1047
		skillForce = 10000
	}	


	return
}

func randSkillWithEnemy_1015(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1598
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_1016(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 45
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1017(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1732
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1018(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1733
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1019(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1512
		skillForce = 3000
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1035
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1020(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1736
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1021(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 35
		skillForce = 5500
	}	


	return
}

func randSkillWithEnemy_1022(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 39
		skillForce = 5500
	}	


	return
}

func randSkillWithEnemy_1023(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1730
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1025(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 38
		skillForce = 5700
	} else if randNum > 30 && randNum <= 60 {
		skillId = 39
		skillForce = 5700
	}	


	return
}

func randSkillWithEnemy_1026(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 15 {
		skillId = 1715
		skillForce = 20
	}	


	return
}

func randSkillWithEnemy_1028(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1703
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_1029(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1508
		skillForce = 7200
	}	


	return
}

func randSkillWithEnemy_1030(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1477
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1031(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1032(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1729
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1034(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1728
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1035(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1732
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1036(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1715
		skillForce = 30
	}	


	return
}

func randSkillWithEnemy_1037(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 38
		skillForce = 8000
	} else if randNum > 50 && randNum <= 100 {
		skillId = 39
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_1038(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1737
		skillForce = 200
	}	


	return
}

func randSkillWithEnemy_1039(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1738
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1040(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1703
		skillForce = 200
	}	


	return
}

func randSkillWithEnemy_1041(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1419
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_1042(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 38
		skillForce = 1040
	}	


	return
}

func randSkillWithEnemy_1045(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1740
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1046(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 35
		skillForce = 2160
	}	


	return
}

func randSkillWithEnemy_1048(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1702
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_1049(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1720
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_1050(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1741
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1051(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1597
		skillForce = 1800
	}	


	return
}

func randSkillWithEnemy_1052(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1742
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_1054(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1743
		skillForce = 1000
	} else if randNum > 20 && randNum <= 40 {
		skillId = 1734
		skillForce = 20
	}	


	return
}

func randSkillWithEnemy_1055(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1508
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1056(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1057(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1038
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1058(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1744
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_1059(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1745
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1060(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1746
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1061(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1747
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1062(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1508
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1063(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 2800
	}	


	return
}

func randSkillWithEnemy_1064(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 39
		skillForce = 2600
	}	


	return
}

func randSkillWithEnemy_1065(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1732
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1066(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 39
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1067(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 38
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1069(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1748
		skillForce = 3500
	}	


	return
}

func randSkillWithEnemy_1070(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1749
		skillForce = 3800
	} else if randNum > 50 && randNum <= 100 {
		skillId = 1750
		skillForce = 20
	}	


	return
}

func randSkillWithEnemy_1071(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1751
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1072(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 108
		skillForce = 20000
	}	


	return
}

func randSkillWithEnemy_1073(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1732
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_1074(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1738
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1075(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 21
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_1076(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 19
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_1077(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1639
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1082(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1480
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1083(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1698
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1084(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 39
		skillForce = 5500
	}	


	return
}

func randSkillWithEnemy_1086(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1500
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_1087(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1147
		skillForce = 30000
	}	


	return
}

func randSkillWithEnemy_1088(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 80 {
		skillId = 11
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1089(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 41
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_1090(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 35
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_1091(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1507
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_1093(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 36
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1095(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 45
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_1099(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 31
		skillForce = 900
	}	


	return
}

func randSkillWithEnemy_1100(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 15
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1101(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1048
		skillForce = 1180
	}	


	return
}

func randSkillWithEnemy_1102(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 108
		skillForce = 1260
	}	


	return
}

func randSkillWithEnemy_1103(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 31
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_1104(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 31
		skillForce = 1250
	}	


	return
}

func randSkillWithEnemy_1105(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 60 {
		skillId = 31
		skillForce = 1500
	}	


	return
}

func randSkillWithEnemy_1106(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 1 {
		skillId = 1512
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1109(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1698
		skillForce = 1500
	}	


	return
}

func randSkillWithEnemy_1110(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 15
		skillForce = 1700
	}	


	return
}

func randSkillWithEnemy_1111(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1639
		skillForce = 1700
	}	


	return
}

func randSkillWithEnemy_1112(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1138
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1113(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1138
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_1114(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1700
		skillForce = 1800
	}	


	return
}

func randSkillWithEnemy_1115(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1701
		skillForce = 2100
	}	


	return
}

func randSkillWithEnemy_1116(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1702
		skillForce = 2300
	}	


	return
}

func randSkillWithEnemy_1117(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 31
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1118(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 3800
	}	


	return
}

func randSkillWithEnemy_1119(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1786
		skillForce = 100
	}	


	return
}

func randSkillWithEnemy_1121(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 33
		skillForce = 2900
	} else if randNum > 30 && randNum <= 60 {
		skillId = 32
		skillForce = 2900
	}	


	return
}

func randSkillWithEnemy_1122(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 33
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_1123(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1514
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1124(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 100 {
		skillId = 1739
		skillForce = 1500
	}	


	return
}

func randSkillWithEnemy_1125(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1477
		skillForce = 3400
	} else if randNum > 20 && randNum <= 50 {
		skillId = 1704
		skillForce = 3400
	}	


	return
}

func randSkillWithEnemy_1126(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 10
		skillForce = 4100
	}	


	return
}

func randSkillWithEnemy_1127(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1705
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1128(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1706
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1129(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1631
		skillForce = 2000
	} else if randNum > 20 && randNum <= 50 {
		skillId = 1639
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1130(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1708
		skillForce = 3000
	} else if randNum > 30 && randNum <= 50 {
		skillId = 1704
		skillForce = 2500
	}	


	return
}

func randSkillWithEnemy_1131(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 22
		skillForce = 4000
	}	


	return
}

func randSkillWithEnemy_1132(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1708
		skillForce = 5200
	}	


	return
}

func randSkillWithEnemy_1133(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1709
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1134(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1605
		skillForce = 5000
	}	


	return
}

func randSkillWithEnemy_1135(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1712
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_1136(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1712
		skillForce = 6000
	}	


	return
}

func randSkillWithEnemy_1137(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 23
		skillForce = 3000
	} else if randNum > 50 && randNum <= 100 {
		skillId = 1712
		skillForce = 4500
	}	


	return
}

func randSkillWithEnemy_1138(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1735
		skillForce = 6300
	}	


	return
}

func randSkillWithEnemy_1140(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 10
		skillForce = 6600
	}	


	return
}

func randSkillWithEnemy_1141(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1477
		skillForce = 6600
	}	


	return
}

func randSkillWithEnemy_1142(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 1504
		skillForce = 7000
	}	


	return
}

func randSkillWithEnemy_1144(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1515
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1146(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 20 {
		skillId = 1713
		skillForce = 8000
	} else if randNum > 20 && randNum <= 35 {
		skillId = 1709
		skillForce = 8000
	}	


	return
}

func randSkillWithEnemy_1150(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 84 {
		skillId = 36
		skillForce = 9999999
	} else if randNum > 84 && randNum <= 99 {
		skillId = 1775
		skillForce = 999999
	} else if randNum > 99 && randNum <= 100 {
		skillId = 1780
		skillForce = 9999999
	}	


	return
}

func randSkillWithEnemy_1161(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1489
		skillForce = 3310
	}	


	return
}

func randSkillWithEnemy_1162(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 108
		skillForce = 3000
	}	


	return
}

func randSkillWithEnemy_1163(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 10
		skillForce = 370
	}	


	return
}

func randSkillWithEnemy_1164(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 29
		skillForce = 500
	}	


	return
}

func randSkillWithEnemy_1166(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 35 {
		skillId = 1751
		skillForce = 880
	}	


	return
}

func randSkillWithEnemy_1167(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 70 {
		skillId = 1736
		skillForce = 1000
	}	


	return
}

func randSkillWithEnemy_1168(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 44
		skillForce = 1320
	}	


	return
}

func randSkillWithEnemy_1169(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 30 {
		skillId = 1477
		skillForce = 1600
	} else if randNum > 30 && randNum <= 60 {
		skillId = 1708
		skillForce = 1600
	}	


	return
}

func randSkillWithEnemy_1170(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 15
		skillForce = 2000
	}	


	return
}

func randSkillWithEnemy_1172(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1487
		skillForce = 2100
	}	


	return
}

func randSkillWithEnemy_1173(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1096
		skillForce = 2200
	}	


	return
}

func randSkillWithEnemy_1174(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 50 {
		skillId = 1505
		skillForce = 2450
	}	


	return
}

func randSkillWithEnemy_1175(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 25 {
		skillId = 1477
		skillForce = 2710
	} else if randNum > 25 && randNum <= 50 {
		skillId = 1639
		skillForce = 2710
	}	


	return
}

func randSkillWithEnemy_1176(f *Fighter) (skillId, skillForce int) {
	randNum := rand.Intn(100) + 1

	if randNum > 0 && randNum <= 40 {
		skillId = 25
		skillForce = 3000
	}	


	return
}

