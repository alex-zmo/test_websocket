package server

import (
	"math/rand"
	"test_websocket/model"
)

func GetData() model.DashboardInfo {
	randomData := model.DashboardInfo{
		Count:      rand.Intn(300),
		Enterprise: false,
	}
	return randomData

}
