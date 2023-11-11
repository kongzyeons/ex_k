package services_test

import (
	"errors"
	"fmt"
	"go_beer/models"
	"go_beer/repository"
	"go_beer/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatBeer(t *testing.T) {
	type testCase struct {
		nameTest          string
		name              string
		category          string
		detail            string
		errRepoBeerGet    error
		beerDB            []repository.Beer
		errRepoBeerCreate error
		expect            error
	}
	cases := []testCase{
		{nameTest: "test create beer success", name: "name", category: "category", detail: "detail", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerCreate: nil, expect: nil},
		{nameTest: "test err repo ceate", name: "name", category: "category", detail: "detail", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerCreate: errors.New(""), expect: services.ErrRepoBeerCreate},
		{nameTest: "test err repo get", name: "name", category: "category", detail: "detail", errRepoBeerGet: errors.New(""), beerDB: []repository.Beer{}, errRepoBeerCreate: nil, expect: services.ErrRepoBeerGet},
		{nameTest: "test check name", name: "name", category: "category", detail: "detail", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, errRepoBeerCreate: nil, expect: services.ErrCheckName},
		{nameTest: "test check struct name", name: "", category: "category", detail: "detail", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerCreate: nil, expect: services.ErrCreateBeerRequest},
		{nameTest: "test check struct category", name: "name", category: "", detail: "detail", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerCreate: nil, expect: services.ErrCreateBeerRequest},
		{nameTest: "test check struct detail", name: "name", category: "category", detail: "", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerCreate: nil, expect: services.ErrCreateBeerRequest},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			model := models.CreateBeerRequest{
				Name:     cases[i].name,
				Category: cases[i].category,
				Detail:   cases[i].detail,
			}

			beerRepo := repository.NewBeerRepoMock()
			query := fmt.Sprintf("Select * From beers Where name = '%v'", model.Name)
			beerRepo.On("GetQuery", query).Return(
				cases[i].beerDB,
				cases[i].errRepoBeerGet,
			)
			beerNew := repository.Beer{
				Name:     model.Name,
				Category: model.Category,
				Detail:   model.Detail,
				// CreateTime: time.Now(),
				// UpdateTime: time.Now(),
			}
			beerRepo.On("Create", beerNew).Return(
				cases[i].errRepoBeerCreate,
			)

			beerSrv := services.NewBeerSrv(beerRepo)
			err := beerSrv.CreateBeer(model)
			assert.ErrorIs(t, err, cases[i].expect)
		})
	}
}

func TestGetBeer(t *testing.T) {
	type testCase struct {
		nameTest       string
		errRepoBeerGet error
		beerDB         []repository.Beer
		expect         error
	}
	cases := []testCase{
		{nameTest: "test get beer success", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, expect: nil},
		{nameTest: "test err repo beer get", errRepoBeerGet: errors.New(""), beerDB: []repository.Beer{{}}, expect: services.ErrRepoBeerGet},
		{nameTest: "test check len beerDB", errRepoBeerGet: nil, beerDB: []repository.Beer{}, expect: services.ErrNotFound},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			model := models.GetBeerRequest{}

			beerRepo := repository.NewBeerRepoMock()
			query := fmt.Sprintf("Select * From beers Where name = '%v'", model.Name)
			beerRepo.On("GetQuery", query).Return(
				cases[i].beerDB,
				cases[i].errRepoBeerGet,
			)

			beerSrv := services.NewBeerSrv(beerRepo)
			_, err := beerSrv.GetBeer(model)
			assert.ErrorIs(t, err, cases[i].expect)

		})
	}
}

