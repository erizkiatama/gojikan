package gojikan

import "net/http"

// Client is an interface for the Jikan client and responsible
// for all API calls to Jikan API
type Client interface {
	GetAnime(id int) (anime Anime, err error)
	GetAnimeCharacterStaff(id int) (animeCharStaff AnimeCharacterStaff, err error)
	GetAnimeAllEpisodes(id, page int) (animeEpisodes AnimeEpisodes, err error)
	GetAnimeRelatedNews(id int) (animeNews AnimeNews, err error)
	GetAnimeRelatedPictures(id int) (animePictures AnimePictures, err error)
	GetAnimeRelatedVideos(id int) (animeVideos AnimeVideos, err error)
	GetAnimeRelatedStats(id int) (animeStats AnimeStats, err error)
	GetAnimeRelatedForum(id int) (animeForum AnimeForum, err error)
	GetAnimeRecommendations(id int) (animeRecommendations AnimeRecommendations, err error)
}

// HTTPClient is an interface for mocking http library calls
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type jikanClient struct {
	baseURL string
	client  HTTPClient
}

// NewJikanClient will return jikanClient that implements Client interface
func NewJikanClient() Client {
	return &jikanClient{
		baseURL: "https://api.jikan.moe/v3",
		client:  &http.Client{},
	}
}
