package main

import (
	"core/log"
	"core/redis"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"platform_server"
	"time"
)

const (
	REDIS_KEY_HASH_ACCOUNT = "key_hash_account"
	REDIS_KEY_HASH_SESSION = "key_hash_session"
	ACTION_LOGIN           = "/login"

	PASSWORD_SALT = "xinshuishendong"

	REGISTER_PRIVATE_KEY = "hechuluoqingchun"
)

type Account struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type Session struct {
	Id         string
	User       string
	LastUpdate string
}

func login(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		httpResponse(resp, 10001, nil, "body error")
		return
	}

	reqAccount := new(Account)
	err = json.Unmarshal(body, &reqAccount)
	if err != nil {
		httpResponse(resp, 10002, nil, "body must be json string")
		return
	}

	//index 0 db
	client := platform_server.GetDBConn("global")
	defer client.Close()

	val, err := redis.String(client.Do("HGET", REDIS_KEY_HASH_ACCOUNT, reqAccount.User))
	if err != nil || val == "" {
		httpResponse(resp, 10003, nil, "the user not exist")
		return
	}

	dbAccount := new(Account)
	err = json.Unmarshal([]byte(val), &dbAccount)
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	if dbAccount.Password != md5Hash(reqAccount.Password) {
		httpResponse(resp, 10005, nil, "the password is wrong")
		return
	}

	_updateSession(client, resp, reqAccount.User)

	httpResponse(resp, 0, nil, "ok")
}

func md5Hash(str string) string {
	hash := md5.New()
	hash.Write([]byte(str + "-" + PASSWORD_SALT))

	return hex.EncodeToString(hash.Sum(nil))
}

func register(resp http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("private_key")
	if url != REGISTER_PRIVATE_KEY {
		resp.WriteHeader(http.StatusForbidden)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		httpResponse(resp, 10001, nil, "body error")
		return
	}

	reqAccount := new(Account)
	err = json.Unmarshal(body, &reqAccount)
	if err != nil {
		httpResponse(resp, 10002, nil, "body must be json string")
		return
	}

	reqAccount.Password = md5Hash(reqAccount.Password)
	//index 0 db
	client := platform_server.GetDBConn("global")
	defer client.Close()

	b, err := json.Marshal(reqAccount)
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	_, err = client.Do("HSET", REDIS_KEY_HASH_ACCOUNT, reqAccount.User, string(b))
	if err != nil {
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	_updateSession(client, resp, reqAccount.User)
	httpResponse(resp, 0, nil, "ok")
}

func _updateSession(client redis.Conn, resp http.ResponseWriter, user string) {
	_uuid, _ := NewV4()
	id := _uuid.String()

	session_cookie := &http.Cookie{
		Name:    "session_id",
		Value:   id,
		Path:    "/",
		Expires: time.Now().AddDate(0, 1, 0),
	}
	user_cookie := &http.Cookie{
		Name:    "user",
		Value:   user,
		Path:    "/",
		Expires: time.Now().AddDate(0, 1, 0),
	}

	http.SetCookie(resp, session_cookie)
	http.SetCookie(resp, user_cookie)

	session := &Session{
		Id:         id,
		User:       user,
		LastUpdate: time.Now().Local().String(),
	}

	b, err := json.Marshal(session)
	_, err = client.Do("HSET", REDIS_KEY_HASH_SESSION, id, string(b))
	if err != nil {
		log.Error("set chookie error: " + err.Error())
	}
}

func logout(resp http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session_id")
	if err != nil {
		resp.WriteHeader(http.StatusNotAcceptable)
		return
	}

	//index 0 db
	client := platform_server.GetDBConn("global")
	defer client.Close()
	_, err = client.Do("HDEL", REDIS_KEY_HASH_SESSION, cookie.Value)
	if err != nil {
		log.Error("set chookie error: " + err.Error())
		httpResponse(resp, 10004, nil, "server inner error")
		return
	}

	http.Redirect(resp, req, "/login", http.StatusSeeOther)
}

func check_cookie(req *http.Request) bool {
	cookie, err := req.Cookie("session_id")
	if err != nil {
		return false
	}

	//index 0 db
	client := platform_server.GetDBConn("global")
	defer client.Close()

	b, err := redis.Bytes(client.Do("HGET", REDIS_KEY_HASH_SESSION, cookie.Value))
	if err != nil || len(b) == 0 {
		return false
	}

	s := new(Session)
	json.Unmarshal(b, s)

	s.LastUpdate = time.Now().String()
	b, _ = json.Marshal(s)
	_, err = client.Do("HSET", REDIS_KEY_HASH_SESSION, cookie.Value, string(b))
	if err != nil {
		log.Error("set chookie error: " + err.Error())
	}

	return true
}
