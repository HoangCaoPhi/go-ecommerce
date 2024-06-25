package settings

type Config struct {
	MySql  MySqlSettings `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySqlSettings struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	UserName           string `mapstructure:"userName"`
	Password           string `mapstructure:"password"`
	DbName             string `mapstructure:"dbName"`
	MaxIdleConnection  int    `mapstructure:"maxIdleConnection"`
	MaxOpenConnect     int    `mapstructure:"maxOpenConnect"`
	ConnectMaxLifeTime int    `mapstructure:"connectMaxLifeTime"`
}

type LoggerSetting struct {
	LoggerLevel string `mapstructure:"loggerLevel"`
	FileName    string `mapstructure:"fileName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
