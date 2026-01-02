package cook

import "fmt"

// 如何制作一道美味的鱼料理
func CookFish() {
	AddSalt()
	Deboning()
	fmt.Println("鱼已经煮完了")
}
func Deboning() {
	fmt.Println("拔掉鱼刺")
}
