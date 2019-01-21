package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) Get(conn redis.Conn, username string) (user *User, err error)  {
	rsp, err := redis.String(conn.Do("HGet", "users", username))
	if err == redis.ErrNil {
		err = ERROR_USER_NOTEXISTS
	}
	user = &User{}
	err = json.Unmarshal([]byte(rsp), user)
	if err != nil {
		return
	}

	return 
}
