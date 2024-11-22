package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	_,err := c.AddFunc("* * * * *",func() {
		fmt.Println("Cron job executed at:",time.Now())
	})
	if err != nil{
		log.Fatalf("Error adding cron job: %v",err)
	}

	c.Start()

	select {}
}
