package quote_details

import (
	"github.com/DANCANKARANI/QVP/controllers/quote_detail"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)
func SetQuoteDetailsRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/quote-details")
	quoteDetailGroup := auth.Group("/",user.JWTMiddleware)
	quoteDetailGroup.Post("/",quote_detail.AddQuoteDetailHandler)
	quoteDetailGroup.Get("/:id", quote_detail.GetQuoteDetailWithPrescriptionHandler)
	quoteDetailGroup.Patch("/:id", quote_detail.UpdateQuoteDetailHandler)
	quoteDetailGroup.Delete("/:id", quote_detail.DeleteQuoteDetailHandler)
}