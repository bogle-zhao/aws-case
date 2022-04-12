package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func LoadAwsConfig() aws.Config {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println(err)
	}
	return cfg
}
