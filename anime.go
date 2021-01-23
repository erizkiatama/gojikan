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

// ===================================================================================================================================

// AnimeVideos is a struct of related promotional and episodes videos of the anime
type AnimeVideos struct {
	RequestHash        string              `json:"request_hash"`
	RequestCached      bool                `json:"request_cached"`
	RequestCacheExpiry int                 `json:"request_cache_expiry"`
	Promo              []AnimeVideoPromo   `json:"promo"`
	Episodes           []AnimeVideoEpisode `json:"episodes"`
}

// AnimeVideoPromo is a struct of details of related promotional video of the anime
type AnimeVideoPromo struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	VideoURL string `json:"video_url"`
}

// AnimeVideoEpisode is a struct of details of related episode video of the anime
type AnimeVideoEpisode struct {
	Title    string `json:"title"`
	Episode  string `json:"episode"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

func (ths *jikanClient) GetAnimeRelatedVideos(id int) (animeVideos AnimeVideos, err error) {
	url := fmt.Sprintf("%s/anime/%d/videos", ths.baseURL, id)

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

	err = json.Unmarshal(body, &animeVideos)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeStats is a struct of related stats of the anime
type AnimeStats struct {
	Watching    int         `json:"watching"`
	Completed   int         `json:"completed"`
	OnHold      int         `json:"on_hold"`
	Dropped     int         `json:"dropped"`
	PlanToWatch int         `json:"plan_to_watch"`
	Total       int         `json:"total"`
	Scores      AnimeScores `json:"scores"`
}

// AnimeScoreValue is a struct details of the scores given
type AnimeScoreValue struct {
	Votes      int     `json:"votes"`
	Percentage float64 `json:"percentage"`
}

// AnimeScores is a struct of anime's score stats from one to ten
type AnimeScores struct {
	One   AnimeScoreValue `json:"1"`
	Two   AnimeScoreValue `json:"2"`
	Three AnimeScoreValue `json:"3"`
	Four  AnimeScoreValue `json:"4"`
	Five  AnimeScoreValue `json:"5"`
	Six   AnimeScoreValue `json:"6"`
	Seven AnimeScoreValue `json:"7"`
	Eight AnimeScoreValue `json:"8"`
	Nine  AnimeScoreValue `json:"9"`
	Ten   AnimeScoreValue `json:"10"`
}

func (ths *jikanClient) GetAnimeRelatedStats(id int) (animeStats AnimeStats, err error) {
	url := fmt.Sprintf("%s/anime/%d/stats", ths.baseURL, id)

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

	err = json.Unmarshal(body, &animeStats)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeForum is a struct of related forum topics of the anime
type AnimeForum struct {
	Topics []AnimeForumTopic `json:"topics"`
}

// AnimeForumTopicLastPost is a struct details of last post of an anime forum topic
type AnimeForumTopicLastPost struct {
	URL        string    `json:"url"`
	AuthorName string    `json:"author_name"`
	AuthorURL  string    `json:"author_url"`
	DatePosted time.Time `json:"date_posted"`
}

// AnimeForumTopic is a struct details of of an anime forum topic
type AnimeForumTopic struct {
	TopicID    int                     `json:"topic_id"`
	URL        string                  `json:"url"`
	Title      string                  `json:"title"`
	DatePosted time.Time               `json:"date_posted"`
	AuthorName string                  `json:"author_name"`
	AuthorURL  string                  `json:"author_url"`
	Replies    int                     `json:"replies"`
	LastPost   AnimeForumTopicLastPost `json:"last_post"`
}

func (ths *jikanClient) GetAnimeRelatedForum(id int) (animeForum AnimeForum, err error) {
	url := fmt.Sprintf("%s/anime/%d/stats", ths.baseURL, id)

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

	err = json.Unmarshal(body, &animeForum)
	if err != nil {
		return
	}

	return
}

// ===================================================================================================================================

// AnimeRecommendations is a struct list of recommendations for the related anime
type AnimeRecommendations struct {
	Recommendations []AnimeRecommendation `json:"recommendations"`
}

// AnimeRecommendation is a struct details of recommendation for the related anime
type AnimeRecommendation struct {
	MalID               int    `json:"mal_id"`
	URL                 string `json:"url"`
	ImageURL            string `json:"image_url"`
	RecommendationURL   string `json:"recommendation_url"`
	Title               string `json:"title"`
	RecommendationCount int    `json:"recommendation_count"`
}

func (ths *jikanClient) GetAnimeRecommendations(id int) (animeRecommendations AnimeRecommendations, err error) {
	url := fmt.Sprintf("%s/anime/%d/recommendations", ths.baseURL, id)

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

	err = json.Unmarshal(body, &animeRecommendations)
	if err != nil {
		return
	}

	return
}
