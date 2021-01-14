package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

var (
	dyn *dynamodb.DynamoDB
)

const (
	REGION = "eu-central-1"
	TABLE  = "myMembers"
)

type Member struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"location"`
	FriendId []int  `json:"friendId"`
	Likes    int    `json:"likes"`
}

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	}))

	dyn = dynamodb.New(sess)

	err := ensureTableExists()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	members := getData()

	for _, member := range members {
		av, err := dynamodbattribute.MarshalMap(member)
		if err != nil {
			log.Println(err)
		}

		_, err = dyn.PutItem(&dynamodb.PutItemInput{
			TableName: aws.String(TABLE),
			Item:      av,
		})
		if err != nil {
			log.Println(err)
		}
	}

	output, err := dyn.Scan(&dynamodb.ScanInput{
		TableName: aws.String(TABLE),
	})
	if err != nil {
		log.Println(err)
	}

	var membersOut []Member
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &membersOut)
	if err != nil {
		log.Println(err)
	}

	for _, member := range members {
		log.Printf("%+v\n", member)
	}

}
