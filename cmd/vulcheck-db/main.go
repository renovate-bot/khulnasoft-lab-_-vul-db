package main

import (
	"log"
	"os"

	"github.com/khulnasoft-labs/vulcheck-db/pkg"
)

var (
	version = "0.0.1"
)

func main() {
	ac := pkg.AppConfig{}
	app := ac.NewApp(version)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
