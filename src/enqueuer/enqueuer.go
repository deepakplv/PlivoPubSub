package enqueuer

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"enqueuer/config"
)


type client struct {
	kinesis *kinesis.Kinesis;
	config   *config.Config

}

func New(config *config.Config) (*client, error){
	return nil, nil
}

func (c *client) PutRecord(data []byte, key string) error {
	return nil
}

func (c *client) PutRecords(data [][]byte, key string) (map[string][][]byte, string) {
	return nil, ""
}
