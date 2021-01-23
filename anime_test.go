package gojikan

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAnimeEndpoints(t *testing.T) {
	Convey("Testing Anime Endpoints Method", t, func() {
		client := NewJikanClient()

		Convey("Testing GetAnime Method", func() {
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
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})
		})

		Convey("Testing GetAnimeCharacterStaff Method", func() {
			animeID := 1
			animeCharStaffStruct := AnimeCharacterStaff{}

			Convey("GetAnimeCharacterStaff should return an AnimeCharacterStaff given valid ID", func() {
				animeCharStaff, err := client.GetAnimeCharacterStaff(animeID)

				So(animeCharStaff, ShouldHaveSameTypeAs, animeCharStaffStruct)
				So(len(animeCharStaff.Characters), ShouldBeGreaterThanOrEqualTo, 0)
				So(len(animeCharStaff.Staff), ShouldBeGreaterThanOrEqualTo, 0)
				So(err, ShouldBeNil)
			})

			Convey("GetAnimeCharacterStaff should return an error given invalid ID", func() {
				animeCharStaff, err := client.GetAnimeCharacterStaff(0)

				So(animeCharStaff, ShouldBeZeroValue)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, ResourceNotFoundError)
			})
		})
	})

}
