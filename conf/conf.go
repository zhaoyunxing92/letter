package conf

import (
	"github.com/spf13/viper"
	"log"
)

var MongodbConfig = new(DataBase)

var ApplicationConfig = new(Application)

var LoggerConfig = new(Logger)

//数据配置
type DataBase struct {
	Name        string //数据库名称
	Url         string //应用名称
	Username    string //用户名
	Password    string //密码
	Timeout     uint64 //链接超时时间,单位秒
	MaxPoolSize uint64 //链接超时时间,单位秒
	MinPoolSize uint64 //链接超时时间,单位秒
}

// 应用信息
type Application struct {
	Mode string //模式
	Name string //应用名称
}

//日志
type Logger struct {
	Path  string //路径
	Level string //日志级别
}

func Setup(path string) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read [%s] config file fail: %s", path, err.Error())
	}
	// 应用信息
	application := viper.Sub("settings.application")
	if application == nil {
		panic("No found settings.application in the configuration")
	}
	ApplicationConfig = initApplication(application)

	// 数据库配置
	database := viper.Sub("settings.database")
	if database == nil {
		panic("No found settings.database in the configuration")
	}
	MongodbConfig = initDataBase(database)
	//logger
	logger := viper.Sub("settings.logger")
	if logger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = initLogger(logger)
}

//初始化数据库
func initDataBase(cfg *viper.Viper) *DataBase {

	cfg.SetDefault("max-pool-size", 100)
	cfg.SetDefault("min-pool-size", 5)

	return &DataBase{
		Name:        cfg.GetString("name"),
		Url:         cfg.GetString("url"),
		Username:    cfg.GetString("username"),
		Password:    cfg.GetString("password"),
		Timeout:     cfg.GetUint64("timeout"),
		MaxPoolSize: cfg.GetUint64("max-pool-size"),
		MinPoolSize: cfg.GetUint64("min-pool-size"),
	}
}

func initApplication(cfg *viper.Viper) *Application {

	return &Application{
		Name: cfg.GetString("name"),
		Mode: cfg.GetString("mode"),
	}
}

func initLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		Path: cfg.GetString("path"),
		Level: cfg.GetString("level"),
	}
}
