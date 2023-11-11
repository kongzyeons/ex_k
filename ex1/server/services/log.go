package services

import (
	"go_beer/models"
	"go_beer/repository"
	"log"
	"time"
)

type LogSrv interface {
	CreateLog(model models.CreateLogRequst) (err error)
}

type logSrv struct {
	logRepo repository.LogRepo
}

func NewLogSrv(logRepo repository.LogRepo) LogSrv {
	return logSrv{logRepo}
}

func (obj logSrv) CreateLog(model models.CreateLogRequst) (err error) {
	logDB := repository.BearLog{
		FunctionName: model.FunctionName,
		Method:       model.Method,
		Message:      model.Message,
		Status:       model.Status,
		Timestamp:    time.Now(),
	}
	err = obj.logRepo.CreateLog(logDB)
	if err != nil {
		log.Println("error", err)
		return ErrRepoLogCreate
	}
	log.Println("log create success")
	return nil
}
