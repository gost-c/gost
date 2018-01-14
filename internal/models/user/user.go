package user

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/gost-c/gost/mongo"
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
	Password string `json:"-"`
	// Tokens tokens
	Tokens []string `json:"-"`
	// Joined joined
	Joined string
}

var (
	table                = "users"
	client               *mgo.Collection
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
	client = mongo.Mongo.DB(table).C(table)
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
	log.Debugf("create user %#v", u)
	return client.Insert(u)
}

func (u *User) Remove() error {
	return client.Remove(bson.M{"username": u.Username})
}

func (u *User) AddToken(token string) error {
	return client.Update(bson.M{"username": u.Username}, bson.M{"$push": bson.M{"tokens": token}})
}

func (u *User) RemoveToken(token string) error {
	return client.Update(bson.M{"username": u.Username}, bson.M{"$pull": bson.M{"tokens": token}})
}

func (u *User) GetUserByName() error {
	log.Debugf("try find user %#v", u)
	return client.Find(bson.M{"username": u.Username}).One(u)
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
