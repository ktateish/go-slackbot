package brain

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	iface "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamoDB struct {
	iface.DynamoDBAPI

	table map[string]map[string][]byte
}

func newMockDynamoDB() *mockDynamoDB {
	res := &mockDynamoDB{
		table: make(map[string]map[string][]byte),
	}
	return res
}

func (mock *mockDynamoDB) DescribeTableWithContext(ctx context.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	if input == nil {
		input = &dynamodb.DescribeTableInput{}
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	_, ok := mock.table[*input.TableName]
	if !ok {
		return nil, awserr.New(dynamodb.ErrCodeTableNotFoundException, fmt.Sprintf("%s", *input.TableName), nil)
	}

	output := &dynamodb.DescribeTableOutput{}
	return output, nil
}

func (mock *mockDynamoDB) CreateTableWithContext(ctx context.Context, input *dynamodb.CreateTableInput, opts ...request.Option) (*dynamodb.CreateTableOutput, error) {
	if input == nil {
		input = &dynamodb.CreateTableInput{}
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	_, ok := mock.table[*input.TableName]
	if ok {
		return nil, awserr.New(dynamodb.ErrCodeTableAlreadyExistsException, fmt.Sprintf("%s", *input.TableName), nil)
	}
	mock.table[*input.TableName] = make(map[string][]byte)

	output := &dynamodb.CreateTableOutput{}
	return output, nil
}

func (mock *mockDynamoDB) GetItemWithContext(ctx context.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	if input == nil {
		input = &dynamodb.GetItemInput{}
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	db, ok := mock.table[*input.TableName]
	if !ok {
		return nil, awserr.New(dynamodb.ErrCodeTableNotFoundException, fmt.Sprintf("%s", *input.TableName), nil)
	}

	item := dyndbItem{}
	err := dynamodbattribute.UnmarshalMap(input.Key, &item)
	if err != nil {
		keys := make([]string, 0)
		for k := range input.Key {
			keys = append(keys, k)
		}
		return nil, awserr.New(dynamodb.ErrCodeIndexNotFoundException, strings.Join(keys, ","), nil)
	}

	val, ok := db[item.Key]
	if !ok {
		return nil, awserr.New(dynamodb.ErrCodeIndexNotFoundException, fmt.Sprintf("%s", item.Key), nil)
	}
	item.Value = val

	av, err := dynamodbattribute.MarshalMap(item)

	output := &dynamodb.GetItemOutput{}
	output.Item = av

	return output, nil
}

func (mock *mockDynamoDB) PutItemWithContext(ctx context.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	if input == nil {
		input = &dynamodb.PutItemInput{}
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	db, ok := mock.table[*input.TableName]
	if !ok {
		return nil, awserr.New(dynamodb.ErrCodeTableNotFoundException, fmt.Sprintf("%s", *input.TableName), nil)
	}

	item := dyndbItem{}
	err := dynamodbattribute.UnmarshalMap(input.Item, &item)
	if err != nil {
		return nil, awserr.New(dynamodb.ErrCodeIndexNotFoundException, "Key", nil)
	}

	db[item.Key] = item.Value

	output := &dynamodb.PutItemOutput{}
	return output, nil
}

func TestDynamoDBBrainManipurateTable(t *testing.T) {
	mock := newMockDynamoDB()

	br, err := NewDynamoDBBrain(mock, "test")
	if err != nil {
		t.Fatalf("NewDynamoDBBrain failed")
	}

	got, err := br.TableExists(context.Background())
	if err != nil {
		t.Fatalf("TableExists returned unexpected error: %s", err)
	}
	if want := false; want != got {
		t.Errorf("TableExists for test table failed: want=%t got=%t\n", want, got)
	}

	err = br.CreateTable(context.Background())
	if err != nil {
		t.Fatalf("CreateTable returned unexpected error: %s", err)
	}

	got, err = br.TableExists(context.Background())
	if err != nil {
		t.Fatalf("TableExists returned unexpected error: %s", err)
	}
	if want := true; want != got {
		t.Errorf("TableExists for test table failed: want=%t got=%t\n", want, got)
	}

}

func TestDynamoDBBrainManipurateItem(t *testing.T) {
	mock := newMockDynamoDB()

	br, err := NewDynamoDBBrain(mock, "test")
	if err != nil {
		t.Fatalf("NewDynamoDBBrain failed")
	}

	err = br.CreateTable(context.Background())
	if err != nil {
		t.Fatalf("CreateTable returned unexpected error: %s", err)
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("/test/key-%d", i)
		val := []byte(fmt.Sprintf("test-value-%d", i))
		err := br.Save(context.Background(), key, val)
		if err != nil {
			t.Fatalf("Save returned unexpected error: %s", err)
		}
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("/test/key-%d", i)
		want := []byte(fmt.Sprintf("test-value-%d", i))
		got, err := br.Load(context.Background(), key)
		if err != nil {
			t.Fatalf("Load returned unexpected error: %s", err)
		}
		if !bytes.Equal(want, got) {
			t.Errorf("Load returned wrong value for key '%s': want=%s, got=%s", key, string(want), string(got))
		}
	}

}
