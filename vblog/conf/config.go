package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "vblog",
			Username: "vblog",
			Password: "123456",
		},
		App: &App{
			HttpHost: "127.0.0.1",
			HttpPort: 7080,
		},
	}
}

// 这个对象维护整个程序配置
type Config struct {
	MySQL *MySQL `json:"mysql" toml:"mysql"`
	App   *App   `json:"app" toml:"app"`
}

func (c *Config) String() string {
	dj, _ := json.Marshal(c)
	return string(dj)
}

type App struct {
	HttpHost string `json:"http_host" env:"HTTP_HOST"`
	HttpPort int64  `json:"http_port" env:"HTTP_PORT"`
}

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%d", a.HttpHost, a.HttpPort)
}

// [mysql]
// host="localhost"
// port=3306
// database="demo"
// username="demo"
// password="demo"
type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"database" toml:"database" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

	// 缓存一个对象
	lock sync.Mutex
	conn *gorm.DB
}

// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 返回一个数据库链接, 返回一个全局单列
func (m *MySQL) GetConn() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 没有就赋值
	if m.conn == nil {
		// 在进行m.conn = conn 赋值操作时 由锁存在不会冲突
		// https://gorm.io/zh_CN/docs/index.html
		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.conn = conn
	}

	return m.conn
}
