package config

type KinesisConf struct {
	ShardCount 	int;
	StreamName 	string;
	Region		string;
	SecretKey	string;
	AccessKey 	string;
}

type Config struct {
	Kinesis *KinesisConf;
}
