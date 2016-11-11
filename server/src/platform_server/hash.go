package platform_server

import (
	"crypto/md5"
	"strconv"
)

var (
	HASH_PLATFORM_SPLIT = []byte("-")
	HASH_PLATFORM_SALT  = []byte("Uh981-jsdO!%&^#kahdbp0;v/'@$1*UOysd")
)

func hashResponse(typeVal uint8, roleId int8, openId string, nick string, time int64) []byte {

	// hash = type-openid-salt-time-nick-roleId

	hash := md5.New()

	hash.Write([]byte(strconv.FormatInt(int64(typeVal), 10)))
	hash.Write(HASH_PLATFORM_SPLIT)

	hash.Write([]byte(openId))
	hash.Write(HASH_PLATFORM_SPLIT)

	hash.Write(HASH_PLATFORM_SALT)
	hash.Write(HASH_PLATFORM_SPLIT)

	hash.Write([]byte(strconv.FormatInt(time, 10)))
	hash.Write(HASH_PLATFORM_SPLIT)

	hash.Write([]byte(nick))
	hash.Write(HASH_PLATFORM_SPLIT)

	hash.Write([]byte(strconv.FormatInt(int64(roleId), 10)))

	return hash.Sum(nil)
}
