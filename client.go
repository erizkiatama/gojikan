package gojikan

// Client is an interface for the Jikan client and responsible
// for all API calls to Jikan API
type Client interface {
	GetAnime(id int) (anime Anime, err error)
}

type jikanClient struct {
	baseURL string
}

// NewJikanClient will return jikanClient that implements Client interface
func NewJikanClient() Client {
	return &jikanClient{
		baseURL: "https://api.jikan.moe/v3",
	}
}
