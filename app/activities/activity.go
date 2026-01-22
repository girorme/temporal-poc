package activities

import (
	"context"
	"net/http"
)

func GetOnGoogle(ctx context.Context) (string, error) {
	res, err := http.Get("https://google.com")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "google down", err
	}

	return "google up", nil
}
