package main

import (
	"fmt"
	"github.com/q2rd/sbpQR/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg.QrID)
}
