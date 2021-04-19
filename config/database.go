package config

type DBOptional struct {
	DriverName   string
	User         string
	Password     string
	DBHostname   string
	DBPort       string
	DBName       string
	DBCharset    string
	Timeout      string
	ReadTimeout  string
	WriteTimeout string
}

func GetDefaultDBOptional() *DBOptional {
	return &DBOptional{
		DriverName:   MysqlSetting.DriverName,
		User:         MysqlSetting.User,
		Password:     MysqlSetting.Password,
		DBHostname:   MysqlSetting.DBHostname,
		DBPort:       MysqlSetting.DBPort,
		DBName:       MysqlSetting.DBName,
		DBCharset:    "utf8",
		Timeout:      "1000ms",
		ReadTimeout:  "2.0s",
		WriteTimeout: "5.0s",
	}
}
