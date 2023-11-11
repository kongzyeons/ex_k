package services

import (
	"fmt"
	"go_beer/models"
	"go_beer/repository"
	"go_beer/utils"
	"io"
	"os"
	"strings"
)

type BeerSrv interface {
	CreateBeer(model models.CreateBeerRequest) (err error)
	GetBeer(model models.GetBeerRequest) (result repository.Beer, err error)
	GetPaginatBeer(model models.GetPaginatBeerRequest) (results []repository.Beer, err error)
	DeleteBeer(beer_id int) (err error)
	UpdateBeer(beer_id int, model models.UpdateBeerRequest) (err error)
}

type beerSrv struct {
	beerRepo repository.BeerRepo
}

func NewBeerSrv(beerRepo repository.BeerRepo) BeerSrv {
	return beerSrv{beerRepo}
}

func (obj beerSrv) CreateBeer(model models.CreateBeerRequest) (err error) {
	// handle request
	if model.Name == "" || model.Category == "" || model.Detail == "" {
		return ErrCreateBeerRequest
	}
	// check file name
	file_name := strings.ToLower(model.FileImg.Filename)
	if !utils.StringInSlice(file_name, []string{".png", ".jpg", ".jpeg"}) {
		return ErrCreateBeerRequest
	}

	// test logic
	query := fmt.Sprintf("Select * From beers Where name = '%v'", model.Name)
	beerDB, err := obj.beerRepo.GetQuery(query)
	if err != nil {
		return ErrRepoBeerGet
	}
	if len(beerDB) > 0 {
		return ErrCheckName
	}

	beerNew := repository.Beer{
		Name:     model.Name,
		Category: model.Category,
		Detail:   model.Detail,
	}

	// create
	err = obj.beerRepo.Create(beerNew)
	if err != nil {
		return ErrRepoBeerCreate
	}

	// save file
	// create file name
	savePath := "uploadImg/" + model.Name + ".png"
	out, err := os.Create(savePath)
	if err != nil {
		return ErrcCreateFolder
	}
	defer out.Close()

	// open file
	src, err := model.FileImg.Open()
	if err != nil {
		return ErrOepnFile
	}
	defer src.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return ErrSaveFile
	}
	return nil
}

func (obj beerSrv) GetBeer(model models.GetBeerRequest) (result repository.Beer, err error) {
	query := fmt.Sprintf("Select * From beers Where name = '%v'", model.Name)

	beerDB, err := obj.beerRepo.GetQuery(query)
	if err != nil {
		return result, ErrRepoBeerGet
	}
	if len(beerDB) == 0 {
		return result, ErrNotFound
	}
	result = beerDB[0]
	return result, nil
}

func (obj beerSrv) GetPaginatBeer(model models.GetPaginatBeerRequest) (results []repository.Beer, err error) {
	if model.PageNumber <= 0 || model.PageSize <= 0 {
		return results, ErrGetPaginatBeerRequest
	}
	offset := (model.PageNumber - 1) * model.PageSize
	limit := model.PageSize
	beerDB, err := obj.beerRepo.GetPagination(offset, limit)
	if err != nil {
		return results, ErrRepoBeerGetPaginate
	}
	if len(beerDB) == 0 {
		return results, ErrNotFound
	}
	results = beerDB
	return results, nil
}

func (obj beerSrv) DeleteBeer(beer_id int) (err error) {
	query := fmt.Sprintf("Select * From beers Where id = '%v'", beer_id)
	beerDB, err := obj.beerRepo.GetQuery(query)
	if err != nil {
		return ErrRepoBeerGet
	}
	if len(beerDB) == 0 {
		return ErrNotFound
	}
	err = obj.beerRepo.Delete(beer_id)
	if err != nil {
		return ErrRepoBeerDelete
	}

	//Remove file folder
	filePath := "uploadImg/" + beerDB[0].Name + ".png"
	err = os.Remove(filePath)
	if err != nil {
		return ErrRemoveFile
	}

	return nil
}

func (obj beerSrv) UpdateBeer(beer_id int, model models.UpdateBeerRequest) (err error) {

	// check file name
	if model.FileImg != nil {
		file_name := strings.ToLower(model.FileImg.Filename)
		if !utils.StringInSlice(file_name, []string{".png", ".jpg", ".jpeg"}) {
			return ErrCreateBeerRequest
		}
	}

	query := fmt.Sprintf("Select * From beers Where id = '%v'", beer_id)
	beerDB, err := obj.beerRepo.GetQuery(query)
	if err != nil {
		return ErrRepoBeerGet
	}
	if len(beerDB) == 0 {
		return ErrNotFound
	}

	beerUpdate := repository.Beer{
		ID:       beer_id,
		Name:     model.Name,
		Category: model.Category,
		Detail:   model.Detail,
	}
	if model.Name == "" {
		beerUpdate.Name = beerDB[0].Name
	}
	if model.Category == "" {
		beerUpdate.Category = beerDB[0].Category
	}
	if model.Detail == "" {
		beerUpdate.Detail = beerDB[0].Detail
	}

	err = obj.beerRepo.Update(beerUpdate)
	if err != nil {
		return ErrrepoBeerUpdate
	}

	// update file
	// check update file name
	savePath := "uploadImg/" + beerDB[0].Name + ".png"
	if model.Name != "" {
		oldFilePath := "uploadImg/" + beerDB[0].Name + ".png"
		newFilePath := "uploadImg/" + beerUpdate.Name + ".png"
		err = os.Rename(oldFilePath, newFilePath)
		if err != nil {
			return ErrRenameFile
		}
		savePath = "uploadImg/" + beerUpdate.Name + ".png"
	}
	if model.FileImg != nil {
		// Check if the file already exists
		if _, err := os.Stat(savePath); err == nil {
			// File exists, so delete it before saving the new one
			if err := os.Remove(savePath); err != nil {
				return ErrRemoveFile
			}
		}
		out, err := os.Create(savePath)
		if err != nil {
			return ErrcCreateFolder
		}
		defer out.Close()

		// open file
		src, err := model.FileImg.Open()
		if err != nil {
			return ErrOepnFile
		}
		defer src.Close()

		_, err = io.Copy(out, src)
		if err != nil {
			return ErrSaveFile
		}

	}

	return nil
}
