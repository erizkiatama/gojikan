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

		Convey("Testing GetAnimeAllEpisodes Method", func() {
			expectedAnimeAllEpisodes := AnimeEpisodes{
				EpisodesLastPage: 1,
				Episodes: []AnimeEpisode{
					AnimeEpisode{
						EpisodeID:     1,
						Title:         "Asteroid Blues",
						TitleJapanese: "アステロイド・ブルース",
						TitleRomanji:  "Asteroid Blues ",
						Filler:        false,
						Recap:         false,
						VideoURL:      "https://myanimelist.net/anime/1/Cowboy_Bebop/episode/1",
						ForumURL:      "https://myanimelist.net/forum/?topicid=29264",
					},
					AnimeEpisode{
						EpisodeID:     2,
						Title:         "Stray Dog Strut",
						TitleJapanese: "野良犬のストラット",
						TitleRomanji:  "Nora Inu no Strut ",
						Filler:        false,
						Recap:         false,
						VideoURL:      "https://myanimelist.net/anime/1/Cowboy_Bebop/episode/2",
						ForumURL:      "https://myanimelist.net/forum/?topicid=29323",
					},
				},
			}

			expectedAnimeAllEpisodesBytes, err := json.Marshal(expectedAnimeAllEpisodes)
			So(err, ShouldBeNil)

			Convey("GetAnimeAllEpisodes should return an AnimeEpisodes given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeAllEpisodesBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeAllEpisodes, err := jikan.GetAnimeAllEpisodes(animeID, 0)

				So(animeAllEpisodes, ShouldResemble, expectedAnimeAllEpisodes)
				So(animeAllEpisodes.EpisodesLastPage, ShouldEqual, 1)
				So(len(animeAllEpisodes.Episodes), ShouldEqual, 2)
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeAllEpisodes should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeAllEpisodes, err := jikan.GetAnimeAllEpisodes(animeID, 0)

				So(animeAllEpisodes, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeAllEpisodes should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeAllEpisodes, err := jikan.GetAnimeAllEpisodes(0, 0)

				So(animeAllEpisodes, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeAllEpisodes should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeAllEpisodes, err := jikan.GetAnimeAllEpisodes(0, 0)

				So(animeAllEpisodes, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

			Convey("GetAnimeAllEpisodes should return zero AnimeEpisodes given page number 2 when anime's episode less than 100", func() {
				expectedAnimeAllEpisodes.Episodes = []AnimeEpisode{}
				expectedAnimeAllEpisodesBytes, err := json.Marshal(expectedAnimeAllEpisodes)
				So(err, ShouldBeNil)

				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeAllEpisodesBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeAllEpisodes, err := jikan.GetAnimeAllEpisodes(animeID, 2)

				So(animeAllEpisodes, ShouldResemble, expectedAnimeAllEpisodes)
				So(animeAllEpisodes.EpisodesLastPage, ShouldEqual, 1)
				So(len(animeAllEpisodes.Episodes), ShouldEqual, 0)
				So(err, ShouldBeNil)
			})
		})

		Convey("Testing GetAnimeRelatedNews Method", func() {
			expectedAnimeRelatedNews := AnimeNews{
				Articles: []AnimeNewsArticle{
					AnimeNewsArticle{
						URL:        "https://myanimelist.net/news/60609964",
						Title:      "North American Anime & Manga Releases for September",
						AuthorName: "ImperfectBlue",
						AuthorURL:  "https://myanimelist.net/profile/ImperfectBlue",
						ForumURL:   "https://myanimelist.net/forum/?topicid=1862079",
						ImageURL:   "https://cdn.myanimelist.net/s/common/uploaded_files/1598909553-a6f9acc1b6c36cd7b792e5bd67321c13.png?s=3b52b4fe7a2670d33b32d8397d2776bb",
						Comments:   0,
						Intro:      "Here are the North American anime & manga releases for September Week 1: September 1 - 7 Anime Releases Africa no Salaryman (TV) (Africa Salaryman) Complete Coll...",
					},
					AnimeNewsArticle{
						URL:        "https://myanimelist.net/news/56156936",
						Title:      "North American Anime & Manga Releases for November",
						AuthorName: "Sakana-san",
						AuthorURL:  "https://myanimelist.net/profile/Sakana-san",
						ForumURL:   "https://myanimelist.net/forum/?topicid=1749894",
						ImageURL:   "https://cdn.myanimelist.net/s/common/uploaded_files/1541455779-8da9e27ca7b6d6d699bbaec5a537b143.jpeg?s=3ee323c78fd67a75e74f36026e032c33",
						Comments:   9,
						Intro:      "Here are the North American anime & manga releases for November Week 1: November 6 - 12 Anime Releases Black Clover Part 2 Blu-ray & DVD Combo Galaxy Angel Z...",
					},
				},
			}

			expectedAnimeRelatedNewsBytes, err := json.Marshal(expectedAnimeRelatedNews)
			So(err, ShouldBeNil)

			Convey("GetAnimeRelatedNews should return an AnimeNews given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRelatedNewsBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedNews, err := jikan.GetAnimeRelatedNews(animeID)

				So(animeRelatedNews, ShouldResemble, expectedAnimeRelatedNews)
				So(len(animeRelatedNews.Articles), ShouldEqual, 2)
				So(animeRelatedNews.Articles[0].Comments, ShouldEqual, 0)
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRelatedNews should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRelatedNews, err := jikan.GetAnimeRelatedNews(animeID)

				So(animeRelatedNews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRelatedNews should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRelatedNews, err := jikan.GetAnimeRelatedNews(0)

				So(animeRelatedNews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRelatedNews should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedNews, err := jikan.GetAnimeRelatedNews(0)

				So(animeRelatedNews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeRelatedPictures Method", func() {
			expectedAnimeRelatedPictures := AnimePictures{
				Pictures: []AnimePicture{
					AnimePicture{
						Large: "https://cdn.myanimelist.net/images/anime/7/3791l.jpg",
						Small: "https://cdn.myanimelist.net/images/anime/7/3791.jpg",
					},
					AnimePicture{
						Large: "https://cdn.myanimelist.net/images/anime/12/19609l.jpg",
						Small: "https://cdn.myanimelist.net/images/anime/12/19609.jpg",
					},
				},
			}

			expectedAnimeRelatedPicturesBytes, err := json.Marshal(expectedAnimeRelatedPictures)
			So(err, ShouldBeNil)

			Convey("GetAnimeRelatedPictures should return an AnimePictures given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRelatedPicturesBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedPictures, err := jikan.GetAnimeRelatedPictures(animeID)

				So(animeRelatedPictures, ShouldResemble, expectedAnimeRelatedPictures)
				So(len(animeRelatedPictures.Pictures), ShouldEqual, 2)
				So(animeRelatedPictures.Pictures[0].Large, ShouldEqual, "https://cdn.myanimelist.net/images/anime/7/3791l.jpg")
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRelatedPictures should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRelatedPictures, err := jikan.GetAnimeRelatedPictures(animeID)

				So(animeRelatedPictures, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRelatedPictures should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRelatedPictures, err := jikan.GetAnimeRelatedPictures(0)

				So(animeRelatedPictures, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRelatedPictures should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedPictures, err := jikan.GetAnimeRelatedPictures(0)

				So(animeRelatedPictures, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeRelatedVideos Method", func() {
			expectedAnimeRelatedVideos := AnimeVideos{
				Promo: []AnimeVideoPromo{
					AnimeVideoPromo{
						Title:    "PV Blu-ray Box version",
						ImageURL: "https://i.ytimg.com/vi/qig4KOK2R2g/mqdefault.jpg",
						VideoURL: "https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1",
					},
					AnimeVideoPromo{
						Title:    "PV 2",
						ImageURL: "https://i.ytimg.com/vi/QCaEJZqLeTU/mqdefault.jpg",
						VideoURL: "https://www.youtube.com/embed/QCaEJZqLeTU?enablejsapi=1&wmode=opaque&autoplay=1",
					},
					AnimeVideoPromo{
						Title:    "PV 1 English dub version",
						ImageURL: "https://cdn.myanimelist.net/images/icon-banned-youtube.png",
						VideoURL: "https://www.youtube.com/embed/gY5nDXOtv_o?enablejsapi=1&wmode=opaque&autoplay=1",
					},
				},
				Episodes: []AnimeVideoEpisode{
					AnimeVideoEpisode{
						Title:    "The Real Folk Blues (part 2)",
						Episode:  "Episode 26",
						URL:      "https://myanimelist.net/anime/1/Cowboy_Bebop/episode/26",
						ImageURL: "https://cdn.myanimelist.net/images/icon-banned-youtube.png",
					},
					AnimeVideoEpisode{
						Title:    "The Real Folk Blues (part 1)",
						Episode:  "Episode 25",
						URL:      "https://myanimelist.net/anime/1/Cowboy_Bebop/episode/25",
						ImageURL: "https://cdn.myanimelist.net/images/icon-banned-youtube.png",
					},
				},
			}

			expectedAnimeRelatedVideosBytes, err := json.Marshal(expectedAnimeRelatedVideos)
			So(err, ShouldBeNil)

			Convey("GetAnimeRelatedVideos should return AnimeVideos given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRelatedVideosBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedVideos, err := jikan.GetAnimeRelatedVideos(animeID)

				So(animeRelatedVideos, ShouldResemble, expectedAnimeRelatedVideos)
				So(len(animeRelatedVideos.Promo), ShouldEqual, 3)
				So(len(animeRelatedVideos.Episodes), ShouldEqual, 2)
				So(animeRelatedVideos.Promo[0].Title, ShouldEqual, "PV Blu-ray Box version")
				So(animeRelatedVideos.Episodes[0].Episode, ShouldEqual, "Episode 26")
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRelatedVideos should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRelatedVideos, err := jikan.GetAnimeRelatedVideos(animeID)

				So(animeRelatedVideos, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRelatedVideos should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRelatedVideos, err := jikan.GetAnimeRelatedVideos(0)

				So(animeRelatedVideos, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRelatedVideos should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedVideos, err := jikan.GetAnimeRelatedVideos(0)

				So(animeRelatedVideos, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeRelatedStats Method", func() {
			expectedAnimeRelatedStats := AnimeStats{
				Watching:    101786,
				Completed:   694120,
				OnHold:      68627,
				Dropped:     25584,
				PlanToWatch: 314633,
				Total:       1204750,
				Scores: AnimeScores{
					One: AnimeScoreValue{
						Votes:      1528,
						Percentage: 0.2,
					},
					Two: AnimeScoreValue{
						Votes:      724,
						Percentage: 0.1,
					},
					Three: AnimeScoreValue{
						Votes:      1309,
						Percentage: 0.2,
					},
					Four: AnimeScoreValue{
						Votes:      3045,
						Percentage: 0.5,
					},
					Five: AnimeScoreValue{
						Votes:      8499,
						Percentage: 1.4,
					},
					Six: AnimeScoreValue{
						Votes:      19715,
						Percentage: 3.2,
					},
					Seven: AnimeScoreValue{
						Votes:      59291,
						Percentage: 9.6,
					},
					Eight: AnimeScoreValue{
						Votes:      125361,
						Percentage: 20.3,
					},
					Nine: AnimeScoreValue{
						Votes:      174841,
						Percentage: 28.3,
					},
					Ten: AnimeScoreValue{
						Votes:      222688,
						Percentage: 36.1,
					},
				},
			}

			expectedAnimeRelatedStatsBytes, err := json.Marshal(expectedAnimeRelatedStats)
			So(err, ShouldBeNil)

			Convey("GetAnimeRelatedStats should return AnimeStats given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRelatedStatsBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedStats, err := jikan.GetAnimeRelatedStats(animeID)

				So(animeRelatedStats, ShouldResemble, expectedAnimeRelatedStats)
				So(animeRelatedStats.Completed, ShouldEqual, 694120)
				So(animeRelatedStats.Scores, ShouldHaveSameTypeAs, AnimeScores{})
				So(animeRelatedStats.Scores.One.Percentage, ShouldEqual, 0.2)
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRelatedStats should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRelatedStats, err := jikan.GetAnimeRelatedStats(animeID)

				So(animeRelatedStats, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRelatedStats should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRelatedStats, err := jikan.GetAnimeRelatedStats(0)

				So(animeRelatedStats, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRelatedStats should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedStats, err := jikan.GetAnimeRelatedStats(0)

				So(animeRelatedStats, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeRelatedForum Method", func() {
			expectedAnimeRelatedForum := AnimeForum{
				Topics: []AnimeForumTopic{
					AnimeForumTopic{
						TopicID:    24838,
						URL:        "https://myanimelist.net/forum/?topicid=24838",
						Title:      "Cowboy Bebop Episode 26 Discussion",
						AuthorName: "Metty",
						AuthorURL:  "https://myanimelist.net/profile/Metty",
						Replies:    478,
						LastPost: AnimeForumTopicLastPost{
							URL:        "https://myanimelist.net/forum/?topicid=24838&goto=lastpost",
							AuthorName: "YonduOdonta",
							AuthorURL:  "https://myanimelist.net/profile/YonduOdonta",
						},
					},
					AnimeForumTopic{
						TopicID:    40846,
						URL:        "https://myanimelist.net/forum/?topicid=40846",
						Title:      "Cowboy Bebop Episode 20 Discussion",
						AuthorName: "Issun",
						AuthorURL:  "https://myanimelist.net/profile/Issun",
						Replies:    192,
						LastPost: AnimeForumTopicLastPost{
							URL:        "https://myanimelist.net/forum/?topicid=40846&goto=lastpost",
							AuthorName: "bigDreamkiller",
							AuthorURL:  "https://myanimelist.net/profile/bigDreamkiller",
						},
					},
				},
			}

			expectedAnimeRelatedForumBytes, err := json.Marshal(expectedAnimeRelatedForum)
			So(err, ShouldBeNil)

			Convey("GetAnimeRelatedForum should return AnimeForum given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRelatedForumBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedForum, err := jikan.GetAnimeRelatedForum(animeID)

				So(animeRelatedForum, ShouldResemble, expectedAnimeRelatedForum)
				So(len(animeRelatedForum.Topics), ShouldEqual, 2)
				So(animeRelatedForum.Topics[0].Replies, ShouldEqual, 478)
				So(animeRelatedForum.Topics[1].LastPost.AuthorName, ShouldEqual, "bigDreamkiller")
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRelatedForum should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRelatedForum, err := jikan.GetAnimeRelatedForum(animeID)

				So(animeRelatedForum, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRelatedForum should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRelatedForum, err := jikan.GetAnimeRelatedForum(0)

				So(animeRelatedForum, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRelatedForum should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRelatedForum, err := jikan.GetAnimeRelatedForum(0)

				So(animeRelatedForum, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeRecommendations Method", func() {
			expectedAnimeRecommendations := AnimeRecommendations{
				Recommendations: []AnimeRecommendation{
					AnimeRecommendation{
						MalID:               205,
						URL:                 "https://myanimelist.net/anime/205/Samurai_Champloo",
						ImageURL:            "https://cdn.myanimelist.net/images/anime/11/29134.jpg?s=f1f4802d4403c077bc159591f056aee1",
						RecommendationURL:   "https://myanimelist.net/recommendations/anime/1-205",
						Title:               "Samurai Champloo",
						RecommendationCount: 90,
					},
					AnimeRecommendation{
						MalID:               6,
						URL:                 "https://myanimelist.net/anime/6/Trigun",
						ImageURL:            "https://cdn.myanimelist.net/images/anime/7/20310.jpg?s=0a1be11b2b831d3b50747ec526e5f8fd",
						RecommendationURL:   "https://myanimelist.net/recommendations/anime/1-6",
						Title:               "Trigun",
						RecommendationCount: 68,
					},
				},
			}

			expectedAnimeRecommendationsBytes, err := json.Marshal(expectedAnimeRecommendations)
			So(err, ShouldBeNil)

			Convey("GetAnimeRecommendations should return AnimeRecommendations given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeRecommendationsBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRecommendations, err := jikan.GetAnimeRecommendations(animeID)

				So(animeRecommendations, ShouldResemble, expectedAnimeRecommendations)
				So(len(animeRecommendations.Recommendations), ShouldEqual, 2)
				So(animeRecommendations.Recommendations[0].MalID, ShouldEqual, 205)
				So(animeRecommendations.Recommendations[1].Title, ShouldEqual, "Trigun")
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeRecommendations should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeRecommendations, err := jikan.GetAnimeRecommendations(animeID)

				So(animeRecommendations, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeRecommendations should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeRecommendations, err := jikan.GetAnimeRecommendations(0)

				So(animeRecommendations, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeRecommendations should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeRecommendations, err := jikan.GetAnimeRecommendations(0)

				So(animeRecommendations, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

		})

		Convey("Testing GetAnimeReviews Method", func() {
			expectedAnimeReviews := AnimeReviews{
				Reviews: []AnimeReview{
					AnimeReview{
						MalID:        7406,
						URL:          "https://myanimelist.net/reviews.php?id=7406",
						Type:         nil,
						HelpfulCount: 1809,
						Reviewer: AnimeReviewer{
							URL:          "https://myanimelist.net/profile/TheLlama",
							ImageURL:     "https://cdn.myanimelist.net/images/userimages/11081.jpg?t=1600353000",
							Username:     "TheLlama",
							EpisodesSeen: 26,
							Scores: AnimeReviewScore{
								Overall:   10,
								Story:     10,
								Animation: 9,
								Sound:     10,
								Character: 10,
								Enjoyment: 9,
							},
						},
						Content: "This is a first review",
					},
					AnimeReview{
						MalID:        104803,
						URL:          "https://myanimelist.net/reviews.php?id=104803",
						Type:         nil,
						HelpfulCount: 1295,
						Reviewer: AnimeReviewer{
							URL:          "https://myanimelist.net/profile/Polyphemus",
							ImageURL:     "https://cdn.myanimelist.net/images/userimages/1872716.jpg?t=1609398600",
							Username:     "Polyphemus",
							EpisodesSeen: 26,
							Scores: AnimeReviewScore{
								Overall:   7,
								Story:     5,
								Animation: 9,
								Sound:     8,
								Character: 7,
								Enjoyment: 8,
							},
						},
						Content: "This is a second review",
					},
				},
			}

			expectedAnimeReviewsBytes, err := json.Marshal(expectedAnimeReviews)
			So(err, ShouldBeNil)

			Convey("GetAnimeReviews should return an AnimeReviews given valid ID", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeReviewsBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeReviews, err := jikan.GetAnimeReviews(animeID, 0)

				So(animeReviews, ShouldResemble, expectedAnimeReviews)
				So(len(animeReviews.Reviews), ShouldEqual, 2)
				So(animeReviews.Reviews[0].MalID, ShouldEqual, 7406)
				So(animeReviews.Reviews[1].Content, ShouldEqual, "This is a second review")
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeReviews should return error when the API call failed", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return nil, errors.New("Something happened when requesting")
					},
				}

				animeReviews, err := jikan.GetAnimeReviews(animeID, 0)

				So(animeReviews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "Something happened when requesting")
			})

			Convey("GetAnimeReviews should return ResourceNotFoundError given unknown ID", func() {
				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 404,
							Body:       nil,
						}, nil
					},
				}

				animeReviews, err := jikan.GetAnimeReviews(0, 0)

				So(animeReviews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})

			Convey("GetAnimeReviews should return error when unmarshaling unknown data type", func() {
				r := ioutil.NopCloser(bytes.NewReader([]byte("Unknown Data")))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeReviews, err := jikan.GetAnimeReviews(0, 0)

				So(animeReviews, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
			})

			Convey("GetAnimeReviews should return AnimeReviews given valid ID and page number 2", func() {
				r := ioutil.NopCloser(bytes.NewReader(expectedAnimeReviewsBytes))

				jikan.client = &MockClient{
					MockDo: func(*http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: 200,
							Body:       r,
						}, nil
					},
				}

				animeReviews, err := jikan.GetAnimeReviews(animeID, 2)

				So(animeReviews, ShouldResemble, expectedAnimeReviews)
				So(len(animeReviews.Reviews), ShouldEqual, 2)
				So(animeReviews.Reviews[1].MalID, ShouldEqual, 104803)
				So(animeReviews.Reviews[0].Content, ShouldEqual, "This is a first review")
				So(err, ShouldBeNil)
			})
		})
	})
}
