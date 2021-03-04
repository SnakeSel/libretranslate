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
	uri = "https://libretranslate.com"
)

func Translate(source, sourceLang, targetLang string) (string, error) {
	params := url.Values{}
	params.Set("q", source)
	params.Add("source", sourceLang)
	params.Add("target", targetLang)

	res, err := http.Post(uri+"/translate", "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
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

// Detect the language of the text
// Return: confidence, language, error
func Detect(text string) (float32, string, error) {
	params := url.Values{}
	params.Set("q", text)

	res, err := http.Post(uri+"/detect", "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		return -1, "", err
	}
	defer res.Body.Close()

	// Decode the JSON response
	type Detection struct {
		Confidence float32
		Language   string
	}

	var result []Detection
	if err := json.NewDecoder(res.Body).Decode(&result); err == nil {
		if len(result) == 1 {
			return result[0].Confidence, result[0].Language, nil
		} else {
			return -1, "", fmt.Errorf("unknown number of responses ")
		}
	}

	// If not return, then error
	var result2 interface{}
	if err := json.NewDecoder(res.Body).Decode(&result2); err != nil {
		return -1, "", err
	}
	m := result2.(map[string]interface{})
	if val, ok := m["error"]; ok {
		return -1, "", errors.New(fmt.Sprintf("%v", val))
	}

	return -1, "", errors.New("unknown answer")
}
