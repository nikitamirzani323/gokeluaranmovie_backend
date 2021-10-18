package main

import (
	"log"

	"github.com/nikitamirzani323/gokeluaranmovie_backend/db"
	"github.com/nikitamirzani323/gokeluaranmovie_backend/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
