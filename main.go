package main

import (
	"log"

	"github.com/girorme/temporal-poc/app"
	"github.com/girorme/temporal-poc/app/activities"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "ping-tasks", worker.Options{})
	w.RegisterWorkflow(app.Ping)
	w.RegisterActivity(activities.PingHost)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
