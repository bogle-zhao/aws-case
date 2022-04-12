package db

import (
	"aws-case/entity"
	"aws-case/log"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type dynamoDBRepo struct {
	tableName string
	client    *dynamodb.Client
}

func (dbRepo dynamoDBRepo) Save(account *entity.Account) (*entity.Account, error) {
	data, err := attributevalue.MarshalMap(account)
	if err != nil {
		return nil, err
	}
	_, err = dbRepo.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(dbRepo.tableName),
		Item:      data,
	})
	if err != nil {
		log.Error("PutItem:%v", err)
		return nil, err
	}

	return account, nil
}

func (dbRepo dynamoDBRepo) FindByUserName(userName string) (*entity.Account, error) {
	data, err := dbRepo.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(dbRepo.tableName),
		Key: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: userName},
		},
	})

	if err != nil {
		log.Error("GetItem:%v", err)
		return nil, err
	}
	if data.Item == nil {
		log.Error("username:%v not exist", userName)
		return nil, nil
	}
	item := &entity.Account{}
	err = attributevalue.UnmarshalMap(data.Item, &item)
	if err != nil {
		log.Error("UnmarshalMap:%v", err)
		return nil, err
	}

	return item, nil
}

func (dbRepo dynamoDBRepo) UpdateAvatar(username string, avatar string) error {

	_, err := dynamodb.NewFromConfig(LoadAwsConfig()).UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(dbRepo.tableName),
		Key: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: username},
		},
		UpdateExpression: aws.String("set avatar = :avatar"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":avatar": &types.AttributeValueMemberS{Value: avatar},
		},
	})
	if err != nil {
		log.Error("UpdateItem:%v", err)
		return err
	}
	return nil
}

func NewDynamoDBRepository() AccountRepository {
	return &dynamoDBRepo{
		tableName: "t_user",
		client:    dynamodb.NewFromConfig(LoadAwsConfig()),
	}
}
