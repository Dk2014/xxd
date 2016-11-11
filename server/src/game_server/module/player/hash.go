package player

import (
	"core/fail"
	"core/time"
	"encoding/base64"
	"game_server/module"
	"strconv"
)

var (
	HASH_DELIMITER = []byte("-")
	HASH_SALT      = []byte("Uh981-jsdO!%&^#kahdbp0;v/'@$1*UOysd")
)

func platformLoginTest(typeVal, roleId int8, user, nick, hash []byte, time int64, wegames_platform_id []byte) bool {
	// hash = type-openid-salt-time-nick-roleId
	newHash := GetPlatformLoginHash(typeVal, roleId, user, nick, time, wegames_platform_id)
	hashcode, err := base64.StdEncoding.DecodeString(string(hash))
	fail.When(err != nil, err)

	for i, v := range newHash {
		if v != hashcode[i] {
			return false
		}
	}

	return true
}

func crossServerLoginTest(pid, time int64, openid, nick, hash []byte, roleId int8, wegames_platform_id []byte) bool {
	hashSource := [][]byte{
		[]byte(strconv.FormatInt(pid, 10)),
		nick,
		openid,
		[]byte(strconv.FormatInt(time, 10)),
		HASH_SALT,
		[]byte(strconv.FormatInt(int64(roleId), 10)),
	}
	if len(wegames_platform_id) > 0 {
		hashSource = append(hashSource, wegames_platform_id)
	}

	newHash := module.HashNow(HASH_DELIMITER, hashSource)
	hashcode, err := base64.StdEncoding.DecodeString(string(hash))
	fail.When(err != nil, err)

	for i, v := range newHash {
		if v != hashcode[i] {
			return false
		}
	}

	return true
}

func getLoginInfo(state *module.SessionState) ([]byte, int64) {
	nowTime := time.GetNowTime()

	player := state.Database.Lookup.Player(state.PlayerId)
	hashSource := [][]byte{
		[]byte(strconv.FormatInt(state.PlayerId, 10)),
		state.PlayerNick,
		[]byte(player.User),
		[]byte(strconv.FormatInt(nowTime, 10)),
		HASH_SALT,
		[]byte(strconv.FormatInt(int64(state.RoleId), 10)),
	}
	if len(state.WegamesPlatformUid) > 0 {
		hashSource = append(hashSource, state.WegamesPlatformUid)
	}
	hash := module.HashNow(HASH_DELIMITER, hashSource)

	return []byte(base64.StdEncoding.EncodeToString(hash)), nowTime
}

func GetPlatformLoginHash(typeVal, roleId int8, user, nick []byte, time int64, wegames_platform_id []byte) []byte {
	hashSource :=
		[][]byte{
			[]byte(strconv.FormatInt(int64(typeVal), 10)),
			user,
			HASH_SALT,
			[]byte(strconv.FormatInt(time, 10)),
			nick,
			[]byte(strconv.FormatInt(int64(roleId), 10)),
		}
	if len(wegames_platform_id) > 0 {
		hashSource = append(hashSource, wegames_platform_id)
	}
	return module.HashNow(HASH_DELIMITER, hashSource)
}
