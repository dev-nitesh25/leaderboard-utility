/*
Package common - Cache file which is having all redis related functions
*/
package common

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Cache provides the connection to the main data cache server.
var Cache *redis.Pool

// DBCon is the pointer to the database connection resource.
var DBCon *sql.DB

// InitDB initialises the database pools with
func InitDB(host, port, user, password string) (dbcon *sql.DB, err error) {
	DBCon, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/lbdb")

	if err != nil {
		return nil, errors.New("DB Connect error :" + err.Error())
	}
	if err = DBCon.Ping(); err != nil {
		return nil, errors.New("DB Connect error :" + err.Error())
	}

	return DBCon, nil
}

// InitCache sets up the two cache pools for the entire API.
func InitCache(cacheHost, cachePort string) error {
	fmt.Println("Connecting to cache server - ", cacheHost)
	Cache = &redis.Pool{
		MaxIdle:     50,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cacheHost+":"+cachePort)
			if err != nil {
				return nil, errors.New("Cache error : " + err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Wait: true,
	}

	return nil
}
