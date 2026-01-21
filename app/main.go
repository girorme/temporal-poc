package app

import "go.temporal.io/sdk/workflow"

func Greet(ctx workflow.Context, name string) (string, error) {
	return "Hello, " + name + "!", nil
}
