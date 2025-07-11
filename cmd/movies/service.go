package movies

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/flags"
	"github.com/FrancoMusolino/film-cli/cmd/utils"
)

const url = "https://api.themoviedb.org/3"

type Header struct {
	key, value string
}

type QueryString struct {
	key, value string
}

type TmdbService struct {
	accessToken string
	lang        flags.Lang
}

func NewMoviesService(lang flags.Lang) MoviesService {
	at, ok := os.LookupEnv("TMDB_API_ACCESS_TOKEN")
	if !ok {
		log.Fatal("Missing Movie API Access Token")
	}

	return TmdbService{
		accessToken: at,
		lang:        lang,
	}
}

func doRequest[T any](method, path string, headers []Header, qs []QueryString) (T, error) {
	var zeroValue T

	url := strings.Join([]string{url, path, "?"}, "")
	for _, query := range qs {
		url += fmt.Sprintf("%s=%s&", query.key, query.value)
	}

	req, _ := http.NewRequest(method, url, nil)

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

func (s TmdbService) GetNowPlayingMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/now_playing", s.withDefaultHeaders(nil), s.withDefaultQuery([]QueryString{{key: "page", value: "6"}}))
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (s TmdbService) GetPopularMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/popular", s.withDefaultHeaders(nil), s.withDefaultQuery([]QueryString{{key: "page", value: "6"}}))
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (s TmdbService) GetTopRatedMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/top_rated", s.withDefaultHeaders(nil), s.withDefaultQuery([]QueryString{{key: "page", value: "6"}}))
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (s TmdbService) GetUpcomingMovies() []Movie {
	data, err := doRequest[ResponseWithPagination[Movie]]("GET", "/movie/upcoming", s.withDefaultHeaders(nil), s.withDefaultQuery([]QueryString{{key: "page", value: "6"}}))
	if err != nil {
		return []Movie{}
	}

	return data.Results
}

func (s TmdbService) GetMovieDetail(id int) (MovieDetail, error) {
	var zeroValue MovieDetail

	data, err := doRequest[MovieDetail]("GET", fmt.Sprintf("/movie/%v", id), s.withDefaultHeaders(nil), s.withDefaultQuery(nil))
	if err != nil {
		return zeroValue, fmt.Errorf("no se ha encontrado a la película con ID %v", id)
	}

	return data, nil
}

func (s *TmdbService) withDefaultHeaders(headers []Header) []Header {
	headers = append(headers, Header{
		key:   "Authorization",
		value: fmt.Sprintf("Bearer %s", s.accessToken),
	})
	return headers
}

func (s *TmdbService) withDefaultQuery(qs []QueryString) []QueryString {
	qs = append(qs, QueryString{key: "language", value: string(s.lang)})
	return qs
}
