package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/gost-c/gost/internal/utils"
	"github.com/gost-c/gost/logger"
)

// Mongo is mongo db instance
var Mongo *mgo.Session

// DBName is mongo database name we used
var DBName string

func init() {
	logger.Logger.Debugf("use mongo %s", utils.GetEnvOrDefault("MONGOURL", "localhost"))
	session, err := mgo.Dial(utils.GetEnvOrDefault("MONGOURL", "localhost"))
	if err != nil {
		logger.Logger.Fatalf("open mongo error: %v", err)
	}
	session.SetMode(mgo.Monotonic, true)
	Mongo = session
	DBName = utils.GetEnvOrDefault("DBNAME", "gost")
}
