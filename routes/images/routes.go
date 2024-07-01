package images
import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers/user"
	"main.go/controllers/image"
)

func SetImageRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/uploads")
	uploadGroup := auth.Group("/",user.JWTMiddleware)
	uploadGroup.Post("/",image.UploadImages)
}