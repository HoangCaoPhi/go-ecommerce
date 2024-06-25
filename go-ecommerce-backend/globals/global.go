package globals

import (
	"github.com/HoangCaoPhi/go-ecommerce/pkg/loggers"
	"github.com/HoangCaoPhi/go-ecommerce/pkg/settings"
)

var (
	Config settings.Config
	Logger *loggers.LoggerZap
)
