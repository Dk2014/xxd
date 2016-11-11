package platform_server

import (
	"core/log"
	"core/redis"
	"errors"
	"strings"
	"time"
)

var redisPool_ *redis.Pool
var dbMap_ = make(map[string]int)

func initDBMap(apps map[string]*App) {
	if len(apps) > 0 {
		for key, app := range apps {
			dbMap_[key] = app.DBIndex
		}
	} else {
		dbMap_["xxd_qq"] = 0
	}
}

func RedisPool() *redis.Pool {
	return redisPool_
}

func InitRedis(flag string, apps map[string]*App) (err error) {
	//init db index map
	initDBMap(apps)

	arr := strings.Split(flag, ",")
	if len(arr) == 0 {
		return errors.New("redis options is empty.")
	}

	server := arr[0]
	redisPool_ = &redis.Pool{
		MaxIdle:     200,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}

			//password
			if len(arr) == 2 {
				password := arr[1]
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	_, err = redisPool_.Dial()
	return err
}

func GetDBConn(app string) redis.Conn {
	if redisPool_ == nil {
		return nil
	}

	conn := redisPool_.Get()
	dbindex := getDBIndex(app)
	_, err := conn.Do("SELECT", dbindex)
	if err != nil {
		log.Errorf("change db index error: %v", err)
	}

	return conn
}

func GetDBMap() map[string]int {
	return dbMap_
}

func getDBIndex(app string) int {
	if v, ok := dbMap_[app]; ok {
		return v
	} else {
		return 0
	}
}
