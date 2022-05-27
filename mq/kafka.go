package mq

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/prometheus/common/log"
	gjson "github.com/shootz-developer/gtool/json"
)

// InitKafkaProducer Kafka生产者初始化
func InitKafkaProducer(kafkaBrokers []string) (kafkaProducer sarama.SyncProducer) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	var err error
	kafkaProducer, err = sarama.NewSyncProducer(kafkaBrokers, config)
	if err != nil {
		log.Errorf("Kafka生产者初始化失败:%+v \n", err)
	}
	log.Infof("Kafka生产者初始化成功")
	return
}

// ProduceKafkaMsg 生产Kafka消息
func ProduceKafkaMsg(kafkaProducer sarama.SyncProducer, kafkaTopic string, msg string) (retMsg string) {
	msgX := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Value: sarama.StringEncoder(msg),
	}
	partition, offset, err := kafkaProducer.SendMessage(msgX)
	if err != nil {
		retMsg = fmt.Sprintf("消息发送(%+v)出错:%+v \n", gjson.DumpString(msgX), err)
	} else {
		retMsg = fmt.Sprintf("消息发送(%+v)成功并存储在topic(%s)/partition(%d)/offset(%d)\n",
			gjson.DumpString(msgX), kafkaTopic, partition, offset)
	}
	return
}

// InitKafkaConsumer Kafka消费者初始化
func InitKafkaConsumer(
	kafkaBrokers []string,
	kafkaTopic string,
	groupID string) (
	kafkaConsumer *cluster.Consumer) {
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	kafkaConsumer, err = cluster.NewConsumer(kafkaBrokers, groupID, []string{kafkaTopic}, config)
	if err != nil {
		log.Errorf(err.Error())
	}
	if kafkaConsumer == nil {
		log.Errorf(fmt.Sprintf("Kafka消费者为空 {brokers:%v, topic:%v, group:%v}", kafkaBrokers, kafkaTopic, groupID))
	} else {
		log.Infof("Kafka消费者初始化成功 {consumer:%v, topic:%v}", kafkaConsumer, kafkaTopic)
	}
	return
}

// KeepKafkaConsuming 保持消费
func KeepKafkaConsuming(
	kafkaConsumer *cluster.Consumer,
	handleKafkaConsumerMsg func(*cluster.Consumer, *sarama.ConsumerMessage)) {
	for {
		log.Infof("循环读取kafka")
		select {
		case msg, ok := <-kafkaConsumer.Messages():
			if ok {
				handleKafkaConsumerMsg(kafkaConsumer, msg)
			} else {
				log.Errorf("Kafka监听服务失败")
			}
		case err, ok := <-kafkaConsumer.Errors():
			if ok {
				log.Errorf("Kafka消费者报错: %+v", err)
			}
		case ntf, ok := <-kafkaConsumer.Notifications():
			if ok {
				log.Infof("Kafka消费者提醒: %+v", ntf)
			}
		}
	}
}
