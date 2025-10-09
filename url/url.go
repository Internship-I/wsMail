package url

import (
	"github.com/Internship-I/wsMail/controller"
	"github.com/Internship-I/wsMail/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	// page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
	page.All("/", controller.Sink)

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

	page.Get("/checkip", controller.Homepage)
	page.Get("/transaction", controller.GetAllTransaction)
	page.Post("/insertTransaction", controller.InsertDataTransaction)
	page.Get("/transactionConnote/:connote", controller.GetTransactionByConnote)
	page.Get("/transactionPhone/:phoneNumber", controller.GetTransactionByPhoneNumber)
	page.Get("/transactionAddress/:address", controller.GetTransactionByAddress)

	page.Get("/user", controller.GetAllUser)
	page.Get("/user/:id", controller.GetUserID)
	page.Post("/insertUser", controller.InsertDataUser)
	page.Put("/user/updateUser/:id", controller.UpdateDataUser)
	page.Delete("/user/deleteUser/:id", controller.DeleteUserByID)
	page.Post("/registeruser", handler.Register)
	page.Post("/login", handler.Login)
	page.Post("/loginCust", handler.CustomerLogin)

	page.Get("/docs/*", swagger.HandlerDefault)

	// page.Use(middleware.AuthMiddleware())
	page.Get("/dashboard", handler.DashboardPage)
}
