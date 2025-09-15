package backend

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestAddConn(t *testing.T) {
	InitDB()
	time.Sleep(2 * time.Second)
	var db = &DbService{ctx: context.Background(), db: dbConn}
	conn := Conn{
		Name:  "testName",
		URL:   "testUrl",
		Token: "testToken",
	}
	err := db.AddConn(&conn)
	if err != nil {
		t.Error(err)
	}
	conns, err := db.QueryConns()
	if err != nil {
		t.Error(err)
	}
	log.Println(conns)
}
