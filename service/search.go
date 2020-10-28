package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	tmdb "github.com/TMDB/TMDB_config"
	"github.com/TMDB/view"
)

var t = tmdb.New()

// PopularMovie ...
func PopularMovie() (result view.SearchMovieResult, err error) {
	s := fmt.Sprintf("%smovie/popular?api_key=%s", t.BaseURL, t.APIKey)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// PopularTV ...
func PopularTV() (result view.SearchTVShowsResult, err error) {
	s := fmt.Sprintf("%stv/popular?api_key=%s", t.BaseURL, t.APIKey)
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// SearchMovie ...
func SearchMovie(query string) (result view.SearchMovieResult, err error) {
	s := fmt.Sprintf("%ssearch/movie?api_key=%s&query=%s", t.BaseURL, t.APIKey, url.QueryEscape(query))

	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}

	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

//SearchTVshow ...
func SearchTVshow(query string) (result view.SearchTVShowsResult, err error) {
	s := fmt.Sprintf("%ssearch/tv?api_key=%s&query=%s", t.BaseURL, t.APIKey, url.QueryEscape(query))
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

//SearchCompany ...
func SearchCompany(query string) (result view.SearchCompaniesResult, err error) {
	s := fmt.Sprintf("%ssearch/company?api_key=%s&query=%s", t.BaseURL, t.APIKey, url.QueryEscape(query))
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

//SearchPerson ...
func SearchPerson(query string) (result view.SearchPeopleResult, err error) {
	s := fmt.Sprintf("%ssearch/person?api_key=%s&query=%s", t.BaseURL, t.APIKey, url.QueryEscape(query))
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

//SearchKeyWord ...
func SearchKeyWord(query string) (result view.SearchKeywordResult, err error) {
	s := fmt.Sprintf("%ssearch/keyword?api_key=%s&query=%s", t.BaseURL, t.APIKey, url.QueryEscape(query))
	u, err := url.Parse(s)
	if err != nil {
		return result, err
	}
	body, resp, err := t.GetContent(u)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != 200 {
		return result, errors.New(resp.Status)
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

//GeneratePhotoURL ...
func GeneratePhotoURL(path string) string {
	p := "https://image.tmdb.org/t/p/w500"
	url := p + path
	return url
}
