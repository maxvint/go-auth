package app

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

var DBSession *mgo.Session

var mgoTimeout = 5 * time.Second

func InitApp() {
	InitDB()
	MigrateDB()
}

func InitDB() {
	var err error
	DBSession, err = mgo.DialWithTimeout(os.Getenv("MONGO_HOST")+":"+os.Getenv("MONGO_PORT"), 10*time.Second)

	if err != nil {
		panic(err)
	}
}

func CloseApp() {
	CloseDB()
}

func GetDB(c *gin.Context) *mgo.Database {
	return c.MustGet("DB").(*mgo.Database)
}

func MigrateDB() {
	session := DBSession.Clone()
	defer session.Close()
	db := session.DB(os.Getenv("MONGO_DB"))

	db.C("todos").EnsureIndex(mgo.Index{
		Key:        []string{"title"},
		Unique:     false,
		Sparse:     false,
		Background: true,
	})
}

func CloseDB() {
	DBSession.Close()
}
