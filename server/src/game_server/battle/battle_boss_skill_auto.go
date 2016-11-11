package battle



func skillTriggerCondition_default(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	return skill
}

func createTriggerSkill(roleId int) funcBossTriggerSkill {
	switch roleId {
	case 1:
		return skillTriggerCondition_1
	case 2:
		return skillTriggerCondition_2
	case 3:
		return skillTriggerCondition_3
	case 4:
		return skillTriggerCondition_4
	case 5:
		return skillTriggerCondition_5
	case 8:
		return skillTriggerCondition_8
	case 10:
		return skillTriggerCondition_10
	case 11:
		return skillTriggerCondition_11
	case 14:
		return skillTriggerCondition_14
	case 18:
		return skillTriggerCondition_18
	case 19:
		return skillTriggerCondition_19
	case 23:
		return skillTriggerCondition_23
	case 25:
		return skillTriggerCondition_25
	case 26:
		return skillTriggerCondition_26
	case 38:
		return skillTriggerCondition_38
	case 88:
		return skillTriggerCondition_88
	case 89:
		return skillTriggerCondition_89
	case 96:
		return skillTriggerCondition_96
	case 97:
		return skillTriggerCondition_97
	case 100:
		return skillTriggerCondition_100
	case 101:
		return skillTriggerCondition_101
	case 103:
		return skillTriggerCondition_103
	case 104:
		return skillTriggerCondition_104
	case 105:
		return skillTriggerCondition_105
	case 106:
		return skillTriggerCondition_106
	case 107:
		return skillTriggerCondition_107
	case 110:
		return skillTriggerCondition_110
	case 111:
		return skillTriggerCondition_111
	case 112:
		return skillTriggerCondition_112
	case 113:
		return skillTriggerCondition_113
	case 114:
		return skillTriggerCondition_114
	case 115:
		return skillTriggerCondition_115
	case 151:
		return skillTriggerCondition_151
	case 181:
		return skillTriggerCondition_181
	case 182:
		return skillTriggerCondition_182
	case 183:
		return skillTriggerCondition_183
	case 184:
		return skillTriggerCondition_184
	case 185:
		return skillTriggerCondition_185
	case 186:
		return skillTriggerCondition_186
	case 187:
		return skillTriggerCondition_187
	case 188:
		return skillTriggerCondition_188
	case 199:
		return skillTriggerCondition_199
	case 204:
		return skillTriggerCondition_204
	case 205:
		return skillTriggerCondition_205
	case 206:
		return skillTriggerCondition_206
	case 208:
		return skillTriggerCondition_208
	case 210:
		return skillTriggerCondition_210
	case 214:
		return skillTriggerCondition_214
	case 215:
		return skillTriggerCondition_215
	case 216:
		return skillTriggerCondition_216
	case 217:
		return skillTriggerCondition_217
	case 218:
		return skillTriggerCondition_218
	case 219:
		return skillTriggerCondition_219
	case 220:
		return skillTriggerCondition_220
	case 221:
		return skillTriggerCondition_221
	case 222:
		return skillTriggerCondition_222
	case 223:
		return skillTriggerCondition_223
	case 224:
		return skillTriggerCondition_224
	case 225:
		return skillTriggerCondition_225
	case 226:
		return skillTriggerCondition_226
	case 227:
		return skillTriggerCondition_227
	case 272:
		return skillTriggerCondition_272
	case 273:
		return skillTriggerCondition_273
	case 274:
		return skillTriggerCondition_274
	case 275:
		return skillTriggerCondition_275
	case 276:
		return skillTriggerCondition_276
	case 281:
		return skillTriggerCondition_281
	case 282:
		return skillTriggerCondition_282
	case 289:
		return skillTriggerCondition_289
	case 290:
		return skillTriggerCondition_290
	case 291:
		return skillTriggerCondition_291
	case 292:
		return skillTriggerCondition_292
	case 308:
		return skillTriggerCondition_308
	case 310:
		return skillTriggerCondition_310
	case 360:
		return skillTriggerCondition_360
	case 361:
		return skillTriggerCondition_361
	case 362:
		return skillTriggerCondition_362
	case 363:
		return skillTriggerCondition_363
	case 364:
		return skillTriggerCondition_364
	case 365:
		return skillTriggerCondition_365
	case 366:
		return skillTriggerCondition_366
	case 367:
		return skillTriggerCondition_367
	case 368:
		return skillTriggerCondition_368
	case 369:
		return skillTriggerCondition_369
	case 370:
		return skillTriggerCondition_370
	case 371:
		return skillTriggerCondition_371
	case 372:
		return skillTriggerCondition_372
	case 373:
		return skillTriggerCondition_373
	case 380:
		return skillTriggerCondition_380
	case 387:
		return skillTriggerCondition_387
	case 394:
		return skillTriggerCondition_394
	case 395:
		return skillTriggerCondition_395
	case 396:
		return skillTriggerCondition_396
	case 397:
		return skillTriggerCondition_397
	case 398:
		return skillTriggerCondition_398
	case 399:
		return skillTriggerCondition_399
	case 400:
		return skillTriggerCondition_400
	case 401:
		return skillTriggerCondition_401
	case 402:
		return skillTriggerCondition_402
	case 403:
		return skillTriggerCondition_403
	case 412:
		return skillTriggerCondition_412
	case 419:
		return skillTriggerCondition_419
	case 426:
		return skillTriggerCondition_426
	case 433:
		return skillTriggerCondition_433
	case 437:
		return skillTriggerCondition_437
	case 438:
		return skillTriggerCondition_438
	case 439:
		return skillTriggerCondition_439
	case 440:
		return skillTriggerCondition_440
	case 441:
		return skillTriggerCondition_441
	case 442:
		return skillTriggerCondition_442
	case 443:
		return skillTriggerCondition_443
	case 500:
		return skillTriggerCondition_500
	case 501:
		return skillTriggerCondition_501
	case 510:
		return skillTriggerCondition_510
	case 511:
		return skillTriggerCondition_511
	case 520:
		return skillTriggerCondition_520
	case 521:
		return skillTriggerCondition_521
	case 530:
		return skillTriggerCondition_530
	case 531:
		return skillTriggerCondition_531
	case 540:
		return skillTriggerCondition_540
	case 541:
		return skillTriggerCondition_541
	case 542:
		return skillTriggerCondition_542
	case 543:
		return skillTriggerCondition_543
	case 544:
		return skillTriggerCondition_544
	case 545:
		return skillTriggerCondition_545
	case 546:
		return skillTriggerCondition_546
	case 573:
		return skillTriggerCondition_573
	case 574:
		return skillTriggerCondition_574
	case 583:
		return skillTriggerCondition_583
	case 584:
		return skillTriggerCondition_584
	case 593:
		return skillTriggerCondition_593
	case 594:
		return skillTriggerCondition_594
	case 603:
		return skillTriggerCondition_603
	case 604:
		return skillTriggerCondition_604
	case 613:
		return skillTriggerCondition_613
	case 614:
		return skillTriggerCondition_614
	case 616:
		return skillTriggerCondition_616
	case 617:
		return skillTriggerCondition_617
	case 618:
		return skillTriggerCondition_618
	case 619:
		return skillTriggerCondition_619
	case 620:
		return skillTriggerCondition_620
	case 623:
		return skillTriggerCondition_623
	case 626:
		return skillTriggerCondition_626
	case 629:
		return skillTriggerCondition_629
	case 632:
		return skillTriggerCondition_632
	case 635:
		return skillTriggerCondition_635
	case 643:
		return skillTriggerCondition_643
	case 644:
		return skillTriggerCondition_644
	case 645:
		return skillTriggerCondition_645
	case 646:
		return skillTriggerCondition_646
	case 647:
		return skillTriggerCondition_647
	case 648:
		return skillTriggerCondition_648
	case 713:
		return skillTriggerCondition_713
	case 718:
		return skillTriggerCondition_718
	case 719:
		return skillTriggerCondition_719
	case 723:
		return skillTriggerCondition_723
	case 724:
		return skillTriggerCondition_724
	case 725:
		return skillTriggerCondition_725
	case 726:
		return skillTriggerCondition_726
	case 727:
		return skillTriggerCondition_727
	case 732:
		return skillTriggerCondition_732
	case 733:
		return skillTriggerCondition_733
	case 738:
		return skillTriggerCondition_738
	case 739:
		return skillTriggerCondition_739
	case 744:
		return skillTriggerCondition_744
	case 745:
		return skillTriggerCondition_745
	case 750:
		return skillTriggerCondition_750
	case 751:
		return skillTriggerCondition_751
	case 752:
		return skillTriggerCondition_752
	case 754:
		return skillTriggerCondition_754
	case 758:
		return skillTriggerCondition_758
	case 761:
		return skillTriggerCondition_761
	case 775:
		return skillTriggerCondition_775
	case 777:
		return skillTriggerCondition_777
	case 779:
		return skillTriggerCondition_779
	case 795:
		return skillTriggerCondition_795
	case 797:
		return skillTriggerCondition_797
	case 799:
		return skillTriggerCondition_799
	case 801:
		return skillTriggerCondition_801
	case 803:
		return skillTriggerCondition_803
	case 808:
		return skillTriggerCondition_808
	case 809:
		return skillTriggerCondition_809
	case 814:
		return skillTriggerCondition_814
	case 815:
		return skillTriggerCondition_815
	case 820:
		return skillTriggerCondition_820
	case 821:
		return skillTriggerCondition_821
	case 826:
		return skillTriggerCondition_826
	case 827:
		return skillTriggerCondition_827
	case 832:
		return skillTriggerCondition_832
	case 833:
		return skillTriggerCondition_833
	case 836:
		return skillTriggerCondition_836
	case 839:
		return skillTriggerCondition_839
	case 842:
		return skillTriggerCondition_842
	case 843:
		return skillTriggerCondition_843
	case 845:
		return skillTriggerCondition_845
	case 847:
		return skillTriggerCondition_847
	case 849:
		return skillTriggerCondition_849
	case 851:
		return skillTriggerCondition_851
	case 859:
		return skillTriggerCondition_859
	case 861:
		return skillTriggerCondition_861
	case 865:
		return skillTriggerCondition_865
	case 868:
		return skillTriggerCondition_868
	case 951:
		return skillTriggerCondition_951
	case 952:
		return skillTriggerCondition_952
	case 957:
		return skillTriggerCondition_957
	case 958:
		return skillTriggerCondition_958
	case 963:
		return skillTriggerCondition_963
	case 964:
		return skillTriggerCondition_964
	case 969:
		return skillTriggerCondition_969
	case 970:
		return skillTriggerCondition_970
	case 975:
		return skillTriggerCondition_975
	case 976:
		return skillTriggerCondition_976
	case 983:
		return skillTriggerCondition_983
	case 988:
		return skillTriggerCondition_988
	case 996:
		return skillTriggerCondition_996
	case 1003:
		return skillTriggerCondition_1003
	case 1004:
		return skillTriggerCondition_1004
	case 1007:
		return skillTriggerCondition_1007
	case 1018:
		return skillTriggerCondition_1018
	case 1019:
		return skillTriggerCondition_1019
	case 1025:
		return skillTriggerCondition_1025
	case 1029:
		return skillTriggerCondition_1029
	case 1037:
		return skillTriggerCondition_1037
	case 1045:
		return skillTriggerCondition_1045
	case 1046:
		return skillTriggerCondition_1046
	case 1052:
		return skillTriggerCondition_1052
	case 1055:
		return skillTriggerCondition_1055
	case 1059:
		return skillTriggerCondition_1059
	case 1061:
		return skillTriggerCondition_1061
	case 1062:
		return skillTriggerCondition_1062
	case 1065:
		return skillTriggerCondition_1065
	case 1068:
		return skillTriggerCondition_1068
	case 1075:
		return skillTriggerCondition_1075
	case 1085:
		return skillTriggerCondition_1085
	case 1086:
		return skillTriggerCondition_1086
	case 1093:
		return skillTriggerCondition_1093
	case 1095:
		return skillTriggerCondition_1095
	case 1097:
		return skillTriggerCondition_1097
	case 1101:
		return skillTriggerCondition_1101
	case 1103:
		return skillTriggerCondition_1103
	case 1109:
		return skillTriggerCondition_1109
	case 1113:
		return skillTriggerCondition_1113
	case 1122:
		return skillTriggerCondition_1122
	case 1129:
		return skillTriggerCondition_1129
	case 1130:
		return skillTriggerCondition_1130
	case 1135:
		return skillTriggerCondition_1135
	case 1137:
		return skillTriggerCondition_1137
	case 1146:
		return skillTriggerCondition_1146
	case 1148:
		return skillTriggerCondition_1148
	case 1149:
		return skillTriggerCondition_1149
	case 1150:
		return skillTriggerCondition_1150
	case 1151:
		return skillTriggerCondition_1151
	case 1153:
		return skillTriggerCondition_1153
	case 1158:
		return skillTriggerCondition_1158
	case 1159:
		return skillTriggerCondition_1159
		
	default:
		return skillTriggerCondition_default
	}
	return nil
}


