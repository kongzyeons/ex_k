package services_test

import (
	"errors"
	"fmt"
	"go_admin/models"
	"go_admin/repository"
	"go_admin/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminGetQuery(t *testing.T) {
	type testCase struct {
		nameTest            string
		errRepoChilGetQuery error
		childDB             []repository.Children
		errRepoIconGetQuery error
		iconDB              []repository.Icon
		errRepoParentQuery  error
		parentDB            []repository.Parent
		expect              error
	}
	cases := []testCase{
		{nameTest: "test admin get query success", errRepoChilGetQuery: nil, childDB: []repository.Children{{}}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{{}}, errRepoParentQuery: nil, parentDB: []repository.Parent{{}}, expect: nil},
		{nameTest: "test err repo chil get query", errRepoChilGetQuery: errors.New(""), childDB: []repository.Children{{}}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{{}}, errRepoParentQuery: nil, parentDB: []repository.Parent{{}}, expect: services.ErrRepoChilGetQuery},
		{nameTest: "test len chil", errRepoChilGetQuery: nil, childDB: []repository.Children{}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{{}}, errRepoParentQuery: nil, parentDB: []repository.Parent{{}}, expect: services.ErrNotFound},
		{nameTest: "test admin get query success", errRepoChilGetQuery: nil, childDB: []repository.Children{{}}, errRepoIconGetQuery: errors.New(""), iconDB: []repository.Icon{{}}, errRepoParentQuery: nil, parentDB: []repository.Parent{{}}, expect: services.ErrRepoIconGetQuery},
		{nameTest: "test len iconDB", errRepoChilGetQuery: nil, childDB: []repository.Children{{}}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{}, errRepoParentQuery: nil, parentDB: []repository.Parent{{}}, expect: services.ErrNotFound},
		{nameTest: "test err repo parent get query", errRepoChilGetQuery: nil, childDB: []repository.Children{{}}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{{}}, errRepoParentQuery: errors.New(""), parentDB: []repository.Parent{{}}, expect: services.ErrRepoParentGetQuery},
		{nameTest: "test len parentDB", errRepoChilGetQuery: nil, childDB: []repository.Children{{}}, errRepoIconGetQuery: nil, iconDB: []repository.Icon{{}}, errRepoParentQuery: nil, parentDB: []repository.Parent{}, expect: services.ErrNotFound},
	}
	for i := range cases {
		t.Run(cases[i].nameTest, func(t *testing.T) {
			model := models.AdminGetQueryRequest{}

			childrenRepo := repository.NewChildrenRepoMock()
			queryChilden := fmt.Sprintf("Select icon_id, parent_id, name, route From childrens Where parent_id = '%v'And icon_id = '%v'And is_children = '%v'",
				model.ParentID, model.IconID, model.IsChildren,
			)
			childrenRepo.On("GetQuery", queryChilden).Return(
				cases[i].childDB,
				cases[i].errRepoChilGetQuery,
			)

			iconRepo := repository.NewIconRepoMock()
			queryIcon := fmt.Sprintf("Select name From icons Where id = '%v'",
				model.IconID,
			)
			iconRepo.On("GetQuery", queryIcon).Return(
				cases[i].iconDB,
				cases[i].errRepoIconGetQuery,
			)

			parentRepo := repository.NewParentRepoMock()
			queryParent := fmt.Sprintf("Select name, route From parents Where id = '%v'",
				model.ParentID,
			)
			parentRepo.On("GetQuery", queryParent).Return(
				cases[i].parentDB,
				cases[i].errRepoParentQuery,
			)

			adminSrv := services.NewAdminSrv(parentRepo, iconRepo, childrenRepo)
			_, err := adminSrv.AdminGetQuery(model)
			assert.ErrorIs(t, err, cases[i].expect)
		})
	}
}
