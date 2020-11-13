package main

import (
	"time"

	"github.com/juju/ratelimit"
)

func main() {
	bucket := ratelimit.NewBucket(time.Second * 5, 1)
	bucket.Available()
	
}