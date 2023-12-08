package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	// 设置 Kafka 服务器地址
	brokers := []string{"172.16.160.96:9092"}
	// 创建 Kafka 生产者配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println("Error creating producer:", err)
		return
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Println("Error closing producer:", err)
		}
	}()

	// 发送消息
	topic := "my-topic"
	message := "Hello, Kafka!"

	// 构建消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// 发送消息并处理结果
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Failed to send message:", err)
	} else {
		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}
}
