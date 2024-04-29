package binance

import (
	"testing"
)

const proxyURL = "http://localhost:7890"
const cnBinanceUrl = "binancezh.com"

func TestPing(t *testing.T) {
	// Test case 1: Sufficient funds

	t.Run("test ping with proxy", func(t *testing.T) {
		client := NewClientWithProxy(proxyURL)
		client.ping()
	})

	t.Run("test cn url", func(t *testing.T) {
		client := NewClientWithBase(cnBinanceUrl)
		client.ping()
	})

}
