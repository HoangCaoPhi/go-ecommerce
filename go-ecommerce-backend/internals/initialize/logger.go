package initialize

import (
	"github.com/HoangCaoPhi/go-ecommerce/globals"
	"github.com/HoangCaoPhi/go-ecommerce/pkg/loggers"
)

func InitLogger() {
	globals.Logger = loggers.NewLogger(globals.Config.Logger)
}
