package shows

import (
	"H2EBack/packages/globals"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var config globals.Config = globals.GetConfig()

type Genre struct {
	Label string `json:"name"`
	Value string `json:"slug"`
}

type TraktIDs struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	IMDB  string `json:"imdb"`
	TMDB  int    `json:"tmdb"`
}

type ActualShow struct {
	Title      string   `json:"title"`
	Year       int      `json:"year"`
	IDs        TraktIDs `json:"ids"`
	Overview   string   `json:"overview"`
	FirstAired string   `json:"first_aired"`
	Runtime    int      `json:"runtime"`
	Trailer    string   `json:"trailer"`
	Status     string   `json:"status"`
	Rating     float64  `json:"rating"`
}
type Show struct {
	Watchers int        `json:"watchers"`
	Show     ActualShow `json:"show"`
}

type SingleImageData struct {
	PosterPath string `json:"poster_path"`
}

type ImageData struct {
	ShowResults []SingleImageData `json:"tv_results"`
}

var genres []Genre
var shows []Show
var imageData ImageData

func addTraktHeadersToRequest(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", config.TRAKT_API_KEY)
}

func addShowFetchParams(params *url.Values, genres, page string) {
	params.Add("limit", "30")
	params.Add("page", page)
	params.Add("genres", genres)
	params.Add("extended", "full")
}

func GetGenres(c *gin.Context) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.trakt.tv/genres/shows", nil)
	addTraktHeadersToRequest(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(bodyBytes, &genres)

	c.IndentedJSON(http.StatusOK, genres)
}

func GetShows(c *gin.Context) {
	client := &http.Client{}
	paramsMap := c.Request.URL.Query()

	base, err := url.Parse("https://api.trakt.tv/shows/" + paramsMap.Get("subLink"))
	if err != nil {
		return
	}
	params := url.Values{}
	addShowFetchParams(&params, paramsMap.Get("genres"), paramsMap.Get("page"))

	base.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
	addTraktHeadersToRequest(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	json.Unmarshal(bodyBytes, &shows)

	c.IndentedJSON(http.StatusOK, shows)
}

func GetImage(c *gin.Context) {
	client := &http.Client{}
	paramsMap := c.Request.URL.Query()

	var url string = fmt.Sprintf("https://api.themoviedb.org/3/find/%s?api_key=%s&language=en-US&external_source=imdb_id", paramsMap.Get("id"), config.TMDB_API_KEY)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(bodyBytes, &imageData)

	c.IndentedJSON(http.StatusOK, imageData)
}
