package main

import (
	// "log"
	"fmt"
	"net/http"
	"os"

	// "github.com/Internship1/wsMail/config"

	// "github.com/aiteung/musik"
	// "github.com/gofiber/fiber/v2/middleware/cors"


	// "github.com/Internship1/wsMail/url"

	// "github.com/gofiber/fiber/v2"
)

// func main() {
// 	site := fiber.New(config.Iteung)
// 	site.Use(cors.New(config.Cors))
// 	url.Web(site)
// 	log.Fatal(site.Listen(musik.Dangdut()))
// }

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}