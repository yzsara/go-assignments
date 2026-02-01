package main

import (
	"time"
)

func main(){
	allow := tokenBucketLimiter(5, 10) // 5 req/sec, burst up to 10

	for i := 1; i <= 15; i++ {
		ok := allow()
		if ok {
			println("request", i, "allowed")
		} else {
			println("request", i, "blocked")
		}
	}
}

type AllowFunc func() (allowed bool)

func tokenBucketLimiter(ratePerSecond float64, burst int) AllowFunc{
	startTime := time.Now()
	
	tokens := float64(burst)

	return func()bool{
		now := time.Now()
		elapsed := now.Sub(startTime)

		tokens += elapsed.Seconds() * ratePerSecond
		if tokens > float64(burst) {
			tokens = float64(burst)
		}

		if(tokens >= 1){
			tokens --
			return true
		}

		return false
	}
}