package main

import (
	models "holamundo/Models"
	"holamundo/database"
	"holamundo/handlers"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/contacts", getContacts)
	router.POST("/contact", createContact)

	router.Run("localhost:8081")
}

func getContacts(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	contacts := handlers.MyList(db)
	c.IndentedJSON(http.StatusOK, contacts)

	defer db.Close()
}

func createContact(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	var contact models.Contact
	c.BindJSON(&contact)
	handlers.CreateContact(db, contact)
	c.IndentedJSON(http.StatusCreated, contact)
	defer db.Close()
}
