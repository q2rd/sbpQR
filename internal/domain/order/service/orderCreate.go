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

func CreateOrder(order *entity.OrderCreateReq, cfg *config.Config) error {
	const (
		op       = "CreateOrder"
		scope    = "https://api.sberbank.ru/qr/order.create"
		url      = "https://mc.api.sberbank.ru:443/prod/qr/order/v3/creation"
		currency = "643"
	)
	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return err
	}
	token, client := sber.NewTokenRequest(scope, cfg)
	// fmt.Println(token.AccessToken, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	req, err := utils.CreateRequest(order.RequestUID, jsonOrder, "POST", url, token)
	if err != nil {
		return err
	}
	// fmt.Println(req)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	// fmt.Println(response.StatusCode)

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
