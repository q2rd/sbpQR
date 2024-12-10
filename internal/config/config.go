package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppCredentials    string
	ClientCertificate string
	ClientPrivateKey  string
	ServerCertificate string
	MemberID          string
	SBPMemberID       string
	QrID              string
	TID               string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("error loadig .env file %s", err))
	}
	return &Config{
		AppCredentials:    mustGetEnv("CREDS"),
		ClientCertificate: mustGetEnv("CLIENT_CERT"),
		ClientPrivateKey:  mustGetEnv("CLIENT_KEY"),
		ServerCertificate: mustGetEnv("SERVER_CERT"),
		MemberID:          mustGetEnv("MEMBER_ID"),
		SBPMemberID:       mustGetEnv("SBP_MEMBER_ID"),
		QrID:              mustGetEnv("QR_ID"),
		TID:               mustGetEnv("TID"),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
