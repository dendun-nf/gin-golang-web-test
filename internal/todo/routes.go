package todo

import (
	"github.com/dendun-nf/gin-golang-web-test/internal/commons"
	"github.com/gin-gonic/gin"
	"net/http"
)

const routePath = "/todos"

type Endpoints struct {
	basePath string
}

func NewEndpoints() *Endpoints {
	return &Endpoints{
		basePath: routePath,
	}
}

func (e *Endpoints) MapEndpoint(r *gin.RouterGroup) {
	base := r.Group(e.basePath)
	base.GET("", index)
	base.GET("/:id", getById)
	base.POST("", add)
	base.PATCH("/:id", update)
	base.DELETE("/:id", remove)
}

func index(c *gin.Context) {
	var todos []Model
	if err := commons.GormDb.Find(&todos).Error; err != nil {
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

func getById(c *gin.Context) {
	var todo Model
	tx := commons.GormDb.Where("id = ?", c.Param("id")).First(&todo).Where("deleted_at = ?", nil)

	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": tx.Error,
		})
		return
	}

	if &todo == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Model not found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func add(c *gin.Context) {
	var todoRequests []AddRequestDto
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

	todos := &[]Model{}
	for _, request := range todoRequests {
		t := Model{
			Title: request.Title,
			Done:  request.Done,
		}

		*todos = append(*todos, t)
	}
	commons.GormDb.Create(todos)
	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}

func update(c *gin.Context) {
	var todoTarget Model
	var todoUpdate UpdateRequestDto

	if err := c.BindJSON(&todoUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	db := commons.GormDb.Where("id = ?", todoUpdate.ID).First(&todoTarget).Where("deleted_at = ?", nil)
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

	commons.GormDb.Save(&todoTarget)

	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}

func remove(c *gin.Context) {
	var todoTarget *Model
	db := commons.GormDb.Where("id = ?", c.Param("id")).First(&todoTarget).Where("deleted_at = ?", nil)
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": db.Error,
		})
		return
	}
	if err := commons.GormDb.Delete(&todoTarget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"action": "Executed Successfully",
	})
}
