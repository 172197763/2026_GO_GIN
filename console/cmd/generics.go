package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 泛型演示
var GenericsCmd = &cobra.Command{
	Use:   "GenericsCmd",
	Short: "泛型演示",
	Long:  `泛型演示`,
	Run: func(cmd *cobra.Command, args []string) {
		handle()
	},
}

// 这里定义了一个 Number 接口，表示类型参数必须是 int 或 float64 或它们的别名。
type Number interface {
	~int | ~float64
}

/**
 * @description:comparable 是 Go 1.18 引入的一个预定义标识符，它表示可以使用 == 和 != 运算符进行比较的类型，包括所有基本类型（如 int, float64, string 等）和某些复合类型（如数组、结构体等，但不包括切片、映射和函数）
 */
func handle() {
	type Person struct {
		Name string
		Age  int
	}
	var persons = []Person{
		{"张三", 18},
		{"李四", 19},
		{"王五", 20},
	}
	persons = Filter(persons, func(v any) bool {
		person := v.(Person)
		return person.Age >= 19
	})
	fmt.Println(persons)
}

/**
 * @description:过滤数组
 * @param arr 任意类型的切片
 * @param fun 过滤函数
 */
func Filter[T any](arr []T, fun func(v any) bool) []T {
	var result []T
	for _, v := range arr {
		if fun(v) == true {
			result = append(result, v)
		}
	}
	return result
}
