package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

// Define the kafka trigger event struct to be received
type StructEvent struct {
	Data            KafkaData
	Id              string
	Source          string
	SpecVersion     string
	Type            string
	DataContentType string
	Time            string
	Subject         string
	AliyunAccountId string
}
type KafkaData struct {
	Topic     string
	Partition int
	Offset    int
	Timestamp int
	Headers   DataHeader
	Value     string
}

type DataHeader struct {
	Headers    []string
	IsReadOnly bool
}

func HandleRequest(ctx context.Context, events []StructEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	flog := fctx.GetLogger()

	for _, evt := range events {
		// The trigger event data is in the `Data` json object from the json array
		flog.Info("kafka topic:", evt.Data.Topic)
		flog.Info("kafka messgae:", evt.Data.Value)
	}

	return fmt.Sprintf("Receive Kafka Trigger Event: %v", events), nil
}

func main() {
	fc.Start(HandleRequest)
}
