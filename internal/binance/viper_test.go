package binance

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestViper(t *testing.T) {
	//Viper的测试

	t.Run("read a value", func(t *testing.T) {
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Println(viper.Get("test"))
	})
}
