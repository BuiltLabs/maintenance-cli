package main

import (
	"flag"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"time"
	"fmt"
)
type Maintenance struct {
	ready bool

	enabled bool
	timestamp time.Time
}

var maintenance Maintenance

var fileTarget string
var storageEngine string

var tableName string
var primaryPartitionKeyName string
var primaryPartitionKey string

func init() {
	flag.StringVar(&fileTarget, "target", "./maintenance.enable", "the file which is to be added/removed")
	flag.StringVar(&storageEngine, "engine", "dynamodb", "The data store model used. DynamoDB is the only option right now.")
	flag.StringVar(&tableName, "tableName", "maintenanceFlags", "The table which our key/data is stored in")
	flag.StringVar(&primaryPartitionKey, "key", "ops", "Primary Partition Key Value")
	flag.StringVar(&primaryPartitionKeyName, "keyName", "environment", "Primary Partion Key Name")
}

func main() {
	var output string

	flag.Parse()

	maintenance.timestamp = time.Now()

	if err := checkFlags(); err != nil {
		panic(err)
	}

	if maintenance.ready {
		if maintenance.enabled {
			if err := enable(); err != nil {
				panic(err)
			}

			output = "maintenance enabled"
		} else {
			if err := disable(); err != nil {
				panic(err)
			}

			output = "maintenance disabled"
		}
	} else {
		output = "unknown error"
	}

	fmt.Printf("[%s] - %s", maintenance.timestamp.Format(time.RFC3339), output)
}

func checkFlags() (err error) {
	switch storageEngine {
	case "dynamodb":
		item, err := checkDynamoDB()

		if (err == nil && item.Item["enabled"] != nil) {
			maintenance.ready   = true
			maintenance.enabled = *item.Item["enabled"].BOOL
		}

		break;
	case "local":
		break;
	}

	return;
}

func checkDynamoDB() (*dynamodb.GetItemOutput, error) {
	svc := dynamodb.New(session.New(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			primaryPartitionKeyName: {
				S: aws.String(primaryPartitionKey),
			},
		},
		TableName: aws.String(tableName),
	}

	resp, err := svc.GetItem(params);

	return resp, err
}

func enable() (err error) {
	file, err := os.Create(fileTarget);

	if err != nil {
		log.Printf("File creation failed (reason: %s)", err)
		return
	}

	if err := file.Close(); err != nil {
		log.Printf("File close failed (reason: %s)", err)
	}

	return
}

func disable() (err error) {
	if _, err := os.Stat(fileTarget); err != nil {
		return nil
	}

	err = os.Remove(fileTarget)

	return
}