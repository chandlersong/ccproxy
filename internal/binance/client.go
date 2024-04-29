package binance

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/proxy"
)

var (
	baseURL  = "https://api.binance.com"
	proxyURL = "http://localhost:7890"
)

type Client struct {
}

func (p *Client) ping() {
	// Create a new client
	cli := gentleman.New()
	cli.URL(baseURL)
	// Define the default HTTP transport

	servers := map[string]string{"http": proxyURL, "https": proxyURL}
	cli.Use(proxy.Set(servers))
	// Perform the request
	req := cli.Request()
	req.Path("/api/v3/ping")
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}
	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}

func NewClient() *Client {
	return &Client{}
}
