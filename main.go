package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int
	Title  string
	Status string
}

func getTodosHandler(c *gin.Context) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt, _ := db.Prepare("SELECT id, title, status FROM todos")
	rows, _ := stmt.Query()

	todos := []Todo{}
	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

func main() {
	r := gin.Default()

	r.GET("/api/todos", getTodosHandler)

	r.Run(":1234")
}
