package models

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/pkg/errors"
	"regexp"
	"time"
)

// User is gost user
type User struct {
	// ID is uuid
	ID string
	// Username username
	Username string
	// Password password
	Password string
	// Tokens tokens
	Tokens []string
	// Joined joined
	Joined string
}

var (
	table                = "users"
	client               *dynamodb.DynamoDB
	log                  = logger.Logger
	validUsername        = regexp.MustCompile(`^[a-zA-Z0-9_]{6,20}$`)
	validaPassword       = regexp.MustCompile(`^[a-zA-Z0-9!"#$%&'()*+,-./:;<=>?@\[\\\]^_{|} ~]{6,20}$`)
	ErrUsernameInvalid   = errors.New("Username length should > 6 and < 20, only support character, numbers and '_'")
	ErrPasswordInvalid   = errors.New("Password's length should > 6 and < 20")
	ErrUserAlreadyExists = errors.New("User already exists, please try another username.")
)

type UserModel interface {
	Create() error
	Remove() error
	AddToken(token string) error
	GetUserByName() error
	Validate() (bool, error)
}

func init() {
	ss, err := session.NewSession(aws.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	client = dynamodb.New(ss)
}

func NewUser(username, password string) *User {
	hashedPass, _ := utils.HashPassword(password)
	return &User{
		ID:       utils.Uuid(),
		Username: username,
		Password: hashedPass,
		Joined:   time.Now().String(),
	}
}

func (u *User) New() *User {
	return NewUser(u.Username, u.Password)
}

func (u *User) Create() error {
	err := u.GetUserByName()
	if err != nil {
		log.Debugf("find user error: %v", err)
		return err
	}
	if u != nil {
		log.Debugf("username exists, %s", u.Username)
		return ErrUserAlreadyExists
	}
	item := u.createUserItem()
	log.Debugf("try create user %#v", u)
	_, err = client.PutItem(&dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	})

	return err
}

func (u *User) Remove() error {
	key := map[string]*dynamodb.AttributeValue{
		"username": {
			S: &u.Username,
		},
	}

	_, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: &table,
		Key:       key,
	})

	return err
}

func (u *User) AddToken(token string) error {
	return nil
}

func (u *User) GetUserByName() error {
	key := map[string]*dynamodb.AttributeValue{
		"username": {
			S: &u.Username,
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

	if err := dynamodbattribute.UnmarshalMap(res.Item, &u); err != nil {
		return errors.Wrap(err, "unmarshaling item")
	}

	return nil
}

func (u *User) Validate() (bool, error) {
	if !validUsername.MatchString(u.Username) {
		return false, ErrUsernameInvalid
	}
	if !validaPassword.MatchString(u.Password) {
		return false, ErrPasswordInvalid
	}
	return true, nil
}

func (u *User) createUserItem() map[string]*dynamodb.AttributeValue {
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: &u.ID,
		},
		"username": {
			S: &u.Username,
		},
		"password": {
			S: &u.Password,
		},
		"joined": {
			S: &u.Joined,
		},
	}
	return item
}
