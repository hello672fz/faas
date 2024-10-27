package policies

import (
	"math/rand"
	"net/url"
	"time"
)

// TODO: Policy
func SelectNodeByPolicy(urls []*url.URL) int {
	if len(urls) == 0 {
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(urls))

	return randomIndex
}
