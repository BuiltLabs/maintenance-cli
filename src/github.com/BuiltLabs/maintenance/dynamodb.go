package maintenance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
)

func (m *Maintenance) lookup() (*dynamodb.GetItemOutput, error) {
	svc := dynamodb.New(session.New())

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			m.KeyName: {
				S: aws.String(m.Key),
			},
		},
		TableName: aws.String(m.TableName),
	}

	resp, err := svc.GetItem(params)

	return resp, err
}

func (m *Maintenance) tableExists() (err error) {
	svc := dynamodb.New(session.New())

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String(m.TableName),
	}

	_, err = svc.DescribeTable(params)

	return err
}

func (m *Maintenance) createTable() (err error) {
	svc := dynamodb.New(session.New())

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(m.KeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(m.KeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(m.TableName),
	}

	_, err = svc.CreateTable(params)

	return err
}

func (m *Maintenance) waitTable() (err error) {
	svc := dynamodb.New(session.New())

	params := &dynamodb.DescribeTableInput {
		TableName: aws.String(m.TableName),
	}

	err = svc.WaitUntilTableExists(params)

	return err
}

func (m *Maintenance) enableRecord() (err error) {
	svc := dynamodb.New(session.New())

	items := map[string]*dynamodb.AttributeValue{}

	items[m.KeyName]    = &dynamodb.AttributeValue{ S: aws.String(m.Key) }
	items["created_at"] = &dynamodb.AttributeValue{ N: aws.String(strconv.FormatInt(m.timestamp.Unix(), 10)) }
	items["disable_at"] = &dynamodb.AttributeValue{ N: aws.String("0") } // not implemented
	items["enable_at"]  = &dynamodb.AttributeValue{ N: aws.String("0") } // not implemented
	items["enabled"]    = &dynamodb.AttributeValue{ BOOL:  aws.Bool(true) }

	if (m.MetaData != "") {
		items["meta"] = &dynamodb.AttributeValue{ S: aws.String(m.MetaData) }
	}

	params := &dynamodb.PutItemInput{
		Item: items,
		TableName: aws.String(m.TableName),
	}

	_, err = svc.PutItem(params)

	return err
}


func (m *Maintenance) disableRecord() (err error) {
	svc := dynamodb.New(session.New())

	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			m.KeyName: {
				S:  aws.String(m.Key),
			},
			"created_at": {
				N:  aws.String(strconv.FormatInt(m.timestamp.Unix(), 10)),
			},
			"disable_at": {
				N:  aws.String("0"), // not implemented
			},
			"enable_at": {
				N:  aws.String("0"), // not implemented
			},
			"enabled": {
				BOOL:  aws.Bool(false),
			},
		},
		TableName: aws.String(m.TableName),
	}

	_, err = svc.PutItem(params)

	return err
}