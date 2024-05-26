package controllers

import (
	"log"
	// "os/exec"

	"github.com/robfig/cron/v3"
)

// myScheduledTask is the task that will be run every minute
func myScheduledTask() {

	// dumpPath := "./dump/"
	// uri := "mongodb://localhost:27017"
	// cmd := exec.Command("mongodump", "--uri", uri, "--out", dumpPath)

	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Data dumped successfully...")

}

// StartCronJob starts the cron job
func InitCronJob() {
	// Initialize a new cron instance
	c := cron.New(cron.WithSeconds())

	// Schedule the task to run every minute
	_, err := c.AddFunc("0 */1 * * * *", myScheduledTask)
	if err != nil {
		log.Fatalf("Error scheduling task: %v", err)
	}

	// // Funcs may also be added to a running Cron
	// c.AddFunc("@daily", func() { fmt.Println("Every day") })

	// Start the cron scheduler
	c.Start()

}
