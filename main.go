package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
// 	"github.com/Internship1/wsMail/config"
// 	"github.com/aiteung/musik"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/Internship1/wsMail/url"
// 	"github.com/gofiber/fiber/v2"
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
    // Ambil connection string dari env
    mongoSTRING := os.Getenv("MONGOSTRING")
    if mongoSTRING == "" {
        log.Fatal("MONGOSTRING not set")
    }

     // Connect ke MongoDB dengan timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoSTRING))
    if err != nil {
        log.Fatal("MongoDB connection error:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

    log.Println("Connected to MongoDB!")

    // Simple HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Listening on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}