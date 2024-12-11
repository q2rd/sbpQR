package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/q2rd/sbpQR/pkg/types"
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
func CreateRequest(requestUID string, data []byte, requestType string, url string, token *types.TokenScopeResponse) (*http.Request, error) {

	req, err := http.NewRequest(
		"POST", url,
		bytes.NewBuffer(data),
	)
	// fmt.Println(req)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("RqUID", requestUID)
	return req, nil
}
