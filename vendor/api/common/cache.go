/*
Package common - Cache file which is having all redis related functions
*/
package common

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/shopspring/decimal"
)

// CacheZADDWithFloat64 :
func CacheZADDWithFloat64(key string, values map[string]float64) error {
	c := Cache.Get()
	defer c.Close()
	for id, score := range values {
		val, _ := decimal.NewFromFloat(score).Round(2).Float64()
		val = val * -1
		if err := c.Send("ZADD", key, val, id); err != nil {
			continue
		}
	}
	if err := c.Flush(); err != nil {
		return err
	}
	return nil
}

// CacheZCOUNT :
func CacheZCOUNT(key string, min, max interface{}) (int, error) {
	c := Cache.Get()
	count, _ := redis.Int(c.Do("ZCOUNT", key, min, max))
	defer c.Close()
	return count, nil
}

// CacheZRANGEBYSCORE :
func CacheZRANGEBYSCORE(key string, min, max float64, offset, limit int, reverse bool) ([]UserRank, error) {
	results := []UserRank{}
	c := Cache.Get()

	v, err := redis.Values(c.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset-1, limit))
	if err != nil || v == nil {
		fmt.Println(err)
		return results, err
	}
	defer c.Close()

	var prevscore float64
	var prevrank int
	var userstr string
	prevscore = -9999

	for len(v) > 0 {
		a := UserRank{}
		v, err = redis.Scan(v, &userstr, &a.Score)
		uservals := strings.Split(userstr, "|")
		if len(uservals) > 1 {
			a.NickName = uservals[0]
			a.UserID = uservals[1]
		} else {
			a.NickName = uservals[0]
		}

		if a.Score != prevscore {
			r1, _ := CacheZCOUNT(key, "-inf", a.Score)
			r2, _ := CacheZCOUNT(key, a.Score, a.Score)
			a.Rank = r1 - r2 + 1
			prevrank = a.Rank
			prevscore = a.Score
			if reverse {
				a.Score = a.Score * -1
			}

			results = append(results, a)
		} else {
			// Same rank
			a.Rank = prevrank
			if reverse {
				a.Score = a.Score * -1
			}

			results = append(results, a)
		}
	}

	return results, nil
}
