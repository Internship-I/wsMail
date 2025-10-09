package main

import (
	"log"
	"os"

	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Internship-I/wsMail/config"
	"github.com/Internship-I/wsMail/url"

	// "github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintln(w, "Hello, World!")
// }

// @title TES SWAGGER MAIL APP
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Nidasakinaa
// @contact.email 714220040@std.ulbi.ac.id

// @host https://mailbe-3edd125fb8b1.herokuapp.com
// @BasePath /
// @schemes https http
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // fallback untuk local
	}

	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)

	log.Fatal(site.Listen(":" + port))
}