package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ensureTableExists() error {
	_, descrerr := dyn.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(TABLE),
	})

	if descrerr != nil {
		if aerr, ok := descrerr.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				// table doesn't exist
				// creating one

				_, err := dyn.CreateTable(&dynamodb.CreateTableInput{
					AttributeDefinitions: []*dynamodb.AttributeDefinition{
						{
							AttributeName: aws.String("userId"),
							AttributeType: aws.String("N"),
						},
					},
					KeySchema: []*dynamodb.KeySchemaElement{
						{
							AttributeName: aws.String("userId"),
							KeyType:       aws.String("HASH"),
						},
					},
					BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
					TableName:   aws.String(TABLE),
				})
				if err != nil {
					return err
				}

				// wait until the table is actually available
				err = dyn.WaitUntilTableExists(&dynamodb.DescribeTableInput{
					TableName: aws.String(TABLE),
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
