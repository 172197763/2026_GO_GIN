package database

import (
	"fmt"
	"gin_test/config"
	"path/filepath"
	"runtime"

	_ "github.com/go-sql-driver/mysql" // 匿名导入
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
)

var engine *xorm.Engine
var session *xorm.Session

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	// 从环境变量读取配置
	host := config.Get("mysql.host", "")
	port := config.Get("mysql.port", 3306)
	user := config.Get("mysql.user", "")
	password := config.Get("mysql.password", "")
	dbname := config.Get("mysql.dbname", "product")
	driver := config.Get("mysql.driver", "mysql")

	// 构建数据源名称
	var dataSourceName string
	if driver == "mysql" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, dbname)
	}

	// 创建数据库引擎
	engine, err := xorm.NewEngine(driver, dataSourceName)
	if err != nil {
		return err
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(xormlog.LOG_DEBUG) // 已修正

	// 测试连接
	if err = engine.Ping(); err != nil {
		return err
	}
	session = engine.NewSession()
	return nil
}
func GetRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "../..")
}
func GetDbEngine() *xorm.Engine {
	return engine
}
func GetDbSession() *xorm.Session {
	return session
}
