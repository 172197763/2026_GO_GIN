package cook

import "fmt"

// 如何制作一道美味的鱼料理
func CookFish() {
	AddSalt()
	//AddSalt()
	Deboning()
	fmt.Println("小A我要负责挑一条鲜活的鱼")
	fmt.Println("小A我还要负责用最少的钱买最好的鱼")
	fmt.Println("鱼已经煮完了")
}
func AddSalt() {
	fmt.Println("加盐")
}
func Deboning() {
	fmt.Println("拔掉鱼刺")
}
