package sber

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"github.com/q2rd/sbpQR/internal/config"
)

// NewTransportWithCert читает сертификатов клиента и сервера создает конфиг для tsl протакола.
// возвращает транспорт с настроенным конфигом.
func NewTransportWithCert(cfg *config.Config) *http.Transport {
	const op = "NewTransportWithCert"

	cert, err := tls.LoadX509KeyPair(cfg.ClientCertificate, cfg.ClientPrivateKey)
	if err != nil {
		log.Fatalf("err loadig client certificates: %v\n op: %s", err, op)
	}

	caCert, err := os.ReadFile(cfg.ServerCertificate)
	if err != nil {
		log.Fatalf("err loading server certificate: %v\n op: %s", err, op)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("error appending certificates to the certificate pool\n op: %s", op)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	return &http.Transport{TLSClientConfig: tlsConfig}
}
