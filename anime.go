package gojikan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// AnimeResource is a struct of anime resources like Related Anime, Producers,
// Licensors, Studios, and Genres
type AnimeResource struct {
	MalID int    `json:"mal_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// AiredTimeline is a struct of anime's aired timeline like its time and date
type AiredTimeline struct {
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Prop   AiredProp `json:"prop"`
	String string    `json:"string"`
}

// AiredProp is a struct of when anime's aired date started and ended
type AiredProp struct {
	From AiredDate `json:"from"`
	To   AiredDate `json:"to"`
}

// AiredDate is a struct of anime's aired date
type AiredDate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// RelatedAnime is a struct of other anime related to this anime
type RelatedAnime struct {
	Adaptation []AnimeResource `json:"Adaptation"`
	SideStory  []AnimeResource `json:"Side story"`
	Summary    []AnimeResource `json:"Summary"`
}

// Anime is a struct of anime details from MyAnimeList
type Anime struct {
	MalID         int             `json:"mal_id"`
	URL           string          `json:"url"`
	ImageURL      string          `json:"image_url"`
	TrailerURL    string          `json:"trailer_url"`
	Title         string          `json:"title"`
	TitleEnglish  string          `json:"title_english"`
	TitleJapanese string          `json:"title_japanese"`
	TitleSynonyms []string        `json:"title_synonyms"`
	Type          string          `json:"type"`
	Source        string          `json:"source"`
	Episodes      int             `json:"episodes"`
	Status        string          `json:"status"`
	Airing        bool            `json:"airing"`
	Aired         AiredTimeline   `json:"aired"`
	Duration      string          `json:"duration"`
	Rating        string          `json:"rating"`
	Score         float64         `json:"score"`
	ScoredBy      int             `json:"scored_by"`
	Rank          int             `json:"rank"`
	Popularity    int             `json:"popularity"`
	Members       int             `json:"members"`
	Favorites     int             `json:"favorites"`
	Synopsis      string          `json:"synopsis"`
	Premiered     string          `json:"premiered"`
	Broadcast     string          `json:"broadcast"`
	Related       RelatedAnime    `json:"related"`
	Producers     []AnimeResource `json:"producers"`
	Licensors     []AnimeResource `json:"licensors"`
	Studios       []AnimeResource `json:"studios"`
	Genres        []AnimeResource `json:"genres"`
	OpeningThemes []string        `json:"opening_themes"`
	EndingThemes  []string        `json:"ending_themes"`
}

func (ths *jikanClient) GetAnime(id int) (anime Anime, err error) {
	url := fmt.Sprintf("%s/anime/%d", ths.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &anime)
	if err != nil {
		return
	}

	return
}
