package config

import (
	"fmt"
	"strings"
)

const BASE_URL = "https://api.themoviedb.org/3"
const API_KEY = "8d8a280cf9527622879fd0dd6197a4ef"
const BASE_IMAGE_URL = "http://image.tmdb.org/t/p"

func ImageURL(poster string) (string, bool) {
	return fmt.Sprintf("%s/w185/%s", BASE_IMAGE_URL, strings.TrimPrefix(poster, "/")), true
}
