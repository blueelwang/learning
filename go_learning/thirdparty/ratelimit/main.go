package main

import (
	"time"

	"github.com/juju/ratelimit"
)

func main() {
	ratelimit.NewBucket(time.Second * 5, 1)
	
}