func skillTriggerCondition_1(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 1214, 640})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 1274, 640})
	}

	if f.Health <= 25000 {
		skill = append(skill, triggerSkill{3, 1214, 640})
	}

	if f.Health <= 35000 {
		skill = append(skill, triggerSkill{4, 1214, 640})
	}
	
	return skill
}

func skillTriggerCondition_2(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 39600 {
		skill = append(skill, triggerSkill{1, 30, 1760})
	}

	if f.Health <= 66000 {
		skill = append(skill, triggerSkill{2, 34, 1760})
	}

	if f.Health <= 105600 {
		skill = append(skill, triggerSkill{3, 30, 1760})
	}
	
	return skill
}

func skillTriggerCondition_3(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 36000 {
		skill = append(skill, triggerSkill{1, 44, 1720})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{2, 18, 1720})
	}

	if f.Health <= 96000 {
		skill = append(skill, triggerSkill{3, 43, 1720})
	}
	
	return skill
}

func skillTriggerCondition_4(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 22000 {
		skill = append(skill, triggerSkill{1, 1147, 30000})
	}

	if f.Health <= 44000 {
		skill = append(skill, triggerSkill{2, 30, 1680})
	}

	if f.Health <= 66000 {
		skill = append(skill, triggerSkill{3, 29, 1680})
	}

	if f.Health <= 88000 {
		skill = append(skill, triggerSkill{4, 28, 1680})
	}
	
	return skill
}

func skillTriggerCondition_5(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 29600 {
		skill = append(skill, triggerSkill{1, 22, 2080})
	}

	if f.Health <= 59200 {
		skill = append(skill, triggerSkill{2, 45, 2080})
	}

	if f.Health <= 88800 {
		skill = append(skill, triggerSkill{3, 45, 2080})
	}

	if f.Health <= 118400 {
		skill = append(skill, triggerSkill{4, 26, 2080})
	}
	
	return skill
}

func skillTriggerCondition_8(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 46500 {
		skill = append(skill, triggerSkill{1, 26, 2120})
	}

	if f.Health <= 77500 {
		skill = append(skill, triggerSkill{2, 26, 2120})
	}

	if f.Health <= 124000 {
		skill = append(skill, triggerSkill{3, 26, 2120})
	}
	
	return skill
}

func skillTriggerCondition_10(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 42000 {
		skill = append(skill, triggerSkill{1, 22, 2600})
	}

	if f.Health <= 84000 {
		skill = append(skill, triggerSkill{2, 45, 2600})
	}

	if f.Health <= 126000 {
		skill = append(skill, triggerSkill{3, 45, 2600})
	}

	if f.Health <= 168000 {
		skill = append(skill, triggerSkill{4, 26, 2600})
	}
	
	return skill
}

func skillTriggerCondition_11(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 2000 {
		skill = append(skill, triggerSkill{1, 1214, 150})
	}

	if f.Health <= 3000 {
		skill = append(skill, triggerSkill{2, 1214, 150})
	}

	if f.Health <= 4000 {
		skill = append(skill, triggerSkill{3, 1214, 150})
	}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{4, 1214, 150})
	}
	
	return skill
}

func skillTriggerCondition_14(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 2000 {
		skill = append(skill, triggerSkill{1, 1215, 100})
	}

	if f.Health <= 4000 {
		skill = append(skill, triggerSkill{2, 1215, 100})
	}

	if f.Health <= 6000 {
		skill = append(skill, triggerSkill{3, 1215, 100})
	}

	if f.Health <= 8000 {
		skill = append(skill, triggerSkill{4, 1215, 100})
	}
	
	return skill
}

func skillTriggerCondition_18(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 4000 {
		skill = append(skill, triggerSkill{1, 30, 200})
	}

	if f.Health <= 7000 {
		skill = append(skill, triggerSkill{2, 30, 200})
	}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{3, 1223, 200})
	}

	if f.Health <= 12000 {
		skill = append(skill, triggerSkill{4, 37, 200})
	}
	
	return skill
}

func skillTriggerCondition_19(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{1, 1213, 200})
	}

	if f.Health <= 9000 {
		skill = append(skill, triggerSkill{2, 1213, 200})
	}

	if f.Health <= 12000 {
		skill = append(skill, triggerSkill{3, 1213, 200})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{4, 1213, 200})
	}
	
	return skill
}

func skillTriggerCondition_23(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 6000 {
		skill = append(skill, triggerSkill{1, 40, 330})
	}

	if f.Health <= 8000 {
		skill = append(skill, triggerSkill{2, 108, 6000})
	}

	if f.Health <= 22000 {
		skill = append(skill, triggerSkill{3, 40, 330})
	}
	
	return skill
}

func skillTriggerCondition_25(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 22, 1640})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{2, 26, 1640})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{3, 45, 1640})
	}

	if f.Health <= 80000 {
		skill = append(skill, triggerSkill{4, 45, 1640})
	}
	
	return skill
}

func skillTriggerCondition_26(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 42000 {
		skill = append(skill, triggerSkill{1, 22, 2000})
	}

	if f.Health <= 70000 {
		skill = append(skill, triggerSkill{2, 30, 2000})
	}

	if f.Health <= 112000 {
		skill = append(skill, triggerSkill{3, 14, 2000})
	}
	
	return skill
}

func skillTriggerCondition_38(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{1, 45, 2440})
	}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{2, 22, 2440})
	}

	if f.Health <= 160000 {
		skill = append(skill, triggerSkill{3, 45, 2440})
	}
	
	return skill
}

func skillTriggerCondition_88(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 1000 {
		skill = append(skill, triggerSkill{1, 42, 50})
	}
	
	return skill
}

func skillTriggerCondition_89(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 1000 {
		skill = append(skill, triggerSkill{1, 42, 50})
	}
	
	return skill
}

func skillTriggerCondition_96(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 33000 {
		skill = append(skill, triggerSkill{1, 40, 2240})
	}

	if f.Health <= 66000 {
		skill = append(skill, triggerSkill{2, 40, 2240})
	}

	if f.Health <= 99000 {
		skill = append(skill, triggerSkill{3, 40, 2240})
	}

	if f.Health <= 132000 {
		skill = append(skill, triggerSkill{4, 40, 2240})
	}
	
	return skill
}

func skillTriggerCondition_97(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 8000 {
		skill = append(skill, triggerSkill{1, 1428, 200})
	}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{2, 1205, 200})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{3, 1428, 200})
	}
	
	return skill
}

func skillTriggerCondition_100(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 5850 {
		skill = append(skill, triggerSkill{1, 1473, 750})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{2, 40, 750})
	}

	if f.Health <= 19500 {
		skill = append(skill, triggerSkill{3, 108, 15000})
	}

	if f.Health <= 23400 {
		skill = append(skill, triggerSkill{4, 40, 750})
	}

	if f.Health <= 31200 {
		skill = append(skill, triggerSkill{5, 40, 750})
	}
	
	return skill
}

func skillTriggerCondition_101(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 6000 {
		skill = append(skill, triggerSkill{1, 108, 4000})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{2, 38, 300})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{3, 38, 300})
	}
	
	return skill
}

func skillTriggerCondition_103(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 42000 {
		skill = append(skill, triggerSkill{1, 40, 1800})
	}

	if f.Health <= 70000 {
		skill = append(skill, triggerSkill{2, 40, 1800})
	}

	if f.Health <= 112000 {
		skill = append(skill, triggerSkill{3, 40, 1800})
	}
	
	return skill
}

func skillTriggerCondition_104(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{1, 45, 300})
	}

	if f.Health <= 13000 {
		skill = append(skill, triggerSkill{2, 1472, 300})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{3, 22, 300})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{4, 45, 300})
	}
	
	return skill
}

func skillTriggerCondition_105(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 7000 {
		skill = append(skill, triggerSkill{1, 1215, 600})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{2, 1215, 600})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{3, 1215, 600})
	}

	if f.Health <= 22500 {
		skill = append(skill, triggerSkill{4, 1434, 600})
	}

	if f.Health <= 25000 {
		skill = append(skill, triggerSkill{5, 1215, 600})
	}
	
	return skill
}

func skillTriggerCondition_106(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 1213, 800})
	}

	if f.Health <= 12000 {
		skill = append(skill, triggerSkill{2, 1436, 800})
	}

	if f.Health <= 14000 {
		skill = append(skill, triggerSkill{3, 1213, 800})
	}

	if f.Health <= 23000 {
		skill = append(skill, triggerSkill{4, 1213, 800})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{5, 1213, 800})
	}
	
	return skill
}

func skillTriggerCondition_107(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{1, 40, 1800})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 40, 1800})
	}

	if f.Health <= 25000 {
		skill = append(skill, triggerSkill{3, 40, 1800})
	}
	
	return skill
}

func skillTriggerCondition_110(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{1, 1039, 8000})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{2, 1039, 8000})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{3, 1039, 8000})
	}
	
	return skill
}

func skillTriggerCondition_111(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 400 {
		skill = append(skill, triggerSkill{1, 108, 1000})
	}
	
	return skill
}

func skillTriggerCondition_112(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 500 {
		skill = append(skill, triggerSkill{1, 39, 20})
	}

	if f.Health <= 1500 {
		skill = append(skill, triggerSkill{2, 39, 20})
	}

	if f.Health <= 2400 {
		skill = append(skill, triggerSkill{3, 39, 20})
	}
	
	return skill
}

func skillTriggerCondition_113(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 40, 2200})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{2, 40, 2200})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{3, 40, 2200})
	}

	if f.Health <= 80000 {
		skill = append(skill, triggerSkill{4, 40, 2200})
	}
	
	return skill
}

func skillTriggerCondition_114(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 70000 {
		skill = append(skill, triggerSkill{1, 22, 5440})
	}

	if f.Health <= 140000 {
		skill = append(skill, triggerSkill{2, 45, 5440})
	}

	if f.Health <= 210000 {
		skill = append(skill, triggerSkill{3, 45, 5440})
	}

	if f.Health <= 280000 {
		skill = append(skill, triggerSkill{4, 26, 5440})
	}
	
	return skill
}

func skillTriggerCondition_115(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 2000 {
		skill = append(skill, triggerSkill{1, 1055, 120})
	}
	
	return skill
}

func skillTriggerCondition_151(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 1039, 560})
	}

	if f.Health <= 21000 {
		skill = append(skill, triggerSkill{2, 1039, 560})
	}

	if f.Health <= 32000 {
		skill = append(skill, triggerSkill{3, 1039, 560})
	}
	
	return skill
}

func skillTriggerCondition_181(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 7800 {
		skill = append(skill, triggerSkill{1, 22, 600})
	}

	if f.Health <= 15600 {
		skill = append(skill, triggerSkill{2, 44, 600})
	}

	if f.Health <= 23400 {
		skill = append(skill, triggerSkill{3, 38, 600})
	}

	if f.Health <= 31200 {
		skill = append(skill, triggerSkill{4, 22, 600})
	}
	
	return skill
}

func skillTriggerCondition_182(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 14400 {
		skill = append(skill, triggerSkill{1, 1049, 500})
	}

	if f.Health <= 24000 {
		skill = append(skill, triggerSkill{2, 1050, 25000})
	}

	if f.Health <= 38400 {
		skill = append(skill, triggerSkill{3, 1049, 500})
	}
	
	return skill
}

func skillTriggerCondition_183(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 13200 {
		skill = append(skill, triggerSkill{1, 30, 840})
	}

	if f.Health <= 22000 {
		skill = append(skill, triggerSkill{2, 10, 840})
	}

	if f.Health <= 35200 {
		skill = append(skill, triggerSkill{3, 30, 840})
	}
	
	return skill
}

