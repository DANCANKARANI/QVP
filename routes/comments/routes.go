package comments

import (
	"github.com/DANCANKARANI/QVP/controllers/comment"
	"github.com/DANCANKARANI/QVP/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func SetCommentRoutes(app *fiber.App) {
	// Group routes under /api/v1/dependants
	auth := app.Group("/api/v1/comments")
	commentGroup := auth.Group("/", user.JWTMiddleware)
	commentGroup.Post("/",comment.CreateCommentHandler)
	commentGroup.Patch("/:id",comment.UpdateCommentHandler)
	commentGroup.Get("/",comment.GetEntityCommentHandler)
	commentGroup.Delete("/:id",comment.DeleteCommentHandler)
}