package app

import (
	"time"

	"github.com/girorme/temporal-poc/app/activities"
	"go.temporal.io/sdk/workflow"
)

func Greet(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var googleStatus string
	err := workflow.ExecuteActivity(ctx, activities.GetOnGoogle).Get(ctx, &googleStatus)
	if err != nil {
		return "", err
	}

	return googleStatus, nil
}