func skillTriggerCondition_184(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10600 {
		skill = append(skill, triggerSkill{1, 40, 1080})
	}

	if f.Health <= 21200 {
		skill = append(skill, triggerSkill{2, 40, 1080})
	}

	if f.Health <= 31800 {
		skill = append(skill, triggerSkill{3, 22, 1080})
	}

	if f.Health <= 42400 {
		skill = append(skill, triggerSkill{4, 18, 1080})
	}
	
	return skill
}

func skillTriggerCondition_185(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 13400 {
		skill = append(skill, triggerSkill{1, 18, 1160})
	}

	if f.Health <= 26800 {
		skill = append(skill, triggerSkill{2, 18, 1160})
	}

	if f.Health <= 40200 {
		skill = append(skill, triggerSkill{3, 16, 1160})
	}

	if f.Health <= 53600 {
		skill = append(skill, triggerSkill{4, 17, 1160})
	}
	
	return skill
}

func skillTriggerCondition_186(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 23400 {
		skill = append(skill, triggerSkill{1, 1051, 1280})
	}

	if f.Health <= 39000 {
		skill = append(skill, triggerSkill{2, 1051, 1280})
	}

	if f.Health <= 62400 {
		skill = append(skill, triggerSkill{3, 1051, 1280})
	}
	
	return skill
}

func skillTriggerCondition_187(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 26400 {
		skill = append(skill, triggerSkill{1, 22, 1360})
	}

	if f.Health <= 44000 {
		skill = append(skill, triggerSkill{2, 26, 1360})
	}

	if f.Health <= 70400 {
		skill = append(skill, triggerSkill{3, 30, 1360})
	}
	
	return skill
}

func skillTriggerCondition_188(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 23500 {
		skill = append(skill, triggerSkill{1, 40, 1560})
	}

	if f.Health <= 37600 {
		skill = append(skill, triggerSkill{2, 39, 1560})
	}

	if f.Health <= 51700 {
		skill = append(skill, triggerSkill{3, 38, 1560})
	}

	if f.Health <= 65800 {
		skill = append(skill, triggerSkill{4, 42, 1560})
	}

	if f.Health <= 79900 {
		skill = append(skill, triggerSkill{5, 45, 1560})
	}
	
	return skill
}

func skillTriggerCondition_199(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 8000 {
		skill = append(skill, triggerSkill{1, 1214, 400})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{2, 1214, 400})
	}

	if f.Health <= 22000 {
		skill = append(skill, triggerSkill{3, 1214, 400})
	}
	
	return skill
}

func skillTriggerCondition_204(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 126000 {
		skill = append(skill, triggerSkill{1, 1432, 7000})
	}

	if f.Health <= 210000 {
		skill = append(skill, triggerSkill{2, 40, 7000})
	}

	if f.Health <= 336000 {
		skill = append(skill, triggerSkill{3, 1432, 7000})
	}
	
	return skill
}

func skillTriggerCondition_205(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 31500 {
		skill = append(skill, triggerSkill{1, 18, 1600})
	}

	if f.Health <= 52500 {
		skill = append(skill, triggerSkill{2, 18, 1600})
	}

	if f.Health <= 84000 {
		skill = append(skill, triggerSkill{3, 18, 1600})
	}
	
	return skill
}

func skillTriggerCondition_206(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 1214, 4000})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{2, 1214, 4000})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{3, 1214, 4000})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{4, 1214, 4000})
	}
	
	return skill
}

func skillTriggerCondition_208(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{1, 1214, 600})
	}

	if f.Health <= 8000 {
		skill = append(skill, triggerSkill{2, 1214, 600})
	}

	if f.Health <= 12000 {
		skill = append(skill, triggerSkill{3, 1214, 600})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{4, 1214, 600})
	}
	
	return skill
}

func skillTriggerCondition_210(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{1, 40, 8000})
	}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{2, 40, 8000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{3, 40, 8000})
	}
	
	return skill
}

func skillTriggerCondition_214(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 1039, 1500})
	}

	if f.Health <= 35000 {
		skill = append(skill, triggerSkill{2, 1039, 1500})
	}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{3, 1039, 1500})
	}
	
	return skill
}

func skillTriggerCondition_215(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 12500 {
		skill = append(skill, triggerSkill{1, 22, 1350})
	}

	if f.Health <= 16500 {
		skill = append(skill, triggerSkill{2, 1438, 1350})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{3, 22, 1350})
	}

	if f.Health <= 27500 {
		skill = append(skill, triggerSkill{4, 44, 1350})
	}

	if f.Health <= 35000 {
		skill = append(skill, triggerSkill{5, 38, 1350})
	}

	if f.Health <= 42500 {
		skill = append(skill, triggerSkill{6, 22, 1350})
	}
	
	return skill
}

func skillTriggerCondition_216(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 12000 {
		skill = append(skill, triggerSkill{1, 30, 1550})
	}

	if f.Health <= 24000 {
		skill = append(skill, triggerSkill{2, 1439, 1550})
	}

	if f.Health <= 36000 {
		skill = append(skill, triggerSkill{3, 9, 1550})
	}

	if f.Health <= 48000 {
		skill = append(skill, triggerSkill{4, 30, 1550})
	}
	
	return skill
}

func skillTriggerCondition_217(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 22000 {
		skill = append(skill, triggerSkill{1, 18, 1900})
	}

	if f.Health <= 35200 {
		skill = append(skill, triggerSkill{2, 18, 1900})
	}

	if f.Health <= 48400 {
		skill = append(skill, triggerSkill{3, 1441, 1900})
	}

	if f.Health <= 61600 {
		skill = append(skill, triggerSkill{4, 16, 1900})
	}

	if f.Health <= 74800 {
		skill = append(skill, triggerSkill{5, 17, 1900})
	}
	
	return skill
}

func skillTriggerCondition_218(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 22, 2250})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{2, 26, 2250})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{3, 1442, 2250})
	}

	if f.Health <= 80000 {
		skill = append(skill, triggerSkill{4, 30, 2250})
	}
	
	return skill
}

func skillTriggerCondition_219(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{1, 22, 2400})
	}

	if f.Health <= 39600 {
		skill = append(skill, triggerSkill{2, 1444, 2400})
	}

	if f.Health <= 48000 {
		skill = append(skill, triggerSkill{3, 45, 2400})
	}

	if f.Health <= 66000 {
		skill = append(skill, triggerSkill{4, 45, 2400})
	}

	if f.Health <= 84000 {
		skill = append(skill, triggerSkill{5, 1443, 2400})
	}

	if f.Health <= 102000 {
		skill = append(skill, triggerSkill{6, 26, 2400})
	}
	
	return skill
}

func skillTriggerCondition_220(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 18000 {
		skill = append(skill, triggerSkill{1, 1214, 640})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{2, 1214, 640})
	}

	if f.Health <= 48000 {
		skill = append(skill, triggerSkill{3, 1214, 640})
	}
	
	return skill
}

func skillTriggerCondition_221(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 34000 {
		skill = append(skill, triggerSkill{1, 1215, 760})
	}

	if f.Health <= 51000 {
		skill = append(skill, triggerSkill{2, 1215, 760})
	}

	if f.Health <= 68000 {
		skill = append(skill, triggerSkill{3, 1215, 760})
	}
	
	return skill
}

func skillTriggerCondition_222(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 23000 {
		skill = append(skill, triggerSkill{1, 1213, 1000})
	}

	if f.Health <= 46000 {
		skill = append(skill, triggerSkill{2, 1070, 1000})
	}

	if f.Health <= 69000 {
		skill = append(skill, triggerSkill{3, 1213, 1000})
	}

	if f.Health <= 92000 {
		skill = append(skill, triggerSkill{4, 1070, 1000})
	}
	
	return skill
}

func skillTriggerCondition_223(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 31000 {
		skill = append(skill, triggerSkill{1, 108, 44000})
	}

	if f.Health <= 62000 {
		skill = append(skill, triggerSkill{2, 40, 1600})
	}

	if f.Health <= 93000 {
		skill = append(skill, triggerSkill{3, 40, 1600})
	}

	if f.Health <= 124000 {
		skill = append(skill, triggerSkill{4, 40, 1600})
	}
	
	return skill
}

func skillTriggerCondition_224(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 58500 {
		skill = append(skill, triggerSkill{1, 1039, 1760})
	}

	if f.Health <= 97500 {
		skill = append(skill, triggerSkill{2, 1039, 1760})
	}

	if f.Health <= 156000 {
		skill = append(skill, triggerSkill{3, 1039, 1760})
	}
	
	return skill
}

func skillTriggerCondition_225(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 75000 {
		skill = append(skill, triggerSkill{1, 1508, 2240})
	}

	if f.Health <= 125000 {
		skill = append(skill, triggerSkill{2, 1740, 2240})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{3, 1508, 2240})
	}
	
	return skill
}

func skillTriggerCondition_226(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 90000 {
		skill = append(skill, triggerSkill{1, 1791, 2700})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{2, 1791, 2700})
	}

	if f.Health <= 240000 {
		skill = append(skill, triggerSkill{3, 1791, 2700})
	}
	
	return skill
}

func skillTriggerCondition_227(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 69000 {
		skill = append(skill, triggerSkill{1, 1070, 3200})
	}

	if f.Health <= 138000 {
		skill = append(skill, triggerSkill{2, 1070, 3200})
	}

	if f.Health <= 207000 {
		skill = append(skill, triggerSkill{3, 22, 3200})
	}

	if f.Health <= 276000 {
		skill = append(skill, triggerSkill{4, 18, 3200})
	}
	
	return skill
}

func skillTriggerCondition_272(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 44, 2500})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{2, 18, 2500})
	}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{3, 1445, 2500})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{4, 43, 2500})
	}
	
	return skill
}

func skillTriggerCondition_273(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 29600 {
		skill = append(skill, triggerSkill{1, 30, 2900})
	}

	if f.Health <= 59200 {
		skill = append(skill, triggerSkill{2, 34, 2900})
	}

	if f.Health <= 88800 {
		skill = append(skill, triggerSkill{3, 1446, 2900})
	}

	if f.Health <= 118400 {
		skill = append(skill, triggerSkill{4, 30, 2900})
	}
	
	return skill
}

func skillTriggerCondition_274(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 41250 {
		skill = append(skill, triggerSkill{1, 22, 3100})
	}

	if f.Health <= 66000 {
		skill = append(skill, triggerSkill{2, 1448, 3100})
	}

	if f.Health <= 90750 {
		skill = append(skill, triggerSkill{3, 30, 3100})
	}

	if f.Health <= 115500 {
		skill = append(skill, triggerSkill{4, 14, 3100})
	}

	if f.Health <= 140250 {
		skill = append(skill, triggerSkill{5, 1447, 3100})
	}
	
	return skill
}

func skillTriggerCondition_275(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 52500 {
		skill = append(skill, triggerSkill{1, 26, 3550})
	}

	if f.Health <= 84000 {
		skill = append(skill, triggerSkill{2, 1449, 3550})
	}

	if f.Health <= 115500 {
		skill = append(skill, triggerSkill{3, 26, 3550})
	}

	if f.Health <= 147000 {
		skill = append(skill, triggerSkill{4, 1449, 3550})
	}

	if f.Health <= 178500 {
		skill = append(skill, triggerSkill{5, 26, 3550})
	}
	
	return skill
}

func skillTriggerCondition_276(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 48000 {
		skill = append(skill, triggerSkill{1, 45, 3200})
	}

	if f.Health <= 96000 {
		skill = append(skill, triggerSkill{2, 1450, 3200})
	}

	if f.Health <= 144000 {
		skill = append(skill, triggerSkill{3, 22, 3200})
	}

	if f.Health <= 192000 {
		skill = append(skill, triggerSkill{4, 45, 3200})
	}
	
	return skill
}

func skillTriggerCondition_281(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 1147, 20000})
	}

	if f.Health <= 54000 {
		skill = append(skill, triggerSkill{2, 30, 2600})
	}

	if f.Health <= 90000 {
		skill = append(skill, triggerSkill{3, 29, 2600})
	}

	if f.Health <= 144000 {
		skill = append(skill, triggerSkill{4, 28, 2600})
	}
	
	return skill
}

