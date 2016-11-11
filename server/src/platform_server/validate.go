package platform_server

import (
	"core/fail"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	TYPE_MOBILE_IOS_WEIXIN     = 1   // IOS微信平台
	TYPE_MOBILE_IOS_QQ         = 2   // IOS手Q平台
	TYPE_MOBILE_IOS_GUEST      = 5   // IOS游客平台
	TYPE_MOBILE_ANDROID_WEIXIN = 17  // 安卓微信平台
	TYPE_MOBILE_ANDROID_QQ     = 18  // 安卓手Q平台
	TYPE_MOBILE_ANDROID_GUEST  = 21  // 安卓游客平台
	TYPE_MOBILE_AUDIT_WEIXIN   = 252 // ios微信审核
	TYPE_MOBILE_AUDIT_QQ       = 253 // ios手Q审核
	TYPE_MOBILE_AUDIT_GUEST    = 254 // ios游客审核
	TYPE_MOBILE_SANDBOX        = 255 // 内部测试
)

const (
	GAME_SERVER_STATUS_MAINTAIN = 0 // 维护
	GAME_SERVER_STATUS_CLEAR    = 1 // 通畅
	GAME_SERVER_STATUS_BUSY     = 2 // 繁忙
	GAME_SERVER_STATUS_CROWDING = 3 // 拥挤
)

const (
	GAME_ROLE_YIFENG_ID = 1
	GAME_ROLE_XINRAN_ID = 2
)

var (
	illegalNick = regexp.MustCompile(`肛交|口交`)
)

func ValidateServerType(typeVal uint8) bool {
	switch typeVal {
	case TYPE_MOBILE_IOS_WEIXIN, TYPE_MOBILE_IOS_QQ, TYPE_MOBILE_IOS_GUEST,
		TYPE_MOBILE_ANDROID_WEIXIN, TYPE_MOBILE_ANDROID_QQ,
		TYPE_MOBILE_ANDROID_GUEST, TYPE_MOBILE_SANDBOX:
		return true
	}

	return false
}
func ValidateMobileType(typeVal uint8) {
	fail.When(!ValidateServerType(typeVal), "incorrect mobile platform type: "+strconv.Itoa(int(typeVal)))
}

func ValidateRoleId(roleId int8) {
	fail.When(roleId != GAME_ROLE_YIFENG_ID && roleId != GAME_ROLE_XINRAN_ID, "incorrect role id")
}

func ValidateNickname(nickname string) (err error) {
	runes := []rune(nickname)

	if strings.ContainsAny(nickname, " .<>&'\"_@#") {
		return errors.New("could not contains ' .<>&'\"_@#' ")
	}

	if illegalNick.MatchString(nickname) {
		return errors.New("could not contains illegal words ")
	}

	for _, v := range runes {
		if v >= 0x10000 && v <= 0x10FFFF {
			return errors.New("could not contains emoji")
		}
	}
	return
}
