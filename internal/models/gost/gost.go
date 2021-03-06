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

// Gost is struct for gost
type Gost struct {
	// ID is gost uuid
	ID string `json:"id"`
	// Public is if the gost is public
	Public bool `json:"public"`
	// Description is gost description message
	Description string `json:"description"`
	// Version is gost version
	Version int `json:"version"`
	// Files is files contained by gost
	Files []*File `bson:"filesArray" json:"files"`
	// CreatedAt is gost created time
	CreatedAt string `json:"created_at"`
	// User is gost owner user
	User user.User `json:"user"`
	// Status is gost status
	Status int `json:"-"`
}

// File is gost file struct
type File struct {
	// ID is uuid of file
	ID string `json:"id"`
	// Filename is the file name
	Filename string `json:"filename"`
	// Content is file content
	Content string `json:"content"`
}

var (
	table  = "gosts"
	client *mgo.Collection
	log    = logger.Logger
	// MaxFilesCount is max files count allowed
	MaxFilesCount = 10
)

const (
	// STATUSWELL mean gost can be show
	STATUSWELL = 1 + iota
	// STATUSDELETEDBYUSER mean gost is already deleted by user
	STATUSDELETEDBYUSER
	// STATUSDELETEDBYSYSTEM mean gost is already deleted by system
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

// NewGost is gost contructor
func NewGost(description string, files []*File, user user.User, version int, public bool) *Gost {
	gost := &Gost{
		ID:          utils.Uuid(),
		Public:      public,
		Description: description,
		Version:     version,
		Files:       files,
		CreatedAt:   time.Now().Format(time.RFC3339),
		User:        user,
		Status:      STATUSWELL,
	}
	gost.injectFileUuid()
	return gost
}

// NewDefaultGost create gost with some default fields
func NewDefaultGost(description string, files []*File, user user.User) *Gost {
	gost := &Gost{
		ID:          utils.Uuid(),
		Public:      true,
		Description: description,
		Version:     1,
		Files:       files,
		CreatedAt:   time.Now().Format(time.RFC3339),
		User:        user,
		Status:      STATUSWELL,
	}
	gost.injectFileUuid()
	return gost
}

// WithUser add author for gost and some init fields
func (g *Gost) WithUser(user user.User) {
	g.User = user
	g.Public = true
	g.CreatedAt = time.Now().Format(time.RFC3339)
	g.Version = 1
	g.ID = utils.Uuid()
	g.Status = STATUSWELL
	g.injectFileUuid()
	log.Debugf("=========%#v", g.Files)
}

// Create method store gost to db
func (g *Gost) Create() error {
	log.Debugf("create gost %#v", g)
	return client.Insert(g)
}

// Remove method soft delete a gost
func (g *Gost) Remove(isUser bool) error {
	status := STATUSDELETEDBYUSER
	if !isUser {
		status = STATUSDELETEDBYSYSTEM
	}
	return client.Update(bson.M{"id": g.ID}, bson.M{"$set": bson.M{"status": status}})
}

// GetGostById find a gost from db by id
func (g *Gost) GetGostById(id string) error {
	return client.Find(bson.M{"id": id, "status": bson.M{"$lte": STATUSWELL}}).One(g)
}

// GetGostsByUsername find gosts from db by author name
func (g *Gost) GetGostsByUsername(username string) ([]Gost, error) {
	var gosts []Gost
	err := client.Find(bson.M{"user.username": username, "status": bson.M{"$lte": STATUSWELL}}).All(&gosts)
	return gosts, err
}

// IgnoreEmpty will ignore empty file
func (g *Gost) IgnoreEmpty() {
	var files []*File
	for _, v := range g.Files {
		if v.Content != "" {
			files = append(files, v)
		}
	}
	g.Files = files
}

// Validate make sure if the gost is validate
func (g *Gost) Validate() bool {
	if g.Description == "" {
		return false
	}
	if len(g.Files) == 0 || len(g.Files) > MaxFilesCount {
		return false
	}
	for _, v := range g.Files {
		if v.Content == "" || v.Filename == "" {
			return false
		}
	}
	return true
}

// FindFile find file from gost by file uuid
func (g *Gost) FindFile(fileId string) *File {
	for _, v := range g.Files {
		log.Debug(v.ID)
		if v.ID == fileId {
			return v
		}
	}
	return nil
}

func (g *Gost) injectFileUuid() {
	for _, v := range g.Files {
		v.ID = utils.Uuid()
	}
}
