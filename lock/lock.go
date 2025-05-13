package lock

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func acquireLock(client *redis.Client, lockKey string, timeout time.Duration) bool {
	// Try to acquire the lock with SETNX command (SET if Not eXists)
	lockAcquired, err := client.SetNX(lockKey, 1, timeout).Result()
	if err != nil {
		fmt.Println("Error acquiring lock: ", err)
		return false
	}
	return lockAcquired
}

func releaseLock(client *redis.Client, lockKey string) {
	client.Del(lockKey)
}

func Lock() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer client.Close()

	// Define the lock key and lock timeout
	lockKey := "my_lock"
	lockTimeout := 20 * time.Second

	// Acquire the lock
	if acquireLock(client, lockKey, lockTimeout) {
		fmt.Println("Lock acquired successfully!")
		// simulate some work withh the lock
		time.Sleep(20 * time.Second)
		fmt.Println("Work done!")

		// release the lock
		releaseLock(client, lockKey)
		fmt.Println("Lock released successfully!")
	} else {
		fmt.Println("Failed to acquire lock. Resource is already locked.")
	}
}
