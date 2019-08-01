package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Todo struct {
	ID        int
	Title     string
	Content   string
	Completed bool
	CreatedAt time.Time
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:password@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("cannot connect to database")
	}
	db.LogMode(true)
	defer db.Close()

	err = db.AutoMigrate(Todo{}).Error
	if err != nil {
		log.Fatalln("cannot create todo table")
	}

	router := gin.Default()
	router.GET("/todolists", listTodos)
	router.POST("/todolist", createTodo)
	router.Run(":8088")
}

func listTodos(c *gin.Context) {
	var todos []Todo
	err := db.Find(&todos).Error
	if err != nil {
		c.String(500, "failed to load all todos")
		return
	}

	c.JSON(200, todos)
}

func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid praram")
		return
	}

	todo := Todo{
		Title: argument.Title,
	}

	err = db.Create(&todo).Error
	if err != nil {
		c.String(500, "failed to create todo")
		return
	}
	c.JSON(200, todo)
}

func updateTodo(c *gin.Context) {
	var argument struct {
		ID    int
		Title string
	}

	err := c.BindQuery(&argument)
	if err != nil {
		c.String(400, "failed to get todo")
	}

	todo := Todo{
		ID:    argument.ID,
		Title: argument.Title,
	}
	err = db.Save(&todo).Error
	if err != nil {
		c.String(500, "failed to update todo")
	}

	c.JSON(200, todo)
}

func deleteTodo(c *gin.Context) {
	var argument struct {
		ID int
	}
	err := c.BindQuery(&argument)
	if err != nil {
		c.String(400, "failed to get todo")
	}

	todo := Todo{
		ID: argument.ID,
	}
	err = db.Delete(&todo).Error
	if err != nil {
		c.String(500, "failed to delete todo")
	}

	c.JSON(200, todo)
}
