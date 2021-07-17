package movie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gusti-andika/movie.go/config"
)

type Movie struct {
	Title  string `json:title`
	Desc   string `json:overview`
	Poster string `json:poster_path`
	Rate   int    `json:vote_average`
	Genres []int  `json:genre_ids`
}

func Search(query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/search/movie?api_key=%s&query=%s", config.BASE_URL, config.API_KEY, query)
	r, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %s", err)
		return nil, err
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var result map[string]interface{}
	if err := decoder.Decode(&result); err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil

}