func skillTriggerCondition_282(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 135000 {
		skill = append(skill, triggerSkill{1, 1218, 7800})
	}

	if f.Health <= 225000 {
		skill = append(skill, triggerSkill{2, 1218, 7800})
	}

	if f.Health <= 360000 {
		skill = append(skill, triggerSkill{3, 1218, 7800})
	}
	
	return skill
}

func skillTriggerCondition_289(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{1, 40, 2500})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 40, 2500})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{3, 40, 2500})
	}
	
	return skill
}

func skillTriggerCondition_290(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{1, 40, 7000})
	}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{2, 40, 7000})
	}

	if f.Health <= 70000 {
		skill = append(skill, triggerSkill{3, 40, 7000})
	}
	
	return skill
}

func skillTriggerCondition_291(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{1, 1039, 2500})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{2, 1039, 2500})
	}

	if f.Health <= 90000 {
		skill = append(skill, triggerSkill{3, 1039, 2500})
	}
	
	return skill
}

func skillTriggerCondition_292(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{1, 1039, 7000})
	}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{2, 1039, 7000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{3, 1039, 7000})
	}
	
	return skill
}

func skillTriggerCondition_308(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{1, 1214, 1500})
	}

	if f.Health <= 15000 {
		skill = append(skill, triggerSkill{2, 1214, 1500})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{3, 1214, 1500})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{4, 1214, 1500})
	}
	
	return skill
}

func skillTriggerCondition_310(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 1214, 2500})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 1214, 2500})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{3, 1214, 2500})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{4, 1214, 2500})
	}
	
	return skill
}

func skillTriggerCondition_360(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 69000 {
		skill = append(skill, triggerSkill{1, 26, 2800})
	}

	if f.Health <= 115000 {
		skill = append(skill, triggerSkill{2, 22, 2800})
	}

	if f.Health <= 184000 {
		skill = append(skill, triggerSkill{3, 18, 2800})
	}
	
	return skill
}

func skillTriggerCondition_361(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 72000 {
		skill = append(skill, triggerSkill{1, 22, 2920})
	}

	if f.Health <= 120000 {
		skill = append(skill, triggerSkill{2, 22, 2920})
	}

	if f.Health <= 192000 {
		skill = append(skill, triggerSkill{3, 1047, 100000})
	}
	
	return skill
}

func skillTriggerCondition_362(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 76500 {
		skill = append(skill, triggerSkill{1, 30, 3160})
	}

	if f.Health <= 127500 {
		skill = append(skill, triggerSkill{2, 22, 3160})
	}

	if f.Health <= 204000 {
		skill = append(skill, triggerSkill{3, 30, 3160})
	}
	
	return skill
}

func skillTriggerCondition_363(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 79500 {
		skill = append(skill, triggerSkill{1, 43, 3320})
	}

	if f.Health <= 132500 {
		skill = append(skill, triggerSkill{2, 44, 3320})
	}

	if f.Health <= 212000 {
		skill = append(skill, triggerSkill{3, 43, 3320})
	}
	
	return skill
}

func skillTriggerCondition_364(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 93000 {
		skill = append(skill, triggerSkill{1, 43, 4360})
	}

	if f.Health <= 155000 {
		skill = append(skill, triggerSkill{2, 43, 4360})
	}

	if f.Health <= 248000 {
		skill = append(skill, triggerSkill{3, 43, 4360})
	}
	
	return skill
}

func skillTriggerCondition_365(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 96000 {
		skill = append(skill, triggerSkill{1, 34, 4800})
	}

	if f.Health <= 160000 {
		skill = append(skill, triggerSkill{2, 30, 4800})
	}

	if f.Health <= 256000 {
		skill = append(skill, triggerSkill{3, 34, 4800})
	}
	
	return skill
}

func skillTriggerCondition_366(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 81000 {
		skill = append(skill, triggerSkill{1, 26, 3720})
	}

	if f.Health <= 135000 {
		skill = append(skill, triggerSkill{2, 22, 3720})
	}

	if f.Health <= 216000 {
		skill = append(skill, triggerSkill{3, 38, 3720})
	}
	
	return skill
}

func skillTriggerCondition_367(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 84000 {
		skill = append(skill, triggerSkill{1, 18, 4000})
	}

	if f.Health <= 140000 {
		skill = append(skill, triggerSkill{2, 26, 4000})
	}

	if f.Health <= 224000 {
		skill = append(skill, triggerSkill{3, 14, 4000})
	}
	
	return skill
}

func skillTriggerCondition_368(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 102000 {
		skill = append(skill, triggerSkill{1, 39, 5160})
	}

	if f.Health <= 170000 {
		skill = append(skill, triggerSkill{2, 30, 5160})
	}

	if f.Health <= 272000 {
		skill = append(skill, triggerSkill{3, 9, 5160})
	}
	
	return skill
}

func skillTriggerCondition_369(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 53000 {
		skill = append(skill, triggerSkill{1, 22, 4550})
	}

	if f.Health <= 106000 {
		skill = append(skill, triggerSkill{2, 1451, 4550})
	}

	if f.Health <= 159000 {
		skill = append(skill, triggerSkill{3, 22, 4550})
	}

	if f.Health <= 212000 {
		skill = append(skill, triggerSkill{4, 22, 4550})
	}
	
	return skill
}

func skillTriggerCondition_370(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 56000 {
		skill = append(skill, triggerSkill{1, 1484, 5500})
	}

	if f.Health <= 112000 {
		skill = append(skill, triggerSkill{2, 1051, 5500})
	}

	if f.Health <= 168000 {
		skill = append(skill, triggerSkill{3, 1484, 5500})
	}

	if f.Health <= 224000 {
		skill = append(skill, triggerSkill{4, 1051, 5500})
	}
	
	return skill
}

func skillTriggerCondition_371(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 80000 {
		skill = append(skill, triggerSkill{1, 18, 6600})
	}

	if f.Health <= 128000 {
		skill = append(skill, triggerSkill{2, 1453, 6600})
	}

	if f.Health <= 176000 {
		skill = append(skill, triggerSkill{3, 26, 6600})
	}

	if f.Health <= 224000 {
		skill = append(skill, triggerSkill{4, 1453, 6600})
	}

	if f.Health <= 272000 {
		skill = append(skill, triggerSkill{5, 14, 6600})
	}
	
	return skill
}

func skillTriggerCondition_372(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 70000 {
		skill = append(skill, triggerSkill{1, 1483, 7500})
	}

	if f.Health <= 140000 {
		skill = append(skill, triggerSkill{2, 1483, 7500})
	}

	if f.Health <= 210000 {
		skill = append(skill, triggerSkill{3, 1454, 7500})
	}

	if f.Health <= 280000 {
		skill = append(skill, triggerSkill{4, 1483, 7500})
	}
	
	return skill
}

func skillTriggerCondition_373(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 95000 {
		skill = append(skill, triggerSkill{1, 22, 7500})
	}

	if f.Health <= 152000 {
		skill = append(skill, triggerSkill{2, 45, 7500})
	}

	if f.Health <= 209000 {
		skill = append(skill, triggerSkill{3, 1455, 7500})
	}

	if f.Health <= 266000 {
		skill = append(skill, triggerSkill{4, 45, 7500})
	}

	if f.Health <= 323000 {
		skill = append(skill, triggerSkill{5, 26, 7500})
	}
	
	return skill
}

func skillTriggerCondition_380(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 18, 1800})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 1210, 1800})
	}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{3, 22, 1800})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{4, 1210, 1800})
	}
	
	return skill
}

func skillTriggerCondition_387(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 10000 {
		skill = append(skill, triggerSkill{1, 1147, 20000})
	}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{2, 30, 3000})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{3, 30, 3000})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{4, 1211, 3000})
	}
	
	return skill
}

func skillTriggerCondition_394(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 20000 {
		skill = append(skill, triggerSkill{1, 22, 4000})
	}

	if f.Health <= 40000 {
		skill = append(skill, triggerSkill{2, 1212, 4000})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{3, 22, 4000})
	}

	if f.Health <= 90000 {
		skill = append(skill, triggerSkill{4, 1212, 4000})
	}
	
	return skill
}

func skillTriggerCondition_395(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1191, 1})
	}
	
	return skill
}

func skillTriggerCondition_396(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1209, 0})
	}
	
	return skill
}

func skillTriggerCondition_397(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1205, 0})
	}
	
	return skill
}

func skillTriggerCondition_398(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1206, 0})
	}
	
	return skill
}

func skillTriggerCondition_399(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1207, 0})
	}
	
	return skill
}

func skillTriggerCondition_400(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1208, 0})
	}
	
	return skill
}

func skillTriggerCondition_401(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1210, 0})
	}
	
	return skill
}

func skillTriggerCondition_402(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1211, 0})
	}
	
	return skill
}

func skillTriggerCondition_403(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99999 {
		skill = append(skill, triggerSkill{1, 1212, 0})
	}
	
	return skill
}

func skillTriggerCondition_412(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{1, 26, 8000})
	}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{2, 1219, 8000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{3, 26, 8000})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{4, 1219, 8000})
	}
	
	return skill
}

func skillTriggerCondition_419(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 80000 {
		skill = append(skill, triggerSkill{1, 40, 10000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{2, 40, 10000})
	}

	if f.Health <= 250000 {
		skill = append(skill, triggerSkill{3, 40, 10000})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{4, 1220, 10000})
	}
	
	return skill
}

func skillTriggerCondition_426(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{1, 1218, 13000})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{2, 1221, 13000})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{3, 1218, 13000})
	}

	if f.Health <= 400000 {
		skill = append(skill, triggerSkill{4, 1221, 13000})
	}
	
	return skill
}

func skillTriggerCondition_433(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 100000 {
		skill = append(skill, triggerSkill{1, 22, 20000})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{2, 22, 20000})
	}

	if f.Health <= 500000 {
		skill = append(skill, triggerSkill{3, 1047, 500000})
	}

	if f.Health <= 800000 {
		skill = append(skill, triggerSkill{4, 1222, 20000})
	}
	
	return skill
}

func skillTriggerCondition_437(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 78000 {
		skill = append(skill, triggerSkill{1, 1512, 3700})
	}

	if f.Health <= 156000 {
		skill = append(skill, triggerSkill{2, 1512, 3700})
	}

	if f.Health <= 234000 {
		skill = append(skill, triggerSkill{3, 18, 3700})
	}

	if f.Health <= 312000 {
		skill = append(skill, triggerSkill{4, 1070, 3700})
	}
	
	return skill
}

func skillTriggerCondition_438(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 135000 {
		skill = append(skill, triggerSkill{1, 1711, 4700})
	}

	if f.Health <= 225000 {
		skill = append(skill, triggerSkill{2, 1711, 4700})
	}

	if f.Health <= 360000 {
		skill = append(skill, triggerSkill{3, 1711, 4700})
	}
	
	return skill
}

func skillTriggerCondition_439(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 98000 {
		skill = append(skill, triggerSkill{1, 1147, 50000})
	}

	if f.Health <= 196000 {
		skill = append(skill, triggerSkill{2, 1519, 4600})
	}

	if f.Health <= 294000 {
		skill = append(skill, triggerSkill{3, 1147, 50000})
	}

	if f.Health <= 392000 {
		skill = append(skill, triggerSkill{4, 1519, 4600})
	}
	
	return skill
}

func skillTriggerCondition_440(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 162000 {
		skill = append(skill, triggerSkill{1, 1218, 4900})
	}

	if f.Health <= 270000 {
		skill = append(skill, triggerSkill{2, 1218, 4900})
	}

	if f.Health <= 432000 {
		skill = append(skill, triggerSkill{3, 1218, 4900})
	}
	
	return skill
}

func skillTriggerCondition_441(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 180000 {
		skill = append(skill, triggerSkill{1, 1606, 5420})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{2, 1606, 5420})
	}

	if f.Health <= 480000 {
		skill = append(skill, triggerSkill{3, 1606, 5420})
	}
	
	return skill
}

func skillTriggerCondition_442(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 198000 {
		skill = append(skill, triggerSkill{1, 1503, 6000})
	}

	if f.Health <= 330000 {
		skill = append(skill, triggerSkill{2, 1275, 6000})
	}

	if f.Health <= 528000 {
		skill = append(skill, triggerSkill{3, 1275, 6000})
	}
	
	return skill
}

