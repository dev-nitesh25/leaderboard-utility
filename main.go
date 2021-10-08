/*
Package main - Entry Point of API where intialized all prerequisite
*/

package main

import (
	"common"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	cacheDataHost := "localhost"
	cacheDataPort := "6379"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""

	// Connect the Database
	_, err := common.InitDB(dbHost, dbPort, dbUser, dbPass)
	if err != nil {
		log.Panic(err)
	}

	// Connect the Redis Cache server
	err = common.InitCache(cacheDataHost, cacheDataPort)
	if err != nil {
		log.Panic(err)
	}

	router := httprouter.New()
	//router.RedirectTrailingSlash = true
	addRouteHandlers(router)

	fmt.Println("Setup complete. Running API server...")

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "Authorization"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	log.Fatal(http.ListenAndServe(":5000", c.Handler(router)))
}
