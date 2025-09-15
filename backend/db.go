package backend

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var dbConn *sql.DB

type DbService struct {
	ctx context.Context
	db  *sql.DB
}

func (db *DbService) Startup(ctx context.Context) {
	db.ctx = ctx
	db.db = dbConn
}

func (db *DbService) Shutdown() {
	_ = db.db.Close()
}

// InitDB 初始化
func InitDB() {
	var err error
	dbConn, err = sql.Open("sqlite3", "./pulsar_desk.db")
	if err != nil {
		log.Fatal(err)
	}
	go initConn()
	go initTopic()
	go initMsg()
	go initConfig()
}

// 初始化连接
func initConn() {
	connTableSql := `CREATE TABLE IF NOT EXISTS conn (
    	Id INTEGER PRIMARY KEY AUTOINCREMENT,
    	Name TEXT,
    	url TEXT,
    	token TEXT
	);`
	execSql(connTableSql)
}

// 初始化topic
func initTopic() {
	topicTableSql := `CREATE TABLE IF NOT EXISTS topic (
    	Id INTEGER PRIMARY KEY AUTOINCREMENT,
    	conn_id INTEGER,
    	Name TEXT,
    	Topic TEXT
	);`
	execSql(topicTableSql)
}

// 初始化msg
func initMsg() {
	msgTableSql := `CREATE TABLE IF NOT EXISTS msg (
    	Id INTEGER PRIMARY KEY AUTOINCREMENT,
    	topic_id INTEGER,
    	content TEXT,
    	type TEXT,
    	Time INTEGER
	);
	CREATE INDEX IF NOT EXISTS topicIdIdx ON msg(topic_id);`
	execSql(msgTableSql)
}

// 初始化msg
func initConfig() {
	configTableSql := `CREATE TABLE IF NOT EXISTS config (
    	Id INTEGER PRIMARY KEY AUTOINCREMENT,
    	Name TEXT,
    	Value TEXT
	);
	CREATE INDEX IF NOT EXISTS ConfigNameIdx ON config(Name);
	INSERT INTO config (Id, Name, Value) VALUES (1, 'timeout', '5'),
	                                            (2, 'subscriptionType', 'Exclusive'),
	                                            (3,'theme','Auto')
	                                     ON CONFLICT(Id) DO NOTHING;
`
	execSql(configTableSql)
}

func execSql(sql string) {
	_, err := dbConn.Exec(sql)
	if err != nil {
		LogError(err.Error())
		log.Fatal(err)
	}
}
