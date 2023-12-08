package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"zrDispatch/core/slog"

	"github.com/segmentio/kafka-go"
)

func main() {
	// createTopicByConn()
	// writeByConn()
	readByConn()
	// topiclist()
}

func createTopicByConn() {
	// 指定要创建的topic名称
	topic := "my-topic"

	// 连接至任意kafka节点
	conn, err := kafka.Dial("tcp", "172.16.160.96:9092")
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	defer conn.Close()

	// 获取当前控制节点信息
	controller, err := conn.Controller()
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	var controllerConn *kafka.Conn
	// 连接至leader节点
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	// 创建topic
	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		slog.Println(slog.DEBUG, err)
	}
}

func writeByConn() {
	topic := "my-topic"
	partition := 0

	// 连接至Kafka集群的Leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.16.160.96:9092", topic, partition)
	if err != nil {
		slog.Println(slog.DEBUG, "failed to dial leader:", err)
	}

	// 设置发送消息的超时时间
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	// 发送消息
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		slog.Println(slog.DEBUG, "failed to write messages:", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		slog.Println(slog.DEBUG, "failed to close writer:", err)
	}
}

// readByConn 连接至kafka后接收消息
func readByConn() {
	// 指定要连接的topic和partition
	topic := "my-topic"
	partition := 0

	// 连接至Kafka的leader节点
	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.16.160.96:9092", topic, partition)
	if err != nil {
		slog.Println(slog.DEBUG, "failed to dial leader:", err)
	}

	// 设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	// 读取一批消息，得到的batch是一系列消息的迭代器
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	// 遍历读取消息
	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		slog.Println(slog.DEBUG, "读消息", string(b[:n]))
	}

	// 关闭batch
	if err := batch.Close(); err != nil {
		slog.Println(slog.DEBUG, "failed to close batch:", err)
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		slog.Println(slog.DEBUG, "failed to close connection:", err)
	}
}

func topiclist() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}
	// 遍历所有分区取topic
	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}
