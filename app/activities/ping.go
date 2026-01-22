package activities

import (
	"context"
	"fmt"
	"net/http"
)

type PingHostInput struct {
	Host string
}

type PingHostOutput struct {
	Result string
}

func PingHost(ctx context.Context, input PingHostInput) (PingHostOutput, error) {
	res, err := http.Get(input.Host)
	if err != nil {
		return PingHostOutput{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PingHostOutput{Result: fmt.Sprintf("%s down", input.Host)}, err
	}

	return PingHostOutput{Result: fmt.Sprintf("%s up", input.Host)}, nil
}
