package handler

import (
	"html/template"
	"net/http"

	"github.com/TMDB/view"

	"github.com/TMDB/service"
)

var query string
var tpl = template.Must(template.ParseGlob("templates/*gohtml"))

//Dashboard ...
func Dashboard(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet || req.Method == http.MethodPost {
		if req.Method == http.MethodGet {
			tpl.ExecuteTemplate(res, "dashboard.gohtml", nil)
		}
		if req.Method == http.MethodPost {
			queryMethod := req.FormValue("querymethod")
			query = req.FormValue("query")
			switch queryMethod {
			case "person":
				http.Redirect(res, req, "/search/person", http.StatusSeeOther)
				return
			case "movie":
				http.Redirect(res, req, "/search/movie", http.StatusSeeOther)
				return
			case "popular movie":
				http.Redirect(res, req, "/search/movie/popular", http.StatusSeeOther)
				return
			case "popular TV":
				http.Redirect(res, req, "/search/tvshow/popular", http.StatusSeeOther)
				return
			case "TV show":
				http.Redirect(res, req, "/search/tvshow", http.StatusSeeOther)
				return
			case "company":
				http.Redirect(res, req, "/search/company", http.StatusSeeOther)
				return
			case "keyword":
				http.Redirect(res, req, "/search/keyword", http.StatusSeeOther)
				return
			default:
				http.Error(res, "Bad Arguments", http.StatusNotAcceptable)
				return
			}

		}

	} else {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
}

// RoutePopularMovie ...
func RoutePopularMovie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.PopularMovie()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].PosterPath = service.GeneratePhotoURL(reslt.Results[i].PosterPath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "popmovie.gohtml", data)

}

// RoutePopularTV ...
func RoutePopularTV(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.PopularTV()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].PosterPath = service.GeneratePhotoURL(reslt.Results[i].PosterPath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "poptv.gohtml", data)

}

// RouteMovie ...
func RouteMovie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchMovie(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].PosterPath = service.GeneratePhotoURL(reslt.Results[i].PosterPath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "movie.gohtml", data)

}

// RouteTVShow ...
func RouteTVShow(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchTVshow(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].PosterPath = service.GeneratePhotoURL(reslt.Results[i].PosterPath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "poptv.gohtml", data)

}

//RouteCompany ...
func RouteCompany(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchCompany(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].LogoPath = service.GeneratePhotoURL(reslt.Results[i].LogoPath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "company.gohtml", data)
}

//RoutePerson ..
func RoutePerson(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchPerson(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	for i := range reslt.Results {
		reslt.Results[i].ProfilePath = service.GeneratePhotoURL(reslt.Results[i].ProfilePath)

	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "person.gohtml", data)
}

// RouteKnownFor ...
func RouteKnownFor(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchPerson(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	data := []view.KnownFor{}
	for _, v := range reslt.Results {
		for _, i := range v.KnownFor {
			i.PosterPath = service.GeneratePhotoURL(i.PosterPath)
			data = append(data, i)
		}

	}
	tpl.ExecuteTemplate(res, "personknownfor.gohtml", data)
}

// RouteKeyWord ...
func RouteKeyWord(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	reslt, err := service.SearchKeyWord(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusConflict)
		return
	}
	data := reslt.Results
	tpl.ExecuteTemplate(res, "keyword.gohtml", data)
}

func Index(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Bad Request Method", http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
