package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/q2rd/sbpQR/app/types"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// NewTransportWithCert читает сертификатов клиента и сервера создает конфиг для tsl протакола.
// возвращает транспорт с настроенным конфигом.
func NewTransportWithCert() *http.Transport {
	const op = "NewTransportWithCert"
	// чтнение сертификата + ключа
	cert, err := tls.LoadX509KeyPair(os.Getenv("CLIENT_CERT"), os.Getenv("CLIENT_KEY"))
	if err != nil {
		log.Fatalf("ошибка загрузки сертификатов: %v\n op: %s", err, op)
	}

	// чтение сертификата с сервера сбербанка
	caCert, err := os.ReadFile(os.Getenv("SERVER_CERT"))
	if err != nil {
		log.Fatalf("ошибка чтения сертификата минцифры: %v\n op: %s", err, op)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("ошибка добовления сертификата минцифры в пул доверенных\n op: %s", op)
	}

	// TLS клиент с сертификатами
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return &http.Transport{TLSClientConfig: tlsConfig}
}

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

	authCreds := fmt.Sprintf("Basic %s", ToBase64(os.Getenv("CREDS")))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", authCreds)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("rquid", GenerateCleanUUID())

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
	err = ReadJson(resp, tokenResponse)
	if err != nil {
		log.Fatal("error reading respons: ", err)
	}
	return tokenResponse, client
}
