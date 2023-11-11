package routes

import (
	"go_beer/controller"
	"go_beer/repository"
	"go_beer/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func BeerRouter(app *fiber.App, mariaDB *gorm.DB, mongoDB *mongo.Database) {
	beerRepo := repository.NewBeerRepo(mariaDB)
	beerSrv := services.NewBeerSrv(beerRepo)

	logRepo := repository.NewLogRepo(mongoDB)
	logSrv := services.NewLogSrv(logRepo)

	beerRest := controller.NewBeerRest(beerSrv, logSrv)

	app.Post("/beer", beerRest.CreateBeer)
	app.Get("/beer", beerRest.GetBeer)
	app.Get("/beer/pagination", beerRest.GetPaginatBeer)
	app.Delete("/beer/:id", beerRest.DeleteBeer)
	app.Put("/beer/:id", beerRest.UpdateBeer)

}
