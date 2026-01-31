package cmd

import (
	"errors"
	"fmt"
	"gin_test/api/common"

	"github.com/spf13/cobra"
)

// 定义包级别的哨兵错误
var (
	ErrUserNotLogin = errors.New("用户未登录")
	ErrParamsError  = errors.New("参数错误")
	ErrTimeout      = errors.New("请求超时")
)

var errorCmd = &cobra.Command{
	Use:   "errorCmd",
	Short: "错误展示",
	Long:  `错误展示`,
	Run: func(cmd *cobra.Command, args []string) {
		errHandle()
	},
}

func errHandle() {
	err := checkUserToken()
	if err != nil {
		fmt.Println(err)
		//检测错误是否包含特定的哨兵错误
		fmt.Println(errors.Is(err, ErrUserNotLogin))
		fmt.Println(errors.Is(err, ErrParamsError))
	}
}
func checkUserToken() error {
	//具体检测用户登录状态
	isLogin := false
	randNum := common.RandInt(0, 100)
	fmt.Println(randNum)
	isExpired := false
	if randNum > 50 {
		isExpired = true
	}
	isDel := !isExpired
	if !isLogin {
		if isExpired {
			return fmt.Errorf("%w:用户token失效", ErrUserNotLogin)
		}
		if isDel {
			return fmt.Errorf("%w:用户已被删除", ErrUserNotLogin)
		}
		return ErrUserNotLogin
	}
	return nil
}
func ErrorsNew(str string) error {
	return errors.New(str)
}
func ErrorsFmt(str string) error {
	return fmt.Errorf(str, 1)
}