func skillTriggerCondition_443(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 216000 {
		skill = append(skill, triggerSkill{1, 45, 6620})
	}

	if f.Health <= 360000 {
		skill = append(skill, triggerSkill{2, 1731, 6620})
	}

	if f.Health <= 576000 {
		skill = append(skill, triggerSkill{3, 1747, 6620})
	}
	
	return skill
}

func skillTriggerCondition_500(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 111000 {
		skill = append(skill, triggerSkill{1, 1218, 5960})
	}

	if f.Health <= 185000 {
		skill = append(skill, triggerSkill{2, 1218, 5960})
	}

	if f.Health <= 296000 {
		skill = append(skill, triggerSkill{3, 1218, 5960})
	}
	
	return skill
}

func skillTriggerCondition_501(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 114000 {
		skill = append(skill, triggerSkill{1, 1218, 6040})
	}

	if f.Health <= 190000 {
		skill = append(skill, triggerSkill{2, 1218, 6040})
	}

	if f.Health <= 304000 {
		skill = append(skill, triggerSkill{3, 1218, 6040})
	}
	
	return skill
}

func skillTriggerCondition_510(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 120000 {
		skill = append(skill, triggerSkill{1, 45, 6160})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{2, 45, 6160})
	}

	if f.Health <= 320000 {
		skill = append(skill, triggerSkill{3, 45, 6160})
	}
	
	return skill
}

func skillTriggerCondition_511(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 123000 {
		skill = append(skill, triggerSkill{1, 26, 6240})
	}

	if f.Health <= 205000 {
		skill = append(skill, triggerSkill{2, 22, 6240})
	}

	if f.Health <= 328000 {
		skill = append(skill, triggerSkill{3, 22, 6240})
	}
	
	return skill
}

func skillTriggerCondition_520(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 129000 {
		skill = append(skill, triggerSkill{1, 1223, 6360})
	}

	if f.Health <= 215000 {
		skill = append(skill, triggerSkill{2, 1223, 6360})
	}

	if f.Health <= 344000 {
		skill = append(skill, triggerSkill{3, 1223, 6360})
	}
	
	return skill
}

func skillTriggerCondition_521(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 88000 {
		skill = append(skill, triggerSkill{1, 22, 6440})
	}

	if f.Health <= 176000 {
		skill = append(skill, triggerSkill{2, 30, 6440})
	}

	if f.Health <= 264000 {
		skill = append(skill, triggerSkill{3, 26, 6440})
	}

	if f.Health <= 352000 {
		skill = append(skill, triggerSkill{4, 18, 6440})
	}
	
	return skill
}

func skillTriggerCondition_530(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 138000 {
		skill = append(skill, triggerSkill{1, 26, 6520})
	}

	if f.Health <= 230000 {
		skill = append(skill, triggerSkill{2, 26, 6520})
	}

	if f.Health <= 368000 {
		skill = append(skill, triggerSkill{3, 26, 6520})
	}
	
	return skill
}

func skillTriggerCondition_531(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 141000 {
		skill = append(skill, triggerSkill{1, 45, 6600})
	}

	if f.Health <= 235000 {
		skill = append(skill, triggerSkill{2, 45, 6600})
	}

	if f.Health <= 376000 {
		skill = append(skill, triggerSkill{3, 45, 6600})
	}
	
	return skill
}

func skillTriggerCondition_540(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 147000 {
		skill = append(skill, triggerSkill{1, 30, 7000})
	}

	if f.Health <= 245000 {
		skill = append(skill, triggerSkill{2, 38, 7000})
	}

	if f.Health <= 392000 {
		skill = append(skill, triggerSkill{3, 30, 7000})
	}
	
	return skill
}

func skillTriggerCondition_541(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{1, 18, 7200})
	}

	if f.Health <= 250000 {
		skill = append(skill, triggerSkill{2, 18, 7200})
	}

	if f.Health <= 400000 {
		skill = append(skill, triggerSkill{3, 18, 7200})
	}
	
	return skill
}

func skillTriggerCondition_542(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 102500 {
		skill = append(skill, triggerSkill{1, 1218, 8550})
	}

	if f.Health <= 164000 {
		skill = append(skill, triggerSkill{2, 1457, 8550})
	}

	if f.Health <= 225500 {
		skill = append(skill, triggerSkill{3, 1218, 8550})
	}

	if f.Health <= 287000 {
		skill = append(skill, triggerSkill{4, 1456, 8550})
	}

	if f.Health <= 348500 {
		skill = append(skill, triggerSkill{5, 1218, 8550})
	}
	
	return skill
}

func skillTriggerCondition_543(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 88000 {
		skill = append(skill, triggerSkill{1, 1490, 8850})
	}

	if f.Health <= 176000 {
		skill = append(skill, triggerSkill{2, 14, 8850})
	}

	if f.Health <= 264000 {
		skill = append(skill, triggerSkill{3, 14, 8850})
	}

	if f.Health <= 352000 {
		skill = append(skill, triggerSkill{4, 14, 8850})
	}
	
	return skill
}

func skillTriggerCondition_544(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 117500 {
		skill = append(skill, triggerSkill{1, 22, 9150})
	}

	if f.Health <= 188000 {
		skill = append(skill, triggerSkill{2, 30, 9150})
	}

	if f.Health <= 329000 {
		skill = append(skill, triggerSkill{3, 26, 9150})
	}

	if f.Health <= 399500 {
		skill = append(skill, triggerSkill{4, 18, 9150})
	}
	
	return skill
}

func skillTriggerCondition_545(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{1, 26, 9900})
	}

	if f.Health <= 275000 {
		skill = append(skill, triggerSkill{2, 26, 9900})
	}

	if f.Health <= 350000 {
		skill = append(skill, triggerSkill{3, 1461, 9900})
	}

	if f.Health <= 425000 {
		skill = append(skill, triggerSkill{4, 26, 9900})
	}
	
	return skill
}

func skillTriggerCondition_546(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 196000 {
		skill = append(skill, triggerSkill{1, 1486, 10450})
	}

	if f.Health <= 252000 {
		skill = append(skill, triggerSkill{2, 1464, 10450})
	}

	if f.Health <= 308000 {
		skill = append(skill, triggerSkill{3, 38, 10450})
	}

	if f.Health <= 392000 {
		skill = append(skill, triggerSkill{4, 30, 10450})
	}

	if f.Health <= 476000 {
		skill = append(skill, triggerSkill{5, 1463, 10450})
	}
	
	return skill
}

func skillTriggerCondition_573(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 156000 {
		skill = append(skill, triggerSkill{1, 1147, 100000})
	}

	if f.Health <= 260000 {
		skill = append(skill, triggerSkill{2, 1047, 100000})
	}

	if f.Health <= 416000 {
		skill = append(skill, triggerSkill{3, 1317, 7120})
	}
	
	return skill
}

func skillTriggerCondition_574(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 159000 {
		skill = append(skill, triggerSkill{1, 22, 7280})
	}

	if f.Health <= 265000 {
		skill = append(skill, triggerSkill{2, 22, 7280})
	}

	if f.Health <= 424000 {
		skill = append(skill, triggerSkill{3, 22, 7280})
	}
	
	return skill
}

func skillTriggerCondition_583(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 110000 {
		skill = append(skill, triggerSkill{1, 1218, 7440})
	}

	if f.Health <= 220000 {
		skill = append(skill, triggerSkill{2, 18, 7440})
	}

	if f.Health <= 330000 {
		skill = append(skill, triggerSkill{3, 1218, 7440})
	}

	if f.Health <= 440000 {
		skill = append(skill, triggerSkill{4, 18, 7440})
	}
	
	return skill
}

func skillTriggerCondition_584(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 112000 {
		skill = append(skill, triggerSkill{1, 22, 7600})
	}

	if f.Health <= 224000 {
		skill = append(skill, triggerSkill{2, 45, 7600})
	}

	if f.Health <= 336000 {
		skill = append(skill, triggerSkill{3, 45, 7600})
	}

	if f.Health <= 448000 {
		skill = append(skill, triggerSkill{4, 26, 7600})
	}
	
	return skill
}

func skillTriggerCondition_593(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 116000 {
		skill = append(skill, triggerSkill{1, 40, 7800})
	}

	if f.Health <= 232000 {
		skill = append(skill, triggerSkill{2, 1223, 7800})
	}

	if f.Health <= 348000 {
		skill = append(skill, triggerSkill{3, 40, 7800})
	}

	if f.Health <= 464000 {
		skill = append(skill, triggerSkill{4, 1223, 7800})
	}
	
	return skill
}

func skillTriggerCondition_594(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 118000 {
		skill = append(skill, triggerSkill{1, 40, 8000})
	}

	if f.Health <= 236000 {
		skill = append(skill, triggerSkill{2, 40, 8000})
	}

	if f.Health <= 354000 {
		skill = append(skill, triggerSkill{3, 18, 8000})
	}

	if f.Health <= 472000 {
		skill = append(skill, triggerSkill{4, 40, 8000})
	}
	
	return skill
}

func skillTriggerCondition_603(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 183000 {
		skill = append(skill, triggerSkill{1, 30, 8200})
	}

	if f.Health <= 305000 {
		skill = append(skill, triggerSkill{2, 30, 8200})
	}

	if f.Health <= 488000 {
		skill = append(skill, triggerSkill{3, 30, 8200})
	}
	
	return skill
}

func skillTriggerCondition_604(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 124000 {
		skill = append(skill, triggerSkill{1, 1218, 9600})
	}

	if f.Health <= 248000 {
		skill = append(skill, triggerSkill{2, 1147, 100000})
	}

	if f.Health <= 372000 {
		skill = append(skill, triggerSkill{3, 1218, 9600})
	}

	if f.Health <= 496000 {
		skill = append(skill, triggerSkill{4, 1147, 100000})
	}
	
	return skill
}

func skillTriggerCondition_613(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 128000 {
		skill = append(skill, triggerSkill{1, 30, 8600})
	}

	if f.Health <= 256000 {
		skill = append(skill, triggerSkill{2, 18, 8600})
	}

	if f.Health <= 384000 {
		skill = append(skill, triggerSkill{3, 1223, 8600})
	}

	if f.Health <= 512000 {
		skill = append(skill, triggerSkill{4, 30, 8600})
	}
	
	return skill
}

func skillTriggerCondition_614(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 195000 {
		skill = append(skill, triggerSkill{1, 26, 8800})
	}

	if f.Health <= 325000 {
		skill = append(skill, triggerSkill{2, 18, 8800})
	}

	if f.Health <= 520000 {
		skill = append(skill, triggerSkill{3, 22, 8800})
	}
	
	return skill
}

func skillTriggerCondition_616(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 122000 {
		skill = append(skill, triggerSkill{1, 22, 11250})
	}

	if f.Health <= 244000 {
		skill = append(skill, triggerSkill{2, 22, 11250})
	}

	if f.Health <= 366000 {
		skill = append(skill, triggerSkill{3, 22, 11250})
	}

	if f.Health <= 488000 {
		skill = append(skill, triggerSkill{4, 1465, 11250})
	}
	
	return skill
}

func skillTriggerCondition_617(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 227500 {
		skill = append(skill, triggerSkill{1, 1218, 12000})
	}

	if f.Health <= 292500 {
		skill = append(skill, triggerSkill{2, 18, 12000})
	}

	if f.Health <= 357500 {
		skill = append(skill, triggerSkill{3, 1218, 12000})
	}

	if f.Health <= 455000 {
		skill = append(skill, triggerSkill{4, 1218, 12000})
	}

	if f.Health <= 552500 {
		skill = append(skill, triggerSkill{5, 18, 12000})
	}
	
	return skill
}

func skillTriggerCondition_618(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 238000 {
		skill = append(skill, triggerSkill{1, 40, 12750})
	}

	if f.Health <= 306000 {
		skill = append(skill, triggerSkill{2, 40, 12750})
	}

	if f.Health <= 374000 {
		skill = append(skill, triggerSkill{3, 1469, 12750})
	}

	if f.Health <= 476000 {
		skill = append(skill, triggerSkill{4, 18, 12750})
	}

	if f.Health <= 578000 {
		skill = append(skill, triggerSkill{5, 40, 12750})
	}
	
	return skill
}

func skillTriggerCondition_619(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 142000 {
		skill = append(skill, triggerSkill{1, 30, 13000})
	}

	if f.Health <= 284000 {
		skill = append(skill, triggerSkill{2, 1470, 13000})
	}

	if f.Health <= 426000 {
		skill = append(skill, triggerSkill{3, 30, 13000})
	}

	if f.Health <= 568000 {
		skill = append(skill, triggerSkill{4, 30, 13000})
	}
	
	return skill
}

