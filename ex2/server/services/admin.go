package services

import (
	"fmt"
	"go_admin/models"
	"go_admin/repository"
)

type AdminSrv interface {
	AdminGetQuery(model models.AdminGetQueryRequest) (results []models.AdminGetQueryResponse, err error)
}

type adminSrv struct {
	parentRepo   repository.ParentRepo
	iconRepo     repository.IconRepo
	childrenRepo repository.ChildrenRepo
}

func NewAdminSrv(
	parentRepo repository.ParentRepo,
	iconRepo repository.IconRepo,
	childrenRepo repository.ChildrenRepo,
) AdminSrv {
	return adminSrv{parentRepo, iconRepo, childrenRepo}
}

func (obj adminSrv) AdminGetQuery(model models.AdminGetQueryRequest) (results []models.AdminGetQueryResponse, err error) {
	var isChildren int
	if model.IsChildren {
		isChildren = 1
	}
	queryChilden := fmt.Sprintf("Select icon_id, parent_id, name, route From childrens Where parent_id = '%v'And icon_id = '%v'And is_children = '%v'",
		model.ParentID, model.IconID, isChildren,
	)

	childDB, err := obj.childrenRepo.GetQuery(queryChilden)
	if err != nil {
		return results, ErrRepoChilGetQuery
	}

	if len(childDB) == 0 {
		return results, ErrNotFound
	}
	queryIcon := fmt.Sprintf("Select name From icons Where id = '%v'",
		model.IconID,
	)
	iconDB, err := obj.iconRepo.GetQuery(queryIcon)
	if err != nil {
		return results, ErrRepoIconGetQuery
	}
	if len(iconDB) == 0 {
		return results, ErrNotFound
	}

	queryParent := fmt.Sprintf("Select name, route From parents Where id = '%v'",
		model.ParentID,
	)
	parentDB, err := obj.parentRepo.GetQuery(queryParent)
	if err != nil {
		return results, ErrRepoParentGetQuery
	}
	if len(parentDB) == 0 {
		return results, ErrNotFound
	}

	var resultChildren []models.Children
	for i := range childDB {
		resultChil := models.Children{
			IconID:   childDB[i].IconID,
			ParentID: childDB[i].ParentID,
			Name:     childDB[i].Name,
			Route:    childDB[i].Route,
		}
		resultChildren = append(resultChildren, resultChil)
	}
	result := models.AdminGetQueryResponse{
		ParentID: model.IconID,
		Name:     parentDB[0].Name,
		Route:    parentDB[0].Route,
		Icon:     iconDB[0].Name,
		Children: resultChildren,
	}
	results = append(results, result)

	return results, nil
}
