package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	host := "localhost"
	port := 6379
	connectTimeout := 3000
	readTimeout := 3000
	writeTimeout := 3000
	auth := "redispass"

	
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := redis.DialTimeout("tcp", address, 
			time.Duration(connectTimeout) * time.Millisecond,
			time.Duration(readTimeout) * time.Millisecond,
			time.Duration(writeTimeout) * time.Millisecond)
	if err != nil {
		log.Fatalf("redis connect failed, host:%s, port:%d, err:%s", host, port, err)
	}
	defer conn.Close()

	reply, err := conn.Do("auth", auth)
	if err != nil {
		log.Fatalf("redis auth failed, auth:%s, reply:%v, err:%v", auth, reply, err)
	}

	reply, err = conn.Do("get", "key")
	fmt.Printf("reply:%s, err:%s", reply, err)

}