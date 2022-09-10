package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GitCollabCode/profiles/data"
	"github.com/GitCollabCode/profiles/handlers"
	"github.com/GitCollabCode/profiles/models"
	"github.com/GitCollabCode/profiles/router"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var docs = flag.Bool("docs", false, "Expose documentation endpoint")

func main() {
	flag.Parse()
	log := logrus.New() //customize logger later

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		return
	}

	err = godotenv.Load(filepath.Join(dir, ".env"))
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := data.InitMongoDBDriver(log)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Client.Disconnect(context.TODO())

	v := models.NewValidtor()

	p := handlers.NewProfiles(log, db, v)

	r := router.InitRouter(p)

	if *docs {
		fmt.Println("add swagger docs endpoint here")
	}

	http.ListenAndServe(":3000", r)

	//use graceful down here
}
