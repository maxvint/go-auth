package models

type Todo struct {
	Id    int32
	Title string
}

func (s *Todo) TableName() string {
	return "todos"
}
