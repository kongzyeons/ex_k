package config

import (
	"fmt"
	"go_admin/repository"
	"log"

	"gorm.io/gorm"
)

func GenDatSimple(db *gorm.DB) {
	parentRepo := repository.NewParentRepo(db)
	parentDB := repository.Parent{
		Name:  "ข้อมูลสรุป",
		Route: "/dashboards",
	}
	err := parentRepo.Create(parentDB)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 11; i++ {
		parentDB := repository.Parent{
			Name:  fmt.Sprintf("parent_name_%v", i),
			Route: fmt.Sprintf("parent_route_%v", i),
		}
		err := parentRepo.Create(parentDB)
		if err != nil {
			log.Fatal(err)
		}
	}

	iconRepo := repository.NewIconRepo(db)

	for i := 1; i < 16; i++ {
		iconDB := repository.Icon{
			Name: fmt.Sprintf("icon_name_%v", i),
		}
		err := iconRepo.Create(iconDB)
		if err != nil {
			log.Fatal(err)
		}
	}
	// id 16
	iconDB := repository.Icon{
		Name: "fi-rr-chart-pie-alt",
	}
	err = iconRepo.Create(iconDB)
	if err != nil {
		log.Fatal(err)
	}

	childRepo := repository.NewChildrenRepo(db)
	childrenDB := repository.Children{
		ParentID:   1,
		IconID:     16,
		Name:       "สินทรัพย์",
		Route:      "/dashboards/assert",
		IsChildren: false,
	}
	err = childRepo.Create(childrenDB)
	if err != nil {
		log.Fatal(err)
	}
	childrenDB = repository.Children{
		ParentID:   1,
		IconID:     16,
		Name:       "สภาพทาง",
		Route:      "/dashboards/condition",
		IsChildren: false,
	}
	err = childRepo.Create(childrenDB)
	if err != nil {
		log.Fatal(err)
	}

	for i := 3; i < 11; i++ {
		childrenDB := repository.Children{
			ParentID:   i,
			IconID:     i,
			Name:       fmt.Sprintf("chil_name_%v", i),
			Route:      fmt.Sprintf("chil_route_%v", i),
			IsChildren: true,
		}
		err = childRepo.Create(childrenDB)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("create data simple success")
}
