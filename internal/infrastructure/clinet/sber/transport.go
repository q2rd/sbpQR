package sber

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
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