func skillTriggerCondition_620(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 262500 {
		skill = append(skill, triggerSkill{1, 1485, 13500})
	}

	if f.Health <= 337500 {
		skill = append(skill, triggerSkill{2, 1485, 13500})
	}

	if f.Health <= 412500 {
		skill = append(skill, triggerSkill{3, 1485, 13500})
	}

	if f.Health <= 525000 {
		skill = append(skill, triggerSkill{4, 1485, 13500})
	}

	if f.Health <= 637500 {
		skill = append(skill, triggerSkill{5, 1471, 13500})
	}
	
	return skill
}

func skillTriggerCondition_623(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 36000 {
		skill = append(skill, triggerSkill{1, 1277, 2000})
	}

	if f.Health <= 60000 {
		skill = append(skill, triggerSkill{2, 1277, 2000})
	}

	if f.Health <= 96000 {
		skill = append(skill, triggerSkill{3, 1277, 2000})
	}
	
	return skill
}

func skillTriggerCondition_626(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 99000 {
		skill = append(skill, triggerSkill{1, 26, 4680})
	}

	if f.Health <= 165000 {
		skill = append(skill, triggerSkill{2, 26, 4680})
	}

	if f.Health <= 264000 {
		skill = append(skill, triggerSkill{3, 26, 4680})
	}
	
	return skill
}

func skillTriggerCondition_629(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 90000 {
		skill = append(skill, triggerSkill{1, 22, 4000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{2, 22, 4000})
	}

	if f.Health <= 240000 {
		skill = append(skill, triggerSkill{3, 22, 4000})
	}
	
	return skill
}

func skillTriggerCondition_632(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 144000 {
		skill = append(skill, triggerSkill{1, 40, 8400})
	}

	if f.Health <= 240000 {
		skill = append(skill, triggerSkill{2, 40, 8400})
	}

	if f.Health <= 384000 {
		skill = append(skill, triggerSkill{3, 40, 8400})
	}
	
	return skill
}

func skillTriggerCondition_635(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 117000 {
		skill = append(skill, triggerSkill{1, 1432, 6360})
	}

	if f.Health <= 195000 {
		skill = append(skill, triggerSkill{2, 40, 6360})
	}

	if f.Health <= 312000 {
		skill = append(skill, triggerSkill{3, 40, 6360})
	}
	
	return skill
}

func skillTriggerCondition_643(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 27000 {
		skill = append(skill, triggerSkill{1, 1213, 1240})
	}

	if f.Health <= 45000 {
		skill = append(skill, triggerSkill{2, 1213, 1240})
	}

	if f.Health <= 72000 {
		skill = append(skill, triggerSkill{3, 1213, 1240})
	}
	
	return skill
}

func skillTriggerCondition_644(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 45000 {
		skill = append(skill, triggerSkill{1, 18, 2360})
	}

	if f.Health <= 75000 {
		skill = append(skill, triggerSkill{2, 22, 2360})
	}

	if f.Health <= 120000 {
		skill = append(skill, triggerSkill{3, 18, 2360})
	}
	
	return skill
}

func skillTriggerCondition_645(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 57000 {
		skill = append(skill, triggerSkill{1, 22, 2800})
	}

	if f.Health <= 95000 {
		skill = append(skill, triggerSkill{2, 22, 2800})
	}

	if f.Health <= 152000 {
		skill = append(skill, triggerSkill{3, 22, 2800})
	}
	
	return skill
}

func skillTriggerCondition_646(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 25000 {
		skill = append(skill, triggerSkill{1, 1147, 30000})
	}

	if f.Health <= 72000 {
		skill = append(skill, triggerSkill{2, 14, 3040})
	}

	if f.Health <= 120000 {
		skill = append(skill, triggerSkill{3, 30, 3040})
	}

	if f.Health <= 192000 {
		skill = append(skill, triggerSkill{4, 14, 3040})
	}
	
	return skill
}

func skillTriggerCondition_647(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 81000 {
		skill = append(skill, triggerSkill{1, 26, 3520})
	}

	if f.Health <= 135000 {
		skill = append(skill, triggerSkill{2, 26, 3520})
	}

	if f.Health <= 216000 {
		skill = append(skill, triggerSkill{3, 26, 3520})
	}
	
	return skill
}

func skillTriggerCondition_648(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{1, 108, 30000})
	}

	if f.Health <= 108000 {
		skill = append(skill, triggerSkill{2, 14, 5520})
	}

	if f.Health <= 180000 {
		skill = append(skill, triggerSkill{3, 14, 5520})
	}

	if f.Health <= 288000 {
		skill = append(skill, triggerSkill{4, 14, 5520})
	}
	
	return skill
}

func skillTriggerCondition_713(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{1, 14, 22500})
	}

	if f.Health <= 600000 {
		skill = append(skill, triggerSkill{2, 14, 22500})
	}

	if f.Health <= 750000 {
		skill = append(skill, triggerSkill{3, 1050, 500000})
	}

	if f.Health <= 1200000 {
		skill = append(skill, triggerSkill{4, 1499, 22500})
	}
	
	return skill
}

func skillTriggerCondition_718(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 201000 {
		skill = append(skill, triggerSkill{1, 1218, 9000})
	}

	if f.Health <= 335000 {
		skill = append(skill, triggerSkill{2, 1218, 9000})
	}

	if f.Health <= 536000 {
		skill = append(skill, triggerSkill{3, 1218, 9000})
	}
	
	return skill
}

func skillTriggerCondition_719(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 204000 {
		skill = append(skill, triggerSkill{1, 1486, 3000})
	}

	if f.Health <= 340000 {
		skill = append(skill, triggerSkill{2, 1486, 3000})
	}

	if f.Health <= 544000 {
		skill = append(skill, triggerSkill{3, 1486, 3000})
	}
	
	return skill
}

func skillTriggerCondition_723(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{1, 1509, 14000})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{2, 1486, 2000})
	}

	if f.Health <= 450000 {
		skill = append(skill, triggerSkill{3, 1509, 14000})
	}

	if f.Health <= 600000 {
		skill = append(skill, triggerSkill{4, 1486, 2000})
	}
	
	return skill
}

func skillTriggerCondition_724(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 269500 {
		skill = append(skill, triggerSkill{1, 1509, 14500})
	}

	if f.Health <= 346500 {
		skill = append(skill, triggerSkill{2, 1509, 14500})
	}

	if f.Health <= 423500 {
		skill = append(skill, triggerSkill{3, 1520, 14500})
	}

	if f.Health <= 539000 {
		skill = append(skill, triggerSkill{4, 1509, 14500})
	}

	if f.Health <= 654500 {
		skill = append(skill, triggerSkill{5, 1509, 14500})
	}
	
	return skill
}

func skillTriggerCondition_725(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 276500 {
		skill = append(skill, triggerSkill{1, 1517, 15000})
	}

	if f.Health <= 355500 {
		skill = append(skill, triggerSkill{2, 1517, 15000})
	}

	if f.Health <= 434500 {
		skill = append(skill, triggerSkill{3, 1517, 15000})
	}

	if f.Health <= 553000 {
		skill = append(skill, triggerSkill{4, 1521, 15000})
	}

	if f.Health <= 671500 {
		skill = append(skill, triggerSkill{5, 1517, 15000})
	}
	
	return skill
}

func skillTriggerCondition_726(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 283500 {
		skill = append(skill, triggerSkill{1, 1518, 1700})
	}

	if f.Health <= 364500 {
		skill = append(skill, triggerSkill{2, 1510, 15500})
	}

	if f.Health <= 445500 {
		skill = append(skill, triggerSkill{3, 1518, 1700})
	}

	if f.Health <= 567000 {
		skill = append(skill, triggerSkill{4, 1510, 15500})
	}

	if f.Health <= 688500 {
		skill = append(skill, triggerSkill{5, 1518, 1700})
	}
	
	return skill
}

func skillTriggerCondition_727(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 255500 {
		skill = append(skill, triggerSkill{1, 1485, 16000})
	}

	if f.Health <= 328500 {
		skill = append(skill, triggerSkill{2, 1485, 16000})
	}

	if f.Health <= 401500 {
		skill = append(skill, triggerSkill{3, 1485, 16000})
	}

	if f.Health <= 511000 {
		skill = append(skill, triggerSkill{4, 1485, 16000})
	}

	if f.Health <= 620500 {
		skill = append(skill, triggerSkill{5, 1522, 16000})
	}
	
	return skill
}

func skillTriggerCondition_732(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 207000 {
		skill = append(skill, triggerSkill{1, 1509, 9400})
	}

	if f.Health <= 345000 {
		skill = append(skill, triggerSkill{2, 1509, 9400})
	}

	if f.Health <= 552000 {
		skill = append(skill, triggerSkill{3, 1509, 9400})
	}
	
	return skill
}

func skillTriggerCondition_733(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 210000 {
		skill = append(skill, triggerSkill{1, 1485, 9600})
	}

	if f.Health <= 350000 {
		skill = append(skill, triggerSkill{2, 1485, 9600})
	}

	if f.Health <= 560000 {
		skill = append(skill, triggerSkill{3, 1485, 9600})
	}
	
	return skill
}

func skillTriggerCondition_738(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 216000 {
		skill = append(skill, triggerSkill{1, 18, 9600})
	}

	if f.Health <= 360000 {
		skill = append(skill, triggerSkill{2, 18, 9600})
	}

	if f.Health <= 576000 {
		skill = append(skill, triggerSkill{3, 18, 9600})
	}
	
	return skill
}

func skillTriggerCondition_739(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 219000 {
		skill = append(skill, triggerSkill{1, 45, 9800})
	}

	if f.Health <= 365000 {
		skill = append(skill, triggerSkill{2, 45, 9800})
	}

	if f.Health <= 584000 {
		skill = append(skill, triggerSkill{3, 45, 9800})
	}
	
	return skill
}

func skillTriggerCondition_744(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 225000 {
		skill = append(skill, triggerSkill{1, 34, 10000})
	}

	if f.Health <= 375000 {
		skill = append(skill, triggerSkill{2, 34, 10000})
	}

	if f.Health <= 600000 {
		skill = append(skill, triggerSkill{3, 34, 10000})
	}
	
	return skill
}

func skillTriggerCondition_745(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 228000 {
		skill = append(skill, triggerSkill{1, 1515, 1020})
	}

	if f.Health <= 380000 {
		skill = append(skill, triggerSkill{2, 1515, 1020})
	}

	if f.Health <= 608000 {
		skill = append(skill, triggerSkill{3, 1515, 1020})
	}
	
	return skill
}

func skillTriggerCondition_750(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 234000 {
		skill = append(skill, triggerSkill{1, 1485, 10400})
	}

	if f.Health <= 390000 {
		skill = append(skill, triggerSkill{2, 1485, 10400})
	}

	if f.Health <= 624000 {
		skill = append(skill, triggerSkill{3, 1485, 10400})
	}
	
	return skill
}

func skillTriggerCondition_751(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 237000 {
		skill = append(skill, triggerSkill{1, 18, 10600})
	}

	if f.Health <= 395000 {
		skill = append(skill, triggerSkill{2, 18, 10600})
	}

	if f.Health <= 632000 {
		skill = append(skill, triggerSkill{3, 18, 10600})
	}
	
	return skill
}

func skillTriggerCondition_752(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 229500 {
		skill = append(skill, triggerSkill{1, 1476, 6900})
	}

	if f.Health <= 382500 {
		skill = append(skill, triggerSkill{2, 1476, 6900})
	}

	if f.Health <= 612000 {
		skill = append(skill, triggerSkill{3, 1476, 6900})
	}
	
	return skill
}

func skillTriggerCondition_754(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 243000 {
		skill = append(skill, triggerSkill{1, 1790, 7260})
	}

	if f.Health <= 405000 {
		skill = append(skill, triggerSkill{2, 1752, 7260})
	}

	if f.Health <= 648000 {
		skill = append(skill, triggerSkill{3, 1752, 7260})
	}
	
	return skill
}

func skillTriggerCondition_758(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 153000 {
		skill = append(skill, triggerSkill{1, 1214, 9000})
	}

	if f.Health <= 255000 {
		skill = append(skill, triggerSkill{2, 1214, 9000})
	}

	if f.Health <= 408000 {
		skill = append(skill, triggerSkill{3, 1214, 9000})
	}
	
	return skill
}

