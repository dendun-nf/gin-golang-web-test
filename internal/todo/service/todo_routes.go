package service

import (
	"github.com/dendun-nf/gin-golang-web-test/database"
	"github.com/dendun-nf/gin-golang-web-test/internal/todo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	var todos *[]todo.Todo
	if err := database.GormDb.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	if &todos == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todos not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})

}

func GetById(c *gin.Context) {
	var todo *todo.Todo
	tx := database.GormDb.Where("id = ?", c.Param("id")).First(&todo).Where("deleted_at = ?", nil)

	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
		})
		return
	}

	if &todo == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func Add(c *gin.Context) {
	var todoRequests *[]todo.AddRequestDto
	if err := c.BindJSON(&todoRequests); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if &todoRequests == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Request couldn't be processed",
		})
		return
	}

	todos := &[]todo.Todo{}
	for _, request := range *todoRequests {
		t := todo.Todo{
			Title: request.Title,
			Done:  request.Done,
		}

		*todos = append(*todos, t)
	}
	database.GormDb.Create(todos)
	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}

func Update(c *gin.Context) {
	var todoTarget todo.Todo
	var todoUpdate *todo.UpdateRequestDto

	if err := c.BindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := database.GormDb.Where("id = ?", todoUpdate.ID).First(&todoTarget).Where("deleted_at = ?", nil)
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": db.Error,
		})
		return
	}

	//database.GormDb.Where("id = ?", todoUpdate.ID).First(&todoTarget)
	todoTarget.Update(
		todoUpdate.Title,
		todoUpdate.Done,
	)

	database.GormDb.Save(&todoTarget)

	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}

func Delete(c *gin.Context) {
	var todoTarget *todo.Todo
	db := database.GormDb.Where("id = ?", c.Param("id")).First(&todoTarget).Where("deleted_at = ?", nil)
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": db.Error,
		})
		return
	}
	if err := database.GormDb.Delete(&todoTarget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}
