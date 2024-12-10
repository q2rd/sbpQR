package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/q2rd/sbpQR/internal/config"
	"github.com/q2rd/sbpQR/internal/domain/order/entity"
	"github.com/q2rd/sbpQR/internal/infrastructure/clinet/sber"
	"github.com/q2rd/sbpQR/pkg/utils"
)

const (
	op       = "CreateOrder"
	scope    = "https://api.sberbank.ru/qr/order.create"
	url      = "https://mc.api.sberbank.ru:443/prod/qr/order/v3/creation"
	currency = "643"
)

func CreateOrder(order *entity.Order, cfg *config.Config) error {
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
	fmt.Println(req)
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
func OrderCreateTest() {
	const (
		op       = "CreateOrder"
		scope    = "https://api.sberbank.ru/qr/order.create"
		url      = "https://mc.api.sberbank.ru:443/prod/qr/order/v3/creation"
		currency = "643"
	)
	cfg := config.LoadConfig()
	rqUID := utils.GenerateCleanUUID()
	token, client := sber.NewTokenRequest(scope, cfg)

	productParams := []*OrderParams{
		{
			PositionName:        "water sparkline",
			PositionCount:       2,
			PositionSum:         100,
			PositionDescription: "газированая минеральная вода",
		},
	}
	order := &SBPOrderCreationRequest{
		RequestUID:      rqUID,
		RequestTime:     utils.GenerateTimestamp(),
		MemberID:        cfg.MemberID,
		OrderNumber:     "768656",
		OrderCreateDate: utils.GenerateTimestamp(),
		OrderParamsType: productParams,
		IDQr:            cfg.TID,
		OrderSum:        123,
		Currency:        currency,
		Description:     "Товар не со склада.",
		SbpMemberID:     cfg.SBPMemberID,
	}
	fmt.Println(order.RequestUID, utils.GenerateTimestamp(), order.SbpMemberID)

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		"POST", url,
		bytes.NewBuffer(jsonOrder),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RqUID", rqUID)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

type OrderParams struct {
	PositionName        string `json:"position_name"`
	PositionCount       int    `json:"position_count"`
	PositionSum         int    `json:"position_sum"`
	PositionDescription string `json:"position_description"`
}

type SBPOrderCreationRequest struct {
	RequestUID      string         `json:"rq_uid"`
	RequestTime     string         `json:"rq_tm"`
	MemberID        string         `json:"member_id"`
	OrderNumber     string         `json:"order_number"`
	OrderCreateDate string         `json:"order_create_date"`
	OrderParamsType []*OrderParams `json:"order_params_type"`
	IDQr            string         `json:"id_qr"`
	OrderSum        int            `json:"order_sum"`
	Currency        string         `json:"currency"`
	Description     string         `json:"description"`
	SbpMemberID     string         `json:"sbp_member_id"`
}
