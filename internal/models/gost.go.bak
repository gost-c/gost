package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gost-c/gost/internal/utils"
	"github.com/pkg/errors"
	"time"
)

type Gost struct {
	ID          string
	Public      bool
	Description string
	Version     int
	Files       []File
	CreatedAt   string
	User
}

type File struct {
	Filename string
	Content  string
}

var (
	tableGost = "gosts"
)

func NewGost(description string, files []File, user User, version int, public bool) *Gost {
	return &Gost{
		ID:          utils.Uuid(),
		Public:      public,
		Description: description,
		Version:     version,
		Files:       files,
		CreatedAt:   time.Now().String(),
		User:        user,
	}
}

func (g *Gost) Create() error {
	item := g.createGost()

	log.Debugf("try create gost %#v", g)

	_, err := client.PutItem(&dynamodb.PutItemInput{
		TableName: &tableGost,
		Item:      item,
	})

	return err
}

func (g *Gost) Remove() error {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: &g.ID,
		},
	}

	_, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: &tableGost,
		Key:       key,
	})

	return err
}

func (g *Gost) GetGostById(id string) error {
	key := map[string]*dynamodb.AttributeValue{
		"username": {
			S: &g.Username,
		},
	}

	res, err := client.GetItem(&dynamodb.GetItemInput{
		TableName:      &table,
		Key:            key,
		ConsistentRead: aws.Bool(true),
	})

	if err != nil {
		return errors.Wrap(err, "getting item")
	}

	if err := dynamodbattribute.UnmarshalMap(res.Item, &g); err != nil {
		return errors.Wrap(err, "unmarshaling item")
	}

	return nil
}

func (g *Gost) createGost() map[string]*dynamodb.AttributeValue {
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: &g.ID,
		},
		"public": {
			BOOL: &g.Public,
		},
		"description": {
			S: &g.Description,
		},
		"version": {
			N: aws.String(fmt.Sprintf("%d", g.Version)),
		},
		"user": {
			M: g.User.createUserItem(),
		},
		"files": {
			L: g.createFileAttributes(),
		},
		"created_at": {
			S: &g.CreatedAt,
		},
	}
	return item
}

func (g *Gost) createFileAttributes() []*dynamodb.AttributeValue {
	var res []*dynamodb.AttributeValue
	for _, v := range g.Files {
		res = append(res, &dynamodb.AttributeValue{
			M: map[string]*dynamodb.AttributeValue{
				"filename": {
					S: &v.Filename,
				},
				"content": {
					S: &v.Content,
				},
			},
		})
	}
	return res
}
