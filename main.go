package main

import (
	"ClubHouseClone/config"
	"ClubHouseClone/middleware"
	"ClubHouseClone/routes"
	"entgo.io/ent"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	conf := config.New()
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s", conf.Database.Host))

	if err != nil {
		utils.
	}

	app := fiber.New()

	middleware.SetMiddleware(app)

	routes.SetupApiV1(app)

	port := "8000"
	addr := flag.String("addres", port, "http service")
	flag.Parse()

	log.Fatal(app.Listen(":" + *addr))
}
