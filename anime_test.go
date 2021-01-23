package gojikan

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAnimeEndpoints(t *testing.T) {
	Convey("Testing Anime Endpoints Method", t, func() {
		jikan := NewJikanClient().(*jikanClient)
		animeID := 1

		Convey("Testing GetAnime Method", func() {
			expectedAnime := Anime{
				MalID:         animeID,
				URL:           "https://myanimelist.net/anime/1/Cowboy_Bebop",
				ImageURL:      "https://cdn.myanimelist.net/images/anime/4/19644.jpg",
				TrailerURL:    "https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1",
				Title:         "Cowboy Bebop",
				TitleEnglish:  "Cowboy Bebop",
				TitleJapanese: "カウボーイビバップ",
				Type:          "TV",
				Source:        "Original",
				Episodes:      26,
				Status:        "Finished Airing",
				Airing:        false,
				Aired:         AiredTimeline{},
			}

			expectedAnimeBytes, err := json.Marshal(expectedAnime)
			So(err, ShouldBeNil)

			Convey("GetAnime should return an Anime given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				anime, err := jikan.GetAnime(animeID)

				So(anime, ShouldResemble, expectedAnime)
				So(anime.MalID, ShouldEqual, animeID)
				So(anime.Title, ShouldEqual, expectedAnime.Title)
				So(err, ShouldBeNil)
			})

			Convey("GetAnime should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				anime, err := jikan.GetAnime(animeID)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnime should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				anime, err := jikan.GetAnime(0)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnime should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				anime, err := jikan.GetAnime(0)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Testing GetAnimeCharacterStaff Method", func() {
			expectedAnimeCharStaff := AnimeCharacterStaff{
				Characters: []AnimeCharacter{
					AnimeCharacter{
						MalID:    3,
						URL:      "https://myanimelist.net/character/3/Jet_Black",
						ImageURL: "https://cdn.myanimelist.net/images/characters/11/253723.jpg?s=6c8a19a79a88c46ae15f30e3ef5fd839",
						Name:     "Black, Jet",
						Role:     "Main",
						VoiceActors: []AnimeVoiceActor{
							AnimeVoiceActor{
								MalID:    357,
								Name:     "Ishizuka, Unshou",
								URL:      "https://myanimelist.net/people/357/Unshou_Ishizuka",
								ImageURL: "https://cdn.myanimelist.net/r/42x62/images/voiceactors/2/17135.jpg?s=5925123b8a7cf9b51a445c225442f0ef",
								Language: "Japanese",
							},
						},
					},
				},
			}

			expectedAnimeCharStaffBytes, err := json.Marshal(expectedAnimeCharStaff)
			So(err, ShouldBeNil)

			Convey("GetAnimeCharacterStaff should return an AnimeCharacterStaff given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeCharStaffBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeCharStaff, err := jikan.GetAnimeCharacterStaff(animeID)

				So(animeCharStaff, ShouldResemble, expectedAnimeCharStaff)
				So(len(animeCharStaff.Characters), ShouldBeGreaterThanOrEqualTo, 0)
				So(len(animeCharStaff.Staff), ShouldBeGreaterThanOrEqualTo, 0)
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeCharacterStaff should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				anime, err := jikan.GetAnimeCharacterStaff(animeID)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeCharacterStaff should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				anime, err := jikan.GetAnimeCharacterStaff(0)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeCharacterStaff should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				anime, err := jikan.GetAnimeCharacterStaff(0)

				So(anime, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})
		})

	})

}
