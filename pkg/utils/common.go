package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

// ReadJson декодирует тело ответа в данный тип.
func ReadJson(res *http.Response, v any) error {
	if res.Body == nil {
		return fmt.Errorf("response body is nil")
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
}

// GenerateCleanUUID генерирует uuid без разделителя.
func GenerateCleanUUID() string {
	uuidStr := uuid.New().String()
	return strings.ReplaceAll(uuidStr, "-", "")
}

// ToBase64 кодирует client_id и app_secret в base64.
func ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func GenerateTimestamp() string {
	currentTime := time.Now().UTC()
	return currentTime.Format(time.RFC3339)
}
