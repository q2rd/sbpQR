package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppCredentials    string
	ClientCertificate string
	ClientPrivateKey  string
	ServerCertificate string
	MemberID          int
	SBPMemberID       int
	QrID              string
	TID               int
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
		MemberID:          toIntValue(mustGetEnv("SBP_MEMBER_ID")),
		SBPMemberID:       toIntValue(mustGetEnv("MEMBER_ID")),
		QrID:              mustGetEnv("QR_ID"),
		TID:               toIntValue(mustGetEnv("TID")),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func toIntValue(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		panic("recived wrng value type")
	}
	return i
}