func skillTriggerCondition_761(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 162000 {
		skill = append(skill, triggerSkill{1, 1213, 9600})
	}

	if f.Health <= 270000 {
		skill = append(skill, triggerSkill{2, 1213, 9600})
	}

	if f.Health <= 432000 {
		skill = append(skill, triggerSkill{3, 1213, 9600})
	}
	
	return skill
}

func skillTriggerCondition_775(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 267000 {
		skill = append(skill, triggerSkill{1, 1504, 8000})
	}

	if f.Health <= 445000 {
		skill = append(skill, triggerSkill{2, 34, 8000})
	}

	if f.Health <= 712000 {
		skill = append(skill, triggerSkill{3, 1504, 8000})
	}
	
	return skill
}

func skillTriggerCondition_777(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 291000 {
		skill = append(skill, triggerSkill{1, 1509, 8900})
	}

	if f.Health <= 485000 {
		skill = append(skill, triggerSkill{2, 1509, 8900})
	}

	if f.Health <= 776000 {
		skill = append(skill, triggerSkill{3, 1509, 8900})
	}
	
	return skill
}

func skillTriggerCondition_779(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{1, 1490, 9600})
	}

	if f.Health <= 312000 {
		skill = append(skill, triggerSkill{2, 14, 9600})
	}

	if f.Health <= 520000 {
		skill = append(skill, triggerSkill{3, 14, 9600})
	}

	if f.Health <= 832000 {
		skill = append(skill, triggerSkill{4, 14, 9600})
	}
	
	return skill
}

func skillTriggerCondition_795(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 160000 {
		skill = append(skill, triggerSkill{1, 1428, 16500})
	}

	if f.Health <= 320000 {
		skill = append(skill, triggerSkill{2, 1428, 16500})
	}

	if f.Health <= 480000 {
		skill = append(skill, triggerSkill{3, 1600, 16500})
	}

	if f.Health <= 640000 {
		skill = append(skill, triggerSkill{4, 1428, 16500})
	}
	
	return skill
}

func skillTriggerCondition_797(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 290500 {
		skill = append(skill, triggerSkill{1, 1595, 17100})
	}

	if f.Health <= 373500 {
		skill = append(skill, triggerSkill{2, 1595, 17100})
	}

	if f.Health <= 456500 {
		skill = append(skill, triggerSkill{3, 1601, 17100})
	}

	if f.Health <= 581000 {
		skill = append(skill, triggerSkill{4, 1595, 17100})
	}

	if f.Health <= 705500 {
		skill = append(skill, triggerSkill{5, 1595, 17100})
	}
	
	return skill
}

func skillTriggerCondition_799(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 301000 {
		skill = append(skill, triggerSkill{1, 18, 17700})
	}

	if f.Health <= 387000 {
		skill = append(skill, triggerSkill{2, 18, 17700})
	}

	if f.Health <= 473000 {
		skill = append(skill, triggerSkill{3, 18, 17700})
	}

	if f.Health <= 602000 {
		skill = append(skill, triggerSkill{4, 1602, 17700})
	}

	if f.Health <= 731000 {
		skill = append(skill, triggerSkill{5, 18, 17700})
	}
	
	return skill
}

func skillTriggerCondition_801(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 311500 {
		skill = append(skill, triggerSkill{1, 1598, 18300})
	}

	if f.Health <= 400500 {
		skill = append(skill, triggerSkill{2, 1598, 18300})
	}

	if f.Health <= 489000 {
		skill = append(skill, triggerSkill{3, 1598, 18300})
	}

	if f.Health <= 623000 {
		skill = append(skill, triggerSkill{4, 1598, 18300})
	}

	if f.Health <= 756500 {
		skill = append(skill, triggerSkill{5, 1603, 18300})
	}
	
	return skill
}

func skillTriggerCondition_803(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 322000 {
		skill = append(skill, triggerSkill{1, 30, 18900})
	}

	if f.Health <= 414000 {
		skill = append(skill, triggerSkill{2, 1510, 18900})
	}

	if f.Health <= 506000 {
		skill = append(skill, triggerSkill{3, 30, 18900})
	}

	if f.Health <= 644000 {
		skill = append(skill, triggerSkill{4, 1510, 18900})
	}

	if f.Health <= 782000 {
		skill = append(skill, triggerSkill{5, 30, 18900})
	}
	
	return skill
}

func skillTriggerCondition_808(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 240000 {
		skill = append(skill, triggerSkill{1, 1428, 10800})
	}

	if f.Health <= 400000 {
		skill = append(skill, triggerSkill{2, 1428, 10800})
	}

	if f.Health <= 640000 {
		skill = append(skill, triggerSkill{3, 1428, 10800})
	}
	
	return skill
}

func skillTriggerCondition_809(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 243000 {
		skill = append(skill, triggerSkill{1, 1218, 3000})
	}

	if f.Health <= 405000 {
		skill = append(skill, triggerSkill{2, 1218, 3000})
	}

	if f.Health <= 648000 {
		skill = append(skill, triggerSkill{3, 1218, 3000})
	}
	
	return skill
}

func skillTriggerCondition_814(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 249000 {
		skill = append(skill, triggerSkill{1, 18, 11280})
	}

	if f.Health <= 415000 {
		skill = append(skill, triggerSkill{2, 18, 11280})
	}

	if f.Health <= 664000 {
		skill = append(skill, triggerSkill{3, 18, 11280})
	}
	
	return skill
}

func skillTriggerCondition_815(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 252000 {
		skill = append(skill, triggerSkill{1, 1213, 11520})
	}

	if f.Health <= 420000 {
		skill = append(skill, triggerSkill{2, 1213, 11520})
	}

	if f.Health <= 672000 {
		skill = append(skill, triggerSkill{3, 1213, 11520})
	}
	
	return skill
}

func skillTriggerCondition_820(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 258000 {
		skill = append(skill, triggerSkill{1, 1485, 11760})
	}

	if f.Health <= 430000 {
		skill = append(skill, triggerSkill{2, 1485, 11760})
	}

	if f.Health <= 688000 {
		skill = append(skill, triggerSkill{3, 1485, 11760})
	}
	
	return skill
}

func skillTriggerCondition_821(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 261000 {
		skill = append(skill, triggerSkill{1, 1595, 12000})
	}

	if f.Health <= 435000 {
		skill = append(skill, triggerSkill{2, 1595, 12000})
	}

	if f.Health <= 696000 {
		skill = append(skill, triggerSkill{3, 1595, 12000})
	}
	
	return skill
}

func skillTriggerCondition_826(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 267000 {
		skill = append(skill, triggerSkill{1, 34, 12240})
	}

	if f.Health <= 445000 {
		skill = append(skill, triggerSkill{2, 34, 12240})
	}

	if f.Health <= 712000 {
		skill = append(skill, triggerSkill{3, 34, 12240})
	}
	
	return skill
}

func skillTriggerCondition_827(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 270000 {
		skill = append(skill, triggerSkill{1, 34, 12480})
	}

	if f.Health <= 450000 {
		skill = append(skill, triggerSkill{2, 34, 12480})
	}

	if f.Health <= 720000 {
		skill = append(skill, triggerSkill{3, 34, 12480})
	}
	
	return skill
}

func skillTriggerCondition_832(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 276000 {
		skill = append(skill, triggerSkill{1, 1215, 12720})
	}

	if f.Health <= 460000 {
		skill = append(skill, triggerSkill{2, 1215, 12720})
	}

	if f.Health <= 736000 {
		skill = append(skill, triggerSkill{3, 1215, 12720})
	}
	
	return skill
}

func skillTriggerCondition_833(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 279000 {
		skill = append(skill, triggerSkill{1, 30, 12960})
	}

	if f.Health <= 465000 {
		skill = append(skill, triggerSkill{2, 30, 12960})
	}

	if f.Health <= 744000 {
		skill = append(skill, triggerSkill{3, 30, 12960})
	}
	
	return skill
}

func skillTriggerCondition_836(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 180000 {
		skill = append(skill, triggerSkill{1, 18, 10200})
	}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{2, 18, 10200})
	}

	if f.Health <= 480000 {
		skill = append(skill, triggerSkill{3, 18, 10200})
	}
	
	return skill
}

func skillTriggerCondition_839(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 189000 {
		skill = append(skill, triggerSkill{1, 1277, 10800})
	}

	if f.Health <= 315000 {
		skill = append(skill, triggerSkill{2, 1277, 10800})
	}

	if f.Health <= 504000 {
		skill = append(skill, triggerSkill{3, 1277, 10800})
	}
	
	return skill
}

func skillTriggerCondition_842(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 198000 {
		skill = append(skill, triggerSkill{1, 26, 11400})
	}

	if f.Health <= 330000 {
		skill = append(skill, triggerSkill{2, 22, 11400})
	}

	if f.Health <= 528000 {
		skill = append(skill, triggerSkill{3, 18, 11400})
	}
	
	return skill
}

func skillTriggerCondition_843(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 190000 {
		skill = append(skill, triggerSkill{1, 1501, 19500})
	}

	if f.Health <= 380000 {
		skill = append(skill, triggerSkill{2, 1501, 19500})
	}

	if f.Health <= 570000 {
		skill = append(skill, triggerSkill{3, 1623, 19500})
	}

	if f.Health <= 760000 {
		skill = append(skill, triggerSkill{4, 1501, 19500})
	}
	
	return skill
}

func skillTriggerCondition_845(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 343000 {
		skill = append(skill, triggerSkill{1, 1428, 20000})
	}

	if f.Health <= 441000 {
		skill = append(skill, triggerSkill{2, 1428, 20000})
	}

	if f.Health <= 539000 {
		skill = append(skill, triggerSkill{3, 1624, 20000})
	}

	if f.Health <= 686000 {
		skill = append(skill, triggerSkill{4, 1428, 20000})
	}

	if f.Health <= 833000 {
		skill = append(skill, triggerSkill{5, 1624, 20000})
	}
	
	return skill
}

func skillTriggerCondition_847(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 353500 {
		skill = append(skill, triggerSkill{1, 1634, 20500})
	}

	if f.Health <= 454500 {
		skill = append(skill, triggerSkill{2, 1634, 20500})
	}

	if f.Health <= 555500 {
		skill = append(skill, triggerSkill{3, 1634, 20500})
	}

	if f.Health <= 707000 {
		skill = append(skill, triggerSkill{4, 1625, 20500})
	}

	if f.Health <= 858500 {
		skill = append(skill, triggerSkill{5, 1634, 20500})
	}
	
	return skill
}

func skillTriggerCondition_849(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 364000 {
		skill = append(skill, triggerSkill{1, 1636, 21000})
	}

	if f.Health <= 468000 {
		skill = append(skill, triggerSkill{2, 1636, 21000})
	}

	if f.Health <= 572000 {
		skill = append(skill, triggerSkill{3, 1636, 21000})
	}

	if f.Health <= 728000 {
		skill = append(skill, triggerSkill{4, 1636, 21000})
	}

	if f.Health <= 884000 {
		skill = append(skill, triggerSkill{5, 1626, 21000})
	}
	
	return skill
}

func skillTriggerCondition_851(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 374500 {
		skill = append(skill, triggerSkill{1, 1432, 21500})
	}

	if f.Health <= 481500 {
		skill = append(skill, triggerSkill{2, 1432, 21500})
	}

	if f.Health <= 588500 {
		skill = append(skill, triggerSkill{3, 1432, 21500})
	}

	if f.Health <= 749000 {
		skill = append(skill, triggerSkill{4, 1432, 21500})
	}

	if f.Health <= 909500 {
		skill = append(skill, triggerSkill{5, 1432, 21500})
	}
	
	return skill
}

func skillTriggerCondition_859(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 330000 {
		skill = append(skill, triggerSkill{1, 1480, 10200})
	}

	if f.Health <= 550000 {
		skill = append(skill, triggerSkill{2, 1485, 10200})
	}

	if f.Health <= 880000 {
		skill = append(skill, triggerSkill{3, 1480, 10200})
	}
	
	return skill
}

