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
// test build
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Buat instance Fiber
	site := fiber.New()

	// Aktifkan CORS (bisa custom pakai config.Cors kalau sudah dibuat)
	site.Use(cors.New())

	config.InitConfig()

	if config.MongoString == "" {
		log.Println("[WARNING] MongoString kosong! Periksa environment variable Anda.")
	} else {
		log.Println("[INFO] MongoString terdeteksi:", config.MongoString)
	}

	// Cek koneksi database
	if config.Ulbimongoconn == nil {
		log.Println("[WARNING] MongoDB belum terhubung.")
	} else {
		log.Println("[INFO] MongoDB aktif dan siap digunakan.")
	}

	// Load semua route
	url.Web(site)

	// Jalankan server
	log.Fatal(site.Listen(":" + port))
}
