package initialize

import (
	"fmt"

	"github.com/HoangCaoPhi/go-ecommerce/globals"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()

	fmt.Println("Load configuration host mysql", globals.Config.MySql.Host)

	InitLogger()

	globals.Logger.Info("ok", zap.String("oke", "success"))
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()

	r.Run(":8002")
}
