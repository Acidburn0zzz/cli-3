package svc

import (
	"fmt"
	"net/http"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/gqlclient"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/machinebox/graphql"
)

type Client struct {
	*gqlclient.Client
	baseUrl string
}

// New will create a new API client using default settings (for an authenticated version use the NewWithAuth version)
func New(cfg *config.Instance) (*Client, error) {
	port := cfg.GetInt(constants.SvcConfigPort)
	if port <= 0 {
		return nil, locale.NewError("err_svc_no_port", "The State Tool service does not appear to be running (no local port was configured).")
	}
	baseUrl := fmt.Sprintf("http://127.0.0.1:%d", port)
	return &Client{
		Client:  gqlclient.New(fmt.Sprintf("%s/query", baseUrl), 0),
		baseUrl: baseUrl,
	}, nil
}

// NewWithoutRetry returns a client based on an HTTP client that does not retry on failure
func NewWithoutRetry(cfg *config.Instance) (*Client, error) {
	port := cfg.GetInt(constants.SvcConfigPort)
	if port <= 0 {
		return nil, locale.NewError("err_svc_no_port", "The State Tool service does not appear to be running (no local port was configured).")
	}

	baseUrl := fmt.Sprintf("http://127.0.0.1:%d", port)
	return &Client{
		Client:  gqlclient.NewWithOpts(fmt.Sprintf("%s/query", baseUrl), 0, graphql.WithHTTPClient(http.DefaultClient)),
		baseUrl: baseUrl,
	}, nil
}

func (c *Client) BaseUrl() string {
	return c.baseUrl
}
