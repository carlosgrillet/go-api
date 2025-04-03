package db

import (
	"context"
	"errors"
	"time"
  "log"

	etcd "go.etcd.io/etcd/client/v3"
)

var cli *etcd.Client

const (
	TIMEOUT  = 5 * time.Second
	ENDPOINT = "http://127.0.0.1:2379"
)

func Init() {
	endpoint := []string{ENDPOINT}

	config := etcd.Config{
		Endpoints:   endpoint,
		DialTimeout: 5 * time.Second,
	}

  var err error
  cli, err = etcd.New(config)
	if err != nil {
    log.Println("Failed to open a connection to the DB")
    panic(err)
	}

}

func Close() {
	if cli != nil {
    log.Println("Closing connection to the database")
		cli.Close()
	}
}

func Put(prefix, key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	_, err := cli.Put(ctx, prefix+key, value)
	if err != nil {
    log.Println("Writting to db key", prefix+key)
	}
}

func Get(prefix, key string, all bool) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

	var resp *etcd.GetResponse
	var err error

  log.Println("Quering db for key", prefix+key)

	if all {
		resp, err = cli.Get(ctx, prefix+key, etcd.WithPrefix())
	} else {
		resp, err = cli.Get(ctx, prefix+key)
	}

	if err != nil {
		return nil, errors.New("Failed to get key from database")
	}

	if len(resp.Kvs) == 0 {
		return nil, errors.New("Key not found")
	}

	result := make(map[string]string)
	for _, event := range resp.Kvs {
		result[string(event.Key)] = string(event.Value)
	}

	return result, nil
}

func Delete(prefix, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT)
	defer cancel()

  log.Println("Deleting from db key", prefix+key)

	_, err := cli.Delete(ctx, prefix+key)
	if err != nil {
		return errors.New("Failed to delete key")
	}
  return nil
}
