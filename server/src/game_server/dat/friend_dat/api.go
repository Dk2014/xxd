package friend_dat

func GetPlatformFriendAward(origin, current int32) (awards []*PlatformFriendAward) {
	for _, award := range platformFriendAwardConfigs {
		if origin < award.RequireFriendNum && award.RequireFriendNum <= current {
			awards = append(awards, award)
		}
	}
	return
}
