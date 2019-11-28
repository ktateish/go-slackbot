package brain

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	iface "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// DynamoDBBrain represents Brain implementation on DynamoDB
type DynamoDBBrain struct {
	db                      iface.DynamoDBAPI
	tableName               string
	tableExistenceConfirmed bool
}

// Creates new DynamoDBBrain object backed by db and returns it.
// The brain uses the tableName for storing data.
func NewDynamoDBBrain(db iface.DynamoDBAPI, tableName string) (*DynamoDBBrain, error) {
	ddb := &DynamoDBBrain{
		db:        db,
		tableName: tableName,
	}

	return ddb, nil
}

// TableExists inspects whether the tabile for br exists.  It returns true
// if it exists, false if it doesn't.  The error is not nil when something
// go wrong on accessing DynamoDB API.
func (br *DynamoDBBrain) TableExists(ctx context.Context) (bool, error) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(br.tableName),
	}
	_, err := br.db.DescribeTableWithContext(ctx, input)
	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == dynamodb.ErrCodeTableNotFoundException {
			return false, nil
		}
		return false, fmt.Errorf("calling API 'DescribeTable': %w", err)
	}
	return true, nil
}

// CreateTable creates the table for br.
func (br *DynamoDBBrain) CreateTable(ctx context.Context) error {
	keys := []*dynamodb.KeySchemaElement{
		&dynamodb.KeySchemaElement{
			AttributeName: aws.String("Key"),
			KeyType:       aws.String("HASH"),
		},
	}

	attrs := []*dynamodb.AttributeDefinition{
		&dynamodb.AttributeDefinition{
			AttributeName: aws.String("Value"),
			AttributeType: aws.String("B"),
		},
	}

	sse := &dynamodb.SSESpecification{
		Enabled: aws.Bool(true),
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attrs,
		BillingMode:          aws.String("PAY_PER_REQUEST"),
		KeySchema:            keys,
		SSESpecification:     sse,
		TableName:            aws.String(br.tableName),
	}
	_, err := br.db.CreateTableWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("calling API 'CreateTable': %w", err)
	}
	return nil
}

type dyndbItem struct {
	Key   string
	Value []byte
}

// Retrieve value for the key from DynamoDB
func (br *DynamoDBBrain) Load(ctx context.Context, key string) ([]byte, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(br.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{
				N: aws.String(key),
			},
		},
	}

	res, err := br.db.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("calling API 'GetItem': %w", err)
	}

	item := dyndbItem{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling returned item: %w", err)
	}

	return item.Value, nil
}

// Put value for the key into DynamoDB
func (br *DynamoDBBrain) Save(ctx context.Context, key string, val []byte) error {
	item := dyndbItem{
		Key:   key,
		Value: val,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("marshalling new table item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(br.tableName),
	}

	_, err = br.db.PutItemWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("calling API 'PutItem': %w", err)
	}

	return nil
}
