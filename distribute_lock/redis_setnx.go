package distribute_lock

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

func incr() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.50:30477",
		Password: "A9bXYIoC0M",
	})

	var lockKey = "counter_lock"
	var counterKey = "counter"

	resp := client.SetNX(lockKey, 1, 5*time.Second)
	ok, err := resp.Result()
	log.Println(err, "try lock result: ", ok)

	if err != nil || !ok {
		return
	}

	getResp := client.Get(counterKey)
	cnt, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cnt++
		resp := client.Set(counterKey, cnt, time.Minute)
		if _, err := resp.Result(); err != nil {
			log.Println("set value err")
		}
	}

	log.Println("current counter is: ", cnt)

	delResp := client.Del(lockKey)
	if unlockOk, err := delResp.Result(); err == nil || unlockOk > 0 {
		log.Println("unlock ok")
	} else {
		log.Println("fail to unlock")
	}
}
