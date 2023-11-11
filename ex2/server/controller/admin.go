package controller

import (
	"go_admin/models"
	"go_admin/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AdminRest interface {
	AdminGetQuery(c *fiber.Ctx) error
}

type adminRest struct {
	adminSrv services.AdminSrv
}

func NewAdminRest(adminSrv services.AdminSrv) AdminRest {
	return adminRest{adminSrv}
}

func (obj adminRest) AdminGetQuery(c *fiber.Ctx) error {
	model := models.AdminGetQueryRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	results, err := obj.adminSrv.AdminGetQuery(model)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("get success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(models.Response(results, http.StatusOK, "get success"))

}
