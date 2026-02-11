package infra

import (
	"fmt"
	"strconv"
	"time"

	tb "github.com/tigerbeetle/tigerbeetle-go"
	"github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

const healthCheckTimeout = 5 * time.Second

// Client wraps the TigerBeetle Go client with a health-check on connect.
type Client struct {
	raw tb.Client
}

// Connect creates a TigerBeetle client and verifies connectivity with a
// health-check query. NewClient itself retries in the background and won't
// fail immediately when the server is down, so we run a QueryAccounts(Limit:1)
// behind a timeout to confirm real connectivity.
func Connect(clusterID string, addresses []string) (*Client, error) {
	id, err := parseClusterID(clusterID)
	if err != nil {
		return nil, fmt.Errorf("invalid cluster ID %q: %w", clusterID, err)
	}

	raw, err := tb.NewClient(id, addresses)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	// Health check with timeout — the client retries indefinitely,
	// so we need an external deadline.
	type result struct{ err error }
	ch := make(chan result, 1)
	go func() {
		_, qErr := raw.QueryAccounts(types.QueryFilter{Limit: 1})
		ch <- result{err: qErr}
	}()

	select {
	case r := <-ch:
		if r.err != nil {
			raw.Close()
			return nil, fmt.Errorf("health check failed: %w", r.err)
		}
	case <-time.After(healthCheckTimeout):
		raw.Close()
		return nil, fmt.Errorf("connection timed out after %s", healthCheckTimeout)
	}

	return &Client{raw: raw}, nil
}

// Close closes the underlying TigerBeetle client.
func (c *Client) Close() {
	if c != nil && c.raw != nil {
		c.raw.Close()
	}
}

// Raw returns the underlying TigerBeetle client for direct queries.
func (c *Client) Raw() tb.Client {
	return c.raw
}

func parseClusterID(s string) (types.Uint128, error) {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return types.Uint128{}, err
	}
	return types.ToUint128(n), nil
}
