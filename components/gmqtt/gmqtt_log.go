package gmqtt

import (
	"fmt"
	"slices"
	"time"
)

type mqttLogClient struct {
	config _config
}
type _config struct {
	userArr  []int  //特定用户列表
	enable   bool   //是否启用日志
	deadLine string //日志过期时间
}

var MqttLogClient *mqttLogClient

func init() {
	MqttLogClient = &mqttLogClient{_config{
		userArr:  []int{},
		enable:   false,
		deadLine: "",
	}}
}
func (c *mqttLogClient) SetConfig(userArr []int, enable bool, deadLine string) {
	c.config.userArr = userArr
	c.config.enable = enable
	c.config.deadLine = deadLine
}

func (c *mqttLogClient) Log(userID int, topic string, payload string) {
	if !c.config.enable {
		return
	}
	if len(c.config.userArr) > 0 && userID != 0 && !slices.Contains(c.config.userArr, userID) {
		return
	}

	if c.config.deadLine != "" {
		layout := "2006-01-02 15:04:05" // 示例格式，可根据实际需求调整
		deadline, err := time.Parse(layout, c.config.deadLine)
		if err != nil {
			return // 解析失败则直接返回
		}
		if time.Now().After(deadline) {
			c.config.enable = false
			return
		}
	}
	mq, err := GetMQTTClient()
	if err != nil {
		return
	}
	err = mq.Publish(topic, 0, false, payload)
	if err != nil {
		fmt.Println("mqtt日志推送失败", err)
	}
}
