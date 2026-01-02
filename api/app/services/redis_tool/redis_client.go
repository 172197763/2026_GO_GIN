package redis_tool

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisStreamListener Redis流监听器结构体
type RedisStreamListener struct {
	client *redis.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// NewRedisStreamListener 创建新的Redis流监听器
func NewRedisStreamListener() *RedisStreamListener {
	// 初始化Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:9002", // Redis地址
		Password: "",               // Redis密码
		DB:       0,                // 使用的数据库
	})

	ctx, cancel := context.WithCancel(context.Background())

	return &RedisStreamListener{
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}
}

// StartListening 启动监听Redis Stream
func (rsl *RedisStreamListener) StartListening(streamName string, groupName string, consumerName string) {
	// 创建消费者组（如果不存在）
	err := rsl.createConsumerGroup(streamName, groupName)
	if err != nil {
		log.Printf("[redis]创建消费者组失败: %v", err)
	}

	// 启动后台协程监听消息
	go rsl.listenForMessages(streamName, groupName, consumerName)
}

// createConsumerGroup 创建消费者组
func (rsl *RedisStreamListener) createConsumerGroup(streamName string, groupName string) error {
	err := rsl.client.XGroupCreateMkStream(rsl.ctx, streamName, groupName, "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return err
	}
	return nil
}

// listenForMessages 监听消息
func (rsl *RedisStreamListener) listenForMessages(streamName string, groupName string, consumerName string) {
	for {
		select {
		case <-rsl.ctx.Done():
			return
		default:
			// 读取消息
			msgs, err := rsl.client.XReadGroup(rsl.ctx, &redis.XReadGroupArgs{
				Group:    groupName,
				Consumer: consumerName,
				Streams:  []string{streamName, ">"},
				Block:    time.Second * 5,
				Count:    10,
			}).Result()

			if err != nil && err != redis.Nil {
				log.Printf("[redis]读取消息出错: %v", err)
				continue
			}

			if len(msgs) > 0 {
				for _, msg := range msgs[0].Messages {
					// 打印消息内容
					log.Printf("[redis]收到消息 - Stream: %s, ID: %s, Data: %v\n",
						msgs[0].Stream, msg.ID, msg.Values)

					// 确认消息已处理
					err := rsl.client.XAck(rsl.ctx, streamName, groupName, msg.ID).Err()
					if err != nil {
						log.Printf("[redis]确认消息失败: %v", err)
					}
				}
			}
		}
	}
}

// Stop 停止监听
func (rsl *RedisStreamListener) Stop() {
	rsl.cancel()
	rsl.client.Close()
}
