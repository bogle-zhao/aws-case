package db

import (
	"aws-case/log"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

type S3Repo struct {
	client *s3.Client
	bucket string
}

func (s S3Repo) PutFile(filename string, file multipart.File) string {
	filename = uuid.NewV4().String() + filename

	url, err := preUrl(filename, s)
	if err != nil {
		log.Error("上传异常:%v", err)
		return ""
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
		Body:   file,
	}

	_, err2 := s.client.PutObject(context.TODO(), input)
	if err2 != nil {
		log.Error("上传异常:%v", err)
		return ""
	}
	return url
}

func preUrl(filename string, s S3Repo) (string, error) {
	psClient := s3.NewPresignClient(s.client)

	resp, err := psClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Error("PutItem:%v", err)
		return "", err
	}
	return resp.URL, nil
}

func NewS3Repo() S3Repo {
	return S3Repo{
		client: s3.NewFromConfig(LoadAwsConfig()),
		bucket: "user-info-test",
	}
}
