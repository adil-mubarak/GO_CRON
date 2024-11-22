package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	logger := log.New(os.Stdout,"cron: ",log.LstdFlags)

	c := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(logger)))

	_,err := c.AddFunc("* * * * *",func() {
		if time.Now().Second()%2 == 0{
			fmt.Println("Job succeeded at:",time.Now())
		}else{
			log.Println("Job failed at:",time.Now())
		}
	})

	if err != nil{
		log.Fatalf("Error adding cron job: %v",err)
	}

	c.Start()

	select {}
}
