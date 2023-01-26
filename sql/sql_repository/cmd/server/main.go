package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/go-sql-driver/mysql"

	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/cmd/server/routes"
)

func main() {
	err := godotenv.Load()
	var (
		storageDB *sql.DB
		HOST      = os.Getenv("HOST")
		DB_ADDR   = os.Getenv("DB_ADDR")
		DB_USER   = os.Getenv("DB_USER")
		DB_PWD    = os.Getenv("DB_PWD")
		DB_NAME   = os.Getenv("DB_NAME")
	)

	if err != nil {
		panic(err)
	}
	log.Printf("user:%v, password:%v\n", os.Getenv("${DB_USER}"), os.Getenv("${DB_PWD}"))

	databaseConfig := mysql.Config{
		User:                 DB_USER,
		Passwd:               DB_PWD,
		Addr:                 DB_ADDR,
		DBName:               DB_NAME,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	storageDB, err = sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	if err := storageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")

	server := gin.Default()
	router := routes.NewRouter(server, storageDB)
	router.MapRoutes()

	server.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "pong") })

	err = server.Run(HOST)
	if err != nil {
		log.Println(err)
	}
	log.Printf("server listening on %s", HOST)
}
