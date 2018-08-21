package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json: "id" bson: "_id, omitempty"`
	UserName string        `json: "username"`
	Password string        `json: "password, omitempty"`
	Token    string        `json "token"`
}

func FindUser(db *mgo.Database, query *User) (*User, error) {

	user := User{}
	err := db.C("users").
		Find(bson.M{"username": query.UserName, "password": query.Password}).
		One(&user)

	return &user, err
}

func (user *User) Create(db *mgo.Database) error {
	fmt.Println(user)
	user.Id = bson.NewObjectId()
	err := db.C("users").Insert(&user)

	if err != nil {
		return err
	}
	return nil
}
