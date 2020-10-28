package view

// SearchMovieResult type represents The Movie Database Movies search results
type SearchMovieResult struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		Popularity       float64 `json:"popularity"`
		VoteCount        int     `json:"vote_count"`
		Video            bool    `json:"video"`
		PosterPath       string  `json:"poster_path"`
		ID               int     `json:"id"`
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIds         []int   `json:"genre_ids"`
		Title            string  `json:"title"`
		VoteAverage      float64 `json:"vote_average"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

// SearchPeopleResult type represents The person Database ppl search results
type SearchPeopleResult struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		Popularity         float64 `json:"popularity"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		ID                 int     `json:"id"`
		ProfilePath        string  `json:"profile_path"`
		Adult              bool    `json:"adult"`
		KnownFor           []struct {
			VoteCount        int     `json:"vote_count"`
			ID               int     `json:"id"`
			Video            bool    `json:"video"`
			MediaType        string  `json:"media_type"`
			VoteAverage      float64 `json:"vote_average"`
			Title            string  `json:"title"`
			ReleaseDate      string  `json:"release_date"`
			OriginalLanguage string  `json:"original_language"`
			OriginalTitle    string  `json:"original_title"`
			GenreIds         []int   `json:"genre_ids"`
			BackdropPath     string  `json:"backdrop_path"`
			Adult            bool    `json:"adult"`
			Overview         string  `json:"overview"`
			PosterPath       string  `json:"poster_path"`
		} `json:"known_for"`
		Gender int `json:"gender"`
	} `json:"results"`
}

//SearchTVShowsResult  represents The TVshows Database TVshows search results
type SearchTVShowsResult struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name"`
		GenreIds         []int    `json:"genre_ids"`
		Name             string   `json:"name"`
		Popularity       float64  `json:"popularity"`
		OriginCountry    []string `json:"origin_country"`
		VoteCount        int      `json:"vote_count"`
		FirstAirDate     string   `json:"first_air_date"`
		BackdropPath     string   `json:"backdrop_path"`
		OriginalLanguage string   `json:"original_language"`
		ID               int      `json:"id"`
		VoteAverage      float64  `json:"vote_average"`
		Overview         string   `json:"overview"`
		PosterPath       string   `json:"poster_path"`
	} `json:"results"`
}

// SearchKeywordResult ..
type SearchKeywordResult struct {
	Page    int `json:"page"`
	Results []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

//SearchCompaniesResult ...
type SearchCompaniesResult struct {
	Page    int `json:"page"`
	Results []struct {
		ID       int    `json:"id"`
		LogoPath string `json:"logo_path"`
		Name     string `json:"name"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// SearchMultiResult type represents The Movie Database multiple media
// search result.
//
// * Search on Movies
// * Search on TV Shows
// * Search on Persons
type SearchMultiResult struct {
	Page    int `json:"page"`
	Results []struct {
		BackdropPath     string   `json:"backdrop_path"`
		ProfilePath      string   `json:"profile_path"`
		FirstAirDate     string   `json:"first_air_date"`
		GenreIds         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name"`
		Title            string   `json:"title"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:"popularity"`
		PosterPath       string   `json:"poster_path"`
		VoteAverage      float64  `json:"vote_average"`
		VoteCount        int      `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

//KnownFor ...
type KnownFor struct {
	VoteCount        int     `json:"vote_count"`
	ID               int     `json:"id"`
	Video            bool    `json:"video"`
	MediaType        string  `json:"media_type"`
	VoteAverage      float64 `json:"vote_average"`
	Title            string  `json:"title"`
	ReleaseDate      string  `json:"release_date"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	GenreIds         []int   `json:"genre_ids"`
	BackdropPath     string  `json:"backdrop_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
}
