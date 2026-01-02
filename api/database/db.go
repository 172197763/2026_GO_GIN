package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/go-sql-driver/mysql" // 匿名导入
	"github.com/joho/godotenv"
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"
)

var engine *xorm.Engine
var session *xorm.Session

// InitDatabase 初始化数据库连接
func InitDatabase() error {
	fpath := filepath.Join(GetRootDir(), ".env")
	fmt.Println("读取", fpath)
	// 检查文件是否存在
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		log.Fatal("Error: .env file does not exist at path:", fpath)
	}
	// 加载 .env 文件
	err := godotenv.Load(filepath.Join(GetRootDir(), ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 从环境变量读取配置
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	// 构建数据源名称
	var dataSourceName string
	if driver == "mysql" {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, dbname)
	}

	// 创建数据库引擎
	engine, err = xorm.NewEngine(driver, dataSourceName)
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
