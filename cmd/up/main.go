package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var filename = "up.json"

func main() {
	var force bool
	flag.BoolVar(&force, "f", false, "force")
	flag.Parse()

	mongoUri := os.Getenv("MONGOURL")
	if mongoUri == "" {
		log.Fatal("env `MONGOURL` is required!!!")
	}

	if _, err := os.Stat(filename); err == nil && !force {
		log.Fatal("up.json already exists")
	}

	data := fmt.Sprintf(upTpl, mongoUri)

	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

var upTpl = `{
  "name": "gost",
  "profile": "cong",
  "regions": [
    "ap-southeast-1"
  ],
  "environment": {
    "MONGOURL": "%s"
  }
}
`
