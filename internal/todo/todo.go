package todo

import (
	"github.com/dendun-nf/gin-golang-web-test/internal/helper"
	"gorm.io/gorm"
)

type Model struct {
	*gorm.Model
	Title string
	Done  bool
}

func (*Model) TableName() string {
	return "todo"
}

func (todo *Model) Update(title string, done bool) {
	todo.Title = helper.GetNotDefault(todo.Title, title)
	todo.Done = helper.GetNotDefault(todo.Done, done)
}
