package model

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

var MyUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}

	return
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

func (this *UserDao) Login(username, password string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.Get(conn, username)
	if err != nil {
		return
	}

	if user.PassWord != password {
		err = ERROR_PWD_FAILED
		return
	}

	return
}
