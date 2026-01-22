package main

import (
	"context"
	"log"

	"github.com/girorme/temporal-poc/app"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "greeting-tasks",
	}

	we, err := c.ExecuteWorkflow(
		context.Background(),
		options,
		app.Greet,
		"Girorme",
	)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
