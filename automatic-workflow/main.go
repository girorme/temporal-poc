package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/girorme/temporal-poc/app"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	randomId := rand.Intn(100000)
	options := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("ping-workflow-%d", randomId),
		TaskQueue: "ping-tasks",
	}

	we, err := c.ExecuteWorkflow(
		context.Background(),
		options,
		app.Ping,
		"https://google.com",
	)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result app.PingWorkflowOutput
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