func TestGetPaginatBeer(t *testing.T) {
	type testCase struct {
		nameTest              string
		pageNumber            int
		pageSize              int
		errRepoBeerGetPaginat error
		beerDB                []repository.Beer
		expect                error
	}
	cases := []testCase{
		{nameTest: "test get paginate beer success", pageNumber: 1, pageSize: 1, errRepoBeerGetPaginat: nil, beerDB: []repository.Beer{{}}, expect: nil},
		{nameTest: "test err repo beer get paginate", pageNumber: 1, pageSize: 1, errRepoBeerGetPaginat: errors.New(""), beerDB: []repository.Beer{{}}, expect: services.ErrRepoBeerGetPaginate},
		{nameTest: "test check len beerDB", pageNumber: 1, pageSize: 1, errRepoBeerGetPaginat: nil, beerDB: []repository.Beer{}, expect: services.ErrNotFound},
		{nameTest: "test check struct pageNumber", pageNumber: 0, pageSize: 1, errRepoBeerGetPaginat: nil, beerDB: []repository.Beer{{}}, expect: services.ErrGetPaginatBeerRequest},
		{nameTest: "test check struct pageSize", pageNumber: 1, pageSize: 0, errRepoBeerGetPaginat: nil, beerDB: []repository.Beer{{}}, expect: services.ErrGetPaginatBeerRequest},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			model := models.GetPaginatBeerRequest{
				PageNumber: cases[i].pageNumber,
				PageSize:   cases[i].pageSize,
			}

			beerRepo := repository.NewBeerRepoMock()
			offset := (model.PageNumber - 1) * model.PageSize
			limit := model.PageSize
			beerRepo.On("GetPagination", offset, limit).Return(
				cases[i].beerDB,
				cases[i].errRepoBeerGetPaginat,
			)

			beerSrv := services.NewBeerSrv(beerRepo)
			_, err := beerSrv.GetPaginatBeer(model)
			assert.ErrorIs(t, err, cases[i].expect)
		})
	}
}

func TestDeleteBeer(t *testing.T) {
	type testCase struct {
		nameTest          string
		errRepoBeerGet    error
		beerDB            []repository.Beer
		errRepoBeerDelete error
		expect            error
	}
	cases := []testCase{
		{nameTest: "test delete beer success", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, errRepoBeerDelete: nil, expect: nil},
		{nameTest: "test err repo beer delete", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, errRepoBeerDelete: nil, expect: services.ErrRepoBeerDelete},
		{nameTest: "test err repo beer get", errRepoBeerGet: errors.New(""), beerDB: []repository.Beer{{}}, errRepoBeerDelete: nil, expect: services.ErrRepoBeerGet},
		{nameTest: "test check len beerDB", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerDelete: nil, expect: services.ErrNotFound},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			var beer_id int

			beerRepo := repository.NewBeerRepoMock()
			query := fmt.Sprintf("Select * From beers Where id = '%v'", beer_id)
			beerRepo.On("GetQuery", query).Return(
				cases[i].beerDB,
				cases[i].errRepoBeerGet,
			)

			beerRepo.On("Delete", beer_id).Return(
				cases[i].expect,
			)
			beerSrv := services.NewBeerSrv(beerRepo)
			err := beerSrv.DeleteBeer(beer_id)
			assert.ErrorIs(t, err, cases[i].expect)

		})
	}
}

func TestUpdateBeer(t *testing.T) {
	type testCase struct {
		nameTest          string
		errRepoBeerGet    error
		beerDB            []repository.Beer
		errRepoBeerUpdate error
		expect            error
	}
	cases := []testCase{
		{nameTest: "test update beer success", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, errRepoBeerUpdate: nil, expect: nil},
		{nameTest: "test err repo beer update", errRepoBeerGet: nil, beerDB: []repository.Beer{{}}, errRepoBeerUpdate: errors.New(""), expect: services.ErrrepoBeerUpdate},
		{nameTest: "test err repo beer get ", errRepoBeerGet: errors.New(""), beerDB: []repository.Beer{{}}, errRepoBeerUpdate: nil, expect: services.ErrRepoBeerGet},
		{nameTest: "test check len beerDB", errRepoBeerGet: nil, beerDB: []repository.Beer{}, errRepoBeerUpdate: nil, expect: services.ErrNotFound},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			var beer_id int
			model := models.UpdateBeerRequest{}

			beerRepo := repository.NewBeerRepoMock()
			query := fmt.Sprintf("Select * From beers Where id = '%v'", beer_id)
			beerRepo.On("GetQuery", query).Return(
				cases[i].beerDB,
				cases[i].errRepoBeerGet,
			)
			beerUpdate := repository.Beer{
				ID:       beer_id,
				Name:     model.Name,
				Category: model.Category,
				Detail:   model.Detail,
				// UpdateTime: time.Now(),
			}
			beerRepo.On("Update", beerUpdate).Return(
				cases[i].errRepoBeerUpdate,
			)

			beerSrv := services.NewBeerSrv(beerRepo)
			err := beerSrv.UpdateBeer(beer_id, model)
			assert.ErrorIs(t, err, cases[i].expect)
		})
	}
}
