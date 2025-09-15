package backend

import (
	"errors"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Conn struct {
	Id    int
	Name  string
	URL   string
	Token string
}

// AddConn 添加连接
func (db *DbService) AddConn(conn *Conn) error {
	stmt, err := db.db.Prepare(`INSERT INTO conn (Name, url, token) VALUES (?, ?, ?)`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(conn.Name, conn.URL, conn.Token)
	if err != nil {
		LogError(err.Error())
	}
	return err
}

// QueryConns 查询所有
func (db *DbService) QueryConns() (conns []Conn, err error) {
	stmt, err := db.db.Prepare(`SELECT * FROM conn`)
	if err != nil {
		LogError(err.Error())
		return nil, err
	}
	defer stmt.Close()
	//查询所有数据
	rows, err := stmt.Query()
	if err != nil {
		LogError(err.Error())
		return nil, err
	}
	//添加数据
	for rows.Next() {
		var conn Conn
		if err := rows.Scan(&conn.Id, &conn.Name, &conn.URL, &conn.Token); err != nil {
			LogError(err.Error())
			continue
		}
		conns = append(conns, conn)
	}
	return conns, nil
}

// DeleteConn 删除指定id
func (db *DbService) DeleteConn(id int) error {
	//先删除其下的topic
	err := db.delByConn(id)
	if err != nil {
		LogError(err.Error())
		return err
	}
	//删除conn
	stmt, err := db.db.Prepare(`DELETE FROM conn WHERE Id=?`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		LogError(err.Error())
		return err
	}
	return nil
}

// client 缓存
var clientCache = make(map[int]pulsar.Client)

func (db *DbService) CreateClient(conn *Conn) error {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               conn.URL,
		ConnectionTimeout: db.queryTimeout(),
	})
	if err != nil {
		LogError(err.Error())
		return err
	}
	LogInfo(fmt.Sprintf(`url: %s,client created`, conn.URL))
	clientCache[conn.Id] = client
	return nil
}

func (db *DbService) Disconnect(id int) {
	clientCache[id].Close()
	LogInfo(fmt.Sprintf(`Id: %v,client created`, id))
	delete(clientCache, id)
}

func getClient(connId int) (pulsar.Client, error) {
	client, ok := clientCache[connId]
	if !ok {
		return nil, errors.New("client not connected")
	}
	return client, nil
}
