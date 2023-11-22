package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type TokenInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Typ           string `json:"typ"`
}

func FormatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}
