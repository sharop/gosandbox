// Example function-based Apache Kafka producer
package main

/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type transaction struct {
	Name   string  `json:"name"` //Key
	Amount float64 `json:"amount"`
	Time   string  `json:"time"` // RFC3339
}

func main() {

	/*if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <topic>\n",
			os.Args[0])
		os.Exit(1)
	}*/

	broker := "localhost"  //os.Args[1]
	topic := "transaction" //os.Args[2]

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)

	for index := 0; index < 10000; index++ {
		keyTrans, valueTrans := getKV()

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(valueTrans),
			Key:            []byte(keyTrans),
			Headers:        []kafka.Header{{Key: keyTrans, Value: []byte(valueTrans)}},
		}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

	}

	close(deliveryChan)
}

func getKV() (string, string) {
	rand.Seed(time.Now().UnixNano())

	accountname := []string{"Sergio", "Marina", "Merari", "Endora", "Morgana"}
	ammountList := []float64{0.5, 0.1, 1, 2, 5, 10, 20, 50, 100, 200, 500, 1000}
	dep := []int{-1, 1}
	indexName := rand.Intn(len(accountname))
	indexAmmount := rand.Intn(len(ammountList))

	nameKey := accountname[indexName]
	ammount := ammountList[indexAmmount] * float64(dep[rand.Intn(len(dep))])
	time := time.Now().Format(time.RFC3339)

	trans := transaction{
		nameKey,
		ammount,
		time,
	}
	jsonTrans, err := json.Marshal(trans)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Tipo marshal %T\n", jsonTrans)
		//fmt.Printf("Sin + %v\n", jbd1)
		//fmt.Printf("con + %+v\n", jbd1)
		fmt.Println(string(jsonTrans))
	}
	return nameKey, string(jsonTrans)

}
