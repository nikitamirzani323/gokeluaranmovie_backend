package main

import (
	"log"

	"github.com/nikitamirzani323/gosveltemdb/db"
	"github.com/nikitamirzani323/gosveltemdb/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
