package binance

import "testing"

func TestPing(t *testing.T) {
	// Test case 1: Sufficient funds
	client := Client{}
	t.Run("hello world", func(t *testing.T) {
		client.ping()
	})

}