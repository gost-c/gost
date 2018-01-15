package gost

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gost-c/gost/internal/models/user"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
	"github.com/gost-c/gost/mongo"
	"time"
)

type Gost struct {
	ID          string    `json:"id"`
	Public      bool      `json:"public"`
	Description string    `json:"description"`
	Version     int       `json:"version"`
	Files       []File    `bson:"filesArray" json:"files"`
	CreatedAt   string    `json:"created_at"`
	User        user.User `json:"user"`
	Status      int       `json:"-"`
}

type File struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

var (
	table  = "gosts"
	client *mgo.Collection
	log    = logger.Logger
)

const (
	STATUSWELL = 1 + iota
	STATUSDELETEDBYUSER
	STATUSDELETEDBYSYSTEM
)

func init() {
	client = mongo.Mongo.DB(mongo.DBName).C(table)
	err := client.EnsureIndex(mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	if err != nil {
		log.Fatalf("ensure table gosts index error: %v", err)
	}
}

func NewGost(description string, files []File, user user.User, version int, public bool) *Gost {
	return &Gost{
		ID:          utils.Uuid(),
		Public:      public,
		Description: description,
		Version:     version,
		Files:       files,
		CreatedAt:   time.Now().String(),
		User:        user,
		Status:      STATUSWELL,
	}
}

func NewDefaultGost(description string, files []File, user user.User) *Gost {
	return &Gost{
		ID:          utils.Uuid(),
		Public:      true,
		Description: description,
		Version:     1,
		Files:       files,
		CreatedAt:   time.Now().String(),
		User:        user,
		Status:      STATUSWELL,
	}
}

func (g *Gost) WithUser(user user.User) {
	g.User = user
	g.Public = true
	g.CreatedAt = time.Now().String()
	g.Version = 1
	g.ID = utils.Uuid()
	g.Status = STATUSWELL
}

func (g *Gost) Create() error {
	log.Debugf("create gost %#v", g)
	return client.Insert(g)
}

func (g *Gost) Remove(isUser bool) error {
	status := STATUSDELETEDBYUSER
	if !isUser {
		status = STATUSDELETEDBYSYSTEM
	}
	return client.Update(bson.M{"id": g.ID}, bson.M{"$set": bson.M{"status": status}})
}

func (g *Gost) GetGostById(id string) error {
	return client.Find(bson.M{"id": id, "status": bson.M{"$lte": STATUSWELL}}).One(g)
}

func (g *Gost) GetGostsByUsername(username string) ([]Gost, error) {
	var gosts []Gost
	err := client.Find(bson.M{"user.username": username}).All(&gosts)
	log.Debugf("%s ----- %#v", username, gosts)
	return gosts, err
}

func (g *Gost) Validate() bool {
	if g.Description == "" {
		return false
	}
	if len(g.Files) == 0 {
		return false
	}
	for _, v := range g.Files {
		if v.Content == "" || v.Filename == "" {
			return false
		}
	}
	return true
}
