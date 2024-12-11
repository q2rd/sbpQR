package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/q2rd/sbpQR/internal/config"
	"github.com/q2rd/sbpQR/internal/domain/order/entity"
	"github.com/q2rd/sbpQR/internal/infrastructure/clinet/sber"
	"github.com/q2rd/sbpQR/pkg/utils"
)

func OrderStatus(order *entity.OrderStatusReq, cfg *config.Config) error {
	const (
		op    = "OrederStatus"
		scope = "https://api.sberbank.ru/qr/order.status"
		url   = "https://mc.api.sberbank.ru:443/prod/qr/order/v3/status"
	)

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}
	token, client := sber.NewTokenRequest(scope, cfg)
	req, err := utils.CreateRequest(order.RequestUID, jsonOrder, "POST", url, token)
	if err != nil {
		return err
	}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
