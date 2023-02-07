package configs

type DB struct {
	Driver string `mapstructure:"DATABASE_DRIVER" default:"mysql"`
	Mysql  Mysql  `mapstructure:",squash"`
}
type Mysql struct {
	Host     string `mapstructure:"MYSQL_HOST" default:"127.0.0.1"`
	Port     string `mapstructure:"MYSQL_PORT" default:"3306"`
	UserName string `mapstructure:"MYSQL_USER_NAME" default:"root"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	Database string `mapstructure:"MYSQL_DATABASE"`
}
