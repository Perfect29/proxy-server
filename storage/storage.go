package storage

import (
	"sync"
	"github.com/Perfect29/proxy-server/models"
)

var Logs sync.Map

func SaveLog(id string, log models.ProxyLog) {
	Logs.Store(id, log)
}

func GetLog(id string) (models.ProxyLog, bool){
	val, ok := Logs.Load(id)

	if !ok {
		return models.ProxyLog{}, false
	}

	log, ok := val.(models.ProxyLog)

	return log, ok
}

