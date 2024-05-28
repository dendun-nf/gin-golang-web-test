package todo

import (
	"github.com/dendun-nf/gin-golang-web-test/helper"
	"gorm.io/gorm"
)

type Todo struct {
	*gorm.Model
	Title string
	Done  bool
}

func (todo *Todo) Update(title string, done bool) {
	todo.Title = helper.GetNotDefault(todo.Title, title)
	todo.Done = helper.GetNotDefault(todo.Done, done)
}
