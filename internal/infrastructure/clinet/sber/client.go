package sber

import (
	"fmt"
	"github.com/q2rd/sbpQR/pkg/types"
	"github.com/q2rd/sbpQR/pkg/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// NewTokenRequest делает запрос для получения токена на указаную ручку. возвращает структуру bearer токен
// и клинта с установленными сертификатами.
func NewTokenRequest(scope string) (*types.TokenScopeResponse, *http.Client) {
	const op = "NewTokenRequest"
	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("scope", scope)

	req, err := http.NewRequest(
		"POST", "https://mc.api.sberbank.ru/prod/tokens/v3/oauth",
		strings.NewReader(form.Encode()),
	)

	if err != nil {
		log.Fatalf("ошибка создания запроса: %s\n op: %s", err, op)
	}

	authCreds := fmt.Sprintf("Basic %s", utils.ToBase64(os.Getenv("CREDS")))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authCreds)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("rquid", utils.GenerateCleanUUID())

	client := &http.Client{Transport: NewTransportWithCert()}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error making request", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("error closing body")
		}
	}(resp.Body)
	tokenResponse := new(types.TokenScopeResponse)
	err = utils.ReadJson(resp, tokenResponse)
	if err != nil {
		log.Fatal("error reading respons: ", err)
	}
	return tokenResponse, client
}
