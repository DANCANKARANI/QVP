package images
import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/DANCANKARANI/QVP/controllers/image"
)

func SetImageRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/uploads")
	uploadGroup := auth.Group("/",user.JWTMiddleware)
	uploadGroup.Post("/",image.UploadImagesHandler)
	uploadGroup.Put("/:id",image.UploadImagesHandler)
	uploadGroup.Delete("/:id",image.DeleteImageHandler)
}