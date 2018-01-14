package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
)

var Mongo *mgo.Session

func init() {
	session, err := mgo.Dial(utils.GetEnvOrDefault("MONGOURL", "localhost"))
	if err != nil {
		logger.Logger.Fatalf("open mongo error: %v", err)
	}
	session.SetMode(mgo.Monotonic, true)
	Mongo = session
}
