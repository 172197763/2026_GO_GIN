package cmd

import (
	"github.com/spf13/cobra"
)

// interface相关知识
var interfCmd = &cobra.Command{
	Use:   "interf",
	Short: "命令行工具",
	Long:  `略`,
	Run: func(cmd *cobra.Command, args []string) {
		interHandle()
		var a interface{}
		if a == nil {
			println("a is nil") //输出
		}
		var b *int = nil
		var data interface{} = b
		if data == nil {
			println("data is nil") //不输出
		}
	},
}

func interHandle() {
	duck := &Duck{Name: "小鸭子"}
	show(duck)
	duck.Iam()
	show(&Chicken{})
}
func show(b Bird) {
	b.Eat()
	b.Sound()
	if duck, ok := b.(*Duck); ok {
		duck.Iam()
	}
}

type Bird interface {
	Eat()
	Sound()
}
type Duck struct {
	Name string
}

func (d *Duck) Iam() {
	println("Duck my name is ", d.Name)
}
func (d *Duck) Eat() {
	println("Duck eat")
}
func (d *Duck) Sound() {
	println("Duck sound")
}

type Chicken struct {
}

func (d *Chicken) Eat() {
	println("Chicken eat")
}
func (d *Chicken) Sound() {
	println("Chicken sound")
}
