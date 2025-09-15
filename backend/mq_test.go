package backend

import (
	"context"
	"fmt"
	"testing"
)

func TestSendMq(t *testing.T) {
	InitDB()
	var db = &DbService{ctx: context.Background(), db: dbConn}
	conn := &Conn{
		Id:    1,
		Name:  "test",
		URL:   "pulsar://localhost:6650",
		Token: "",
	}
	err := db.CreateClient(conn)
	handleErr(err, t)

	topic := &Topic{
		ConnID: 1,
		Name:   "test1",
		Topic:  "first",
	}

	err = db.CreateProducer(topic)
	handleErr(err, t)

	go func() {
		for i := 0; i < 10; i++ {
			msg := &Msg{
				Id:      i,
				TopicId: 1,
				Content: fmt.Sprintf("msg:%v", i),
				Type:    To,
			}
			err = db.SendMsg(msg)
			handleErr(err, t)
		}
	}()

	err = db.CreateConsumer(topic)
	handleErr(err, t)

	err = db.Receive(topic.Id)
	handleErr(err, t)
}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
