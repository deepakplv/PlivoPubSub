package subscriber

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"subscriber/config"
)


type client struct {
	kinesis *kinesis.Kinesis;
	config   *config.Config

}

func New() (*client, error){
	return nil, nil
}

func (c *client) GetRecords(checkpoint string) (map[string][][]byte, string) {
	return nil, ""
}
