package audits

// import (
// 	"github.com/DANCANKARANI/QVP/controllers/audit"
// 	"github.com/DANCANKARANI/QVP/controllers/user"
// 	"github.com/gofiber/fiber/v2"
// )

// func SetAuditsRoutes(app *fiber.App) {
// 	// Group routes under /api/v1/auidits
// 	auth := app.Group("/api/v1/audits")
// 	auditGroup := auth.Group("/", user.JWTMiddleware)
// 	auditGroup.Get("/",audit.GetUserAuditsHandler)
// 	auditGroup.Post("/",audit.AddAuditHandler)
// 	auditGroup.Patch("/:id",audit.UpdateAuditHandler)
// 	auditGroup.Delete("/:id",audit.DeleteAuditHandler)
// }