func skillTriggerCondition_861(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{1, 1501, 11140})
	}

	if f.Health <= 345000 {
		skill = append(skill, triggerSkill{2, 1721, 11140})
	}

	if f.Health <= 575000 {
		skill = append(skill, triggerSkill{3, 1747, 11140})
	}

	if f.Health <= 920000 {
		skill = append(skill, triggerSkill{4, 1501, 11140})
	}
	
	return skill
}

func skillTriggerCondition_865(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 210000 {
		skill = append(skill, triggerSkill{1, 1277, 13600})
	}

	if f.Health <= 350000 {
		skill = append(skill, triggerSkill{2, 1277, 13600})
	}

	if f.Health <= 560000 {
		skill = append(skill, triggerSkill{3, 1277, 13600})
	}
	
	return skill
}

func skillTriggerCondition_868(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 219000 {
		skill = append(skill, triggerSkill{1, 26, 14800})
	}

	if f.Health <= 365000 {
		skill = append(skill, triggerSkill{2, 22, 14800})
	}

	if f.Health <= 584000 {
		skill = append(skill, triggerSkill{3, 18, 14800})
	}
	
	return skill
}

func skillTriggerCondition_951(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 285000 {
		skill = append(skill, triggerSkill{1, 1501, 13400})
	}

	if f.Health <= 475000 {
		skill = append(skill, triggerSkill{2, 1501, 13400})
	}

	if f.Health <= 760000 {
		skill = append(skill, triggerSkill{3, 1501, 13400})
	}
	
	return skill
}

func skillTriggerCondition_952(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 288000 {
		skill = append(skill, triggerSkill{1, 1218, 13520})
	}

	if f.Health <= 480000 {
		skill = append(skill, triggerSkill{2, 1218, 13520})
	}

	if f.Health <= 768000 {
		skill = append(skill, triggerSkill{3, 1218, 13520})
	}
	
	return skill
}

func skillTriggerCondition_957(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 294000 {
		skill = append(skill, triggerSkill{1, 1428, 13760})
	}

	if f.Health <= 490000 {
		skill = append(skill, triggerSkill{2, 1428, 13760})
	}

	if f.Health <= 784000 {
		skill = append(skill, triggerSkill{3, 1428, 13760})
	}
	
	return skill
}

func skillTriggerCondition_958(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 291000 {
		skill = append(skill, triggerSkill{1, 1595, 13640})
	}

	if f.Health <= 485000 {
		skill = append(skill, triggerSkill{2, 1595, 13640})
	}

	if f.Health <= 776000 {
		skill = append(skill, triggerSkill{3, 1595, 13640})
	}
	
	return skill
}

func skillTriggerCondition_963(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 300000 {
		skill = append(skill, triggerSkill{1, 1634, 13880})
	}

	if f.Health <= 500000 {
		skill = append(skill, triggerSkill{2, 1634, 13880})
	}

	if f.Health <= 800000 {
		skill = append(skill, triggerSkill{3, 1634, 13880})
	}
	
	return skill
}

func skillTriggerCondition_964(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 303000 {
		skill = append(skill, triggerSkill{1, 1484, 14000})
	}

	if f.Health <= 505000 {
		skill = append(skill, triggerSkill{2, 1484, 14000})
	}

	if f.Health <= 808000 {
		skill = append(skill, triggerSkill{3, 1484, 14000})
	}
	
	return skill
}

func skillTriggerCondition_969(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 309000 {
		skill = append(skill, triggerSkill{1, 1485, 14120})
	}

	if f.Health <= 515000 {
		skill = append(skill, triggerSkill{2, 1485, 14120})
	}

	if f.Health <= 824000 {
		skill = append(skill, triggerSkill{3, 1485, 14120})
	}
	
	return skill
}

func skillTriggerCondition_970(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 312000 {
		skill = append(skill, triggerSkill{1, 1636, 14240})
	}

	if f.Health <= 520000 {
		skill = append(skill, triggerSkill{2, 1636, 14240})
	}

	if f.Health <= 832000 {
		skill = append(skill, triggerSkill{3, 1636, 14240})
	}
	
	return skill
}

func skillTriggerCondition_975(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 318000 {
		skill = append(skill, triggerSkill{1, 1432, 14360})
	}

	if f.Health <= 530000 {
		skill = append(skill, triggerSkill{2, 1432, 14360})
	}

	if f.Health <= 848000 {
		skill = append(skill, triggerSkill{3, 1432, 14360})
	}
	
	return skill
}

func skillTriggerCondition_976(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 321000 {
		skill = append(skill, triggerSkill{1, 1638, 14480})
	}

	if f.Health <= 535000 {
		skill = append(skill, triggerSkill{2, 1638, 14480})
	}

	if f.Health <= 856000 {
		skill = append(skill, triggerSkill{3, 1638, 14480})
	}
	
	return skill
}

func skillTriggerCondition_983(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{1, 18, 3040})
	}
	
	return skill
}

func skillTriggerCondition_988(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 7 - 1 {
		skill = append(skill, triggerSkill{1, 1727, 1500})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1727, 1500})
	}
	
	return skill
}

func skillTriggerCondition_996(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1724, 100})
	}

	if f.battle.GetRounds() >= 7 - 1 {
		skill = append(skill, triggerSkill{2, 1724, 100})
	}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{3, 1724, 100})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{4, 1724, 100})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{5, 1724, 100})
	}
	
	return skill
}

func skillTriggerCondition_1003(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1728, 2900})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1728, 2900})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1728, 2900})
	}
	
	return skill
}

func skillTriggerCondition_1004(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1729, 2900})
	}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{2, 1729, 2900})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{3, 1729, 2900})
	}
	
	return skill
}

func skillTriggerCondition_1007(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 34, 5000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 34, 5000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 34, 5000})
	}
	
	return skill
}

func skillTriggerCondition_1018(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1734, 15})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1734, 15})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1734, 15})
	}
	
	return skill
}

func skillTriggerCondition_1019(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 7 - 1 {
		skill = append(skill, triggerSkill{1, 1735, 8000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1735, 8000})
	}
	
	return skill
}

func skillTriggerCondition_1025(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1485, 8000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1485, 8000})
	}
	
	return skill
}

func skillTriggerCondition_1029(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1509, 10000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1509, 10000})
	}
	
	return skill
}

func skillTriggerCondition_1037(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1714, 100})
	}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{2, 1485, 12000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{3, 1714, 100})
	}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{4, 1485, 12000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{5, 1714, 100})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{6, 1485, 12000})
	}
	
	return skill
}

func skillTriggerCondition_1045(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1509, 2600})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1509, 2600})
	}
	
	return skill
}

func skillTriggerCondition_1046(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	
	return skill
}

func skillTriggerCondition_1052(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	
	return skill
}

func skillTriggerCondition_1055(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1509, 5000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1509, 5000})
	}
	
	return skill
}

func skillTriggerCondition_1059(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{1, 34, 6000})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{2, 34, 6000})
	}
	
	return skill
}

func skillTriggerCondition_1061(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1729, 500})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1729, 5000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1729, 5000})
	}
	
	return skill
}

func skillTriggerCondition_1062(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1714, 1000})
	}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{2, 1714, 1000})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{3, 1714, 1000})
	}
	
	return skill
}

func skillTriggerCondition_1065(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1729, 2000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1729, 2000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{3, 1787, 1000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{4, 1729, 2000})
	}

	if f.Health <= 150000 {
		skill = append(skill, triggerSkill{5, 1788, 1000})
	}
	
	return skill
}

func skillTriggerCondition_1068(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 30000 {
		skill = append(skill, triggerSkill{1, 1474, 230000})
	}
	
	return skill
}

func skillTriggerCondition_1075(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1752, 9000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1752, 9000})
	}
	
	return skill
}

func skillTriggerCondition_1085(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	
	return skill
}

func skillTriggerCondition_1086(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1743, 10000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 1743, 10000})
	}
	
	return skill
}

func skillTriggerCondition_1093(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1748, 13000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1748, 13000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1748, 13000})
	}
	
	return skill
}

func skillTriggerCondition_1095(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1714, 15})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1714, 15})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1714, 15})
	}
	
	return skill
}

func skillTriggerCondition_1097(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{1, 18, 3600})
	}
	
	return skill
}

func skillTriggerCondition_1101(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{1, 1428, 40000})
	}
	
	return skill
}

func skillTriggerCondition_1103(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{1, 34, 6400})
	}
	
	return skill
}

func skillTriggerCondition_1109(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{1, 1697, 100})
	}
	
	return skill
}

func skillTriggerCondition_1113(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1699, 5000})
	}

	if f.battle.GetRounds() >= 5 - 1 {
		skill = append(skill, triggerSkill{2, 1699, 5000})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{3, 1699, 5000})
	}
	
	return skill
}

func skillTriggerCondition_1122(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 50000 {
		skill = append(skill, triggerSkill{1, 34, 4000})
	}

	if f.Health <= 130000 {
		skill = append(skill, triggerSkill{2, 34, 4000})
	}

	if f.Health <= 200000 {
		skill = append(skill, triggerSkill{3, 34, 4000})
	}
	
	return skill
}

func skillTriggerCondition_1129(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1707, 10000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1707, 10000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{3, 1707, 10000})
	}
	
	return skill
}

func skillTriggerCondition_1130(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.Health <= 500000 {
		skill = append(skill, triggerSkill{1, 22, 5000})
	}
	
	return skill
}

func skillTriggerCondition_1135(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1711, 10000})
	}
	
	return skill
}

func skillTriggerCondition_1137(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 26, 12000})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{2, 26, 12000})
	}
	
	return skill
}

func skillTriggerCondition_1146(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 9 - 1 {
		skill = append(skill, triggerSkill{1, 1714, 15})
	}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{2, 30, 12000})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{3, 1714, 15})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{4, 30, 12000})
	}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{5, 1714, 15})
	}
	
	return skill
}

func skillTriggerCondition_1148(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{1, 1697, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{2, 1628, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{3, 1628, 9999999})
	}
	
	return skill
}

func skillTriggerCondition_1149(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{1, 1779, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{2, 1776, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{3, 1776, 9999999})
	}
	
	return skill
}

func skillTriggerCondition_1150(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 8 - 1 {
		skill = append(skill, triggerSkill{1, 1780, 99999})
	}

	if f.battle.GetRounds() >= 6 - 1 {
		skill = append(skill, triggerSkill{2, 1780, 9999999})
	}

	if f.battle.GetRounds() >= 4 - 1 {
		skill = append(skill, triggerSkill{3, 1780, 9999999})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{4, 1780, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{5, 1775, 9999999})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{6, 1775, 9999999})
	}
	
	return skill
}

func skillTriggerCondition_1151(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	
	return skill
}

func skillTriggerCondition_1153(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}
	
	return skill
}

func skillTriggerCondition_1158(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{1, 15, 100})
	}

	if f.Health <= 5000 {
		skill = append(skill, triggerSkill{2, 11, 100})
	}
	
	return skill
}

func skillTriggerCondition_1159(f *Fighter) []triggerSkill {
	skill := []triggerSkill{}

	if f.battle.GetRounds() >= 3 - 1 {
		skill = append(skill, triggerSkill{1, 26, 500})
	}

	if f.battle.GetRounds() >= 2 - 1 {
		skill = append(skill, triggerSkill{2, 18, 500})
	}

	if f.battle.GetRounds() >= 1 - 1 {
		skill = append(skill, triggerSkill{3, 14, 500})
	}
	
	return skill
}

//获取灵宠概率
func getBattlePetCatchRate(pet *Fighter) int {
	health_rate := float64(pet.Health) / float64(pet.MaxHealth)
	_ = health_rate 
	switch pet.RoleId {
	case 91:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 92:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 117:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 118:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 119:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 296:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 297:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 299:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 444:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 445:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 446:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 447:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 450:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 451:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 452:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 453:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 454:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 455:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 456:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 457:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 458:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 459:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 460:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 461:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 462:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 463:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 464:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 465:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 466:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 467:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 468:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 469:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 470:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 471:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 472:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 473:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 474:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 475:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 476:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 477:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 478:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 479:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 480:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 481:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 482:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 483:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 484:
		if health_rate <= 10 {
			return 30
		}
		if health_rate <= 30 {
			return 15
		}
		if health_rate <= 60 {
			return 10
		}
	case 771:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 772:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 773:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	case 774:
		if health_rate <= 10 {
			return 50
		}
		if health_rate <= 50 {
			return 20
		}
		if health_rate <= 80 {
			return 10
		}
	default: return 0 
	}
	return 0
}
