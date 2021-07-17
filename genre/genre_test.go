package genre

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGenreMovies(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	jsonResponse := `{
		"page": 1,
		"results": [
		  {
			"adult": false,
			"backdrop_path": "/dq18nCTTLpy9PmtzZI6Y2yAgdw5.jpg",
			"genre_ids": [
			  28,
			  12,
			  53,
			  878
			],
			"id": 497698,
			"original_language": "en",
			"original_title": "Black Widow",
			"overview": "Natasha Romanoff, also known as Black Widow, confronts the darker parts of her ledger when a dangerous conspiracy with ties to her past arises. Pursued by a force that will stop at nothing to bring her down, Natasha must deal with her history as a spy and the broken relationships left in her wake long before she became an Avenger.",
			"popularity": 9692.351,
			"poster_path": "/qAZ0pzat24kLdO3o8ejmbLxyOac.jpg",
			"release_date": "2021-07-07",
			"title": "Black Widow",
			"video": false,
			"vote_average": 8.1,
			"vote_count": 2446
		  },
		  {
			"adult": false,
			"backdrop_path": "/yizL4cEKsVvl17Wc1mGEIrQtM2F.jpg",
			"genre_ids": [
			  28,
			  878
			],
			"id": 588228,
			"original_language": "en",
			"original_title": "The Tomorrow War",
			"overview": "The world is stunned when a group of time travelers arrive from the year 2051 to deliver an urgent message: Thirty years in the future, mankind is losing a global war against a deadly alien species. The only hope for survival is for soldiers and civilians from the present to be transported to the future and join the fight. Among those recruited is high school teacher and family man Dan Forester. Determined to save the world for his young daughter, Dan teams up with a brilliant scientist and his estranged father in a desperate quest to rewrite the fate of the planet.",
			"popularity": 4299.876,
			"poster_path": "/34nDCQZwaEvsy4CFO5hkGRFDCVU.jpg",
			"release_date": "2021-06-30",
			"title": "The Tomorrow War",
			"video": false,
			"vote_average": 8.3,
			"vote_count": 2405
		  }
		]
	}
	`
	g := Genre{1, "Action"}
	httpmock.RegisterResponder("GET", g.MoviesURL(),
		newResponder(200, jsonResponse, "application/json"))

	movies, _ := g.Movies()
	assert.NotNil(t, movies)
	assert.Equal(t, 2, len(movies))
	assert.Equal(t, 497698, movies[0].Id)
	assert.Equal(t, "Black Widow", movies[0].Title)
	assert.Equal(t, "Natasha Romanoff, also known as Black Widow, confronts the darker parts of her ledger when a dangerous conspiracy with ties to her past arises. Pursued by a force that will stop at nothing to bring her down, Natasha must deal with her history as a spy and the broken relationships left in her wake long before she became an Avenger.", movies[0].Desc)
	assert.Equal(t, "/qAZ0pzat24kLdO3o8ejmbLxyOac.jpg", movies[0].Poster)
}

func newResponder(s int, c string, ct string) httpmock.Responder {
	resp := httpmock.NewStringResponse(s, c)
	resp.Header.Set("Content-Type", ct)
	return httpmock.ResponderFromResponse(resp)
}
