package controllers

import (
	"log"

	"github.com/robfig/cron/v3"
)

// myScheduledTask is the task that will be run every minute
func myScheduledTask() {
	log.Println("This task is run every minute")
}



// StartCronJob starts the cron job
func init() {
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
