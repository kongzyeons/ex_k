package routes

import (
	"go_admin/controller"
	"go_admin/repository"
	"go_admin/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AdminRouter(app *fiber.App, db *gorm.DB) {
	parentRepo := repository.NewParentRepo(db)
	iconRepo := repository.NewIconRepo(db)
	childRepo := repository.NewChildrenRepo(db)

	adminSrv := services.NewAdminSrv(parentRepo, iconRepo, childRepo)
	adminRest := controller.NewAdminRest(adminSrv)
	app.Get("/admin/getquery", adminRest.AdminGetQuery)
}
