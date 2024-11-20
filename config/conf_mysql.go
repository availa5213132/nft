package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"` //服务器地址：端口
	Port         int    `yaml:"port"`
	Config       string `yaml:"config"` //高级配置
	DB           string `yaml:"db"`     //数据库名
	User         string `yaml:"user"`   //数据库用户名
	Password     string `yaml:"password"`
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` //空闲着中的最大连接数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` //打开到数据库的最大连接数
	LogLevel     string `yaml:"logLevel"`                             //日志等级  debug是输出全部sql，dev，release
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
