package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/gost-c/gost/mongo"
	"time"
)

// User is gost user
type User struct {
	// ID is uuid
	ID string `json:"id"`
	// Username username
	Username string `json:"username"`
	// Password password
	Password string `json:"-"`
	// Tokens tokens
	Tokens []string `json:"-"`
	// Joined joined
	Joined string `json:"joined"`
}

var (
	table  = "users"
	client *mgo.Collection
	log    = logger.Logger
)

// UserModel is user model interface
type UserModel interface {
	Create() error
	Remove() error
	AddToken(token string) error
	GetUserByName() error
}

func init() {
	client = mongo.Mongo.DB(mongo.DBName).C(table)
	err := client.EnsureIndex(mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	if err != nil {
		log.Fatalf("ensure table users index error: %v", err)
	}
}

// NewUser create a user with certain username and password
func NewUser(username, password string) *User {
	hashedPass, _ := utils.HashPassword(password)
	return &User{
		ID:       utils.Uuid(),
		Username: username,
		Password: hashedPass,
		Joined:   time.Now().Format(time.RFC3339),
	}
}

// New fill the default fields
func (u *User) New() *User {
	return NewUser(u.Username, u.Password)
}

// Create store user to db
func (u *User) Create() error {
	log.Debugf("create user %#v", u)
	return client.Insert(u)
}

// Remove remove a user
func (u *User) Remove() error {
	return client.Remove(bson.M{"username": u.Username})
}

// AddToken add a token to user
func (u *User) AddToken(token string) error {
	return client.Update(bson.M{"username": u.Username}, bson.M{"$push": bson.M{"tokens": token}})
}

// RemoveToken remove a token from user
func (u *User) RemoveToken(token string) error {
	return client.Update(bson.M{"username": u.Username}, bson.M{"$pull": bson.M{"tokens": token}})
}

// GetUserByName can find a user by username
func (u *User) GetUserByName() error {
	log.Debugf("try find user %#v", u)
	return client.Find(bson.M{"username": u.Username}).One(u)
}
