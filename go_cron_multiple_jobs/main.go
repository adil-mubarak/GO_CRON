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
		fmt.Println("Running job every minute:",time.Now())
	})

	if err != nil{
		log.Fatalf("Error adding cron job: %v",err)
	}

	_,err = c.AddFunc("0 12 * * *",func() {
		fmt.Println("Running daily job at 12 pm:",time.Now())
	})

	if err != nil{
		log.Fatalf("Error adding cron job: %v",err)
	}

	c.Start()

	select {}
}