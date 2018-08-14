package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id      bson.ObjectId `json: "id" bson: "_id, omitempty"`
	Title   string        `json: "title" bson: "title" binding: "required, max=16"`
	DueDate int64         `json: "due_date" bson: "due_date, omitempty"`
	Color   string        `json: "color" bson: "color"`
	Marked  int8          `json: "marked" bson:"marked"`
}

type TodoGroup struct {
	Id    bson.ObjectId   `json: "id" bson: "_id, omitempty"`
	Title string          `json: "title" bson: "title" binding: "required"`
	Todos []bson.ObjectId `json: "todos" bson: "todos, omitempty"`
}

func FindTodoGroup(db *mgo.Database) (*TodoGroup, error) {
	group := TodoGroup{}
	err := db.C("todo_groups").
		Find(bson.M{}).
		Select(bson.M{"todos": 1}).
		One(&group)

	if err == mgo.ErrNotFound {
		group.Id = bson.NewObjectId()
		group.Title = "root"
		err = db.C("todo_groups").Insert(&group)
		if err != nil {
			return nil, err
		}
	}

	return &group, err
}

func FindTodos(db *mgo.Database) ([]Todo, error) {
	group, err := FindTodoGroup(db)

	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = db.C("todos").
		Find(bson.M{"_id": bson.M{"$in": group.Todos}}).
		All(&todos)

	if err != nil {
		return nil, err
	}

	return todos, err
}

func (todo *Todo) Create(db *mgo.Database) error {
	fmt.Println(todo)
	todo.Id = bson.NewObjectId()
	err := db.C("todos").Insert(&todo)

	if err != nil {
		return err
	}
	return nil
}
