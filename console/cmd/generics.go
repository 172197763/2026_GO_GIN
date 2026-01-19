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
type Person struct {
	Name    string
	Age     int
	UseTool bool
}
type Animal struct {
	Name string
	Age  int
}
type Resp[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

/**
 * @description:comparable 是 Go 1.18 引入的一个预定义标识符，它表示可以使用 == 和 != 运算符进行比较的类型，包括所有基本类型（如 int, float64, string 等）和某些复合类型（如数组、结构体等，但不包括切片、映射和函数）
 */
func handle() {
	var persons = []Person{
		{"张三", 18, true},
		{"李四", 19, true},
		{"王五", 20, true},
	}
	persons = Filter(persons, func(v any) bool {
		person := v.(Person)
		return person.Age >= 19
	})
	fmt.Println(persons)
	response()
}
func response() {
	p := Person{Name: "张三", Age: 18, UseTool: true}
	a := Animal{Name: "小黄", Age: 1}
	r1 := Resp[Person]{Data: p}
	a1 := Resp[Animal]{Data: a}
	//可直接调用data中的属性
	fmt.Println(r1.Data.UseTool)
	fmt.Println(a1.Data.Name)
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
