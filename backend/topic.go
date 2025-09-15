package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Topic struct {
	Id     int
	ConnID int
	Name   string
	Topic  string
}

func (db *DbService) AddTopic(topic *Topic) error {
	stmt, err := db.db.Prepare(`INSERT INTO Topic (conn_id, Name, Topic) VALUES(?, ?, ?)`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(topic.ConnID, topic.Name, topic.Topic)
	if err != nil {
		LogError(err.Error())
		return err
	}
	return nil
}

func (db *DbService) QueryTopics(connId int) (res []Topic, err error) {
	stmt, err := db.db.Prepare(`SELECT * FROM Topic WHERE conn_id=?`)
	if err != nil {
		LogError(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(connId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var topic Topic
		if err := rows.Scan(&topic.Id, &topic.ConnID, &topic.Name, &topic.Topic); err != nil {
			LogError(err.Error())
			continue
		}
		res = append(res, topic)
	}
	return res, nil
}

func (db *DbService) DelTopic(id int) error {
	stmt, err := db.db.Prepare(`DELETE FROM Topic WHERE Id=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	//先删除其下的msg
	err = delByTopic(id)
	if err != nil {
		return err
	}
	//然后删除topic
	_, err = stmt.Exec(id)
	if err != nil {
		LogError(err.Error())
		return err
	}
	return nil
}

func (db *DbService) delByConn(connId int) error {
	stmt, err := db.db.Prepare(`SELECT ID FROM Topic WHERE conn_id=?`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(connId)
	if err != nil {
		LogError(err.Error())
		return err
	}
	var topics []Topic
	for rows.Next() {
		var topic Topic
		if err := rows.Scan(&topic.Id); err != nil {
			LogError(err.Error())
			continue
		}
		topics = append(topics, topic)
	}
	for _, topic := range topics {
		err := db.DelTopic(topic.Id)
		if err != nil {
			LogError(err.Error())
			continue
		}
	}
	return nil
}

// 生产者缓存
var producerCache = make(map[int]pulsar.Producer)

func (db *DbService) CreateProducer(topic *Topic) error {
	client, err := getClient(topic.ConnID)
	if err != nil {
		return err
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic.Topic,
	})
	if err != nil {
		LogError(err.Error())
		return err
	}
	LogInfo(fmt.Sprintf("Create new producer for topic Id %v", topic.Topic))
	producerCache[topic.Id] = producer
	return nil
}

func (db *DbService) RemoveProducer(topicId int) error {
	producer, ok := producerCache[topicId]
	if !ok {
		return errors.New("not found producer")
	}
	producer.Close()
	delete(producerCache, topicId)
	return nil
}

func (db *DbService) SendMsg(msg *Msg) error {
	producer, ok := producerCache[msg.TopicId]
	if !ok {
		return errors.New("producer not exist")
	}
	_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msg.Content),
	})
	if err != nil {
		LogError(err.Error())
		return err
	}
	LogDebug(fmt.Sprintf("Send message: %v to topic Id %v", msg.Content, msg.TopicId))
	msg.Type = To
	err = msg.add()
	return err
}

// 消费者缓存
var consumerCache = make(map[int]pulsar.Consumer)

func (db *DbService) CreateConsumer(topic *Topic) error {
	client, err := getClient(topic.ConnID)
	if err != nil {
		LogError(err.Error())
		return err
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic.Topic,
		SubscriptionName: topic.Topic + "Consumer",
		Type:             db.querySubscription(),
	})
	if err != nil {
		LogError(err.Error())
		return err
	}
	LogInfo(fmt.Sprintf("Begin to consume topic %v Msg", topic.Topic))
	consumerCache[topic.Id] = consumer
	return nil
}

func (db *DbService) RemoveConsumer(topicId int) error {
	consumer, ok := consumerCache[topicId]
	if !ok {
		return errors.New("not found consumer")
	}
	consumer.Close()
	LogInfo(fmt.Sprintf("Stop consume for topic Id %v", topicId))
	delete(consumerCache, topicId)
	return nil
}

func (db *DbService) Receive(topicId int) error {
	consumer, ok := consumerCache[topicId]
	if !ok {
		return errors.New("consumer not exist")
	}
	go func() {
		for consumerCache[topicId] != nil {
			message, err := consumer.Receive(context.Background())
			if err != nil {
				LogWarn(err.Error())
				continue
			}
			//ack消息
			err = consumer.Ack(message)
			if err != nil {
				LogWarn(err.Error())
				continue
			}
			content := string(message.Payload())
			LogDebug(fmt.Sprintf("Receive Msg: %v from topic Id %v", content, topicId))
			//入库
			msg := Msg{
				TopicId: topicId,
				Content: content,
				Type:    From,
				Time:    message.PublishTime().UnixMilli(),
			}
			//通知前端
			go runtime.EventsEmit(db.ctx, "mqMsg", msg)
			//入库
			go func() {
				err = msg.add()
				if err != nil {
					LogWarn(err.Error())
				}
			}()
		}
	}()
	return nil
}
