package main

import (
	"log"
	"os"

	"github.com/jfirme-sys/whats-the-ip/app"
)

func main() {
	app := app.GenApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
