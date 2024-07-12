package images
import (
	"github.com/gofiber/fiber/v2"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/DANCANKARANI/QVP/controllers/image"
)

func SetImageRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/upload/profile")
	//user profile images
	uploadGroup := auth.Group("/",user.JWTMiddleware)
	uploadGroup.Post("/",image.UploadProfileImage)
	uploadGroup.Put("/:id",image.UploadImagesHandler)
	uploadGroup.Delete("/:id",image.DeleteImageHandler)

	//insurance icons
	auth= app.Group("/api/v1/upload/insurance")
	uploadGroup = auth.Group("/",user.JWTMiddleware)
	uploadGroup.Post("/:id",image.UploadInsuranceIcon)

	//insurance icon
	auth= app.Group("/api/v1/upload/payment-methods")
	uploadGroup = auth.Group("/",user.JWTMiddleware)
	uploadGroup.Post("/:id",image.UploadPaymentMethod)
}