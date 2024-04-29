package binance

import (
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/proxy"
)

var (
	baseURL = "https://api.binance.com"
)

type Client struct {
	baseUrl  string
	proxyUrl string
}

func (p *Client) ping() {
	// Create a new client
	cli := gentleman.New()
	cli.URL(baseURL)
	// Define the default HTTP transport

	p.addProxyIfNecessary(cli)
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

func (p *Client) addProxyIfNecessary(cli *gentleman.Client) bool {

	hasProxy := p.proxyUrl != ""
	if hasProxy {
		servers := map[string]string{"http": p.proxyUrl, "https": p.proxyUrl}
		cli.Use(proxy.Set(servers))
	}

	return hasProxy

}

func NewClientWithProxy(proxyUrl string) *Client {
	return &Client{baseUrl: baseURL, proxyUrl: proxyUrl}
}

func NewClientWithBase(base string) *Client {
	return &Client{baseUrl: base}
}
