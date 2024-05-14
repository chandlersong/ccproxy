package binance

import (
	"context"
	"fmt"
	"github.com/binance/binance-connector-go"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"testing"
)

func TestBinanceAPI(t *testing.T) {
	//完全是测试Binance官方的API

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	apiKey := viper.GetString("binance_api_key")
	secretKey := viper.GetString("binance_secret_key")
	proxyUrl, err := url.Parse("http://localhost:7890")
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	binanceClient := binance_connector.NewClient(apiKey, secretKey)
	// Debug Mode
	binanceClient.Debug = true

	// TimeOffset (in milliseconds) - used to adjust the request timestamp by subtracting/adding the current time with it:
	binanceClient.TimeOffset = -1000 // implies adding: request timestamp = current time - (-1000)
	binanceClient.HTTPClient = myClient

	t.Run("test ping with proxy", func(t *testing.T) {
		// NewPingService
		ping := binanceClient.NewPingService().Do(context.Background())
		if ping == nil {
			fmt.Println("Success")
			return
		}
		fmt.Println(binance_connector.PrettyPrint(ping))
	})

	t.Run("test cn url", func(t *testing.T) {
		client := NewClientWithBase(cnBinanceUrl)
		client.ping()
	})

}
