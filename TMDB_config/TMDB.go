package tmdb

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

// TMDB is a type that implements a client to The Movie Database API
//
// More info : https://www.themoviedb.org/documentation/api
type TMDB struct {
	Client  *http.Client
	APIKey  string
	BaseURL *url.URL
}

// New allocates and initializes a new TMDB.
func New() *TMDB {
	logger, _ := thoth.Init("log")

	apiKey, exist := os.LookupEnv("API_Key")

	if !exist {
		logger.Log(errors.New("API_Key not set in .env"))
		log.Fatal("API_Key not set in .env")
	}

	baseURL, exist := os.LookupEnv("Base_URL")

	if !exist {
		logger.Log(errors.New("Base_URL not set in .env"))
		log.Fatal("Base_URL not set in .env")
	}

	u, _ := url.Parse(baseURL)

	return &TMDB{Client: http.DefaultClient, APIKey: apiKey, BaseURL: u}

}

// GetContent ...
func (tmdb *TMDB) GetContent(u *url.URL) (body []byte, res *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Accept", "application/json")
	res, err = tmdb.Client.Do(req)
	if err != nil {
		return body, res, err
	}
	if res != nil {
		defer res.Body.Close()
	}
	req.Close = true
	body, err = ioutil.ReadAll(res.Body)
	return body, res, err
}
