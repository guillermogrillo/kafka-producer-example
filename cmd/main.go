package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guillermogrillo/kafka-producer-example/pkg/events"
	"github.com/guillermogrillo/kafka-producer-example/pkg/producer"
)

const (
	bootstrapServers = "localhost:9092"
)

func main() {
	p, err := producer.NewTollProducer(bootstrapServers)
	if err != nil {
		panic("kafka producer could not be created")
	}
	ctx := context.Background()
	topic := "toll-records"
	tollId := uuid.New().String()
	for {
		payload, err := json.Marshal(events.TollRecord{
			Id:        uuid.New().String(),
			CreatedAt: time.Now(),
			TollId:    tollId,
		})
		if err != nil {
			fmt.Sprintf("error marshalling event for toll id %s", tollId)
		}
		p.Produce(ctx, topic, payload)
		time.Sleep(5 * time.Second)
	}
}
