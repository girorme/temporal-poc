package app

import (
	"time"

	"github.com/girorme/temporal-poc/app/activities"
	"go.temporal.io/sdk/workflow"
)

type PingWorkflowOutput struct {
	Result string
}

func Ping(ctx workflow.Context, input string) (PingWorkflowOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var hostStatus activities.PingHostOutput
	err := workflow.ExecuteActivity(ctx, activities.PingHost, activities.PingHostInput{Host: input}).Get(ctx, &hostStatus)
	if err != nil {
		return PingWorkflowOutput{}, err
	}

	return PingWorkflowOutput{Result: hostStatus.Result}, nil
}
