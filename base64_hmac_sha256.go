package goonung

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func CalculateBase64Hmac256(dataMessage, apiSecret string) (string, error) {
	key := []byte(apiSecret)

	h := hmac.New(sha256.New, key)

	h.Write([]byte(dataMessage))

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
