package backend

import (
	"errors"
	"strconv"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Config struct {
	Id    int
	Name  string
	Value string
}

type ConfigId int

const (
	Timeout ConfigId = iota + 1
	SubscriptionType
)

func (db *DbService) UpdateConfig(config Config) error {
	stmt, err := db.db.Prepare(`UPDATE config SET Value = ? where Id = ?`)
	if err != nil {
		LogError(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(config.Value, config.Id)
	if err != nil {
		LogError(err.Error())
	}
	return err
}

func (db *DbService) QueryConfigs() ([]Config, error) {
	stmt, err := db.db.Prepare(`SELECT * FROM config`)
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
	var res []Config
	for rows.Next() {
		var config Config
		if err := rows.Scan(&config.Id, &config.Name, &config.Value); err != nil {
			LogError(err.Error())
			continue
		}
		res = append(res, config)
	}
	return res, nil
}

func (db *DbService) queryConfig(id ConfigId) (string, error) {
	configs, err := db.QueryConfigs()
	if err != nil {
		LogError(err.Error())
		return "", err
	}
	for _, config := range configs {
		if config.Id == int(id) {
			return config.Value, nil
		}
	}
	return "", errors.New("config not found")
}

func (db *DbService) queryTimeout() time.Duration {
	const defaultTimeout = 5 * time.Second
	timeout, err := db.queryConfig(Timeout)
	if err != nil {
		return defaultTimeout
	}
	timeoutNum, err := strconv.Atoi(timeout)
	if err != nil {
		return defaultTimeout
	}
	return time.Duration(timeoutNum) * time.Second
}

func (db *DbService) querySubscription() pulsar.SubscriptionType {
	const defaultType = pulsar.Exclusive
	subType, err := db.queryConfig(SubscriptionType)
	if err != nil {
		return defaultType
	}
	subTypeNum, err := strconv.Atoi(subType)
	if err != nil {
		return defaultType
	}
	return pulsar.SubscriptionType(subTypeNum)
}
