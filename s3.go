package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3Object(bucket, key string) (io.ReadCloser, error) {

	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("error creating AWS session (check access keys): %s", err)
	}

	in := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	obj, err := s3.New(sess).GetObject(in)
	if err != nil {
		return nil, fmt.Errorf("error getting S3 object: %s", err)
	}
	return obj.Body, nil
}

func putS3Object(bucket, key string, buf []byte) error {

	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("error creating AWS session (check access keys): %s", err)
	}

	in := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf),
	}

	_, err = s3.New(sess).PutObject(in)
	if err != nil {
		return fmt.Errorf("error putting S3 object: %s", err)
	}
	return nil
}
