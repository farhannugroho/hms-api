package main

import (
	"andhiga.com/dhimasprajaya/go-vue-rs/config"
	"andhiga.com/dhimasprajaya/go-vue-rs/router"
	"andhiga.com/dhimasprajaya/go-vue-rs/util"
	"log"
)

func main() {
	// Setup Config and JWT
	config.Setup()
	util.Setup()

	// Init DB
	config.InitDatabase()

	// Init Router
	r := router.InitRouter()

	// Start Server
	log.Fatal(r.Run())

	// Start Server Port :8000 Development
	//log.Fatal(r.Run(":8000"))
}
