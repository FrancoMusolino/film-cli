package movies

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/utils"
)

const url = "https://api.themoviedb.org/3"

type Header struct {
	key, value string
}

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

func (ts TmdbService) getAccessTokenBearerHeader() Header {
	return Header{
		key:   "Authorization",
		value: fmt.Sprintf("Bearer %s", ts.accessToken),
	}
}

func doRequest[T any](method, path string, headers []Header) (T, error) {
	var zeroValue T

	req, _ := http.NewRequest(method, strings.Join([]string{url, path}, ""), nil)

	req.Header.Add("accept", "application/json")
	for _, header := range headers {
		req.Header.Add(header.key, header.value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return zeroValue, err
	}
	defer res.Body.Close()

	data, err := utils.Decode[T](res.Body)
	if err != nil {
		return zeroValue, err
	}

	return data, nil
}

func (ts TmdbService) GetNowPlayingMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/now_playing?language=es-AR&page=1", []Header{ts.getAccessTokenBearerHeader()})
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (ts TmdbService) GetPopularMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/popular?language=es-AR&page=1", []Header{ts.getAccessTokenBearerHeader()})
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (ts TmdbService) GetTopRatedMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/top_rated?language=es-AR&page=1", []Header{ts.getAccessTokenBearerHeader()})
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (ts TmdbService) GetUpcomingMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/upcoming?language=es-AR&page=1", []Header{ts.getAccessTokenBearerHeader()})
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (ts TmdbService) GetMovieDetail(id int) (MovieDetail, error) {
	var zeroValue MovieDetail

	data, err := doRequest[MovieDetail]("GET", fmt.Sprintf("/movie/%v?language=es-AR", id), []Header{ts.getAccessTokenBearerHeader()})
	if err != nil {
		return zeroValue, fmt.Errorf("no se ha encontrado a la película con ID %v", id)
	}

	return data, nil
}
