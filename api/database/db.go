package database

//初始化数据库连接
import (
	"fmt"
	"gin_test/api/config"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var (
	engines   = make(map[string]*xorm.Engine)
	enginesMu sync.RWMutex
)

// GetDb 获取指定数据库的引擎
// @param dbName 数据库名称 例:db1.name
// @return *xorm.Engine 数据库引擎
func GetDb(dbName string) *xorm.Engine {

	enginesMu.RLock()
	if engine, ok := engines[dbName]; ok {
		enginesMu.RUnlock()
		return engine
	}
	enginesMu.RUnlock()

	enginesMu.Lock()
	defer enginesMu.Unlock()

	engine, err := initDatabase(dbName)
	if err != nil {
		panic(err)
	}
	engines[dbName] = engine
	return engine
}

// GetDbSession 获取指定数据库的会话
func GetDbSession(dbName string) *xorm.Session {
	return GetDb(dbName).NewSession()
}

// initDatabase 初始化数据库连接
func initDatabase(dbName string) (*xorm.Engine, error) {
	user, _ := config.GetConfigString("DB.USER")
	password, _ := config.GetConfigString("DB.PASSWORD")
	host, _ := config.GetConfigString("DB.HOST")
	port, _ := config.GetConfigInt("DB.PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	driver, _ := config.GetConfigString("DB.DRIVER")
	engine, err := xorm.NewEngine(driver, dsn)
	// 如果需要设置连接池的空闲数大小，可以使用 engine.SetMaxIdleConns() 来实现。
	// 如果需要设置最大打开连接数，则可以使用 engine.SetMaxOpenConns() 来实现。
	// 如果需要设置连接的最大生存时间，则可以使用 engine.SetConnMaxLifetime() 来实现。
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.SetConnMaxLifetime(time.Duration(10 * 60))
	if err != nil {
		return nil, err
	}
	env, _ := config.GetConfigString("APP.ENV")
	if env != "pro" {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)
	}

	if err = engine.Ping(); err != nil {
		return nil, err
	}
	return engine, nil
}

// CloseAll 关闭所有数据库连接
func CloseAll() {
	enginesMu.Lock()
	defer enginesMu.Unlock()
	for _, engine := range engines {
		engine.Close()
	}
}
