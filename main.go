package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	ID       int    `json:"id"`
	AppName  string `json:"appname"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string
}

// / Handlers for authentication of the App
type AuthHandler struct {
	db *sql.DB
	//token Token
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{
		db: db,
	}
}

// Connect to authentication database
func initDB() *sql.DB {
	appDB, err := sql.Open("sqlite3", "./auth.db")
	if err != nil {
		log.Fatal("Could not connect to the app credentials databse :", err)
	}
	return appDB
}

func main() {
	// Initializing the router
	router := gin.Default()

	// Initializing the credentials database connection
	dbConn := initDB()
	defer dbConn.Close()

	// AuthHandler init
	authHandler := NewAuthHandler(dbConn)

	router.POST("/login", authHandler.GetUserFromDB)
	router.Run(":8080")

}

func (h *AuthHandler) GetUserFromDB(c *gin.Context) {
	var app App
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request. App did not bind."})
		return
	}
	fmt.Println("App", app.AppName, app.Password)

	var appFromDB App
	err := h.db.QueryRow("SELECT id, appname, password FROM apps WHERE appname = ?", app.AppName).
		Scan(&appFromDB.ID, &appFromDB.AppName, &appFromDB.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get App from DB"})
		return
	}

	if app.Password != appFromDB.Password { // TODO : replace with hashing
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password do not match"})
		return
	}

	// Generate Token
	fmt.Println("Generating token ! blipbloop")

}
