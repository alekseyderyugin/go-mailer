package public_api

import (
	"fmt"
	"go-mailer/letters/app/http/public_api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	fmt.Println("Сервер запущен.")
	log.Fatal(app.Listen(":3000"))
}
