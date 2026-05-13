package config

import (
	"fmt"
	"gin_test/gerrors"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	globalConfig map[string]any
	configOnce   sync.Once
)

// GetConfigInt 获取指定配置项的整数值
func GetConfigInt(name string) (val int, err error) {
	val = viper.GetInt(name)
	if !viper.IsSet(name) {
		err = gerrors.ConfigNotExist
	}
	return
}

// GetConfigString 获取指定配置项的字符串值
func GetConfigString(name string) (val string, err error) {
	val = viper.GetString(name)
	if !viper.IsSet(name) {
		return val, gerrors.ConfigNotExist
	}
	return
}

// GetConfigMap 获取指定配置项的映射
func GetConfigMap(name string) (val map[string]any, err error) {
	if name == "" {
		return globalConfig, nil
	}
	val = viper.GetStringMap(name)
	if !viper.IsSet(name) {
		return val, gerrors.ConfigNotExist
	}
	return
}

// init 初始化配置
func init() {
	configOnce.Do(func() {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("读取配置文件失败: %v", err))
		}

		// 将 viper 中所有的配置项，直接 Unmarshal 到一个通用的 map 中
		globalConfig = viper.AllSettings()
		fmt.Printf("配置文件加载成功: %+v\n", globalConfig)
	})
}
