package Cache

import (
	"github.com/gomodule/redigo/redis"
)

var Cache redis.Conn

func init() {
	initCache()
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")

	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
}
