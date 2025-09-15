package backend

type Msg struct {
	Id      int
	TopicId int
	Content string
	Type    MsgType
	Time    int64
}

type MsgType string

const (
	From MsgType = "from"
	To   MsgType = "to"
)

func (db *DbService) QueryByTopic(topicId int) (res []Msg, err error) {
	stmt, err := db.db.Prepare(`SELECT * FROM Msg WHERE topic_id=?`)
	if err != nil {
		LogError(err.Error())
		return res, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var msg Msg
		if err := rows.Scan(&msg.Id, &msg.TopicId, &msg.Content, &msg.Type, &msg.Time); err != nil {
			LogError(err.Error())
			continue
		}
		res = append(res, msg)
	}
	return res, nil
}

func (msg *Msg) add() error {
	stmt, err := dbConn.Prepare(`INSERT INTO Msg (topic_id, content, type,Time) VALUES (?,?,?,?)`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(msg.TopicId, msg.Content, msg.Type, msg.Time)
	if err != nil {
		LogError(err.Error())
		return err
	}
	return nil
}

func delByTopic(topicId int) error {
	stmt, err := dbConn.Prepare(`DELETE FROM Msg WHERE topic_id=?`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(topicId)
	if err != nil {
		LogError(err.Error())
		return err
	}
	return nil
}
