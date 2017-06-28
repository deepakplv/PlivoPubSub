package main

import "enqueuer"
import "enqueuer/config"


func main() {


	con := config.Config{
		Kinesis: config.KinesisConf{
		ShardCount: 1,
	}}

	c, _ :=enqueuer.New(con)

	//c.PutRecord()
	//c.PutRecords()

}