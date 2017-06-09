package main

import (
	"gitlab.com/zcong1993/gost/server"
	"log"
	"os"
)

func init() {
	mysqlConfig := os.Getenv("MYSQL_DB_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if mysqlConfig == "" || jwtSecret == "" {
		log.Fatal("env `MYSQL_DB_URL` and `JWT_SECRET` are required!")
	}
}

func main() {
	s := server.GinEngine()
	s.Run(":8000")
}
