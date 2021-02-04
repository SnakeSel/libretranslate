package libretranslate

import (
	"bytes"
	"encoding/json"

	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	uri = "https://libretranslate.com/translate"
)

func Translate(source, sourceLang, targetLang string) (string, error) {
	params := url.Values{}
	params.Set("q", source)
	params.Add("source", sourceLang)
	params.Add("target", targetLang)

	res, err := http.Post(uri, "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Decode the JSON response
	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}

	m := result.(map[string]interface{})
	if val, ok := m["translatedText"]; ok {
		return fmt.Sprintf("%v", val), nil
	}

	if val, ok := m["error"]; ok {
		return "", errors.New(fmt.Sprintf("%v", val))
	}

	return "", errors.New("unknown answer")

}
