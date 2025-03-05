package bitwarden

import (
	"encoding/json"
	"github.com/rollicks-c/secretblendproviders/bitwarden"
	"strings"
)

const (
	KeyTOTP = "totp"
	Key2FA  = "2fa"
)

func isTOTP(key string) bool {
	switch key {
	case KeyTOTP:
		return true
	case Key2FA:
		return true
	}
	return false
}

func fuzzySearchKey(exp string) (string, bool) {

	// items
	items := []string{
		"username",
		"password",
		KeyTOTP,
		Key2FA,
	}

	// find candidates
	keyList := []string{}
	exp = strings.ToLower(exp)
	for _, k := range items {
		k = strings.ToLower(k)
		if strings.HasPrefix(k, exp) {
			keyList = append(keyList, k)
		}
	}
	if len(keyList) == 0 {
		return "", false
	}
	if len(keyList) > 1 {
		return "", false
	}

	// match
	key := keyList[0]
	return key, true
}

func lookupKey(item bitwarden.ItemData, exp string) (string, string, bool) {

	// make flat
	raw, err := json.Marshal(item.Login)
	if err != nil {
		panic(err)
	}
	flat := map[string]interface{}{}
	if err := json.Unmarshal(raw, &flat); err != nil {
		panic(err)
	}

	// lookup
	exp = strings.ToLower(exp)
	for k, v := range flat {
		if _, ok := v.(string); !ok {
			continue
		}
		k = strings.ToLower(k)
		if k == exp {
			return k, v.(string), true
		}
	}

	// not found
	return "", "", false
}

func filterOut(exp string, exclude ...string) bool {
	for _, ex := range exclude {
		if exp == ex {
			return true
		}
	}
	return false
}
