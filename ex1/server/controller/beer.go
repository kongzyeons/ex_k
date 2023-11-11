package controller

import (
	"go_beer/models"
	"go_beer/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type BeerRest interface {
	CreateBeer(c *fiber.Ctx) error
	GetBeer(c *fiber.Ctx) error
	GetPaginatBeer(c *fiber.Ctx) error
	DeleteBeer(c *fiber.Ctx) error
	UpdateBeer(c *fiber.Ctx) error
}

type beerRest struct {
	beerSrv services.BeerSrv
	logSrv  services.LogSrv
}

func NewBeerRest(beerSrv services.BeerSrv, logSrv services.LogSrv) BeerRest {
	return beerRest{beerSrv, logSrv}
}

func (obj beerRest) CreateBeer(c *fiber.Ctx) error {
	model := models.CreateBeerRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "CreateBeer",
				Method:       "Post",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	// Form image
	if _, err := c.MultipartForm(); err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "CreateBeer",
				Method:       "Post",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	file_image, err := c.FormFile("file_image")
	if err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "CreateBeer",
				Method:       "Post",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	model.FileImg = file_image

	err = obj.beerSrv.CreateBeer(model)
	if err != nil {
		if err == services.ErrCreateBeerRequest {
			log.Println("error", err)
			obj.logSrv.CreateLog(
				models.CreateLogRequst{
					FunctionName: "CreateBeer",
					Method:       "Post",
					Message:      err.Error(),
					Status:       false,
				},
			)
			return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
		}
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "CreateBeer",
				Method:       "Post",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("create success", c.Status(http.StatusCreated))
	obj.logSrv.CreateLog(
		models.CreateLogRequst{
			FunctionName: "CreateBeer",
			Method:       "Post",
			Message:      "create success",
			Status:       true,
		},
	)
	var data []interface{}
	return c.Status(http.StatusCreated).JSON(models.Response(data, http.StatusCreated, "create success"))
}

func (obj beerRest) GetBeer(c *fiber.Ctx) error {
	model := models.GetBeerRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	result, err := obj.beerSrv.GetBeer(model)
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("get success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(models.Response(result, http.StatusOK, "get success"))

}

func (obj beerRest) GetPaginatBeer(c *fiber.Ctx) error {
	model := models.GetPaginatBeerRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	results, err := obj.beerSrv.GetPaginatBeer(model)
	if err != nil {
		if err == services.ErrGetPaginatBeerRequest {
			log.Println("error", err)
			return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
		}
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("get pagination success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(models.Response(results, http.StatusOK, "get pagination success"))
}

func (obj beerRest) DeleteBeer(c *fiber.Ctx) error {
	id := c.Params("id")
	beer_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "DeleteBeer",
				Method:       "Delete",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	err = obj.beerSrv.DeleteBeer(beer_id)
	if err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "DeleteBeer",
				Method:       "Delete",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("delete success", c.Status(http.StatusOK))
	obj.logSrv.CreateLog(
		models.CreateLogRequst{
			FunctionName: "DeleteBeer",
			Method:       "Delete",
			Message:      "delete success",
			Status:       true,
		},
	)
	var data []interface{}
	return c.Status(http.StatusOK).JSON(models.Response(data, http.StatusOK, "delete success"))
}

func (obj beerRest) UpdateBeer(c *fiber.Ctx) error {
	id := c.Params("id")
	beer_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "UpdateBeer",
				Method:       "Put",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}
	model := models.UpdateBeerRequest{}
	if err := c.BodyParser(&model); err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "UpdateBeer",
				Method:       "Put",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusBadRequest).JSON(models.Err_response(err, http.StatusBadRequest))
	}

	// Form image
	if _, err := c.MultipartForm(); err == nil {
		file_image, err := c.FormFile("file_image")
		if err == nil {
			model.FileImg = file_image
		}
	}
	err = obj.beerSrv.UpdateBeer(beer_id, model)
	if err != nil {
		log.Println("error", err)
		obj.logSrv.CreateLog(
			models.CreateLogRequst{
				FunctionName: "UpdateBeer",
				Method:       "Put",
				Message:      err.Error(),
				Status:       false,
			},
		)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err, http.StatusInternalServerError))
	}
	log.Println("update success", c.Status(http.StatusOK))
	obj.logSrv.CreateLog(
		models.CreateLogRequst{
			FunctionName: "UpdateBeer",
			Method:       "Put",
			Message:      "update success",
			Status:       true,
		},
	)
	var data []interface{}
	return c.Status(http.StatusOK).JSON(models.Response(data, http.StatusOK, "update success"))
}
