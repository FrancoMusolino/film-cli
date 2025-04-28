package movies

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const url = "https://api.themoviedb.org/3"

type TmdbService struct {
	accessToken string
}

func NewMoviesService() MoviesService {
	at, ok := os.LookupEnv("TMDB_API_ACCESS_TOKEN")
	if !ok {
		log.Fatal("Missing Movie API Access Token")
	}

	return TmdbService{
		accessToken: at,
	}
}

func (ts TmdbService) doRequest(method, path string) ([]byte, error) {
	req, _ := http.NewRequest(method, strings.Join([]string{url, path}, ""), nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ts.accessToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body, nil
}

func (ts TmdbService) GetNowPlayingMovies() []Movie {
	var response BaseResponse[Movie]

	b, err := ts.doRequest("GET", "/movie/now_playing?language=es-AR&page=1")
	if err != nil {
		return response.Results
	}

	if err := json.Unmarshal(b, &response); err != nil {
		log.Fatal(err)
	}

	return response.Results
}

func (ts TmdbService) GetPopularMovies() []Movie {
	var response BaseResponse[Movie]

	b, err := ts.doRequest("GET", "/movie/popular?language=es-AR&page=1")
	if err != nil {
		return response.Results
	}

	if err := json.Unmarshal(b, &response); err != nil {
		log.Fatal(err)
	}

	return response.Results
}

func (ts TmdbService) GetTopRatedMovies() []Movie {
	var response BaseResponse[Movie]

	b, err := ts.doRequest("GET", "/movie/top_rated?language=es-AR&page=1")
	if err != nil {
		return response.Results
	}

	if err := json.Unmarshal(b, &response); err != nil {
		log.Fatal(err)
	}

	return response.Results
}

func (ts TmdbService) GetUpcomingMovies() []Movie {
	var response BaseResponse[Movie]

	b, err := ts.doRequest("GET", "/movie/upcoming?language=es-AR&page=1")
	if err != nil {
		return response.Results
	}

	if err := json.Unmarshal(b, &response); err != nil {
		log.Fatal(err)
	}

	return response.Results
}

type BaseResponse[T any] struct {
	Page    int `json:"page"`
	Results []T `json:"results"`
}
