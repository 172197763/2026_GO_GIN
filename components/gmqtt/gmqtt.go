package gmqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"gin_test/config"
	"net"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client mqtt.Client
}

var (
	instance *MQTTClient
	mutex    sync.Mutex
)

func GetMQTTClient() (*MQTTClient, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if instance != nil && instance.client.IsConnected() {
		// fmt.Println("连接健康状态", instance.client.IsConnected())
		return instance, nil
	}
	broker := config.Get("mqtt.host", "") + ":" + config.Get("mqtt.port", "")
	clientID := fmt.Sprintf("admin_%d", time.Now().UnixNano())
	username := config.Get("mqtt.user", "")
	password := config.Get("mqtt.password", "")
	useTLS := false
	var err error
	//实例化
	instance, err = NewMQTTClient(broker, clientID, username, password, useTLS)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
func NewMQTTClient(broker, clientID, username, password string, useTLS bool) (*MQTTClient, error) {
	opts := mqtt.NewClientOptions()

	if useTLS {
		opts.AddBroker("ssl://" + broker)
		// 加载 CA 证书（生产环境）
		caCert, _ := os.ReadFile("ca.crt")
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(caCert)
		opts.TLSConfig = &tls.Config{
			RootCAs:    certPool,
			ServerName: extractHost(broker),
		}
	} else {
		opts.AddBroker("tcp://" + broker)
	}

	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(1 * time.Second)
	opts.SetConnectTimeout(2 * time.Second)

	timeout := 2 * time.Second
	var client mqtt.Client
	var lastErr error
	for attempt := 1; attempt <= 3; attempt++ {
		if attempt > 1 {
			fmt.Printf("尝试连接 MQTT Broker[%s]... (第 %d 次)\n", broker, attempt)
		}
		client = mqtt.NewClient(opts)
		token := client.Connect()

		// 等待连接完成（最多 timeout 时间，由 SetConnectTimeout 保证）
		if token.WaitTimeout(timeout) { // WaitTimeout 返回 true 表示未超时
			if token.Error() == nil {
				fmt.Println("✅ MQTT 连接成功")
				return &MQTTClient{client: client}, nil
			}
			lastErr = token.Error()
			fmt.Printf("❌ 连接失败: %v\n", lastErr)
		} else {
			lastErr = fmt.Errorf("连接超时（> %v）", timeout)
			fmt.Printf("⏰ %v\n", lastErr)
		}
	}

	return nil, fmt.Errorf("连接mq服务失败")
}

// 订阅主题
func (m *MQTTClient) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) error {
	token := m.client.Subscribe(topic, qos, handler)
	token.Wait()
	return token.Error()
}

// @description: 发布消息
//
// @param {string} topic 主题
//
// @param {byte} qos QoS等级 0至多一次 1至少一次 2 恰好一次
func (m *MQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	if m.client == nil {
		return fmt.Errorf("mqtt client is nil")
	}
	if !m.client.IsConnected() {
		return fmt.Errorf("mqtt client is not connected")
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	newpayload := string(data)
	if newpayload == "" {
		return fmt.Errorf("payload is empty")
	}
	token := m.client.Publish(topic, qos, retained, newpayload)
	if token.WaitTimeout(5 * time.Second) { // 使用超时机制避免无限等待
		return token.Error()
	} else {
		return fmt.Errorf("publish operation timed out")
	}
}

func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
}

// extractHost 从 "host:port" 中提取 host 部分
func extractHost(address string) string {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		// 如果格式不是 host:port（比如只有 host），直接返回原值
		return address
	}
	return host
}
