package user

import "time"

func SendPhoneMsg(num string) string {
	time.Sleep(1 * time.Second)
	return num + " send msg success!"
}
