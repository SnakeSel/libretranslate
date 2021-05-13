package libretranslate

import (
	"bytes"
	"encoding/json"

	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// Config is the configuration struct you should pass to New().
type Config struct {
	Url string
	Key string
	// Debug is an optional writer which will be used for debug output.
	Debug io.Writer
}
type Translation struct {
	log *log.Logger
	Config
}

// New returns a new Translation.
func New(conf Config) *Translation {

	tr := new(Translation)

	if conf.Url != "" {
		tr.Url = conf.Url
	} else {
		tr.Url = "https://libretranslate.com"
	}

	tr.Key = conf.Key

	if conf.Debug == nil {
		conf.Debug = ioutil.Discard
	}

	tr.log = log.New(conf.Debug, "[LibreTr]\t", log.LstdFlags)

	return tr
}

// Translate text from a language to another
func (tr *Translation) Translate(source, sourceLang, targetLang string) (string, error) {
	params := url.Values{}
	params.Set("q", source)
	params.Add("source", sourceLang)
	params.Add("target", targetLang)
	if len(tr.Key) > 0 {
		tr.log.Println("add api key to param")
		params.Add("api_key", tr.Key)
	}
	uri, err := url.Parse(tr.Url)
	if err != nil {
		tr.log.Println("Error parse url")
		return "", err
	}

	uri.Path = path.Join(uri.Path, "/translate")
	tr.log.Println(uri.String())

	res, err := http.Post(uri.String(), "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		fmt.Println("Post error")
		return "", err
	}
	defer res.Body.Close()
	tr.log.Printf("Response code: %d", res.StatusCode)

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
		return "", fmt.Errorf("%v", val)

	}

	return "", errors.New("unknown answer")

}

// Detect the language of the text
// Return: confidence, language, error
func (tr *Translation) Detect(text string) (float32, string, error) {
	params := url.Values{}
	params.Set("q", text)
	if len(tr.Key) > 0 {
		tr.log.Println("add api key to param")
		params.Add("api_key", tr.Key)
	}
	uri, err := url.Parse(tr.Url)
	if err != nil {
		tr.log.Println("Error parse url")
		return -1, "", err
	}

	uri.Path = path.Join(uri.Path, "/detect")
	tr.log.Println(uri.String())

	res, err := http.Post(uri.String(), "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
	if err != nil {
		return -1, "", err
	}
	defer res.Body.Close()
	tr.log.Printf("Response code: %d", res.StatusCode)

	switch res.StatusCode {
	case http.StatusOK: //200
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
	//case StatusBadRequest: //400 	Invalid request
	//case StatusTooManyRequests: //429 Slow down
	//case StatusInternalServerError: //500 Detection error
	default:
		var result2 interface{}
		if err := json.NewDecoder(res.Body).Decode(&result2); err != nil {
			return -1, "", err
		}
		m := result2.(map[string]interface{})
		if val, ok := m["error"]; ok {
			return -1, "", fmt.Errorf("%v", val)
		}
	}

	return -1, "", errors.New("unknown answer")
}
