package cron

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	cronApp *cron.Cron
	once    sync.Once
)

func ProvideCron() *cron.Cron {
	once.Do(func() {
		fmt.Println("Initializing Cron")

		indLocation, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			log.Fatalf("failed to load location: %v", err)
		}

		cronApp = cron.New(cron.WithLocation(indLocation), cron.WithSeconds())
	})
	return cronApp
}

func Get() *cron.Cron {
	return ProvideCron()
}
