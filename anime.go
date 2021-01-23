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

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := ths.client.Do(req)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &anime)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeCharacter is a struct of characters' details in the anime
type AnimeCharacter struct {
	MalID       int               `json:"mal_id"`
	URL         string            `json:"url"`
	ImageURL    string            `json:"image_url"`
	Name        string            `json:"name"`
	Role        string            `json:"role"`
	VoiceActors []AnimeVoiceActor `json:"voice_actors"`
}

// AnimeVoiceActor is a struct of voice actors' details of the anime character
type AnimeVoiceActor struct {
	MalID    int    `json:"mal_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
	Language string `json:"language"`
}

// AnimeStaff is a struct of staff's details responsible for the anime
type AnimeStaff struct {
	MalID     int      `json:"mal_id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	ImageURL  string   `json:"image_url"`
	Positions []string `json:"positions"`
}

// AnimeCharacterStaff is a struct of characters and staffs of the anime
type AnimeCharacterStaff struct {
	Characters []AnimeCharacter `json:"characters"`
	Staff      []AnimeStaff     `json:"staff"`
}

func (ths *jikanClient) GetAnimeCharacterStaff(id int) (animeCharStaff AnimeCharacterStaff, err error) {
	url := fmt.Sprintf("%s/anime/%d/characters_staff", ths.baseURL, id)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := ths.client.Do(req)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &animeCharStaff)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeEpisodes is a struct of all episodes in the anime with pagination
// One page consist of max 100 episodes
type AnimeEpisodes struct {
	EpisodesLastPage int            `json:"episodes_last_page"`
	Episodes         []AnimeEpisode `json:"episodes"`
}

// AnimeEpisode is a struct of episode details in the anime
type AnimeEpisode struct {
	EpisodeID     int       `json:"episode_id"`
	Title         string    `json:"title"`
	TitleJapanese string    `json:"title_japanese"`
	TitleRomanji  string    `json:"title_romanji"`
	Aired         time.Time `json:"aired"`
	Filler        bool      `json:"filler"`
	Recap         bool      `json:"recap"`
	VideoURL      string    `json:"video_url"`
	ForumURL      string    `json:"forum_url"`
}

// GetAnimeAllEpisodes return all anime's episode per page
// Maximum 100 episode per page, if there are more you have to call next page
// Put 0 in page parameter if don't want to use the page
func (ths *jikanClient) GetAnimeAllEpisodes(id, page int) (animeEpisodes AnimeEpisodes, err error) {
	url := fmt.Sprintf("%s/anime/%d/episodes", ths.baseURL, id)
	if page > 0 {
		url = fmt.Sprintf("%s/%d", url, page)
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := ths.client.Do(req)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &animeEpisodes)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeNews is a struct of related news articles of the anime
type AnimeNews struct {
	Articles []AnimeNewsArticle `json:"articles"`
}

// AnimeNewsArticle is a struct of related news article details of the anime
type AnimeNewsArticle struct {
	URL        string    `json:"url"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	AuthorName string    `json:"author_name"`
	AuthorURL  string    `json:"author_url"`
	ForumURL   string    `json:"forum_url"`
	ImageURL   string    `json:"image_url"`
	Comments   int       `json:"comments"`
	Intro      string    `json:"intro"`
}

func (ths *jikanClient) GetAnimeRelatedNews(id int) (animeNews AnimeNews, err error) {
	url := fmt.Sprintf("%s/anime/%d/news", ths.baseURL, id)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := ths.client.Do(req)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &animeNews)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimePictures is a struct of related pictures of the anime
type AnimePictures struct {
	Pictures []AnimePicture `json:"pictures"`
}

// AnimePicture is a struct of large and small picture of the anime
type AnimePicture struct {
	Large string `json:"large"`
	Small string `json:"small"`
}

func (ths *jikanClient) GetAnimeRelatedPictures(id int) (animePictures AnimePictures, err error) {
	url := fmt.Sprintf("%s/anime/%d/pictures", ths.baseURL, id)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	resp, err := ths.client.Do(req)
	if err != nil {
		return
	}

	err = ths.checkStatusError(resp.StatusCode)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &animePictures)
	if err != nil {
		return
	}

	return
}
