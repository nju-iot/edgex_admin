package config

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

var (
	ServerSetting = serverSetting{}
	LogSetting    = logSetting{}
	MysqlSetting  = mysqlSetting{}
	MongoSetting  = mongoSetting{}
	RedisSetting  = redisSetting{}
)

type serverSetting struct {
	RunMode  string
	HttpPort int
	Port     string
}

type logSetting struct {
	LogLevel   string
	FileName   string // 日志文件名
	MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    // 日志文件最多保存多少个备份
	MaxAge     int    // 文件最多保存多少天
	Compress   bool   // 日志是否压缩
}

type mysqlSetting struct {
	DriverName string
	User       string
	Password   string
	DBHostname string
	DBPort     string
	DBName     string
}

type mongoSetting struct {
}

type redisSetting struct {
	Address  string
	Password string
	DB       int
}

func InitConfig() {

	cfg, err := ini.Load("/Users/bytedance/go/src/github.com/nju-iot/edgex_admin/app.ini")
	if err != nil {
		panic(err)
	}

	mapTo("Log", &LogSetting, cfg)
	mapTo("Mongo", &MongoSetting, cfg)
	mapTo("Mysql", &MysqlSetting, cfg)
	mapTo("Redis", &RedisSetting, cfg)
	mapTo("Server", &ServerSetting, cfg)

	if ServerSetting.HttpPort != 0 {
		ServerSetting.Port = fmt.Sprintf(":%d", ServerSetting.HttpPort)
	}
}

func mapTo(section string, v interface{}, cfg *ini.File) {
	if cfg == nil || section == "" {
		log.Fatalf("section=%v, iniFile=%v", section, cfg)
		return
	}
	if err := cfg.Section(section).MapTo(v); err != nil {
		panic(err)
	}
}
