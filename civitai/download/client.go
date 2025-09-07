package download

import (
	"github.com/Simmons135/civitai-cli/civitai/api"
)

type Client struct {
	api *api.CivitaiClient
}

func NewClient(api *api.CivitaiClient) *Client {
	return &Client{
		api: api,
	}
}

func DefaultClient() *Client {
	return NewClient(api.NewClient(api.GetAPIToken()))
}
// API returns the underlying Civitai API client.
func (c *Client) API() *api.CivitaiClient {
	return c.api
}
