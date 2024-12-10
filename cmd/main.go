package main

import (
	// "encoding/json"

	"github.com/q2rd/sbpQR/internal/domain/order/service"
)

func main() {
	// cfg := config.LoadConfig()
	// fmt.Println(cfg.QrID)
	// mockOrder := entity.GetMockOrder(cfg)
	//
	// jsonData, err := json.MarshalIndent(mockOrder, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(jsonData))
	// errr := service.CreateOrder(mockOrder, cfg)
	// if err != nil {
	// 	log.Fatal(errr)
	// }
	service.OrderCreateTest()
}
