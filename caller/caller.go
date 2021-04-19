package caller

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/nju-iot/edgex_admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	RedisClient *redis.Client
	EdgexDB     *gorm.DB
)

func InitClient() {
	initRedisClient()
	initMysqlClient()
}

func initRedisClient() {
	redisOpt := &redis.Options{
		Addr:     config.RedisSetting.Address,
		Password: config.RedisSetting.Password,
		DB:       config.RedisSetting.DB,
	}
	RedisClient = redis.NewClient(redisOpt)
}

func initMysqlClient() {

	optional := config.GetDefaultDBOptional()

	format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s"
	dbConfig := fmt.Sprintf(format, optional.User, optional.Password, optional.DBHostname, optional.DBPort,
		optional.DBName, optional.DBCharset, optional.Timeout, optional.ReadTimeout, optional.WriteTimeout)

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var err error
	EdgexDB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dbConfig,
	}), &gormConfig)

	if err != nil {
		panic(err)
	}
}
