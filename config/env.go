package config

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

var once sync.Once

// 配置
func Get[T any](key string, defaultValue T) T {
	once.Do(func() {
		loadConfigFile()
	})
	keyArr := strings.Split(key, ".")
	config := viper.AllSettings()
	group, ok := config[keyArr[0]]
	if !ok {
		fmt.Println("group:", key, "not found")
		return defaultValue
	}
	item, ok := (group.(map[string]interface{}))[keyArr[1]]
	if !ok {
		fmt.Println("config item:", key, "not found")
		return defaultValue
	}
	// 类型断言并尝试转换
	switch any(defaultValue).(type) {
	case string:
		if str, ok := item.(string); ok {
			return any(str).(T)
		}
	case int:
		if str, ok := item.(string); ok {
			if num, err := strconv.Atoi(str); err == nil {
				return any(num).(T)
			}
		}
	}
	return item.(T)
}
func loadConfigFile() {
	// 加载 .ini 文件
	cfg, err := ini.Load(".ini")
	if err != nil {
		log.Fatalf("Failed to read .ini file: %v", err)
	}

	// 将 .ini 文件内容注入 viper
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			viper.Set(section.Name()+"."+key.Name(), key.Value())
		}
	}

	// fmt.Println("Loaded config:", viper.AllSettings())
}
