package gojikan

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAnimeEndpoints(t *testing.T) {
	Convey("Testing Anime Endpoints Method", t, func() {
		client := NewJikanClient()
		animeID := 1
		animeStruct := Anime{}

		Convey("GetAnime should return an Anime given valid ID", func() {
			anime, err := client.GetAnime(animeID)

			So(anime, ShouldHaveSameTypeAs, animeStruct)
			So(anime.MalID, ShouldEqual, animeID)
			So(err, ShouldBeNil)
		})

		Convey("GetAnime should return an error given invalid ID", func() {
			anime, err := client.GetAnime(0)

			So(anime, ShouldBeZeroValue)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Resource does not exist")
		})
	})
}
