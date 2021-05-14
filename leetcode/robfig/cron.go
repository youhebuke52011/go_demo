package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	fmt.Println("---->>",time.Now().Format("2006-01-02 15:04:05"))
	c.AddFunc("* * * * * *", func() {
		time.Sleep(5 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
	time.Sleep(time.Minute)
}
